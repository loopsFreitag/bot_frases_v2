package utilis

import (
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

const VERBO = "verbo"
const SUBISTANTIVO = "substantivo"
const BASEURL = "https://dicionario.aizeta.com/verbetes"

func BuildPhrase() []string {
	//var plural bool

	rand.Seed(time.Now().UnixNano())
	c := colly.NewCollector(
		colly.AllowedDomains("www.dicionario.aizeta.com", "dicionario.aizeta.com"),
	)
	//plural = rand.Intn(2) != 0

	//set a random delay so if we dont find a word it wont break
	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	letter := getRandomLetter(VERBO)
	maxPage := getMaxPage(c, letter)
	verbo := getVerb(c, letter, maxPage)

	//couldnt find a better way to do that
	grammarTenses := getGrammaticalTense()
	grammarTenseFn := grammarTenses[rand.Intn(len(grammarTenses))]
	sufixo := callGrammarFn(grammarTenseFn, verbo)

	return sufixo
}

func getGrammaticalTense() []string {
	return []string{"pret_imper", "pret_maisque", "futuro_pres", "futuro_pret"}
}

func callGrammarFn(fn string, verbo string) []string {
	var sufixo []string
	switch fn {
	case "pret_imper":
		sufixo = pret_imper(verbo)
	case "pret_maisque":
		sufixo = pret_maisque(verbo)
	case "futuro_pres":
		sufixo = futuro_pres(verbo)
	case "futuro_pret":
		sufixo = futuro_pret(verbo)
	}
	return sufixo
}

func pret_imper(verbo string) []string {
	var sufixos []string
	if verbo[len(verbo)-2:] == "ar" {
		sufixos = []string{"ava", "avas", "ava", "ávamos", "áveis", "avam"}
	} else {
		sufixos = []string{"ia", "ias", "ia", "íamos", "íeis", "iam"}
	}
	return sufixos
}

func pret_maisque(verbo string) []string {
	var sufixos []string
	ending := verbo[len(verbo)-2:]
	if ending == "ar" {
		sufixos = []string{"ara", "aras", "ara", "áramos", "áreis", "aram"}
	} else if ending == "er" {
		sufixos = []string{"era", "eras", "era", "êramos", "êreis", "eram"}
	} else {
		sufixos = []string{"ira", "iras", "ira", "íramos", "íreis", "iram"}
	}
	return sufixos
}

func futuro_pres(verbo string) []string {
	var sufixos []string
	ending := verbo[len(verbo)-2:]
	if ending == "ar" {
		sufixos = []string{"arei", "arás", "ará", "aremos", "areis", "arão"}
	} else if ending == "er" {
		sufixos = []string{"erei", "erás", "erá", "eremos", "ereis", "erão"}
	} else {
		sufixos = []string{"irei", "irás", "irá", "iremos", "ireis", "irão"}
	}
	return sufixos
}

func futuro_pret(verbo string) []string {
	var sufixos []string
	ending := verbo[len(verbo)-2:]
	if ending == "ar" {
		sufixos = []string{"aria", "arias", "aria", "aríamos", "aríeis", "ariam"}
	} else if ending == "er" {
		sufixos = []string{"eria", "erias", "eria", "eríamos", "eríeis", "eriam"}
	} else {
		sufixos = []string{"iria", "irias", "iria", "iríamos", "iríeis", "iriam"}
	}
	return sufixos
}
