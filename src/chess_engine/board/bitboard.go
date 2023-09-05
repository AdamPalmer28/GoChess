package board

type Bitboard uint64

const (
	FileA Bitboard = 0x0101010101010101
	FileB Bitboard = FileA << 1
	FileC Bitboard = FileA << 2
	FileD Bitboard = FileA << 3
	FileE Bitboard = FileA << 4
	FileF Bitboard = FileA << 5
	FileG Bitboard = FileA << 6
	FileH Bitboard = FileA << 7

	Rank1 Bitboard = 0xFF
	Rank2 Bitboard = Rank1 << 8
	Rank3 Bitboard = Rank1 << 16
	Rank4 Bitboard = Rank1 << 24
	Rank5 Bitboard = Rank1 << 32
	Rank6 Bitboard = Rank1 << 40
	Rank7 Bitboard = Rank1 << 48
	Rank8 Bitboard = Rank1 << 56

	EmptyBoard Bitboard = 0
)

func (b Bitboard) Print() {

	for rank := uint(7); rank < 8; rank-- {
		for file := uint(0); file < 8; file++ {
			square := (rank * 8) + file
			mask := Bitboard(1) << square
			if (b & mask) != 0 {
				print("X ")
			} else {
				print("- ")
			}
		}
		println()
	}
	println()
}

func (b Bitboard) Index() []uint {
	// convert a bitboard to slide index

	ind := []uint{}
	var bitPosition uint = 0

	for b > 0 {
		if b&1 == 1 {
			ind = append(ind, bitPosition)
		}
		bitPosition++
		b >>= 1
	}

	return ind
}

// slower method
func (b Bitboard) BB_to_index2() []uint {
	// convert a bitboard to slide index

	var ind []uint

	for i := uint(0); i < 64; i++ {
		mask := Bitboard(1) << i
		if (b & mask) != 0 {
			ind = append(ind, uint(i))
		}
	}
	return ind
}

// flip bb index
func (b *Bitboard) flip(ind uint) {

	*b ^= 1 << ind
}

func Make_bitboard(ind []uint) *Bitboard {
	var BB Bitboard
	for _, i := range ind {
		BB.flip(i)
	}
	return &BB
}