package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"strconv"
	"strings"
)

func main() {
	sudokuApp := app.New()
	sudokuWin := sudokuApp.NewWindow("Sudoku")

	grid := container.NewGridWithRows(9) // Use Rows to create 9 rows

	entries := make([]*widget.Entry, 81)

	for i := 0; i < 9; i++ {
		row := container.NewGridWithColumns(9) // Create a new row with 9 columns
		for j := 0; j < 9; j++ {
			entry := widget.NewEntry()
			entry.SetText("")
			entry.TextStyle = fyne.TextStyle{Bold: true} // Make text bold for better visibility.

			entry.OnChanged = func(s string) {
				value := strings.TrimSpace(s)
				if len(value) > 0 {
					if _, err := strconv.Atoi(value); err != nil || len(value) > 1 {
						entry.SetText("")
					}
				} else {
					entValue := "  " + value
					entry.SetText(entValue)
				}
			}

			entryContainer := container.New(layout.NewCenterLayout(), entry) // Center the entry

			entries[i*9+j] = entry
			row.Add(entryContainer)
		}
		grid.Add(row)
	}

	solveButton := widget.NewButton("Solve", func() {
		solveSudoku(entries)
	})

	content := container.NewVBox(grid, solveButton)

	sudokuWin.SetContent(content)
	sudokuWin.Resize(fyne.NewSize(400, 400))
	sudokuWin.CenterOnScreen()
	sudokuWin.ShowAndRun()
}

func solveSudoku(entries []*widget.Entry) {
}
