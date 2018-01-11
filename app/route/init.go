package route

import "github.com/DSMdongly/siksa-bot/app"

func Init() {
	User(app.Echo)
	Alarm(app.Echo)
}
