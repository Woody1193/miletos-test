package data

import (
	"database/sql/driver"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/xefino/protobuf-gen-go/utils"
	"gopkg.in/yaml.v3"
)

// ClientTypeAlternates contains alternative values for the DynamoDB.ClientType enum
var ClientTypeAlternates = map[string]ClientType{
	"":      ClientType_Invalid,
	"web":   ClientType_Web,
	"cli":   ClientType_CLI,
	"other": ClientType_Other,
}

// FrequencyAlternates contains alternative values for the OHLC.v1.Frequency enum
var FrequencyAlternates = map[string]Frequency{
	"":        Frequency_InvalidFrequency,
	"s":       Frequency_Second,
	"S":       Frequency_Second,
	"second":  Frequency_Second,
	"m":       Frequency_Minute,
	"M":       Frequency_Minute,
	"minute":  Frequency_Minute,
	"h":       Frequency_Hour,
	"H":       Frequency_Hour,
	"hour":    Frequency_Hour,
	"d":       Frequency_Day,
	"D":       Frequency_Day,
	"day":     Frequency_Day,
	"w":       Frequency_Week,
	"W":       Frequency_Week,
	"week":    Frequency_Week,
	"mn":      Frequency_Month,
	"MN":      Frequency_Month,
	"month":   Frequency_Month,
	"q":       Frequency_Quarter,
	"Q":       Frequency_Quarter,
	"quarter": Frequency_Quarter,
	"y":       Frequency_Year,
	"Y":       Frequency_Year,
	"year":    Frequency_Year,
}

// FrequencyMapping contains alternate names for the data.Frequency enum
var FrequencyMapping = map[Frequency]string{
	Frequency_InvalidFrequency: "",
	Frequency_Second:           "s",
	Frequency_Minute:           "m",
	Frequency_Hour:             "h",
	Frequency_Day:              "d",
	Frequency_Week:             "w",
	Frequency_Month:            "mn",
	Frequency_Quarter:          "q",
	Frequency_Year:             "y",
}

// UnmarshalJSON converts JSON data into a Data.ClientType
func (enum *ClientType) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalValue(data, ClientType_value, ClientTypeAlternates, enum)
}

// MarshalJSON converts a data.Frequency value to a JSON value
func (enum Frequency) MarshalJSON() ([]byte, error) {
	return []byte(utils.MarshalString(enum, Frequency_name, FrequencyMapping, true)), nil
}

// MarshalCSV converts a data.Frequency value to CSV cell value
func (enum Frequency) MarshalCSV() (string, error) {
	return utils.MarshalString(enum, Frequency_name, FrequencyMapping, false), nil
}

// MarshalYAML converts a data.Frequency value to a YAML node value
func (enum Frequency) MarshalYAML() (interface{}, error) {
	return utils.MarshalString(enum, Frequency_name, FrequencyMapping, false), nil
}

// MarshalDynamoDBAttributeValue converts a data.Frequency value to a DynamoDB AttributeValue
func (enum Frequency) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberS{Value: utils.MarshalString(enum, Frequency_name, FrequencyMapping, false)}, nil
}

// Value converts a data.Frequency value to an SQL driver value
func (enum Frequency) Value() (driver.Value, error) {
	return utils.MarshalString(enum, Frequency_name, FrequencyMapping, false), nil
}

// UnmarshalJSON attempts to convert a JSON value to a new data.Frequency value
func (enum *Frequency) UnmarshalJSON(data []byte) error {
	return utils.UnmarshalValue(data, Frequency_value, FrequencyAlternates, enum)
}

// UnmarshalCSV attempts to convert a CSV cell value to a new data.Frequency value
func (enum *Frequency) UnmarshalCSV(raw string) error {
	return utils.UnmarshalString(raw, Frequency_value, FrequencyAlternates, enum)
}

// UnmarshalYAML attempts to convert a YAML node to a new data.Frequency value
func (enum *Frequency) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return fmt.Errorf("YAML node had an invalid kind (expected scalar value)")
	} else {
		return utils.UnmarshalString(value.Value, Frequency_value, FrequencyAlternates, enum)
	}
}

// UnmarshalParam attempts to converts a form or query parameter into a new data.Frequency value
func (enum *Frequency) UnmarshalParam(param string) error {
	return utils.UnmarshalString(param, Frequency_value, FrequencyAlternates, enum)
}

// UnmarshalDynamoDBAttributeValue attempts to convert a DynamoDB AttributeVAlue to a data.Frequency
// value. This function can handle []bytes, numerics, or strings. If the AttributeValue is NULL then
// the Frequency value will not be modified.
func (enum *Frequency) UnmarshalDynamoDBAttributeValue(value types.AttributeValue) error {
	switch casted := value.(type) {
	case *types.AttributeValueMemberB:
		return utils.UnmarshalValue(casted.Value, Frequency_value, FrequencyAlternates, enum)
	case *types.AttributeValueMemberN:
		return utils.UnmarshalString(casted.Value, Frequency_value, FrequencyAlternates, enum)
	case *types.AttributeValueMemberNULL:
		return nil
	case *types.AttributeValueMemberS:
		return utils.UnmarshalString(casted.Value, Frequency_value, FrequencyAlternates, enum)
	default:
		return fmt.Errorf("Attribute value of %T could not be converted to a data.Frequency", value)
	}
}

// Scan attempts to convert an SQL driver value to a new data.Frequency value
func (enum *Frequency) Scan(value interface{}) error {
	return utils.ScanValue(value, Frequency_value, FrequencyAlternates, enum)
}
