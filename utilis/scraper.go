package utilis

import (
	"math/rand"
	"strconv"

	"github.com/gocolly/colly"
)

type GrammarClass struct {
	word string
	//Verbo is transitivade
	//Subsitantivo is genero
	util string
}

func getMaxPage(c *colly.Collector, letter string) int {
	var maxPage int

	c.OnHTML("div.pagination", func(e *colly.HTMLElement) {
		e.ForEach("li", func(i int, h *colly.HTMLElement) {
			pageNumber, err := strconv.Atoi(h.Text)
			if err == nil && pageNumber > maxPage {
				maxPage = pageNumber
			}
		})

		if maxPage >= 5 {
			c.Visit(getUrl(VERBO, letter, maxPage))
		}
	})
	c.Visit(getUrl(VERBO, letter))

	return maxPage
}

func getVerbOrSub(c *colly.Collector, grammarClass string, plural bool) GrammarClass {
	letter := getRandomLetter(grammarClass)
	maxPage := getMaxPage(c, letter)

	var wordStruct GrammarClass
	grammarstr := make([]GrammarClass, 0)
	ulrPage := getUrl(grammarClass, letter)

	if maxPage > 0 {
		ulrPage = getUrl(grammarClass, letter, rand.Intn(maxPage))
	}

	c.OnHTML(".info-feat", func(e *colly.HTMLElement) {
		// get all verbos from this page
		// e.ChildText("p") is the word and e.ChildText("div") is the util transitivade/genero
		grammarstr = append(grammarstr, GrammarClass{e.ChildText("p"), e.ChildText("div")})
	})

	c.OnScraped(func(r *colly.Response) {
		if len(grammarstr) > 0 {
			randomIndex := rand.Intn(len(grammarstr))
			wordStruct = grammarstr[randomIndex]
		}
	})

	c.Visit(ulrPage)
	c.Wait()
	//if its an verbo we see if we can handle the transitividade
	// get transitividade first if we can't handle we just call build phrase again
	if grammarClass == VERBO {
		if !stringInSlice(wordStruct.util, getSupportedTrasitividades()) {
			return getVerbOrSub(c, grammarClass, plural)
		}
	}

	return wordStruct
}
