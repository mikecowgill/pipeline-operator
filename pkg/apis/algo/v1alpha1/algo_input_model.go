/*
 * Algo.Run API 1.0
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0
 * Contact: support@algohub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package v1alpha1

type AlgoInputModel struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	IsRequired bool `json:"isRequired,omitempty"`

	InputDeliveryType string `json:"inputDeliveryType"`

	Parameter string `json:"parameter,omitempty"`

	ParameterDelimiter string `json:"parameterDelimiter,omitempty"`

	EnvironmentVariable string `json:"environmentVariable,omitempty"`

	HttpVerb string `json:"httpVerb,omitempty"`

	HttpPort int32 `json:"httpPort,omitempty"`

	HttpPath string `json:"httpPath,omitempty"`

	HttpHeaders string `json:"httpHeaders,omitempty"`

	ContentTypes []ContentTypeModel `json:"contentTypes"`
}
