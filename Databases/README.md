# Working with Databases
- SQL Database API in its standard library `database/sql` package.
- Different SQL Services provides different drivers to work on and install them by
    - MySQL - `go get "github.com/go-sql-driver/mysql"`
    - PostGres - `go get "github.com/lib/pq"`

All drivers are listed here - [Drivers](https://github.com/golang/go/wiki/SQLDrivers)

- `sql.Open()` function specifies the type of database server we are connecting to. The second argument is a `Data Source Name (DSN)` indicating the following :-

    >>>
    User Account

    Password

    IP Address of the database server

    Port Number of the database server

    Database name
    >>>

- `Data Source Name (DSN)` format :
> <user>:<password>@tcp(<ip_address>:<port>)/<DB_name>

## Retrieving a record
```
>> Create a type called Course with two fields: ID and Details
>> Define a function named Records to fetch records from the database and then print them out.
```
- `Query()` function to perform a query on the table. 
- `Next()` method to move from one row to another.
- `Scan()` method from the Rows struct to read the records from the table and map it to the fields in the Course struct.

## Adding, Modifying and Deleting a record
- `Exec()` function, which executes a query without returning any rows. The second argument of the `Exec()` function contains the values to pass into the query’s placeholders. This is known as a `parameterized SQL statement`. It returns a `Result` struct. Using the result, you can
determine if the record is inserted successfully using the `RowsAffected()` method.
- If you need to know the ID of the row inserted, you can use the `LastInsertId()` method of the `Result` struct