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
// HookConfig struct for HookConfig
type HookConfig struct {
	DeploymentOwner string `json:"deploymentOwner,omitempty"`
	DeploymentName string `json:"deploymentName,omitempty"`
	PipelineOwner string `json:"pipelineOwner,omitempty"`
	PipelineName string `json:"pipelineName,omitempty"`
	Image *ContainerImageModel `json:"image,omitempty"`
	Resource *ResourceModel `json:"resource,omitempty"`
	WebHooks []WebHookModel `json:"webHooks,omitempty"`
	Pipes []PipeModel `json:"pipes,omitempty"`
	Topics []TopicConfigModel `json:"topics,omitempty"`
	LivenessProbe *ProbeModel `json:"livenessProbe,omitempty"`
	ReadinessProbe *ProbeModel `json:"readinessProbe,omitempty"`
}
