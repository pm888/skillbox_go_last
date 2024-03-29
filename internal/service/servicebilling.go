package service

import (
	"mymod/internal/data"
	"mymod/internal/helpers"
)

const (
	CreateCustomerPosition = 1 << iota
	PurchasePosition
	PayoutPosition
	RecurringPosition
	FraudControlPosition
	CheckoutPagePosition
)

func ReadBilling() ([]data.BillingData, error) {
	var BillingDataSlice = make([]data.BillingData, 0)
	sliceFile, err := helpers.ReadFile(data.FileBillingRead)
	ch := helpers.Interpretation(sliceFile)
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
	return BillingDataSlice, err

}
