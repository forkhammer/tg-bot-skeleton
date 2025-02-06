package bot

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/getsentry/sentry-go"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ConnectionConfig struct {
	host   string
	port   int
	dbName string
	user   string
	pass   string
}

type DbConnection struct {
	config ConnectionConfig
	Db     *gorm.DB
}

func NewDbConnection(host string, port int, dbName string, user string, pass string) (*DbConnection, error) {
	connection := &DbConnection{
		config: ConnectionConfig{
			host:   host,
			port:   port,
			dbName: dbName,
			user:   user,
			pass:   pass,
		},
	}

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  1,           // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  connection.getDSN(),
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	connection.Db = db

	if err := connection.migrate(); err != nil {
		sentry.CaptureException(err)
		panic(err)
	}

	return connection, nil
}

func (c *DbConnection) getDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=UTC", c.config.host, c.config.port, c.config.user, c.config.pass, c.config.dbName)
}

func (c *DbConnection) migrate() error {
	return c.Db.AutoMigrate(&User{})
}
