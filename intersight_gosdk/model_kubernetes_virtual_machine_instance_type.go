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

// KubernetesVirtualMachineInstanceType A policy specifying CPU, Memory and Disk size configuration for a Virtual Machine.
type KubernetesVirtualMachineInstanceType struct {
	PolicyAbstractPolicy
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Number of CPUs allocated to virtual machine.
	Cpu *int64 `json:"Cpu,omitempty"`
	// Ephemeral disk capacity to be provided with units example - 10Gi.
	DiskSize *int64 `json:"DiskSize,omitempty"`
	// Virtual machine memory defined in mebibytes (MiB).
	Memory       *int64                                `json:"Memory,omitempty"`
	Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
	// An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources.
	Profiles             []KubernetesVirtualMachineInfrastructureProviderRelationship `json:"Profiles,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesVirtualMachineInstanceType KubernetesVirtualMachineInstanceType

// NewKubernetesVirtualMachineInstanceType instantiates a new KubernetesVirtualMachineInstanceType object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesVirtualMachineInstanceType(classId string, objectType string) *KubernetesVirtualMachineInstanceType {
	this := KubernetesVirtualMachineInstanceType{}
	this.ClassId = classId
	this.ObjectType = objectType
	var cpu int64 = 4
	this.Cpu = &cpu
	var memory int64 = 16384
	this.Memory = &memory
	return &this
}

// NewKubernetesVirtualMachineInstanceTypeWithDefaults instantiates a new KubernetesVirtualMachineInstanceType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesVirtualMachineInstanceTypeWithDefaults() *KubernetesVirtualMachineInstanceType {
	this := KubernetesVirtualMachineInstanceType{}
	var classId string = "kubernetes.VirtualMachineInstanceType"
	this.ClassId = classId
	var objectType string = "kubernetes.VirtualMachineInstanceType"
	this.ObjectType = objectType
	var cpu int64 = 4
	this.Cpu = &cpu
	var memory int64 = 16384
	this.Memory = &memory
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesVirtualMachineInstanceType) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesVirtualMachineInstanceType) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesVirtualMachineInstanceType) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesVirtualMachineInstanceType) SetObjectType(v string) {
	o.ObjectType = v
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceType) GetCpu() int64 {
	if o == nil || o.Cpu == nil {
		var ret int64
		return ret
	}
	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetCpuOk() (*int64, bool) {
	if o == nil || o.Cpu == nil {
		return nil, false
	}
	return o.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceType) HasCpu() bool {
	if o != nil && o.Cpu != nil {
		return true
	}

	return false
}

// SetCpu gets a reference to the given int64 and assigns it to the Cpu field.
func (o *KubernetesVirtualMachineInstanceType) SetCpu(v int64) {
	o.Cpu = &v
}

// GetDiskSize returns the DiskSize field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceType) GetDiskSize() int64 {
	if o == nil || o.DiskSize == nil {
		var ret int64
		return ret
	}
	return *o.DiskSize
}

// GetDiskSizeOk returns a tuple with the DiskSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetDiskSizeOk() (*int64, bool) {
	if o == nil || o.DiskSize == nil {
		return nil, false
	}
	return o.DiskSize, true
}

// HasDiskSize returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceType) HasDiskSize() bool {
	if o != nil && o.DiskSize != nil {
		return true
	}

	return false
}

// SetDiskSize gets a reference to the given int64 and assigns it to the DiskSize field.
func (o *KubernetesVirtualMachineInstanceType) SetDiskSize(v int64) {
	o.DiskSize = &v
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceType) GetMemory() int64 {
	if o == nil || o.Memory == nil {
		var ret int64
		return ret
	}
	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetMemoryOk() (*int64, bool) {
	if o == nil || o.Memory == nil {
		return nil, false
	}
	return o.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceType) HasMemory() bool {
	if o != nil && o.Memory != nil {
		return true
	}

	return false
}

// SetMemory gets a reference to the given int64 and assigns it to the Memory field.
func (o *KubernetesVirtualMachineInstanceType) SetMemory(v int64) {
	o.Memory = &v
}

// GetOrganization returns the Organization field value if set, zero value otherwise.
func (o *KubernetesVirtualMachineInstanceType) GetOrganization() OrganizationOrganizationRelationship {
	if o == nil || o.Organization == nil {
		var ret OrganizationOrganizationRelationship
		return ret
	}
	return *o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesVirtualMachineInstanceType) GetOrganizationOk() (*OrganizationOrganizationRelationship, bool) {
	if o == nil || o.Organization == nil {
		return nil, false
	}
	return o.Organization, true
}

// HasOrganization returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceType) HasOrganization() bool {
	if o != nil && o.Organization != nil {
		return true
	}

	return false
}

// SetOrganization gets a reference to the given OrganizationOrganizationRelationship and assigns it to the Organization field.
func (o *KubernetesVirtualMachineInstanceType) SetOrganization(v OrganizationOrganizationRelationship) {
	o.Organization = &v
}

// GetProfiles returns the Profiles field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *KubernetesVirtualMachineInstanceType) GetProfiles() []KubernetesVirtualMachineInfrastructureProviderRelationship {
	if o == nil {
		var ret []KubernetesVirtualMachineInfrastructureProviderRelationship
		return ret
	}
	return o.Profiles
}

// GetProfilesOk returns a tuple with the Profiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *KubernetesVirtualMachineInstanceType) GetProfilesOk() (*[]KubernetesVirtualMachineInfrastructureProviderRelationship, bool) {
	if o == nil || o.Profiles == nil {
		return nil, false
	}
	return &o.Profiles, true
}

// HasProfiles returns a boolean if a field has been set.
func (o *KubernetesVirtualMachineInstanceType) HasProfiles() bool {
	if o != nil && o.Profiles != nil {
		return true
	}

	return false
}

// SetProfiles gets a reference to the given []KubernetesVirtualMachineInfrastructureProviderRelationship and assigns it to the Profiles field.
func (o *KubernetesVirtualMachineInstanceType) SetProfiles(v []KubernetesVirtualMachineInfrastructureProviderRelationship) {
	o.Profiles = v
}

func (o KubernetesVirtualMachineInstanceType) MarshalJSON() ([]byte, error) {
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
	if o.Cpu != nil {
		toSerialize["Cpu"] = o.Cpu
	}
	if o.DiskSize != nil {
		toSerialize["DiskSize"] = o.DiskSize
	}
	if o.Memory != nil {
		toSerialize["Memory"] = o.Memory
	}
	if o.Organization != nil {
		toSerialize["Organization"] = o.Organization
	}
	if o.Profiles != nil {
		toSerialize["Profiles"] = o.Profiles
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesVirtualMachineInstanceType) UnmarshalJSON(bytes []byte) (err error) {
	type KubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Number of CPUs allocated to virtual machine.
		Cpu *int64 `json:"Cpu,omitempty"`
		// Ephemeral disk capacity to be provided with units example - 10Gi.
		DiskSize *int64 `json:"DiskSize,omitempty"`
		// Virtual machine memory defined in mebibytes (MiB).
		Memory       *int64                                `json:"Memory,omitempty"`
		Organization *OrganizationOrganizationRelationship `json:"Organization,omitempty"`
		// An array of relationships to kubernetesVirtualMachineInfrastructureProvider resources.
		Profiles []KubernetesVirtualMachineInfrastructureProviderRelationship `json:"Profiles,omitempty"`
	}

	varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct := KubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct)
	if err == nil {
		varKubernetesVirtualMachineInstanceType := _KubernetesVirtualMachineInstanceType{}
		varKubernetesVirtualMachineInstanceType.ClassId = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.ClassId
		varKubernetesVirtualMachineInstanceType.ObjectType = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.ObjectType
		varKubernetesVirtualMachineInstanceType.Cpu = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.Cpu
		varKubernetesVirtualMachineInstanceType.DiskSize = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.DiskSize
		varKubernetesVirtualMachineInstanceType.Memory = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.Memory
		varKubernetesVirtualMachineInstanceType.Organization = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.Organization
		varKubernetesVirtualMachineInstanceType.Profiles = varKubernetesVirtualMachineInstanceTypeWithoutEmbeddedStruct.Profiles
		*o = KubernetesVirtualMachineInstanceType(varKubernetesVirtualMachineInstanceType)
	} else {
		return err
	}

	varKubernetesVirtualMachineInstanceType := _KubernetesVirtualMachineInstanceType{}

	err = json.Unmarshal(bytes, &varKubernetesVirtualMachineInstanceType)
	if err == nil {
		o.PolicyAbstractPolicy = varKubernetesVirtualMachineInstanceType.PolicyAbstractPolicy
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Cpu")
		delete(additionalProperties, "DiskSize")
		delete(additionalProperties, "Memory")
		delete(additionalProperties, "Organization")
		delete(additionalProperties, "Profiles")

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

type NullableKubernetesVirtualMachineInstanceType struct {
	value *KubernetesVirtualMachineInstanceType
	isSet bool
}

func (v NullableKubernetesVirtualMachineInstanceType) Get() *KubernetesVirtualMachineInstanceType {
	return v.value
}

func (v *NullableKubernetesVirtualMachineInstanceType) Set(val *KubernetesVirtualMachineInstanceType) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesVirtualMachineInstanceType) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesVirtualMachineInstanceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesVirtualMachineInstanceType(val *KubernetesVirtualMachineInstanceType) *NullableKubernetesVirtualMachineInstanceType {
	return &NullableKubernetesVirtualMachineInstanceType{value: val, isSet: true}
}

func (v NullableKubernetesVirtualMachineInstanceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesVirtualMachineInstanceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
