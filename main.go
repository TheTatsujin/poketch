package main

import (
  "fmt"
  //"github.com/TheTatsujin/poketch/mod/apifetch/pokeapi"
  //"github.com/TheTatsujin/poketch/mod/apifetch/cache"
  "github.com/TheTatsujin/poketch/mod/cli"
  "github.com/TheTatsujin/poketch/mod/cli/graphics"
  //"github.com/gdamore/tcell/v2"
)


func main() {
  cli.GetParams()
  if err := graphics.NewLayout(); err != nil {
    fmt.Println(err)
  }
}
