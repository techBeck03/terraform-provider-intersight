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

// ChassisConfigImport Configuration import action will import the existing configuration from chassis and generate Intersight policies profile from it. At end of successful import, chassis will be assigned to the generated profile which has policies associated with it. No chassis profile or policies will be generated if configuration import fails.
type ChassisConfigImport struct {
	MoBaseMo
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Description of the imported profile.
	Description *string `json:"Description,omitempty"`
	// Policy prefix for the policies of the imported chassis profile.
	PolicyPrefix *string  `json:"PolicyPrefix,omitempty"`
	PolicyTypes  []string `json:"PolicyTypes,omitempty"`
	// Profile name for the imported chassis profile.
	ProfileName          *string                               `json:"ProfileName,omitempty"`
	Chassis              *EquipmentChassisRelationship         `json:"Chassis,omitempty"`
	ChassisProfile       *ChassisProfileRelationship           `json:"ChassisProfile,omitempty"`
	Organization         *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ChassisConfigImport ChassisConfigImport

// NewChassisConfigImport instantiates a new ChassisConfigImport object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewChassisConfigImport(classId string, objectType string) *ChassisConfigImport {
	this := ChassisConfigImport{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewChassisConfigImportWithDefaults instantiates a new ChassisConfigImport object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewChassisConfigImportWithDefaults() *ChassisConfigImport {
	this := ChassisConfigImport{}
	var classId string = "chassis.ConfigImport"
	this.ClassId = classId
	var objectType string = "chassis.ConfigImport"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *ChassisConfigImport) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ChassisConfigImport) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ChassisConfigImport) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ChassisConfigImport) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ChassisConfigImport) SetDescription(v string) {
	o.Description = &v
}

// GetPolicyPrefix returns the PolicyPrefix field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetPolicyPrefix() string {
	if o == nil || o.PolicyPrefix == nil {
		var ret string
		return ret
	}
	return *o.PolicyPrefix
}

// GetPolicyPrefixOk returns a tuple with the PolicyPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetPolicyPrefixOk() (*string, bool) {
	if o == nil || o.PolicyPrefix == nil {
		return nil, false
	}
	return o.PolicyPrefix, true
}

// HasPolicyPrefix returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasPolicyPrefix() bool {
	if o != nil && o.PolicyPrefix != nil {
		return true
	}

	return false
}

// SetPolicyPrefix gets a reference to the given string and assigns it to the PolicyPrefix field.
func (o *ChassisConfigImport) SetPolicyPrefix(v string) {
	o.PolicyPrefix = &v
}

// GetPolicyTypes returns the PolicyTypes field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ChassisConfigImport) GetPolicyTypes() []string {
	if o == nil {
		var ret []string
		return ret
	}
	return o.PolicyTypes
}

// GetPolicyTypesOk returns a tuple with the PolicyTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ChassisConfigImport) GetPolicyTypesOk() (*[]string, bool) {
	if o == nil || o.PolicyTypes == nil {
		return nil, false
	}
	return &o.PolicyTypes, true
}

// HasPolicyTypes returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasPolicyTypes() bool {
	if o != nil && o.PolicyTypes != nil {
		return true
	}

	return false
}

// SetPolicyTypes gets a reference to the given []string and assigns it to the PolicyTypes field.
func (o *ChassisConfigImport) SetPolicyTypes(v []string) {
	o.PolicyTypes = v
}

// GetProfileName returns the ProfileName field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetProfileName() string {
	if o == nil || o.ProfileName == nil {
		var ret string
		return ret
	}
	return *o.ProfileName
}

// GetProfileNameOk returns a tuple with the ProfileName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetProfileNameOk() (*string, bool) {
	if o == nil || o.ProfileName == nil {
		return nil, false
	}
	return o.ProfileName, true
}

// HasProfileName returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasProfileName() bool {
	if o != nil && o.ProfileName != nil {
		return true
	}

	return false
}

// SetProfileName gets a reference to the given string and assigns it to the ProfileName field.
func (o *ChassisConfigImport) SetProfileName(v string) {
	o.ProfileName = &v
}

// GetChassis returns the Chassis field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetChassis() EquipmentChassisRelationship {
	if o == nil || o.Chassis == nil {
		var ret EquipmentChassisRelationship
		return ret
	}
	return *o.Chassis
}

// GetChassisOk returns a tuple with the Chassis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetChassisOk() (*EquipmentChassisRelationship, bool) {
	if o == nil || o.Chassis == nil {
		return nil, false
	}
	return o.Chassis, true
}

// HasChassis returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasChassis() bool {
	if o != nil && o.Chassis != nil {
		return true
	}

	return false
}

// SetChassis gets a reference to the given EquipmentChassisRelationship and assigns it to the Chassis field.
func (o *ChassisConfigImport) SetChassis(v EquipmentChassisRelationship) {
	o.Chassis = &v
}

// GetChassisProfile returns the ChassisProfile field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetChassisProfile() ChassisProfileRelationship {
	if o == nil || o.ChassisProfile == nil {
		var ret ChassisProfileRelationship
		return ret
	}
	return *o.ChassisProfile
}

// GetChassisProfileOk returns a tuple with the ChassisProfile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetChassisProfileOk() (*ChassisProfileRelationship, bool) {
	if o == nil || o.ChassisProfile == nil {
		return nil, false
	}
	return o.ChassisProfile, true
}

// HasChassisProfile returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasChassisProfile() bool {
	if o != nil && o.ChassisProfile != nil {
		return true
	}

	return false
}

// SetChassisProfile gets a reference to the given ChassisProfileRelationship and assigns it to the ChassisProfile field.
func (o *ChassisConfigImport) SetChassisProfile(v ChassisProfileRelationship) {
	o.ChassisProfile = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *ChassisConfigImport) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ChassisConfigImport) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *ChassisConfigImport) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *ChassisConfigImport) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

func (o ChassisConfigImport) MarshalJSON() ([]byte, error) {
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
	if o.Description != nil {
		toSerialize["Description"] = o.Description
	}
	if o.PolicyPrefix != nil {
		toSerialize["PolicyPrefix"] = o.PolicyPrefix
	}
	if o.PolicyTypes != nil {
		toSerialize["PolicyTypes"] = o.PolicyTypes
	}
	if o.ProfileName != nil {
		toSerialize["ProfileName"] = o.ProfileName
	}
	if o.Chassis != nil {
		toSerialize["Chassis"] = o.Chassis
	}
	if o.ChassisProfile != nil {
		toSerialize["ChassisProfile"] = o.ChassisProfile
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ChassisConfigImport) UnmarshalJSON(bytes []byte) (err error) {
	type ChassisConfigImportWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Description of the imported profile.
		Description *string `json:"Description,omitempty"`
		// Policy prefix for the policies of the imported chassis profile.
		PolicyPrefix *string  `json:"PolicyPrefix,omitempty"`
		PolicyTypes  []string `json:"PolicyTypes,omitempty"`
		// Profile name for the imported chassis profile.
		ProfileName    *string                               `json:"ProfileName,omitempty"`
		Chassis        *EquipmentChassisRelationship         `json:"Chassis,omitempty"`
		ChassisProfile *ChassisProfileRelationship           `json:"ChassisProfile,omitempty"`
		Organization   *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	}

	varChassisConfigImportWithoutEmbeddedStruct := ChassisConfigImportWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varChassisConfigImportWithoutEmbeddedStruct)
	if err == nil {
		varChassisConfigImport := _ChassisConfigImport{}
		varChassisConfigImport.ClassId = varChassisConfigImportWithoutEmbeddedStruct.ClassId
		varChassisConfigImport.ObjectType = varChassisConfigImportWithoutEmbeddedStruct.ObjectType
		varChassisConfigImport.Description = varChassisConfigImportWithoutEmbeddedStruct.Description
		varChassisConfigImport.PolicyPrefix = varChassisConfigImportWithoutEmbeddedStruct.PolicyPrefix
		varChassisConfigImport.PolicyTypes = varChassisConfigImportWithoutEmbeddedStruct.PolicyTypes
		varChassisConfigImport.ProfileName = varChassisConfigImportWithoutEmbeddedStruct.ProfileName
		varChassisConfigImport.Chassis = varChassisConfigImportWithoutEmbeddedStruct.Chassis
		varChassisConfigImport.ChassisProfile = varChassisConfigImportWithoutEmbeddedStruct.ChassisProfile
		varChassisConfigImport.Organization = varChassisConfigImportWithoutEmbeddedStruct.Organization
		*o = ChassisConfigImport(varChassisConfigImport)
	} else {
		return err
	}

	varChassisConfigImport := _ChassisConfigImport{}

	err = json.Unmarshal(bytes, &varChassisConfigImport)
	if err == nil {
		o.MoBaseMo = varChassisConfigImport.MoBaseMo
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Description")
		delete(additionalProperties, "PolicyPrefix")
		delete(additionalProperties, "PolicyTypes")
		delete(additionalProperties, "ProfileName")
		delete(additionalProperties, "Chassis")
		delete(additionalProperties, "ChassisProfile")
		delete(additionalProperties, "Organization")

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

type NullableChassisConfigImport struct {
	value *ChassisConfigImport
	isSet bool
}

func (v NullableChassisConfigImport) Get() *ChassisConfigImport {
	return v.value
}

func (v *NullableChassisConfigImport) Set(val *ChassisConfigImport) {
	v.value = val
	v.isSet = true
}

func (v NullableChassisConfigImport) IsSet() bool {
	return v.isSet
}

func (v *NullableChassisConfigImport) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChassisConfigImport(val *ChassisConfigImport) *NullableChassisConfigImport {
	return &NullableChassisConfigImport{value: val, isSet: true}
}

func (v NullableChassisConfigImport) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChassisConfigImport) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
