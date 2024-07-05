package cli


import (
  "fmt"
  "os"
)

/*
  poketch [command] [flags] params

COMMANDS:
  help - get information about the commands available and command specific information

  pokemon - shows info about the pokemon. the info is divided into different categories    
    default: type, evolutions, abilities, base stats, dex#, generation, icon
    
    Categories:
    [Filters can be used for all categories]: generation, games,  
    - moves: learnset
    - pokedex: description, egg group, shape, height, weight, ...
    - in-game: encounter location, encounter method, ...
  
  move - shows info about moves. some options can be added for more information
    default: type, base power, accuracy, secondary effects, other interactions
    pokemon: pokemon that learn the move

  item - shows info about items
    - battleitems
*/


func GetParams() {
  args := os.Args
  if len(args) < 2 {
    fmt.Fprintf(os.Stderr, "%s: missing command\nRun poketch help for more info\n", args[0])
    return 
  }
  
  fmt.Printf("command: %s\n", args[1])
}
