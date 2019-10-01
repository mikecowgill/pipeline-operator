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

type EndpointConfig struct {

	DeploymentOwnerUserName string `json:"deploymentOwnerUserName,omitempty"`

	DeploymentName string `json:"deploymentName,omitempty"`

	PipelineOwnerUserName string `json:"pipelineOwnerUserName,omitempty"`

	PipelineName string `json:"pipelineName,omitempty"`

	Outputs []EndpointOutputModel `json:"outputs,omitempty"`
}
