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

// MoTagKeySummary A summary of the Tag usage for an API resource. Tags are Key, Value pairs that can be assigned to API resources.
type MoTagKeySummary struct {
	// The tag key for which usage information is provided.
	Key *string `json:"Key,omitempty"`
	// The number of times this tag Key has been set in an API resource.
	NumKeys *int32 `json:"NumKeys,omitempty"`
	// A list of all Tag values that have been assigned to this tag Key.
	Values               *[]string `json:"Values,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MoTagKeySummary MoTagKeySummary

// NewMoTagKeySummary instantiates a new MoTagKeySummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMoTagKeySummary() *MoTagKeySummary {
	this := MoTagKeySummary{}
	return &this
}

// NewMoTagKeySummaryWithDefaults instantiates a new MoTagKeySummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMoTagKeySummaryWithDefaults() *MoTagKeySummary {
	this := MoTagKeySummary{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *MoTagKeySummary) GetKey() string {
	if o == nil || o.Key == nil {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTagKeySummary) GetKeyOk() (*string, bool) {
	if o == nil || o.Key == nil {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *MoTagKeySummary) HasKey() bool {
	if o != nil && o.Key != nil {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *MoTagKeySummary) SetKey(v string) {
	o.Key = &v
}

// GetNumKeys returns the NumKeys field value if set, zero value otherwise.
func (o *MoTagKeySummary) GetNumKeys() int32 {
	if o == nil || o.NumKeys == nil {
		var ret int32
		return ret
	}
	return *o.NumKeys
}

// GetNumKeysOk returns a tuple with the NumKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTagKeySummary) GetNumKeysOk() (*int32, bool) {
	if o == nil || o.NumKeys == nil {
		return nil, false
	}
	return o.NumKeys, true
}

// HasNumKeys returns a boolean if a field has been set.
func (o *MoTagKeySummary) HasNumKeys() bool {
	if o != nil && o.NumKeys != nil {
		return true
	}

	return false
}

// SetNumKeys gets a reference to the given int32 and assigns it to the NumKeys field.
func (o *MoTagKeySummary) SetNumKeys(v int32) {
	o.NumKeys = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MoTagKeySummary) GetValues() []string {
	if o == nil || o.Values == nil {
		var ret []string
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MoTagKeySummary) GetValuesOk() (*[]string, bool) {
	if o == nil || o.Values == nil {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MoTagKeySummary) HasValues() bool {
	if o != nil && o.Values != nil {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *MoTagKeySummary) SetValues(v []string) {
	o.Values = &v
}

func (o MoTagKeySummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Key != nil {
		toSerialize["Key"] = o.Key
	}
	if o.NumKeys != nil {
		toSerialize["NumKeys"] = o.NumKeys
	}
	if o.Values != nil {
		toSerialize["Values"] = o.Values
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MoTagKeySummary) UnmarshalJSON(bytes []byte) (err error) {
	varMoTagKeySummary := _MoTagKeySummary{}

	if err = json.Unmarshal(bytes, &varMoTagKeySummary); err == nil {
		*o = MoTagKeySummary(varMoTagKeySummary)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "Key")
		delete(additionalProperties, "NumKeys")
		delete(additionalProperties, "Values")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMoTagKeySummary struct {
	value *MoTagKeySummary
	isSet bool
}

func (v NullableMoTagKeySummary) Get() *MoTagKeySummary {
	return v.value
}

func (v *NullableMoTagKeySummary) Set(val *MoTagKeySummary) {
	v.value = val
	v.isSet = true
}

func (v NullableMoTagKeySummary) IsSet() bool {
	return v.isSet
}

func (v *NullableMoTagKeySummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMoTagKeySummary(val *MoTagKeySummary) *NullableMoTagKeySummary {
	return &NullableMoTagKeySummary{value: val, isSet: true}
}

func (v NullableMoTagKeySummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMoTagKeySummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
