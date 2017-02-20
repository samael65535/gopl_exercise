package main
// 编写多参数版本的strings.Join。
import (
	"fmt"
)

func join(str ...string) (string, error){
	if (len(str) <= 1) {
		return "", fmt.Errorf("error")
	}
	var sep = str[len(str)-1]
	result := str[0]
	for _,v := range str[1:len(str) - 1] {
		result += sep + v
	}
	return result, nil

}

func main() {
	s, _ := join("dfdf", "adf", "dfdfd", ",")
	fmt.Println(s)
}
