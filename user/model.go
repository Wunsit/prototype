package user

//User is made to send user info
type User struct {
	ID        int    `json:"id"`        //id
	Email     string `json:"email"`     //email
	Passhash  string `json:"Passhash"`  //passhash
	Firstname string `json:"firstname"` //firsname
	Lastname  string `json:"lastname"`  //lastname
}
