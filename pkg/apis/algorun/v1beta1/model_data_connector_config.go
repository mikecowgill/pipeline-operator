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
// DataConnectorConfig struct for DataConnectorConfig
type DataConnectorConfig struct {
	DeploymentName string `json:"deploymentName,omitempty"`
	DataConnectorType *DataConnectorTypes `json:"dataConnectorType,omitempty"`
	Name string `json:"name,omitempty"`
	VersionTag string `json:"versionTag,omitempty"`
	Index int32 `json:"index,omitempty"`
	ImageRepository string `json:"imageRepository,omitempty"`
	ImageTag string `json:"imageTag,omitempty"`
	ConnectorClass string `json:"connectorClass,omitempty"`
	TasksMax int32 `json:"tasksMax,omitempty"`
	TopicConfigs []TopicConfigModel `json:"topicConfigs,omitempty"`
	Options []DataConnectorOptionModel `json:"options,omitempty"`
}
