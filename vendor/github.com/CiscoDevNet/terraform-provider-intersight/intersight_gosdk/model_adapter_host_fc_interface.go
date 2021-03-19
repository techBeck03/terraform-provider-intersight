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
	"reflect"
	"strings"
)

// AdapterHostFcInterface Host facing fibre channel interface on a server adapter.
type AdapterHostFcInterface struct {
	EquipmentBase
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Admin Configured State of Host Fibre Channel Interface.
	AdminState *string `json:"AdminState,omitempty"`
	// The Endpoint Config Dn of the Host Fibre Channel Interface.
	EpDn *string `json:"EpDn,omitempty"`
	// Identifier of Host Fibre Channel Interface.
	HostFcInterfaceId *int64 `json:"HostFcInterfaceId,omitempty"`
	// Name of Host Fibre Channel Interface.
	Name       *string  `json:"Name,omitempty"`
	OperReason []string `json:"OperReason,omitempty"`
	// Operational State of Host Fibre Channel Interface.
	OperState *string `json:"OperState,omitempty"`
	// Operability status of Host Fibre Channel Interface.
	Operability *string `json:"Operability,omitempty"`
	// The uniquely distinguishable factory default  World Wide Node Name of the Host.
	OriginalWwnn *string `json:"OriginalWwnn,omitempty"`
	// The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface.
	OriginalWwpn *string `json:"OriginalWwpn,omitempty"`
	// PeerPort Dn of Host Fibre Channel Interface.
	PeerDn *string `json:"PeerDn,omitempty"`
	// The uniquely distinguishable user configured World Wide Node Name of the Host.
	Wwnn *string `json:"Wwnn,omitempty"`
	// The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface.
	Wwpn                 *string                              `json:"Wwpn,omitempty"`
	AdapterUnit          *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
	InventoryDeviceInfo  *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AdapterHostFcInterface AdapterHostFcInterface

// NewAdapterHostFcInterface instantiates a new AdapterHostFcInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAdapterHostFcInterface(classId string, objectType string) *AdapterHostFcInterface {
	this := AdapterHostFcInterface{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAdapterHostFcInterfaceWithDefaults instantiates a new AdapterHostFcInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAdapterHostFcInterfaceWithDefaults() *AdapterHostFcInterface {
	this := AdapterHostFcInterface{}
	var classId string = "adapter.HostFcInterface"
	this.ClassId = classId
	var objectType string = "adapter.HostFcInterface"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AdapterHostFcInterface) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AdapterHostFcInterface) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AdapterHostFcInterface) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AdapterHostFcInterface) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAdminState returns the AdminState field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetAdminState() string {
	if o == nil || o.AdminState == nil {
		var ret string
		return ret
	}
	return *o.AdminState
}

// GetAdminStateOk returns a tuple with the AdminState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetAdminStateOk() (*string, bool) {
	if o == nil || o.AdminState == nil {
		return nil, false
	}
	return o.AdminState, true
}

// HasAdminState returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasAdminState() bool {
	if o != nil && o.AdminState != nil {
		return true
	}

	return false
}

// SetAdminState gets a reference to the given string and assigns it to the AdminState field.
func (o *AdapterHostFcInterface) SetAdminState(v string) {
	o.AdminState = &v
}

// GetEpDn returns the EpDn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetEpDn() string {
	if o == nil || o.EpDn == nil {
		var ret string
		return ret
	}
	return *o.EpDn
}

// GetEpDnOk returns a tuple with the EpDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetEpDnOk() (*string, bool) {
	if o == nil || o.EpDn == nil {
		return nil, false
	}
	return o.EpDn, true
}

// HasEpDn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasEpDn() bool {
	if o != nil && o.EpDn != nil {
		return true
	}

	return false
}

// SetEpDn gets a reference to the given string and assigns it to the EpDn field.
func (o *AdapterHostFcInterface) SetEpDn(v string) {
	o.EpDn = &v
}

// GetHostFcInterfaceId returns the HostFcInterfaceId field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetHostFcInterfaceId() int64 {
	if o == nil || o.HostFcInterfaceId == nil {
		var ret int64
		return ret
	}
	return *o.HostFcInterfaceId
}

// GetHostFcInterfaceIdOk returns a tuple with the HostFcInterfaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetHostFcInterfaceIdOk() (*int64, bool) {
	if o == nil || o.HostFcInterfaceId == nil {
		return nil, false
	}
	return o.HostFcInterfaceId, true
}

// HasHostFcInterfaceId returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasHostFcInterfaceId() bool {
	if o != nil && o.HostFcInterfaceId != nil {
		return true
	}

	return false
}

// SetHostFcInterfaceId gets a reference to the given int64 and assigns it to the HostFcInterfaceId field.
func (o *AdapterHostFcInterface) SetHostFcInterfaceId(v int64) {
	o.HostFcInterfaceId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AdapterHostFcInterface) SetName(v string) {
	o.Name = &v
}

// GetOperReason returns the OperReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *AdapterHostFcInterface) GetOperReason() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.OperReason
}

// GetOperReasonOk returns a tuple with the OperReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AdapterHostFcInterface) GetOperReasonOk() (*[]string, bool) {
	if o == nil || o.OperReason == nil {
		return nil, false
	}
	return &o.OperReason, true
}

// HasOperReason returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasOperReason() bool {
	if o != nil && o.OperReason != nil {
		return true
	}

	return false
}

// SetOperReason gets a reference to the given []string and assigns it to the OperReason field.
func (o *AdapterHostFcInterface) SetOperReason(v []string) {
	o.OperReason = v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *AdapterHostFcInterface) SetOperState(v string) {
	o.OperState = &v
}

// GetOperability returns the Operability field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetOperability() string {
	if o == nil || o.Operability == nil {
		var ret string
		return ret
	}
	return *o.Operability
}

// GetOperabilityOk returns a tuple with the Operability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetOperabilityOk() (*string, bool) {
	if o == nil || o.Operability == nil {
		return nil, false
	}
	return o.Operability, true
}

// HasOperability returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasOperability() bool {
	if o != nil && o.Operability != nil {
		return true
	}

	return false
}

// SetOperability gets a reference to the given string and assigns it to the Operability field.
func (o *AdapterHostFcInterface) SetOperability(v string) {
	o.Operability = &v
}

// GetOriginalWwnn returns the OriginalWwnn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetOriginalWwnn() string {
	if o == nil || o.OriginalWwnn == nil {
		var ret string
		return ret
	}
	return *o.OriginalWwnn
}

// GetOriginalWwnnOk returns a tuple with the OriginalWwnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetOriginalWwnnOk() (*string, bool) {
	if o == nil || o.OriginalWwnn == nil {
		return nil, false
	}
	return o.OriginalWwnn, true
}

// HasOriginalWwnn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasOriginalWwnn() bool {
	if o != nil && o.OriginalWwnn != nil {
		return true
	}

	return false
}

// SetOriginalWwnn gets a reference to the given string and assigns it to the OriginalWwnn field.
func (o *AdapterHostFcInterface) SetOriginalWwnn(v string) {
	o.OriginalWwnn = &v
}

// GetOriginalWwpn returns the OriginalWwpn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetOriginalWwpn() string {
	if o == nil || o.OriginalWwpn == nil {
		var ret string
		return ret
	}
	return *o.OriginalWwpn
}

// GetOriginalWwpnOk returns a tuple with the OriginalWwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetOriginalWwpnOk() (*string, bool) {
	if o == nil || o.OriginalWwpn == nil {
		return nil, false
	}
	return o.OriginalWwpn, true
}

// HasOriginalWwpn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasOriginalWwpn() bool {
	if o != nil && o.OriginalWwpn != nil {
		return true
	}

	return false
}

// SetOriginalWwpn gets a reference to the given string and assigns it to the OriginalWwpn field.
func (o *AdapterHostFcInterface) SetOriginalWwpn(v string) {
	o.OriginalWwpn = &v
}

// GetPeerDn returns the PeerDn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetPeerDn() string {
	if o == nil || o.PeerDn == nil {
		var ret string
		return ret
	}
	return *o.PeerDn
}

// GetPeerDnOk returns a tuple with the PeerDn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetPeerDnOk() (*string, bool) {
	if o == nil || o.PeerDn == nil {
		return nil, false
	}
	return o.PeerDn, true
}

// HasPeerDn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasPeerDn() bool {
	if o != nil && o.PeerDn != nil {
		return true
	}

	return false
}

// SetPeerDn gets a reference to the given string and assigns it to the PeerDn field.
func (o *AdapterHostFcInterface) SetPeerDn(v string) {
	o.PeerDn = &v
}

// GetWwnn returns the Wwnn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetWwnn() string {
	if o == nil || o.Wwnn == nil {
		var ret string
		return ret
	}
	return *o.Wwnn
}

// GetWwnnOk returns a tuple with the Wwnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetWwnnOk() (*string, bool) {
	if o == nil || o.Wwnn == nil {
		return nil, false
	}
	return o.Wwnn, true
}

// HasWwnn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasWwnn() bool {
	if o != nil && o.Wwnn != nil {
		return true
	}

	return false
}

// SetWwnn gets a reference to the given string and assigns it to the Wwnn field.
func (o *AdapterHostFcInterface) SetWwnn(v string) {
	o.Wwnn = &v
}

// GetWwpn returns the Wwpn field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetWwpn() string {
	if o == nil || o.Wwpn == nil {
		var ret string
		return ret
	}
	return *o.Wwpn
}

// GetWwpnOk returns a tuple with the Wwpn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetWwpnOk() (*string, bool) {
	if o == nil || o.Wwpn == nil {
		return nil, false
	}
	return o.Wwpn, true
}

// HasWwpn returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasWwpn() bool {
	if o != nil && o.Wwpn != nil {
		return true
	}

	return false
}

// SetWwpn gets a reference to the given string and assigns it to the Wwpn field.
func (o *AdapterHostFcInterface) SetWwpn(v string) {
	o.Wwpn = &v
}

// GetAdapterUnit returns the AdapterUnit field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetAdapterUnit() AdapterUnitRelationship {
	if o == nil || o.AdapterUnit == nil {
		var ret AdapterUnitRelationship
		return ret
	}
	return *o.AdapterUnit
}

// GetAdapterUnitOk returns a tuple with the AdapterUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetAdapterUnitOk() (*AdapterUnitRelationship, bool) {
	if o == nil || o.AdapterUnit == nil {
		return nil, false
	}
	return o.AdapterUnit, true
}

// HasAdapterUnit returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasAdapterUnit() bool {
	if o != nil && o.AdapterUnit != nil {
		return true
	}

	return false
}

// SetAdapterUnit gets a reference to the given AdapterUnitRelationship and assigns it to the AdapterUnit field.
func (o *AdapterHostFcInterface) SetAdapterUnit(v AdapterUnitRelationship) {
	o.AdapterUnit = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *AdapterHostFcInterface) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *AdapterHostFcInterface) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AdapterHostFcInterface) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *AdapterHostFcInterface) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *AdapterHostFcInterface) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o AdapterHostFcInterface) MarshalJSON() ([]byte, error) {
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
	if o.AdminState != nil {
		toSerialize["AdminState"] = o.AdminState
	}
	if o.EpDn != nil {
		toSerialize["EpDn"] = o.EpDn
	}
	if o.HostFcInterfaceId != nil {
		toSerialize["HostFcInterfaceId"] = o.HostFcInterfaceId
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.OperReason != nil {
		toSerialize["OperReason"] = o.OperReason
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.Operability != nil {
		toSerialize["Operability"] = o.Operability
	}
	if o.OriginalWwnn != nil {
		toSerialize["OriginalWwnn"] = o.OriginalWwnn
	}
	if o.OriginalWwpn != nil {
		toSerialize["OriginalWwpn"] = o.OriginalWwpn
	}
	if o.PeerDn != nil {
		toSerialize["PeerDn"] = o.PeerDn
	}
	if o.Wwnn != nil {
		toSerialize["Wwnn"] = o.Wwnn
	}
	if o.Wwpn != nil {
		toSerialize["Wwpn"] = o.Wwpn
	}
	if o.AdapterUnit != nil {
		toSerialize["AdapterUnit"] = o.AdapterUnit
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

func (o *AdapterHostFcInterface) UnmarshalJSON(bytes []byte) (err error) {
	type AdapterHostFcInterfaceWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Admin Configured State of Host Fibre Channel Interface.
		AdminState *string `json:"AdminState,omitempty"`
		// The Endpoint Config Dn of the Host Fibre Channel Interface.
		EpDn *string `json:"EpDn,omitempty"`
		// Identifier of Host Fibre Channel Interface.
		HostFcInterfaceId *int64 `json:"HostFcInterfaceId,omitempty"`
		// Name of Host Fibre Channel Interface.
		Name       *string  `json:"Name,omitempty"`
		OperReason []string `json:"OperReason,omitempty"`
		// Operational State of Host Fibre Channel Interface.
		OperState *string `json:"OperState,omitempty"`
		// Operability status of Host Fibre Channel Interface.
		Operability *string `json:"Operability,omitempty"`
		// The uniquely distinguishable factory default  World Wide Node Name of the Host.
		OriginalWwnn *string `json:"OriginalWwnn,omitempty"`
		// The uniquely distinguishable factory default World Wide Port Name of the Host Fibre Channel Interface.
		OriginalWwpn *string `json:"OriginalWwpn,omitempty"`
		// PeerPort Dn of Host Fibre Channel Interface.
		PeerDn *string `json:"PeerDn,omitempty"`
		// The uniquely distinguishable user configured World Wide Node Name of the Host.
		Wwnn *string `json:"Wwnn,omitempty"`
		// The uniquely distinguishable user configured World Wide Port Name of the Host Fibre Channel Interface.
		Wwpn                *string                              `json:"Wwpn,omitempty"`
		AdapterUnit         *AdapterUnitRelationship             `json:"AdapterUnit,omitempty"`
		InventoryDeviceInfo *InventoryDeviceInfoRelationship     `json:"InventoryDeviceInfo,omitempty"`
		RegisteredDevice    *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varAdapterHostFcInterfaceWithoutEmbeddedStruct := AdapterHostFcInterfaceWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAdapterHostFcInterfaceWithoutEmbeddedStruct)
	if err == nil {
		varAdapterHostFcInterface := _AdapterHostFcInterface{}
		varAdapterHostFcInterface.ClassId = varAdapterHostFcInterfaceWithoutEmbeddedStruct.ClassId
		varAdapterHostFcInterface.ObjectType = varAdapterHostFcInterfaceWithoutEmbeddedStruct.ObjectType
		varAdapterHostFcInterface.AdminState = varAdapterHostFcInterfaceWithoutEmbeddedStruct.AdminState
		varAdapterHostFcInterface.EpDn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.EpDn
		varAdapterHostFcInterface.HostFcInterfaceId = varAdapterHostFcInterfaceWithoutEmbeddedStruct.HostFcInterfaceId
		varAdapterHostFcInterface.Name = varAdapterHostFcInterfaceWithoutEmbeddedStruct.Name
		varAdapterHostFcInterface.OperReason = varAdapterHostFcInterfaceWithoutEmbeddedStruct.OperReason
		varAdapterHostFcInterface.OperState = varAdapterHostFcInterfaceWithoutEmbeddedStruct.OperState
		varAdapterHostFcInterface.Operability = varAdapterHostFcInterfaceWithoutEmbeddedStruct.Operability
		varAdapterHostFcInterface.OriginalWwnn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.OriginalWwnn
		varAdapterHostFcInterface.OriginalWwpn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.OriginalWwpn
		varAdapterHostFcInterface.PeerDn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.PeerDn
		varAdapterHostFcInterface.Wwnn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.Wwnn
		varAdapterHostFcInterface.Wwpn = varAdapterHostFcInterfaceWithoutEmbeddedStruct.Wwpn
		varAdapterHostFcInterface.AdapterUnit = varAdapterHostFcInterfaceWithoutEmbeddedStruct.AdapterUnit
		varAdapterHostFcInterface.InventoryDeviceInfo = varAdapterHostFcInterfaceWithoutEmbeddedStruct.InventoryDeviceInfo
		varAdapterHostFcInterface.RegisteredDevice = varAdapterHostFcInterfaceWithoutEmbeddedStruct.RegisteredDevice
		*o = AdapterHostFcInterface(varAdapterHostFcInterface)
	} else {
		return err
	}

	varAdapterHostFcInterface := _AdapterHostFcInterface{}

	err = json.Unmarshal(bytes, &varAdapterHostFcInterface)
	if err == nil {
		o.EquipmentBase = varAdapterHostFcInterface.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AdminState")
		delete(additionalProperties, "EpDn")
		delete(additionalProperties, "HostFcInterfaceId")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "OperReason")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "Operability")
		delete(additionalProperties, "OriginalWwnn")
		delete(additionalProperties, "OriginalWwpn")
		delete(additionalProperties, "PeerDn")
		delete(additionalProperties, "Wwnn")
		delete(additionalProperties, "Wwpn")
		delete(additionalProperties, "AdapterUnit")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "RegisteredDevice")

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

type NullableAdapterHostFcInterface struct {
	value *AdapterHostFcInterface
	isSet bool
}

func (v NullableAdapterHostFcInterface) Get() *AdapterHostFcInterface {
	return v.value
}

func (v *NullableAdapterHostFcInterface) Set(val *AdapterHostFcInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableAdapterHostFcInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableAdapterHostFcInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdapterHostFcInterface(val *AdapterHostFcInterface) *NullableAdapterHostFcInterface {
	return &NullableAdapterHostFcInterface{value: val, isSet: true}
}

func (v NullableAdapterHostFcInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdapterHostFcInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
