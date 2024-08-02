package main

import (
	"bufio"
	"example.com/note/note"
	"fmt"
	"os"
	"strings"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}
	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Saving the note failed.")
		return
	}

	fmt.Scan("File saved successfully!")
}

func getNoteData() (string, string) {
	title := getUserInput("Title of the note: ")
	content := getUserInput("Content of the note: ")
	return title, content
}

func getUserInput(propmt string) string {
	fmt.Printf("%v ", propmt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
