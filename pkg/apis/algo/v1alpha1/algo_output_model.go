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

type AlgoOutputModel struct {

	Name string `json:"name"`

	Description string `json:"description,omitempty"`

	OutputDeliveryType string `json:"outputDeliveryType"`

	Parameter string `json:"parameter,omitempty"`

	ContentType *ContentTypeModel `json:"contentType"`
}
