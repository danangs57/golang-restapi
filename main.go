package main

import (
	"database/sql"
	"net/http"

	merch "test/merch"
	user "test/user"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var err error

	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/ecommerce")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// •	2 types of users (buyers, and sellers)

	// •	User sellers can list his merchs
	// •	User seller can update his merchs quantity

	// •	User buyers can see list of merchs
	// •	User buyers can make a purchase

	// these may not the best approaches but trust me ill become so handy in short amount of time

	http.HandleFunc("/merchs", merch.MerchHnadler)
	http.HandleFunc("/merchs/getMerchsByUserId", merch.GetMerchsByUserId)
	http.HandleFunc("/merchs/updateMerch/", merch.UpdateMerchById) // idk how to use it without third party app :D /merchs/updateMerch/{merch_id}
	http.HandleFunc("/merchs/purchase", user.MakePurchase)         // goroutine implenetation here
	http.ListenAndServe(":8080", nil)
}
