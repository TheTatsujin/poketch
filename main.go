package main

import (
  "fmt"
  //"github.com/TheTatsujin/poketch/mod/apifetch/pokeapi"
  //"github.com/TheTatsujin/poketch/mod/apifetch/cache"
  "github.com/TheTatsujin/poketch/mod/cli"
)


func main() {
  cli.GetParams()
  mainWindow := cli.Window{}
  mainWindow.NewLayout()
  mainWindow.AddEntry(cli.NewPokemonEntry())
  if err := cli.Start(mainWindow); err != nil {
    fmt.Println(err)
    return
  }
}
