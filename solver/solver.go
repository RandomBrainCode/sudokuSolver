package solver

type Block struct {
	row      int
	col      int
	box      int
	value    int
	possible []int
}

func (b *Block) SetBox() {
	if b.box < 1 || b.box > 9 {
		b.box = GetBoxNumber(b.row, b.col)
	}
}

func GetBoxNumber(row, col int) int {
	return (((row-1)/3)*3 + (col-1)/3) + 1
}

func (b *Block) SetPossible(Grid []Block) {
	b.possible = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, block := range Grid {
		if b.row == block.row || b.col == block.col || b.box == block.box {
			if block.value >= 1 || block.value <= 9 {
				b.RemovePossible(block.value)
			}
		}
	}
}

func (b *Block) RemovePossible(value int) {
	for i, possible := range b.possible {
		if possible == value {
			b.possible = append(b.possible[:i], b.possible[i+1:]...)
			break
		}
	}
}

func Solver(Grid []Block) []Block {

}
