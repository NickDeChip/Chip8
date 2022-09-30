package cpu

import (
	"fmt"
	"math/rand"
)

type CPU struct {
	PC uint16

	//Registers
	I uint16
	V [16]uint8

	//STACK
	Stack [16]uint16
	SP    uint8
}

func New() *CPU {
	return &CPU{
		PC:    0,
		I:     0,
		V:     [16]uint8{},
		Stack: [16]uint16{},
		SP:    0,
	}
}

func (cpu *CPU) HandleOpcode(opcode uint16) {
	x, y := getXY(opcode)

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
		cpu.op4XNN(opcode, x)
	case 0x7000:
		cpu.op7XNN(opcode, x)
	case 0x8000:
		cpu.handle8Opcode(opcode, x, y)
	case 0x9000:
		cpu.op09XY0(x, y)
	case 0xA000:
		cpu.opANNN(opcode)
	case 0xB000:
		cpu.opBNNN(opcode)
	case 0xC000:
		cpu.opCXNN(opcode, x)
	case 0xD000:

	case 0xE000:
		cpu.handleEOpcodes(opcode, x)
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
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handle8Opcode(opcode uint16, x uint8, y uint8) {
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

func (cpu *CPU) handleEOpcodes(opcode uint16, x uint8) {
	switch opcode & 0x000F {
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
}

func (cpu *CPU) handleFopcodes(opcode uint16, x uint8) {
	switch opcode & 0x000F {
	case 0x000E:
		cpu.opFX1E(x)
	default:
		panic(fmt.Errorf("unkown opcode: %x", opcode))
	}
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
	cpu.V[x] |= cpu.V[y]
}

func (cpu *CPU) op8XY2(x uint8, y uint8) {
	cpu.V[x] &= cpu.V[y]
}

func (cpu *CPU) op8XY3(x uint8, y uint8) {
	cpu.V[x] ^= cpu.V[y]
}

func (cpu *CPU) op8XY4(x uint8, y uint8) {
	cpu.V[x] += cpu.V[y]
	overFlowCheck := uint16(cpu.V[x]) + uint16(cpu.V[y])
	if overFlowCheck > 255 {
		cpu.V[0xF] = 1
	}
}

func (cpu *CPU) op8XY5(x uint8, y uint8) {
	if cpu.V[y] > cpu.V[x] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}

	cpu.V[x] -= cpu.V[y]
}

func (cpu *CPU) op8XY6(x uint8) {
	cpu.V[0xF] = cpu.V[x] & 0x000F
	cpu.V[x] >>= 1
}

func (cpu *CPU) op8XY7(x uint8, y uint8) {
	if cpu.V[x] > cpu.V[y] {
		cpu.V[0xF] = 1
	} else {
		cpu.V[0xF] = 0
	}

	cpu.V[y] -= cpu.V[x]
}

func (cpu *CPU) op8XYE(x uint8) {
	cpu.V[0xF] = cpu.V[x] & 0x000F
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

	cpu.V[x] = uint8(rand.Intn(255)) & uint8(NN)
}

func (cpu *CPU) opFX1E(x uint8) {
	cpu.I += uint16(cpu.V[x])
}

// Get X and Y for cpu V`s
func getXY(opcode uint16) (x uint8, y uint8) {
	x = uint8(opcode & (0x0F00) >> 8)
	y = uint8(opcode & (0x00F0) >> 4)

	return x, y
}
