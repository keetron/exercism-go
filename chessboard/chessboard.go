package chessboard

// File Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File []bool

// Chessboard Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	a, e := cb[file]
	if !e {
		return 0
	}
	c := 0
	for _, b := range a {
		if b {
			c++
		}
	}
	return c
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}

	c := 0
	for _, v := range cb {
		if v[rank-1] {
			c++
		}
	}
	return c
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	return len(cb) * 8
	//s := 0
	//for range cb {
	//	s += 1
	//}
	//s = s * 8
	//return s
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	c := 0
	for _, v := range cb {
		for _, b := range v {
			if b {
				c++
			}
		}
	}
	return c
}
