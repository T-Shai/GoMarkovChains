package mchains

// mchains package contains the core algorithm to create
// a markov chain using a txt file then generate a output string
// SEED MUST BE INITIALISED IN MAIN

import (
	"fmt"
	"math/rand"
)

// Mchains : struct
type Mchains struct {
	order     int                 // markov chains order
	data      string              // analysed string
	ngrams    map[string][]string // possibilities of completion
	maxLength int                 // output string's max length
}

// NewMchains : creates a new markov chains
// with given order, name, maxlength
func NewMchains(order int, name string, maxl int) *Mchains {
	m := Mchains{order, ReadFile(name), make(map[string][]string), maxl}
	return &m
}

// Train : trains the model
// construct the ngram with data
func (m *Mchains) Train() {
	for i := 0; i < len(m.data)-m.order; i++ {
		substr := m.data[i : i+m.order]
		s := string(m.data[i+m.order])
		m.ngrams[substr] = append(m.ngrams[substr], s)
	}
}

// Generate : generate text from the markov chain.
// string can't exceed maxlength
func (m Mchains) Generate(start string, length int) string {
	result := start
	for i := 0; true; i++ {
		if i > m.maxLength-1 {
			fmt.Printf("Max interation %d is reached !", i)
			break // max length
		}
		l := len(result)
		nexts := m.ngrams[result[l-m.order:l]]
		ln := len(nexts)
		if ln == 0 {
			fmt.Printf("No more possibility after %d iteration(s)", i)
			break // no more possibilities
		}
		rindex := rand.Intn(ln)
		next := nexts[rindex]
		result += next
	}
	return result
}
