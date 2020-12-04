package repositories

import (
	"database/sql"
	"log"

	"github.com/Matias-Barrios/GoUnitTestingExample/configproviders"
	_ "github.com/lib/pq" // postgres golang driver
)

type IDBProvider interface {
	Connect()
	Close() error
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var instance *DBProvider

type DBProvider struct {
	db                          *sql.DB
	environmentVariableProvider configproviders.IEnvironmentVariableProvider
}

func (db DBProvider) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.db.Query(query, args...)
}

func (db DBProvider) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.db.QueryRow(query, args...)
}

func (db DBProvider) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.db.Exec(query, args...)
}
func (db *DBProvider) Connect() {
	uri := "postgres://postgres:" + db.environmentVariableProvider.Get("DBPASSWORD") + "@localhost/gotest?sslmode=disable"
	database, err := sql.Open("postgres", uri)
	if err != nil {
		log.Fatal(err.Error())
	}
	db.db = database
}

func (db DBProvider) Close() error {
	return db.db.Close()
}

func NewDBProvider(evp configproviders.IEnvironmentVariableProvider) IDBProvider {
	if instance == nil {
		instance = &DBProvider{
			environmentVariableProvider: evp,
		}
		instance.Connect()
		if err := instance.db.Ping(); err != nil {
			log.Fatalln(err.Error())
		}
	}
	return instance
}

func NewMockDBProvider(data *sql.DB, evp configproviders.IEnvironmentVariableProvider) IDBProvider {
	return &DBProvider{
		environmentVariableProvider: evp,
		db:                          data,
	}
}
