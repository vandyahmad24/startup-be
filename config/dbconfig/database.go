package dbconfig

import "fmt"

type DatabaseList struct {
	Startup struct {
		Mysql Database
	}
}

type Database struct {
	Host     string
	Port     string
	Username string
	Password string
	Dbname   string
	Adapter  string
}

func Halo() {
	fmt.Println("Halo database")
}
