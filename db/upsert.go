package db

import (
	"bitbucket.org/okestrolab/baton-om-sdk/btodb"
	"gorm.io/gorm/clause"
)

func UpsertOne(data btodb.BtResource, chanVals map[string]interface{}) error {
	res := Dblinker.DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.Assignments(chanVals),
	}).Create(&data)

	if res != nil {
		return res.Error
	}
	return nil
}
