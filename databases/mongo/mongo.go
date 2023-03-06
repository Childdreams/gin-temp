package mongo

import (
	"app/settings"
	"github.com/globalsign/mgo"
	"log"
	"os"
)

var Mongo *mgo.Session

var (
	MongodbUri string
)

func init() {
	settings.RequireEnvs([]string{
		"MONGODB_URI",
	})

	MongodbUri = os.Getenv("MONGODB_URI")

	var err error
	Mongo, err = mgo.Dial(MongodbUri)
	if err != nil {
		log.Fatalln(err)
	}
}
