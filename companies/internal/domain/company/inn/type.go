package inn

import (
	"fmt"
	"strconv"
)

const (
	DigitsCnt = 10
)

var (
	ErrWrongLength = fmt.Errorf("number of digits in inn must be %d", DigitsCnt)
)

type Inn uint32

func (n Inn) String() string {
	return strconv.FormatUint(uint64(n), 10)
}

func (n Inn) Inn() uint32 {
	return uint32(n)
}

func NewInn(inn uint32) (*Inn, error) {
	innString := strconv.FormatUint(uint64(inn), 10)
	if len([]rune(innString)) < DigitsCnt || len([]rune(innString)) > DigitsCnt {
		return nil, ErrWrongLength
	}

	n := Inn(inn)

	return &n, nil
}
