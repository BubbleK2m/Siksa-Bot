package config

import "fmt"

var (
	HTTP     map[string]string
	JWT      map[string]string
	Postgres map[string]string
)

func Init() {
	HTTP = make(map[string]string)
	HTTP["PORT"] = "8000"

	JWT = make(map[string]string)
	JWT["SECRET"] = "8FC536E4DE2ABADF47558D9F2C2E4"

	Postgres = make(map[string]string)
	Postgres["HOST"] = "localhost"
	Postgres["USER"] = "root"
	Postgres["PASSWORD"] = "ehdgus0608"
	Postgres["DB"] = "glove"
	Postgres["PATH"] = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", Postgres["HOST"], Postgres["USER"], Postgres["DB"], Postgres["PASSWORD"])
}
