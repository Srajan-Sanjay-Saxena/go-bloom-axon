package main

import (
	"bufio"
	"fmt"
	bf "go-bloom-axon/bloomFilter"
	"os"
	"strings"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
	colorBold   = "\033[1m"
)

func printBanner() {
	fmt.Println()
	fmt.Println(colorCyan + colorBold + "  ╔══════════════════════════════════════╗" + colorReset)
	fmt.Println(colorCyan + colorBold + "  ║       USERNAME AVAILABILITY CLI      ║" + colorReset)
	fmt.Println(colorCyan + colorBold + "  ║       powered by Bloom Filter        ║" + colorReset)
	fmt.Println(colorCyan + colorBold + "  ╚══════════════════════════════════════╝" + colorReset)
	fmt.Println(colorYellow + "  Type a username to check availability." + colorReset)
	fmt.Println(colorYellow + "  Type 'exit' to quit." + colorReset)
	fmt.Println()
}

func main() {
	filter := bf.New(10)
	scanner := bufio.NewScanner(os.Stdin)

	printBanner()

	for {
		fmt.Print(colorBold + "  → Enter username: " + colorReset)

		if !scanner.Scan() {
			break
		}

		username := strings.TrimSpace(scanner.Text())

		if username == "" {
			fmt.Println(colorYellow + "  ⚠  Username cannot be empty." + colorReset)
			fmt.Println()
			continue
		}

		if strings.ToLower(username) == "exit" {
			fmt.Println(colorCyan + "\n  Goodbye! 👋" + colorReset)
			fmt.Println()
			break
		}

		if filter.Contains([]byte(username)) {
			fmt.Printf(colorRed+"  ✗  '%s' is already taken.\n"+colorReset, username)
		} else {
			filter.Add([]byte(username))
			fmt.Printf(colorGreen+"  ✓  '%s' is available — username saved!\n"+colorReset, username)
		}

		fmt.Println()
	}
}
