/*
 * Algo.Run API 1.0-beta1
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0-beta1
 * Contact: support@algohub.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package v1beta1
// DataTypes the model 'DataTypes'
type DataTypes string

// List of DataTypes
const (
	DATATYPES_STRING DataTypes = "String"
	DATATYPES_SELECT DataTypes = "Select"
	DATATYPES_SWITCH DataTypes = "Switch"
	DATATYPES_INTEGER DataTypes = "Integer"
	DATATYPES_DECIMAL DataTypes = "Decimal"
	DATATYPES_FLOAT DataTypes = "Float"
	DATATYPES_CURRENCY DataTypes = "Currency"
	DATATYPES_IPV4 DataTypes = "Ipv4"
	DATATYPES_DATE DataTypes = "Date"
	DATATYPES_DATETIME DataTypes = "Datetime"
	DATATYPES_BOOLEAN DataTypes = "Boolean"
	DATATYPES_GUID DataTypes = "Guid"
)