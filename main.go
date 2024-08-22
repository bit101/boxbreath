package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/bit101/go-ansi"
)

func main() {
	reps := flag.Int("reps", 8, "how many reps of breathing in and out to do.")
	count := flag.Int("count", 4, "how many beats to breath and hold")
	flag.Parse()

	messages := []string{"BREATH IN ", "HOLD      ", "BREATH OUT", "HOLD      "}
	colors := []ansi.AnsiColor{ansi.Yellow, ansi.Blue, ansi.Green, ansi.Blue}
	msg := 0
	for cycle := range *reps {
		ansi.Printf(ansi.White, "ROUND %d/%d\n", cycle+1, *reps)
		for range 4 {
			ansi.Printf(colors[msg%4], "  %s", messages[msg%4])
			for second := range *count {
				fmt.Printf(" %d", *count-second)
				time.Sleep(time.Second * 1)
			}
			msg++
			fmt.Println()
		}
	}
	ansi.Println(ansi.Yellow, "DONE!")
}
