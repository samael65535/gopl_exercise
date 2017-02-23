package main
/*
练习 7.16：
编写一个基于web的计算器程序。
*/
import (
	_"io"
	_"bytes"
	"fmt"
	"../ex7.14/eval"
	"net/http"
	"strings"
)
func cal(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	t := r.Form.Get("expr")
	s := strings.Replace(t, " ", "+", -1)
	fmt.Println(t)
	result, err := eval.Parse(s)
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "%s=%f", s,result.Eval(eval.Env{}))
	return
}

func main() {
	http.HandleFunc("/", cal)
	http.ListenAndServe("localhost:8000", nil)
}
