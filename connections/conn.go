package connections

import(
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"tutorial2/databases"
	"github.com/joho/godotenv"
)

// DB Global Variable
var DB *gorm.DB
var err interface{}

func init() {
	errenv := godotenv.Load()
	if errenv != nil {
		panic("Error loading .env file")
	}

	dbHost := os.Getenv("HOST")
	dbName := os.Getenv("DATABASE")
	dbUser := os.Getenv("USER")
	dbPass := os.Getenv("PASSWORD")

	dsn := ``+ dbUser +`:`+ dbPass +`@tcp(`+ dbHost +`)/`+ dbName +`?charset=utf8mb4&parseTime=True&loc=Local`
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
    	panic("failed to connect database")
  	}else{
		  fmt.Println("Connected")
	}

	migrateShcemaDB := map[string]interface{}{
		"Products" 		: &databases.Product{},
		"Users" 		: &databases.User{},
		"Author" 		: &databases.Author{},
	}

	for key, value := range migrateShcemaDB {
		check := DB.Migrator().HasTable(value) 

		fmt.Println(`Database has `+ key + ` =`, check)
		DB.AutoMigrate(value)
		fmt.Println(`Migrate Database `+ key)
	}

}

