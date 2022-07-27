//This package gnerates random chracters

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	InosukeSpeak()
}

//This function generates random characters.
func InosukeSpeak() {
	Quotes := []string{
		"It wont matter once I destory that thing!",
		"Go ahead and do whatever you want!",
		"so fast!",
		"I am totally fine!",
		"So freaking fast!",
		"Comin... through!!",
	}
	rand.Seed(time.Now().UnixNano())
	z := rand.Intn(6)
	fmt.Println(Quotes[z])

}
