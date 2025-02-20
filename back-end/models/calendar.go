package models

type Calendar struct {
	ServiceID int `gorm:"primaryKey" json:"serviceId"`
	Monday    int    `gorm:"not null" json:"monday"`
	Tuesday   int    `gorm:"not null" json:"tuesday"`
	Wednesday int    `gorm:"not null" json:"wednesday"`
	Thursday  int    `gorm:"not null" json:"thursday"`
	Friday    int    `gorm:"not null" json:"friday"`
	Saturday  int    `gorm:"not null" json:"saturday"`
	Sunday    int    `gorm:"not null" json:"sunday"`
	StartDate string `gorm:"not null" json:"startDate"`
	EndDate   string `gorm:"not null" json:"endDate"`
}
