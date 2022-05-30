package get

import (
	"mymod/internal/data"
	"mymod/internal/servicebilling"
	"mymod/internal/serviceincident"
	"mymod/internal/servicemail"
	"mymod/internal/servicemms"
	"mymod/internal/servicesms"
	"mymod/internal/servicesupport"
	servecevoicecall "mymod/internal/servicevoicecall"
)

func GetResultData() *data.ResultSetT {
	report := data.ResultSetT{
		SMS:       servicesms.FileSMS(),
		MMS:       servicemms.GetMMS(),
		VoiceCall: servecevoicecall.FileVoice(),
		Email:     servicemail.FileEmail(),
		Billing:   servicebilling.ReadBilling(),
		Support:   servicesupport.GetSupport(),
		Incidents: serviceincident.Incident(),
	}
	return &report
}
