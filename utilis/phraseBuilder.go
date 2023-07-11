package utilis

import (
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

const VERBO = "verbo"
const SUBISTANTIVO = "substantivo"
const BASEURL = "https://dicionario.aizeta.com/verbetes"

func BuildPhrase() string {
	rand.Seed(time.Now().UnixNano())
	c := colly.NewCollector(
		colly.AllowedDomains("www.dicionario.aizeta.com", "dicionario.aizeta.com"),
	)

	//set a random delay so if we dont find a word it wont break
	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	letter := getRandomLetter(VERBO)
	maxPage := getMaxPage(c, letter)
	verbo := getVerb(c, letter, maxPage)

	return verbo
}
