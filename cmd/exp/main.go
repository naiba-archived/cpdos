package main

import (
	"fmt"

	"github.com/andlabs/ui"
	_ "github.com/andlabs/ui/winmanifest"
	"github.com/naiba/cpdos/internal"
)

var (
	methods  []string
	mainwin  *ui.Window
	etURL    *ui.Entry
	meBefore *ui.MultilineEntry
	meAfter  *ui.MultilineEntry
)

func init() {
	methods = []string{
		"HTTP Header Oversize (HHO)",
		"HTTP Meta Character (HMC)",
		"HTTP Method Override (HMO)",
	}
}

func setupUI() {
	mainwin = ui.NewWindow("CPDoS Test Tool", 300, 200, true)
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
	etURL = ui.NewEntry()
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
	btnVerify.OnClicked(onClick)
	vbox.Append(btnVerify, false)

	fmBefore := ui.NewForm()
	fmBefore.SetPadded(true)
	meBefore = ui.NewMultilineEntry()
	meBefore.SetReadOnly(true)
	fmBefore.Append("Before", meBefore, true)
	vbox.Append(fmBefore, true)

	fmAfter := ui.NewForm()
	fmAfter.SetPadded(true)
	meAfter = ui.NewMultilineEntry()
	meAfter.SetReadOnly(true)
	fmAfter.Append("   After", meAfter, true)
	vbox.Append(fmAfter, true)

	mainwin.Show()
}

func onClick(b *ui.Button) {
	b.Disable()
	exp := internal.NewCPDoSExp(etURL.Text())
	go func() {
		body, status := exp.Get()
		str := fmt.Sprintf("[%d]%s\n", status, body)
		ui.QueueMain(func() {
			meBefore.Append(str)
		})
		body, status = exp.HHO(250)
		str = fmt.Sprintf("[%d]%s\n", status, body)
		ui.QueueMain(func() {
			meAfter.Append(str)
			b.Enable()
		})
	}()
}

func main() {
	ui.Main(setupUI)
}
