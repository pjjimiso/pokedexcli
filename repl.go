package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"
)


type cliCommand struct { 
    name        string
    description string
    callback    func() error
    config      *Config
}

var config = Config{
    Next: "",
    Previous: "",
}

func getCommands() map[string]cliCommand { 
    return map[string]cliCommand { 
        "exit": {
            name:           "exit", 
            description:    "Exit the Pokedex",
            callback:       commandExit,
            config:         &config,
        },
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
            config:         &config,
        },
        "map": {
            name:           "map", 
            description:    "Paginates through location areas (20 per page)",
            callback:       commandMap,
            config:         &config,
        },
    }
}


func startRepl() { 
    url := "https://pokeapi.co/api/v2/location-area"
    config.Next = url

    scanner := bufio.NewScanner(os.Stdin)

    for {
        fmt.Printf("Pokedex > ")
        scanner.Scan()
        
        if err := scanner.Err(); err != nil { 
            fmt.Errorf("error scanning input string: %w", err)
            continue
        }

        words := cleanInput(scanner.Text())
        if len(words) == 0 { 
            continue 
        }
        
        commandName := words[0]

        command, exists := getCommands()[commandName]
        if exists {
            err := command.callback()
            if err != nil { 
                fmt.Errorf("error running command: %w", err)
            }
            continue
        } else {
            fmt.Println("Unknown command")
            continue
        }
    }
}

func cleanInput(text string) []string {
    return strings.Fields(strings.ToLower(text))
}

