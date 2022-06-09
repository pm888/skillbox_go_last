package service

import (
	"mymod/internal/data"
	"mymod/internal/method"
	"sort"
	"strings"
)

func FileEmail() (map[string][][]data.EmailData, error) {
	var StorageDataEmailCall = make([]data.EmailData, 0)
	var temporaryMap = make(map[string][]data.EmailData)
	var emailMap = make(map[string][][]data.EmailData)
	file, err := method.ReadFile(data.FileEmailRead)
	if err != nil {
		return nil, err
	}
	stringsTemp := strings.Split(string(file), "\n")
	for i := 0; i < len(stringsTemp)-1; i++ {
		email := strings.Split(stringsTemp[i], ";")
		if len(email) == 3 {
			for key, _ := range method.Alfa2Data {
				if key == email[0] && method.IsValidEmailProvider(email[1]) {
					time, err := method.StringIntoInt(email[2])
					if err != nil {
						return nil, err
					}
					newPerson := data.EmailData{
						Country:      email[0],
						Provider:     email[1],
						DeliveryTime: time,
					}
					StorageDataEmailCall = append(StorageDataEmailCall, newPerson)
				}

			}

		}

	}
	for i := range StorageDataEmailCall {
		temporaryMap[StorageDataEmailCall[i].Country] = append(temporaryMap[StorageDataEmailCall[i].Country], StorageDataEmailCall[i])
	}
	for s := range temporaryMap {
		sort.Slice(temporaryMap[s], func(i, j int) (less bool) {
			return temporaryMap[s][i].DeliveryTime < temporaryMap[s][j].DeliveryTime
		})

	}
	for key, val := range temporaryMap {
		emailMap[key] = append(emailMap[key], []data.EmailData{val[0], val[1], val[3]})
		emailMap[key] = append(emailMap[key], []data.EmailData{val[len(val)-1], val[len(val)-2], val[len(val)-3]})

	}
	return emailMap, err

}
