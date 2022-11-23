package parser

import (
	"fmt"
)

type ArabicParser struct{}

func (p *ArabicParser) Parse(inputStr string, operation *string, num *float64) error {
	if _, err := fmt.Sscanf(inputStr, "%s %v", operation, num); err != nil {
		return err
	}

	return nil
}
