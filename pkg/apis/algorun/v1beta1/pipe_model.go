/*
 * Algo.Run API 1.0-beta1
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0-beta1
 * Contact: support@algohub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1beta1

type PipeModel struct {

	SourceComponentType string `json:"sourceComponentType,omitempty"`

	SourceName string `json:"sourceName,omitempty"`

	SourceOutputName string `json:"sourceOutputName,omitempty"`

	SourceOutputMessageDataType string `json:"sourceOutputMessageDataType,omitempty"`

	DestComponentType string `json:"destComponentType,omitempty"`

	DestName string `json:"destName,omitempty"`

	DestInputName string `json:"destInputName,omitempty"`
}