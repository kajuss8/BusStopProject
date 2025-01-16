package models

type FarePayment uint8

const (
    FareOnBoard         FarePayment = 0
    FareBeforeBoarding  FarePayment = 1
)

type TransferPermitted uint8

const (
    NoTransfers      TransferPermitted = 0 
    OneTransfer      TransferPermitted = 1 
    TwoTransfers     TransferPermitted = 2 
    UnlimitedTransfers TransferPermitted = 255 
)

type FareAttribute struct {
	Id 				uint32 			`json:"id"`
	FareId 			uint32 			`json:"fareId"`
	Price 			float32 		`json:"price"`
	CurrencyType 	string 			`json:"currencyType"`
	PaymentMethod 	FarePayment 	`json:"paymentMethod"`
	Transfers 		TransferPermitted `json:"transfers"`
	TransferDuration uint16 		`json:"transferDuration"`
}