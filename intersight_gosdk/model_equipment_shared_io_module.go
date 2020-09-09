/*
 * Cisco Intersight
 *
 * Cisco Intersight is a management platform delivered as a service with embedded analytics for your Cisco and 3rd party IT infrastructure. This platform offers an intelligent level of management that enables IT organizations to analyze, simplify, and automate their environments in more advanced ways than the prior generations of tools. Cisco Intersight provides an integrated and intuitive management experience for resources in the traditional data center as well as at the edge. With flexible deployment options to address complex security needs, getting started with Intersight is quick and easy. Cisco Intersight has deep integration with Cisco UCS and HyperFlex systems allowing for remote deployment, configuration, and ongoing maintenance. The model-based deployment works for a single system in a remote location or hundreds of systems in a data center and enables rapid, standardized configuration and deployment. It also streamlines maintaining those systems whether you are working with small or very large configurations. The Intersight OpenAPI document defines the complete set of properties that are returned in the HTTP response. From that perspective, a client can expect that no additional properties are returned, unless these properties are explicitly defined in the OpenAPI document. However, when a client uses an older version of the Intersight OpenAPI document, the server may send additional properties because the software is more recent than the client. In that case, the client may receive properties that it does not know about. Some generated SDKs perform a strict validation of the HTTP response body against the OpenAPI document. This document was created on 2020-07-31T04:35:53Z.
 *
 * API version: 1.0.9-2110
 * Contact: intersight@cisco.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package intersight

import (
	"encoding/json"
	"reflect"
	"strings"
)

// EquipmentSharedIoModule I/O Controller present inside SIOC to provide data path from S-series server to FI.
type EquipmentSharedIoModule struct {
	EquipmentBase
	// This field identifies the configuration state for this SIOM Unit.
	ConfigState *string `json:"ConfigState,omitempty"`
	// This field identifies the discovery state of SIOM.
	Discovery *string `json:"Discovery,omitempty"`
	// This field identifies the MAC of IOM-A side.
	MacOfSharedIomAside *string `json:"MacOfSharedIomAside,omitempty"`
	// This field identifies the MAC of IOM-B side.
	MacOfSharedIomBside *string `json:"MacOfSharedIomBside,omitempty"`
	// This field identifies the SIOM operational state.
	OperState *string `json:"OperState,omitempty"`
	// This field identifies the Part Number for this SIOM Unit.
	PartNumber *string `json:"PartNumber,omitempty"`
	// This field identifies the reachability to FI-A and B side.
	Reachability *string `json:"Reachability,omitempty"`
	// User label configured for the SIOM.
	UsrLbl *string `json:"UsrLbl,omitempty"`
	// This field identifies the vendor id for this SIOM Unit.
	Vid                         *string                                  `json:"Vid,omitempty"`
	Controller                  *ManagementControllerRelationship        `json:"Controller,omitempty"`
	EquipmentSystemIoController *EquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
	InventoryDeviceInfo         *InventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
	// An array of relationships to portGroup resources.
	PortGroups           []PortGroupRelationship              `json:"PortGroups,omitempty"`
	RegisteredDevice     *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _EquipmentSharedIoModule EquipmentSharedIoModule

// NewEquipmentSharedIoModule instantiates a new EquipmentSharedIoModule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEquipmentSharedIoModule() *EquipmentSharedIoModule {
	this := EquipmentSharedIoModule{}
	return &this
}

// NewEquipmentSharedIoModuleWithDefaults instantiates a new EquipmentSharedIoModule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEquipmentSharedIoModuleWithDefaults() *EquipmentSharedIoModule {
	this := EquipmentSharedIoModule{}
	return &this
}

// GetConfigState returns the ConfigState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetConfigState() string {
	if o == nil || o.ConfigState == nil {
		var ret string
		return ret
	}
	return *o.ConfigState
}

// GetConfigStateOk returns a tuple with the ConfigState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetConfigStateOk() (*string, bool) {
	if o == nil || o.ConfigState == nil {
		return nil, false
	}
	return o.ConfigState, true
}

// HasConfigState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasConfigState() bool {
	if o != nil && o.ConfigState != nil {
		return true
	}

	return false
}

// SetConfigState gets a reference to the given string and assigns it to the ConfigState field.
func (o *EquipmentSharedIoModule) SetConfigState(v string) {
	o.ConfigState = &v
}

// GetDiscovery returns the Discovery field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetDiscovery() string {
	if o == nil || o.Discovery == nil {
		var ret string
		return ret
	}
	return *o.Discovery
}

// GetDiscoveryOk returns a tuple with the Discovery field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetDiscoveryOk() (*string, bool) {
	if o == nil || o.Discovery == nil {
		return nil, false
	}
	return o.Discovery, true
}

// HasDiscovery returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasDiscovery() bool {
	if o != nil && o.Discovery != nil {
		return true
	}

	return false
}

// SetDiscovery gets a reference to the given string and assigns it to the Discovery field.
func (o *EquipmentSharedIoModule) SetDiscovery(v string) {
	o.Discovery = &v
}

// GetMacOfSharedIomAside returns the MacOfSharedIomAside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomAside() string {
	if o == nil || o.MacOfSharedIomAside == nil {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomAside
}

// GetMacOfSharedIomAsideOk returns a tuple with the MacOfSharedIomAside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomAsideOk() (*string, bool) {
	if o == nil || o.MacOfSharedIomAside == nil {
		return nil, false
	}
	return o.MacOfSharedIomAside, true
}

// HasMacOfSharedIomAside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasMacOfSharedIomAside() bool {
	if o != nil && o.MacOfSharedIomAside != nil {
		return true
	}

	return false
}

// SetMacOfSharedIomAside gets a reference to the given string and assigns it to the MacOfSharedIomAside field.
func (o *EquipmentSharedIoModule) SetMacOfSharedIomAside(v string) {
	o.MacOfSharedIomAside = &v
}

// GetMacOfSharedIomBside returns the MacOfSharedIomBside field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomBside() string {
	if o == nil || o.MacOfSharedIomBside == nil {
		var ret string
		return ret
	}
	return *o.MacOfSharedIomBside
}

// GetMacOfSharedIomBsideOk returns a tuple with the MacOfSharedIomBside field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetMacOfSharedIomBsideOk() (*string, bool) {
	if o == nil || o.MacOfSharedIomBside == nil {
		return nil, false
	}
	return o.MacOfSharedIomBside, true
}

// HasMacOfSharedIomBside returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasMacOfSharedIomBside() bool {
	if o != nil && o.MacOfSharedIomBside != nil {
		return true
	}

	return false
}

// SetMacOfSharedIomBside gets a reference to the given string and assigns it to the MacOfSharedIomBside field.
func (o *EquipmentSharedIoModule) SetMacOfSharedIomBside(v string) {
	o.MacOfSharedIomBside = &v
}

// GetOperState returns the OperState field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetOperState() string {
	if o == nil || o.OperState == nil {
		var ret string
		return ret
	}
	return *o.OperState
}

// GetOperStateOk returns a tuple with the OperState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetOperStateOk() (*string, bool) {
	if o == nil || o.OperState == nil {
		return nil, false
	}
	return o.OperState, true
}

// HasOperState returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasOperState() bool {
	if o != nil && o.OperState != nil {
		return true
	}

	return false
}

// SetOperState gets a reference to the given string and assigns it to the OperState field.
func (o *EquipmentSharedIoModule) SetOperState(v string) {
	o.OperState = &v
}

// GetPartNumber returns the PartNumber field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetPartNumber() string {
	if o == nil || o.PartNumber == nil {
		var ret string
		return ret
	}
	return *o.PartNumber
}

// GetPartNumberOk returns a tuple with the PartNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetPartNumberOk() (*string, bool) {
	if o == nil || o.PartNumber == nil {
		return nil, false
	}
	return o.PartNumber, true
}

// HasPartNumber returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasPartNumber() bool {
	if o != nil && o.PartNumber != nil {
		return true
	}

	return false
}

// SetPartNumber gets a reference to the given string and assigns it to the PartNumber field.
func (o *EquipmentSharedIoModule) SetPartNumber(v string) {
	o.PartNumber = &v
}

// GetReachability returns the Reachability field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetReachability() string {
	if o == nil || o.Reachability == nil {
		var ret string
		return ret
	}
	return *o.Reachability
}

// GetReachabilityOk returns a tuple with the Reachability field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetReachabilityOk() (*string, bool) {
	if o == nil || o.Reachability == nil {
		return nil, false
	}
	return o.Reachability, true
}

// HasReachability returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasReachability() bool {
	if o != nil && o.Reachability != nil {
		return true
	}

	return false
}

// SetReachability gets a reference to the given string and assigns it to the Reachability field.
func (o *EquipmentSharedIoModule) SetReachability(v string) {
	o.Reachability = &v
}

// GetUsrLbl returns the UsrLbl field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetUsrLbl() string {
	if o == nil || o.UsrLbl == nil {
		var ret string
		return ret
	}
	return *o.UsrLbl
}

// GetUsrLblOk returns a tuple with the UsrLbl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetUsrLblOk() (*string, bool) {
	if o == nil || o.UsrLbl == nil {
		return nil, false
	}
	return o.UsrLbl, true
}

// HasUsrLbl returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasUsrLbl() bool {
	if o != nil && o.UsrLbl != nil {
		return true
	}

	return false
}

// SetUsrLbl gets a reference to the given string and assigns it to the UsrLbl field.
func (o *EquipmentSharedIoModule) SetUsrLbl(v string) {
	o.UsrLbl = &v
}

// GetVid returns the Vid field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetVid() string {
	if o == nil || o.Vid == nil {
		var ret string
		return ret
	}
	return *o.Vid
}

// GetVidOk returns a tuple with the Vid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetVidOk() (*string, bool) {
	if o == nil || o.Vid == nil {
		return nil, false
	}
	return o.Vid, true
}

// HasVid returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasVid() bool {
	if o != nil && o.Vid != nil {
		return true
	}

	return false
}

// SetVid gets a reference to the given string and assigns it to the Vid field.
func (o *EquipmentSharedIoModule) SetVid(v string) {
	o.Vid = &v
}

// GetController returns the Controller field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetController() ManagementControllerRelationship {
	if o == nil || o.Controller == nil {
		var ret ManagementControllerRelationship
		return ret
	}
	return *o.Controller
}

// GetControllerOk returns a tuple with the Controller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetControllerOk() (*ManagementControllerRelationship, bool) {
	if o == nil || o.Controller == nil {
		return nil, false
	}
	return o.Controller, true
}

// HasController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasController() bool {
	if o != nil && o.Controller != nil {
		return true
	}

	return false
}

// SetController gets a reference to the given ManagementControllerRelationship and assigns it to the Controller field.
func (o *EquipmentSharedIoModule) SetController(v ManagementControllerRelationship) {
	o.Controller = &v
}

// GetEquipmentSystemIoController returns the EquipmentSystemIoController field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetEquipmentSystemIoController() EquipmentSystemIoControllerRelationship {
	if o == nil || o.EquipmentSystemIoController == nil {
		var ret EquipmentSystemIoControllerRelationship
		return ret
	}
	return *o.EquipmentSystemIoController
}

// GetEquipmentSystemIoControllerOk returns a tuple with the EquipmentSystemIoController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetEquipmentSystemIoControllerOk() (*EquipmentSystemIoControllerRelationship, bool) {
	if o == nil || o.EquipmentSystemIoController == nil {
		return nil, false
	}
	return o.EquipmentSystemIoController, true
}

// HasEquipmentSystemIoController returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasEquipmentSystemIoController() bool {
	if o != nil && o.EquipmentSystemIoController != nil {
		return true
	}

	return false
}

// SetEquipmentSystemIoController gets a reference to the given EquipmentSystemIoControllerRelationship and assigns it to the EquipmentSystemIoController field.
func (o *EquipmentSharedIoModule) SetEquipmentSystemIoController(v EquipmentSystemIoControllerRelationship) {
	o.EquipmentSystemIoController = &v
}

// GetInventoryDeviceInfo returns the InventoryDeviceInfo field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetInventoryDeviceInfo() InventoryDeviceInfoRelationship {
	if o == nil || o.InventoryDeviceInfo == nil {
		var ret InventoryDeviceInfoRelationship
		return ret
	}
	return *o.InventoryDeviceInfo
}

// GetInventoryDeviceInfoOk returns a tuple with the InventoryDeviceInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetInventoryDeviceInfoOk() (*InventoryDeviceInfoRelationship, bool) {
	if o == nil || o.InventoryDeviceInfo == nil {
		return nil, false
	}
	return o.InventoryDeviceInfo, true
}

// HasInventoryDeviceInfo returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasInventoryDeviceInfo() bool {
	if o != nil && o.InventoryDeviceInfo != nil {
		return true
	}

	return false
}

// SetInventoryDeviceInfo gets a reference to the given InventoryDeviceInfoRelationship and assigns it to the InventoryDeviceInfo field.
func (o *EquipmentSharedIoModule) SetInventoryDeviceInfo(v InventoryDeviceInfoRelationship) {
	o.InventoryDeviceInfo = &v
}

// GetPortGroups returns the PortGroups field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *EquipmentSharedIoModule) GetPortGroups() []PortGroupRelationship {
	if o == nil {
		var ret []PortGroupRelationship
		return ret
	}
	return o.PortGroups
}

// GetPortGroupsOk returns a tuple with the PortGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *EquipmentSharedIoModule) GetPortGroupsOk() (*[]PortGroupRelationship, bool) {
	if o == nil || o.PortGroups == nil {
		return nil, false
	}
	return &o.PortGroups, true
}

// HasPortGroups returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasPortGroups() bool {
	if o != nil && o.PortGroups != nil {
		return true
	}

	return false
}

// SetPortGroups gets a reference to the given []PortGroupRelationship and assigns it to the PortGroups field.
func (o *EquipmentSharedIoModule) SetPortGroups(v []PortGroupRelationship) {
	o.PortGroups = v
}

// GetRegisteredDevice returns the RegisteredDevice field value if set, zero value otherwise.
func (o *EquipmentSharedIoModule) GetRegisteredDevice() AssetDeviceRegistrationRelationship {
	if o == nil || o.RegisteredDevice == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.RegisteredDevice
}

// GetRegisteredDeviceOk returns a tuple with the RegisteredDevice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EquipmentSharedIoModule) GetRegisteredDeviceOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.RegisteredDevice == nil {
		return nil, false
	}
	return o.RegisteredDevice, true
}

// HasRegisteredDevice returns a boolean if a field has been set.
func (o *EquipmentSharedIoModule) HasRegisteredDevice() bool {
	if o != nil && o.RegisteredDevice != nil {
		return true
	}

	return false
}

// SetRegisteredDevice gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the RegisteredDevice field.
func (o *EquipmentSharedIoModule) SetRegisteredDevice(v AssetDeviceRegistrationRelationship) {
	o.RegisteredDevice = &v
}

func (o EquipmentSharedIoModule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedEquipmentBase, errEquipmentBase := json.Marshal(o.EquipmentBase)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	errEquipmentBase = json.Unmarshal([]byte(serializedEquipmentBase), &toSerialize)
	if errEquipmentBase != nil {
		return []byte{}, errEquipmentBase
	}
	if o.ConfigState != nil {
		toSerialize["ConfigState"] = o.ConfigState
	}
	if o.Discovery != nil {
		toSerialize["Discovery"] = o.Discovery
	}
	if o.MacOfSharedIomAside != nil {
		toSerialize["MacOfSharedIomAside"] = o.MacOfSharedIomAside
	}
	if o.MacOfSharedIomBside != nil {
		toSerialize["MacOfSharedIomBside"] = o.MacOfSharedIomBside
	}
	if o.OperState != nil {
		toSerialize["OperState"] = o.OperState
	}
	if o.PartNumber != nil {
		toSerialize["PartNumber"] = o.PartNumber
	}
	if o.Reachability != nil {
		toSerialize["Reachability"] = o.Reachability
	}
	if o.UsrLbl != nil {
		toSerialize["UsrLbl"] = o.UsrLbl
	}
	if o.Vid != nil {
		toSerialize["Vid"] = o.Vid
	}
	if o.Controller != nil {
		toSerialize["Controller"] = o.Controller
	}
	if o.EquipmentSystemIoController != nil {
		toSerialize["EquipmentSystemIoController"] = o.EquipmentSystemIoController
	}
	if o.InventoryDeviceInfo != nil {
		toSerialize["InventoryDeviceInfo"] = o.InventoryDeviceInfo
	}
	if o.PortGroups != nil {
		toSerialize["PortGroups"] = o.PortGroups
	}
	if o.RegisteredDevice != nil {
		toSerialize["RegisteredDevice"] = o.RegisteredDevice
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *EquipmentSharedIoModule) UnmarshalJSON(bytes []byte) (err error) {
	type EquipmentSharedIoModuleWithoutEmbeddedStruct struct {
		// This field identifies the configuration state for this SIOM Unit.
		ConfigState *string `json:"ConfigState,omitempty"`
		// This field identifies the discovery state of SIOM.
		Discovery *string `json:"Discovery,omitempty"`
		// This field identifies the MAC of IOM-A side.
		MacOfSharedIomAside *string `json:"MacOfSharedIomAside,omitempty"`
		// This field identifies the MAC of IOM-B side.
		MacOfSharedIomBside *string `json:"MacOfSharedIomBside,omitempty"`
		// This field identifies the SIOM operational state.
		OperState *string `json:"OperState,omitempty"`
		// This field identifies the Part Number for this SIOM Unit.
		PartNumber *string `json:"PartNumber,omitempty"`
		// This field identifies the reachability to FI-A and B side.
		Reachability *string `json:"Reachability,omitempty"`
		// User label configured for the SIOM.
		UsrLbl *string `json:"UsrLbl,omitempty"`
		// This field identifies the vendor id for this SIOM Unit.
		Vid                         *string                                  `json:"Vid,omitempty"`
		Controller                  *ManagementControllerRelationship        `json:"Controller,omitempty"`
		EquipmentSystemIoController *EquipmentSystemIoControllerRelationship `json:"EquipmentSystemIoController,omitempty"`
		InventoryDeviceInfo         *InventoryDeviceInfoRelationship         `json:"InventoryDeviceInfo,omitempty"`
		// An array of relationships to portGroup resources.
		PortGroups       []PortGroupRelationship              `json:"PortGroups,omitempty"`
		RegisteredDevice *AssetDeviceRegistrationRelationship `json:"RegisteredDevice,omitempty"`
	}

	varEquipmentSharedIoModuleWithoutEmbeddedStruct := EquipmentSharedIoModuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varEquipmentSharedIoModuleWithoutEmbeddedStruct)
	if err == nil {
		varEquipmentSharedIoModule := _EquipmentSharedIoModule{}
		varEquipmentSharedIoModule.ConfigState = varEquipmentSharedIoModuleWithoutEmbeddedStruct.ConfigState
		varEquipmentSharedIoModule.Discovery = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Discovery
		varEquipmentSharedIoModule.MacOfSharedIomAside = varEquipmentSharedIoModuleWithoutEmbeddedStruct.MacOfSharedIomAside
		varEquipmentSharedIoModule.MacOfSharedIomBside = varEquipmentSharedIoModuleWithoutEmbeddedStruct.MacOfSharedIomBside
		varEquipmentSharedIoModule.OperState = varEquipmentSharedIoModuleWithoutEmbeddedStruct.OperState
		varEquipmentSharedIoModule.PartNumber = varEquipmentSharedIoModuleWithoutEmbeddedStruct.PartNumber
		varEquipmentSharedIoModule.Reachability = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Reachability
		varEquipmentSharedIoModule.UsrLbl = varEquipmentSharedIoModuleWithoutEmbeddedStruct.UsrLbl
		varEquipmentSharedIoModule.Vid = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Vid
		varEquipmentSharedIoModule.Controller = varEquipmentSharedIoModuleWithoutEmbeddedStruct.Controller
		varEquipmentSharedIoModule.EquipmentSystemIoController = varEquipmentSharedIoModuleWithoutEmbeddedStruct.EquipmentSystemIoController
		varEquipmentSharedIoModule.InventoryDeviceInfo = varEquipmentSharedIoModuleWithoutEmbeddedStruct.InventoryDeviceInfo
		varEquipmentSharedIoModule.PortGroups = varEquipmentSharedIoModuleWithoutEmbeddedStruct.PortGroups
		varEquipmentSharedIoModule.RegisteredDevice = varEquipmentSharedIoModuleWithoutEmbeddedStruct.RegisteredDevice
		*o = EquipmentSharedIoModule(varEquipmentSharedIoModule)
	} else {
		return err
	}

	varEquipmentSharedIoModule := _EquipmentSharedIoModule{}

	err = json.Unmarshal(bytes, &varEquipmentSharedIoModule)
	if err == nil {
		o.EquipmentBase = varEquipmentSharedIoModule.EquipmentBase
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ConfigState")
		delete(additionalProperties, "Discovery")
		delete(additionalProperties, "MacOfSharedIomAside")
		delete(additionalProperties, "MacOfSharedIomBside")
		delete(additionalProperties, "OperState")
		delete(additionalProperties, "PartNumber")
		delete(additionalProperties, "Reachability")
		delete(additionalProperties, "UsrLbl")
		delete(additionalProperties, "Vid")
		delete(additionalProperties, "Controller")
		delete(additionalProperties, "EquipmentSystemIoController")
		delete(additionalProperties, "InventoryDeviceInfo")
		delete(additionalProperties, "PortGroups")
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

type NullableEquipmentSharedIoModule struct {
	value *EquipmentSharedIoModule
	isSet bool
}

func (v NullableEquipmentSharedIoModule) Get() *EquipmentSharedIoModule {
	return v.value
}

func (v *NullableEquipmentSharedIoModule) Set(val *EquipmentSharedIoModule) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentSharedIoModule) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentSharedIoModule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentSharedIoModule(val *EquipmentSharedIoModule) *NullableEquipmentSharedIoModule {
	return &NullableEquipmentSharedIoModule{value: val, isSet: true}
}

func (v NullableEquipmentSharedIoModule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentSharedIoModule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}