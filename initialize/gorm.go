package initialize

import (
	"fmt"
	"gin-blog-server/dao"
	"gin-blog-server/global"
	"log"

	"gorm.io/gorm/schema"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitGorm() {
	cfg := global.AppCfg.Mysql
	gormCfg := global.AppCfg.GormCfg
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database, cfg.CharSet,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   gormCfg.TablePrefix,
			SingularTable: true,
		},
	})
	if err != nil {
		log.Fatal("Connect to db failed: ", err.Error())
	}
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(gormCfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(gormCfg.MaxOpenConns)
	dao.Init(db)
}
