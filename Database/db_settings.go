package database

import "fmt"

var (
	Host        string = "localhost"
	User        string = "gorm_u"
	Database    string = "gorm_db"
	Db_password string = "plhfnenb"
	Port        string = "5432"
	Sslmode     string = "disable"
	Timezone    string = "Asia/Omsk"
	ConString          = fmt.Sprintf("host=%v user=%v database=%v password=%v port=%v sslmode=%v TimeZone=%v", Host, User, Database, Db_password, Port, Sslmode, Timezone)
)
