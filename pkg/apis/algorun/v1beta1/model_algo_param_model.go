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
// AlgoParamModel struct for AlgoParamModel
type AlgoParamModel struct {
	SortOrder *int32 `json:"sortOrder,omitempty"`
	Required bool `json:"required,omitempty"`
	Name string `json:"name"`
	Value string `json:"value,omitempty"`
	Description string `json:"description,omitempty"`
	DataType *DataTypeModel `json:"dataType,omitempty"`
	Options []DataTypeOptionModel `json:"options,omitempty"`
}
