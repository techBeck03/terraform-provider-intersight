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

// FirmwareHttpServerAllOf Definition of the list of properties defined in 'firmware.HttpServer', excluding properties defined in parent classes.
type FirmwareHttpServerAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// HTTP/HTTPS link to the image. Accepted formats HTTP[s]://server-hostname/share/image or HTTP[s]://serverip/share/image. For a successful upgrade, the remote share server must have network connectivity with the CIMC of the selected devices.
	LocationLink *string `json:"LocationLink,omitempty"`
	// Mount option as configured on the HTTP server. Empty if nothing is configured.
	MountOptions         *string `json:"MountOptions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FirmwareHttpServerAllOf FirmwareHttpServerAllOf

// NewFirmwareHttpServerAllOf instantiates a new FirmwareHttpServerAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFirmwareHttpServerAllOf(classId string, objectType string) *FirmwareHttpServerAllOf {
	this := FirmwareHttpServerAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewFirmwareHttpServerAllOfWithDefaults instantiates a new FirmwareHttpServerAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFirmwareHttpServerAllOfWithDefaults() *FirmwareHttpServerAllOf {
	this := FirmwareHttpServerAllOf{}
	var classId string = "firmware.HttpServer"
	this.ClassId = classId
	var objectType string = "firmware.HttpServer"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *FirmwareHttpServerAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServerAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *FirmwareHttpServerAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *FirmwareHttpServerAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServerAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *FirmwareHttpServerAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetLocationLink returns the LocationLink field value if set, zero value otherwise.
func (o *FirmwareHttpServerAllOf) GetLocationLink() string {
	if o == nil || o.LocationLink == nil {
		var ret string
		return ret
	}
	return *o.LocationLink
}

// GetLocationLinkOk returns a tuple with the LocationLink field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServerAllOf) GetLocationLinkOk() (*string, bool) {
	if o == nil || o.LocationLink == nil {
		return nil, false
	}
	return o.LocationLink, true
}

// HasLocationLink returns a boolean if a field has been set.
func (o *FirmwareHttpServerAllOf) HasLocationLink() bool {
	if o != nil && o.LocationLink != nil {
		return true
	}

	return false
}

// SetLocationLink gets a reference to the given string and assigns it to the LocationLink field.
func (o *FirmwareHttpServerAllOf) SetLocationLink(v string) {
	o.LocationLink = &v
}

// GetMountOptions returns the MountOptions field value if set, zero value otherwise.
func (o *FirmwareHttpServerAllOf) GetMountOptions() string {
	if o == nil || o.MountOptions == nil {
		var ret string
		return ret
	}
	return *o.MountOptions
}

// GetMountOptionsOk returns a tuple with the MountOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FirmwareHttpServerAllOf) GetMountOptionsOk() (*string, bool) {
	if o == nil || o.MountOptions == nil {
		return nil, false
	}
	return o.MountOptions, true
}

// HasMountOptions returns a boolean if a field has been set.
func (o *FirmwareHttpServerAllOf) HasMountOptions() bool {
	if o != nil && o.MountOptions != nil {
		return true
	}

	return false
}

// SetMountOptions gets a reference to the given string and assigns it to the MountOptions field.
func (o *FirmwareHttpServerAllOf) SetMountOptions(v string) {
	o.MountOptions = &v
}

func (o FirmwareHttpServerAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.LocationLink != nil {
		toSerialize["LocationLink"] = o.LocationLink
	}
	if o.MountOptions != nil {
		toSerialize["MountOptions"] = o.MountOptions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FirmwareHttpServerAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varFirmwareHttpServerAllOf := _FirmwareHttpServerAllOf{}

	if err = json.Unmarshal(bytes, &varFirmwareHttpServerAllOf); err == nil {
		*o = FirmwareHttpServerAllOf(varFirmwareHttpServerAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "LocationLink")
		delete(additionalProperties, "MountOptions")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFirmwareHttpServerAllOf struct {
	value *FirmwareHttpServerAllOf
	isSet bool
}

func (v NullableFirmwareHttpServerAllOf) Get() *FirmwareHttpServerAllOf {
	return v.value
}

func (v *NullableFirmwareHttpServerAllOf) Set(val *FirmwareHttpServerAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFirmwareHttpServerAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFirmwareHttpServerAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFirmwareHttpServerAllOf(val *FirmwareHttpServerAllOf) *NullableFirmwareHttpServerAllOf {
	return &NullableFirmwareHttpServerAllOf{value: val, isSet: true}
}

func (v NullableFirmwareHttpServerAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFirmwareHttpServerAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
