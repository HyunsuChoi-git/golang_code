package db

import (
	"bitbucket.org/okestrolab/baton-om-sdk/btodb"
	"github.com/rs/zerolog/log"
)

func UpdateOne(pkey int64, data btodb.BtResource) error {

	res := Dblinker.DB.Model(&btodb.BtResourceAttr{}).
		Where("id = ?", pkey).Updates(data)
	if res.Error != nil {
		return res.Error
	}

	return nil
}
func UpdateList(dataList []*btodb.BtResourceAttr) error {
	// check
	if len(dataList) == 0 {
		return nil
	}

	// execute
	for _, rcAttr := range dataList {
		res := Dblinker.DB.Model(&btodb.BtResourceAttr{}).Where("resource_id = ? AND key = ?", "abc", "abc").
			Updates(btodb.BtResourceAttr{Value: rcAttr.Value, ModDt: rcAttr.ModDt})

		if res.Error != nil {
			log.Error().Msgf("resource attr update failed: rc(%v) : %s", rcAttr, res.Error)
			continue
		}
	}

	return nil
}
