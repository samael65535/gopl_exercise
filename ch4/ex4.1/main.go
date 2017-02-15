package main

//  编写一个函数，计算两个SHA256哈希码中不同bit的数目。
import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		// 计算有每8位有几个个1
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func count1(x byte) int {
	// 数字中有几个1
	return int(pc[x])

}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	fmt.Println(diffCount(&c1, &c2))
}

func diffCount(ptr1 *[32]byte, ptr2 *[32]byte) int {
	count := 0
	for i := range ptr1 {
		n := ptr2[i] ^ ptr1[i]
		count += count1(n)
	}
	return count
}
