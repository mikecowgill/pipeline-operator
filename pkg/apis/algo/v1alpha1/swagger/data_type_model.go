/*
 * Algo.Run API 1.0
 *
 * API for the Algo.Run Engine
 *
 * API version: 1.0
 * Contact: support@algohub.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DataTypeModel struct {

	Name string `json:"name"`

	Regex string `json:"regex,omitempty"`

	Precision int32 `json:"precision,omitempty"`

	Scale int32 `json:"scale,omitempty"`

	Mask string `json:"mask,omitempty"`
}
