package solver

type Block struct {
	row      int
	col      int
	box      int
	value    int
	possible []int
}

type Grid []*Block

func (b *Block) SetBox() {
	if b.box < 1 || b.box > 9 {
		b.box = GetBoxNumber(b.row, b.col)
	}
}

func GetBoxNumber(row, col int) int {
	return (((row-1)/3)*3 + (col-1)/3) + 1
}

func (b *Block) SetPossible(grid *Grid) {
	b.possible = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for _, block := range *grid {
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

func (g *Grid) Build() {
	for row := 1; row < 10; row++ {
		for col := 1; col < 10; col++ {
			b := &Block{row, col, 0, 0, []int{}}
			b.SetBox()
			*g = append(*g, b)
		}
	}
}

func (g *Grid) ModBlock(row, col, value int) {
	b := (*g)[(row-1)*9+(col-1)]
	b.value = value
	b.possible = []int{}
}

func (g *Grid) Solve() {}
