package connection

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func SetupDB() *sql.DB {

	/*path_dir := ("snippetBox-microservice/basket")*/
	/*err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	/*er := godotenv.Load("snippetBox-microservice/basket/.env")*/

	/*if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbDriver := os.Getenv("DB_DRIVER")
	dbUser := os.Getenv("DB_USERNAME")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")*/
	db, err := sql.Open("mysql", "root"+":duman070601"+""+"@/"+"articles")
	if err != nil {
		panic(err.Error())
	}
	return db
}
