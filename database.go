package hello_go

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func databaseGo() int {
	return 125
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func databaseOpen() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		id   int
		name string
	)

	//	rows, err := db.Query("select id, name from users where id = 1", 1)
	rows, err := db.Query("select id, name from users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		//		fmt.Println(rows)
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
