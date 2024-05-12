package main

import (
	"fmt"
	"os"
	"time"

	"golang.org/x/term"
)

type BPMCalculator struct {
	PrevTime  time.Time
	Count     int
	TotalTime float64
	BPM       float64
}

func bpmCalculator() *BPMCalculator {
	return &BPMCalculator{}
}

func (b *BPMCalculator) updateBpm() {
	if b.Count > 0 {
		elapsedTime := time.Since(b.PrevTime).Seconds()
		b.BPM = 60 / elapsedTime
		b.TotalTime += elapsedTime
	}
	b.PrevTime = time.Now()
	b.Count++
}

func main() {
	fmt.Println("Press any key (press 'q' to terminate):")

	bpmCalculator := bpmCalculator()

	// Disable echoing of input
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)

	for {
		var buffer [1]byte
		_, err := os.Stdin.Read(buffer[:])

		if err != nil {
			fmt.Println("Error reading input:", err)
			break
		}

		if buffer[0] == 'q' {
			fmt.Println("Terminating...")
			break
		}

		bpmCalculator.updateBpm()

		if bpmCalculator.Count > 1 {
			averageBPM := int(float64(bpmCalculator.Count-1) / bpmCalculator.TotalTime * 60)
			fmt.Printf("\rAverage BPM: %d", averageBPM)
		}
	}

	fmt.Println()
}
