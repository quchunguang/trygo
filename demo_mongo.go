package trygo

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
)

// AddressBook struct
type AddressBook struct {
	Name  string
	Phone string
}

// DemoMongo func
func DemoMongo() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	c := session.DB("test").C("people")

	// err = c.Insert(&AddressBook{"Ale", "+55 53 8116 9639"},
	//  &AddressBook{"Cla", "+55 53 8402 8510"})
	// if err != nil {
	//  log.Fatal(err)
	// }

	result := AddressBook{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)
}
