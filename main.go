package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/bobilev/sasds/dbwork"
	"net/http"
	//_ "github.com/go-sql-driver/mysql"
	"encoding/json"
	"strconv"
)


func getRecords(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var str string
	//Тут вылавливаем параметр name в URL если он был передан
	if len(r.URL.RawQuery) > 0 {
		str = r.URL.Query().Get("name")
		if str == "" {
			w.WriteHeader(400)
			return
		}
	}
	//Запрос к БД возвращает структуру structs.Record[id: int, name: string]::: Можно все что взбредет в голову
	recs ,err := dbwork.DboperQuery()

	if err != nil {
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	//Если нету ошибки то отправляем назад
	if err = json.NewEncoder(w).Encode(recs); err != nil {
		w.WriteHeader(500)
	}
}
func getID(w http.ResponseWriter, ps httprouter.Params) (int, bool) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.WriteHeader(400)
		return 0, false
	}
	return id, true
}


func main() {
	router := httprouter.New()
	router.GET("/api/records", getRecords)
	//router.GET("/api/records/:id", getRecord)
	//router.POST("/api/records", addRecord)
	//router.PUT("/api/records/:id", updateRecord)
	//router.DELETE("/api/records/:id", deleteRecord)
	fmt.Println("Server start, port:80")
	http.ListenAndServe(":80", router)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}