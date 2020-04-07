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
import (
	"time"
)
// NotifMessage struct for NotifMessage
type NotifMessage struct {
	Key string `json:"key,omitempty"`
	MessageTimestamp time.Time `json:"messageTimestamp,omitempty"`
	Level *LogLevels `json:"level,omitempty"`
	Title string `json:"title,omitempty"`
	Summary string `json:"summary,omitempty"`
	Url string `json:"url,omitempty"`
	Type *NotifTypes `json:"type,omitempty"`
	DeploymentStatusMessage *DeploymentStatusMessage `json:"deploymentStatusMessage,omitempty"`
}
