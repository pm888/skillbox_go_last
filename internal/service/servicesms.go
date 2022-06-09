package service

import (
	"fmt"
	"mymod/internal/data"
	"mymod/internal/method"
	"sort"
	"strings"
)

func FileSMS() ([][]data.SMSData, error) {
	var StorageDataSMS = make([]data.SMSData, 0)
	file, err := method.ReadFile(data.FileSmsRead)
	if err != nil {
		return nil, err
	}
	stringsTemp := strings.Split(string(file), "\n")
	for i := 0; i < len(stringsTemp)-1; i++ {
		sms := strings.Split(stringsTemp[i], ";")
		if len(sms) == 4 {
			for key := range method.Alfa2Data {
				if key == sms[0] && method.IsValidSmsMmsProvider(sms[3]) {
					newPerson := data.SMSData{
						Country:      sms[0],
						Bandwidth:    sms[1],
						ResponseTime: sms[2],
						Provider:     sms[3],
					}
					StorageDataSMS = append(StorageDataSMS, newPerson)
				}
			}
		}
	}
	fmt.Println("SMS", StorageDataSMS)

	for i := 0; i < len(StorageDataSMS); i++ {
		for key := range method.Alfa2Data {
			if StorageDataSMS[i].Country == key {
				StorageDataSMS[i].Country = method.Alfa2Data[key]

			} else {
				continue
			}
		}

	}
	var sliceSliceSms = make([][]data.SMSData, 2)
	var sliceCopy = make([]data.SMSData, len(StorageDataSMS))
	sort.Slice(StorageDataSMS, func(i, j int) (less bool) {
		return StorageDataSMS[i].Provider < StorageDataSMS[j].Provider
	})
	sliceSliceSms[0] = StorageDataSMS

	copy(sliceCopy, StorageDataSMS)

	sort.Slice(sliceCopy, func(i, j int) (less bool) {
		return sliceCopy[i].Country < sliceCopy[j].Country
	})
	sliceSliceSms[1] = sliceCopy
	return sliceSliceSms, err

}
