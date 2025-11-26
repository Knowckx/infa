package db

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/cockroachdb/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GenDefaultConfig() *gorm.Config {
	logConfig := logger.Config{
		SlowThreshold: time.Second,  // Slow SQL threshold
		LogLevel:      logger.Error, // Log level
		Colorful:      false,        // Disable color
	}
	log := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logConfig,
	)
	out := &gorm.Config{
		Logger:                 log,
		SkipDefaultTransaction: true,
	}
	return out
}

type PGConnInfo struct {
	HOST     string
	PORT     string
	DATABASE string
	USER     string
	PASWD    string
	SSLMode  string
}

func (in PGConnInfo) String() string {
	buf := bytes.Buffer{}
	baseInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s ", in.HOST, in.PORT, in.USER, in.PASWD, in.DATABASE)
	buf.WriteString(baseInfo)
	info := fmt.Sprintf("application_name=%s TimeZone=Asia/Shanghai ", "gorm")
	buf.WriteString(info)

	if in.SSLMode == "" {
		buf.WriteString("sslmode=disable ")
	} else {
		buf.WriteString("sslmode=prefer ")
	}
	dsn := buf.String()
	return dsn
}

func ConnPG(in *PGConnInfo) (db *gorm.DB, err error) {
	dsn := in.String()
	d := postgres.Open(dsn)
	pgConfig := GenDefaultConfig()
	db, err = gorm.Open(d, pgConfig)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return db, nil
}

// local
func NewPGConnInfoLocal() *PGConnInfo {
	tdb := &PGConnInfo{}
	tdb.HOST = "127.0.0.1"
	tdb.PORT = "5432"
	tdb.DATABASE = "postgres"
	return tdb
}
