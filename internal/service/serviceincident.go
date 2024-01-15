package service

import (
	"encoding/json"
	"sort"

	"mymod/internal/data"
	"mymod/internal/helpers"
)

func Incident() ([]data.IncidentData, error) {
	var DataIncident = make([]data.IncidentData, 0)
	body, err := helpers.GetBody(data.UrlAccendent)
	if err := json.Unmarshal(body, &DataIncident); err != nil {
		return nil, err
	}

	sort.Slice(DataIncident, func(i, j int) (less bool) {
		return DataIncident[i].Status < DataIncident[j].Status
	})
	return DataIncident, err

}
