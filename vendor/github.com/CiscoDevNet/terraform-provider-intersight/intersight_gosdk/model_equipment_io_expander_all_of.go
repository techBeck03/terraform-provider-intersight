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

// EquipmentIoExpanderAllOf Definition of the list of properties defined in 'equipment.IoExpander', excluding properties defined in parent classes.
type EquipmentIoExpanderAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Identifies the operational state of I/O expander.
	OperState            *string                              `json:"OperState,omitempty"`
	ComputeBlade         *ComputeBladeRelationship            `json:"ComputeBlade,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentIoExpanderAllOf EquipmentIoExpanderAllOf

// NewEquipmentIoExpanderAllOf instantiates a new EquipmentIoExpanderAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentIoExpanderAllOf(classId string, objectType string) *EquipmentIoExpanderAllOf {
	this := EquipmentIoExpanderAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentIoExpanderAllOfWithDefaults instantiates a new EquipmentIoExpanderAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentIoExpanderAllOfWithDefaults() *EquipmentIoExpanderAllOf {
	this := EquipmentIoExpanderAllOf{}
	var classId string = "equipment.IoExpander"
	this.ClassId = classId
	var objectType string = "equipment.IoExpander"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentIoExpanderAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentIoExpanderAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentIoExpanderAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentIoExpanderAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentIoExpanderAllOf) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentIoExpanderAllOf) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentIoExpanderAllOf) SetOperState(v string) {
	o.OperState = &v
}

// GetComputeBlade returns the ComputeBlade field value if set, zero value otherwise.
func (o *EquipmentIoExpanderAllOf) GetComputeBlade() ComputeBladeRelationship {
	if o == nil || o.ComputeBlade == nil {
		var ret ComputeBladeRelationship
		return ret
	}
	return *o.ComputeBlade
}

// GetComputeBladeOk returns a tuple with the ComputeBlade field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetComputeBladeOk() (*ComputeBladeRelationship, bool) {
	if o == nil || o.ComputeBlade == nil {
		return nil, false
	}
	return o.ComputeBlade, true
}

// HasComputeBlade returns a boolean if a field has been set.
func (o *EquipmentIoExpanderAllOf) HasComputeBlade() bool {
	if o != nil && o.ComputeBlade != nil {
		return true
	}

	return false
}

// SetComputeBlade gets a reference to the given ComputeBladeRelationship and assigns it to the ComputeBlade field.
func (o *EquipmentIoExpanderAllOf) SetComputeBlade(v ComputeBladeRelationship) {
	o.ComputeBlade = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentIoExpanderAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentIoExpanderAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentIoExpanderAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentIoExpanderAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentIoExpanderAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentIoExpanderAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentIoExpanderAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentIoExpanderAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.ComputeBlade != nil {
		toSerialize["ComputeBlade"] = o.ComputeBlade
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentIoExpanderAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentIoExpanderAllOf := _EquipmentIoExpanderAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentIoExpanderAllOf); err == nil {
		*o = EquipmentIoExpanderAllOf(varEquipmentIoExpanderAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "ComputeBlade")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentIoExpanderAllOf struct {
	value *EquipmentIoExpanderAllOf
	isSet bool
}

func (v NullableEquipmentIoExpanderAllOf) Get() *EquipmentIoExpanderAllOf {
	return v.value
}

func (v *NullableEquipmentIoExpanderAllOf) Set(val *EquipmentIoExpanderAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentIoExpanderAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentIoExpanderAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentIoExpanderAllOf(val *EquipmentIoExpanderAllOf) *NullableEquipmentIoExpanderAllOf {
	return &NullableEquipmentIoExpanderAllOf{value: val, isSet: true}
}

func (v NullableEquipmentIoExpanderAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentIoExpanderAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
