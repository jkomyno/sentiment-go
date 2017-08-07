package sentiment

import (
	"bytes"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

/*
Result structure
*/
type Result struct {
	Score       float64
	Comparative float64
	Tokens      []string
	Words       []string
	Positive    []string
	Negative    []string
}

var negators = map[string]int{
	"cant":    1,
	"can't":   1,
	"dont":    1,
	"don't":   1,
	"doesnt":  1,
	"doesn't": 1,
	"not":     1,
	"non":     1,
	"wont":    1,
	"won't":   1,
	"isnt":    1,
	"isn't":   1,
}

var dataset = "build.json"

/*
GetJSON unmarshals dynamically the file indicated by filename and returns it
*/
func getJSON(filename string) map[string]interface{} {
	jsonStr, _ := ioutil.ReadFile(filename)
	var f map[string]interface{}

	if err := json.Unmarshal([]byte(jsonStr), &f); err != nil {
		panic(err)
	}

	return f
}

func keyExists(decoded map[string]interface{}, key string) bool {
	val, ok := decoded[key]
	return ok && val != nil
}

/*
Tokenize removes special characters and return an array of tokens (words)
*/
func Tokenize(input string) []string {
	lower := strings.ToLower(input)

	var buffer bytes.Buffer
	buffer.WriteString(`[.,\/#!$%\^&\*;:{}=_`)
	buffer.WriteString("`\"~()]")

	re := regexp.MustCompile(buffer.String())
	replaced := re.ReplaceAllString(lower, "")
	return strings.Split(replaced, " ")
}

/*
Analyze performs sentiment analysis on the provided input 'phrase'.
*/
func Analyze(phrase string) Result {
	afinn := getJSON(dataset)

	// Storage objects
	tokens := Tokenize(phrase)
	var score float64
	var words, positive, negative []string

	// Iterate over tokens
	length := len(tokens) - 1
	for length >= 0 {
		obj := tokens[length]

		if keyExists(afinn, obj) {
			item := afinn[obj].(float64)

			// Check for negation
			if length >= 0 {
				prevtoken := tokens[length-1]
				if negators[prevtoken] == 1 {
					item = -item
				}
			}

			words = append(words, obj)
			if item > 0 {
				positive = append(positive, obj)
			} else if item < 0 {
				negative = append(negative, obj)
			}

			score += item
		} else {
			// fmt.Println(obj, "does not exist")
		}
		length--
	}

	var comparative float64

	if length > 0 {
		comparative = score / float64(length)
	} else {
		comparative = 0
	}

	res := Result{
		Score:       score,
		Comparative: comparative,
		Tokens:      tokens,
		Words:       words,
		Positive:    positive,
		Negative:    negative,
	}

	return res
}
