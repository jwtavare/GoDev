package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(rand.Intn(5))
	fmt.Println("Hello!! 🍌 👶 💢")
	const name, age = "JT and Iqy 4501100", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Println(`
	
	There once was a lady named muffet, who sat on hter tuffet, eating her curds
	and weight.  Along came a spider and sat down beside her and...yoou know the rest...
	
	`)
}
