package route

import "github.com/DSMdongly/glove/app"

func Init() {
	User(app.Echo)
	Note(app.Echo)
}
