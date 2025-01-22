package database

var connection string

func init() {
	connection = "Mysql\n"
}

func GetDatabase() string {
	return connection
}
