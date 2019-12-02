package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/webpkg/web"
)

var (
	_databaseCluster *web.DatabaseCluster
	_once            sync.Once
)

// Init config
func Init(dc *web.DatabaseCluster) {

	_once.Do(func() {
		_databaseCluster = dc
	})
}

func open(cfg *web.DatabaseCluster) (*sql.DB, error) {
	db, err := sql.Open(cfg.Driver, fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",
		cfg.Username,
		cfg.Password,
		cfg.Write.Host,
		cfg.Write.Port,
		cfg.Database,
		cfg.Charset))

	if err != nil {
		return nil, err
	}

	return db, nil
}
