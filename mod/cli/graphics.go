
package cli


import (
  "github.com/TheTatsujin/poketch/mod/apifetch/pokeapi"
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
  "strconv"
)

const (
  mediumScreenSize = 50
  entryHeight = 16


  // Pokemon Stat Grid
    // Medium Size ratio
    mediumStatLeft = 1
    mediumStatMiddle = 1
    mediumStatRight = 6

    // Small Size ratio
    smallStatLeft = 1
    smallStatMiddle = 1
    smallStatRight = 0

  numberStats = 6
)

type Window struct {
  displayGrid   *tview.Grid
  commandInput  *tview.InputField
  layout        *tview.Flex
}


type PokemonEntry struct {
  layout *tview.Grid
}

func getStatNames() ([6]string) {
  return [numberStats]string{"Hp", "Attack", "Defense", "Sp.Attack", "Sp.Defense", "Speed"}
}

func newPrimitive(text string, alignment int) tview.Primitive {
  return tview.NewTextView().
    SetTextAlign(alignment).
    SetText(text)
}


func (w *Window) NewLayout() {
  w.commandInput = tview.NewInputField().
    SetFieldBackgroundColor(tcell.NewRGBColor(50, 50, 70)).
    SetLabel("> ").
    SetDoneFunc(func (key tcell.Key) {
      if (key == tcell.KeyEnter){
        // Do the query
      }
    })

  // TODO: Incrementing Number of rows and RowSize interaction
  w.displayGrid = tview.NewGrid().SetSize(100, 0, entryHeight, 0)
  w.displayGrid.Box.SetBorder(true).SetTitle("Poketch").
    // Don't allow side scrolling
    SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
      if(event.Key() == tcell.KeyRight || event.Key() == tcell.KeyLeft) {
        return nil
      }
      return event
    })

  w.layout = tview.NewFlex().
  SetDirection(tview.FlexRow).
  AddItem(w.displayGrid, 0, 9, true).
  AddItem(w.commandInput, 0, 1, false)
}

func (w *Window) GetLayout() (*tview.Flex){
  return w.layout 
}

func newPokemonGrid() *tview.Grid {
  return nil
}

func newStatGrid(sampleStats [6]int) *tview.Grid {
  // Base Stat Grid
  // Medium/Large Screen ratio - 2:2:6
  // Small Screen ratio - 1:1:0
  var statNames [numberStats]string = getStatNames()

  statGrid := tview.NewGrid()
  barGrid := tview.NewGrid().SetMinSize(0, 1)
  
  baseStatTotal := 0
  for i := 0; i < numberStats; i++ {
    // Bars for each stat
    barGrid.AddItem(
      tview.NewBox().
      SetBorder(true).
      SetBorderColor(tcell.ColorBlack).
      SetBackgroundColor(tcell.ColorWhite), 
      i, 0, 1, sampleStats[i]/10, 0, 0, false)

    // Stat Names
    statGrid.AddItem(newPrimitive(statNames[i] + string(':'), tview.AlignLeft), i, 0, 1, smallStatLeft, 0, 0, false)
    statGrid.AddItem(newPrimitive(statNames[i] + string(':'), tview.AlignLeft), i, 0, 1, mediumStatLeft, 0, mediumScreenSize, false)
    // Stat Numbers
    statGrid.AddItem(newPrimitive(strconv.Itoa(sampleStats[i]), tview.AlignCenter), i, 1, 1, smallStatMiddle, 0, 0, false)
    statGrid.AddItem(newPrimitive(strconv.Itoa(sampleStats[i]), tview.AlignCenter), i, 1, 1, mediumStatMiddle, 0, mediumScreenSize, false)
    baseStatTotal += sampleStats[i]
  }

  statGrid.AddItem(newPrimitive("Total:", tview.AlignLeft), numberStats, 0, 1, smallStatLeft, 0, 0, false)
  statGrid.AddItem(newPrimitive("Total:", tview.AlignLeft), numberStats, 0, 1, mediumStatLeft, 0, mediumScreenSize, false)
  statGrid.AddItem(newPrimitive(strconv.Itoa(baseStatTotal), tview.AlignCenter), numberStats, 1, 1, smallStatMiddle, 0, 0, false)
  statGrid.AddItem(newPrimitive(strconv.Itoa(baseStatTotal), tview.AlignCenter), numberStats, 1, 1, mediumStatMiddle, 0, mediumScreenSize, false)
  
  statGrid.AddItem(barGrid, 0, 2, 6, smallStatRight, 0, 0, false)
  statGrid.AddItem(barGrid, 0, 2, 6, mediumStatRight, 0, mediumScreenSize, false)
  
  statGrid.SetBorder(true)

  return statGrid
}

func NewPokemonEntry() PokemonEntry {
  var pkm pokeapi.Pokemon = pokeapi.NewPokemon()
  pkm.Name = "pepe"

  sampleStats := [numberStats]int{80, 82, 83, 100, 100, 80}

  sampleLeftPanel := newPrimitive("left", tview.AlignCenter)
  sampleRightPanel := newStatGrid(sampleStats)

  return PokemonEntry {
    layout   :  tview.NewGrid().
      AddItem(sampleLeftPanel, 0, 0, 1, 2, 0, 0, false).
      AddItem(sampleLeftPanel, 0, 0, 1, 1, 0, mediumScreenSize*2, false).
      AddItem(sampleRightPanel, 0, 2, 1, 1, 0, 0, false).
      AddItem(sampleRightPanel, 0, 1, 1, 1, 0, mediumScreenSize*2, false),

  }
}

func (w *Window) AddEntry(newEntry PokemonEntry){
  w.displayGrid.AddItem(newEntry.layout, 0, 0, 1, 5, 0, 0, false)
}

func Start(w Window) error {
  app := tview.NewApplication()
  app.SetRoot(w.GetLayout(), true).
    SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
      if event.Key() == tcell.KeyRune && event.Rune() == '!' {
        app.SetFocus(w.commandInput)
        return nil
      } else if (event.Key() == tcell.KeyEsc){
        if !w.displayGrid.HasFocus() {
          app.SetFocus(w.displayGrid)
        }
        return nil
      }
      return event
    })

  if err := app.Run(); err != nil {
    return err
  }
  return nil
}


