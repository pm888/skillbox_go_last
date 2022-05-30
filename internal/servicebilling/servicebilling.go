package servicebilling

import (
	"mymod/internal/data"
	"mymod/internal/method"
)

const (
	CreateCustomerPosition = 1 << iota
	PurchasePosition
	PayoutPosition
	RecurringPosition
	FraudControlPosition
	CheckoutPagePosition
)

var BillingDataSlice = make([]data.BillingData, 0)

func ReadBilling() []data.BillingData {
	sliceFile := method.ReadFile(data.FileBillingRead)
	ch := method.Interpretation(sliceFile)
	mask := int(ch)
	person := data.BillingData{
		CreateCustomer: CreateCustomerPosition&mask > 0,
		Purchase:       PurchasePosition&mask > 0,
		Payout:         PayoutPosition&mask > 0,
		Recurring:      RecurringPosition&mask > 0,
		FraudControl:   FraudControlPosition&mask > 0,
		CheckoutPage:   CheckoutPagePosition&mask > 0,
	}
	BillingDataSlice = append(BillingDataSlice, person)
	return BillingDataSlice

}
