package utils

import (
	"database/sql/driver"
	"fmt"
)

type CompanyType string

const (
	Corporations CompanyType = "Corporations"
	NonProfit    CompanyType = "NonProfit"
	Sole         CompanyType = "Sole Proprietorship"
)

func (c CompanyType) Value() (driver.Value, error) {
	return string(c), nil
}

// Scan the sql.Scanner interface to convert the Enum value from the db
func (c *CompanyType) Scan(value interface{}) error {
	str, ok := value.(string)
	if !ok {
		return fmt.Errorf("failed to scan CompanyType: expected string, got %T", value)
	}
	*c = CompanyType(str)
	return nil
}
