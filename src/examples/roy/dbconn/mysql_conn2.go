package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
)

func main() {
    db, err := sql.Open("mysql", "testuser:testuser@/test_db?charset=utf8")
    checkErr(err)

    // insert
    stmt, err := db.Prepare("INSERT userinfo_tbl SET username=?,departname=?,created=?")
    checkErr(err)

    t := time.Now()
    res, err := stmt.Exec("ericsson", "Sales", t)
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

    fmt.Println(id)
    // update
    stmt, err = db.Prepare("update userinfo_tbl set username=? where uid=?")
    checkErr(err)

    res, err = stmt.Exec("anthony", id)
    checkErr(err)

    affect, err := res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    // query
    rows, err := db.Query("SELECT * FROM userinfo_tbl")
    checkErr(err)

    for rows.Next() {
        var uid int
        var username string
        var department string
        var created string
        err = rows.Scan(&uid, &username, &department, &created)
        checkErr(err)
        fmt.Println(uid)
        fmt.Println(username)
        fmt.Println(department)
        fmt.Println(created)
    }

    // delete
    stmt, err = db.Prepare("delete from userinfo_tbl where uid=?")
    checkErr(err)

    res, err = stmt.Exec(id)
    checkErr(err)

    affect, err = res.RowsAffected()
    checkErr(err)

    fmt.Println(affect)

    db.Close()

}

func checkErr(err error) {
    if err != nil {
        //panic(err.Error())
        log.Fatal("Error: ", err)
    }
}
