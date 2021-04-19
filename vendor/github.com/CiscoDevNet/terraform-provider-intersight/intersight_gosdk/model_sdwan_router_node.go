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

// SdwanRouterNode Configuration settings for a SDWAN vEdge router.
type SdwanRouterNode struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration.
	DeviceTemplate *string `json:"DeviceTemplate,omitempty"`
	// Name of the router node object.
	Name                 *string                         `json:"Name,omitempty"`
	NetworkConfiguration []SdwanNetworkConfigurationType `json:"NetworkConfiguration,omitempty"`
	TemplateInputs       []SdwanTemplateInputsType       `json:"TemplateInputs,omitempty"`
	// Uniquely identifies the router by its chassis number.
	Uuid                 *string                               `json:"Uuid,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	Profile              *SdwanProfileRelationship             `json:"Profile,omitempty"`
	ServerNode           *AssetDeviceRegistrationRelationship  `json:"ServerNode,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SdwanRouterNode SdwanRouterNode

// NewSdwanRouterNode instantiates a new SdwanRouterNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSdwanRouterNode(classId string, objectType string) *SdwanRouterNode {
	this := SdwanRouterNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewSdwanRouterNodeWithDefaults instantiates a new SdwanRouterNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSdwanRouterNodeWithDefaults() *SdwanRouterNode {
	this := SdwanRouterNode{}
	var classId string = "sdwan.RouterNode"
	this.ClassId = classId
	var objectType string = "sdwan.RouterNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *SdwanRouterNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *SdwanRouterNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *SdwanRouterNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *SdwanRouterNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceTemplate returns the DeviceTemplate field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetDeviceTemplate() string {
	if o == nil || o.DeviceTemplate == nil {
		var ret string
		return ret
	}
	return *o.DeviceTemplate
}

// GetDeviceTemplateOk returns a tuple with the DeviceTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetDeviceTemplateOk() (*string, bool) {
	if o == nil || o.DeviceTemplate == nil {
		return nil, false
	}
	return o.DeviceTemplate, true
}

// HasDeviceTemplate returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasDeviceTemplate() bool {
	if o != nil && o.DeviceTemplate != nil {
		return true
	}

	return false
}

// SetDeviceTemplate gets a reference to the given string and assigns it to the DeviceTemplate field.
func (o *SdwanRouterNode) SetDeviceTemplate(v string) {
	o.DeviceTemplate = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SdwanRouterNode) SetName(v string) {
	o.Name = &v
}

// GetNetworkConfiguration returns the NetworkConfiguration field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdwanRouterNode) GetNetworkConfiguration() []SdwanNetworkConfigurationType {
	if o == nil {
		var ret []SdwanNetworkConfigurationType
		return ret
	}
	return o.NetworkConfiguration
}

// GetNetworkConfigurationOk returns a tuple with the NetworkConfiguration field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdwanRouterNode) GetNetworkConfigurationOk() (*[]SdwanNetworkConfigurationType, bool) {
	if o == nil || o.NetworkConfiguration == nil {
		return nil, false
	}
	return &o.NetworkConfiguration, true
}

// HasNetworkConfiguration returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasNetworkConfiguration() bool {
	if o != nil && o.NetworkConfiguration != nil {
		return true
	}

	return false
}

// SetNetworkConfiguration gets a reference to the given []SdwanNetworkConfigurationType and assigns it to the NetworkConfiguration field.
func (o *SdwanRouterNode) SetNetworkConfiguration(v []SdwanNetworkConfigurationType) {
	o.NetworkConfiguration = v
}

// GetTemplateInputs returns the TemplateInputs field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SdwanRouterNode) GetTemplateInputs() []SdwanTemplateInputsType {
	if o == nil {
		var ret []SdwanTemplateInputsType
		return ret
	}
	return o.TemplateInputs
}

// GetTemplateInputsOk returns a tuple with the TemplateInputs field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SdwanRouterNode) GetTemplateInputsOk() (*[]SdwanTemplateInputsType, bool) {
	if o == nil || o.TemplateInputs == nil {
		return nil, false
	}
	return &o.TemplateInputs, true
}

// HasTemplateInputs returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasTemplateInputs() bool {
	if o != nil && o.TemplateInputs != nil {
		return true
	}

	return false
}

// SetTemplateInputs gets a reference to the given []SdwanTemplateInputsType and assigns it to the TemplateInputs field.
func (o *SdwanRouterNode) SetTemplateInputs(v []SdwanTemplateInputsType) {
	o.TemplateInputs = v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SdwanRouterNode) SetUuid(v string) {
	o.Uuid = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *SdwanRouterNode) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetProfile() SdwanProfileRelationship {
	if o == nil || o.Profile == nil {
		var ret SdwanProfileRelationship
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetProfileOk() (*SdwanProfileRelationship, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given SdwanProfileRelationship and assigns it to the Profile field.
func (o *SdwanRouterNode) SetProfile(v SdwanProfileRelationship) {
	o.Profile = &v
}

// GetServerNode returns the ServerNode field value if set, zero value otherwise.
func (o *SdwanRouterNode) GetServerNode() AssetDeviceRegistrationRelationship {
	if o == nil || o.ServerNode == nil {
		var ret AssetDeviceRegistrationRelationship
		return ret
	}
	return *o.ServerNode
}

// GetServerNodeOk returns a tuple with the ServerNode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SdwanRouterNode) GetServerNodeOk() (*AssetDeviceRegistrationRelationship, bool) {
	if o == nil || o.ServerNode == nil {
		return nil, false
	}
	return o.ServerNode, true
}

// HasServerNode returns a boolean if a field has been set.
func (o *SdwanRouterNode) HasServerNode() bool {
	if o != nil && o.ServerNode != nil {
		return true
	}

	return false
}

// SetServerNode gets a reference to the given AssetDeviceRegistrationRelationship and assigns it to the ServerNode field.
func (o *SdwanRouterNode) SetServerNode(v AssetDeviceRegistrationRelationship) {
	o.ServerNode = &v
}

func (o SdwanRouterNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseMo, errMoBaseMo := json.Marshal(o.MoBaseMo)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	errMoBaseMo = json.Unmarshal([]byte(serializedMoBaseMo), &toSerialize)
	if errMoBaseMo != nil {
		return []byte{}, errMoBaseMo
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceTemplate != nil {
		toSerialize["DeviceTemplate"] = o.DeviceTemplate
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.NetworkConfiguration != nil {
		toSerialize["NetworkConfiguration"] = o.NetworkConfiguration
	}
	if o.TemplateInputs != nil {
		toSerialize["TemplateInputs"] = o.TemplateInputs
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profile != nil {
		toSerialize["Profile"] = o.Profile
	}
	if o.ServerNode != nil {
		toSerialize["ServerNode"] = o.ServerNode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SdwanRouterNode) UnmarshalJSON(bytes []byte) (err error) {
	type SdwanRouterNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Name of the Cisco vManage device template that the current device should be attached to. A device template consists of many feature templates that contain SD-WAN vEdge router configuration.
		DeviceTemplate *string `json:"DeviceTemplate,omitempty"`
		// Name of the router node object.
		Name                 *string                         `json:"Name,omitempty"`
		NetworkConfiguration []SdwanNetworkConfigurationType `json:"NetworkConfiguration,omitempty"`
		TemplateInputs       []SdwanTemplateInputsType       `json:"TemplateInputs,omitempty"`
		// Uniquely identifies the router by its chassis number.
		Uuid         *string                               `json:"Uuid,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		Profile      *SdwanProfileRelationship             `json:"Profile,omitempty"`
		ServerNode   *AssetDeviceRegistrationRelationship  `json:"ServerNode,omitempty"`
	}

	varSdwanRouterNodeWithoutEmbeddedStruct := SdwanRouterNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varSdwanRouterNodeWithoutEmbeddedStruct)
	if err == nil {
		varSdwanRouterNode := _SdwanRouterNode{}
		varSdwanRouterNode.ClassId = varSdwanRouterNodeWithoutEmbeddedStruct.ClassId
		varSdwanRouterNode.ObjectType = varSdwanRouterNodeWithoutEmbeddedStruct.ObjectType
		varSdwanRouterNode.DeviceTemplate = varSdwanRouterNodeWithoutEmbeddedStruct.DeviceTemplate
		varSdwanRouterNode.Name = varSdwanRouterNodeWithoutEmbeddedStruct.Name
		varSdwanRouterNode.NetworkConfiguration = varSdwanRouterNodeWithoutEmbeddedStruct.NetworkConfiguration
		varSdwanRouterNode.TemplateInputs = varSdwanRouterNodeWithoutEmbeddedStruct.TemplateInputs
		varSdwanRouterNode.Uuid = varSdwanRouterNodeWithoutEmbeddedStruct.Uuid
		varSdwanRouterNode.Organization = varSdwanRouterNodeWithoutEmbeddedStruct.Organization
		varSdwanRouterNode.Profile = varSdwanRouterNodeWithoutEmbeddedStruct.Profile
		varSdwanRouterNode.ServerNode = varSdwanRouterNodeWithoutEmbeddedStruct.ServerNode
		*o = SdwanRouterNode(varSdwanRouterNode)
	} else {
		return err
	}

	varSdwanRouterNode := _SdwanRouterNode{}

	err = json.Unmarshal(bytes, &varSdwanRouterNode)
	if err == nil {
		o.MoBaseMo = varSdwanRouterNode.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceTemplate")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "NetworkConfiguration")
		delete(additionalProperties, "TemplateInputs")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profile")
		delete(additionalProperties, "ServerNode")

		// remove fields from embedded structs
		reflectMoBaseMo := reflect.ValueOf(o.MoBaseMo)
		for i := 0; i < reflectMoBaseMo.Type().NumField(); i++ {
			t := reflectMoBaseMo.Type().Field(i)

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

type NullableSdwanRouterNode struct {
	value *SdwanRouterNode
	isSet bool
}

func (v NullableSdwanRouterNode) Get() *SdwanRouterNode {
	return v.value
}

func (v *NullableSdwanRouterNode) Set(val *SdwanRouterNode) {
	v.value = val
	v.isSet = true
}

func (v NullableSdwanRouterNode) IsSet() bool {
	return v.isSet
}

func (v *NullableSdwanRouterNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSdwanRouterNode(val *SdwanRouterNode) *NullableSdwanRouterNode {
	return &NullableSdwanRouterNode{value: val, isSet: true}
}

func (v NullableSdwanRouterNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSdwanRouterNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
