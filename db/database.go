package db

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"

	_ "github.com/lib/pq"
)

type Database struct {
	conn *sql.DB
}

func (this *Database) Close() {
	this.conn.Close()
}

func NewDatabase(path, name string) (db Database, err error) {
	config, err := parseConfig(path, name)
	if err != nil {
		return
	}
	var conn *sql.DB
	if conn, err = sql.Open("postgres", config.String()); err != nil {
		return
	}

	if err = conn.Ping(); err != nil {
		//We are not connected to the db
		log.Fatal("No connection to db found...")
		return
	}

	return Database{conn: conn}, err
}

type Dbconfig struct {
	Database string `json:"database"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	SSL      string `json:"ssl_mode"`
}

func (config Dbconfig) String() string {
	dbConfig := "host=" + config.Host +
		" port=" + config.Port +
		" user=" + config.User +
		" dbname=" + config.Database +
		" sslmode=" + config.SSL
	if len(config.Password) > 0 {
		dbConfig += " password=" + config.Password
	}
	return dbConfig
}

func parseConfig(path, name string) (conf Dbconfig, err error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}

	var configs map[string]Dbconfig
	err = json.Unmarshal(data, &configs)
	if err != nil {
		return
	}
	conf = configs[name]
	return
}
