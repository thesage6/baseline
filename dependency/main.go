package main
import (
	"github.com/thesage6/puppy"
	"fmt"
)

func main(){
	s1 := puppy.Bark()
	s2 := puppy.Barks()
	s3 := puppy.Sit()
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
