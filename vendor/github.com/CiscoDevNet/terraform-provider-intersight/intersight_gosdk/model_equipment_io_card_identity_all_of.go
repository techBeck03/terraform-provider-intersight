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

// EquipmentIoCardIdentityAllOf Definition of the list of properties defined in 'equipment.IoCardIdentity', excluding properties defined in parent classes.
type EquipmentIoCardIdentityAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// MO Reference to equipmentIoCard MO in inventory service.
	IoCardMoid *string `json:"IoCardMoid,omitempty"`
	// IOM/MUX Module ID connected to the FI.
	ModuleId *int64 `json:"ModuleId,omitempty"`
	// MO Reference to networkElement MO in inventory service.
	NetworkElementMoid *string `json:"NetworkElementMoid,omitempty"`
	// Switch ID to which IOM is connected, ID can be either 1 or 2, equalent to A or B.
	SwitchId             *int64 `json:"SwitchId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoCardIdentityAllOf EquipmentIoCardIdentityAllOf

// NewEquipmentIoCardIdentityAllOf instantiates a new EquipmentIoCardIdentityAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoCardIdentityAllOf(classId string, objectType string) *EquipmentIoCardIdentityAllOf {
	this := EquipmentIoCardIdentityAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoCardIdentityAllOfWithDefaults instantiates a new EquipmentIoCardIdentityAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoCardIdentityAllOfWithDefaults() *EquipmentIoCardIdentityAllOf {
	this := EquipmentIoCardIdentityAllOf{}
	var classId string = "equipment.IoCardIdentity"
	this.ClassId = classId
	var objectType string = "equipment.IoCardIdentity"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoCardIdentityAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoCardIdentityAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoCardIdentityAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoCardIdentityAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetIoCardMoid returns the IoCardMoid field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentityAllOf) GetIoCardMoid() string {
	if o == nil || o.IoCardMoid == nil {
		var ret string
		return ret
	}
	return *o.IoCardMoid
}

// GetIoCardMoidOk returns a tuple with the IoCardMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetIoCardMoidOk() (*string, bool) {
	if o == nil || o.IoCardMoid == nil {
		return nil, false
	}
	return o.IoCardMoid, true
}

// HasIoCardMoid returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentityAllOf) HasIoCardMoid() bool {
	if o != nil && o.IoCardMoid != nil {
		return true
	}

	return false
}

// SetIoCardMoid gets a reference to the given string and assigns it to the IoCardMoid field.
func (o *EquipmentIoCardIdentityAllOf) SetIoCardMoid(v string) {
	o.IoCardMoid = &v
}

// GetModuleId returns the ModuleId field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentityAllOf) GetModuleId() int64 {
	if o == nil || o.ModuleId == nil {
		var ret int64
		return ret
	}
	return *o.ModuleId
}

// GetModuleIdOk returns a tuple with the ModuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetModuleIdOk() (*int64, bool) {
	if o == nil || o.ModuleId == nil {
		return nil, false
	}
	return o.ModuleId, true
}

// HasModuleId returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentityAllOf) HasModuleId() bool {
	if o != nil && o.ModuleId != nil {
		return true
	}

	return false
}

// SetModuleId gets a reference to the given int64 and assigns it to the ModuleId field.
func (o *EquipmentIoCardIdentityAllOf) SetModuleId(v int64) {
	o.ModuleId = &v
}

// GetNetworkElementMoid returns the NetworkElementMoid field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentityAllOf) GetNetworkElementMoid() string {
	if o == nil || o.NetworkElementMoid == nil {
		var ret string
		return ret
	}
	return *o.NetworkElementMoid
}

// GetNetworkElementMoidOk returns a tuple with the NetworkElementMoid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetNetworkElementMoidOk() (*string, bool) {
	if o == nil || o.NetworkElementMoid == nil {
		return nil, false
	}
	return o.NetworkElementMoid, true
}

// HasNetworkElementMoid returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentityAllOf) HasNetworkElementMoid() bool {
	if o != nil && o.NetworkElementMoid != nil {
		return true
	}

	return false
}

// SetNetworkElementMoid gets a reference to the given string and assigns it to the NetworkElementMoid field.
func (o *EquipmentIoCardIdentityAllOf) SetNetworkElementMoid(v string) {
	o.NetworkElementMoid = &v
}

// GetSwitchId returns the SwitchId field value if set, zero value otherwise.
func (o *EquipmentIoCardIdentityAllOf) GetSwitchId() int64 {
	if o == nil || o.SwitchId == nil {
		var ret int64
		return ret
	}
	return *o.SwitchId
}

// GetSwitchIdOk returns a tuple with the SwitchId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoCardIdentityAllOf) GetSwitchIdOk() (*int64, bool) {
	if o == nil || o.SwitchId == nil {
		return nil, false
	}
	return o.SwitchId, true
}

// HasSwitchId returns a boolean if a field has been set.
func (o *EquipmentIoCardIdentityAllOf) HasSwitchId() bool {
	if o != nil && o.SwitchId != nil {
		return true
	}

	return false
}

// SetSwitchId gets a reference to the given int64 and assigns it to the SwitchId field.
func (o *EquipmentIoCardIdentityAllOf) SetSwitchId(v int64) {
	o.SwitchId = &v
}

func (o EquipmentIoCardIdentityAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.IoCardMoid != nil {
		toSerialize["IoCardMoid"] = o.IoCardMoid
	}
	if o.ModuleId != nil {
		toSerialize["ModuleId"] = o.ModuleId
	}
	if o.NetworkElementMoid != nil {
		toSerialize["NetworkElementMoid"] = o.NetworkElementMoid
	}
	if o.SwitchId != nil {
		toSerialize["SwitchId"] = o.SwitchId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoCardIdentityAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoCardIdentityAllOf := _EquipmentIoCardIdentityAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoCardIdentityAllOf); err == nil {
		*o = EquipmentIoCardIdentityAllOf(varEquipmentIoCardIdentityAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "IoCardMoid")
		delete(additionalProperties, "ModuleId")
		delete(additionalProperties, "NetworkElementMoid")
		delete(additionalProperties, "SwitchId")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoCardIdentityAllOf struct {
	value *EquipmentIoCardIdentityAllOf
	isSet bool
}

func (v NullableEquipmentIoCardIdentityAllOf) Get() *EquipmentIoCardIdentityAllOf {
	return v.value
}

func (v *NullableEquipmentIoCardIdentityAllOf) Set(val *EquipmentIoCardIdentityAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoCardIdentityAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoCardIdentityAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoCardIdentityAllOf(val *EquipmentIoCardIdentityAllOf) *NullableEquipmentIoCardIdentityAllOf {
	return &NullableEquipmentIoCardIdentityAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoCardIdentityAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoCardIdentityAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
