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
	"reflect"
	"strings"
)

// StorageFlexFlashController Storage Flex Flash Controller.
type StorageFlexFlashController struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// State of the Flex Flash Storage Controller.
	ControllerState *string `json:"ControllerState,omitempty"`
	// Identifier for the Flex Flash Storage Controller.
	FfControllerId *string                   `json:"FfControllerId,omitempty"`
	ComputeBoard   *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
	// An array of relationships to storageFlexFlashControllerProps resources.
	FlexFlashControllerProps []StorageFlexFlashControllerPropsRelationship `json:"FlexFlashControllerProps,omitempty"`
	// An array of relationships to storageFlexFlashPhysicalDrive resources.
	FlexFlashPhysicalDrives []StorageFlexFlashPhysicalDriveRelationship `json:"FlexFlashPhysicalDrives,omitempty"`
	// An array of relationships to storageFlexFlashVirtualDrive resources.
	FlexFlashVirtualDrives []StorageFlexFlashVirtualDriveRelationship `json:"FlexFlashVirtualDrives,omitempty"`
	InventoryDeviceInfo    *InventoryDeviceInfoRelationship           `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice       *AssetDeviceRegistrationRelationship       `json:"RegisteredDevice,omitempty"`
	// An array of relationships to firmwareRunningFirmware resources.
	RunningFirmware      []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageFlexFlashController StorageFlexFlashController

// NewStorageFlexFlashController instantiates a new StorageFlexFlashController object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageFlexFlashController(classId string, objectType string) *StorageFlexFlashController {
	this := StorageFlexFlashController{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageFlexFlashControllerWithDefaults instantiates a new StorageFlexFlashController object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageFlexFlashControllerWithDefaults() *StorageFlexFlashController {
	this := StorageFlexFlashController{}
	var classId string = "storage.FlexFlashController"
	this.ClassId = classId
	var objectType string = "storage.FlexFlashController"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageFlexFlashController) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageFlexFlashController) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageFlexFlashController) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageFlexFlashController) SetObjectType(v string) {
	o.ObjectType = v
}

// GetControllerState returns the ControllerState field value if set, zero value otherwise.
func (o *StorageFlexFlashController) GetControllerState() string {
	if o == nil || o.ControllerState == nil {
		var ret string
		return ret
	}
	return *o.ControllerState
}

// GetControllerStateOk returns a tuple with the ControllerState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetControllerStateOk() (*string, bool) {
	if o == nil || o.ControllerState == nil {
		return nil, false
	}
	return o.ControllerState, true
}

// HasControllerState returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasControllerState() bool {
	if o != nil && o.ControllerState != nil {
		return true
	}

	return false
}

// SetControllerState gets a reference to the given string and assigns it to the ControllerState field.
func (o *StorageFlexFlashController) SetControllerState(v string) {
	o.ControllerState = &v
}

// GetFfControllerId returns the FfControllerId field value if set, zero value otherwise.
func (o *StorageFlexFlashController) GetFfControllerId() string {
	if o == nil || o.FfControllerId == nil {
		var ret string
		return ret
	}
	return *o.FfControllerId
}

// GetFfControllerIdOk returns a tuple with the FfControllerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetFfControllerIdOk() (*string, bool) {
	if o == nil || o.FfControllerId == nil {
		return nil, false
	}
	return o.FfControllerId, true
}

// HasFfControllerId returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasFfControllerId() bool {
	if o != nil && o.FfControllerId != nil {
		return true
	}

	return false
}

// SetFfControllerId gets a reference to the given string and assigns it to the FfControllerId field.
func (o *StorageFlexFlashController) SetFfControllerId(v string) {
	o.FfControllerId = &v
}

// GetComputeBoard returns the ComputeBoard field value if set, zero value otherwise.
func (o *StorageFlexFlashController) GetComputeBoard() ComputeBoardRelationship {
	if o == nil || o.ComputeBoard == nil {
		var ret ComputeBoardRelationship
		return ret
	}
	return *o.ComputeBoard
}

// GetComputeBoardOk returns a tuple with the ComputeBoard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetComputeBoardOk() (*ComputeBoardRelationship, bool) {
	if o == nil || o.ComputeBoard == nil {
		return nil, false
	}
	return o.ComputeBoard, true
}

// HasComputeBoard returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasComputeBoard() bool {
	if o != nil && o.ComputeBoard != nil {
		return true
	}

	return false
}

// SetComputeBoard gets a reference to the given ComputeBoardRelationship and assigns it to the ComputeBoard field.
func (o *StorageFlexFlashController) SetComputeBoard(v ComputeBoardRelationship) {
	o.ComputeBoard = &v
}

// GetFlexFlashControllerProps returns the FlexFlashControllerProps field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashController) GetFlexFlashControllerProps() []StorageFlexFlashControllerPropsRelationship {
	if o == nil {
		var ret []StorageFlexFlashControllerPropsRelationship
		return ret
	}
	return o.FlexFlashControllerProps
}

// GetFlexFlashControllerPropsOk returns a tuple with the FlexFlashControllerProps field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashController) GetFlexFlashControllerPropsOk() (*[]StorageFlexFlashControllerPropsRelationship, bool) {
	if o == nil || o.FlexFlashControllerProps == nil {
		return nil, false
	}
	return &o.FlexFlashControllerProps, true
}

// HasFlexFlashControllerProps returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasFlexFlashControllerProps() bool {
	if o != nil && o.FlexFlashControllerProps != nil {
		return true
	}

	return false
}

// SetFlexFlashControllerProps gets a reference to the given []StorageFlexFlashControllerPropsRelationship and assigns it to the FlexFlashControllerProps field.
func (o *StorageFlexFlashController) SetFlexFlashControllerProps(v []StorageFlexFlashControllerPropsRelationship) {
	o.FlexFlashControllerProps = v
}

// GetFlexFlashPhysicalDrives returns the FlexFlashPhysicalDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashController) GetFlexFlashPhysicalDrives() []StorageFlexFlashPhysicalDriveRelationship {
	if o == nil {
		var ret []StorageFlexFlashPhysicalDriveRelationship
		return ret
	}
	return o.FlexFlashPhysicalDrives
}

// GetFlexFlashPhysicalDrivesOk returns a tuple with the FlexFlashPhysicalDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashController) GetFlexFlashPhysicalDrivesOk() (*[]StorageFlexFlashPhysicalDriveRelationship, bool) {
	if o == nil || o.FlexFlashPhysicalDrives == nil {
		return nil, false
	}
	return &o.FlexFlashPhysicalDrives, true
}

// HasFlexFlashPhysicalDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasFlexFlashPhysicalDrives() bool {
	if o != nil && o.FlexFlashPhysicalDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashPhysicalDrives gets a reference to the given []StorageFlexFlashPhysicalDriveRelationship and assigns it to the FlexFlashPhysicalDrives field.
func (o *StorageFlexFlashController) SetFlexFlashPhysicalDrives(v []StorageFlexFlashPhysicalDriveRelationship) {
	o.FlexFlashPhysicalDrives = v
}

// GetFlexFlashVirtualDrives returns the FlexFlashVirtualDrives field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashController) GetFlexFlashVirtualDrives() []StorageFlexFlashVirtualDriveRelationship {
	if o == nil {
		var ret []StorageFlexFlashVirtualDriveRelationship
		return ret
	}
	return o.FlexFlashVirtualDrives
}

// GetFlexFlashVirtualDrivesOk returns a tuple with the FlexFlashVirtualDrives field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashController) GetFlexFlashVirtualDrivesOk() (*[]StorageFlexFlashVirtualDriveRelationship, bool) {
	if o == nil || o.FlexFlashVirtualDrives == nil {
		return nil, false
	}
	return &o.FlexFlashVirtualDrives, true
}

// HasFlexFlashVirtualDrives returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasFlexFlashVirtualDrives() bool {
	if o != nil && o.FlexFlashVirtualDrives != nil {
		return true
	}

	return false
}

// SetFlexFlashVirtualDrives gets a reference to the given []StorageFlexFlashVirtualDriveRelationship and assigns it to the FlexFlashVirtualDrives field.
func (o *StorageFlexFlashController) SetFlexFlashVirtualDrives(v []StorageFlexFlashVirtualDriveRelationship) {
	o.FlexFlashVirtualDrives = v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *StorageFlexFlashController) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *StorageFlexFlashController) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *StorageFlexFlashController) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageFlexFlashController) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *StorageFlexFlashController) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

// GetRunningFirmware returns the RunningFirmware field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *StorageFlexFlashController) GetRunningFirmware() []FirmwareRunningFirmwareRelationship {
	if o == nil {
		var ret []FirmwareRunningFirmwareRelationship
		return ret
	}
	return o.RunningFirmware
}

// GetRunningFirmwareOk returns a tuple with the RunningFirmware field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *StorageFlexFlashController) GetRunningFirmwareOk() (*[]FirmwareRunningFirmwareRelationship, bool) {
	if o == nil || o.RunningFirmware == nil {
		return nil, false
	}
	return &o.RunningFirmware, true
}

// HasRunningFirmware returns a boolean if a field has been set.
func (o *StorageFlexFlashController) HasRunningFirmware() bool {
	if o != nil && o.RunningFirmware != nil {
		return true
	}

	return false
}

// SetRunningFirmware gets a reference to the given []FirmwareRunningFirmwareRelationship and assigns it to the RunningFirmware field.
func (o *StorageFlexFlashController) SetRunningFirmware(v []FirmwareRunningFirmwareRelationship) {
	o.RunningFirmware = v
}

func (o StorageFlexFlashController) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.ControllerState != nil {
		toSerialize["ControllerState"] = o.ControllerState
	}
	if o.FfControllerId != nil {
		toSerialize["FfControllerId"] = o.FfControllerId
	}
	if o.ComputeBoard != nil {
		toSerialize["ComputeBoard"] = o.ComputeBoard
	}
	if o.FlexFlashControllerProps != nil {
		toSerialize["FlexFlashControllerProps"] = o.FlexFlashControllerProps
	}
	if o.FlexFlashPhysicalDrives != nil {
		toSerialize["FlexFlashPhysicalDrives"] = o.FlexFlashPhysicalDrives
	}
	if o.FlexFlashVirtualDrives != nil {
		toSerialize["FlexFlashVirtualDrives"] = o.FlexFlashVirtualDrives
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}
	if o.RunningFirmware != nil {
		toSerialize["RunningFirmware"] = o.RunningFirmware
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageFlexFlashController) UnmarshalJSON(bytes []byte) (err error) {
	type StorageFlexFlashControllerWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// State of the Flex Flash Storage Controller.
		ControllerState *string `json:"ControllerState,omitempty"`
		// Identifier for the Flex Flash Storage Controller.
		FfControllerId *string                   `json:"FfControllerId,omitempty"`
		ComputeBoard   *ComputeBoardRelationship `json:"ComputeBoard,omitempty"`
		// An array of relationships to storageFlexFlashControllerProps resources.
		FlexFlashControllerProps []StorageFlexFlashControllerPropsRelationship `json:"FlexFlashControllerProps,omitempty"`
		// An array of relationships to storageFlexFlashPhysicalDrive resources.
		FlexFlashPhysicalDrives []StorageFlexFlashPhysicalDriveRelationship `json:"FlexFlashPhysicalDrives,omitempty"`
		// An array of relationships to storageFlexFlashVirtualDrive resources.
		FlexFlashVirtualDrives []StorageFlexFlashVirtualDriveRelationship `json:"FlexFlashVirtualDrives,omitempty"`
		InventoryDeviceInfo    *InventoryDeviceInfoRelationship           `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice       *AssetDeviceRegistrationRelationship       `json:"RegisteredDevice,omitempty"`
		// An array of relationships to firmwareRunningFirmware resources.
		RunningFirmware []FirmwareRunningFirmwareRelationship `json:"RunningFirmware,omitempty"`
	}

	varStorageFlexFlashControllerWithoutEmbeddedStruct := StorageFlexFlashControllerWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageFlexFlashControllerWithoutEmbeddedStruct)
	if err == nil {
		varStorageFlexFlashController := _StorageFlexFlashController{}
		varStorageFlexFlashController.ClassId = varStorageFlexFlashControllerWithoutEmbeddedStruct.ClassId
		varStorageFlexFlashController.ObjectType = varStorageFlexFlashControllerWithoutEmbeddedStruct.ObjectType
		varStorageFlexFlashController.ControllerState = varStorageFlexFlashControllerWithoutEmbeddedStruct.ControllerState
		varStorageFlexFlashController.FfControllerId = varStorageFlexFlashControllerWithoutEmbeddedStruct.FfControllerId
		varStorageFlexFlashController.ComputeBoard = varStorageFlexFlashControllerWithoutEmbeddedStruct.ComputeBoard
		varStorageFlexFlashController.FlexFlashControllerProps = varStorageFlexFlashControllerWithoutEmbeddedStruct.FlexFlashControllerProps
		varStorageFlexFlashController.FlexFlashPhysicalDrives = varStorageFlexFlashControllerWithoutEmbeddedStruct.FlexFlashPhysicalDrives
		varStorageFlexFlashController.FlexFlashVirtualDrives = varStorageFlexFlashControllerWithoutEmbeddedStruct.FlexFlashVirtualDrives
		varStorageFlexFlashController.InventoryDeviceInfo = varStorageFlexFlashControllerWithoutEmbeddedStruct.InventoryDeviceInfo
		varStorageFlexFlashController.RegisteredDevice = varStorageFlexFlashControllerWithoutEmbeddedStruct.RegisteredDevice
		varStorageFlexFlashController.RunningFirmware = varStorageFlexFlashControllerWithoutEmbeddedStruct.RunningFirmware
		*o = StorageFlexFlashController(varStorageFlexFlashController)
	} else {
		return err
	}

	varStorageFlexFlashController := _StorageFlexFlashController{}

	err = json.Unmarshal(bytes, &varStorageFlexFlashController)
	if err == nil {
		o.EquipmentBase = varStorageFlexFlashController.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "ControllerState")
		delete(additionalProperties, "FfControllerId")
		delete(additionalProperties, "ComputeBoard")
		delete(additionalProperties, "FlexFlashControllerProps")
		delete(additionalProperties, "FlexFlashPhysicalDrives")
		delete(additionalProperties, "FlexFlashVirtualDrives")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")
		delete(additionalProperties, "RunningFirmware")

		// remove fields from embedded structs
		reflectEquipmentBase := reflect.ValueOf(o.EquipmentBase)
		for i := 0; i < reflectEquipmentBase.Type().NumField(); i++ {
			t := reflectEquipmentBase.Type().Field(i)

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

type NullableStorageFlexFlashController struct {
	value *StorageFlexFlashController
	isSet bool
}

func (v NullableStorageFlexFlashController) Get() *StorageFlexFlashController {
	return v.value
}

func (v *NullableStorageFlexFlashController) Set(val *StorageFlexFlashController) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageFlexFlashController) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageFlexFlashController) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageFlexFlashController(val *StorageFlexFlashController) *NullableStorageFlexFlashController {
	return &NullableStorageFlexFlashController{value: val, isSet: true}
}

func (v NullableStorageFlexFlashController) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageFlexFlashController) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
