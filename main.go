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
	// •	2 types of users (buyers, and sellers) -done

	// •	User sellers can list his merchs -done
	// •	User seller can update his merchs quantity -done

	// •	User buyers can see list of merchs done
	// •	User buyers can make a purchase

	http.HandleFunc("/merchs", merch.MerchHnadler)
	http.HandleFunc("/merchs/getMerchsByUserId", merch.GetMerchsByUserId)
	http.HandleFunc("/merchs/updateMerch/", merch.UpdateMerchById) // idk how to use it without third party app :D /merchs/updateMerch/{merch_id}
	http.HandleFunc("/merchs/purchase", user.MakePurchase)
	http.ListenAndServe(":8080", nil)
}
