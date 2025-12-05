package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"
)

func startRepl() { 
    for {
        fmt.Printf("Pokedex > ")
        input, err := getInput()
        if err != nil { 
            fmt.Errorf("error scanning input string '%s': %w", input, err)
            continue
        }

        words := cleanInput(input)
        if len(words) == 0 { 
            continue 
        }
        
        commandName := words[0]

        fmt.Printf("Your command was: %s\n", commandName)
    }
}

func getInput() (string, error) {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    if err := scanner.Err(); err != nil { 
        return "", err
    }
    input := scanner.Text()
    return input, nil
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

