package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"log"
	"net/http"
	"strconv"
)

// Purchase struct (Model)
type Purchase struct {
	ID       string `json:"item_id"`
	Name     string `json:"item_name"`
	Quantity string `json:"item_quantity"`
	Rate     string `json:"item_rate"`
	Date     string `json:"item_purchase_date"`
}

// Login struct (Model)
type Login struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password1 string `json:"password"`
}

// PasswordReset struct (Model)
type PasswordReset struct {
	Email string `json:"email"`
	Code    string `json:"code"`
	VerificationStatus string `json:"verification_status"`
}

// Init allpurchases var as a slice Purchase struct
var allPurchases []Purchase

// Init allLogins var as a slice Login struct
var allLogins []Login

// Reverse string
// func Reverse(s string) string {
//     r := []rune(s)
//     for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
//         r[i], r[j] = r[j], r[i]
//     }
//     return string(r)
// }

func createConnection() *sql.DB {

	// Connect to the DB, panic if failed
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/items?sslmode=disable")
	if err != nil {
		fmt.Println(`Could not connect to db`)
		panic(err)
	}

	// check the connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// fmt.Println("Successfully connected!")
	// return the connection
	return db
}

// Get all allpurchases
func getPurchases(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM purchases`)
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	var col1 string
	var col2 string
	var col3 string
	var col4 string
	var col5 string
	allPurchases = nil
	for rows.Next() {
		rows.Scan(&col1, &col2, &col3, &col4, &col5)
		// fmt.Println(col1, col2, col3, col4)
		allPurchases = append(allPurchases, Purchase{ID: col1, Name: col2, Quantity: col3, Rate: col4, Date: col5})
	}
	json.NewEncoder(w).Encode(allPurchases)
}

// Get all allLogins
func getLogins(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM logins`)
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)

	var col1 string
	var col2 string
	var col3 string
	allLogins = nil
	for rows.Next() {
		rows.Scan(&col1, &col2, &col3)
		// fmt.Println(col1, col2, col3, col4)
		allLogins = append(allLogins, Login{Username: col1, Email: col2, Password1: col3})
	}
	json.NewEncoder(w).Encode(allLogins)
}

// Get single post
func getPurchase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// get the postid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	rows := db.QueryRow(`SELECT * FROM logins WHERE item_id=$1`, id)

	var col1 string
	var col2 string
	var col3 string
	var col4 string
	var col5 string
	allPurchases = nil

	rows.Scan(&col1, &col2, &col3, &col4, &col5)
	// fmt.Println(col1, col2, col3, col4)
	allPurchases = append(allPurchases, Purchase{ID: col1, Name: col2, Quantity: col3, Rate: col4, Date: col5})

	json.NewEncoder(w).Encode(allPurchases)

}

// Get single login
func getLogin1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// get the postid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Unable to convert the string into int.  %v", err)
	// }

	rows := db.QueryRow(`SELECT * FROM logins WHERE username=$1`, params["id"])

	var col1 string
	var col2 string
	var col3 string

	allLogins = nil

	rows.Scan(&col1, &col2, &col3)
	// fmt.Println(col1, col2, col3)
	allLogins = append(allLogins, Login{Username: col1, Email: col2, Password1: col3})

	json.NewEncoder(w).Encode(allLogins)

}

// Get single login
func getLogin2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// get the postid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	// id, err := strconv.Atoi(params["id"])

	// if err != nil {
	// 	log.Fatalf("Unable to convert the string into int.  %v", err)
	// }

	rows := db.QueryRow(`SELECT * FROM logins WHERE email=$1`, params["id"])

	var col1 string
	var col2 string
	var col3 string

	allLogins = nil

	rows.Scan(&col1, &col2, &col3)
	// fmt.Println(col1, col2, col3)
	allLogins = append(allLogins, Login{Username: col1, Email: col2, Password1: col3})

	json.NewEncoder(w).Encode(allLogins)

}

// Add new post
func createPurchase(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var p Purchase
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println("from api : in create purchase : this is post p : ", p)
	fmt.Println("from api : in create purchase : this is error: ", err)

	// IDConv, err := strconv.Atoi(p.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// QuantityConv, err := strconv.Atoi(p.Quantity)
	// if err != nil {
	// 	panic(err)
	// }
	// RateConv, err := strconv.Atoi(p.Rate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("here is id1: ", useridConv)
	// fmt.Println("here is id2: ", idConv)

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// fmt.Println(`from api : INSERT INTO purchases (item_id, item_name, item_quantity, item_rate, item_purchase_date) VALUES ($1, $2, $3, $4, $5)`, p.ID, p.Name, p.Quantity, p.Rate, p.Date)
	row, err := db.Exec(`INSERT INTO purchases (item_id, item_name, item_quantity, item_rate, item_purchase_date) VALUES ($1, $2, $3, $4, $5)`, p.ID, p.Name, p.Quantity, p.Rate, p.Date)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
	fmt.Printf("Inserted a single record %v", p.ID)
}

// Add new login
func createLogin(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var l Login
	err := json.NewDecoder(r.Body).Decode(&l)
	fmt.Println("from api : in create purchase : this is post p : ", l)
	fmt.Println("from api : in create purchase : this is error: ", err)

	// IDConv, err := strconv.Atoi(p.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// QuantityConv, err := strconv.Atoi(p.Quantity)
	// if err != nil {
	// 	panic(err)
	// }
	// RateConv, err := strconv.Atoi(p.Rate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("here is id1: ", useridConv)
	// fmt.Println("here is id2: ", idConv)

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// fmt.Println(`from api : INSERT INTO purchases (item_id, item_name, item_quantity, item_rate, item_purchase_date) VALUES ($1, $2, $3, $4, $5)`, p.ID, p.Name, p.Quantity, p.Rate, p.Date)
	row, err := db.Exec(`INSERT INTO logins (username, email, password1) VALUES ($1, $2, $3)`, l.Username, l.Email, l.Password1)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
	fmt.Printf("Inserted a single record %v", l.Username)
}

// Add new login
func createPR(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var p PasswordReset
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println("from api : in create purchase : this is post p : ", p)
	fmt.Println("from api : in create purchase : this is error: ", err)

	// IDConv, err := strconv.Atoi(p.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// QuantityConv, err := strconv.Atoi(p.Quantity)
	// if err != nil {
	// 	panic(err)
	// }
	// RateConv, err := strconv.Atoi(p.Rate)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("here is id1: ", useridConv)
	// fmt.Println("here is id2: ", idConv)

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// fmt.Println(`from api : INSERT INTO purchases (item_id, item_name, item_quantity, item_rate, item_purchase_date) VALUES ($1, $2, $3, $4, $5)`, p.ID, p.Name, p.Quantity, p.Rate, p.Date)
	row, err := db.Exec(`INSERT INTO passwordreset (email, code, verification_status) VALUES ($1, $2, $3)`, p.Email, p.Code, p.VerificationStatus)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
	fmt.Printf("Inserted a single record %v", p.Email)
}

// Update post
func updatePurchase(w http.ResponseWriter, r *http.Request) {
	// w.Header().Set("Content-Type", "application/json")
	// w.Header().Set("Access-Control-Allow-Origin", "*")

	var p Purchase
	err := json.NewDecoder(r.Body).Decode(&p)
	fmt.Println("this is post p : ", p)
	fmt.Println("this is error: ", err)
	// useridConv, err := strconv.Atoi(p.Userid)
	// if err != nil {
	// 	panic(err)
	// }
	// idConv, err := strconv.Atoi(p.ID)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("here is id1: ", useridConv)
	// fmt.Println("here is id2: ", idConv)

	//get the postid from the request params, key is "id"
	params := mux.Vars(r)

	//convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}
	fmt.Println("this is id: ", id)
	// fmt.Println()
	//create the postgres db connection
	db := createConnection()
	//close the db connection
	defer db.Close()

	// fmt.Println(`from api : update purchase : UPDATE purchases SET item_id=$1, item_name=$2, item_quantity=$3, item_rate=$4, item_purchase_date=$5 WHERE item_id=$6`, p.ID, p.Name, p.Quantity, p.Rate, p.Date, id)
	row, err := db.Exec(`UPDATE purchases SET item_id=$1, item_name=$2, item_quantity=$3, item_rate=$4, item_purchase_date=$5 WHERE item_id=$6`, p.ID, p.Name, p.Quantity, p.Rate, p.Date, id)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
	fmt.Printf("Inserted a single record %v", p.ID)
}

// Delete post
func deletePurchase(w http.ResponseWriter, r *http.Request) {

	// w.Header().Set("Access-Control-Allow-Origin", "*")
	// w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	// w.Header().Add("Access-Control-Allow-Headers", "Content-Type")
	// w.Header().Add("Content-Type", "application/json")

	// create the postgres db connection
	db := createConnection()
	// close the db connection
	defer db.Close()

	// get the postid from the request params, key is "id"
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])
	fmt.Println(id)

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	rows, err := db.Exec(`DELETE FROM purchases WHERE item_id=$1`, id)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}

// Main function
func main() {

	// pgConString := fmt.Sprintf("host=%s port=%d user=%s "+
	// 							"password=%s dbname=%s sslmode=disable",
	// 							hostname, port, username, password, database)
	// Init router
	r := mux.NewRouter()
	// fmt.Println(r)

	// Hardcoded data - @todo: add database
	// allpurchases = append(allpurchases, Purchase{Title: "hi", Body: "card", Userid: 200, ID: 66})
	// allpurchases = append(allpurchases, Purchase{Title: "hello", Body: "hi there", Userid: 2400, ID: 656})

	// Route handles & endpoints
	r.HandleFunc("/logins", getLogins).Methods("GET")
	r.HandleFunc("/logins/1/{id}", getLogin1).Methods("GET")
	r.HandleFunc("/logins/2/{id}", getLogin2).Methods("GET")
	r.HandleFunc("/purchases", getPurchases).Methods("GET")
	r.HandleFunc("/purchases/{id}", getPurchase).Methods("GET")
	r.HandleFunc("/purchases", createPurchase).Methods("POST")
	r.HandleFunc("/logins", createLogin).Methods("POST")
	r.HandleFunc("/purchases/{id}", updatePurchase).Methods("PUT")
	r.HandleFunc("/purchases/{id}", deletePurchase).Methods("DELETE")

	
	// Start server
	handler := cors.AllowAll().Handler(r)
	// handler := cors.Default().Handler(r)
	//fmt.Println("ok")
	log.Fatal(http.ListenAndServe(":3000", handler))

}
