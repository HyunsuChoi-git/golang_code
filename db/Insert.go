package db

import (
	"bitbucket.org/okestrolab/baton-om-sdk/btodb"
	"log"
)

func InsertOne(data *btodb.BtResource) error {
	fnc := "InsertOne"

	// check
	if data == nil {
		return nil
	}

	// execute
	res := Dblinker.DB.Create(data)
	if res.Error != nil {
		return res.Error
	}

	log.Printf("%s: completed: affCnt(%d)", fnc, res.RowsAffected)
	return nil
}

func InsertList(dataList []*btodb.BtResourceAttr) error {
	fnc := "InsertList"

	// check
	if len(dataList) == 0 {
		return nil
	}

	// execute
	res := Dblinker.DB.Create(dataList)
	if res.Error != nil {
		return res.Error
	}

	log.Printf("%s: completed: affCnt(%d)", fnc, res.RowsAffected)
	return nil
}
