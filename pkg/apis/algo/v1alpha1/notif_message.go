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

import (
	"time"
)

type NotifMessage struct {
	Key string `json:"key,omitempty"`

	MessageTimestamp time.Time `json:"messageTimestamp,omitempty"`

	NotifLevel string `json:"notifLevel,omitempty"`

	Title string `json:"title,omitempty"`

	Summary string `json:"summary,omitempty"`

	Url string `json:"url,omitempty"`

	LogMessageType string `json:"logMessageType,omitempty"`

	EndpointStatusMessage *EndpointStatusMessage `json:"endpointStatusMessage,omitempty"`
}
