package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func startRepl() {
    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("Pokedex >")
        scanner.Scan()

        words := cleanInput(scanner.Text())
        if len(words) == 0 {
            continue
        }

        commandName := words[0]
        fmt.Printf("Your command was: %s\n", commandName)
    }
}

func cleanInput(text string) []string {
    output := strings.ToLower(text)
    words := strings.Fields(output)
    return words
}

