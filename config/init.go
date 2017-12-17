package config

import "fmt"

var (
	HTTP     map[string]string
	JWT      map[string]string
	Postgres map[string]string
)

func Init() {
	HTTP = make(map[string]string)
	HTTP["PORT"] = "your http ports"

	JWT = make(map[string]string)
	JWT["SECRET"] = "your jwt keys"

	Postgres = make(map[string]string)
	Postgres["HOST"] = "your postgres host"
	Postgres["USER"] = "your postgres user name"
	Postgres["PASSWORD"] = "your postgres user password"
	Postgres["DB"] = "your postgres db name"
	Postgres["PATH"] = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", Postgres["HOST"], Postgres["USER"], Postgres["DB"], Postgres["PASSWORD"])
}
