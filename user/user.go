package user

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"test/conf"
)

type PurchaseDetail struct {
	MerchId    int `json:"merch_id"`
	Quantity   int `json:"quantity"`
	Status     int `json:"status"`
	PurchaseId int `json:"purchase_id"`
}

type PurchaseData struct {
	SellerId       int              `json:"seller_id"`
	BuyerId        int              `json:"buyer_id"`
	PurchaseDetail []PurchaseDetail `json:"purchase_data"`
}

func MakePurchase(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, http.StatusText(405), 405)
	}
	var p PurchaseData
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = json.Unmarshal(body, &p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	res, err := conf.Db.Exec("INSERT INTO purchases (buyer_id, seller_id) VALUES (?, ?)", p.BuyerId, p.SellerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	lastID, err := res.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// making purchases and decrease the item quantity as it go
	for _, data := range p.PurchaseDetail {

		var quantity int
		err := conf.Db.QueryRow("SELECT quantity from merchs where id = ? ", data.MerchId).Scan(&quantity)
		if err != nil {
			http.Error(w, "invalid merch data", http.StatusInternalServerError)
		}

		// merch out of quantity
		if quantity <= 0 {
			continue
		}

		conf.Db.Exec("INSERT INTO user_purchase (merch_id, quantity, status, purchase_id) VALUES (?, ? ,? ,?)", data.MerchId, data.Quantity, data.Status, lastID)

		// update stock/quantity
		go func(data PurchaseDetail) {
			conf.Db.Exec("UPDATE merchs SET quantity = quantity - ? WHERE id = ?", data.Quantity, data.MerchId)
		}(data)

	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Purchase Created!")

}
