package controllers

import (
	"github.com/hculpan/go-scape/app/model"
	"github.com/hculpan/go-scape/app/pages"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/hculpan/go-sdl-lib/game"
	"github.com/veandco/go-sdl2/sdl"
)

type SugarscapeController struct {
	game.GameController
}

var Controller *SugarscapeController

func NewSugarscapeController(windowWidth, windowHeight int32) SugarscapeController {
	result := SugarscapeController{}

	windowBackground := sdl.Color{R: 0, G: 0, B: 0, A: 0}

	result.Game = model.NewSugarscape(100, 100)
	result.Window = component.NewWindow(windowWidth, windowHeight, "GoScape - an implementation of Sugarscape", windowBackground)

	result.RegisterPages()

	Controller = &result

	return result
}

func (s *SugarscapeController) RegisterPages() {
	component.RegisterPage(pages.NewGamePage("GamePage", 0, 0, s.Window.Width, s.Window.Height))
}
