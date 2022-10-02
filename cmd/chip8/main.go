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

		handleKeyPress(chip8)
		handleKeyUp(chip8)

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

func handleKeyPress(chip8 *cpu.CPU) {
	if rl.IsKeyDown(rl.KeyOne) {
		chip8.Key[0] = 1
	}
	if rl.IsKeyDown(rl.KeyTwo) {
		chip8.Key[1] = 1
	}
	if rl.IsKeyDown(rl.KeyThree) {
		chip8.Key[2] = 1
	}
	if rl.IsKeyDown(rl.KeyFour) {
		chip8.Key[3] = 1
	}
	if rl.IsKeyDown(rl.KeyQ) {
		chip8.Key[4] = 1
	}
	if rl.IsKeyDown(rl.KeyW) {
		chip8.Key[5] = 1
	}
	if rl.IsKeyDown(rl.KeyE) {
		chip8.Key[6] = 1
	}
	if rl.IsKeyDown(rl.KeyR) {
		chip8.Key[7] = 1
	}
	if rl.IsKeyDown(rl.KeyA) {
		chip8.Key[8] = 1
	}
	if rl.IsKeyDown(rl.KeyS) {
		chip8.Key[9] = 1
	}
	if rl.IsKeyDown(rl.KeyD) {
		chip8.Key[0xA] = 1
	}
	if rl.IsKeyDown(rl.KeyF) {
		chip8.Key[0xB] = 1
	}
	if rl.IsKeyDown(rl.KeyZ) {
		chip8.Key[0xC] = 1
	}
	if rl.IsKeyDown(rl.KeyX) {
		chip8.Key[0xD] = 1
	}
	if rl.IsKeyDown(rl.KeyC) {
		chip8.Key[0xE] = 1
	}
	if rl.IsKeyDown(rl.KeyV) {
		chip8.Key[0xF] = 1
	}
}

func handleKeyUp(chip8 *cpu.CPU) {
	if rl.IsKeyUp(rl.KeyOne) {
		chip8.Key[0] = 0
	}
	if rl.IsKeyUp(rl.KeyTwo) {
		chip8.Key[1] = 0
	}
	if rl.IsKeyUp(rl.KeyThree) {
		chip8.Key[2] = 0
	}
	if rl.IsKeyUp(rl.KeyFour) {
		chip8.Key[3] = 0
	}
	if rl.IsKeyUp(rl.KeyQ) {
		chip8.Key[4] = 0
	}
	if rl.IsKeyUp(rl.KeyW) {
		chip8.Key[5] = 0
	}
	if rl.IsKeyUp(rl.KeyE) {
		chip8.Key[6] = 0
	}
	if rl.IsKeyUp(rl.KeyR) {
		chip8.Key[7] = 0
	}
	if rl.IsKeyUp(rl.KeyA) {
		chip8.Key[8] = 0
	}
	if rl.IsKeyUp(rl.KeyS) {
		chip8.Key[9] = 0
	}
	if rl.IsKeyUp(rl.KeyD) {
		chip8.Key[0xA] = 0
	}
	if rl.IsKeyUp(rl.KeyF) {
		chip8.Key[0xB] = 0
	}
	if rl.IsKeyUp(rl.KeyZ) {
		chip8.Key[0xC] = 0
	}
	if rl.IsKeyUp(rl.KeyX) {
		chip8.Key[0xD] = 0
	}
	if rl.IsKeyUp(rl.KeyC) {
		chip8.Key[0xE] = 0
	}
	if rl.IsKeyUp(rl.KeyV) {
		chip8.Key[0xF] = 0
	}
}
