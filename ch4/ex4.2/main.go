package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main() {

	usage := `
        Usage: 
        ./main -f sha256 require_input
        ./main -f sha512 require_input
        ./main -f sha384 require_input
        ./main require_input
    `
	fmt.Println(usage)
	f := os.Args[1]
	switch f {
	case "-h":
		fmt.Fprint(os.Stdout, usage)
		break
	case "-f":
		s := []byte(os.Args[3])
		outputSHA(os.Args[2], &s)
		break
	default:
		s := []byte(os.Args[1])
		outputSHA("sha256", &s)
	}
}
func outputSHA(flag string, input *[]byte) int {
	switch flag {
	case "sha512":
		fmt.Fprintf(os.Stdout, "sha512 %x", sha512.Sum512(*input))
		break
	case "sha384":
		fmt.Fprintf(os.Stdout, "sha384 %x", sha512.Sum384(*input))
		break
	default:
		fmt.Fprintf(os.Stdout, "sha256 %x", sha256.Sum256(*input))
	}
	return 0
}
