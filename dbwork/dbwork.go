package dbwork

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"github.com/bobilev/sasds/structs"
	"fmt"
)

func dbConnect() *sql.DB{
	dbname := "test"
	dbuser := "bobilev"
	dbpassword := "130215"
	db, err := sql.Open("mysql", ""+dbuser+":"+dbpassword+"@tcp(localhost:3306)/"+dbname+"?charset=utf8")
	checkErr(err)
	return db
}
func DboperInsert(name string) {//вставить - Insert
	db := dbConnect()
	defer db.Close()

	// insert
	stmt, err := db.Prepare("INSERT series SET name=?")
	checkErr(err)

	res, err := stmt.Exec(name)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)


}
func DboperUpdate(id int) {
	db := dbConnect()
	defer db.Close()
	// update
	stmt, err := db.Prepare("update series set text=? where id=?")
	checkErr(err)

	res, err := stmt.Exec("update", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}
func DboperQuery() ([]structs.Record, error) {//Выборка - Query
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT id,name FROM series")
	checkErr(err)
	var rs = make([]structs.Record, 0)
	var rec structs.Record
	for rows.Next() {
		var id int
		var name string

		err = rows.Scan(&rec.Id, &rec.Name)
		checkErr(err)
		fmt.Println("------------")
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println("------------")
		rs = append(rs, rec)
	}
	return rs, err
}
func DboperQueryLast() map[string]map[string]int {
	mapLast := make(map[string]map[string]int)
	db := dbConnect()
	defer db.Close()

	rows, err := db.Query("SELECT name,season,episode FROM series")
	checkErr(err)

	for rows.Next() {
		var name string
		var season int
		var episode int

		err = rows.Scan(&name, &season, &episode)
		checkErr(err)
		//fmt.Println(name)
		//fmt.Println(season)
		//fmt.Println(episode)

		mapLast[name] = make(map[string]int)

		mapLast[name]["season"] = season
		mapLast[name]["episode"] = episode
		//fmt.Println(mapLast[name]["season"])
		//fmt.Println(mapLast[name]["episode"])
	}
	return mapLast
}
func DboperDelet(id int) {
	db := dbConnect()
	defer db.Close()

	stmt, err := db.Prepare("delete from series where id=?")
	checkErr(err)

	res, err := stmt.Exec(id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
