package internal

import "fmt"

type Money int64

func (m Money) String() string {
	rubles := int64(m) / 100
	kopeks := int64(m) % 100
	return fmt.Sprintf("%d.%02d", rubles, kopeks)
}
