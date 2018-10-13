package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

//User is made to send user info
type User struct {
	ID          int      `json:"User_ID"`        //ID
	Name        string   `json:"User_name"`      //name
	Firstname   string   `json:"User_firstname"` //firsname
	Lastname    string   `json:"User_lastname"`  //lastname
	Email       string   `json:"User_email"`     //email
	ProductList []string `json:"products"`       //productlist
}

// type Product struct {
// 	ID      int    `json:"Product_id"`
// 	Website string `json:"Product_site"`
// }
func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "hi")
	case http.MethodPost:
		fmt.Fprint(w, "Updating List")
	case http.MethodPut:
		fmt.Fprint(w, "Creating List")
	case http.MethodDelete:
		fmt.Fprint(w, "Deleting List")
	default:
		fmt.Fprint(w, "404 Error List")
	}
	// l := &List{
	//     ID:   1,
	//     Products: []string{"apple", "peach", "pear"}}
	// jsonList, _ := json.Marshal(l)
	// fmt.Fprint(w, jsonList)
	// json.NewEncoder(w).Encode(l)
}

func userMain(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/"):]
	fmt.Fprint(w, r.Method)
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprint(w, r.URL.Query())
	if name == "" {
		switch r.Method {
		case http.MethodGet:
			l := &User{
				ID:          10000000,
				Name:        "hello",
				Firstname:   "Nathalie",
				Lastname:    "desune",
				Email:       "hotstuff@hotmail.com",
				ProductList: []string{"website", "herro"}}
			json.NewEncoder(w).Encode(l)


			// base, _ := json.MarshalIndent(l, "", " ")
			// fmt.Println(string(base))

			// _ = ioutil.WriteFile("newinfo.json", base, 0644)

			// file, err := os.OpenFile("allusers.log", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0660)
			// if err != nil {
			// 	log.Fatal(err)
			// }
			// file.Write(base)
			op := &api.Operation{
				Schema: `name: string @index(exact) .`,
			}
			err := DB.Alter(r.Context(), op)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Fprint(w, "Show all User")
		case http.MethodPost:
			fmt.Fprint(w, "Updating User")
		case http.MethodPut:
			fmt.Fprint(w, "Creating User")
		default:
			fmt.Fprint(w, "404 Error User")
		}
	} else {
		testname, _ := strconv.ParseFloat(name, 64)
		if testname == 0 {
			fmt.Println("404")
		} else {
			switch r.Method {
			case http.MethodGet:
				fmt.Fprint(w, "Show specic user")
			case http.MethodDelete:
				fmt.Fprint(w, "Deleting User")
			default:
				fmt.Fprint(w, "404 Error User")
			}
		}
		// fmt.Fprint(w, name)
		// fmt.Fprint(w, "This is a generic welcome message.")
		// fmt.Println("This is a generic welcome message.")
	}
}

func user(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Path[len("/user/"):]
	fmt.Fprint(w, r.Method)
	fmt.Fprint(w, r.URL.Path)
	fmt.Fprint(w, r.URL.Query())
	if name == "" {
		switch r.Method {
		case http.MethodGet:
			fmt.Fprint(w, "Show all User")
		case http.MethodPost:
			fmt.Fprint(w, "Updating User")
		case http.MethodPut:
			fmt.Fprint(w, "Creating User")
		default:
			fmt.Fprint(w, "404 Error User")
		}
	} else {
		testname, _ := strconv.ParseFloat(name, 64)
		if testname == 0 {
			fmt.Println("404")
		} else {
			switch r.Method {
			case http.MethodGet:
				fmt.Fprint(w, "Show specic user")
			case http.MethodDelete:
				fmt.Fprint(w, "Deleting User")
			default:
				fmt.Fprint(w, "404 Error User")
			}
		}
		// fmt.Fprint(w, name)
		// fmt.Fprint(w, "This is a generic welcome message.")
		// fmt.Println("This is a generic welcome message.")
	}
}

func welcome(w http.ResponseWriter, r *http.Request)  {
	welcomestring := "Hello and welcome to this api, to get to users fo /user/"
	fmt.Fprint(w, welcomestring)
}


func databasconnetion(addr string) (*dgo.Dgraph, func() error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Trying to dial grpc")
	}
	dgraphClient := dgo.NewDgraphClient(api.NewDgraphClient(conn))

	fmt.Println("Connected to database")
	// Check error
	return dgraphClient, conn.Close
}

//DB is
var DB *dgo.Dgraph

func main() {
	addr := "localhost:9080"
	var closer func() error
	DB, closer = databasconnetion(addr)
	defer closer()

	http.HandleFunc("/", welcome)
	http.HandleFunc("/user/", user)
	// // http.HandleFunc("/product/", welcome)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
