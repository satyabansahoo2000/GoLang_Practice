package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type Course struct {
	ID string
	Details string
}

func GetRecords(db *sql.DB) {
	results, err := db.Query("SELECT * FROM Course")
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var course Course
		err = results.Scan(&course.ID, "-", &course.Details)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(course.ID, "-", course.Details)
	}
}

func InsertRecord(db *sql.DB, ID string, Details string) {
    // use parameterized SQL statement
    result, err := db.Exec(
      "INSERT INTO Course VALUES (?, ?)", ID, Details)
    if err != nil {
        fmt.Println(err.Error())
    } else {
        if count, err := result.RowsAffected(); err == nil {
            fmt.Println(count, "row(s) affected")
        }
    }
}

func EditRecord(db *sql.DB, ID string, Details string) {
    result, err := db.Exec(
        "UPDATE Course SET Details=? WHERE ID=?", Details, ID,
    )
    if err != nil {
        fmt.Println(err.Error())
    } else {
        if count, err := result.RowsAffected(); err == nil {
            fmt.Println(count, "row(s) affected!!")
        }
    }
}

func DeleteRecord(db *sql.DB, ID string) {
    result, err := db.Exec(
        "DELETE FROM Course WHERE ID=?", ID,
    )
    if err != nil {
        fmt.Println(err.Error())
    } else {
        if count, err := result.RowsAffected(); err == nil {
            fmt.Println(count, "row(s) affected!!!")
        }
    }
}

func main() {
    db, err := sql.Open("postgres","gouser:password@tcp(127.0.0.1:5432)/CoursesDB")
        
    // handle error
    if err != nil {
        panic(err.Error())
    } else {
        fmt.Println("Database object created")
		InsertRecord(db, "IOS101", "iOS Programming")
		EditRecord(db, "IOS101", "SwiftUI Programming")
        DeleteRecord(db, "IOS101")
        GetRecords(db)
    }

    // defer the close till after the main function has
    // finished executing
    defer db.Close()
}