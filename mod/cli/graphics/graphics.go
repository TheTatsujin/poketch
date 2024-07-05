package graphics


import (
  "math/rand"
  "github.com/gdamore/tcell/v2"
  "github.com/rivo/tview"
)



func DrawText(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
  row := y1
  col := x1

  for _, r := range []rune(text) {
    s.SetContent(col, row, r, nil, style)
    col++
    if col >= x2 {
      row++
      col = x1
    }
    if row > y2 {
      break
    }
  }
}


func DrawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
	if y2 < y1 {
		y1, y2 = y2, y1
	}
	if x2 < x1 {
		x1, x2 = x2, x1
	}

	// Fill background
	for row := y1; row <= y2; row++ {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	// Draw borders
	for col := x1; col <= x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
	}

	// Only draw corners if necessary
	if y1 != y2 && x1 != x2 {
		s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
		s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
		s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
		s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
	}

	DrawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

// Using tview

func makebox(s tcell.Screen) {
	w, h := s.Size()

	if w == 0 || h == 0 {
		return
	}

	glyphs := []rune{'@', '#', '&', '*', '=', '%', 'Z', 'A'}

	lx := rand.Int() % w
	ly := rand.Int() % h
	lw := rand.Int() % (w - lx)
	lh := rand.Int() % (h - ly)
	st := tcell.StyleDefault
	gl := ' '
	if s.Colors() > 256 {
		rgb := tcell.NewHexColor(int32(rand.Int() & 0xffffff))
		st = st.Background(rgb)
	} else if s.Colors() > 1 {
		st = st.Background(tcell.Color(rand.Int() % s.Colors()))
	} else {
		st = st.Reverse(rand.Int()%2 == 0)
		gl = glyphs[rand.Int()%len(glyphs)]
	}

	for row := 0; row < lh; row++ {
		for col := 0; col < lw; col++ {
			s.SetCell(lx+col, ly+row, st, gl)
		}
	}
	s.Show()
}


func NewLayout() error {
  app := tview.NewApplication()
  
  var commandInput = tview.NewInputField().
    SetFieldBackgroundColor(tcell.NewRGBColor(50, 50, 80)).
    SetLabel("> ").
    SetDoneFunc(func(key tcell.Key){
      // Do whatever
      if key == tcell.KeyEnter {
        commandInput.GetText()
      }
    })  


  var displayBox = tview.NewBox().SetBorder(true)
  screen, err := tcell.NewScreen()
  if err != nil {
    return err
  }

  makebox(screen)

  displayBox.Draw(screen)
  var flexLayout = tview.NewFlex().
    AddItem(displayBox, 0, 6, false).
    SetDirection (tview.FlexRow).
    AddItem(commandInput, 0, 1, true)

  
 
  if err = app.SetRoot(flexLayout, true).Run(); err != nil {
    return err
  }
  return nil
}



