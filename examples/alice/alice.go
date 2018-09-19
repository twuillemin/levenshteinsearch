package main

import (
	"bitbucket.org/twuillemin/levenshteinsearch/pkg/levenshteinsearch"
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	// Get an array of words from Alice In Wonderlands
	aliceWords, err := readAlice()
	if err != nil {
		log.Fatal(err)
	}

	// Create a dictionary
	dict := levenshteinsearch.CreateDictionary()

	// Add alice to the dictionary
	for _, word := range aliceWords {
		dict.Put(word)
	}

	// Get information about the dictionary
	log.Printf("Number of words in Alice In Wonderlands: %v", dict.WordCount)
	log.Printf("Number of unique words in Alice In Wonderlands: %v", dict.UniqueWordCount)

	// Get information about rabbit
	wordInformation := dict.Get("rabbit")
	if wordInformation!=nil {
		log.Printf("Number of times the word 'rabbit' is present: %v", wordInformation.Count)
	}	else {
		log.Printf("The word 'rabbit' is not part of Alice In Wonderlands")
	}

	// Get information about flat worm
	wordInformation = dict.Get("platyhelminth")
	if wordInformation!=nil {
		log.Printf("Number of times the word 'platyhelminth' is present: %v", wordInformation.Count)
	}	else {
		log.Printf("The word 'platyhelminth' is not part of Alice In Wonderlands")
	}

	// Get information about word similar to rabbit
	for distance := 0; distance < 4; distance++ {
		wordInformationByWord := dict.SearchAll("rabbit", distance)
		log.Printf("Number of words close to \"rabbit\" with a distance of %v: %v", distance, len(wordInformationByWord))
		for key,value := range wordInformationByWord {
			log.Printf("\tWord: '%v' count: %v", key,value.Count)
		}
	}
}

func readAlice() ([]string, error) {

	file, err := os.Open("assets/alice/alice.txt")

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	aliceWords := make([]string, 0, 0)

	for scanner.Scan() {
		nextWord := strings.TrimSpace(scanner.Text())
		nextWord = strings.Replace(nextWord, `"`, "", -1)
		nextWord = strings.Replace(nextWord, `.`, "", -1)
		nextWord = strings.Replace(nextWord, `,`, "", -1)
		nextWord = strings.Replace(nextWord, `;`, "", -1)
		nextWord = strings.Replace(nextWord, `:`, "", -1)
		nextWord = strings.Replace(nextWord, `!`, "", -1)
		nextWord = strings.Replace(nextWord, `?`, "", -1)
		for _, word := range strings.Split(nextWord, " ") {
			word = strings.TrimSpace(word)
			word = strings.ToLower(word)
			aliceWords = append(aliceWords, word)
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return aliceWords, nil
}
