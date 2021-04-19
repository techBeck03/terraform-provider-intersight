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
)

// WorkflowWaitTaskAllOf Definition of the list of properties defined in 'workflow.WaitTask', excluding properties defined in parent classes.
type WorkflowWaitTaskAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                   `json:"ObjectType"`
	Prompts              []WorkflowWaitTaskPrompt `json:"Prompts,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowWaitTaskAllOf WorkflowWaitTaskAllOf

// NewWorkflowWaitTaskAllOf instantiates a new WorkflowWaitTaskAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowWaitTaskAllOf(classId string, objectType string) *WorkflowWaitTaskAllOf {
	this := WorkflowWaitTaskAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewWorkflowWaitTaskAllOfWithDefaults instantiates a new WorkflowWaitTaskAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWaitTaskAllOfWithDefaults() *WorkflowWaitTaskAllOf {
	this := WorkflowWaitTaskAllOf{}
	var classId string = "workflow.WaitTask"
	this.ClassId = classId
	var objectType string = "workflow.WaitTask"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *WorkflowWaitTaskAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *WorkflowWaitTaskAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *WorkflowWaitTaskAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *WorkflowWaitTaskAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *WorkflowWaitTaskAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetPrompts returns the Prompts field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WorkflowWaitTaskAllOf) GetPrompts() []WorkflowWaitTaskPrompt {
	if o == nil {
		var ret []WorkflowWaitTaskPrompt
		return ret
	}
	return o.Prompts
}

// GetPromptsOk returns a tuple with the Prompts field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WorkflowWaitTaskAllOf) GetPromptsOk() (*[]WorkflowWaitTaskPrompt, bool) {
	if o == nil || o.Prompts == nil {
		return nil, false
	}
	return &o.Prompts, true
}

// HasPrompts returns a boolean if a field has been set.
func (o *WorkflowWaitTaskAllOf) HasPrompts() bool {
	if o != nil && o.Prompts != nil {
		return true
	}

	return false
}

// SetPrompts gets a reference to the given []WorkflowWaitTaskPrompt and assigns it to the Prompts field.
func (o *WorkflowWaitTaskAllOf) SetPrompts(v []WorkflowWaitTaskPrompt) {
	o.Prompts = v
}

func (o WorkflowWaitTaskAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Prompts != nil {
		toSerialize["Prompts"] = o.Prompts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowWaitTaskAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowWaitTaskAllOf := _WorkflowWaitTaskAllOf{}

	if err = json.Unmarshal(bytes, &varWorkflowWaitTaskAllOf); err == nil {
		*o = WorkflowWaitTaskAllOf(varWorkflowWaitTaskAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Prompts")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowWaitTaskAllOf struct {
	value *WorkflowWaitTaskAllOf
	isSet bool
}

func (v NullableWorkflowWaitTaskAllOf) Get() *WorkflowWaitTaskAllOf {
	return v.value
}

func (v *NullableWorkflowWaitTaskAllOf) Set(val *WorkflowWaitTaskAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowWaitTaskAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowWaitTaskAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowWaitTaskAllOf(val *WorkflowWaitTaskAllOf) *NullableWorkflowWaitTaskAllOf {
	return &NullableWorkflowWaitTaskAllOf{value: val, isSet: true}
}

func (v NullableWorkflowWaitTaskAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowWaitTaskAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
