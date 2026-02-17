package functions

import (
	"github.com/krendel52/go-astm/v3/models/astmmodels"
	"strings"
)

func LoadConfiguration(configuration ...astmmodels.Configuration) (config *astmmodels.Configuration, err error) {
	if len(configuration) > 0 {
		config = &configuration[0]
	} else {
		config = &astmmodels.DefaultConfiguration
	}
	if config.Delimiters.Field == "" ||
		config.Delimiters.Repeat == "" ||
		config.Delimiters.Component == "" ||
		config.Delimiters.Escape == "" {
		config.Delimiters = astmmodels.DefaultDelimiters
	}
	config.TimeLocation, err = config.TimeZone.GetLocation()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func ExtractSignature(lines []string) string {
	// Extract the first characters from each line
	firstChars := ""
	for _, line := range lines {
		if len(line) > 0 {
			firstChars += string(line[0])
		}
	}
	// Remove M and C characters
	var signature strings.Builder
	for _, r := range firstChars {
		if r != 'M' && r != 'C' {
			signature.WriteRune(r)
		}
	}
	// Return final signature
	return signature.String()
}
