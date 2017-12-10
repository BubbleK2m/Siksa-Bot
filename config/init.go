package config

var (
	HTTP   map[string]string
	JWT    map[string]string
	Sqlite map[string]string
)

func Init() {
	HTTP = make(map[string]string)
	HTTP["PORT"] = "5000"

	JWT = make(map[string]string)
	JWT["SECRET"] = "JWTSECRETKEY"

	Sqlite = make(map[string]string)
	Sqlite["PATH"] = "glove.db"
}
