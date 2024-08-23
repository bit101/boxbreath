package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"slices"
	"time"

	"github.com/bit101/go-ansi"
)

func main() {
	boxFlag := flag.Bool("box", false, "Do a box breathing session.")
	breath478Flag := flag.Bool("478", false, "Do a 4-7-8 breathing session.")
	flag.Usage = info

	flag.Parse()

	if *boxFlag {
		box()
		return
	}

	if *breath478Flag {
		breath478()
		return
	}

	choices := []string{
		"1. Box Breathing",
		"2. 4-7-8 Breathing",
		"Q. Quit",
	}
	ansi.Println(ansi.Yellow, "Choose an activity type")
	for _, choice := range choices {
		fmt.Println(choice)
	}

	switch getChoice() {
	case "1":
		box()
		return
	case "2":
		breath478()
		return
	case "q":
		os.Exit(0)
	}
}

func box() {
	ansi.Println(ansi.Green, "\nBox Breathing")
	reps := getReps()
	count := getCount()
	printTime(reps * count * 4)
	ansi.Println(ansi.Yellow, "Press ENTER to start")
	fmt.Scanln()
	messages := []string{"BREATH IN ", "HOLD      ", "BREATH OUT", "HOLD      "}
	colors := []ansi.AnsiColor{ansi.Yellow, ansi.Blue, ansi.Green, ansi.Blue}
	msg := 0
	for cycle := range reps {
		ansi.Printf(ansi.White, "ROUND %d/%d\n", cycle+1, reps)
		for range 4 {
			ansi.Printf(colors[msg], "  %s", messages[msg])
			for second := range count {
				fmt.Printf(" %d", count-second)
				time.Sleep(time.Second * 1)
			}
			msg++
			msg %= 4
			fmt.Println()
		}
	}
	ansi.Println(ansi.Yellow, "DONE!")
}

func breath478() {
	ansi.Println(ansi.Green, "\n4-7-8 Breathing")
	reps := getReps()
	printTime(reps * 19)
	ansi.Println(ansi.Yellow, "Press ENTER to start")
	fmt.Scanln()
	for cycle := range reps {
		ansi.Printf(ansi.White, "ROUND %d/%d\n", cycle+1, reps)

		ansi.Print(ansi.Yellow, "  BREATH IN ")
		for second := range 4 {
			fmt.Printf(" %d", 4-second)
			time.Sleep(time.Second * 1)
		}
		fmt.Println()

		ansi.Print(ansi.Blue, "  HOLD      ")
		for second := range 7 {
			fmt.Printf(" %d", 7-second)
			time.Sleep(time.Second * 1)
		}
		fmt.Println()

		ansi.Print(ansi.Green, "  BREATH OUT")
		for second := range 8 {
			fmt.Printf(" %d", 8-second)
			time.Sleep(time.Second * 1)
		}
		fmt.Println()
	}
	ansi.Println(ansi.Yellow, "DONE!")

}

func discardBuffer(r *bufio.Reader) {
	r.Discard(r.Buffered())
}

func getChoice() string {
	stdin := bufio.NewReader(os.Stdin)
	var choice string
	validChoices := []string{"1", "2", "q"}
	for true {
		ansi.Print(ansi.Yellow, "Choice > ")
		if _, err := fmt.Fscanln(stdin, &choice); err != nil {
			discardBuffer(stdin)
			ansi.Println(ansi.Red, "Please enter a valid choice.")
			continue
		}
		if slices.Contains(validChoices, choice) {
			break
		}
		ansi.Println(ansi.Red, "Please enter a valid choice.")
		continue
	}
	return choice
}

func getReps() int {
	stdin := bufio.NewReader(os.Stdin)
	var reps int
	for true {
		ansi.Print(ansi.Yellow, "How many breaths?")
		ansi.Print(ansi.White, " (try 4 to 10) ")
		if _, err := fmt.Fscanln(stdin, &reps); err != nil {
			discardBuffer(stdin)
			ansi.Println(ansi.Red, "Please enter a valid number.")
			continue
		}
		if reps < 1 {
			ansi.Println(ansi.Red, "Please enter a number greater than 0.")
			continue
		}
		break
	}
	return reps
}

func getCount() int {
	stdin := bufio.NewReader(os.Stdin)
	var count int
	for true {
		ansi.Print(ansi.Yellow, "Count for each segment?")
		ansi.Print(ansi.White, " (4 or 5 is usual) ")
		if _, err := fmt.Fscanln(stdin, &count); err != nil {
			discardBuffer(stdin)
			ansi.Println(ansi.Red, "Please enter a valid number.")
			continue
		}
		if count < 3 {
			ansi.Println(ansi.Red, "Please enter a count of at least 3 seconds.")
			continue
		}
		if count > 10 {
			ansi.Println(ansi.Red, "That's probably too much. Try something from 3 to 10 seconds.")
			continue
		}
		break
	}
	return count
}

func printTime(seconds int) {
	mins := seconds / 60
	seconds -= mins * 60
	minString := "minutes"
	if mins < 1 {
		fmt.Printf("This activity will take %d seconds\n", seconds)
		return
	} else if mins == 1 {
		minString = "minute"
	}
	fmt.Printf("This activity will take %d %s and %d seconds\n", mins, minString, seconds)

}

func info() {
	fmt.Println()
	ansi.Println(ansi.Yellow, "Usage:")

	fmt.Println("boxbreath [flags]")
	fmt.Println("If started without flags, a menu of options will be shown.")
	fmt.Println()

	ansi.Println(ansi.Yellow, "Flags:")

	ansi.Println(ansi.Yellow, "-h ")
	fmt.Println("Shows this help.")
	fmt.Println()

	ansi.Println(ansi.Yellow, "-box ")
	fmt.Println("Starts a box breathing session")
	fmt.Println("In box breathing, you will inhale for a certain number of seconds, then hold that breath for the same amount of time.")
	fmt.Println("Then you will exhale for that same count, and finally hold once again.")
	fmt.Println("You can choose how many breaths to take, and the count for each segment.")
	fmt.Println()

	ansi.Println(ansi.Yellow, "-478 ")
	fmt.Println("Starts a 4-7-8 breathing session")
	fmt.Println("In 4-7-8 breathing, you will breathe in for four seconds, hold that breath for seven seconds and exhale for eight seconds.")
	fmt.Println("You can choose how many breaths to take.")
}
