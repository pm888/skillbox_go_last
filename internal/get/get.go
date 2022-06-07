package get

import (
	"mymod/internal/data"
	"mymod/internal/service"
)

func GetResultData() (*data.ResultSetT, error) {
	sms, err := service.FileSMS()
	if err != nil {
		return nil, err
	}
	mms, err := service.GetMMS()
	if err != nil {
		return nil, err
	}
	voiceCall, err := service.FileVoice()
	if err != nil {
		return nil, err
	}
	email, err := service.FileEmail()
	if err != nil {
		return nil, err
	}
	billing, err := service.ReadBilling()
	if err != nil {
		return nil, err
	}
	support, err := service.GetSupport()
	if err != nil {
		return nil, err
	}
	incidents, err := service.Incident()
	if err != nil {
		return nil, err
	}
	report := data.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: voiceCall,
		Email:     email,
		Billing:   billing,
		Support:   support,
		Incidents: incidents,
	}
	return &report, err
}
