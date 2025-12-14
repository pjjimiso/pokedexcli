package main

import (
    "strings"
    "bufio"
    "os"
    "fmt"

    "github.com/pjjimiso/pokedexcli/internal/pokeapi"
)


type config struct { 
    pokeapiClient       pokeapi.Client
    nextLocationsURL    *string
    prevLocationsURL    *string
}

type cliCommand struct { 
    name        string
    description string
    callback    func(*config) error
}

func getCommands() map[string]cliCommand { 
    return map[string]cliCommand { 
        "exit": {
            name:           "exit", 
            description:    "Exit the Pokedex",
            callback:       commandExit,
        },
        "help": {
            name:           "help",
            description:    "Displays a help message",
            callback:       commandHelp,
        },
        "map": {
            name:           "map", 
            description:    "Paginates forwards through location areas (20 per page)",
            callback:       commandMap,
        },
        "mapb": {
            name:           "mapb",
            description:    "Paginates backwards through location areas (20 per page)",
            callback:       commandMapb,
        },
        "explore": { 
            name:           "explore",
            description:    "Lists pokemon found in the given location",
            callback:       commandExplore,
        }
    }
}


func startRepl(cfg *config) { 
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
            if commandName == "explore" {
                err := command.callback(cfg, words[1])
            } else {
                err := command.callback(cfg, nil)
            }
            if err != nil { 
                fmt.Println(err)
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

