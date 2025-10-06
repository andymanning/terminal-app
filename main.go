package main

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func main() {
	// Set up green-on-black monochrome theme
	tview.Styles.PrimitiveBackgroundColor = tcell.ColorBlack
	tview.Styles.ContrastBackgroundColor = tcell.ColorBlack
	tview.Styles.MoreContrastBackgroundColor = tcell.ColorBlack
	tview.Styles.PrimaryTextColor = tcell.ColorGreen
	tview.Styles.SecondaryTextColor = tcell.ColorGreen
	tview.Styles.TertiaryTextColor = tcell.ColorGreen
	tview.Styles.InverseTextColor = tcell.ColorBlack
	tview.Styles.ContrastSecondaryTextColor = tcell.ColorGreen

	app := tview.NewApplication()

	// Create the main layout
	mainFlex := tview.NewFlex().SetDirection(tview.FlexRow)

	// Create the menu bar
	menuBar := tview.NewFlex().SetDirection(tview.FlexColumn)

	// Create a blank main area
	mainArea := tview.NewBox().
		SetBorder(false).
		SetBackgroundColor(tcell.ColorBlack)

	// Store the original UI to return to after dropdown closes
	originalRoot := mainFlex

	// Create menu buttons
	fileButton := tview.NewButton("File").SetSelectedFunc(func() {
		showDropdown(app, "File", []string{"New", "Open", "Save", "Exit"}, originalRoot)
	})

	editButton := tview.NewButton("Edit").SetSelectedFunc(func() {
		showDropdown(app, "Edit", []string{"Cut", "Copy", "Paste"}, originalRoot)
	})

	helpButton := tview.NewButton("Help").SetSelectedFunc(func() {
		showDropdown(app, "Help", []string{"About", "Documentation"}, originalRoot)
	})

	// Add buttons to the menu bar
	menuBar.AddItem(fileButton, 10, 0, false)
	menuBar.AddItem(editButton, 10, 0, false)
	menuBar.AddItem(helpButton, 10, 0, false)
	menuBar.AddItem(tview.NewBox(), 0, 1, false) // Filler to push menus to the left

	// Style the menu bar
	menuBar.SetBackgroundColor(tcell.ColorBlack)
	fileButton.SetBackgroundColor(tcell.ColorBlack)
	fileButton.SetLabelColor(tcell.ColorGreen)
	editButton.SetBackgroundColor(tcell.ColorBlack)
	editButton.SetLabelColor(tcell.ColorGreen)
	helpButton.SetBackgroundColor(tcell.ColorBlack)
	helpButton.SetLabelColor(tcell.ColorGreen)

	// Add the menu bar and main area to the main layout
	mainFlex.AddItem(menuBar, 1, 0, false)
	mainFlex.AddItem(mainArea, 0, 1, true) // Give the main area all remaining space

	// Set up key handling for the application
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Key() == tcell.KeyEscape {
			// Return to the main UI when Escape is pressed
			app.SetRoot(originalRoot, true)
			return nil
		}
		return event
	})

	// Set the main flex as the root of the application
	if err := app.SetRoot(mainFlex, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

// showDropdown displays a dropdown menu with the given items
func showDropdown(app *tview.Application, title string, items []string, mainUI tview.Primitive) {
	// Create the dropdown menu
	dropdown := tview.NewList().
		SetMainTextColor(tcell.ColorGreen).
		SetSelectedTextColor(tcell.ColorBlack).
		SetSelectedBackgroundColor(tcell.ColorGreen)
	for i, item := range items {
		itemText := item // Create a new variable to avoid closure issues
		dropdown.AddItem(itemText, "", rune('a'+i), func() {
			// Handle menu item selection (just close the dropdown for now)
			app.SetRoot(mainUI, true)
		})
	}

	// Set up the dropdown position and appearance
	dropdown.SetBorder(true).
		SetTitle(title).
		SetBackgroundColor(tcell.ColorBlack).
		SetTitleColor(tcell.ColorGreen).
		SetBorderColor(tcell.ColorGreen)

	// Create a flex to position the dropdown
	dropdownFlex := tview.NewFlex().
		AddItem(nil, 0, 1, false).
		AddItem(
			tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 1, 0, false).
				AddItem(dropdown, len(items)+2, 0, true).
				AddItem(nil, 0, 1, false),
			25, 0, true).
		AddItem(nil, 0, 1, false)

	// Show the dropdown
	app.SetRoot(dropdownFlex, true)
	app.SetFocus(dropdown)
}
