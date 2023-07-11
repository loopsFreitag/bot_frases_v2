package utilis

import (
	"fmt"
	"math/rand"
)

func getRandomLetter(class string) string {
	var letters []string

	// we do that because those classes doesent have words for all letters
	if class == VERBO {
		letters = getLettersVerb()
		return letters[rand.Intn(len(letters))]
	}
	letters = getLegetLettersSub()
	return letters[rand.Intn(len(letters))]
}

func getLettersVerb() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "i", "l", "m", "n", "o", "p", "r", "s", "t", "v"}
}

func getLegetLettersSub() []string {
	return []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "l", "m", "n", "o", "p", "q", "r", "s", "t", "v"}
}

func getSupportedTrasitividades() []string {
	return []string{"v.t.d.", "v.intr.", "v.t.d. e intr."}
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func getUrl(grammarClass string, letter string, pages ...int) string {
	page := 1
	if len(pages) > 0 {
		page = pages[0]
	}

	return fmt.Sprintf("%s/%s/%s/%d", BASEURL, grammarClass, letter, page)
}
