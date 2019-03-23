package levenshteinsearch

import (
	"bufio"
	"log"
	"os"
	"strings"
	"testing"
)

var aliceWords []string

const maxSimilaritySearch = 10

func ensureAlice() error {

	if len(aliceWords) > 0 {
		return nil
	}

	file, err := os.Open("../../assets/alice/alice.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

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
		log.Fatal(err)
	}

	return nil
}

func BenchmarkNaive1Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceNaive(aliceWords, "rabbit", i)
		}
	}
}

func BenchmarkNaive3Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceNaive(aliceWords, "rabbit", i)
		}
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceNaive(aliceWords, "eart", i)
		}
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceNaive(aliceWords, "the", i)
		}
	}
}

func BenchmarkMap1Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	aliceMap := map[string]*WordInformation{}

	for _, word := range aliceWords {
		wordInfo := aliceMap[word]
		if wordInfo != nil {
			wordInfo.Count++
		} else {
			aliceMap[word] = &WordInformation{
				Count: 1,
			}
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceMap(aliceMap, "rabbit", i)
		}
	}
}

func BenchmarkMap3Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	aliceMap := map[string]*WordInformation{}

	for _, word := range aliceWords {
		wordInfo := aliceMap[word]
		if wordInfo != nil {
			wordInfo.Count++
		} else {
			aliceMap[word] = &WordInformation{
				Count: 1,
			}
		}
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceMap(aliceMap, "rabbit", i)
		}
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceMap(aliceMap, "eart", i)
		}
	}

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			distanceMap(aliceMap, "the", i)
		}
	}
}

func BenchmarkOptimized1Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	dict := CreateDictionary()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("rabbit", i)
		}
	}
}

func BenchmarkOptimized3Word(b *testing.B) {

	if err := ensureAlice(); err != nil {
		log.Fatal(err)
	}

	dict := CreateDictionary()

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("rabbit", i)
		}
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("eart", i)
		}
	}
	for i := 0; i < b.N; i++ {
		for i := 0; i < maxSimilaritySearch; i++ {
			dict.SearchAll("the", i)
		}
	}
}

func distanceNaive(data []string, wordToSearch string, distanceMax int) map[string]*WordInformation {
	result := map[string]*WordInformation{}

	for _, word := range data {
		distance := levenshtein([]rune(wordToSearch), []rune(word))
		if distance <= distanceMax {
			wordInfo := result[word]
			if wordInfo == nil {
				wordInfo = &WordInformation{
					Count: 1,
				}
				result[word] = wordInfo
			} else {
				wordInfo.Count++
			}
		}
	}
	return result
}

func distanceMap(data map[string]*WordInformation, wordToSearch string, distanceMax int) map[string]*WordInformation {
	result := map[string]*WordInformation{}

	for word, wordInfo := range data {
		distance := levenshtein([]rune(wordToSearch), []rune(word))
		if distance <= distanceMax {

			result[word] = wordInfo
		}
	}
	return result
}

// I admit, I stole this code somewhere in Internet...
func levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
