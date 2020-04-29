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
// InputDeliveryTypes the model 'InputDeliveryTypes'
type InputDeliveryTypes string

// List of InputDeliveryTypes
const (
	INPUTDELIVERYTYPES_STD_IN InputDeliveryTypes = "StdIn"
	INPUTDELIVERYTYPES_HTTP InputDeliveryTypes = "Http"
	INPUTDELIVERYTYPES_HTTPS InputDeliveryTypes = "Https"
	INPUTDELIVERYTYPES_GRPC InputDeliveryTypes = "Grpc"
	INPUTDELIVERYTYPES_FILE_PARAMETER InputDeliveryTypes = "FileParameter"
	INPUTDELIVERYTYPES_FILE_ARRAY_PARAMETER InputDeliveryTypes = "FileArrayParameter"
	INPUTDELIVERYTYPES_KAFKA_TOPIC InputDeliveryTypes = "KafkaTopic"
)
