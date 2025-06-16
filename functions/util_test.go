package functions

import (
	"github.com/blutspende/bloodlab-common/timezone"
	"github.com/blutspende/go-astm/v3/enums/lineseparator"
	"github.com/blutspende/go-astm/v3/models/astmmodels"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoadConfigurationEmpty(t *testing.T) {
	// Arrange
	// Act
	loadedConfig, err := LoadConfiguration()
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, loadedConfig)
}
func TestLoadConfigurationNoDelimiters(t *testing.T) {
	// Arrange
	inputConfig := astmmodels.Configuration{
		Delimiters: astmmodels.Delimiters{},
	}
	// Act
	loadedConfig, err := LoadConfiguration(inputConfig)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, loadedConfig)
	assert.Equal(t, "|", loadedConfig.Delimiters.Field)
}
func TestLoadConfigurationValueKept(t *testing.T) {
	// Arrange
	inputConfig := astmmodels.Configuration{
		LineSeparator: lineseparator.LF,
	}
	// Act
	loadedConfig, err := LoadConfiguration(inputConfig)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, loadedConfig)
	assert.Equal(t, "\n", loadedConfig.LineSeparator)
}
func TestLoadConfigurationLocationLoaded(t *testing.T) {
	// Arrange
	inputConfig := astmmodels.Configuration{
		TimeZone: timezone.EuropeBerlin,
	}
	location, _ := timezone.EuropeBerlin.GetLocation()
	// Act
	loadedConfig, err := LoadConfiguration(inputConfig)
	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, loadedConfig)
	assert.Equal(t, location, loadedConfig.TimeLocation)
}

func TestExtractSignature(t *testing.T) {
	// Arrange
	lines := []string{
		"H|something",
		"P|1|something",
		"M|1|something",
		"C|1|something",
		"C|2|something",
		"O|1|something",
		"C|1|something",
		"R|1|something",
		"M|1|something",
		"C|1|something",
		"R|2|something",
		"C|1|something",
		"L|1|N",
	}
	// Act
	signature := ExtractSignature(lines)
	// Assert
	assert.Equal(t, "HPORRL", signature)
}
