// Simple Go program to prevent AFK timeouts during FFXIV Endwalker launch.
// Presses escape 4x after any 10 minute interval of no user input.
package main

import (
	"fmt"
	"time"

	"github.com/eiannone/keyboard"
	"github.com/go-vgo/robotgo"
)

var (
	lastKeypressTime time.Time = time.Now()
	dateFormat       string    = "Mon Jan _2 15:04:05 2006"
)

func main() {
	fmt.Println("Idler - presses escape 4x after any 10 minute interval of no user input.")

	// Validate keyboard device for reading
	if err := keyboard.Open(); err != nil {
		panic(err)
	}

	// Spawn key listener routine
	go keyListener()

	// Await keypresses, calc time diff, send macros - loop
	for {
		if time.Since(lastKeypressTime) > time.Minute*10 {
			fmt.Printf("Idler activated at %s, pressing keys...\n", time.Now().Format(dateFormat))
			for i := 0; i < 4; i++ {
				// Stop macro inputs if user presses key during
				// macro execution
				if time.Since(lastKeypressTime) < time.Second*1 {
					break
				}

				robotgo.KeyTap("escape")
				time.Sleep(1 * time.Second)
			}
			lastKeypressTime = time.Now()
		}
	}
}

// Background goroutine to record time of latest keypress
func keyListener() {
	for {
		keyEvents, err := keyboard.GetKeys(10)
		if err != nil {
			panic(err)
		}

		for keypress := range keyEvents {
			if keypress.Err != nil {
				panic(keypress.Err)
			}

			// Log formatted user input to stdout
			fmt.Printf("Key Pressed - %q (%s)\n", keypress.Rune, time.Now().Format(dateFormat))

			// Ignore escape, as it is reserved for
			// this script
			if keypress.Key != 27 {
				lastKeypressTime = time.Now()
			}
		}
		time.Sleep(1 * time.Second)
	}
}
