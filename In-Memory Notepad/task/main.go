package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const MaxMemos = 5

type MemoPad struct {
	memos   []string
	maxSize int
}

const DefaultMaxMemos = 5

func NewMemoPad() *MemoPad {
	return &MemoPad{memos: make([]string, 0, DefaultMaxMemos)}
}

func (mp *MemoPad) add(data string) {
	if len(mp.memos) >= mp.maxSize {
		fmt.Println("[Error] Notepad is full")
		return
	}
	if data != "" {
		mp.memos = append(mp.memos, data)
		fmt.Println("[OK] The note was successfully created")
	} else {
		fmt.Println("[Error] Missing note argument")
	}
}

func (mp *MemoPad) list() {
	if len(mp.memos) > 0 {
		for i, val := range mp.memos {
			fmt.Printf("[Info] %d: %s\n", i+1, val)
		}
	} else {
		fmt.Println("[Info] Notepad is empty")
	}
}

func (mp *MemoPad) clear() {
	mp.memos = make([]string, 0, mp.maxSize)
	fmt.Println("[OK] All notes were successfully deleted")
}

func (mp *MemoPad) setMaxSize() {
	fmt.Print("Enter the maximum number of notes: ")
	var size int
	fmt.Scan(&size)
	if size <= 0 {
		fmt.Println("[Error] Size must be positive")
		return
	}

	if size < len(mp.memos) {
		fmt.Printf("[Warning] [d notes will be deleted\n", len(mp.memos)-size)
		mp.memos = mp.memos[:size]
	}

	newMemos := make([]string, len(mp.memos), size)
	copy(newMemos, mp.memos)
	mp.memos = newMemos
	mp.maxSize = size
}

func main() {
	processCommands()
}

func processCommands() {
	memoPad := NewMemoPad()
	memoPad.setMaxSize()
	for {
		command, data := readAndParseInput()
		switch command {
		case "create":
			memoPad.add(data)
		case "list":
			memoPad.list()
		case "clear":
			memoPad.clear()
		case "exit":
			fmt.Println("[Info] Bye!")
			return
		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func readAndParseInput() (string, string) {
	input := readInput()
	command, data := parseInput(input)
	return command, data
}

func readInput() string {
	fmt.Print("Enter a command and data: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func parseInput(input string) (string, string) {
	parsed := strings.Split(input, " ")
	return parsed[0], strings.Join(parsed[1:], " ")
}
