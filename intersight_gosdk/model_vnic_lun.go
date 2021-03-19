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

// VnicLun Lun specific parameters that will be part of a vNIC.
type VnicLun struct {
	MoBaseComplexType
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Specifies LUN is bootable.
	Bootable *bool `json:"Bootable,omitempty"`
	// The Identifier of the LUN.
	LunId                *int64 `json:"LunId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _VnicLun VnicLun

// NewVnicLun instantiates a new VnicLun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVnicLun(classId string, objectType string) *VnicLun {
	this := VnicLun{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewVnicLunWithDefaults instantiates a new VnicLun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVnicLunWithDefaults() *VnicLun {
	this := VnicLun{}
	var classId string = "vnic.Lun"
	this.ClassId = classId
	var objectType string = "vnic.Lun"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *VnicLun) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *VnicLun) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *VnicLun) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *VnicLun) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *VnicLun) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *VnicLun) SetObjectType(v string) {
	o.ObjectType = v
}

// GetBootable returns the Bootable field value if set, zero value otherwise.
func (o *VnicLun) GetBootable() bool {
	if o == nil || o.Bootable == nil {
		var ret bool
		return ret
	}
	return *o.Bootable
}

// GetBootableOk returns a tuple with the Bootable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLun) GetBootableOk() (*bool, bool) {
	if o == nil || o.Bootable == nil {
		return nil, false
	}
	return o.Bootable, true
}

// HasBootable returns a boolean if a field has been set.
func (o *VnicLun) HasBootable() bool {
	if o != nil && o.Bootable != nil {
		return true
	}

	return false
}

// SetBootable gets a reference to the given bool and assigns it to the Bootable field.
func (o *VnicLun) SetBootable(v bool) {
	o.Bootable = &v
}

// GetLunId returns the LunId field value if set, zero value otherwise.
func (o *VnicLun) GetLunId() int64 {
	if o == nil || o.LunId == nil {
		var ret int64
		return ret
	}
	return *o.LunId
}

// GetLunIdOk returns a tuple with the LunId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VnicLun) GetLunIdOk() (*int64, bool) {
	if o == nil || o.LunId == nil {
		return nil, false
	}
	return o.LunId, true
}

// HasLunId returns a boolean if a field has been set.
func (o *VnicLun) HasLunId() bool {
	if o != nil && o.LunId != nil {
		return true
	}

	return false
}

// SetLunId gets a reference to the given int64 and assigns it to the LunId field.
func (o *VnicLun) SetLunId(v int64) {
	o.LunId = &v
}

func (o VnicLun) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedMoBaseComplexType, errMoBaseComplexType := json.Marshal(o.MoBaseComplexType)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	errMoBaseComplexType = json.Unmarshal([]byte(serializedMoBaseComplexType), &toSerialize)
	if errMoBaseComplexType != nil {
		return []byte{}, errMoBaseComplexType
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Bootable != nil {
		toSerialize["Bootable"] = o.Bootable
	}
	if o.LunId != nil {
		toSerialize["LunId"] = o.LunId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *VnicLun) UnmarshalJSON(bytes []byte) (err error) {
	type VnicLunWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Specifies LUN is bootable.
		Bootable *bool `json:"Bootable,omitempty"`
		// The Identifier of the LUN.
		LunId *int64 `json:"LunId,omitempty"`
	}

	varVnicLunWithoutEmbeddedStruct := VnicLunWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varVnicLunWithoutEmbeddedStruct)
	if err == nil {
		varVnicLun := _VnicLun{}
		varVnicLun.ClassId = varVnicLunWithoutEmbeddedStruct.ClassId
		varVnicLun.ObjectType = varVnicLunWithoutEmbeddedStruct.ObjectType
		varVnicLun.Bootable = varVnicLunWithoutEmbeddedStruct.Bootable
		varVnicLun.LunId = varVnicLunWithoutEmbeddedStruct.LunId
		*o = VnicLun(varVnicLun)
	} else {
		return err
	}

	varVnicLun := _VnicLun{}

	err = json.Unmarshal(bytes, &varVnicLun)
	if err == nil {
		o.MoBaseComplexType = varVnicLun.MoBaseComplexType
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Bootable")
		delete(additionalProperties, "LunId")

		// remove fields from embedded structs
		reflectMoBaseComplexType := reflect.ValueOf(o.MoBaseComplexType)
		for i := 0; i < reflectMoBaseComplexType.Type().NumField(); i++ {
			t := reflectMoBaseComplexType.Type().Field(i)

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

type NullableVnicLun struct {
	value *VnicLun
	isSet bool
}

func (v NullableVnicLun) Get() *VnicLun {
	return v.value
}

func (v *NullableVnicLun) Set(val *VnicLun) {
	v.value = val
	v.isSet = true
}

func (v NullableVnicLun) IsSet() bool {
	return v.isSet
}

func (v *NullableVnicLun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVnicLun(val *VnicLun) *NullableVnicLun {
	return &NullableVnicLun{value: val, isSet: true}
}

func (v NullableVnicLun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVnicLun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}