package util

const (
	KNOT_SIZE  = 256
	NUM_ROUNDS = 64
	BLOCK_SIZE = 16
)

type HashKnot struct {
	index    int
	skipSize int
	Ring     []int
}

func NewHashKnot() *HashKnot {
	numbers := make([]int, KNOT_SIZE)
	i := 0
	for i < len(numbers) {
		numbers[i] = i
		i++
	}
	return &HashKnot{
		index:    0,
		skipSize: 0,
		Ring:     numbers,
	}
}

func (hk *HashKnot) Hash() []int {
	numBlocks := len(hk.Ring) / BLOCK_SIZE
	blocks := make([]int, 0, numBlocks)
	i := 0
	for i < numBlocks {
		start := i * BLOCK_SIZE
		end := start + BLOCK_SIZE
		block := hk.Ring[start:end]

		v := block[0]
		j := 1
		for j < BLOCK_SIZE {
			v ^= block[j]
			j++
		}
		blocks = append(blocks, v)
		i++
	}
	return blocks
}

func (hk *HashKnot) RunRound(lengths []int) {
	for _, length := range lengths {
		hk.PinchTwist(length)
	}
}

func (hk *HashKnot) PinchTwist(length int) {
	if length > len(hk.Ring) {
		return
	}

	buf := make([]int, length)
	for j, _ := range buf {
		i := ((hk.index + length) - j - 1) % len(hk.Ring)
		buf[j] = hk.Ring[i]
	}

	for i, v := range buf {
		j := (hk.index + i) % len(hk.Ring)
		hk.Ring[j] = v
	}

	hk.index = (hk.index + hk.skipSize + length) % len(hk.Ring)
	hk.skipSize++
}

func AsASCIICodes(str string) []int {
	buf := make([]int, 0, len(str))
	for _, v := range str {
		buf = append(buf, int(v))
	}
	return buf
}
