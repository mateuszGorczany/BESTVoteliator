package common

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStream(t *testing.T) {
	list := []int{1, 2, 3, 4}
	slist := Stream[int]{list}
	slist2 := slist.
		Map(func(e int) int {
			return 2 * e
		}).
		Filter(func(e int) bool {
			return e%2 == 0
		}).
		Collect()

	fmt.Printf("%v", slist)
	assert.Equal(t, slist2, []int{2, 4, 6, 8})

}
