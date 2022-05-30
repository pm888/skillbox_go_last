package servicemms

import (
	"encoding/json"
	"mymod/internal/data"
	"mymod/internal/method"
	"sort"
)

var DataMMS = make([]data.MMSData, 0)

func RemoveApplications(apps []data.MMSData, i int) []data.MMSData {
	apps = append(apps[:i], apps[i+1:]...)
	return apps
}

func GetMMS() [][]data.MMSData {
	body := method.GetBody(data.UrlMMS)
	if err := json.Unmarshal(body, &DataMMS); err != nil {
	}

	for i := len(DataMMS) - 1; i != 0; i-- {
		if !method.IsValidSmsMmsProvider(DataMMS[i].Provider) {
			DataMMS = RemoveApplications(DataMMS, i)
			continue
		}

		if _, ok := method.Alfa2Data[DataMMS[i].Country]; !ok {
			DataMMS = RemoveApplications(DataMMS, i)

		}
	}

	var sliceSliceMms = make([][]data.MMSData, 2)
	var sliceCopy = make([]data.MMSData, len(DataMMS))
	sort.Slice(DataMMS, func(i, j int) (less bool) {
		return DataMMS[i].Provider < DataMMS[j].Provider
	})
	sliceSliceMms[0] = DataMMS

	copy(sliceCopy, DataMMS)

	sort.Slice(sliceCopy, func(i, j int) (less bool) {
		return sliceCopy[i].Country < sliceCopy[j].Country
	})
	sliceSliceMms[1] = sliceCopy

	return sliceSliceMms

}
