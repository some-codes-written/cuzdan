package models

import (
	"immortality/pkg/common"
	"immortality/pkg/domain/persons"
	"time"
)

// hesap turu gelir mi gider mi ?
type AccountingType struct {
	common.EntityModel

	Name        string `json:"name" example:"name"`
	Description string `json:"description" example:"description"`
}

// gelir / gider turleri
type ExpenseType struct {
	common.EntityModel

	ParentExpenseTypeID uint         `json:"parentExpenseTypeId" example:"1"` // ExpenseType
	ParentExpenseType   *ExpenseType `json:"parentExpenseType" gorm:"foreignKey:ParentExpenseTypeID"`
	Name                string       `json:"name" example:"name"`
	Description         string       `json:"description" example:"description"`
}

// kisilere bagli olarak yapilan islemler
type Accounting struct {
	common.EntityModel

	ExpenseTypeID    uint            `json:"expenseTypeId" example:"1"` // ExpenseType
	ExpenseType      *ExpenseType    `json:"expenseType" gorm:"foreignKey:ExpenseTypeID"`
	AccountingTypeID uint            `json:"accountingTypeId" example:"1"` // AccountingType
	AccountingType   *AccountingType `json:"accountingType" gorm:"foreignKey:AccountingTypeId"`
	PersonId         uint            `json:"personId" example:"1"`              // Person
	Person           *persons.Person `json:"person" gorm:"foreignKey:PersonId"` // Person
	Description      string          `json:"description" example:"description"`
	Amount           float64         `json:"amount" example:"100.00"`
	CurrentTotal     float64         `json:"currentTotal" example:"100.00"`
	ProcessDate      time.Time       `json:"processDate" example:"2021-01-01T00:00:00Z"`
}

// gider yada gelir turu.
type CashflowType struct {
	common.EntityModel

	Name        string `json:"name" example:"name"`
	Description string `json:"description" example:"description"`
}

// para aktarimlarinin yapildigi yer
type Cashflow struct {
	common.EntityModel

	AccountingID   uint          `json:"accountingId" example:"1"`
	Accounting     *Accounting   `json:"accounting" gorm:"foreignKey:AccountingId"` // Accounting
	CashflowTypeId uint          `json:"cashflowTypeId" example:"1"`                // CashflowType
	CashflowType   *CashflowType `json:"cashflowType" gorm:"foreignKey:CashflowTypeId"`
	Amount         float64       `json:"amount" example:"100.00"`
	ProcessDate    time.Time     `json:"processDate" example:"2021-01-01T00:00:00Z"`
}

// bir hesaptan birden fazla hesaba gonderilmis para bulunabilir.
type CashflowAccountingDetail struct {
	common.EntityModel

	CashflowID   uint        `json:"cashflowId" example:"1"`
	Cashflow     *Cashflow   `json:"cashflow" gorm:"foreignKey:CashflowId"` // Cashflow
	AccountingID uint        `json:"accountingId" example:"1"`
	Accounting   *Accounting `json:"accounting" gorm:"foreignKey:AccountingId"` // Accounting
	Amount       float64     `json:"amount" example:"100.00"`
	ProcessDate  time.Time   `json:"processDate" example:"2021-01-01T00:00:00Z"`
}
