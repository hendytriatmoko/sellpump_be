package databases

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"merk_microservices/common"
)

type PostgresDBSellPump struct {
	DB *gorm.DB
}

func (dbS *PostgresDBSellPump) Init() error {

	db, err := gorm.Open("postgres", "host="+common.Config.DbAddrsWarnaCRMPublic+" port="+common.Config.DBPortWarnaCRM+" user="+common.Config.DbUsernameWarnaCRM+" dbname="+common.Config.DbNameWarnaCRM+" sslmode=disable password="+common.Config.DbPasswordWarnaCRM+"")

	db.DB().Ping()
	db.DB().SetMaxIdleConns(0)

	if err != nil {
		return err
	}

	dbS.DB = db

	return err
}
