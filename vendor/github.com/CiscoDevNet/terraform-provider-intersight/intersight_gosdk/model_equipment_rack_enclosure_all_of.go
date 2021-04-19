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

// EquipmentRackEnclosureAllOf Definition of the list of properties defined in 'equipment.RackEnclosure', excluding properties defined in parent classes.
type EquipmentRackEnclosureAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// This represents the Enclosure Identifier for Rack servers.
	EnclosureId *int64 `json:"EnclosureId,omitempty"`
	// An array of relationships to equipmentFanModule resources.
	Fanmodules          []EquipmentFanModuleRelationship `json:"Fanmodules,omitempty"`
	InventoryDeviceInfo *InventoryDeviceInfoRelationship `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to equipmentPsu resources.
	Psus             []EquipmentPsuRelationship           `json:"Psus,omitempty"`
	RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	// An array of relationships to equipmentRackEnclosureSlot resources.
	Slots                []EquipmentRackEnclosureSlotRelationship `json:"Slots,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentRackEnclosureAllOf EquipmentRackEnclosureAllOf

// NewEquipmentRackEnclosureAllOf instantiates a new EquipmentRackEnclosureAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentRackEnclosureAllOf(classId string, objectType string) *EquipmentRackEnclosureAllOf {
	this := EquipmentRackEnclosureAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewEquipmentRackEnclosureAllOfWithDefaults instantiates a new EquipmentRackEnclosureAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentRackEnclosureAllOfWithDefaults() *EquipmentRackEnclosureAllOf {
	this := EquipmentRackEnclosureAllOf{}
	var classId string = "equipment.RackEnclosure"
	this.ClassId = classId
	var objectType string = "equipment.RackEnclosure"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *EquipmentRackEnclosureAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *EquipmentRackEnclosureAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *EquipmentRackEnclosureAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *EquipmentRackEnclosureAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetEnclosureId returns the EnclosureId field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureAllOf) GetEnclosureId() int64 {
	if o == nil || o.EnclosureId == nil {
		var ret int64
		return ret
	}
	return *o.EnclosureId
}

// GetEnclosureIdOk returns a tuple with the EnclosureId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureAllOf) GetEnclosureIdOk() (*int64, bool) {
	if o == nil || o.EnclosureId == nil {
		return nil, false
	}
	return o.EnclosureId, true
}

// HasEnclosureId returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasEnclosureId() bool {
	if o != nil && o.EnclosureId != nil {
		return true
	}

	return false
}

// SetEnclosureId gets a reference to the given int64 and assigns it to the EnclosureId field.
func (o *EquipmentRackEnclosureAllOf) SetEnclosureId(v int64) {
	o.EnclosureId = &v
}

// GetFanmodules returns the Fanmodules field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosureAllOf) GetFanmodules() []EquipmentFanModuleRelationship {
	if o == nil {
		var ret []EquipmentFanModuleRelationship
		return ret
	}
	return o.Fanmodules
}

// GetFanmodulesOk returns a tuple with the Fanmodules field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosureAllOf) GetFanmodulesOk() (*[]EquipmentFanModuleRelationship, bool) {
	if o == nil || o.Fanmodules == nil {
		return nil, false
	}
	return &o.Fanmodules, true
}

// HasFanmodules returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasFanmodules() bool {
	if o != nil && o.Fanmodules != nil {
		return true
	}

	return false
}

// SetFanmodules gets a reference to the given []EquipmentFanModuleRelationship and assigns it to the Fanmodules field.
func (o *EquipmentRackEnclosureAllOf) SetFanmodules(v []EquipmentFanModuleRelationship) {
	o.Fanmodules = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureAllOf) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureAllOf) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentRackEnclosureAllOf) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPsus returns the Psus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosureAllOf) GetPsus() []EquipmentPsuRelationship {
	if o == nil {
		var ret []EquipmentPsuRelationship
		return ret
	}
	return o.Psus
}

// GetPsusOk returns a tuple with the Psus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosureAllOf) GetPsusOk() (*[]EquipmentPsuRelationship, bool) {
	if o == nil || o.Psus == nil {
		return nil, false
	}
	return &o.Psus, true
}

// HasPsus returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasPsus() bool {
	if o != nil && o.Psus != nil {
		return true
	}

	return false
}

// SetPsus gets a reference to the given []EquipmentPsuRelationship and assigns it to the Psus field.
func (o *EquipmentRackEnclosureAllOf) SetPsus(v []EquipmentPsuRelationship) {
	o.Psus = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentRackEnclosureAllOf) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentRackEnclosureAllOf) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentRackEnclosureAllOf) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetSlots returns the Slots field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentRackEnclosureAllOf) GetSlots() []EquipmentRackEnclosureSlotRelationship {
	if o == nil {
		var ret []EquipmentRackEnclosureSlotRelationship
		return ret
	}
	return o.Slots
}

// GetSlotsOk returns a tuple with the Slots field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentRackEnclosureAllOf) GetSlotsOk() (*[]EquipmentRackEnclosureSlotRelationship, bool) {
	if o == nil || o.Slots == nil {
		return nil, false
	}
	return &o.Slots, true
}

// HasSlots returns a boolean if a field has been set.
func (o *EquipmentRackEnclosureAllOf) HasSlots() bool {
	if o != nil && o.Slots != nil {
		return true
	}

	return false
}

// SetSlots gets a reference to the given []EquipmentRackEnclosureSlotRelationship and assigns it to the Slots field.
func (o *EquipmentRackEnclosureAllOf) SetSlots(v []EquipmentRackEnclosureSlotRelationship) {
	o.Slots = v
}

func (o EquipmentRackEnclosureAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.EnclosureId != nil {
		toSerialize["EnclosureId"] = o.EnclosureId
	}
	if o.Fanmodules != nil {
		toSerialize["Fanmodules"] = o.Fanmodules
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.Psus != nil {
		toSerialize["Psus"] = o.Psus
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.Slots != nil {
		toSerialize["Slots"] = o.Slots
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentRackEnclosureAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varEquipmentRackEnclosureAllOf := _EquipmentRackEnclosureAllOf{}

	if err = json.Unmarshal(bytes, &varEquipmentRackEnclosureAllOf); err == nil {
		*o = EquipmentRackEnclosureAllOf(varEquipmentRackEnclosureAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "EnclosureId")
		delete(additionalProperties, "Fanmodules")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "Psus")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "Slots")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableEquipmentRackEnclosureAllOf struct {
	value *EquipmentRackEnclosureAllOf
	isSet bool
}

func (v NullableEquipmentRackEnclosureAllOf) Get() *EquipmentRackEnclosureAllOf {
	return v.value
}

func (v *NullableEquipmentRackEnclosureAllOf) Set(val *EquipmentRackEnclosureAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentRackEnclosureAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentRackEnclosureAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentRackEnclosureAllOf(val *EquipmentRackEnclosureAllOf) *NullableEquipmentRackEnclosureAllOf {
	return &NullableEquipmentRackEnclosureAllOf{value: val, isSet: true}
}

func (v NullableEquipmentRackEnclosureAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentRackEnclosureAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
