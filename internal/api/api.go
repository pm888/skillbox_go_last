package api

import (
	"encoding/json"

	"mymod/internal/data"
	"mymod/internal/get"
)

func GetApi() ([]byte, error) {
	var js []byte
	var resT data.ResultT
	resSet, err := get.Result()
	if err != nil {
		return nil, err
	}
	if resSet.SMS != nil && resSet.MMS != nil &&
		resSet.VoiceCall != nil &&
		resSet.Email != nil &&
		resSet.Billing != nil &&
		resSet.Support != nil &&
		resSet.Incidents != nil {
		resT.Status = true
		resT.Data = *resSet
		js, err = json.Marshal(resT)
		if err != nil {
			return nil, err
		}
	} else {
		resT.Status = false
		resT.Error = "Error on collect data"
		resT.Data = *resSet
		js, err = json.Marshal(resT)
		if err != nil {
			return nil, err
		}

	}
	return js, err

}
