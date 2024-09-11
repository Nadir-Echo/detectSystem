package models

//database
import (
	"fmt"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"Demo/pkg/setting"
)

var db *gorm.DB

type Model struct {
	ResultId    int `gorm:"primary_key" json:"result_id"`
	CreatedTime int `json:"created_time"`
	//ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                       error
		dbName, user, password, host, tablePrefix string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	//dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//用户名:密码@tcp(IP:port)/数据库?charset=utf8mb4&parseTime=True&loc=Local
	datasorcename := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)
	//连接额外配置信息
	gormConfig := &gorm.Config{
		//用于打印sql语句
		Logger: logger.Default.LogMode(logger.Info),
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix, //表前缀
			SingularTable: true,        //使用单数表名，启用该选项时，`User` 的表名应该是 `user`而不是users
		},
	}

	db, err = gorm.Open(mysql.Open(datasorcename), gormConfig)

	if err != nil {
		log.Println(err)
	}

	//表前缀
	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return tablePrefix + defaultTableName
	//}
	sqlDB, err := db.DB()
	if err != nil {
		log.Panicln(err.Error())

	}

	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Hour)
}

// 获取 db *gorm.DB
func GetDB() *gorm.DB {
	return db
}

//gorm自己维护了一个连接池,无需使用close函数关闭函数
//func CloseDB() {
//	defer Db.Close()
//}
