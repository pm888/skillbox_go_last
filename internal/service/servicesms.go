package service

import (
	"os"
	"sort"
	"strings"

	"mymod/internal/data"
	"mymod/internal/helpers"
)

func FileSMS() ([][]data.SMSData, error) {
	var storageDataSMS []data.SMSData
	file, err := os.ReadFile(data.FileSmsRead)
	if err != nil {
		return nil, err
	}

	stringsTemp := strings.Split(string(file), "\n")
	for _, smsString := range stringsTemp {
		sms := strings.Split(smsString, ";")
		if len(sms) < 4 || !helpers.IsValidSmsMmsProvider(sms[3]) {
			continue
		}

		if country, ok := helpers.Alfa2Data[sms[0]]; ok {
			newPerson := data.SMSData{
				Country:      country,
				Bandwidth:    sms[1],
				ResponseTime: sms[2],
				Provider:     sms[3],
			}
			storageDataSMS = append(storageDataSMS, newPerson)
		}
	}

	sort.Slice(storageDataSMS, func(i, j int) bool {
		if storageDataSMS[i].Provider != storageDataSMS[j].Provider {
			return storageDataSMS[i].Provider < storageDataSMS[j].Provider
		}
		return storageDataSMS[i].Country < storageDataSMS[j].Country
	})

	return [][]data.SMSData{storageDataSMS, append([]data.SMSData(nil), storageDataSMS...)}, nil
}
