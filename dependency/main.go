package main

import (
	"fmt"

	"github.com/thesage6/puppy"
)

func main(){
	s1:= puppy.Bark()
	s2:= puppy.Barks()
	s3:= puppy.Sit()
	fmt.Println(s1, s2, s3)
	fmt.Println(puppy.Barks())
	fmt.Println(puppy.BigBarks())

}