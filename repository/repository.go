// Copyright 2023 The GoStartKit Authors. All rights reserved.
// Use of this source code is governed by a AGPL
// license that can be found in the LICENSE file.
// https://gostartkit.com
package repository

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gostartkit/api/config"
	"github.com/gostartkit/api/helper"
	"github.com/webpkg/web"
)

const (
	_pageSize    = 10
	_maxPageSize = 1000
)

var (
	_writeDatabase *sql.DB
	_readDatabases []*sql.DB
	_once          sync.Once
)

// Init config
func Init() error {
	var err error

	_once.Do(func() {
		_writeDatabase, err = openWriteDatabase()

		if err == nil {
			_readDatabases, err = openReadDatabases()
		}
	})

	return err
}

// Close databases
func Close() error {
	var err error

	if _writeDatabase != nil {
		if err1 := _writeDatabase.Close(); err1 != nil {
			err = err1
		}
	}

	for _, r := range _readDatabases {
		if r != nil {
			if err1 := r.Close(); err1 != nil {
				err = err1
			}
		}
	}

	return err
}

// openWriteDatabase of config.database.write
func openWriteDatabase() (*sql.DB, error) {
	db, err := sql.Open(config.Database().Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
		config.Database().Username,
		config.Database().Password,
		config.Database().Write.Host,
		config.Database().Write.Port,
		config.Database().Database,
		config.Database().Charset))

	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(0)

	return db, nil
}

// openReadDatabases of config.database.read
func openReadDatabases() ([]*sql.DB, error) {

	var readDatabases []*sql.DB

	for _, r := range *config.Database().Read {
		db, err := sql.Open(config.Database().Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=true",
			config.Database().Username,
			config.Database().Password,
			r.Host,
			r.Port,
			config.Database().Database,
			config.Database().Charset))

		if err != nil {
			return nil, err
		}

		db.SetMaxIdleConns(0)

		readDatabases = append(readDatabases, db)
	}

	return readDatabases, nil
}

func selectDB(databaseCluster []*sql.DB) *sql.DB {
	return databaseCluster[helper.RandMax(len(databaseCluster))]
}

func query(query string, args ...interface{}) (*sql.Rows, error) {

	rows, err := selectDB(_readDatabases).Query(query, args...)

	if err != nil {
		if config.App().AppDebug {
			log.Printf("query: %s\n", query)
		}
		return nil, err
	}

	return rows, nil
}

func queryRow(query string, args ...interface{}) *sql.Row {
	return selectDB(_readDatabases).QueryRow(query, args...)
}

func prepare(query string) (*sql.Stmt, error) {

	stmt, err := _writeDatabase.Prepare(query)

	if err != nil {
		if config.App().AppDebug {
			log.Printf("prepare: %s\n", query)
		}
		return nil, err
	}

	return stmt, nil
}

func txPrepare(tx *sql.Tx, query string) (*sql.Stmt, error) {

	stmt, err := tx.Prepare(query)

	if err != nil {
		if config.App().AppDebug {
			log.Printf("txPrepare: %s\n", query)
		}
		return nil, err
	}

	return stmt, nil
}

func stmtExec(stmt *sql.Stmt, args ...interface{}) (sql.Result, error) {

	result, err := stmt.Exec(args...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func exec(query string, args ...interface{}) (sql.Result, error) {

	result, err := _writeDatabase.Exec(query, args...)

	if err != nil {
		if config.App().AppDebug {
			log.Printf("exec: %s\n", query)
			log.Printf("args: %v\n", args)
		}
		return nil, err
	}

	return result, nil
}

func begin() (*sql.Tx, error) {
	return _writeDatabase.Begin()
}

func txExec(tx *sql.Tx, query string, args ...interface{}) (sql.Result, error) {

	result, err := tx.Exec(query, args...)

	if err != nil {
		if config.App().AppDebug {
			log.Printf("txExec: %s\n", query)
			log.Printf("args: %v\n", args)
		}
		return nil, err
	}

	return result, nil
}

func rollback(tx *sql.Tx) error {
	return tx.Rollback()
}

func commit(tx *sql.Tx) error {
	return tx.Commit()
}

func now() *time.Time {
	return web.Now()
}

// max get max key value
func max(tableName string, key string, appID uint64, appNum uint64) (uint64, error) {

	sqlx := "SELECT MAX(`" + key + "`) FROM `" + tableName + "` WHERE `" + key + "` % ? = ? "

	row := queryRow(sqlx, appNum, appID)

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
