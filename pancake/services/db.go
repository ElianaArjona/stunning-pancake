package services

import "fmt"

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "pancake"
)

type DataBase struct {
	Conector string `json:"conector,omitempty"`
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	User     string `json:"user,omitempty"`
	Password string `json:"password,omitempty"`
	DBname   string `json:"d_bname,omitempty"`
}

func (db *DataBase) CreateConnector() {

	db.Host = host
	db.Port = port
	db.User = user
	db.Password = password
	db.DBname = dbname

	db.Conector = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

}
