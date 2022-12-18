package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/jinzhu/copier"
)

type Userr struct {
	Username  string
	Password  string
	UserEmail string
	Phone     string
	ID        string
	//should not be copied => mention in destination copier:"-"
	MID string
}

type UserClaims struct {
	Username string
	Password string
	Email    string `copier:"UserEmail"` //copy this fiels from certain field from src struct
	Phone    int    //does not copy since type is diff
	PhoneNum string //does not copy since variable names does not match
	ID       string `copier:"must",nopanic` //throws error (but does not panic) if field not copied
	MID      string `copier:"-"`            //does not copy this field
}

func main() {

	u := &Userr{"shiva", "psw", "shiva@email.com", "987987987987", "123", "m123"}
	uc := &UserClaims{}

	err := copier.Copy(uc, u)
	if err != nil {
		log.Println(err.Error())
	}
	printjson(u)
	printjson(uc)
}

func printjson(cls any) {
	b, _ := json.MarshalIndent(cls, " ", "  ")
	fmt.Println(string(b))
}
