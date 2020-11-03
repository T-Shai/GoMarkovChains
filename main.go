package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"time"

	"./mchains"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the start of the sentence :")
	start, _, _ := reader.ReadLine()
	fmt.Println(len(start))
	rand.Seed(time.Now().UTC().UnixNano())
	fmt.Println("creating markov chains...")
	m := mchains.NewMchains(4, "la-princesse-de-cleve-la-fayette.txt", 10000)
	fmt.Println("Training...")
	m.Train()
	gen := m.Generate(string(start), 600)
	err := ioutil.WriteFile("markov_output.txt", []byte(gen), 0644)
	if err != nil {
		panic(err)
	}
}
