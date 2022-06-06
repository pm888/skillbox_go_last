package api

import (
	"encoding/json"
	"fmt"
	"log"
	"mymod/internal/data"
	"mymod/internal/get"
)

func GetApi() []byte {
	var js []byte
	var err error
	var resT data.ResultT
	var resSet data.ResultSetT
	var resSetnil data.ResultSetT
	resSet = *get.GetResultData()
	//resSet.SMS = nil
	fmt.Println(resT)
	if resSet.SMS != nil && resSet.MMS != nil &&
		resSet.VoiceCall != nil &&
		resSet.Email != nil &&
		resSet.Billing != nil &&
		resSet.Support != nil &&
		resSet.Incidents != nil {
		resT.Status = true
		resT.Data = resSet
		js, err = json.Marshal(resT)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		resT.Status = false
		resT.Error = "Error on collect data"
		resT.Data = resSetnil
		js, err = json.Marshal(resT)
		if err != nil {
			log.Fatal(err)
		}

	}

	return js

}
