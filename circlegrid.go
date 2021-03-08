package generativeart

import (
	"github.com/fogleman/gg"
	"image/color"
	"math"
	"math/rand"
)

type circleGrid struct {
	circleNumMin, circleNumMax int
}

// NewCircleGrid returns a circleGrid object.
func NewCircleGrid(circleNumMin, circleNumMax int) *circleGrid {
	return &circleGrid{
		circleNumMin: circleNumMin,
		circleNumMax: circleNumMax,
	}
}

// Generative draws a circle grid image.
func (cg *circleGrid) Generative(c *canva) {
	ctex := gg.NewContextForRGBA(c.img)
	cg.grid(ctex, c)
	ctex.Translate(float64(c.width)/2, float64(c.height)/2)
	ctex.Scale(0.9, 0.9)
	ctex.Translate(-float64(c.width)/2, -float64(c.height)/2)

	seg := RandomRangeInt(cg.circleNumMin, cg.circleNumMax)
	w := float64(c.width) / float64(seg)

	for i := 0; i < seg; i++ {
		for j := 0; j < seg; j++ {
			x := float64(i)*w + w/2
			y := float64(j)*w + w/2
			ctex.SetColor(c.opts.colorSchema[rand.Intn(len(c.opts.colorSchema))])
			ctex.DrawCircle(x, y, w/2*RandomRangeFloat64(0.1, 0.5))
			ctex.Fill()
			cg.draw(ctex, c, x, y, w/2*RandomRangeFloat64(0.6, 0.95))
		}
	}
}

func (cg *circleGrid) draw(ctex *gg.Context, c *canva, x, y, r float64) {
	rnd := rand.Intn(4)
	col := c.opts.colorSchema[rand.Intn(len(c.opts.colorSchema))]
	ctex.Push()
	ctex.Translate(x, y)
	ctex.Rotate(float64(rand.Intn(10)))
	ctex.SetColor(col)
	ctex.SetLineWidth(c.opts.lineWidth)

	switch rnd {
	case 0:
		ctex.DrawCircle(0, 0, r)
		ctex.Stroke()
	case 1:
		n := RandomRangeInt(1, 4) * 2
		ctex.DrawCircle(0, 0, r)
		ctex.Stroke()
		for i := 0; i < n; i++ {
			ctex.Rotate(math.Pi * 2 / float64(n))
			ctex.DrawCircle(r, 0, r*0.1)
			ctex.Fill()
		}
	case 2:
		n := RandomRangeInt(8, 20)
		theta := math.Pi * 0.5 * float64(RandomRangeInt(1, 5))
		for i := 0; i < n; i++ {
			d := float64(i) / float64(n)
			if d > r*0.1 {
				d = r * 0.1
			}
			ctex.Rotate(theta / float64(n))
			ctex.DrawCircle(r/2, 0, d*2)
			ctex.Fill()
		}
	case 3:
		n := RandomRangeInt(5, 20)
		for i := 0; i < n; i++ {
			ctex.Rotate(math.Pi * 2 / float64(n))
			ctex.DrawLine(r/2, 0, (r*2/3)-(r*0.05), 0)
			ctex.Stroke()
		}

	}
	ctex.Pop()
}

func (cg *circleGrid) grid(ctex *gg.Context, c *canva) {
	var segment int = 100
	w := float64(c.width) / float64(segment)

	ctex.SetColor(color.RGBA{255, 255, 255, 255})
	ctex.SetLineWidth(0.6)
	for i := 0; i < segment; i++ {
		ctex.DrawLine(0, float64(i)*w, float64(c.width), float64(i)*w)
		ctex.Stroke()
	}

	for j := 0; j < segment; j++ {
		ctex.DrawLine(float64(j)*w, 0, float64(j)*w, float64(c.height))
		ctex.Stroke()
	}
}