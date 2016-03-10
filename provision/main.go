package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"tools-support/screen"

	"github.com/lentregu/Equinox/goops"
)

var clear screen.ClearWindow

var actions map[int]string

var stdin *bufio.Reader

var (
	bdServer string
	bdPort   int
	bdName   string
	bdColl   string
)

func init() {

	var fatalErr error
	defer func() {
		if fatalErr != nil {
			flag.PrintDefaults()
			log.Fatalln(fatalErr)
		}
	}()

	flag.Parse()

	if fatalErr != nil {
		return
	}

	clear = screen.NewClearScreenFunction(screen.DARWIN)
	stdin = bufio.NewReader(os.Stdin)

	actions = map[int]string{
		1: "createFaceList",
		2: "listFacesList",
		3: "addFace",
		4: "end",
	}
}

func main() {

	var option string
	for option != "end" {
		option = menu()
		switch {
		case option == "createFaceList":
			id, err := createFaceList()
			if err != nil {
				log.Fatal(err)
			} else {
				goops.Info("The list %s has been created", id)
			}
		case option == "listFacesList":
			list, err := getFaceList()
			if err != nil {
				log.Fatal(err)
			} else {
				fmt.Printf("Lists: %s\n", list)
			}
		case option == "addFace":
			goops.Info("AddFace......")
		case option == "end":
			return
		}
		fmt.Println("\nPress <ENTER>......")
		stdin.ReadLine()
	}
	os.Exit(0)
}

func menu() string {
	option := 0
	for option < 1 || option > 4 {
		clear := screen.NewClearScreenFunction(screen.DARWIN)
		clear()
		fmt.Println("1. Create Face List")
		fmt.Println("2. List of Faces lists")
		fmt.Println("3. Add Face")
		fmt.Println("4. Exit")
		fmt.Printf("\nChoose an option....:")
		if _, err := fmt.Fscanf(stdin, "%d", &option); err != nil {
			// In case of not introducing a number
			option = 0
		}
		stdin.ReadLine() //This line is necessary to flush the buffer because there is a "\n" left

	}
	return actions[option]
}
