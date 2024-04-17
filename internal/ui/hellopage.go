package ui

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type HelloComponent struct {
	app.Compo
}

func (h *HelloComponent) Render() app.UI {

	h1 := app.H1().Text("Welcome to some go wasm ui...")
	
	return h1
}

