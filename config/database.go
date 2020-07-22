package config

var (
	_connection = "mysql"
	_host       = "database"
	_port       = 3306
	_database   = "api"
	_username   = "api"
	_password   = "ba3ljA1M25hepA"
	_charset    = "utf8"
	_collation  = "utf8_general_ci"
)

// CreateDatabaseClusterConfig create database cluster config
func CreateDatabaseClusterConfig() *DatabaseCluster {
	cfg := &DatabaseCluster{
		Driver:    _connection,
		Database:  _database,
		Username:  _username,
		Password:  _password,
		Charset:   _charset,
		Collation: _collation,
	}

	cfg.Write = &DatabaseHostConfig{
		Host: _host,
		Port: _port,
	}

	cfg.Read = &[]DatabaseHostConfig{
		{
			Host: _host,
			Port: _port,
		},
		{
			Host: _host,
			Port: _port,
		},
	}

	return cfg
}

// DatabaseCluster struct
type DatabaseCluster struct {
	Driver    string
	Database  string
	Username  string
	Password  string
	Charset   string
	Collation string
	Write     *DatabaseHostConfig
	Read      *[]DatabaseHostConfig
}

// DatabaseHostConfig struct
type DatabaseHostConfig struct {
	Host string
	Port int
}
