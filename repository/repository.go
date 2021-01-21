package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/webpkg/api/config"
	"github.com/webpkg/api/helper"
)

const (
	_pageSize    = 20
	_maxPageSize = 1000
)

var (
	_webConfig     *config.WebConfig
	_writeDatabase *sql.DB
	_readDatabases []*sql.DB
	_once          sync.Once
)

// Init config
func Init(WebConfig *config.WebConfig) {
	_once.Do(func() {
		_webConfig = WebConfig
		_writeDatabase = openWriteDatabase()
		_readDatabases = openReadDatabases()
	})
}

// WebConfig get WebConfig
func WebConfig() *config.WebConfig {
	return _webConfig
}

// Close databases
func Close() error {

	if _writeDatabase != nil {
		err := _writeDatabase.Close()
		log.Printf("close: %v\n", err)
	}

	for i, r := range _readDatabases {
		if r != nil {
			err := r.Close()
			log.Printf("close(%d): %v\n", i, err)
		}
	}

	return nil
}

// openWriteDatabase of config.database.write
func openWriteDatabase() *sql.DB {
	db, err := sql.Open(_webConfig.Database.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		_webConfig.Database.Username,
		_webConfig.Database.Password,
		_webConfig.Database.Write.Host,
		_webConfig.Database.Write.Port,
		_webConfig.Database.Database,
		_webConfig.Database.Charset))

	if err != nil {
		log.Fatalf("openWriteDatabase: %s\n", err)
	}

	return db
}

// openReadDatabases of config.database.read
func openReadDatabases() []*sql.DB {

	var readDatabases []*sql.DB

	for _, r := range *_webConfig.Database.Read {
		db, err := sql.Open(_webConfig.Database.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
			_webConfig.Database.Username,
			_webConfig.Database.Password,
			r.Host,
			r.Port,
			_webConfig.Database.Database,
			_webConfig.Database.Charset))

		if err != nil {
			log.Printf("openReadDatabases(%s:%d): %v\n", r.Host, r.Port, err)
			continue
		}

		readDatabases = append(readDatabases, db)
	}

	return readDatabases
}

func selectDB(databaseCluster []*sql.DB) *sql.DB {
	return databaseCluster[helper.RandMax(len(databaseCluster))]
}

func query(query string, args ...interface{}) (*sql.Rows, error) {
	return selectDB(_readDatabases).Query(query, args...)
}

func queryRow(query string, args ...interface{}) *sql.Row {
	return selectDB(_readDatabases).QueryRow(query, args...)
}

func prepare(query string) (*sql.Stmt, error) {
	return _writeDatabase.Prepare(query)
}

func txPrepare(tx *sql.Tx, query string) (*sql.Stmt, error) {
	return tx.Prepare(query)
}

func stmtExec(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {
	return stmt.Exec(args...)
}

func exec(query string, args ...interface{}) (sql.Result, error) {
	return _writeDatabase.Exec(query, args...)
}

func begin() (*sql.Tx, error) {
	return _writeDatabase.Begin()
}

func txExec(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {
	return tx.Exec(query, args...)
}

func rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func commit(tx *sql.Tx) error {
	return tx.Commit()
}

func now() *time.Time {
	return helper.Now()
}

// max get max key value
func max(tableName string, key string) (uint64, error) {

	sqlx := "SELECT MAX(`" + key + "`) FROM `" + tableName + "` WHERE `" + key + "` % ? = ? "

	row := queryRow(sqlx, WebConfig().App.AppNum, WebConfig().App.AppID)

	var val *uint64

	err := row.Scan(&val)

	if err != nil {
		return 0, err
	}

	if val == nil {
		return 0, nil
	}

	return *val, nil
}
