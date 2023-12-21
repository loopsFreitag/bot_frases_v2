package utilis

import (
	"math/rand"
	"strconv"

	"github.com/gocolly/colly"
)

type Verbo struct {
	verbo          string
	transitividade string
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

func getVerb(c *colly.Collector, letter string, maxPage int) Verbo {
	var verboStruct Verbo
	verbos := make([]Verbo, 0)
	ulrVerbPage := getUrl(VERBO, letter)

	if maxPage > 0 {
		ulrVerbPage = getUrl(VERBO, letter, rand.Intn(maxPage))
	}

	c.OnHTML(".info-feat", func(e *colly.HTMLElement) {
		// get all verbos from this page
		// e.ChildText("p") is the word and e.ChildText("div") is the transitividade
		verbos = append(verbos, Verbo{e.ChildText("p"), e.ChildText("div")})
	})

	c.OnScraped(func(r *colly.Response) {
		// get a random verbo
		randomIndex := rand.Intn(len(verbos))
		verboStruct = verbos[randomIndex]
	})

	c.Visit(ulrVerbPage)
	c.Wait()

	// get transitividade first if we can't handle we just call build phrase again
	if !stringInSlice(verboStruct.transitividade, getSupportedTrasitividades()) {
		return getVerb(c, letter, maxPage)
	}

	return verboStruct
}
