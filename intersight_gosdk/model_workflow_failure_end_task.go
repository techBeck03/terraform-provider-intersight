/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2021-03-10T06:51:24Z.
 *
 * API version: 1.0.9-3942
 * Contact: intersight@cisco.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// WorkflowFailureEndTask A FailureEndTask denotes the failed completion of a workflow.
type WorkflowFailureEndTask struct {
	WorkflowEndTask
	AdditionalProperties map[string]interface{}
}

type _WorkflowFailureEndTask WorkflowFailureEndTask

// NewWorkflowFailureEndTask instantiates a new WorkflowFailureEndTask object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowFailureEndTask() *WorkflowFailureEndTask {
	this := WorkflowFailureEndTask{}
	return &this
}

// NewWorkflowFailureEndTaskWithDefaults instantiates a new WorkflowFailureEndTask object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowFailureEndTaskWithDefaults() *WorkflowFailureEndTask {
	this := WorkflowFailureEndTask{}
	return &this
}

func (o WorkflowFailureEndTask) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedWorkflowEndTask, errWorkflowEndTask := json.Marshal(o.WorkflowEndTask)
	if errWorkflowEndTask != nil {
		return []byte{}, errWorkflowEndTask
	}
	errWorkflowEndTask = json.Unmarshal([]byte(serializedWorkflowEndTask), &toSerialize)
	if errWorkflowEndTask != nil {
		return []byte{}, errWorkflowEndTask
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowFailureEndTask) UnmarshalJSON(bytes []byte) (err error) {
	type WorkflowFailureEndTaskWithoutEmbeddedStruct struct {
	}

	varWorkflowFailureEndTaskWithoutEmbeddedStruct := WorkflowFailureEndTaskWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varWorkflowFailureEndTaskWithoutEmbeddedStruct)
	if err == nil {
		varWorkflowFailureEndTask := _WorkflowFailureEndTask{}
		*o = WorkflowFailureEndTask(varWorkflowFailureEndTask)
	} else {
		return err
	}

	varWorkflowFailureEndTask := _WorkflowFailureEndTask{}

	err = json.Unmarshal(bytes, &varWorkflowFailureEndTask)
	if err == nil {
		o.WorkflowEndTask = varWorkflowFailureEndTask.WorkflowEndTask
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {

		// remove fields from embedded structs
		reflectWorkflowEndTask := reflect.ValueOf(o.WorkflowEndTask)
		for i := 0; i < reflectWorkflowEndTask.Type().NumField(); i++ {
			t := reflectWorkflowEndTask.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowFailureEndTask struct {
	value *WorkflowFailureEndTask
	isSet bool
}

func (v NullableWorkflowFailureEndTask) Get() *WorkflowFailureEndTask {
	return v.value
}

func (v *NullableWorkflowFailureEndTask) Set(val *WorkflowFailureEndTask) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowFailureEndTask) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowFailureEndTask) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowFailureEndTask(val *WorkflowFailureEndTask) *NullableWorkflowFailureEndTask {
	return &NullableWorkflowFailureEndTask{value: val, isSet: true}
}

func (v NullableWorkflowFailureEndTask) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowFailureEndTask) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
