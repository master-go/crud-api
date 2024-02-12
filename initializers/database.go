package initializers

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() *gorm.DB {
	var err error
	// postgres://iwhdhaot:1Jj4BVNbKfxtSHDipfSIr_wndNs_qnto@silly.db.elephantsql.com/iwhdhaot
	DB, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  os.Getenv("DB_URL"), // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect with database")
	}
	return DB

}
