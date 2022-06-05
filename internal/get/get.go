package get

import (
	"mymod/internal/data"
	"mymod/internal/service"
)

func GetResultData() *data.ResultSetT {
	report := data.ResultSetT{
		SMS:       service.FileSMS(),
		MMS:       service.GetMMS(),
		VoiceCall: service.FileVoice(),
		Email:     service.FileEmail(),
		Billing:   service.ReadBilling(),
		Support:   service.GetSupport(),
		Incidents: service.Incident(),
	}
	return &report
}
