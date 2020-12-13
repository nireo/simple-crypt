package main

import (
	"fmt"
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {
	gtk.Init(nil)

	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("GTK Go Input Dialog example!")
	window.SetIconName("gtk-dialog-info")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "Happy coding!")

	dialog := gtk.NewDialog()
	dialog.SetTitle("User input")

	vbox := dialog.GetVBox()

	label := gtk.NewLabel("Enter some characters here :")
	vbox.Add(label)

	input := gtk.NewEntry()
	input.SetEditable(true)
	vbox.Add(input)

	button := gtk.NewButtonWithLabel("OK")
	button.Connect("clicked", func() {
		fmt.Println("Input : ", input.GetText())
		//gtk.MainQuit()
	})

	vbox.Add(button)

	quitButton := gtk.NewButtonWithLabel("Quit")
	quitButton.Connect("clicked", func() {
		fmt.Println("Quiting now ...")
		gtk.MainQuit()
	})

	vbox.Add(quitButton)

	dialog.ShowAll()
	window.Add(dialog)
	gtk.Main()
}
