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
// HttpGetActionV1 struct for HttpGetActionV1
type HttpGetActionV1 struct {
	Scheme string `json:"scheme,omitempty"`
	Path string `json:"path,omitempty"`
	HttpHeaders []HttpHeaderV1 `json:"httpHeaders,omitempty"`
	Host string `json:"host,omitempty"`
	Port *Int32OrStringV1 `json:"port,omitempty"`
}