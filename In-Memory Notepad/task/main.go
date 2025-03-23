package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter a command and data: ")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := strings.Split(scanner.Text(), " ")
		command, data := input[0], strings.Join(input[1:], " ")

		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		default:
			fmt.Println(command, data)
			fmt.Print("Enter a command and data:")
		}
	}
}
