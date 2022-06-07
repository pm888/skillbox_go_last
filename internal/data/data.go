package data

const (
	FileNameAlpha2  = "data/alpha-2.txt"
	FileVoiceRead   = "simulator/voice.data"
	FileSmsRead     = "simulator/sms.data"
	FileEmailRead   = "simulator/email.data"
	FileBillingRead = "simulator/billing.data"
	UrlMMS          = "http://127.0.0.1:8383/mms"
	UrlSuport       = "http://127.0.0.1:8383/support"
	UrlAccendent    = "http://127.0.0.1:8383/accendent"
	UrlServer       = "127.0.0.1:8080"
)

var Report *ResultSetT

var ValidProviders = map[string]struct{}{
	"Topolo": {},
	"Rond":   {},
	"Kildy":  {},
}

var ValidProvidersVoice = map[string]struct{}{
	"E-Voice":          {},
	"TransparentCalls": {},
	"JustPhone":        {},
}

var ValidProvidersEmail = map[string]struct{}{
	"Gmail":      {},
	"Yahoo":      {},
	"Hotmail":    {},
	"MSN":        {},
	"Orange":     {},
	"Comcast":    {},
	"AOL":        {},
	"Live":       {},
	"RediffMail": {},
	"GMX":        {},
	"Protonmail": {},
	"Yandex":     {},
	"Mail.ru":    {},
}

type SMSData struct {
	Country      string `json:"country"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
	Provider     string `json:"provider"`
}

type VoiceCallData struct {
	Country             string  `json:"country"`
	Bandwidth           string  `json:"bandwidth"`
	ResponseTime        string  `json:"response_time"`
	Provider            string  `json:"provider"`
	ConnectionStability float32 `json:"connection_stability"`
	TTFB                int     `json:"ttfb"`
	VoicePurity         int     `json:"voice_purity"`
	MedianOfCallsTime   int     `json:"median_of_calls_time"`
}
type MMSData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	Bandwidth    string `json:"bandwidth"`
	ResponseTime string `json:"response_time"`
}

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}
type SupportData struct {
	Topic         string `json:"topic"`
	ActiveTickets int    `json:"active_tickets"`
}

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"`
}
type BillingData struct {
	CreateCustomer bool `json:"createCustomer"`
	Purchase       bool `json:"purchase"`
	Payout         bool `json:"payout"`
	Recurring      bool `json:"recurring"`
	FraudControl   bool `json:"fraudControl"`
	CheckoutPage   bool `json:"checkoutPage"`
}

type ResultT struct {
	Status bool       `json:"status"`
	Data   ResultSetT `json:"data"`
	Error  string     `json:"error"`
}
type ResultSetT struct {
	SMS       [][]SMSData              `json:"sms"`
	MMS       [][]MMSData              `json:"mms"`
	VoiceCall []VoiceCallData          `json:"voice_call"`
	Email     map[string][][]EmailData `json:"email"`
	Billing   []BillingData            `json:"billing"`
	Support   []int                    `json:"support"`
	Incidents []IncidentData           `json:"incident"`
}
