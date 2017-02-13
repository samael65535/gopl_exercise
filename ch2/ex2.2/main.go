package main

import "fmt"
import "./lenconv"
import "os"
import "strconv"

func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}

		ft := lenconv.Foot(t)
		m := lenconv.Meter(t)

		fmt.Printf("%s = %s, %s = %s\n", ft, lenconv.FTToM(ft).String(), m, lenconv.MToFT(m).String())
	}
}
