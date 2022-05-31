package sql3

import (
   //"fmt"
   "database/sql"
   "log"
   _ "github.com/mattn/go-sqlite3"
)

func OpenSqlite3(fname string) *sql.DB {
   db, err := sql.Open("sqlite3", fname)
   if err != nil {
      log.Printf("Cannot open file '%s'\n", fname)
      log.Fatal(err)
   }
   return db
}

func ExecSqlOnly(db *sql.DB, sql string) {
   _, err := db.Exec(sql)
   if err != nil {
      log.Printf("%q: %s\n", err, sql)
      log.Fatal(err)
   }
}

func ExecSql(db *sql.DB, sql string) *sql.Rows {
   rows, err := db.Query(sql)
   if err != nil {
      log.Printf("sql: '%s'\n", sql)
      log.Fatal(err)
   }
   err = rows.Err()
   if err != nil {
      log.Fatal(err)
   }
   return rows
}

func Scan(rows *sql.Rows, dest... interface{}) {
   err := rows.Scan(dest...)
   if err != nil {
      log.Fatal(err)
   }
}

/* Not allowed
func (rows *sql.Rows) ScanRows(dest... interface{}) {
   err := rows.Scan(dest...)
   if err != nil {
      log.Fatal(err)
   }
}
*/

// Prepare SQL template
func PrepTempl(db *sql.DB, sql string) *sql.Stmt {
   stmt, err := db.Prepare(sql)
   if err != nil {
      log.Printf("sql='%s'\n", sql)
      log.Fatal(err)
   }
   return stmt
}

func TemplExec(stmt *sql.Stmt, args... interface{}) {
    _, err :=  stmt.Exec(args...)
    if err != nil {
       log.Fatal(err)
    }
}

func TemplGetRows(stmt *sql.Stmt, args... interface{}) *sql.Rows {
   rows, err := stmt.Query(args...)
   if err != nil {
      log.Fatal(err)
   }
   return rows
}

