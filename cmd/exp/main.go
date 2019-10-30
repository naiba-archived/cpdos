package main

import (
	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
)

var (
	methods []string
)

func init() {
	methods = []string{
		"HTTP Header Oversize (HHO)",
		"HTTP Meta Character (HMC)",
		"HTTP Method Override (HMO)",
	}
}

func setupUI() {
	mainwin := ui.NewWindow("CPDoS Test Tool", 300, 100, true)
	mainwin.SetMargined(true)
	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})
	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	mainwin.SetChild(vbox)

	urlBox := ui.NewHorizontalBox()
	urlBox.SetPadded(true)
	lbURL := ui.NewLabel("Website URL")
	etURL := ui.NewEntry()
	urlBox.Append(lbURL, false)
	urlBox.Append(etURL, true)
	vbox.Append(urlBox, false)

	methodBox := ui.NewHorizontalBox()
	methodBox.SetPadded(true)
	lbMethod := ui.NewLabel("CPDoS Type")
	cbMethod := ui.NewCombobox()
	for i := 0; i < len(methods); i++ {
		cbMethod.Append(methods[i])
	}
	cbMethod.SetSelected(0)
	methodBox.Append(lbMethod, false)
	methodBox.Append(cbMethod, true)
	vbox.Append(methodBox, false)

	btnVerify := ui.NewButton("Verify")
	btnVerify.OnClicked(func(btn *ui.Button) {
		ui.MsgBox(mainwin, "Hola", "You clicked the button!")
	})
	vbox.Append(btnVerify, false)

	mainwin.Show()
}

func main() {
	ui.Main(setupUI)
}
