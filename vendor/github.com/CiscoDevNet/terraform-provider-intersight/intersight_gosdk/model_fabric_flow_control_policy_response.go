/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-04-12T05:47:20Z.
 *
 * API version: 1.0.9-4240
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"fmt"
)

// FabricFlowControlPolicyResponse - The response body of a HTTP GET request for the 'fabric.FlowControlPolicy' resource. The value may be one of the following types. 1. When 'tag' is specified in the URL query, the response schema     is a summary of the tag usage. 1. When '$apply' is specified in the URL query, the response schema     is dynamically-generated schema based on the $apply value. 1. When '$count' is specified in the URL query, the response is     a simple object providing the count of the resources. 1. In all other cases, the response is a list of 'fabric.FlowControlPolicy' resources.
type FabricFlowControlPolicyResponse struct {
	FabricFlowControlPolicyList *FabricFlowControlPolicyList
	MoAggregateTransform        *MoAggregateTransform
	MoDocumentCount             *MoDocumentCount
	MoTagSummary                *MoTagSummary
}

// FabricFlowControlPolicyListAsFabricFlowControlPolicyResponse is a convenience function that returns FabricFlowControlPolicyList wrapped in FabricFlowControlPolicyResponse
func FabricFlowControlPolicyListAsFabricFlowControlPolicyResponse(v *FabricFlowControlPolicyList) FabricFlowControlPolicyResponse {
	return FabricFlowControlPolicyResponse{FabricFlowControlPolicyList: v}
}

// MoAggregateTransformAsFabricFlowControlPolicyResponse is a convenience function that returns MoAggregateTransform wrapped in FabricFlowControlPolicyResponse
func MoAggregateTransformAsFabricFlowControlPolicyResponse(v *MoAggregateTransform) FabricFlowControlPolicyResponse {
	return FabricFlowControlPolicyResponse{MoAggregateTransform: v}
}

// MoDocumentCountAsFabricFlowControlPolicyResponse is a convenience function that returns MoDocumentCount wrapped in FabricFlowControlPolicyResponse
func MoDocumentCountAsFabricFlowControlPolicyResponse(v *MoDocumentCount) FabricFlowControlPolicyResponse {
	return FabricFlowControlPolicyResponse{MoDocumentCount: v}
}

// MoTagSummaryAsFabricFlowControlPolicyResponse is a convenience function that returns MoTagSummary wrapped in FabricFlowControlPolicyResponse
func MoTagSummaryAsFabricFlowControlPolicyResponse(v *MoTagSummary) FabricFlowControlPolicyResponse {
	return FabricFlowControlPolicyResponse{MoTagSummary: v}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *FabricFlowControlPolicyResponse) UnmarshalJSON(data []byte) error {
	var err error
	// use discriminator value to speed up the lookup
	var jsonDict map[string]interface{}
	err = json.Unmarshal(data, &jsonDict)
	if err != nil {
		return fmt.Errorf("Failed to unmarshal JSON into map for the discrimintor lookup.")
	}

	// check if the discriminator value is 'fabric.FlowControlPolicy.List'
	if jsonDict["ObjectType"] == "fabric.FlowControlPolicy.List" {
		// try to unmarshal JSON data into FabricFlowControlPolicyList
		err = json.Unmarshal(data, &dst.FabricFlowControlPolicyList)
		if err == nil {
			return nil // data stored in dst.FabricFlowControlPolicyList, return on the first match
		} else {
			dst.FabricFlowControlPolicyList = nil
			return fmt.Errorf("Failed to unmarshal FabricFlowControlPolicyResponse as FabricFlowControlPolicyList: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.AggregateTransform'
	if jsonDict["ObjectType"] == "mo.AggregateTransform" {
		// try to unmarshal JSON data into MoAggregateTransform
		err = json.Unmarshal(data, &dst.MoAggregateTransform)
		if err == nil {
			return nil // data stored in dst.MoAggregateTransform, return on the first match
		} else {
			dst.MoAggregateTransform = nil
			return fmt.Errorf("Failed to unmarshal FabricFlowControlPolicyResponse as MoAggregateTransform: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.DocumentCount'
	if jsonDict["ObjectType"] == "mo.DocumentCount" {
		// try to unmarshal JSON data into MoDocumentCount
		err = json.Unmarshal(data, &dst.MoDocumentCount)
		if err == nil {
			return nil // data stored in dst.MoDocumentCount, return on the first match
		} else {
			dst.MoDocumentCount = nil
			return fmt.Errorf("Failed to unmarshal FabricFlowControlPolicyResponse as MoDocumentCount: %s", err.Error())
		}
	}

	// check if the discriminator value is 'mo.TagSummary'
	if jsonDict["ObjectType"] == "mo.TagSummary" {
		// try to unmarshal JSON data into MoTagSummary
		err = json.Unmarshal(data, &dst.MoTagSummary)
		if err == nil {
			return nil // data stored in dst.MoTagSummary, return on the first match
		} else {
			dst.MoTagSummary = nil
			return fmt.Errorf("Failed to unmarshal FabricFlowControlPolicyResponse as MoTagSummary: %s", err.Error())
		}
	}

	return nil
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FabricFlowControlPolicyResponse) MarshalJSON() ([]byte, error) {
	if src.FabricFlowControlPolicyList != nil {
		return json.Marshal(&src.FabricFlowControlPolicyList)
	}

	if src.MoAggregateTransform != nil {
		return json.Marshal(&src.MoAggregateTransform)
	}

	if src.MoDocumentCount != nil {
		return json.Marshal(&src.MoDocumentCount)
	}

	if src.MoTagSummary != nil {
		return json.Marshal(&src.MoTagSummary)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FabricFlowControlPolicyResponse) GetActualInstance() interface{} {
	if obj.FabricFlowControlPolicyList != nil {
		return obj.FabricFlowControlPolicyList
	}

	if obj.MoAggregateTransform != nil {
		return obj.MoAggregateTransform
	}

	if obj.MoDocumentCount != nil {
		return obj.MoDocumentCount
	}

	if obj.MoTagSummary != nil {
		return obj.MoTagSummary
	}

	// all schemas are nil
	return nil
}

type NullableFabricFlowControlPolicyResponse struct {
	value *FabricFlowControlPolicyResponse
	isSet bool
}

func (v NullableFabricFlowControlPolicyResponse) Get() *FabricFlowControlPolicyResponse {
	return v.value
}

func (v *NullableFabricFlowControlPolicyResponse) Set(val *FabricFlowControlPolicyResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFabricFlowControlPolicyResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFabricFlowControlPolicyResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFabricFlowControlPolicyResponse(val *FabricFlowControlPolicyResponse) *NullableFabricFlowControlPolicyResponse {
	return &NullableFabricFlowControlPolicyResponse{value: val, isSet: true}
}

func (v NullableFabricFlowControlPolicyResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFabricFlowControlPolicyResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
