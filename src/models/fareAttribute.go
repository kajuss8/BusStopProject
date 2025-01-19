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
	Id 				int 			`json:"id"`
	FareId 			int 			`json:"fareId"`
	Price 			float64 		`json:"price"`
	CurrencyType 	string 			`json:"currencyType"`
	PaymentMethod 	FarePayment 	`json:"paymentMethod"`
	Transfers 		TransferPermitted `json:"transfers"`
	TransferDuration int 			`json:"transferDuration"`
}