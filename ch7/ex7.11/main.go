package main

/*
练习 7.11：
增加额外的handler让客服端可以创建，读取，更新和删除数据库记录。
例如，一个形如 /update?item=socks&price=6 的请求会更新库存清单里一个货品的价格并且当这个货品不存在或价格无效时返回一个错误值。（注意：这个修改会引入变量同时更新的问题）
*/

import (
	"net/http"
	"fmt"
	"strconv"
)
type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

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
	for item, price := range db {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}
