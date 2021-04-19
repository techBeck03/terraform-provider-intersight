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

// StorageNetAppNode NetApp node is a controller in a NetApp cluster. Services and components are controlled and managed by the NetApp node.
type StorageNetAppNode struct {
	StorageBaseArrayController
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// Health of NetApp Node.
	Health *bool `json:"Health,omitempty"`
	// System id of NetApp Node.
	Systemid *string `json:"Systemid,omitempty"`
	// UUID of NetApp Node.
	Uuid                 *string                           `json:"Uuid,omitempty"`
	Array                *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _StorageNetAppNode StorageNetAppNode

// NewStorageNetAppNode instantiates a new StorageNetAppNode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStorageNetAppNode(classId string, objectType string) *StorageNetAppNode {
	this := StorageNetAppNode{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewStorageNetAppNodeWithDefaults instantiates a new StorageNetAppNode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStorageNetAppNodeWithDefaults() *StorageNetAppNode {
	this := StorageNetAppNode{}
	var classId string = "storage.NetAppNode"
	this.ClassId = classId
	var objectType string = "storage.NetAppNode"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *StorageNetAppNode) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *StorageNetAppNode) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *StorageNetAppNode) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *StorageNetAppNode) SetObjectType(v string) {
	o.ObjectType = v
}

// GetHealth returns the Health field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetHealth() bool {
	if o == nil || o.Health == nil {
		var ret bool
		return ret
	}
	return *o.Health
}

// GetHealthOk returns a tuple with the Health field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetHealthOk() (*bool, bool) {
	if o == nil || o.Health == nil {
		return nil, false
	}
	return o.Health, true
}

// HasHealth returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasHealth() bool {
	if o != nil && o.Health != nil {
		return true
	}

	return false
}

// SetHealth gets a reference to the given bool and assigns it to the Health field.
func (o *StorageNetAppNode) SetHealth(v bool) {
	o.Health = &v
}

// GetSystemid returns the Systemid field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetSystemid() string {
	if o == nil || o.Systemid == nil {
		var ret string
		return ret
	}
	return *o.Systemid
}

// GetSystemidOk returns a tuple with the Systemid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetSystemidOk() (*string, bool) {
	if o == nil || o.Systemid == nil {
		return nil, false
	}
	return o.Systemid, true
}

// HasSystemid returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasSystemid() bool {
	if o != nil && o.Systemid != nil {
		return true
	}

	return false
}

// SetSystemid gets a reference to the given string and assigns it to the Systemid field.
func (o *StorageNetAppNode) SetSystemid(v string) {
	o.Systemid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetUuid() string {
	if o == nil || o.Uuid == nil {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetUuidOk() (*string, bool) {
	if o == nil || o.Uuid == nil {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasUuid() bool {
	if o != nil && o.Uuid != nil {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *StorageNetAppNode) SetUuid(v string) {
	o.Uuid = &v
}

// GetArray returns the Array field value if set, zero value otherwise.
func (o *StorageNetAppNode) GetArray() StorageNetAppClusterRelationship {
	if o == nil || o.Array == nil {
		var ret StorageNetAppClusterRelationship
		return ret
	}
	return *o.Array
}

// GetArrayOk returns a tuple with the Array field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StorageNetAppNode) GetArrayOk() (*StorageNetAppClusterRelationship, bool) {
	if o == nil || o.Array == nil {
		return nil, false
	}
	return o.Array, true
}

// HasArray returns a boolean if a field has been set.
func (o *StorageNetAppNode) HasArray() bool {
	if o != nil && o.Array != nil {
		return true
	}

	return false
}

// SetArray gets a reference to the given StorageNetAppClusterRelationship and assigns it to the Array field.
func (o *StorageNetAppNode) SetArray(v StorageNetAppClusterRelationship) {
	o.Array = &v
}

func (o StorageNetAppNode) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedStorageBaseArrayController, errStorageBaseArrayController := json.Marshal(o.StorageBaseArrayController)
	if errStorageBaseArrayController != nil {
		return []byte{}, errStorageBaseArrayController
	}
	errStorageBaseArrayController = json.Unmarshal([]byte(serializedStorageBaseArrayController), &toSerialize)
	if errStorageBaseArrayController != nil {
		return []byte{}, errStorageBaseArrayController
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Health != nil {
		toSerialize["Health"] = o.Health
	}
	if o.Systemid != nil {
		toSerialize["Systemid"] = o.Systemid
	}
	if o.Uuid != nil {
		toSerialize["Uuid"] = o.Uuid
	}
	if o.Array != nil {
		toSerialize["Array"] = o.Array
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *StorageNetAppNode) UnmarshalJSON(bytes []byte) (err error) {
	type StorageNetAppNodeWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// Health of NetApp Node.
		Health *bool `json:"Health,omitempty"`
		// System id of NetApp Node.
		Systemid *string `json:"Systemid,omitempty"`
		// UUID of NetApp Node.
		Uuid  *string                           `json:"Uuid,omitempty"`
		Array *StorageNetAppClusterRelationship `json:"Array,omitempty"`
	}

	varStorageNetAppNodeWithoutEmbeddedStruct := StorageNetAppNodeWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varStorageNetAppNodeWithoutEmbeddedStruct)
	if err == nil {
		varStorageNetAppNode := _StorageNetAppNode{}
		varStorageNetAppNode.ClassId = varStorageNetAppNodeWithoutEmbeddedStruct.ClassId
		varStorageNetAppNode.ObjectType = varStorageNetAppNodeWithoutEmbeddedStruct.ObjectType
		varStorageNetAppNode.Health = varStorageNetAppNodeWithoutEmbeddedStruct.Health
		varStorageNetAppNode.Systemid = varStorageNetAppNodeWithoutEmbeddedStruct.Systemid
		varStorageNetAppNode.Uuid = varStorageNetAppNodeWithoutEmbeddedStruct.Uuid
		varStorageNetAppNode.Array = varStorageNetAppNodeWithoutEmbeddedStruct.Array
		*o = StorageNetAppNode(varStorageNetAppNode)
	} else {
		return err
	}

	varStorageNetAppNode := _StorageNetAppNode{}

	err = json.Unmarshal(bytes, &varStorageNetAppNode)
	if err == nil {
		o.StorageBaseArrayController = varStorageNetAppNode.StorageBaseArrayController
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Health")
		delete(additionalProperties, "Systemid")
		delete(additionalProperties, "Uuid")
		delete(additionalProperties, "Array")

		// remove fields from embedded structs
		reflectStorageBaseArrayController := reflect.ValueOf(o.StorageBaseArrayController)
		for i := 0; i < reflectStorageBaseArrayController.Type().NumField(); i++ {
			t := reflectStorageBaseArrayController.Type().Field(i)

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

type NullableStorageNetAppNode struct {
	value *StorageNetAppNode
	isSet bool
}

func (v NullableStorageNetAppNode) Get() *StorageNetAppNode {
	return v.value
}

func (v *NullableStorageNetAppNode) Set(val *StorageNetAppNode) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageNetAppNode) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageNetAppNode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageNetAppNode(val *StorageNetAppNode) *NullableStorageNetAppNode {
	return &NullableStorageNetAppNode{value: val, isSet: true}
}

func (v NullableStorageNetAppNode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageNetAppNode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
