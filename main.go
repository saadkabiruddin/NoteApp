package main

import (
	"bufio"
	"fmt"
	"note/note"
	"os"
	"strings"
)

func main() {

	title, content := getNoteData()

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
	}

	err = userNote.SaveToFile()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Note Saved Successfully")

	fmt.Println("Note Saved Successfully")

	fmt.Println("Title:", title)
	fmt.Println("Content:", content)
}

func getNoteData() (string, string) {

	title := getUserInput("Note Title:")
	content := getUserInput("Note Content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Println(prompt)

	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n')

	return strings.TrimSpace(value)
}
