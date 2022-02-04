package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

type InputText struct {
	Text string `json:"text"`
}

type WordCount struct {
	Key   string `json:"word"`
	Value int    `json:"count"`
}

func TopTenWordCount(w http.ResponseWriter, r *http.Request) {

	var inputTextRequest InputText
	err := json.NewDecoder(r.Body).Decode(&inputTextRequest)
	if err != nil {
		fmt.Fprintf(w, "Error in Decode")
		w.WriteHeader(http.StatusBadRequest)
	}
	// Split Words by space
	textArray := strings.Split(inputTextRequest.Text, " ")
	wordOccurrenceMap := make(map[string]int)
	for _, text := range textArray {
		wordOccurrenceMap[text] += 1
	}
	// Sort Map Key-Value Pair
	occurenceWordMap := SortMapByValue(wordOccurrenceMap)

	// Create JSON of Map
	w.WriteHeader(http.StatusOK)
	errMsg := json.NewEncoder(w).Encode(occurenceWordMap)
	if errMsg != nil {
		fmt.Fprintf(w, "Error in creating JSON")
	}
}

func SortMapByValue(wordFrequencyMap map[string]int) []WordCount {

	var wordsCounts []WordCount
	for key, value := range wordFrequencyMap {
		wordsCounts = append(wordsCounts, WordCount{Key: key, Value: value})
	}

	sort.Slice(wordsCounts, func(i, j int) bool {
		return wordsCounts[i].Value > wordsCounts[j].Value
	})

	if len(wordsCounts) > 10 {
		return wordsCounts[:10]
	}

	return wordsCounts
}
