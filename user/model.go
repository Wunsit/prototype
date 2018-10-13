package user

//User is made to send user info
type User struct {
	ID          int      `json:"id"`        //id
	Name        string   `json:"name"`      //name
	Firstname   string   `json:"firstname"` //firsname
	Lastname    string   `json:"lastname"`  //lastname
	Email       string   `json:"email"`     //email
	ProductList []string `json:"products"`  //productlist
}
