
package cli


import (
  "github.com/TheTatsujin/poketch/mod/apifetch/pokeapi"
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)

const (
  // Pokemon Stat Bar
  statBarMaxLen = 20
  statGridMaxLen = 12


  rowSize = 16

  numberStats = 6

)

type Window struct {
  displayGrid   *tview.Grid
  commandInput  *tview.InputField
  layout        *tview.Flex
}


type PokemonEntry struct {
  statGrid *tview.Grid
  layout *tview.Grid
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
  w.displayGrid = tview.NewGrid().SetSize(100, 0, rowSize, 0)
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

func NewPokemonEntry() PokemonEntry {
  var pkm pokeapi.Pokemon = pokeapi.NewPokemon()
  pkm.Name = "pepe"
  e := PokemonEntry {

    statGrid : tview.NewGrid().SetColumns(12, 40, 12).
      // Stat Name
      AddItem(newPrimitive("Hp:", tview.AlignLeft), 0, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Attack:", tview.AlignLeft), 1, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Defense:", tview.AlignLeft), 2, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Sp.Attack:", tview.AlignLeft), 3, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Sp.Defense:", tview.AlignLeft), 4, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Speed:", tview.AlignLeft), 5, 0, 1, 1, 2, 1, false).
      AddItem(newPrimitive("Total:", tview.AlignLeft), 6, 0, 1, 1, 2, 1, false),

    layout  : tview.NewGrid().
      AddItem(newPrimitive("left", tview.AlignCenter), 0, 0, 1, 1, 1, 1, false).
      AddItem(newPrimitive("right", tview.AlignCenter), 0, 2, 1, 1, 1, 1, false),
  }
  
  // Stat Bars
  e.statGrid.
    AddItem(tview.NewGrid().
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 0, 0, 1, statBarMaxLen, 2, 10, false).
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 1, 0, 1, statBarMaxLen/4, 2, 10, false).
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 2, 0, 1, 8, 2, 10, false).
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 3, 0, 1, 2, 2, 10, false).
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 4, 0, 1, 5, 2, 10, false).
      AddItem(tview.NewBox().SetBorder(true).SetBorderColor(tcell.ColorBlack).SetBackgroundColor(tcell.ColorWhite), 5, 0, 1, 16, 2, 10, false),
    0, 1, numberStats, 1, numberStats, statBarMaxLen, false)
  
  // Stat Number
  e.statGrid.
    AddItem(newPrimitive("100", tview.AlignCenter), 0, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("100", tview.AlignCenter), 1, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("100", tview.AlignCenter), 2, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("100", tview.AlignCenter), 3, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("100", tview.AlignCenter), 4, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("100", tview.AlignCenter), 5, 2, 1, 1, 2, 1, false).
    AddItem(newPrimitive("420", tview.AlignLeft), 6, 1, 1, 1, 2, 1, false)

  e.statGrid.SetBorder(true)
  e.layout.AddItem(e.statGrid, 0, 1, 1, 1, 1, 1, false)
  return e
}

func (w *Window) AddEntry(newEntry PokemonEntry){
  w.displayGrid.AddItem(newEntry.layout, 0, 0, 1, 1, 1, 1, false)
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


