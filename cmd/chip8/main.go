package main

import (
	"fmt"
	"github.com/NickDeChip/chip8-go/pkg/cpu"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const scale = 10

func main() {
	rl.InitWindow(64*scale, 32*scale, "Chip8++ | FPS: 0")
	rl.SetTargetFPS(120)

	chip8 := cpu.New()
	chip8.LoadFontsetIntoMemory()

	chip8.LoadFileIntoMemory("./roms/BREAKOUT")

	for !rl.WindowShouldClose() {
		rl.SetWindowTitle(fmt.Sprintf("Chip8++ | FPS: %v", rl.GetFPS()))

		chip8.UpdateTimers(rl.GetFrameTime())
		chip8.HandleOpcode(chip8.GetCurrentOpcode())

		rl.BeginDrawing()
		if chip8.ShouldClearScreen {
			rl.ClearBackground(rl.Black)
		}
		if chip8.ShouldDrawScreen {
			for y := 0; y < len(chip8.GFX); y++ {
				for x := 0; x < len(chip8.GFX[x]); x++ {
					if chip8.GFX[y][x] != 0 {
						rl.DrawRectangle(int32(x)*scale, int32(y)*scale, scale, scale, rl.White)
					} else {
						rl.DrawRectangle(int32(x)*scale, int32(y)*scale, scale, scale, rl.Black)
					}
				}
			}
		}
		rl.EndDrawing()

	}
	rl.CloseWindow()
}
