package main

func main() {
	DataMigration()
	HandlerRouting()

}

// package main

// import (
// 	"database/sql"
// 	"fmt"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	fmt.Print("Go Connect  to  mysql")
// 	db, err := sql.Open("mysql", "root:demo1234@tcp(127.0.0.1:3306)/testdb")
// 	if err != nil {
// 		println((err.Error()))
// 		fmt.Print("not Connected")
// 	}
// 	fmt.Print("Connected")
// 	defer db.Close()
// }

// package main

// import (
//     "database/sql"
//     "fmt"
//     "log"

//     _ "github.com/go-sql-driver/mysql"
// )

// func main() {

//     db, err := sql.Open("mysql", "user7:s$cret@tcp(127.0.0.1:3306)/testdb")
//     defer db.Close()

//     if err != nil {
//         log.Fatal(err)
//     }

//     var version string

//     err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

//     if err2 != nil {
//         log.Fatal(err2)
//     }

//     fmt.Print(version)
// }

// package main

// import (
// 	"log"
// 	"net/http"

// 	"github.com/gorilla/mux"
// )

// func YourHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte("Gorilla!\n"))
// }

// func main() {
// 	r := mux.NewRouter()
// 	// Routes consist of a path and a handler function.
// 	r.HandleFunc("/", YourHandler)

// 	// Bind to a port and pass our router in
// 	log.Fatal(http.ListenAndServe(":8000", r))
// }
