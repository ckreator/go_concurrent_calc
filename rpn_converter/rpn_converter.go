package rpn_converter

import (
	"cmd_calculator/shared"
	"cmd_calculator/splitter"
)

/*
If the token is a number, then push it to the output queue.
If the token is a function token, then push it onto the stack.
If the token is a function argument separator (e.g., a comma):
Until the token at the top of the stack is a left parenthesis, pop operators off the stack onto the output queue. If no left parentheses are encountered, either the separator was misplaced or parentheses were mismatched.


If the token is an operator, o1, then:
	while there is an operator token o2, at the top of the operator stack and either
	o1 is left-associative and its precedence is less than or equal to that of o2, or
	o1 is right associative, and has precedence less than that of o2,
		pop o2 off the operator stack, onto the output queue;
	at the end of iteration push o1 onto the operator stack.

If the token is a left parenthesis (i.e. "("), then push it onto the stack.
If the token is a right parenthesis (i.e. ")"):
Until the token at the top of the stack is a left parenthesis, pop operators off the stack onto the output queue.
Pop the left parenthesis from the stack, but not onto the output queue.
If the token at the top of the stack is a function token, pop it onto the output queue.
If the stack runs out without finding a left parenthesis, then there are mismatched parentheses.
*/

var prec_table = map[string]int{"+": 10, "-": 10, "*": 20, "/": 20, "%": 10}

func prec(op string) int {
	return prec_table[op]
}

// RPNConverter writes an arithmetic expression
// in Reverse Polish Notation, which is a pretty cool way to handle
// arithmetic expressions as well as cache them, because they
// are bracket-independent
func RPNConverter(s string, out chan *shared.Token) {
	//fmt.Println("Started RPN Converter")
	tokens := make(chan *shared.Token)
	stack := shared.New_stack()
	var last_prec int
	var curr_prec int
	var last_op *shared.Token
	go func() { splitter.Split(s, tokens) }()
	for t := range tokens {
		// we'll use the shunting yard algorithm
		if t.Tok_type == "num" {
			out <- t
		} else if t.Tok_type == "punc" {
			if t.Value == "(" {
				stack.Push(t)
			} else {
				// pop off all operators off the stack until we encounter a left parenthesis
				last := stack.Pop()
				for last != nil {
					//fmt.Println("LAST IN STACK: ", stack.Peek(), last)
					if last.Tok_type == "op" {
						out <- last
						last = stack.Pop()
					} else {
						// pop off the last bracket
						//fmt.Println("LAST ELSE: ", last, stack.Pop())
						break
					}
				}
			}
		} else {
			// must be an operator
			//fmt.Printf("PREC of %s: %d\n", t.Value, prec(t.Value))
			// TODO: handle special op: 'minus sign' -> can be negation or minus
			last_op = stack.Peek()
			if last_op != nil && last_op.Value == "op" {
				last_prec = prec(last_op.Value)
				curr_prec = prec(t.Value)
				for last_prec <= curr_prec {
					out <- stack.Pop()
					last_op = stack.Peek()
					if last_op == nil {
						break
					}
					last_prec = prec(last_op.Value)
				}
				stack.Push(t)
			} else {
				stack.Push(t)
			}
			//out <- t
		}
		//fmt.Println("GOT TOKEN: ", t)
	}
	// empty the stack
	last := stack.Pop()
	for last != nil {
		out <- last
		last = stack.Pop()
	}
	close(out)
}
