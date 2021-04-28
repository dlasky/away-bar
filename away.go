package main

import (
	"log"

	"github.com/dlasky/gotk3-layershell/layershell"
	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	css, err := gtk.CssProviderNew()
	if err != nil {
		log.Fatal("unable to create css provider")
	}
	err = css.LoadFromPath("./style.css") //TODO: xdg config
	if err != nil {
		log.Fatal(err)
	}
	screen, err := gdk.ScreenGetDefault()
	if err != nil {
		log.Fatal()
	}
	gtk.AddProviderForScreen(screen, css, gtk.STYLE_PROVIDER_PRIORITY_APPLICATION)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	layershell.InitForWindow(win)

	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_LEFT, true)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_TOP, true)
	layershell.SetAnchor(win, layershell.LAYER_SHELL_EDGE_RIGHT, true)

	layershell.SetLayer(win, layershell.LAYER_SHELL_LAYER_BOTTOM)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_TOP, 0)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_LEFT, 0)
	layershell.SetMargin(win, layershell.LAYER_SHELL_EDGE_RIGHT, 0)

	layershell.AutoExclusiveZoneEnable(win)

	b, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 5)
	if err != nil {
		log.Fatal("ui error")
	}

	ws, err := InitWorkspaces()
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(ws)

	clock, err := InitClock("3:04 PM")
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(clock)

	cpu, err := InitCPU()
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(cpu)
	mem, err := InitMem()
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(mem)

	bat, err := InitBattery()
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(bat)

	vol, err := InitPulseAudio()
	if err != nil {
		log.Fatal("ui error")
	}
	b.Add(vol)

	win.Add(b)
	// Set the default window size.
	win.SetDefaultSize(800, 30)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.

	gtk.Main()
}
