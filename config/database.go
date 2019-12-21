package config

// DatabaseHostConfig struct
type DatabaseHostConfig struct {
	Host string
	Port int
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
