package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    db, err := sql.Open("mysql", "testuser:testuser@/school_db?charset=utf8")
    if err != nil {
        panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()

    // Prepare statement for inserting data
    stmtIns, err := db.Prepare("INSERT INTO class_tbl (id, name, gpa) VALUES( ?, ?, ? )") // ? = placeholder
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtIns.Close() // Close the statement when we leave main() / the program terminates

    // Prepare statement for reading data
    stmtOut, err := db.Prepare("SELECT name FROM class_tbl WHERE id = ?")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()

    // Insert class record into class_tbl
    for i := 5; i < 10; i++ {
        _, err = stmtIns.Exec((i * 11), "New Student "+string((i*11)), 1.0) // Insert tuples (i, i^2)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
    }

    var student_name string // we "scan" the result in here

    // Query name of student where id is 11
    err = stmtOut.QueryRow(55).Scan(&student_name) // WHERE number = 11
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The name of student whose id is 55: %s\n", student_name)

    // Query name of another student where id is 44
    err = stmtOut.QueryRow(99).Scan(&student_name) // WHERE number = 44
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("The name of student whose id is 99: %s\n", student_name)
}
