package database

// import (
// 	"go.uber.org/zap"
// 	"gorm.io/driver/mysql"
// 	"gorm.io/gorm"
// 	// "gorm.io/driver/mysql"
// 	// "gorm.io/gorm"
// )

// func ConnectDatabase(user, host, password, port, db string) *gorm.DB {
// 	logger, _ := zap.NewDevelopment()
// 	defer logger.Sync()
// 	// log = logger.Sugar()

// 	// logger.N

// 	// newLogger := logger.New(
// 	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
// 	// 	logger.Config{
// 	// 		SlowThreshold: time.Second,
// 	// 		LogLevel:      logger.Info,
// 	// 		Colorful:      true,
// 	// 	},
// 	// )
// 	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + db + "?parseTime=true"
// 	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
// 		// Logger: newLogger,
// 	})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	return database
// }
