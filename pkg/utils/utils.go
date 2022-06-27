package utils

type ToolsI interface {
	Reverse(s string) (result string)

	CountVowels(word string) (number int)
	IsVowel(char string) bool
}

type Tools struct {
}

func NewTools() ToolsI {
	t := new(Tools)
	return t
}

//Reverse Returns the string given backwards
func (t *Tools) Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

//IsVowel Returns a boolean based on the character given, knowing if it is a vowel or not
func (t *Tools) IsVowel(char string) bool {
	if (char == "a") || (char == "e") || (char == "i") || (char == "o") || (char == "u") {
		return true
	} else {
		return false
	}
}

//CountVowels Returns the number of vowels into the word given
func (t *Tools) CountVowels(word string) (number int) {
	for _, l := range word {
		isVow := t.IsVowel(string(l))
		if isVow {
			number++
		}
	}
	return
}
