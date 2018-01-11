package config

import (
	"fmt"
	"net"
	"net/url"
	"os"
	
	"github.com/subosito/gotenv"
)

var (
	HTTP     map[string]string
	JWT      map[string]string
	Postgres map[string]string
	Twitter  map[string]string
)

func Init() {
	gotenv.Load()

	HTTP = make(map[string]string)
	HTTP["PORT"] = os.Getenv("PORT")

	JWT = make(map[string]string)
	JWT["SECRET"] = os.Getenv("JWT_SECRET")

	pth, _ := url.Parse(os.Getenv("DATABASE_URL"))

	Postgres = make(map[string]string)
	Postgres["HOST"], _, _ = net.SplitHostPort(pth.Host)
	Postgres["USER"] = pth.User.Username()
	Postgres["PASSWORD"], _ = pth.User.Password()
	Postgres["DB"] = pth.Path[1:]
	Postgres["PATH"] = fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", Postgres["HOST"], Postgres["USER"], Postgres["DB"], Postgres["PASSWORD"])

	Twitter = make(map[string]string)
	Twitter["CONSUMER_KEY"] = os.Getenv("TWITTER_CONSUMER_KEY")
	Twitter["CONSUMER_SECRET"] = os.Getenv("TWITTER_CONSUMER_SECRET")
	Twitter["ACCESS_TOKEN"] = os.Getenv("TWITTER_ACCESS_TOKEN")
	Twitter["ACCESS_SECRET"] = os.Getenv("TWITTER_ACCESS_SECRET")
}
