package pokeapi

import (
  "fmt"
)

type Type struct {
  Name      string
  WeakTo    []*Type
  StrongTo  []*Type
  ImmuneTo  []*Type
}

type Stats struct {
  HP        uint8
  Attack    uint8
  Defense   uint8
  SpAttack  uint8
  SpDefense uint8
  Speed     uint8
  Total     uint16
}

type Ability struct {
  Name    string
  isHidden bool
  Effect  string // The ability description
}


type Pokemon struct {
  // The Dex Number across all generations where
  // The index represents the generation -> 0: National Dex, i: Gen i
  // A value of -1 indicates that the pokemon isnt present in such generation
  DexNumber []int
  Name      string
  BaseStats Stats // The Pokemon's Base Stats
  
  Abilities []*Ability // The Posible abilities, size varies
  
  Types     []*Type // The Types
  
  EggGroup  string // The Egg Group
  Height    float32
  Weight    float32
  
  //HeldItem  Item
  // Does the pokemon have an evolution
  // From showdown's nomenclature: NFE - Not Finally Evolved
  isNFE     bool
}


func NewPokemon() Pokemon {
  var pkm Pokemon
  pkmData,_ := GetJson(SiteURL, true)
  var h bool = pkmData["abilities"].([]interface{})[1].
  (map[string]interface{})["is_hidden"].(bool)
  fmt.Println(h)
  return pkm
}
