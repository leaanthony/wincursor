package main

import (
	"fmt"
	"time"

	"github.com/leaanthony/wincursor"
)

func main() {
	// Ensure the cursor is already visible
	wincursor.Show()
	fmt.Printf("This has a cursor at the end:")
	time.Sleep(time.Second * 5)
	wincursor.Hide()
	fmt.Println()
	fmt.Printf("This doesn't have a cursor at the end:")
	time.Sleep(time.Second * 5)
	fmt.Println()
	wincursor.Show()

}
