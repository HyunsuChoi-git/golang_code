package main

import (
	"bitbucket.org/okestrolab/baton-om-sdk/btodb"
	"golang_code/db"
	"golang_code/test"
	"time"
)

func main() {
	test.OpenstackTest()

	//exit := make(chan os.Signal, 1)
	//<-exit
}

func dbTest() {
	db.Init(db.DbTyps, db.Host, db.Port, db.DbName, db.UserName, db.UserPassword, db.Encoding, db.ConnTimeout)

	data := btodb.BtResource{
		ID:             1,
		ProviderID:     9,
		ResourceTypeID: 1,
		UuID:           "testtesttesttesttesttesttest",
		Name:           "test",
		Description:    "test",
	}

	chanVals := map[string]interface{}{
		"description": "upsert 테스트",
		"mod_dt":      time.Now(),
	}

	db.UpsertOne(data, chanVals)
}
