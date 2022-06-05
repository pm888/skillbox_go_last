package service

import (
	"encoding/json"
	"mymod/internal/data"
	"mymod/internal/method"
)

var DataSupport = make([]*data.SupportData, 0)

func GetSupport() []int {
	var totalTicets int
	var loadSupport int
	body := method.GetBody(data.UrlSuport)
	if err := json.Unmarshal(body, &DataSupport); err != nil {
	}

	for i := range DataSupport {
		totalTicets += DataSupport[i].ActiveTickets

	}
	expectation := float64(totalTicets) * float64(60/18)
	switch {
	case totalTicets < 9:
		loadSupport = 1
	case 9 <= totalTicets && totalTicets <= 16:
		loadSupport = 2
	case totalTicets > 16:
		loadSupport = 3
	}
	return []int{loadSupport, int(expectation)}

}
