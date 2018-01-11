package bot

import (
	"log"
	"time"

	"github.com/DSMdongly/siksa-bot/app/model"
	"github.com/DSMdongly/siksa-bot/config"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

var (
	Client *twitter.Client
)

func Init() {
	cnf := oauth1.NewConfig(config.Twitter["CONSUMER_KEY"], config.Twitter["CONSUMER_SECRET"])
	tok := oauth1.NewToken(config.Twitter["ACCESS_TOKEN"], config.Twitter["ACCESS_SECRET"])

	Client = twitter.NewClient(cnf.Client(oauth1.NoContext, tok))
}

func Start() {
	go func(){
		log.Println("Bot Started Tracking Alarm")

		for {
			db := model.DB
			alms := []model.Alarm{}
			
			now := time.Now()
			now = time.Date(2000, 6, 8, 1, 40, 0, 0, time.UTC)

			log.Println("Track ", now)

			if err := db.Where("time = ?", now).Find(&alms).Error; err != nil {
				break
			}
	
			for _, alm := range alms {
				if _, _, err := Client.Statuses.Update(alm.Note, nil); err != nil {
					log.Println("Error Occured ", err)
					break
				}
				
				log.Println("Sent Alarm ", alm.Note)
			}

			time.Sleep(1 * time.Second)
		}
	}()
}