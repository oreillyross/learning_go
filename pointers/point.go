package main

import "log"

func main() {
	var myString = "Green"
	log.Println("my string is", myString)
        changeMyStringWithPointer(&myString)	
	log.Println("new colour is", myString)
}

func changeMyStringWithPointer(s *string) {
	var newColour string = "Red"
	*s = newColour

}


