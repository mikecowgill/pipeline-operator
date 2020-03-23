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
// DataConnectorOptionModel struct for DataConnectorOptionModel
type DataConnectorOptionModel struct {
	SortOrder *int32 `json:"sortOrder,omitempty"`
	Required bool `json:"required,omitempty"`
	Name string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Value string `json:"value,omitempty"`
	DataType *DataTypeModel `json:"dataType,omitempty"`
}
