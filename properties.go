package properties

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/gershwinlabs/pathlib"
)

const (
	PropertiesExtension = ".properties"
)

type Properties map[string]string

// NewProperties creates a new, empty Properties instance.
func NewProperties() Properties {
	return make(map[string]string)
}

// BytesToProperties takes bytes (usually from a file) and returns a Properties instance.
func BytesToProperties(input []byte) (Properties, error) {
	inputStr := bytes.NewBuffer(input).String()
	lines := make([]string, 0)
	properties := make(Properties)
	errorList := make([]string, 0)

	for _, line := range strings.Split(inputStr, "\n") {
		line = strings.TrimSpace(line)

		if len(line) > 0 {
			lines = append(lines, line)
		}
	}

	for _, line := range lines {
		parts := strings.Split(line, "=")

		if len(parts) == 2 {
			name := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			properties[name] = value
		} else {
			errorList = append(errorList, fmt.Sprintf("Invalid property: %s", line))
		}
	}

	if len(errorList) > 0 {
		return properties, errors.New(strings.Join(errorList, "; "))
	}

	return properties, nil
}

// PropertiesFromFile takes a `pathlib.Path` and returns a Properties instance.
func PropertiesFromFile(path pathlib.Path) (Properties, error) {
	propertiesBytes, err := path.ReadBytes()

	if err != nil {
		return nil, err
	}

	return BytesToProperties(propertiesBytes)
}

// hasReservedCharacters returns true if the string contains any characters that are reserved.
func hasReservedCharacters(s string) bool {
	return strings.ContainsAny(s, "\n=")
}

// PropertiesToFile writes properties to a file.
func PropertiesToFile(properties Properties, path pathlib.Path) error {
	propertiesBytes, err := PropertiesToBytes(properties)

	if err != nil {
		return err
	}

	return path.WriteBytes(propertiesBytes)
}

// PropertiesToBytes returns the byte representation of Properties.
func PropertiesToBytes(properties Properties) ([]byte, error) {
	parts := make([]string, 0, len(properties))

	for key, value := range properties {
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		if hasReservedCharacters(key) {
			return nil, fmt.Errorf("Invalid property key \"%s\"", key)
		}

		if hasReservedCharacters(value) {
			return nil, fmt.Errorf("Invalid property value \"%s\"", value)
		}

		parts = append(parts, fmt.Sprintf("%s = %s", key, value))
	}

	propertiesStr := strings.Join(parts, "\n")
	var buf bytes.Buffer
	buf.WriteString(propertiesStr)
	return buf.Bytes(), nil
}

// ValidPropertiesFile take a `pathlib.Path` and returns true if it's a valid properties file.
func ValidPropertiesFile(path pathlib.Path) bool {
	_, err := PropertiesFromFile(path)

	if err == nil {
		return true
	}

	return false
}
