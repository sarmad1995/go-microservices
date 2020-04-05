package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	els := []int{9, 8, 6, 7, 5, 3, 2, 1}

	BubbleSort(els)
	assert.NotNil(t, els)
	assert.EqualValues(t, 8, len(els))
	assert.EqualValues(t, 1, els[0])
	assert.EqualValues(t, 2, els[1])
	assert.EqualValues(t, 3, els[2])
	assert.EqualValues(t, 5, els[3])
	assert.EqualValues(t, 6, els[4])
	assert.EqualValues(t, 7, els[5])
	assert.EqualValues(t, 8, els[6])
	assert.EqualValues(t, 9, els[7])

}

func getElements(n int) []int {
	result := make([]int, n)
	i := 0
	for j := n - 1; j >= 0; j-- {
		result[i] = j
		i++
	}
	return result
}

func BenchmarkBubbleSort(b *testing.B) {
	els := getElements(2)
	for i := 0; i < b.N; i++ {
		BubbleSort(els)
	}
}
