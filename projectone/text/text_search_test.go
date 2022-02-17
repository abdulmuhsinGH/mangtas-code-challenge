package text

import "testing"

//test Search function in text_search.go
//when a passage is passed in, words and number of occurences should be returned
func TestSearchPassageIsPassed(t *testing.T) {
	passage := "This is a test. This is only a test. I hope this works."
	words := Search(passage)
	expected := map[string]int{"this": 3, "test": 2, "a": 2, "is": 2, "only": 1, "hope": 1, "works": 1, "i": 1}
	if len(words) == 0 {
		t.Errorf("Search() actual = %v, expected %v", words, "not empty")
	}

	for word, count := range words {
		if count != expected[word] {
			t.Errorf("Search() actual = %v, expected %v", words, expected)
		}
	}

}

//test Search function in text_search.go
//when passage is empty, an empty string should be returned
func TestSearchPassageIsEmpty(t *testing.T) {
	passage := ""
	words := Search(passage)
	if len(words) != 0 {
		t.Errorf("Search() actual = %v, expected %v", words, "empty")
	}

}
