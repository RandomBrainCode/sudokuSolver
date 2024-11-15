package solver

import (
	"testing"
)

type boxBlockTest struct {
	row int
	col int
	box int
}

func TestBlock_SetBox(t *testing.T) {
	var blocks []boxBlockTest
	blocks = append(blocks, boxBlockTest{1, 1, 1})
	blocks = append(blocks, boxBlockTest{3, 3, 1})
	blocks = append(blocks, boxBlockTest{1, 4, 2})
	blocks = append(blocks, boxBlockTest{3, 6, 2})
	blocks = append(blocks, boxBlockTest{1, 7, 3})
	blocks = append(blocks, boxBlockTest{1, 9, 3})
	blocks = append(blocks, boxBlockTest{3, 9, 3})
	blocks = append(blocks, boxBlockTest{5, 2, 4})
	blocks = append(blocks, boxBlockTest{5, 5, 5})
	blocks = append(blocks, boxBlockTest{5, 8, 6})
	blocks = append(blocks, boxBlockTest{7, 3, 7})
	blocks = append(blocks, boxBlockTest{7, 6, 8})
	blocks = append(blocks, boxBlockTest{7, 9, 9})
	blocks = append(blocks, boxBlockTest{9, 1, 7})
	blocks = append(blocks, boxBlockTest{9, 4, 8})
	blocks = append(blocks, boxBlockTest{9, 7, 9})
	blocks = append(blocks, boxBlockTest{9, 9, 9})

	for _, testValue := range blocks {
		b := BlockSetbox(testValue.row, testValue.col)
		if b.box != testValue.box {
			t.Fatalf("Given Row %d, Column %d. Expected Box Address of %d, got %d", testValue.row, testValue.col, testValue.box, b.box)
		} else {
			t.Logf("Given Row %d, Column %d. Retrieved Box Address %d.", testValue.row, testValue.col, b.box)
		}
	}
}

func BlockSetbox(row int, col int) Block {
	b := Block{}
	b.row = row
	b.col = col
	b.SetBox()
	return b
}

func TestGrid_Build(t *testing.T) {
	grid := Grid{}
	grid.Build()

	if len(grid) != 81 {
		t.Fatalf("Expected grid length of 81, got %d", len(grid))
	} else {
		t.Log("Grid Block Count is 81 blocks.")
	}

	for pos := 0; pos < 81; pos++ {
		block := grid[pos]
		blockPos := BlockPos(block.row, block.col)
		if pos != blockPos {
			t.Fatalf("Block with Row %d, Column %d expected Position %d but got %d", block.row, block.col, blockPos, pos)
		}
	}
	t.Log("All 81 positions mapped successfully.")
}

func BlockPos(row, col int) int {
	return ((row - 1) * 9) - 1 + col
}
