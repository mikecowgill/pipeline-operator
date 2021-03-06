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
// NotifTypes the model 'NotifTypes'
type NotifTypes string

// List of NotifTypes
const (
	NOTIFTYPES_OPERATOR NotifTypes = "Operator"
	NOTIFTYPES_PIPELINE_DEPLOYMENT_STATUS NotifTypes = "PipelineDeploymentStatus"
	NOTIFTYPES_PIPELINE_DEPLOYMENT_DELETED NotifTypes = "PipelineDeploymentDeleted"
	NOTIFTYPES_PIPELINE_DEPLOYMENT NotifTypes = "PipelineDeployment"
	NOTIFTYPES_PIPELINE_DEPLOYMENT_POD NotifTypes = "PipelineDeploymentPod"
)
