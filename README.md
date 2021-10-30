It's very important that you have to do this at the first:

go get github.com/go-sql-driver/mysql

if you are in china,you need to change the proxy:
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn,direc

Implementation
We’ll begin by connecting to a database we’ve set up on our local machine and then go on to perform some basic insert and select statements.

Connecting to a MySQL database
Let’s create a new main.go file. Within this, we’ll import a few packages and set up a simple connection to an already running local database. For the purpose of this tutorial, I’ve started MySQL using phpmyadmin and I’ve created a database called test to connect to and create tables within.

We’ll use sql.Open to connect to our database and set up our automatic connection pool, this will return either db or an err that we can handle.

package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

}
Performing Basic SQL Commands
So, now that we’ve created a connection, we need to start submitting queries to the database.

Thankfully, db.Query(sql) allows us to perform any SQL command we so desire. We can simply construct the query string and pass it in as a parameter.

package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    fmt.Println("Go MySQL Tutorial")

    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "root:password1@tcp(127.0.0.1:3306)/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()

    // perform a db.Query insert
    insert, err := db.Query("INSERT INTO test VALUES ( 2, 'TEST' )")

    // if there is an error inserting, handle it
    if err != nil {
        panic(err.Error())
    }
    // be careful deferring Queries if you are using transactions
    defer insert.Close()


}
Populating Structs from Results
Retrieving a set of results from the database is all well and good, but we need to be able to read these results or populating existing structs so that we can parse them and modify them easily. In order to parse a number of rows we can use the .Scan(args...) method which takes in any number of arguments and allows us to populate a composite object.

/*
 * Tag... - a very simple struct
 */
type Tag struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
func main() {
    // Open up our database connection.
    db, err := sql.Open("mysql", "root:pass1@tcp(127.0.0.1:3306)/tuts")

    // if there is an error opening the connection, handle it
    if err != nil {
        log.Print(err.Error())
    }
    defer db.Close()

    // Execute the query
    results, err := db.Query("SELECT id, name FROM tags")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }

    for results.Next() {
        var tag Tag
        // for each row, scan the result into our tag composite object
        err = results.Scan(&tag.ID, &tag.Name)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        log.Printf(tag.Name)
    }

}
In this example we retrieved 2 columns from the tags database and then used .Scan to populate our tag object.

Note - If you retrieve 3 fields from the database and Scan only has 2 parameters, it will fail. They need to match up exactly.

Querying a Single Row
Say we wanted to query a single row this time and had an ID and again wanted to populate our struct. We could do that like so:

var tag Tag
// Execute the query
err = db.QueryRow("SELECT id, name FROM tags where id = ?", 2).Scan(&tag.ID, &tag.Name)
if err != nil {
    panic(err.Error()) // proper error handling instead of panic in your app
}

log.Println(tag.ID)
log.Println(tag.Name)