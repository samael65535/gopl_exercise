package lenconv

import (
	"fmt"
)

type Foot float64
type Meter float64

func (f Foot) String() string {
	return fmt.Sprintf("%gft", f)
}

func (m Meter) String() string {
	return fmt.Sprintf("%gm", m)
}
