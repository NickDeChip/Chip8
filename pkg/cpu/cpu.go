package cpu

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type CPU struct {
	PC uint16

	//Registers
	I uint16
	V [16]uint8

	//STACK
	Stack [16]uint16
	SP    uint8

	//Screen
	ShouldClearScreen bool
	ShouldDrawScreen  bool
	GFX               [32][64]uint8 // 64 x 32 screen size

	//Memory
	Memory [4096]uint8

	//Keys
	Key [16]uint8

	//Timers
	DelayTimer uint8
	SoundTimer uint8
	TickTimer  float32
}

func New() *CPU {
	return &CPU{
		PC:                0,
		I:                 0,
		V:                 [16]uint8{},
		Stack:             [16]uint16{},
		SP:                0,
		ShouldClearScreen: false,
		ShouldDrawScreen:  false,
		GFX:               [32][64]uint8{},
		Memory:            [4096]uint8{},
		Key:               [16]uint8{},
		DelayTimer:        0,
		SoundTimer:        0,
		TickTimer:         0.0,
	}
}

func (cpu *CPU) HandleOpcode(opcode uint16) {
	x, y := getXY(opcode)
	//fmt.Printf("PC: %x\n", cpu.PC)
	//fmt.Printf("Opcode: %x\n", opcode)

	cpu.PC += 2

	switch opcode & 0xF000 {
	case 0x0000:
		cpu.handle0Opcode(opcode)
	case 0x1000:
		cpu.op1NNN(opcode)
	case 0x2000:
		cpu.op2NNN(opcode)
	case 0x3000:
		cpu.op3XNN(opcode, x)
	case 0x4000:
		cpu.op4XNN(opcode, x)
	case 0x5000:
		cpu.op5XY0(opcode, x, y)
	case 0x6000:
		cpu.op6XNN(opcode, x)
	case 0x7000:
		cpu.op7XNN(opcode, x)
	case 0x8000:
		cpu.handle8opcode(opcode, x, y)
	case 0x9000:
		cpu.op09XY0(x, y)
	case 0xA000:
		cpu.opANNN(opcode)
	case 0xB000:
		cpu.opBNNN(opcode)
	case 0xC000:
		cpu.opCXNN(opcode, x)
	case 0xD000:
		cpu.opDXYN(opcode, x, y)
	case 0xE000:
		cpu.handleEopcodes(opcode, x)
	case 0xF000:
		cpu.handleFopcodes(opcode, x)
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handle0Opcode(opcode uint16) {
	switch opcode {
	case 0x00EE:
		cpu.op00EE()
	case 0x00E0:
		cpu.op00E0()
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handle8opcode(opcode uint16, x uint8, y uint8) {
	switch opcode & 0x000F {
	case 0x0000:
		cpu.op8XY0(x, y)
	case 0x0001:
		cpu.op8XY1(x, y)
	case 0x0002:
		cpu.op8XY2(x, y)
	case 0x0003:
		cpu.op8XY3(x, y)
	case 0x0004:
		cpu.op8XY4(x, y)
	case 0x0005:
		cpu.op8XY5(x, y)
	case 0x0006:
		cpu.op8XY6(x)
	case 0x0007:
		cpu.op8XY7(x, y)
	case 0x000E:
		cpu.op8XYE(x)
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handleEopcodes(opcode uint16, x uint8) {
	switch opcode & 0x00FF {
	case 0x009E:
		cpu.opEX9E(x)
	case 0x00A1:
		cpu.opEXA1(x)
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handleFopcodes(opcode uint16, x uint8) {
	switch opcode & 0x00FF {
	case 0x0007:
		cpu.opFX07(x)
	case 0x000A:
		cpu.opFX0A(x)
	case 0x001E:
		cpu.opFX1E(x)
	case 0x0015:
		cpu.opFX15(x)
	case 0x0018:
		cpu.opFX18(x)
	case 0x0029:
		cpu.opFX29(x)
	case 0x0033:
		cpu.opFX33(x)
	case 0x0055:
		cpu.opFX55(x)
	case 0x0065:
		cpu.opFX65(x)
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) op00E0() {
	cpu.GFX = [32][64]uint8{}
	cpu.ShouldClearScreen = true
}

func (cpu *CPU) op00EE() {
	cpu.SP -= 1
	cpu.PC = cpu.Stack[cpu.SP]
}

func (cpu *CPU) op1NNN(opcode uint16) {
	cpu.PC = opcode & 0x0FFF
}

func (cpu *CPU) op2NNN(opcode uint16) {
	cpu.Stack[cpu.SP] = cpu.PC
	cpu.SP += 1

	cpu.PC = opcode & 0x0FFF
}

func (cpu *CPU) op3XNN(opcode uint16, x uint8) {
	NN := uint8(opcode & 0x00FF)
	if cpu.V[x] == NN {
		cpu.PC += 2
	}
}

func (cpu *CPU) op4XNN(opcode uint16, x uint8) {
	NN := uint8(opcode & 0x00FF)
	if cpu.V[x] != NN {
		cpu.PC += 2
	}
}

func (cpu *CPU) op5XY0(opcode uint16, x uint8, y uint8) {
	if cpu.V[x] == cpu.V[y] {
		cpu.PC += 2
	}
}

func (cpu *CPU) op6XNN(opcode uint16, x uint8) {
	NN := uint8(opcode & 0x00FF)
	cpu.V[x] = NN
}

func (cpu *CPU) op7XNN(opcode uint16, x uint8) {
	NN := uint8(opcode & 0x00FF)
	cpu.V[x] += NN
}

func (cpu *CPU) op8XY0(x uint8, y uint8) {
	cpu.V[x] = cpu.V[y]
}

func (cpu *CPU) op8XY1(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] | cpu.V[y]
}

func (cpu *CPU) op8XY2(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] & cpu.V[y]
}

func (cpu *CPU) op8XY3(x uint8, y uint8) {
	cpu.V[x] = cpu.V[x] ^ cpu.V[y]
}

func (cpu *CPU) op8XY4(x uint8, y uint8) {
	cpu.V[x] += cpu.V[y]
	overFlowCheck := uint16(cpu.V[x]) + uint16(cpu.V[y])
	if overFlowCheck >= 255 {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}
}

func (cpu *CPU) op8XY5(x uint8, y uint8) {
	if cpu.V[x] > cpu.V[y] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}

	cpu.V[x] -= cpu.V[y]
}

func (cpu *CPU) op8XY6(x uint8) {
	cpu.V[0xF] = cpu.V[x] & 0x1
	cpu.V[x] >>= 1
}

func (cpu *CPU) op8XY7(x uint8, y uint8) {
	if cpu.V[y] > cpu.V[x] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}

	cpu.V[x] = cpu.V[y] - cpu.V[x]
}

func (cpu *CPU) op8XYE(x uint8) {
	cpu.V[0xF] = cpu.V[x] & 0x80
	cpu.V[x] <<= 1
}

func (cpu *CPU) op09XY0(x uint8, y uint8) {
	if cpu.V[x] != cpu.V[y] {
		cpu.PC += 2
	}
}

func (cpu *CPU) opANNN(opcode uint16) {
	NNN := uint16(opcode & 0x0FFF)
	cpu.I = NNN
}

func (cpu *CPU) opBNNN(opcode uint16) {
	NNN := uint16(opcode & 0x0FFF)
	cpu.PC = NNN + uint16(cpu.V[0])
}

func (cpu *CPU) opCXNN(opcode uint16, x uint8) {
	NN := uint16(opcode & 0x00FF)
	rand.Seed(time.Now().UnixNano())
	randInt := rand.Int31n(255)

	cpu.V[x] = (uint8(randInt)) & uint8(NN)
}

func (cpu *CPU) opDXYN(opcode uint16, x uint8, y uint8) {
	height := opcode & 0x000F

	cpu.V[0xF] = 0
	for yLine := uint16(0); yLine < height; yLine++ {
		pixel := cpu.Memory[cpu.I+yLine]
		for xLine := uint16(0); xLine < 8; xLine++ {
			if (pixel & (0x80 >> xLine)) != 0 {
				row := (cpu.V[y] + uint8(yLine)) % 32
				col := (cpu.V[x] + uint8(xLine)) % 64
				if cpu.GFX[row][col] == 1 {
					cpu.V[0xF] = 1
				}
				cpu.GFX[row][col] ^= 1
			}
		}
	}
	cpu.ShouldDrawScreen = true
}

func (cpu *CPU) opEX9E(x uint8) {
	if cpu.Key[cpu.V[x]] != 0 {
		cpu.PC += 2
	}
}

func (cpu *CPU) opEXA1(x uint8) {
	if cpu.Key[cpu.V[x]] == 0 {
		cpu.PC += 2
	}
}

func (cpu *CPU) opFX07(x uint8) {
	cpu.V[x] = cpu.DelayTimer
}

func (cpu *CPU) opFX0A(x uint8) {
	for i, key := range cpu.Key {
		if key != 0 {
			cpu.V[x] = uint8(i)
			goto keyPressed
		}
	}
	cpu.PC -= 2
keyPressed:
}

func (cpu *CPU) opFX15(x uint8) {
	cpu.DelayTimer = cpu.V[x]
}

func (cpu *CPU) opFX18(x uint8) {
	cpu.SoundTimer = cpu.V[x]
}

func (cpu *CPU) opFX1E(x uint8) {
	cpu.I += uint16(cpu.V[x])
}

func (cpu *CPU) opFX29(x uint8) {
	cpu.I = uint16(cpu.V[x] * 5)
}

func (cpu *CPU) opFX33(x uint8) {
	value := uint32(cpu.V[x])
	cpu.Memory[cpu.I+2] = uint8(value % 10)
	value /= 10
	cpu.Memory[cpu.I+1] = uint8(value % 10)
	value /= 10
	cpu.Memory[cpu.I] = uint8(value % 10)
}

func (cpu *CPU) opFX55(x uint8) {
	for i := uint8(0); i <= x; i++ {
		cpu.Memory[cpu.I+uint16(i)] = cpu.V[i]
	}
}

func (cpu *CPU) opFX65(x uint8) {
	for i := uint8(0); i <= x; i++ {
		cpu.V[i] = cpu.Memory[cpu.I+uint16(i)]
	}
}

// Get X and Y for cpu V`s
func getXY(opcode uint16) (x uint8, y uint8) {
	x = uint8(opcode & (0x0F00) >> 8)
	y = uint8(opcode & (0x00F0) >> 4)

	return x, y
}

func (cpu *CPU) UpdateTimers(dt float32) {
	cpu.TickTimer += dt
	if cpu.TickTimer > 0.0167 {
		cpu.TickTimer = 0
		if cpu.DelayTimer > 0 {
			cpu.DelayTimer -= 1
		}
		if cpu.SoundTimer > 0 {
			cpu.SoundTimer -= 1
		}
	}
}

func (cpu *CPU) LoadFontsetIntoMemory() {
	fontset := []uint8{
		0xF0, 0x90, 0x90, 0x90, 0xF0, // 0
		0x20, 0x60, 0x20, 0x20, 0x70, // 1
		0xF0, 0x10, 0xF0, 0x80, 0xF0, // 2
		0xF0, 0x10, 0xF0, 0x10, 0xF0, // 3
		0x90, 0x90, 0xF0, 0x10, 0x10, // 4
		0xF0, 0x80, 0xF0, 0x10, 0xF0, // 5
		0xF0, 0x80, 0xF0, 0x90, 0xF0, // 6
		0xF0, 0x10, 0x20, 0x40, 0x40, // 7
		0xF0, 0x90, 0xF0, 0x90, 0xF0, // 8
		0xF0, 0x90, 0xF0, 0x10, 0xF0, // 9
		0xF0, 0x90, 0xF0, 0x90, 0x90, // A
		0xE0, 0x90, 0xE0, 0x90, 0xE0, // B
		0xF0, 0x80, 0x80, 0x80, 0xF0, // C
		0xE0, 0x90, 0x90, 0x90, 0xE0, // D
		0xF0, 0x80, 0xF0, 0x80, 0xF0, // E
		0xF0, 0x80, 0xF0, 0x80, 0x80, // F
	}

	copy(cpu.Memory[:], fontset)
	cpu.PC = 0x200
}

func (cpu *CPU) LoadFileIntoMemory(filePath string) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	for i, byte := range data {
		cpu.Memory[0x200+uint16(i)] = byte
	}
}

func (cpu *CPU) GetCurrentOpcode() uint16 {
	return (uint16(cpu.Memory[cpu.PC]) << 8) | uint16(cpu.Memory[cpu.PC+1])
}
