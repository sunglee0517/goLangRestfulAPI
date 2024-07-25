// config/db.go

package config

const (
	DBUser     = "root"
	DBPassword = "1234"
	DBName     = "company"
	DBHost     = "localhost"
	DBPort     = "3307"
)

func GetDBConnectionString() string {
	return DBUser + ":" + DBPassword + "@tcp(" + DBHost + ":" + DBPort + ")/" + DBName + "?parseTime=true"
}
