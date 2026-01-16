package dependency2


import(
	"fmt"

	"github.com/thesage6/puppy"
)
func main(){
	line1 := puppy.Bark()
	line2 := puppy.Sit()
	fmt.Println(line1)
	fmt.Println(line2)
}
