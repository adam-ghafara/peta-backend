package petabackend

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"
	userdata.Role = "admin"
	mconn := SetConnection("MONGODATA", "geojsonList")
	CreateNewUserRole(mconn, "userLogin", userdata)
}

func TestDeleteUser(t *testing.T) {
	mconn := SetConnection("MONGODATA", "geojsonList")
	var userdata User
	userdata.Username = "yyy"
	DeleteUser(mconn, "userLogin", userdata)
}

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"
	userdata.Role = "admin"

	// Create a MongoDB connection
	mconn := SetConnection("MONGODATA", "geojsonList")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "userLogin", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGODATA", "geojsonList")
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "userLogin", userdata)
}

// func TestProduct(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petapedia")
// 	var productdata Product
// 	productdata.Nomorid = 1
// 	productdata.Name = "dzikri"
// 	productdata.Description = "haq"
// 	productdata.Price = 1000
// 	productdata.Size = "XL"
// 	productdata.Stock = 100
// 	productdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
// 	CreateNewProduct(mconn, "product", productdata)
// }

// func TestAllProduct(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petapedia")
// 	product := GetAllProduct(mconn, "product")
// 	fmt.Println(product)
// }

// func TestUpdateGetData(t *testing.T) {
// 	mconn := SetConnection("MONGOSTRING", "petapedia")
// 	datagedung := GetAllBangunanLineString(mconn, "petapedia")
// 	fmt.Println(datagedung)
// }

func TestGeneratePasswordHash(t *testing.T) {
	password := "12345"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)
	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("12345", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGODATA", "geojsonList")
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "userLogin", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGODATA", "geojsonList")
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"

	anu := IsPasswordValid(mconn, "userLogin", userdata)
	fmt.Println(anu)
}

// func CreateContent(t *testing.T) {
// 	mconn := SetConnection("MONGODATA", "geojsonList")
// 	var contentdata Content
// 	contentdata.ID = 1
// 	contentdata.Content = "admin"
// 	contentdata.Description = "12345"
// 	contentdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
// 	CreateNewContent(mconn, "content", contentdata)
// }

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGODATA", "geojsonList")
	var userdata User
	userdata.Username = "admin"
	userdata.Password = "12345"
	userdata.Role = "admin"
	CreateUser(mconn, "userLogin", userdata)
}
