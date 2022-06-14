package main

import (
	"bufio"
	"fmt"
	"learngo/books/the_go_prog_language/ch3/smp/mlib"
	"learngo/books/the_go_prog_language/ch3/smp/mp"
	"os"
	"strconv"
	"strings"
)

var (
	lib          *mlib.MusicManager
	id           int = 1
	ctrl, signal chan int
)

func handleLibCommands(tokens []string) {
	switch tokens[1] {
	case "list":
		for i := 0; i < lib.Len(); i++ {
			e, _ := lib.Get(i)
			fmt.Println(i+1, ":", e.Name, e.Artist, e.Source, e.Type)
		}

	case "add":
		if len(tokens) == 6 {
			id++
			lib.Add(&mlib.MusicEntry{strconv.Itoa(id), tokens[2], tokens[3], tokens[4], tokens[5]})
		} else {
			fmt.Println("USAGE: lib add")
		}

	case "remove":
		if len(tokens) == 3 {
			lib.Remove(1)
		} else {
			fmt.Println("USAGE: lib remove <name>")
		}

	default:
		fmt.Println("Unrecognized lib commad:", tokens[1])
	}
}

func handlePlayCommand(tokens []string) {
	if len(tokens) != 2 {
		fmt.Println("USAGE: play <name>")
		return
	}

	e := lib.Find(tokens[1])
	if e == nil {
		fmt.Println("The music does't exist.")
		return
	}

	mp.Play(e.Source, e.Type)
}

func main() {
	fmt.Println(`
		Enter follwing commands to control the player:
		lib list -- View the existing music lib
		lib add <name><artist><source><type> -- Add a music to the music lib
		lib remove <name> -- Remove the specified music from the lib
		play <name> -- Play the specified music
		`)

	lib = mlib.NewMusicManager()

	r := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter command-> ")

		rawLine, _, _ := r.ReadLine()
		line := string(rawLine)

		if line == "q" || line == "e" {
			break
		}

		tokens := strings.Split(line, " ")
		if tokens[0] == "lib" {
			handleLibCommands(tokens)

		} else if tokens[0] == "play" {
			handlePlayCommand(tokens)
		} else {
			fmt.Println("Unrecognized command: ", tokens[0])
		}
	}
}
