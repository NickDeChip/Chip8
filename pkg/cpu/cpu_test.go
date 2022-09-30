package cpu_test

import (
	"github.com/NickDeChip/chip8-go/pkg/cpu"
	"testing"
)

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

	cpu.V[x] = 4

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

	cpu.V[x] = 8

	cpu.HandleOpcode(0x8006)

	if cpu.V[x] != 4 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted 4", x, cpu.V[x])
	}
	if cpu.V[0xF] != 8 {
		t.Errorf("CPU VF is incorrect, got %d but wanted 8", cpu.V[0xF])
	}
}

func TestOp8XY7(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)
	y := uint8(1)

	cpu.V[x] = 114
	cpu.V[y] = 128

	cpu.HandleOpcode(0x8017)

	if cpu.V[y] != 14 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted %d.", y, cpu.V[y], 14)
	}
	if cpu.V[0xF] != 0 {
		t.Errorf("CPU VF was incorrect, got %d wanted 0", cpu.V[0xF])
	}
}

func TestOp8XYE(t *testing.T) {
	cpu := cpu.New()

	x := uint8(0)

	cpu.V[x] = 8

	cpu.HandleOpcode(0x801E)

	if cpu.V[x] != 16 {
		t.Errorf("CPU V%d is incorrect, got %d but wanted 16", x, cpu.V[x])
	}
	if cpu.V[0xF] != 8 {
		t.Errorf("CPU VF is incorrect, got %d but wanted 8", cpu.V[0xF])
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
