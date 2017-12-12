//dbActions.go

package main

import (
    "time"
    "database/sql"
    _ "github.com/lib/pq"
    "log"
    "os"
)

/**
 * does db actions. Called by handlers when requests are made
 **/

 var db *sql.DB

/**
 * initializes sql db
 **/
func openDB() bool {
    var err error
    db, err = sql.Open("postgres", os.Getenv("DATABASE_URL"))
    if err != nil {
        log.Fatalf("Error opening database: %q", err)
        return false
    }
    return true
}

/**
 * reads all rows from DB
 * @return {String} encoded string
 **/
func readAllRows() string {
    if _, err := db.Exec("CREATE TABLE IF NOT EXISTS ticks (tick timestamp)"); err != nil {
        return "Error creating database table"
    }

    if _, err := db.Exec("INSERT INTO ticks VALUES (now())"); err != nil {
        return "INSERT INTO ticks VALUES (now())"
    }

    rows, err := db.Query("SELECT tick FROM ticks")
    if err != nil {
        return "Error reading ticks"
    }

    defer rows.Close()
    for rows.Next() {
        var tick time.Time
        if err := rows.Scan(&tick); err != nil {
            return "Error scanning ticks"
        }
        return tick.String()
    }
    return ""
}
