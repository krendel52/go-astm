package astm

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/messagetype"
	"github.com/krendel52/go-astm/v3/functions"
	"github.com/krendel52/go-astm/v3/models/astmmodels"
	"regexp"
)

func IdentifyMessage(messageData []byte, configuration ...astmmodels.Configuration) (messageType messagetype.MessageType, err error) {
	// Load configuration
	config, err := functions.LoadConfiguration(configuration...)
	if err != nil {
		return "", err
	}
	// Convert encoding to UTF8
	utf8Data, err := encoding.ConvertFromEncodingToUtf8(messageData, config.Encoding)
	if err != nil {
		return "", err
	}
	// Split the message data into lines
	lines, err := functions.SliceLines(utf8Data, config)
	if err != nil {
		return "", err
	}
	// Extract signature
	signature := functions.ExtractSignature(lines)

	// Check the signature against the regexes and return the message type
	switch {
	case regexOrderAndResult.MatchString(signature):
		return messagetype.Result, nil
	case regexQuery.MatchString(signature):
		return messagetype.Query, nil
	case regexManyOrderAndResult.MatchString(signature):
		return messagetype.Result, nil
	case regexOrder.MatchString(signature):
		return messagetype.Order, nil
	}
	// If no match was found return unknown
	return messagetype.Unidentified, nil
}

// Regular expressions to identify message types
var (
	regexQuery              = regexp.MustCompile("^(HQ+)+L?$")
	regexOrder              = regexp.MustCompile("^(H(PO+)+)+L?$")
	regexOrderAndResult     = regexp.MustCompile("^H(P(OR+)+)+L?$")
	regexManyOrderAndResult = regexp.MustCompile("^(H(P(OR+)+)+L?)+$")
)
