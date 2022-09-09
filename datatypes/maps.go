package main

import "log"


 type User struct {
	FirstName string
	LastName string
 }

func main() {
  myMap := make(map[string]User)
  bob := User{
	FirstName: "Bob",
	LastName: "Jones",
  }
  
  myMap["dude"] = bob

  log.Println(myMap["dude"].FirstName)

}