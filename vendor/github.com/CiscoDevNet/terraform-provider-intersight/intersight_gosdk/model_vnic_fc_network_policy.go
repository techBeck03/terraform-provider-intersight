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

// VnicFcNetworkPolicy A Fibre Channel Network policy governs the vSAN configuration for the virtual interfaces.
type VnicFcNetworkPolicy struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType           string                                `json:"ObjectType"`
	VsanSettings         NullableVnicVsanSettings              `json:"VsanSettings,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicFcNetworkPolicy VnicFcNetworkPolicy

// NewVnicFcNetworkPolicy instantiates a new VnicFcNetworkPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicFcNetworkPolicy(classId string, objectType string) *VnicFcNetworkPolicy {
	this := VnicFcNetworkPolicy{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicFcNetworkPolicyWithDefaults instantiates a new VnicFcNetworkPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicFcNetworkPolicyWithDefaults() *VnicFcNetworkPolicy {
	this := VnicFcNetworkPolicy{}
	var classId string = "vnic.FcNetworkPolicy"
	this.ClassId = classId
	var objectType string = "vnic.FcNetworkPolicy"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicFcNetworkPolicy) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicFcNetworkPolicy) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicFcNetworkPolicy) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicFcNetworkPolicy) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicFcNetworkPolicy) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicFcNetworkPolicy) SetObjectType(v string) {
	o.ObjectType = v
}

// GetVsanSettings returns the VsanSettings field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *VnicFcNetworkPolicy) GetVsanSettings() VnicVsanSettings {
	if o == nil || o.VsanSettings.Get() == nil {
		var ret VnicVsanSettings
		return ret
	}
	return *o.VsanSettings.Get()
}

// GetVsanSettingsOk returns a tuple with the VsanSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *VnicFcNetworkPolicy) GetVsanSettingsOk() (*VnicVsanSettings, bool) {
	if o == nil {
		return nil, false
	}
	return o.VsanSettings.Get(), o.VsanSettings.IsSet()
}

// HasVsanSettings returns a boolean if a field has been set.
func (o *VnicFcNetworkPolicy) HasVsanSettings() bool {
	if o != nil && o.VsanSettings.IsSet() {
		return true
	}

	return false
}

// SetVsanSettings gets a reference to the given NullableVnicVsanSettings and assigns it to the VsanSettings field.
func (o *VnicFcNetworkPolicy) SetVsanSettings(v VnicVsanSettings) {
	o.VsanSettings.Set(&v)
}

// SetVsanSettingsNil sets the value for VsanSettings to be an explicit nil
func (o *VnicFcNetworkPolicy) SetVsanSettingsNil() {
	o.VsanSettings.Set(nil)
}

// UnsetVsanSettings ensures that no value is present for VsanSettings, not even an explicit nil
func (o *VnicFcNetworkPolicy) UnsetVsanSettings() {
	o.VsanSettings.Unset()
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *VnicFcNetworkPolicy) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicFcNetworkPolicy) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *VnicFcNetworkPolicy) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *VnicFcNetworkPolicy) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o VnicFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyAbstractPolicy, errPolicyAbstractPolicy := json.Marshal(o.PolicyAbstractPolicy)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	errPolicyAbstractPolicy = json.Unmarshal([]byte(serializedPolicyAbstractPolicy), &toSerialize)
	if errPolicyAbstractPolicy != nil {
		return []byte{}, errPolicyAbstractPolicy
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.VsanSettings.IsSet() {
		toSerialize["VsanSettings"] = o.VsanSettings.Get()
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicFcNetworkPolicy) UnmarshalJSON(bytes []byte) (err error) {
	type VnicFcNetworkPolicyWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType   string                                `json:"ObjectType"`
		VsanSettings NullableVnicVsanSettings              `json:"VsanSettings,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varVnicFcNetworkPolicyWithoutEmbeddedStruct := VnicFcNetworkPolicyWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicFcNetworkPolicyWithoutEmbeddedStruct)
	if err == nil {
		varVnicFcNetworkPolicy := _VnicFcNetworkPolicy{}
		varVnicFcNetworkPolicy.ClassId = varVnicFcNetworkPolicyWithoutEmbeddedStruct.ClassId
		varVnicFcNetworkPolicy.ObjectType = varVnicFcNetworkPolicyWithoutEmbeddedStruct.ObjectType
		varVnicFcNetworkPolicy.VsanSettings = varVnicFcNetworkPolicyWithoutEmbeddedStruct.VsanSettings
		varVnicFcNetworkPolicy.Organization = varVnicFcNetworkPolicyWithoutEmbeddedStruct.Organization
		*o = VnicFcNetworkPolicy(varVnicFcNetworkPolicy)
	} else {
		return err
	}

	varVnicFcNetworkPolicy := _VnicFcNetworkPolicy{}

	err = json.Unmarshal(bytes, &varVnicFcNetworkPolicy)
	if err == nil {
		o.PolicyAbstractPolicy = varVnicFcNetworkPolicy.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "VsanSettings")
		delete(additionalProperties, "Organization")

		// remove fields from embedded structs
		reflectPolicyAbstractPolicy := reflect.ValueOf(o.PolicyAbstractPolicy)
		for i := 0; i < reflectPolicyAbstractPolicy.Type().NumField(); i++ {
			t := reflectPolicyAbstractPolicy.Type().Field(i)

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

type NullableVnicFcNetworkPolicy struct {
	value *VnicFcNetworkPolicy
	isSet bool
}

func (v NullableVnicFcNetworkPolicy) Get() *VnicFcNetworkPolicy {
	return v.value
}

func (v *NullableVnicFcNetworkPolicy) Set(val *VnicFcNetworkPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicFcNetworkPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicFcNetworkPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicFcNetworkPolicy(val *VnicFcNetworkPolicy) *NullableVnicFcNetworkPolicy {
	return &NullableVnicFcNetworkPolicy{value: val, isSet: true}
}

func (v NullableVnicFcNetworkPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicFcNetworkPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
