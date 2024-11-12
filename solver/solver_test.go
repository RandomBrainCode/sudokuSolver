package solver

import "testing"

type testBlock struct {
	row int
	col int
	box int
}

func TestBlock_SetBox(t *testing.T) {
	var blocks []testBlock
	blocks = append(blocks, testBlock{1, 1, 1})
	blocks = append(blocks, testBlock{3, 3, 1})
	blocks = append(blocks, testBlock{1, 4, 2})
	blocks = append(blocks, testBlock{3, 6, 2})
	blocks = append(blocks, testBlock{1, 7, 3})
	blocks = append(blocks, testBlock{1, 9, 3})
	blocks = append(blocks, testBlock{3, 9, 3})
	blocks = append(blocks, testBlock{5, 2, 4})
	blocks = append(blocks, testBlock{5, 5, 5})
	blocks = append(blocks, testBlock{5, 8, 6})
	blocks = append(blocks, testBlock{7, 3, 7})
	blocks = append(blocks, testBlock{7, 6, 8})
	blocks = append(blocks, testBlock{7, 9, 9})
	blocks = append(blocks, testBlock{9, 1, 7})
	blocks = append(blocks, testBlock{9, 4, 8})
	blocks = append(blocks, testBlock{9, 7, 9})
	blocks = append(blocks, testBlock{9, 9, 9})

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

func GridBuilder() []*Block {
	var Grid []*Block
	for r := 1; r < 10; r++ {
		for c := 1; c < 10; c++ {
			Grid = append(Grid, &Block{r, c, 0, 0, []int{}})
		}
	}
	return Grid
}

func TestBlock_Possible(t *testing.T) {
	Grid := GridBuilder()

}
