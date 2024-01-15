package service

import (
	"encoding/json"
	"sort"

	"mymod/internal/data"
	"mymod/internal/helpers"
)

func GetMMS() ([][]data.MMSData, error) {
	var DataMMS = make([]data.MMSData, 0)
	body, err := helpers.GetBody(data.UrlMMS)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &DataMMS); err != nil {
		return nil, err
	}

	// Функция для проверки наличия элемента в мапе
	hasCountry := func(country string) bool {
		_, ok := helpers.Alfa2Data[country]
		return ok
	}

	// Функция для фильтрации элементов
	filterMMS := func(apps []data.MMSData) []data.MMSData {
		result := make([]data.MMSData, 0, len(apps))
		for i := range apps {
			if helpers.IsValidSmsMmsProvider(apps[i].Provider) && hasCountry(apps[i].Country) {
				result = append(result, apps[i])
			}
		}
		return result
	}

	DataMMS = filterMMS(DataMMS)

	var sliceSliceMms = make([][]data.MMSData, 2)
	var sliceCopy = make([]data.MMSData, len(DataMMS))
	copy(sliceCopy, DataMMS)

	// Сортировка исходного среза
	sort.Slice(DataMMS, func(i, j int) (less bool) {
		return DataMMS[i].Provider < DataMMS[j].Provider
	})
	sliceSliceMms[0] = DataMMS

	// Сортировка копии среза
	sort.Slice(sliceCopy, func(i, j int) (less bool) {
		return sliceCopy[i].Country < sliceCopy[j].Country
	})
	sliceSliceMms[1] = sliceCopy

	return sliceSliceMms, nil
}
