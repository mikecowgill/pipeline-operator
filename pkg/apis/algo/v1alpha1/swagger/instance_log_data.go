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

import (
	"time"
)

type InstanceLogData struct {

	OrchEventType string `json:"orchEventType,omitempty"`

	EndpointOwnerUserName string `json:"endpointOwnerUserName,omitempty"`

	EndpointName string `json:"endpointName,omitempty"`

	AlgoOwnerUserName string `json:"algoOwnerUserName,omitempty"`

	AlgoName string `json:"algoName,omitempty"`

	AlgoVersionTag string `json:"algoVersionTag,omitempty"`

	AlgoIndex int32 `json:"algoIndex,omitempty"`

	Name string `json:"name,omitempty"`

	Status string `json:"status,omitempty"`

	Restarts int32 `json:"restarts,omitempty"`

	CreatedTimestamp time.Time `json:"createdTimestamp,omitempty"`

	Ip string `json:"ip,omitempty"`

	Node string `json:"node,omitempty"`
}
