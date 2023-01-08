package merch

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"test/conf"
)

type Merch struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
	Seller   int    `json:"seller"`
}

func MerchHnadler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getMerchs(w, r)
	case "POST":
		createMerch(w, r)
	case "PUT":
		UpdateMerchById(w, r)
	default:
		http.Error(w, http.StatusText(405), 405)
	}
}

func UpdateMerchById(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		http.Error(w, r.Method, 405)
	}
	rawPath := r.URL.Path

	type UpdateMerchRequest struct {
		Name     string `json:"name"`
		Quantity int    `json:"quantity"`
	}

	segments := strings.Split(rawPath, "/")

	merch_id := segments[len(segments)-1]

	var req UpdateMerchRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if merch_id == "" || req.Name == "" || req.Quantity <= 0 {
		http.Error(w, "parameters not satisfied ", http.StatusBadRequest)
		return
	}

	_, err = conf.Db.Exec(`UPDATE merchs
		SET name = ?, quantity = ?
		WHERE id = ?`,
		req.Name, req.Quantity,
		merch_id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Updated!")
}

func GetMerchsByUserId(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, http.StatusText(405), 405)
	}
	query := r.URL.Query()
	q := query.Get("seller_id")
	if q == "" {
		http.Error(w, "seller_id is required!", http.StatusBadRequest)
		return
	}

	rows, err := conf.Db.Query("SELECT * FROM merchs WHERE SELLER = ?", q)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	merchs := []Merch{}

	for rows.Next() {
		var p Merch
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Seller)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		merchs = append(merchs, p)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merchs)

}

func createMerch(w http.ResponseWriter, r *http.Request) {

	var p Merch
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, "invalid body", http.StatusBadRequest)
		return
	}

	if p.Name == "" || p.Quantity <= 0 || p.Seller <= 0 {
		http.Error(w, "Invalid merch request body", http.StatusBadRequest)
		return
	}

	_, err = conf.Db.Exec("INSERT INTO merchs (name, quantity, seller) VALUES (?, ?, ?)", p.Name, p.Quantity, p.Seller)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	io.WriteString(w, "Created!")
}

func getMerchs(w http.ResponseWriter, r *http.Request) {

	rows, err := conf.Db.Query("SELECT id, name, quantity, seller FROM merchs")
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	defer rows.Close()

	merchs := []Merch{}

	for rows.Next() {
		var p Merch
		err := rows.Scan(&p.ID, &p.Name, &p.Quantity, &p.Seller)
		if err != nil {
			http.Error(w, http.StatusText(500), 500)
			return
		}
		merchs = append(merchs, p)
	}

	if err = rows.Err(); err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(merchs)

}
