package interpreter

import (
	"cmd_calculator/rpn_converter"
	"cmd_calculator/shared"
	"fmt"
	"strconv"
)

type Acc struct {
	Value float64
}

type runner struct {
	exec func(float64, ...float64) float64
	args int
}

func make_runner(f func(float64, ...float64) float64, args int) *runner {
	return &runner{exec: f, args: args}
}

// takes two args
func run_plus(acc float64, args ...float64) float64 {
	for _, i := range args {
		acc += i
	}
	return acc
}

//var plus = make_runner(run_plus, 2)

func run_minus(acc float64, args ...float64) float64 {
	for _, i := range args {
		acc -= i
	}
	return acc
}

var OPS = map[string]*runner{"+": make_runner(run_plus, 2), "-": make_runner(run_minus, 2)}

func Interpreter(s string) {
	fmt.Println("STARTING INTERPETER")
	cursor_store := -1
	//cursor_num := -1
	stack := shared.New_stack()
	num_stack := shared.New_num_stack()
	var accumulator *Acc
	rpn := make(chan *shared.Token)
	run_op := func(op string) {
		// TODO: should also support operators with different arguments sizes
		fn, has := OPS[op]
		if has {
			b := num_stack.Pop()
			a := num_stack.Pop()
			// now act upon them
			fmt.Println("PUSHING: ", fn.exec(a, b))
			num_stack.Push(fn.exec(a, b))
		}
	}

	go func() { rpn_converter.RPNConverter(s, rpn) }()
	for t := range rpn {
		stack.Push(t)
		cursor_store++
		if accumulator == nil {
			real, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				panic("Syntax Error")
			}
			// set the initial accumulator - first number
			accumulator = &Acc{Value: real}
		}
		// check whether we have
		if t.Tok_type == "op" {
			// operate on it
			run_op(t.Value)
		} else {
			// must be a number
			val, err := strconv.ParseFloat(t.Value, 64)
			if err != nil {
				panic("Syntax error")
			}
			num_stack.Push(val)
		}
		// now check whether t is an operator and if yes, act upon it
		/*if t.Tok_type == "num" {
			cursor++
		} else if t.Tok_type == "op" {
			// get the previous token and the accumulated token
			run_op(t.Value)
		}*/
		fmt.Println("INTERPRETER: ", t)
	}
	fmt.Println("NUM STACK: ", num_stack)
}

/*
// stack:
| 14 | 5 | + | 9 | - |...
// CURSOR:
// cursor always moves and stores first result
| 14 | 5 | + | 9 | - |...
// next:
| 19 | 9 | - |...
// next:
| 10 |...
// next:
*/

// ---------------
//| 1 | 4 | 9 | * | - |
// second stack
/*
first stack:
| 1 | 4 | 9 | * | - |
cursor:
  ^
      ^
	      ^
		  	  ^
| 1 | ...
| 1 | 4 | ...
| 1 | 4 | 9 | ... <- get multiplication
| 1 | 36 | ... <- get minus
| -35 | ...
*/
//| 1 | 36 | - |
