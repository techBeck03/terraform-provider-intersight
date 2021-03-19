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

// IppoolIpV4BlockAllOf Definition of the list of properties defined in 'ippool.IpV4Block', excluding properties defined in parent classes.
type IppoolIpV4BlockAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// First IPv4 address of the block.
	From *string `json:"From,omitempty"`
	// Last IPv4 address of the block.
	To                   *string `json:"To,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IppoolIpV4BlockAllOf IppoolIpV4BlockAllOf

// NewIppoolIpV4BlockAllOf instantiates a new IppoolIpV4BlockAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIppoolIpV4BlockAllOf(classId string, objectType string) *IppoolIpV4BlockAllOf {
	this := IppoolIpV4BlockAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIppoolIpV4BlockAllOfWithDefaults instantiates a new IppoolIpV4BlockAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIppoolIpV4BlockAllOfWithDefaults() *IppoolIpV4BlockAllOf {
	this := IppoolIpV4BlockAllOf{}
	var classId string = "ippool.IpV4Block"
	this.ClassId = classId
	var objectType string = "ippool.IpV4Block"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IppoolIpV4BlockAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4BlockAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IppoolIpV4BlockAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IppoolIpV4BlockAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IppoolIpV4BlockAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IppoolIpV4BlockAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *IppoolIpV4BlockAllOf) GetFrom() string {
	if o == nil || o.From == nil {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4BlockAllOf) GetFromOk() (*string, bool) {
	if o == nil || o.From == nil {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *IppoolIpV4BlockAllOf) HasFrom() bool {
	if o != nil && o.From != nil {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *IppoolIpV4BlockAllOf) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value if set, zero value otherwise.
func (o *IppoolIpV4BlockAllOf) GetTo() string {
	if o == nil || o.To == nil {
		var ret string
		return ret
	}
	return *o.To
}

// GetToOk returns a tuple with the To field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IppoolIpV4BlockAllOf) GetToOk() (*string, bool) {
	if o == nil || o.To == nil {
		return nil, false
	}
	return o.To, true
}

// HasTo returns a boolean if a field has been set.
func (o *IppoolIpV4BlockAllOf) HasTo() bool {
	if o != nil && o.To != nil {
		return true
	}

	return false
}

// SetTo gets a reference to the given string and assigns it to the To field.
func (o *IppoolIpV4BlockAllOf) SetTo(v string) {
	o.To = &v
}

func (o IppoolIpV4BlockAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.From != nil {
		toSerialize["From"] = o.From
	}
	if o.To != nil {
		toSerialize["To"] = o.To
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IppoolIpV4BlockAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIppoolIpV4BlockAllOf := _IppoolIpV4BlockAllOf{}

	if err = json.Unmarshal(bytes, &varIppoolIpV4BlockAllOf); err == nil {
		*o = IppoolIpV4BlockAllOf(varIppoolIpV4BlockAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "From")
		delete(additionalProperties, "To")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIppoolIpV4BlockAllOf struct {
	value *IppoolIpV4BlockAllOf
	isSet bool
}

func (v NullableIppoolIpV4BlockAllOf) Get() *IppoolIpV4BlockAllOf {
	return v.value
}

func (v *NullableIppoolIpV4BlockAllOf) Set(val *IppoolIpV4BlockAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIppoolIpV4BlockAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIppoolIpV4BlockAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIppoolIpV4BlockAllOf(val *IppoolIpV4BlockAllOf) *NullableIppoolIpV4BlockAllOf {
	return &NullableIppoolIpV4BlockAllOf{value: val, isSet: true}
}

func (v NullableIppoolIpV4BlockAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIppoolIpV4BlockAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
