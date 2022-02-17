package text

import (
	"sort"
	"strings"
)

/**
* Searches a text for the most occuring word
* returns list of words and number of occurences
 */
func Search(text string) map[string]int {
	if len(text) == 0 {
		return map[string]int{}
	}

	wordCounts := mapWordsToCounts(text)

	wordCountsSorted := sortMapByValue(wordCounts)

	return firstTenWords(wordCountsSorted, wordCounts)

}

func mapWordsToCounts(text string) map[string]int {
	//split the text into words
	words := strings.Fields(text)
	//create a map to store the words and their counts
	wordCounts := make(map[string]int)
	//loop through the words
	for _, word := range words {
		//to lowercase the word
		word = strings.ToLower(word)
		//remove punctuation
		word = removeNonAlphaCharacters(word)
		wordCounts[word]++
	}

	return wordCounts
}

func removeNonAlphaCharacters(word string) string {

	var result strings.Builder
	for i := 0; i < len(word); i++ {
		b := word[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') {
			result.WriteByte(b)
		}
	}
	return result.String()
	return word
}

func sortMapByValue(wordCounts map[string]int) map[string]int {
	wordCountsSortedMap := sortWordsByValue(wordCounts)

	if numOfWOrds := len(wordCountsSortedMap) < 10; numOfWOrds {
		return wordCountsSortedMap
	}

	return wordCountsSortedMap
}

func sortWordsByValue(wordCounts map[string]int) map[string]int {
	var wordCountsSorted []string
	for word := range wordCounts {
		wordCountsSorted = append(wordCountsSorted, word)
	}
	//sort the words by the number of occurences
	sort.Slice(wordCountsSorted, func(i, j int) bool {
		return wordCounts[wordCountsSorted[i]] > wordCounts[wordCountsSorted[j]]
	})

	//create a new map to store the sorted words and their counts
	wordCountsSortedMap := make(map[string]int)
	//loop through the sorted words
	for _, word := range wordCountsSorted {
		//add the word and its count to the new map
		wordCountsSortedMap[word] = wordCounts[word]
	}

	return wordCountsSortedMap
}

func firstTenWords(wordCountsSortedMap map[string]int, wordCounts map[string]int) map[string]int {
	wordCountsSortedMapTrimmed := make(map[string]int)
	count := 0
	for word := range wordCountsSortedMap {
		wordCountsSortedMapTrimmed[word] = wordCounts[word]
		count++
		if count == 10 {
			break
		}
	}
	return wordCountsSortedMapTrimmed
}
