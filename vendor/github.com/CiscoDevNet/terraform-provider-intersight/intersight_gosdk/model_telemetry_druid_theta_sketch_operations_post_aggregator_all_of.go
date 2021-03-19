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
)

// TelemetryDruidThetaSketchOperationsPostAggregatorAllOf struct for TelemetryDruidThetaSketchOperationsPostAggregatorAllOf
type TelemetryDruidThetaSketchOperationsPostAggregatorAllOf struct {
	// Output name for the post-aggregator.
	Name *string `json:"name,omitempty"`
	Func *string `json:"func,omitempty"`
	// array of fieldAccess type post aggregators to access the thetaSketch aggregators or thetaSketchSetOp type post aggregators to allow arbitrary combination of set operations.
	Fields *[]string `json:"fields,omitempty"`
	// must be max of size from sketches in fields input.
	Size                 *int32 `json:"size,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidThetaSketchOperationsPostAggregatorAllOf TelemetryDruidThetaSketchOperationsPostAggregatorAllOf

// NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOf instantiates a new TelemetryDruidThetaSketchOperationsPostAggregatorAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOf() *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf {
	this := TelemetryDruidThetaSketchOperationsPostAggregatorAllOf{}
	var size int32 = 16384
	this.Size = &size
	return &this
}

// NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOfWithDefaults instantiates a new TelemetryDruidThetaSketchOperationsPostAggregatorAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidThetaSketchOperationsPostAggregatorAllOfWithDefaults() *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf {
	this := TelemetryDruidThetaSketchOperationsPostAggregatorAllOf{}
	var size int32 = 16384
	this.Size = &size
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetName(v string) {
	o.Name = &v
}

// GetFunc returns the Func field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFunc() string {
	if o == nil || o.Func == nil {
		var ret string
		return ret
	}
	return *o.Func
}

// GetFuncOk returns a tuple with the Func field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFuncOk() (*string, bool) {
	if o == nil || o.Func == nil {
		return nil, false
	}
	return o.Func, true
}

// HasFunc returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasFunc() bool {
	if o != nil && o.Func != nil {
		return true
	}

	return false
}

// SetFunc gets a reference to the given string and assigns it to the Func field.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetFunc(v string) {
	o.Func = &v
}

// GetFields returns the Fields field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFields() []string {
	if o == nil || o.Fields == nil {
		var ret []string
		return ret
	}
	return *o.Fields
}

// GetFieldsOk returns a tuple with the Fields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetFieldsOk() (*[]string, bool) {
	if o == nil || o.Fields == nil {
		return nil, false
	}
	return o.Fields, true
}

// HasFields returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasFields() bool {
	if o != nil && o.Fields != nil {
		return true
	}

	return false
}

// SetFields gets a reference to the given []string and assigns it to the Fields field.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetFields(v []string) {
	o.Fields = &v
}

// GetSize returns the Size field value if set, zero value otherwise.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetSize() int32 {
	if o == nil || o.Size == nil {
		var ret int32
		return ret
	}
	return *o.Size
}

// GetSizeOk returns a tuple with the Size field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) GetSizeOk() (*int32, bool) {
	if o == nil || o.Size == nil {
		return nil, false
	}
	return o.Size, true
}

// HasSize returns a boolean if a field has been set.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) HasSize() bool {
	if o != nil && o.Size != nil {
		return true
	}

	return false
}

// SetSize gets a reference to the given int32 and assigns it to the Size field.
func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) SetSize(v int32) {
	o.Size = &v
}

func (o TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Func != nil {
		toSerialize["func"] = o.Func
	}
	if o.Fields != nil {
		toSerialize["fields"] = o.Fields
	}
	if o.Size != nil {
		toSerialize["size"] = o.Size
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidThetaSketchOperationsPostAggregatorAllOf := _TelemetryDruidThetaSketchOperationsPostAggregatorAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidThetaSketchOperationsPostAggregatorAllOf); err == nil {
		*o = TelemetryDruidThetaSketchOperationsPostAggregatorAllOf(varTelemetryDruidThetaSketchOperationsPostAggregatorAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "func")
		delete(additionalProperties, "fields")
		delete(additionalProperties, "size")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf struct {
	value *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf
	isSet bool
}

func (v NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) Get() *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf {
	return v.value
}

func (v *NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) Set(val *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf(val *TelemetryDruidThetaSketchOperationsPostAggregatorAllOf) *NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf {
	return &NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidThetaSketchOperationsPostAggregatorAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
