package main

import (
	"fmt"
	"time"
	rivescript "github.com/aichaos/rivescript-go"
)

type time_t float64

var TEST_INPUTS [6]string = [6]string{
	"Hello.",
	"Hello Alice",
	"Who is dr. wallace?",
	"What is your name?",
	"What is my name?",
	"What is the capital of Texas?",
}

func main() {
	count := Timereps(5, TestRS)
	fmt.Printf("Average time of run: %08fs\n", count)
}

// TestRS actually does the RiveScript benchmark testing.
func TestRS() {
	t := Now()

	// Load Alice from disk.
	bot := rivescript.New()
	bot.LoadDirectory("./alice")
	t = Delta(t, "Loading A.L.I.C.E. from disk.")

	// Sort the replies.
	bot.SortReplies()
	t = Delta(t, "Sort the replies.")

	// Get some responses.
	for _, message := range TEST_INPUTS {
		reply := bot.Reply("user", message)
		t = Delta(t, "Reply to %s => %s", message, reply)
		// bot.Debug = true
	}

	fmt.Printf("\n")
}

// Delta prints a time delta along with a message.
func Delta(t time_t, desc string, misc ...interface{}) time_t {
	delta := Now() - t
	fmt.Printf("[%0.8fs] %s\n", delta, fmt.Sprintf(desc, misc...))
	return Now()
}

// Timereps runs a function multiple times and returns the mean time.
func Timereps(reps int, test func()) time_t {
	times := []time_t{}
	for i := 0; i < reps; i++ {
		start := Now()
		test()
		times = append(times, Now() - start)
	}
	return Average(times)
}

// Now returns the current Unix time.
func Now() time_t {
	return time_t(float64(time.Now().UnixNano()) / 1000000000.0)
}

// Average returns the average of a set of floats.
func Average(xs []time_t) time_t {
	total := time_t(0.0)
	for _, v := range xs {
		total += v
	}
	return total / time_t(len(xs))
}
