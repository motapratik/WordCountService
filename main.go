package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/motapratik/WordCountService/api"
)

func main() {
	//create a new router
	router := mux.NewRouter()
	//specify endpoints, handler functions and HTTP method
	router.HandleFunc("/TopTenWordCount", api.TopTenWordCount).Methods("POST")
	//start and listen to requests
	http.ListenAndServe(":8080", router)
}

/*
	//s := []string{"this is my is is my wow so is is my my so hey no yes"}
	s := []string{"Create to a project that calls service created above, pass text from input file ‘GoLang_Test.txt’ and prints JSONoutput returned from the service. Program should be optimized to process large text file.This task should be completed using GoLang and uploaded to GitHub"}
	words := strings.Split(strings.Join(s, ""), " ")

	// count same words in s
	m := make(map[string]int)
	for _, word := range words {
		if _, ok := m[word]; ok {
			m[word]++
		} else {
			m[word] = 1
		}
	}

	// create and fill slice of word-count pairs for sorting by count
	wordCounts := make([]WordCount, 0, len(m))
	for key, val := range m {
		wordCounts = append(wordCounts, WordCount{word: key, count: val})
	}

	// sort wordCount slice by decreasing count number
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].count > wordCounts[j].count
	})

	// display the three most frequent words
	for i := 0; i < len(wordCounts) && i < 10; i++ {
		fmt.Println(wordCounts[i].word, ":", wordCounts[i].count)
	}

	w.WriteHeader(http.StatusOK)
*/
