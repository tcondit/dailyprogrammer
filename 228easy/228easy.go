/*
	https://www.reddit.com/r/dailyprogrammer/comments/3h9pde/20150817_challenge_228_easy_letters_in/

	Description

	A handful of words have their letters in alphabetical order, that is
	nowhere in the word do you change direction in the word if you were to
	scan along the English alphabet. An example is the word "almost", which
	has its letters in alphabetical order.

	Your challenge today is to write a program that can determine if the
	letters in a word are in alphabetical order.  As a bonus, see if you
	can find words spelled in reverse alphabetical order.

	Input Description

	You'll be given one word per line, all in standard English. Examples:
	  almost
	  cereal

	Output Description

	Your program should emit the word and if it is in order or not.
	Examples:
	  almost IN ORDER
	  cereal NOT IN ORDER
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	// relative paths are nice, but can get confusing quickly
	// TODO accept input file at cmdline
	file, err := os.Open("./challenge-input")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	reader := bufio.NewReader(file)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)

	/*
		A word list contains at least one word, but setting length to 1
		adds an empty string as the first value. ints and other types
		may not do this, don't know yet.
	*/
	words := make([]string, 0)

	/*
		want to
		- append each word to an array slice
		- for each word in slice
			- read each byte and get its ASCII value
			- if each byte >= the one before, IN ORDER
			- else NOT IN ORDER but maybe IN REVERSE ORDER
			
			- if NOT IN ORDER but maybe IN REVERSE ORDER
				- read each byte and get its ASCII value
				- if each byte <= the one before, IN ORDER
				- else NOT IN ORDER
	*/
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	// Yay range!
	for _, word := range words {
		if in_order(word) {
			fmt.Println(word, "IN ORDER")
		} else if reverse_order(word) {
			fmt.Println(word, "REVERSE ORDER")
		} else {
			fmt.Println(word, "NOT IN ORDER")
		}
//	fmt.Println()
	}
}

func in_order(word string) bool {
	var in_order bool = true
	var current rune
	for _, letter := range word {
		if letter < current {
			in_order = false
			break
		}
		current = letter
	}
	return in_order
}

func reverse_order(word string) bool {
	var reverse_order bool = true
	var current rune = 'z'
	for _, letter := range word {
		if letter > current {
			reverse_order = false
			break
		}
		current = letter
	}
	return reverse_order
}
