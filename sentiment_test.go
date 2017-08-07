package sentiment

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTokenizeEnglish(t *testing.T) {
	var phrase string
	var expected []string

	phrase = "The cat went over the wall."
	expected = []string{"the", "cat", "went", "over", "the", "wall"}
	assert.Equal(t, Tokenize(phrase), expected, "they should be equal")

	phrase = "It should also parse \"ESCAPED\" sentences, should it"
	expected = []string{"it", "should", "also", "parse", "escaped", "sentences", "should", "it"}
	assert.Equal(t, Tokenize(phrase), expected, "they should be equal")
}

func TestTokenizeDiacritic(t *testing.T) {
	var phrase string
	var expected []string

	phrase = "This approach is naïve."
	expected = []string{"this", "approach", "is", "naïve"}
	assert.Equal(t, Tokenize(phrase), expected, "they should be equal")

	phrase = "The puppy bowl team was very coöperative."
	expected = []string{"the", "puppy", "bowl", "team", "was", "very", "coöperative"}
	assert.Equal(t, Tokenize(phrase), expected, "they should be equal")

	phrase = "The soufflé was delicious!"
	expected = []string{"the", "soufflé", "was", "delicious"}
	assert.Equal(t, Tokenize(phrase), expected, "they should be equal")
}

func TestAnalyzeSyncCorpus(t *testing.T) {
	dataset, _ := ioutil.ReadFile("fixtures/corpus.txt")
	result := Analyze(string(dataset))

	assert.Equal(t, result.Score, -2.0)
	assert.Equal(t, len(result.Tokens), 1416)
	assert.Equal(t, len(result.Words), 72)
}

type validateType struct {
	pass int
	fail int
}

/*
func validate(set map[string]interface{}) float64 {
	var obj validateType

	// Iterate over each word/class pair in the dataset
	for _, i := range set {
		fmt.Println(set[i])
	}
}
*/

func TestAnalyzeSyncFuzz(t *testing.T) {
	// dataset := getJSON("fixtures/amazon.txt")
	// result := Analyze(dataset)
}
