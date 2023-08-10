package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println(rand.Intn(5))
	fmt.Println("Time current time is", time.Now().AddDate(-1,2,3))
	fmt.Println("Let's find some squares ",math.Sqrt(255))
	fmt.Println("Hello!! 🍌 👶 💢")
	fmt.Println(math.Pi)
	const name, age = "JT and Iqy 4501100", 22
	fmt.Printf("%s is %d years old.\n", name, age)
	fmt.Println(`
	
	There once was a lady named muffet, who sat on hter tuffet, eating her curds
	and weight.  Along came a spider and sat down beside her and...yoou know the rest...
	
	`)
}
