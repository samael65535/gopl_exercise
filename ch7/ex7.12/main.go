package main

/*
练习 7.11：
改/list的handler让它把输出打印成一个HTML的表格而不是文本。
*/

import (
	"net/http"
	"fmt"
	"strconv"
	"html/template"
	"log"
)
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars


const templ = `
<table>
<tr style='text-align: left' >
<td>Item</td>
<td>Price</td>
</tr>
{{range $item, $price := .}}
<tr>
  <td style="padding-right: 30px;">{{$item}}</td>
  <td style="padding-right: 30px;">{{$price}}</td>
</tr>
{{end}}
</table>

`
func (db database)Update(name string, price dollars) {
	if _, ok := db[name]; ok {
		db[name] = price
	}
}


func (db database)Add(name string, price dollars) {
	db[name] = price
}

func (db database)Remove(name string) {
	delete(db, name)
}

func main() {
	db := database{"test": 12}
	server := http.NewServeMux()
	server.HandleFunc("/update", db.updateHandler)
	server.HandleFunc("/add",db.addHandler)
	server.HandleFunc("/list", db.listHandler)
	server.HandleFunc("/", db.listHandler)
	server.HandleFunc("/remove", db.removeHandler)
	http.ListenAndServe("localhost:8001", server)
}

func (db database)updateHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}

	db.Update(item, dollars(price))
	db.listHandler(w, r)
}

func (db database)addHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	if err != nil {
		fmt.Fprintf(w, "%s", err.Error())
		return
	}
	db.Add(item, dollars(price))

	db.listHandler(w, r)
}

func (db database)removeHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	db.Remove(item)

	db.listHandler(w, r)
}

func (db database)listHandler(w http.ResponseWriter, r *http.Request) {
	result := template.Must(template.New("itemlist").Parse(templ))
	if err := result.Execute(w, db); err != nil {
		log.Fatal(err)
	}
}
