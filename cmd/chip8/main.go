package main

import (
	"fmt"
	"github.com/NickDeChip/chip8-go/pkg/cpu"
	rl "github.com/gen2brain/raylib-go/raylib"
)

const scale = 10

var romLoaded = ""
var chosenGame = uint8(0)

func main() {
	gamePicker()

	FPSCap := true

	rl.InitWindow(64*scale, 32*scale, "Chip8++ | FPS: 0")
	rl.InitAudioDevice()

	chip8 := cpu.New()
	chip8.LoadFontsetIntoMemory()

	chip8.LoadFileIntoMemory(romLoaded)

	beep := rl.LoadMusicStream("./assets/beep.wav")

	isBeeping := false

	for !rl.WindowShouldClose() {
		rl.SetWindowTitle(fmt.Sprintf("Chip8++ | FPS: %v", rl.GetFPS()))

		if rl.IsKeyPressed(rl.KeyFive) {
			FPSCap = !FPSCap
		}

		if FPSCap && rl.IsKeyPressed(rl.KeyFive) {
			rl.SetTargetFPS(200)
		} else if rl.IsKeyPressed(rl.KeyFive) {
			rl.SetTargetFPS(60)
		}

		handleKeyPress(chip8)
		handleKeyUp(chip8)

		chip8.UpdateTimers(rl.GetFrameTime())
		chip8.HandleOpcode(chip8.GetCurrentOpcode())

		rl.UpdateMusicStream(beep)

		if chip8.SoundTimer > 0 {
			timePlayed := rl.GetMusicTimePlayed(beep) / rl.GetMusicTimeLength(beep)
			if timePlayed >= 1 {
				rl.StopMusicStream(beep)
				isBeeping = false
			}
			if !isBeeping {
				rl.PlayMusicStream(beep)
				isBeeping = true
			}
		} else if isBeeping {
			rl.StopMusicStream(beep)
			isBeeping = false
		}

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
	rl.UnloadMusicStream(beep)
	rl.CloseAudioDevice()
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

func gamePicker() {
	fmt.Println("1 - 15PUZZLE")
	fmt.Println("2 - BLINKY")
	fmt.Println("3 - BREAKOUT")
	fmt.Println("4 - BRIX")
	fmt.Println("5 - CONNECT4")
	fmt.Println("6 - GUESS")
	fmt.Println("7 - HIDDEN")
	fmt.Println("8 - INVADERS")
	fmt.Println("9 - KALEID")
	fmt.Println("10 - MAZE")
	fmt.Println("11 - MERLIN")
	fmt.Println("12 - MISSILE")
	fmt.Println("13 - PONG")
	fmt.Println("14 - PONG2")
	fmt.Println("15 - PUZZLE")
	fmt.Println("16 - SQUASH")
	fmt.Println("17 - SYZYGY")
	fmt.Println("18 - TANK")
	fmt.Println("19 - TETRIS")
	fmt.Println("20 - TICTAC")
	fmt.Println("21 - UFO")
	fmt.Println("22 - VBRIX")
	fmt.Println("23 - WALL")
	fmt.Println("24 - WIPEOFF")
	fmt.Println("Input the number for the game you would like to play:")
	fmt.Scan(&chosenGame)

	switch chosenGame {
	case 1:
		romLoaded = "./roms/15PUZZLE"
	case 2:
		romLoaded = "./roms/BLINKY"
	case 3:
		romLoaded = "./roms/BREAKOUT"
	case 4:
		romLoaded = "./roms/BRIX"
	case 5:
		romLoaded = "./roms/CONNECT4"
	case 6:
		romLoaded = "./roms/GUESS"
	case 7:
		romLoaded = "./roms/HIDDEN"
	case 8:
		romLoaded = "./roms/INVADERS"
	case 9:
		romLoaded = "./roms/KALEID"
	case 10:
		romLoaded = "./roms/MAZE"
	case 11:
		romLoaded = "./roms/MERLIN"
	case 12:
		romLoaded = "./roms/MISSILE"
	case 13:
		romLoaded = "./roms/PONG"
	case 14:
		romLoaded = "./roms/PONG2"
	case 15:
		romLoaded = "./roms/PUZZLE"
	case 16:
		romLoaded = "./roms/SQUASH"
	case 17:
		romLoaded = "./roms/SYZYGY"
	case 18:
		romLoaded = "./roms/TANK"
	case 19:
		romLoaded = "./roms/TETRIS"
	case 20:
		romLoaded = "./roms/TICTAC"
	case 21:
		romLoaded = "./roms/UFO"
	case 22:
		romLoaded = "./roms/VBRIX"
	case 23:
		romLoaded = "./roms/WALL"
	case 24:
		romLoaded = "./roms/WIPEOFF"
	}
}
