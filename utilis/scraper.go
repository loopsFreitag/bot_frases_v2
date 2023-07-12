package utilis

import (
	"math/rand"
	"strconv"

	"github.com/gocolly/colly"
)

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

func getVerb(c *colly.Collector, letter string, maxPage int) string {
	var verbo string
	var transitividade string
	verbos := make([][]string, 0)
	ulrVerbPage := getUrl(VERBO, letter)

	if maxPage > 0 {
		ulrVerbPage = getUrl(VERBO, letter, rand.Intn(maxPage))
	}

	c.OnHTML(".info-feat", func(e *colly.HTMLElement) {
		// get all verbos from this page
		// e.ChildText("p") is the word and e.ChildText("div") is the trasitividade
		verbos = append(verbos, []string{e.ChildText("p"), e.ChildText("div")})
	})

	c.OnScraped(func(r *colly.Response) {
		//get a random verbo
		randomIndex := rand.Intn(len(verbos))
		verboslice := verbos[randomIndex]
		verbo = verboslice[0]
		transitividade = verboslice[1]
	})

	c.Visit(ulrVerbPage)
	c.Wait()

	// get transitividade first if we cant handle we just call build phrase again
	if !stringInSlice(transitividade, getSupportedTrasitividades()) {
		return getVerb(c, letter, maxPage)
	}

	return verbo
}
