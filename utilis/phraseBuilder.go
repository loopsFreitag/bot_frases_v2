package utilis

import (
	"math/rand"
	"time"

	"github.com/gocolly/colly"
)

const (
	VERBO        = "verbo"
	SUBISTANTIVO = "substantivo"
	BASEURL      = "https://dicionario.aizeta.com/verbetes"
)

type GrammarFunc func(string) []string

var grammarFuncs = map[string]GrammarFunc{
	"pret_imper":   pret_imper,
	"pret_maisque": pret_maisque,
	"futuro_pres":  futuro_pres,
	"futuro_pret":  futuro_pret,
}

func BuildPhrase() string {
	rand.Seed(time.Now().UnixNano())

	plural := rand.Intn(2) == 0
	c := colly.NewCollector(
		colly.AllowedDomains("www.dicionario.aizeta.com", "dicionario.aizeta.com"),
	)

	c.Limit(&colly.LimitRule{
		RandomDelay: 2 * time.Second,
		Parallelism: 4,
	})

	verbo := getVerbOrSub(c, VERBO, plural)

	grammarTenseFn := getRandomGrammarTense()
	ending := verbo.word[len(verbo.word)-2:]
	sufixo := grammarFuncs[grammarTenseFn](ending)

	verboStr := verbo.word[:len(verbo.word)-2] + sufixo[0]
	if plural {
		verboStr = verbo.word[:len(verbo.word)-2] + sufixo[1]
	}

	// INTRANSITIVO
	frase := getVerbOrSub(c, SUBISTANTIVO, plural).word + " " + verboStr

	switch verbo.util {
	// TRANSITIVO DIRETO
	case getSupportedTrasitividades()[0]:
		frase = frase + " " + getVerbOrSub(c, SUBISTANTIVO, plural).word
		// TRANSITIVO DIRETO OU INTRANSITIVO
	case getSupportedTrasitividades()[2]:
		randomNum := rand.Intn(2)
		if randomNum == 0 {
			frase = frase + " " + getVerbOrSub(c, SUBISTANTIVO, plural).word
		}
	}

	return frase
}

func getRandomGrammarTense() string {
	grammarTenses := []string{"pret_imper", "pret_maisque", "futuro_pres", "futuro_pret"}
	randomIndex := rand.Intn(len(grammarTenses))
	return grammarTenses[randomIndex]
}

func pret_imper(ending string) []string {
	if ending == "ar" {
		return []string{"ava", "avam"}
	}
	return []string{"ia", "iam"}
}

func pret_maisque(ending string) []string {
	if ending == "ar" {
		return []string{"ara", "aram"}
	} else if ending == "er" {
		return []string{"era", "eram"}
	}
	return []string{"ira", "iram"}
}

func futuro_pres(ending string) []string {
	if ending == "ar" {
		return []string{"ará", "arão"}
	} else if ending == "er" {
		return []string{"erá", "erão"}
	}
	return []string{"irá", "irão"}
}

func futuro_pret(ending string) []string {
	if ending == "ar" {
		return []string{"aria", "ariam"}
	} else if ending == "er" {
		return []string{"eria", "eriam"}
	}
	return []string{"iria", "iriam"}
}

func artMascPlural() []string {
	return []string{"os", "uns"}
}

func artMascSingular() []string {
	return []string{"o", "um"}
}

func artFemPlural() []string {
	return []string{"as", "umas"}
}

func artFemSingular() []string {
	return []string{"a", "uma"}
}
