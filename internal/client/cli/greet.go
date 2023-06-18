package cli

import (
	"os"
	"strings"

	"golang.org/x/term"
)

var (
	Greeting string
)

func init() {
	Greeting = greeting()
}

const (
	largeGreeting = `
░██████╗░░█████╗░██████╗░██╗░░██╗██╗░░██╗███████╗███████╗██████╗░███████╗██████╗░
██╔════╝░██╔══██╗██╔══██╗██║░░██║██║░██╔╝██╔════╝██╔════╝██╔══██╗██╔════╝██╔══██╗
██║░░██╗░██║░░██║██████╔╝███████║█████═╝░█████╗░░█████╗░░██████╔╝█████╗░░██████╔╝
██║░░╚██╗██║░░██║██╔═══╝░██╔══██║██╔═██╗░██╔══╝░░██╔══╝░░██╔═══╝░██╔══╝░░██╔══██╗
╚██████╔╝╚█████╔╝██║░░░░░██║░░██║██║░╚██╗███████╗███████╗██║░░░░░███████╗██║░░██║
░╚═════╝░░╚════╝░╚═╝░░░░░╚═╝░░╚═╝╚═╝░░╚═╝╚══════╝╚══════╝╚═╝░░░░░╚══════╝╚═╝░░╚═╝

`
	plainGreeting = "\nWelcome to Gophkeeper!\n\n"
)

func greeting() string {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	width, _, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		return plainGreeting
	}
	if width >= greetingWidth(largeGreeting) {
		return largeGreeting
	} else {
		return plainGreeting
	}
}

func greetingWidth(greet string) int {
	split := strings.SplitAfterN(strings.TrimPrefix(greet, "\n"), "\n", len(greet)/2)
	if len(split) > 0 {
		return len([]rune(split[0]))
	}
	return 0
}
