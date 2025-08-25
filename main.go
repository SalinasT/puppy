package puppy

import "fmt"

func Bark() string {
	return "Woof!"
}

func Barks() string {
	return "Woof!"
}

func main() {
	fmt.Println(Bark())
	fmt.Println(Barks())
}
