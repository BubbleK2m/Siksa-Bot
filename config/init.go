package config

var (
	Mgo  map[string]string
	HTTP map[string]string
	JWT  map[string]string
)

func Init() {
	HTTP = make(map[string]string)
	HTTP["PORT"] = "5000"

	Mgo = make(map[string]string)
	Mgo["PATH"] = "mongodb://localhost"
	Mgo["DB"] = "quickstart"

	JWT = make(map[string]string)
	JWT["SECRET"] = "JWTSECRETKEY"
}
