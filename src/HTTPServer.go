package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Character struct {
	Name  string
	Phone string
}


func handler(w http.ResponseWriter, r *http.Request) {
	session, err := mgo.Dial("")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Character{"Ale", "+55 53 8116 9639"},
		&Character{"Cla", "+55 53 8402 8510"})
	if err != nil {
		log.Fatal(err)
	}

	result := Character{}
	err = c.Find(bson.M{"name": "Ale"}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Phone:", result.Phone)

	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])

}

func main() {




	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
func find() {

}
