package components

import (
	"github.com/hculpan/go-scape/app/model"
	"github.com/hculpan/go-sdl-lib/component"
	"github.com/veandco/go-sdl2/gfx"
	"github.com/veandco/go-sdl2/sdl"
)

type SugarscapeComponent struct {
	component.BaseComponent
}

func NewSugarscapeComponent(x, y, width, height int32) *SugarscapeComponent {
	result := &SugarscapeComponent{}

	result.SetPosition(x, y)
	result.SetSize(width, height)

	return result
}

func (c *SugarscapeComponent) DrawComponent(r *sdl.Renderer) error {
	r.SetDrawColor(255, 255, 255, 255)
	r.FillRect(&sdl.Rect{X: c.X, Y: c.Y, W: c.Width, H: c.Height})

	for x := 0; x < model.Sugarscape.BoardWidth; x++ {
		for y := 0; y < model.Sugarscape.BoardHeight; y++ {
			n := uint8(255 - (model.Sugarscape.Board[x][y].Sugar * 20))
			xloc := int32(x*8) + c.X
			yloc := int32(y*8) + c.Y
			var b uint8 = n / 4
			if n == 255 {
				b = 255
			}
			gfx.RoundedBoxRGBA(r, xloc+1, yloc+1, xloc+7, yloc+7, 2, n, n, b, 255)

			if model.Sugarscape.Board[x][y].Occupant != nil {
				gfx.RoundedBoxRGBA(r, xloc+2, yloc+2, xloc+6, yloc+6, 2, 255, 0, 0, 255)
			}
		}
	}

	return nil
}

func (c *SugarscapeComponent) Draw(r *sdl.Renderer) error {
	if err := component.DrawParentAndChildren(r, c); err != nil {
		return err
	}

	return nil
}
