package get

import (
	"mymod/internal/data"
	"mymod/internal/service"
)

// Result возвращает результаты различных служб.
func Result() (*data.ResultSetT, error) {
	// Получение данных из службы SMS
	sms, err := service.FileSMS()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы MMS
	mms, err := service.GetMMS()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы VoiceCall
	voiceCall, err := service.FileVoice()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы Email
	email, err := service.FileEmail()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы Billing
	billing, err := service.ReadBilling()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы Support
	support, err := service.GetSupport()
	if err != nil {
		return nil, err
	}

	// Получение данных из службы Incidents
	incidents, err := service.Incident()
	if err != nil {
		return nil, err
	}

	// Возвращение результата в виде структуры
	return &data.ResultSetT{
		SMS:       sms,
		MMS:       mms,
		VoiceCall: voiceCall,
		Email:     email,
		Billing:   billing,
		Support:   support,
		Incidents: incidents,
	}, nil
}
