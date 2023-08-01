package board

import (
	"testing"
)


var bb_set = []Bitboard{
			0b101, 
			0b1000000000000, 
			1<<64 -1, 
			1<<23 -1 + 1<<40,
			0b1111111111001111111111111,
			1<<60 + 1<<61 + 1<<62 + 1<<63,
			1<<40 + 1<<28 + 1<<16,
		}

// expected index sets
var bb_ind_set = [][]uint{ 
			[]uint{0, 2},  // 0
			[]uint{12},     // 1
			[]uint{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 
				33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63}, // 2
			[]uint{0,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,40}, // 3
			[]uint{0,1,2,3,4,5,6,7,8,9,10,11,12,15,16,17,18,19,20,21,22,23,24}, // 4
			[]uint{60,61,62,63}, // 5
			[]uint{16,28,40}, // 6
                         }

func Test_BB_to_index(b *testing.T) {

	for i, bb := range bb_set {
		ind := bb.Index()

		// test length
		if len(ind) != len(bb_ind_set[i]) {

			b.Errorf("BB_to_index failed Sets different length. \nSet: %d Result: %d Expected: %d, number: %b", i, ind, bb_ind_set[i], bb)
			
		} else { // test values

			for j, v := range ind {
				if v != uint(bb_ind_set[i][j]) {
					b.Errorf("BB_to_index value error. \nSet: %d Result: %d Expected: %d", i, ind, bb_ind_set[i])
					break
				}
			}
		}
	}
}

func Test_BB_flip(T *testing.T) {

	num := 0b1011_1000_0010_0001_0001
	original := Bitboard(num)
	bb := Bitboard(num)

	bb.flip(3)
	expected := original ^ (1<<3)
	if bb != expected {
		T.Errorf("BB_flip failed. \nResult: %b Expected: %b", bb, num)
	}
	
	bb.flip(12)
	expected = original ^ (1<<3) ^ (1<<12)
	if bb != expected {
		T.Errorf("BB_flip failed. \nResult: %b Expected: %b", bb, num)
	}
}




// benchmarking BB_to_index
// deprecated - here for reference


func BenchmarkBB_to_index(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, bb := range bb_set {

			bb.Index()
		}
	}
	b.ReportAllocs()
}

func BenchmarkBB_to_index2(b *testing.B) {

	for i := 0; i < b.N; i++ {
		for _, bb := range bb_set {

			bb.BB_to_index2()
		}
	}
	b.ReportAllocs()
}