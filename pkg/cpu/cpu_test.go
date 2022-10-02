package cpu_test

import (
	"github.com/NickDeChip/chip8-go/pkg/cpu"
	"testing"
)

func TestOp00E0(t *testing.T) {
	cpu := cpu.New()

	cpu.HandleOpcode(0x00E0)

	if cpu.ShouldClearScreen != true {
		t.Errorf("ShouldClearScreen was incorrect, got false, wanted true")
	}
}

func TestOp00EE(t *testing.T) {
	cpu := cpu.New()

	cpu.PC = 6
	cpu.Stack[cpu.SP] = cpu.PC
	cpu.SP += 1
	cpu.PC = 10

	cpu.HandleOpcode(0x00EE)

	if cpu.PC != 6 {
		t.Errorf("PC did not return to correct index, got %d wanted 6", cpu.PC)
	}
}

func TestOp1NNN(t *testing.T) {
	cpu := cpu.New()

	cpu.HandleOpcode(0x100A)

	if cpu.PC != 10 {
		t.Errorf("PC did not jump to correct index, got %d wanted 10", cpu.PC)
	}
}

func TestOp2NNN(t *testing.T) {
	cpu := cpu.New()

	cpu.HandleOpcode(0x200A)

	if cpu.PC != 10 {
		t.Errorf("PC did not jump to correct index, got %d wanted 10", cpu.PC)
	}

	if cpu.Stack[0] != 2 {
		t.Errorf("Stack did not have correct return address, got %d wanted 2", cpu.Stack[0])
	}

	if cpu.SP != 1 {
		t.Errorf("Stack Pointer did not increase, Got %d wanted 1", cpu.SP)
	}
}

func TestOp3XNN(t *testing.T) {
	cpu := cpu.New()

	x := 0

	cpu.V[x] = 4

	cpu.HandleOpcode(0x3004)

	if cpu.PC != 4 {
		t.Errorf("PC did not skip instructtion, got %d want 4", cpu.PC)
	}
}

func TestOp4XNN(t *testing.T) {
	cpu := cpu.New()

	x := 0

	cpu.V[x] = 2

	cpu.HandleOpcode(0x4004)

	if cpu.PC != 4 {
		t.Errorf("PC did not skip instructtion, got %d want 4", cpu.PC)
	}
}

func TestOp5XY0(t *testing.T) {
	cpu := cpu.New()

	x := 0
	y := 1

	cpu.V[x] = 4
	cpu.V[y] = 4

	cpu.HandleOpcode(0x5010)

	if cpu.PC != 4 {
		t.Errorf("PC did not skip instructtion, got %d want 4", cpu.PC)
	}
}

func TestOp6XNN(t *testing.T) {
	cpu := cpu.New()

	x := 0

	cpu.V[x] = 0

	cpu.HandleOpcode(0x6004)

	if cpu.V[x] != 4 {
		t.Errorf("VX does not eqaul NN, VX equals %d and NN equals 4", cpu.V[x])
	}
}

func TestOp7XNN(t *testing.T) {
	cpu := cpu.New()

	x := 0

	cpu.V[x] = 4

	cpu.HandleOpcode(0x7004)

	if cpu.V[x] != 8 {
		t.Errorf("VX did not add to NN, got %d want 8", cpu.V[x])
	}
}

func TestOp8XY0(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 99
	cpu.V[y] = 69

	cpu.HandleOpcode(0x8010)

	if cpu.V[x] != cpu.V[y] {
		t.Errorf("CPU V%d did not match v%d, got %d but wanted %d", x, y, cpu.V[x], cpu.V[y])
	}
}

func TestOp8XY1(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 0b0101_0101
	cpu.V[y] = 0b0010_0010

	cpu.HandleOpcode(0x8011)

	if cpu.V[x] != 0b0111_0111 {
		t.Errorf("CPU V%d is incorrect, got %b but wanted 0b0111_0111", x, cpu.V[x])
	}
}

func TestOp8XY2(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 0b0111_0111
	cpu.V[y] = 0b0010_0010

	cpu.HandleOpcode(0x8012)

	if cpu.V[x] != 0b0010_0010 {
		t.Errorf("CPU V%d is incorrect, got %b but wanted 0b0010_0010", x, cpu.V[x])
	}
}

func TestOp8XY3(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 0b0111_0111
	cpu.V[y] = 0b0010_0010

	cpu.HandleOpcode(0x8013)

	if cpu.V[x] != 0b0101_0101 {
		t.Errorf("CPU V%d is incorrect, got %b but wanted 0b0101_0101", x, cpu.V[x])
	}
}

func TestOp8XY4(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 150
	cpu.V[y] = 250

	cpu.HandleOpcode(0x8014)

	if cpu.V[x] != 144 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted %d.", x, cpu.V[x], cpu.V[x]-255)
	}
	if cpu.V[0xF] != 1 {
		t.Errorf("CPU VF was incorrect, got %d wanted 1", cpu.V[0xF])
	}
}

func TestOp8XY5(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 128
	cpu.V[y] = 114

	cpu.HandleOpcode(0x8015)

	if cpu.V[x] != 14 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted %d.", x, cpu.V[x], cpu.V[x]-cpu.V[y])
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("CPU VF was incorrect, got %d wanted 0", cpu.V[0xF])
	}
}

func TestOp8XY6(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 1

	cpu.HandleOpcode(0x8006)

	if cpu.V[x] != 0 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted 0", x, cpu.V[x])
	}
	if cpu.V[0xF] != 1 {
		t.Errorf("CPU VF is incorrect, got %d but wanted 1", cpu.V[0xF])
	}
}

func TestOp8XY7(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 114
	cpu.V[y] = 128

	cpu.HandleOpcode(0x8017)

	if cpu.V[x] != 14 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted %d.", y, cpu.V[x], 14)
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("CPU VF was incorrect, got %d wanted 0", cpu.V[0xF])
	}
}

func TestOp8XYE(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 128

	cpu.HandleOpcode(0x801E)

	if cpu.V[x] != 0 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted 0", x, cpu.V[x])
	}
	if cpu.V[0xF] != 128 {
		t.Errorf("CPU VF is incorrect, got %d but wanted 128", cpu.V[0xF])
	}
}

func TestOp9XY0(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 64
	cpu.V[y] = 128

	cpu.HandleOpcode(0x9010)

	if cpu.PC != 4 {
		t.Errorf("PC to be 4, But got %d", cpu.PC)
	}
}

func TestOpANNN(t *testing.T) {
	cpu := cpu.New()

	cpu.HandleOpcode(0xA010)

	if cpu.I != 16 {
		t.Errorf("cpu.I was not set to the NNN address, got %d wanted 18", cpu.PC)
	}
}

func TestOpBNNN(t *testing.T) {
	cpu := cpu.New()

	cpu.V[0] = 2

	cpu.HandleOpcode(0xB010)

	if cpu.PC != 18 {
		t.Errorf("PC to be 16, But got %d", cpu.PC)
	}
}

func TestOpCXNN(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.HandleOpcode(0xC0FF)

	if false {
		t.Errorf("opCXNN: %d", cpu.V[x])
	}

}

func TestOpEX9E(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 10
	cpu.Key[10] = 1

	cpu.HandleOpcode(0xE09E)

	if cpu.PC != 4 {
		t.Errorf("CPU PC was incorrect, got %d wanted 4", cpu.PC)
	}
}

func TestOpEXA1(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 10
	cpu.Key[10] = 0

	cpu.HandleOpcode(0xE0A1)

	if cpu.PC != 4 {
		t.Errorf("CPU PC was incorrect, got %d wanted 4", cpu.PC)
	}
}

func TestOpFX07(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	cpu.DelayTimer = 10

	cpu.V[x] = 0

	cpu.HandleOpcode(0xF007)

	if cpu.V[x] != 10 {
		t.Errorf("CPU V%d was incorrect, got %d wanted 10", x, cpu.V[x])
	}
}

func TestOpFX15(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	cpu.DelayTimer = 0

	cpu.V[x] = 10

	cpu.HandleOpcode(0xF015)

	if cpu.DelayTimer != 10 {
		t.Errorf("CPU DelayTimer was incorrect, got %d wanted 10", cpu.DelayTimer)
	}
}

func TestOpFX18(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	cpu.SoundTimer = 0

	cpu.V[x] = 10

	cpu.HandleOpcode(0xF018)

	if cpu.SoundTimer != 10 {
		t.Errorf("CPU DelayTimer was incorrect, got %d wanted 10", cpu.SoundTimer)
	}
}

func TestOpFX1E(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 8

	cpu.I += 2

	cpu.HandleOpcode(0xF01E)

	if cpu.I != 10 {
		t.Errorf("CPU I is incorrect, got %d but wanted 10", cpu.V[x])
	}
}

func TestOpFX55(t *testing.T) {
	cpu := cpu.New()

	cpu.V[0] = 2
	cpu.V[1] = 4
	cpu.V[2] = 6
	cpu.V[3] = 8
	cpu.V[4] = 10
	cpu.V[5] = 12
	cpu.V[6] = 14
	cpu.V[7] = 16
	cpu.V[8] = 18
	cpu.V[9] = 20
	cpu.V[0xA] = 22
	cpu.V[0xB] = 24
	cpu.V[0xC] = 26
	cpu.V[0xD] = 28
	cpu.V[0xE] = 30
	cpu.V[0xF] = 32

	cpu.I = 1000

	cpu.HandleOpcode(0xFB55)

	if cpu.Memory[cpu.I] != 2 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 2", cpu.Memory[cpu.I])
	}
	if cpu.Memory[cpu.I+1] != 4 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 4", cpu.Memory[cpu.I+1])
	}
	if cpu.Memory[cpu.I+2] != 6 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 6", cpu.Memory[cpu.I+2])
	}
	if cpu.Memory[cpu.I+3] != 8 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 8", cpu.Memory[cpu.I+3])
	}
	if cpu.Memory[cpu.I+4] != 10 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 10", cpu.Memory[cpu.I+4])
	}
	if cpu.Memory[cpu.I+5] != 12 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 12", cpu.Memory[cpu.I+5])
	}
	if cpu.Memory[cpu.I+6] != 14 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 14", cpu.Memory[cpu.I+6])
	}
	if cpu.Memory[cpu.I+7] != 16 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 16", cpu.Memory[cpu.I+7])
	}
	if cpu.Memory[cpu.I+8] != 18 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 18", cpu.Memory[cpu.I+8])
	}
	if cpu.Memory[cpu.I+9] != 20 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 20", cpu.Memory[cpu.I+9])
	}
	if cpu.Memory[cpu.I+0xA] != 22 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 22", cpu.Memory[cpu.I+0xA])
	}
	if cpu.Memory[cpu.I+0xB] != 24 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 24", cpu.Memory[cpu.I+0xB])
	}
	if cpu.Memory[cpu.I+0xC] != 0 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 0", cpu.Memory[cpu.I+0xC])
	}
	if cpu.Memory[cpu.I+0xD] != 0 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 0", cpu.Memory[cpu.I+0xD])
	}
	if cpu.Memory[cpu.I+0xE] != 0 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 0", cpu.Memory[cpu.I+0xE])
	}
	if cpu.Memory[cpu.I+0xF] != 0 {
		t.Errorf("CPU Memory is not set to register correctly, got %d wanted 0", cpu.Memory[cpu.I+0xF])
	}
}

func TestOpFX65(t *testing.T) {
	cpu := cpu.New()

	cpu.I = 1000

	cpu.Memory[cpu.I] = 2
	cpu.Memory[cpu.I+1] = 4
	cpu.Memory[cpu.I+2] = 6
	cpu.Memory[cpu.I+3] = 8
	cpu.Memory[cpu.I+4] = 10
	cpu.Memory[cpu.I+5] = 12
	cpu.Memory[cpu.I+6] = 14
	cpu.Memory[cpu.I+7] = 16
	cpu.Memory[cpu.I+8] = 18
	cpu.Memory[cpu.I+9] = 20
	cpu.Memory[cpu.I+0xA] = 22
	cpu.Memory[cpu.I+0xB] = 24
	cpu.Memory[cpu.I+0xC] = 26
	cpu.Memory[cpu.I+0xD] = 28
	cpu.Memory[cpu.I+0xE] = 30
	cpu.Memory[cpu.I+0xF] = 32

	cpu.HandleOpcode(0xFB65)

	if cpu.V[0] != 2 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 2", cpu.V[0])
	}
	if cpu.V[1] != 4 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 4", cpu.V[1])
	}
	if cpu.V[2] != 6 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 6", cpu.V[2])
	}
	if cpu.V[3] != 8 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 8", cpu.V[3])
	}
	if cpu.V[4] != 10 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 10", cpu.V[4])
	}
	if cpu.V[5] != 12 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 12", cpu.V[5])
	}
	if cpu.V[6] != 14 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 14", cpu.V[6])
	}
	if cpu.V[7] != 16 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 16", cpu.V[7])
	}
	if cpu.V[8] != 18 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 18", cpu.V[8])
	}
	if cpu.V[9] != 20 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 20", cpu.V[9])
	}
	if cpu.V[0xA] != 22 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 22", cpu.V[0xA])
	}
	if cpu.V[0xB] != 24 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 24", cpu.V[0xB])
	}
	if cpu.V[0xC] != 0 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 0", cpu.V[0xC])
	}
	if cpu.V[0xD] != 0 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 0", cpu.V[0xD])
	}
	if cpu.V[0xE] != 0 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 0", cpu.V[0xE])
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("CPU V is not set to register correctly, got %d wanted 0", cpu.V[0xF])
	}
}
