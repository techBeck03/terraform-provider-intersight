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

// HyperflexServerFirmwareVersionInfoAllOf Definition of the list of properties defined in 'hyperflex.ServerFirmwareVersionInfo', excluding properties defined in parent classes.
type HyperflexServerFirmwareVersionInfoAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The platform type for UCS server. * `M5` - M5 generation of UCS server. * `M3` - M3 generation of UCS server. * `M4` - M4 generation of UCS server.
	ServerPlatform *string `json:"ServerPlatform,omitempty"`
	// The server firmware bundle version.
	Version              *string `json:"Version,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _HyperflexServerFirmwareVersionInfoAllOf HyperflexServerFirmwareVersionInfoAllOf

// NewHyperflexServerFirmwareVersionInfoAllOf instantiates a new HyperflexServerFirmwareVersionInfoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHyperflexServerFirmwareVersionInfoAllOf(classId string, objectType string) *HyperflexServerFirmwareVersionInfoAllOf {
	this := HyperflexServerFirmwareVersionInfoAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// NewHyperflexServerFirmwareVersionInfoAllOfWithDefaults instantiates a new HyperflexServerFirmwareVersionInfoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHyperflexServerFirmwareVersionInfoAllOfWithDefaults() *HyperflexServerFirmwareVersionInfoAllOf {
	this := HyperflexServerFirmwareVersionInfoAllOf{}
	var classId string = "hyperflex.ServerFirmwareVersionInfo"
	this.ClassId = classId
	var objectType string = "hyperflex.ServerFirmwareVersionInfo"
	this.ObjectType = objectType
	var serverPlatform string = "M5"
	this.ServerPlatform = &serverPlatform
	return &this
}

// GetClassId returns the ClassId field value
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *HyperflexServerFirmwareVersionInfoAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *HyperflexServerFirmwareVersionInfoAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetServerPlatform returns the ServerPlatform field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetServerPlatform() string {
	if o == nil || o.ServerPlatform == nil {
		var ret string
		return ret
	}
	return *o.ServerPlatform
}

// GetServerPlatformOk returns a tuple with the ServerPlatform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetServerPlatformOk() (*string, bool) {
	if o == nil || o.ServerPlatform == nil {
		return nil, false
	}
	return o.ServerPlatform, true
}

// HasServerPlatform returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) HasServerPlatform() bool {
	if o != nil && o.ServerPlatform != nil {
		return true
	}

	return false
}

// SetServerPlatform gets a reference to the given string and assigns it to the ServerPlatform field.
func (o *HyperflexServerFirmwareVersionInfoAllOf) SetServerPlatform(v string) {
	o.ServerPlatform = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *HyperflexServerFirmwareVersionInfoAllOf) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *HyperflexServerFirmwareVersionInfoAllOf) SetVersion(v string) {
	o.Version = &v
}

func (o HyperflexServerFirmwareVersionInfoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ServerPlatform != nil {
		toSerialize["ServerPlatform"] = o.ServerPlatform
	}
	if o.Version != nil {
		toSerialize["Version"] = o.Version
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *HyperflexServerFirmwareVersionInfoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varHyperflexServerFirmwareVersionInfoAllOf := _HyperflexServerFirmwareVersionInfoAllOf{}

	if err = json.Unmarshal(bytes, &varHyperflexServerFirmwareVersionInfoAllOf); err == nil {
		*o = HyperflexServerFirmwareVersionInfoAllOf(varHyperflexServerFirmwareVersionInfoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ServerPlatform")
		delete(additionalProperties, "Version")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableHyperflexServerFirmwareVersionInfoAllOf struct {
	value *HyperflexServerFirmwareVersionInfoAllOf
	isSet bool
}

func (v NullableHyperflexServerFirmwareVersionInfoAllOf) Get() *HyperflexServerFirmwareVersionInfoAllOf {
	return v.value
}

func (v *NullableHyperflexServerFirmwareVersionInfoAllOf) Set(val *HyperflexServerFirmwareVersionInfoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHyperflexServerFirmwareVersionInfoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHyperflexServerFirmwareVersionInfoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHyperflexServerFirmwareVersionInfoAllOf(val *HyperflexServerFirmwareVersionInfoAllOf) *NullableHyperflexServerFirmwareVersionInfoAllOf {
	return &NullableHyperflexServerFirmwareVersionInfoAllOf{value: val, isSet: true}
}

func (v NullableHyperflexServerFirmwareVersionInfoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHyperflexServerFirmwareVersionInfoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
