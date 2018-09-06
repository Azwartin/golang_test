package validators

import "github.com/golang-collections/collections/stack"

//Validator validate that tag bracers are in correct order and balanced
type Validator struct{}

//Validate html tag
func (validator *Validator) Validate(html string) bool {
	st := stack.New()

	openBracket, closeBracket := '<', '>'
	for _, runeValue := range html {
		if runeValue == closeBracket && st.Pop() != openBracket {
			return false
		} else if runeValue == openBracket {
			st.Push(runeValue)
		}
	}

	return st.Len() == 0
}
