package iter

import "github.com/BooleanCat/go-functional/option"

// RunesIter iterator, see [Runes].
type RunesIter struct {
	BaseIter[rune]
	runes []rune
	index int
}

// Next implements the [Iterator] interface.
func (iter *RunesIter) Next() option.Option[rune] {
	if iter.index >= len(iter.runes) {
		return option.None[rune]()
	}

	iter.index++

	return option.Some(iter.runes[iter.index-1])
}

var _ Iterator[rune] = new(RunesIter)

// Runes instantiates a [*RunesIter] that will yield a rune on iteration.
func Runes[T string | []rune](runes T) *RunesIter {
	iterator := &RunesIter{runes: []rune(runes)}
	iterator.BaseIter = BaseIter[rune]{iterator}
	return iterator
}
