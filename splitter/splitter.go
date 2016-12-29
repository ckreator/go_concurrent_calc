package splitter

import (
	"cmd_calculator/shared"
	"regexp"
)

// Splitter is a function that sends in a stream of 'tokens'. A token is a data structure that contains it's value and type
func Split(s string, write_to chan *shared.Token) {
	r, _ := regexp.Compile("[a-zA-Z_][a-zA-Z_0-9]*|[0-9]+(\\.[0-9]+)?([Ee][\\+\\-][0-9]+)?|[\\+\\-\\/\\*\\%]|[\\(\\)]")
	t := r.FindAllString(s, -1)
	var tok_type string
	for _, x := range t {
		if x == "|" || x == "+" || x == "-" || x == "*" || x == "%" {
			tok_type = "op"
		} else if x == ")" || x == "(" {
			tok_type = "punc"
		} else {
			tok_type = "num"
		}
		write_to <- shared.New_token(tok_type, x)
	}
	close(write_to)
}
