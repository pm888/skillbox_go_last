package service

import (
	"encoding/json"
	"mymod/internal/data"
	"mymod/internal/method"
	"sort"
)

var DataIncident = make([]data.IncidentData, 0)

func Incident() ([]data.IncidentData, error) {
	body, err := method.GetBody(data.UrlAccendent)
	if err := json.Unmarshal(body, &DataIncident); err != nil {
		return nil, err
	}

	sort.Slice(DataIncident, func(i, j int) (less bool) {
		return DataIncident[i].Status < DataIncident[j].Status
	})
	return DataIncident, err

}
