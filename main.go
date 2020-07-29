package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

// create database connection
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

// Usersec struct (Model)
type Usersec struct {
	UserID   string `json:"userid"`
	Menuname string `json:"menuname"`
	Mainmenu string `json:"mainmenu"`
}

// Init allUserSec var as a slice Purchase struct
var allUserSec []Usersec

// GetAllUsersec returns all usersec
func GetAllUsersec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db := createConnection()
	defer db.Close()

	rows, err := db.Query(`SELECT * FROM usersec`)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	var userid string
	var menuname string
	var mainmenu string

	allUserSec = nil
	for rows.Next() {
		rows.Scan(&userid, &menuname, &mainmenu)
		allUserSec = append(allUserSec, Usersec{UserID: userid, Menuname: menuname, Mainmenu: mainmenu})
	}
	json.NewEncoder(w).Encode(allUserSec)
}

// GetOneUsersec returns one usersec
func GetOneUsersec(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	db := createConnection()
	defer db.Close()

	parameters := mux.Vars(r)

	rows, err := db.Query(`SELECT * FROM usersec where userid = $1`, parameters["userid"])
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)

	var userid string
	var menuname string
	var mainmenu string

	allUserSec = nil
	for rows.Next() {
		rows.Scan(&userid, &menuname, &mainmenu)
		allUserSec = append(allUserSec, Usersec{UserID: userid, Menuname: menuname, Mainmenu: mainmenu})
	}
	json.NewEncoder(w).Encode(allUserSec)
}

// DeleteOneUsersecOneAccess deletes one usersec - all access
func DeleteOneUsersecOneAccess(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var usersec Usersec
	err := json.NewDecoder(r.Body).Decode(&usersec)

	db := createConnection()
	defer db.Close()

	rows, err := db.Exec(`DELETE FROM usersec WHERE userid = $1 AND menuname = $2 AND mainmenu = $3`, usersec.UserID, usersec.Menuname, usersec.Mainmenu)
	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}

// DeleteOneUsersecAllAccess deletes one usersec - all access
func DeleteOneUsersecAllAccess(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	parameters := mux.Vars(r)

	db := createConnection()
	defer db.Close()

	rows, err := db.Exec(`DELETE FROM usersec WHERE userid = $1`, parameters["userid"])

	if err != nil {
		panic(err)
	}
	fmt.Println(rows)
}

// CreateOneUsersec creates one usersec
func CreateOneUsersec(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	var usersec Usersec
	err := json.NewDecoder(r.Body).Decode(&usersec)

	db := createConnection()
	defer db.Close()

	row, err := db.Exec(`INSERT INTO usersec (userid, menuname, mainmenu) VALUES ($1, $2, $3)`, usersec.UserID, usersec.Menuname, usersec.Mainmenu)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
}

// UpdateOneUsersec updates on usersec
func UpdateOneUsersec(w http.ResponseWriter, r *http.Request) {
	type CustomUsersec struct {
		NewUserID   string `json:"newuserid"`
		NewMenuname string `json:"newmenuname"`
		NewMainmenu string `json:"newmainmenu"`
		OldUserID   string `json:"olduserid"`
		OldMenuname string `json:"oldmenuname"`
		OldMainmenu string `json:"oldmainmenu"`
	}
	var customusersec CustomUsersec

	err := json.NewDecoder(r.Body).Decode(&customusersec)

	db := createConnection()
	defer db.Close()

	row, err := db.Exec(`UPDATE usersec SET userid=$1, menuname=$2, mainmenu=$3 WHERE userid=$4 AND menuname=$5 AND mainmenu=$6`, customusersec.NewUserID, customusersec.NewMenuname, customusersec.NewMainmenu, customusersec.OldUserID, customusersec.OldMenuname, customusersec.OldMainmenu)
	if err != nil {
		log.Fatalf("Unable to execute the query. %v", err)
	}
	fmt.Println(row)
}

// Login struct (Model)
type Login struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password1 string `json:"password"`
}

// Init allLogins var as a slice Login struct
var allLogins []Login

// GetLogins returns allLogins
func GetLogins(w http.ResponseWriter, r *http.Request) {
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

// GetLogin1 returns single login
func GetLogin1(w http.ResponseWriter, r *http.Request) {
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

// GetLogin2 returns single login
func GetLogin2(w http.ResponseWriter, r *http.Request) {
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

// CreateLogin Adds new login
func CreateLogin(w http.ResponseWriter, r *http.Request) {

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

// Purchase struct (Model)
type Purchase struct {
	ID       string `json:"item_id"`
	Name     string `json:"item_name"`
	Quantity string `json:"item_quantity"`
	Rate     string `json:"item_rate"`
	Date     string `json:"item_purchase_date"`
}

// Init allpurchases var as a slice Purchase struct
var allPurchases []Purchase

// GetPurchases returns allpurchases
func GetPurchases(w http.ResponseWriter, r *http.Request) {
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

// GetPurchase returns single purchase
func GetPurchase(w http.ResponseWriter, r *http.Request) {
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

// CreatePurchase Adds new purchase
func CreatePurchase(w http.ResponseWriter, r *http.Request) {

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

// UpdatePurchase updates single purchase
func UpdatePurchase(w http.ResponseWriter, r *http.Request) {
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

// DeletePurchase deletes single purchase
func DeletePurchase(w http.ResponseWriter, r *http.Request) {

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

	r := mux.NewRouter()

	// Route handles & endpoints
	r.HandleFunc("/usersec", GetAllUsersec).Methods("GET")
	r.HandleFunc("/usersec/{userid}", GetOneUsersec).Methods("GET")
	r.HandleFunc("/usersec", CreateOneUsersec).Methods("POST")
	r.HandleFunc("/usersec", UpdateOneUsersec).Methods("PUT")
	r.HandleFunc("/usersec/{userid}", DeleteOneUsersecAllAccess).Methods("DELETE")



	r.HandleFunc("/logins", GetLogins).Methods("GET")
	r.HandleFunc("/logins/1/{id}", GetLogin1).Methods("GET")
	r.HandleFunc("/logins/2/{id}", GetLogin2).Methods("GET")
	r.HandleFunc("/logins", CreateLogin).Methods("POST")


	
	r.HandleFunc("/purchases", GetPurchases).Methods("GET")
	r.HandleFunc("/purchases/{id}", GetPurchase).Methods("GET")
	r.HandleFunc("/purchases", CreatePurchase).Methods("POST")
	r.HandleFunc("/purchases/{id}", UpdatePurchase).Methods("PUT")
	r.HandleFunc("/purchases/{id}", DeletePurchase).Methods("DELETE")

	handler := cors.AllowAll().Handler(r)

	log.Fatal(http.ListenAndServe(":3000", handler))
}
