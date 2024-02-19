package db

import (
	"fmt"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"strings"
)

var Dblinker dbLinker

type dbLinker struct {
	Info dbInfo
	DB   *gorm.DB
}

type dbInfo struct {
	kind           string
	ip             string
	port           int
	name           string
	authId         string
	authPwd        string
	encoding       string
	dataSrcName    string
	connTimeoutSec int
	maxIdleConns   int
	maxOpenConns   int
}

func Init(kind string, ip string, port int, name string, authId string, authPwd string, encoding string, connTimeoutSec int) error {
	fnc := "db_handler_Init"
	err := error(nil)

	{
		Dblinker.Info.kind = kind
		Dblinker.Info.ip = ip
		Dblinker.Info.port = port
		Dblinker.Info.name = name
		Dblinker.Info.authId = authId
		Dblinker.Info.authPwd = authPwd
		Dblinker.Info.encoding = encoding
		Dblinker.Info.connTimeoutSec = connTimeoutSec
		Dblinker.Info.maxIdleConns = 2
		Dblinker.Info.maxOpenConns = 50

		if strings.ToLower(kind) == "postgres" {
			Dblinker.Info.dataSrcName = fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Dblinker.Info.ip, Dblinker.Info.port, Dblinker.Info.authId, Dblinker.Info.authPwd, Dblinker.Info.name)
		}
	}

	{

		if Dblinker.DB, err = gorm.Open(postgres.Open(Dblinker.Info.dataSrcName), &gorm.Config{}); err != nil {
			log.Error().Msgf("%s: gorm.Open failed: dsn(%s): %s", fnc, Dblinker.Info.dataSrcName, err.Error())
			return err
		}

		if sqlDB, err := Dblinker.DB.DB(); err != nil {
			log.Error().Msgf("%s: DB() failed: %s", fnc, err.Error())
			return err
		} else {
			sqlDB.SetMaxIdleConns(Dblinker.Info.maxIdleConns)
			sqlDB.SetMaxOpenConns(Dblinker.Info.maxOpenConns)
			//sqlDB.SetConnMaxLifetime(time.Hour)
		}
	}

	return nil
}
