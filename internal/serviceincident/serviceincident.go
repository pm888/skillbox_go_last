package serviceincident

import (
	"encoding/json"
	"mymod/internal/data"
	"mymod/internal/method"
	"sort"
)

var DataIncident = make([]data.IncidentData, 0)

func Incident() []data.IncidentData {
	body := method.GetBody(data.UrlAccendent)
	if err := json.Unmarshal(body, &DataIncident); err != nil {
	}

	sort.Slice(DataIncident, func(i, j int) (less bool) {
		return DataIncident[i].Status < DataIncident[j].Status
	})
	return DataIncident

}
