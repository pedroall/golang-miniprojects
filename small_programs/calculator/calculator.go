package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var stop bool = false
	var stacktype string = "number"
	stack := []any{}

	allowedOperations := []string{"sum", "sun", "mul", "div", "res"}

	for stop != true {
		switch stacktype {
		case "operation":
			fmt.Println(
				`Select an operation:
  sum (+)
  sub (-)
  mul (*)
  div (/)
  res (=)\n`,
			)
			operationReader := bufio.NewReader(os.Stdin)
			operation, operationerror := operationReader.ReadString('\n')
			operation = strings.TrimSpace(operation)

			if operationerror != nil {
				fmt.Println("Input error for writing an operation: ", operationerror)
				stop = true
			} else {
				if slices.Contains(allowedOperations, operation) {
					if operation == "res" {
						stop = true
					}
					stack = append(stack, operation)
					stacktype = "number"
				} else {
					fmt.Println("No allowed operation received!")
					stop = true
				}
			}
		case "number":
			fmt.Println("Provide a number: ")
			inputReader := bufio.NewReader(os.Stdin)
			input, inputerror := inputReader.ReadString('\n')
			input = strings.TrimSpace(input)
			if inputerror != nil {
				fmt.Println("InputError: ", inputerror)
				stop = true
			} else {
				number, numbererror := strconv.Atoi(input)
				if numbererror != nil {
					fmt.Println("Number was invalid because of error: ", numbererror)
					stop = true
				}
				stack = append(stack, number)
			}
			stacktype = "operation"
		}
	}
	lastIndex := stack[len(stack)-1]
	if lastIndex != "res" {
		fmt.Println("Operation ended with error.")
	} else {
		var total int = 0
		var operator string
		for _, value := range stack {
			if number, definition := value.(int); definition {
				if total != 0 {
					switch operator {
					case "sum":
						total += number
					case "sub":
						total -= number
					case "mul":
						total *= number
					case "div":
						total /= number
					default:
						fmt.Println("Unexpected error happened")
					}
				} else {
					total = number
				}
			} else {
				var str = value.(string)
				operator = str
			}
		}
		fmt.Println(total)
	}
}
