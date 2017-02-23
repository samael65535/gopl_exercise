package eval

import (
//	"fmt"
//	"strings"
	"strconv"
)

// ex7.13
func (v Var)  String() string{
	return string(v)
}

func (l literal) String() string {
	return strconv.FormatFloat(float64(l), 'g', -1, 64)
}

func (u unary) String() string{
	op := strconv.QuoteRune(u.op)
	if (op == "-") {
		return "-" + u.x.String()
	}
	return u.x.String()
}

func (b binary) String() string {
	return b.x.String() + string(b.op) + b.y.String()
}

func (c call) String() string {
	str := c.fn + "("
	for _, a := range c.args {
		str += a.String()
	}
	str += ")"
	return str
}
