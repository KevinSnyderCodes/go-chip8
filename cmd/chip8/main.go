package main

import (
	"flag"
	"fmt"
	"image/color"
	"log"
	"os"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"github.com/faiface/pixel/pixelgl"
	"github.com/kevinsnydercodes/go-chip8/internal/chip8"
)

var (
	fProgramFilename = flag.String("program-filename", "", "")
)

func runE() error {
	flag.Parse()

	cfg := pixelgl.WindowConfig{
		Title:  "Chip-8",
		Bounds: pixel.R(0, 0, 640, 320),
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		return fmt.Errorf("error creating window: %w", err)
	}

	program, err := os.ReadFile(*fProgramFilename)
	if err != nil {
		return fmt.Errorf("error reading program file: %w", err)
	}

	Chip8 := chip8.Chip8{}
	Chip8.Reset(program)

	for !win.Closed() {
		if err := Chip8.Tick(); err != nil {
			return fmt.Errorf("tick error: %w", err)
		}

		display := Chip8.DumpDisplay()

		win.Clear(color.RGBA{R: 80, G: 69, B: 155})
		for x := 0; x < 64; x++ {
			for y := 0; y < 32; y++ {
				if display[x][y] {
					imd := imdraw.New(nil)
					imd.Color = color.RGBA{R: 136, G: 126, B: 203}
					imd.Push(pixel.V(float64(x*10), float64(320-(y*10))))
					imd.Push(pixel.V(float64((x+1)*10), float64(320-((y+1)*10))))
					imd.Rectangle(0)
					imd.Draw(win)
				}
			}
		}

		win.Update()
	}

	return nil
}

func run() {
	if err := runE(); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	pixelgl.Run(run)
}
