package main

import (
	"encoding/json"
	"math/rand"
	"os"
)

type NamesDictionary struct {
	Creatures  []string   `json:"creatures"`
	Adjectives Adjectives `json:"adjectives"`
}

type Adjectives struct {
	Nice []string `json:"nice"`
	Ugly []string `json:"ugly"`
}

var dictionary NamesDictionary

func generateRandomName(nice bool) string {
	var adjectives *[]string
	if nice {
		adjectives = &dictionary.Adjectives.Nice
	} else {
		adjectives = &dictionary.Adjectives.Ugly
	}
	return (*adjectives)[rand.Intn(len(*adjectives))] + " " + dictionary.Creatures[rand.Intn(len(dictionary.Creatures))]
}

func init() {
	byteValue, _ := os.ReadFile("names-dictionary.json")
	json.Unmarshal(byteValue, &dictionary)
	if len(dictionary.Creatures) == 0 || len(dictionary.Adjectives.Ugly) == 0 || len(dictionary.Adjectives.Nice) == 0 {
		panic("no dictionary")
	}
}
