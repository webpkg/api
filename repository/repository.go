package repository

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/webpkg/api/config"
)

var (
	_databaseCluster *config.DatabaseCluster
	_once            sync.Once
)

// Init config
func Init(dc *config.DatabaseCluster) {

	_once.Do(func() {
		_databaseCluster = dc
	})
}

func open(cfg *config.DatabaseCluster) (*sql.DB, error) {
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
