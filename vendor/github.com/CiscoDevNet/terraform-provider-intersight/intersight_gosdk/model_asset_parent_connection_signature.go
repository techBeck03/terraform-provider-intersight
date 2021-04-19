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

// AssetParentConnectionSignature A signature generated by a device connector with its private key. Signature is passed to leaf device connectors to authenticate connectivity between devices.
type AssetParentConnectionSignature struct {
	AssetClaimSignature
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The moid of the parent device registration.
	DeviceId *string `json:"DeviceId,omitempty"`
	// The node identity of the parent device, corresponds to the parents ClusterMember.memberIdentity. Used on connect to establish through which device in a cluster the child is connected through.
	NodeId               *string `json:"NodeId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetParentConnectionSignature AssetParentConnectionSignature

// NewAssetParentConnectionSignature instantiates a new AssetParentConnectionSignature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetParentConnectionSignature(classId string, objectType string) *AssetParentConnectionSignature {
	this := AssetParentConnectionSignature{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewAssetParentConnectionSignatureWithDefaults instantiates a new AssetParentConnectionSignature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetParentConnectionSignatureWithDefaults() *AssetParentConnectionSignature {
	this := AssetParentConnectionSignature{}
	var classId string = "asset.ParentConnectionSignature"
	this.ClassId = classId
	var objectType string = "asset.ParentConnectionSignature"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *AssetParentConnectionSignature) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *AssetParentConnectionSignature) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *AssetParentConnectionSignature) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *AssetParentConnectionSignature) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *AssetParentConnectionSignature) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *AssetParentConnectionSignature) SetObjectType(v string) {
	o.ObjectType = v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *AssetParentConnectionSignature) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetParentConnectionSignature) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *AssetParentConnectionSignature) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *AssetParentConnectionSignature) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *AssetParentConnectionSignature) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetParentConnectionSignature) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *AssetParentConnectionSignature) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *AssetParentConnectionSignature) SetNodeId(v string) {
	o.NodeId = &v
}

func (o AssetParentConnectionSignature) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedAssetClaimSignature, errAssetClaimSignature := json.Marshal(o.AssetClaimSignature)
	if errAssetClaimSignature != nil {
		return []byte{}, errAssetClaimSignature
	}
	errAssetClaimSignature = json.Unmarshal([]byte(serializedAssetClaimSignature), &toSerialize)
	if errAssetClaimSignature != nil {
		return []byte{}, errAssetClaimSignature
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.DeviceId != nil {
		toSerialize["DeviceId"] = o.DeviceId
	}
	if o.NodeId != nil {
		toSerialize["NodeId"] = o.NodeId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AssetParentConnectionSignature) UnmarshalJSON(bytes []byte) (err error) {
	type AssetParentConnectionSignatureWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
		ObjectType string `json:"ObjectType"`
		// The moid of the parent device registration.
		DeviceId *string `json:"DeviceId,omitempty"`
		// The node identity of the parent device, corresponds to the parents ClusterMember.memberIdentity. Used on connect to establish through which device in a cluster the child is connected through.
		NodeId *string `json:"NodeId,omitempty"`
	}

	varAssetParentConnectionSignatureWithoutEmbeddedStruct := AssetParentConnectionSignatureWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varAssetParentConnectionSignatureWithoutEmbeddedStruct)
	if err == nil {
		varAssetParentConnectionSignature := _AssetParentConnectionSignature{}
		varAssetParentConnectionSignature.ClassId = varAssetParentConnectionSignatureWithoutEmbeddedStruct.ClassId
		varAssetParentConnectionSignature.ObjectType = varAssetParentConnectionSignatureWithoutEmbeddedStruct.ObjectType
		varAssetParentConnectionSignature.DeviceId = varAssetParentConnectionSignatureWithoutEmbeddedStruct.DeviceId
		varAssetParentConnectionSignature.NodeId = varAssetParentConnectionSignatureWithoutEmbeddedStruct.NodeId
		*o = AssetParentConnectionSignature(varAssetParentConnectionSignature)
	} else {
		return err
	}

	varAssetParentConnectionSignature := _AssetParentConnectionSignature{}

	err = json.Unmarshal(bytes, &varAssetParentConnectionSignature)
	if err == nil {
		o.AssetClaimSignature = varAssetParentConnectionSignature.AssetClaimSignature
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "DeviceId")
		delete(additionalProperties, "NodeId")

		// remove fields from embedded structs
		reflectAssetClaimSignature := reflect.ValueOf(o.AssetClaimSignature)
		for i := 0; i < reflectAssetClaimSignature.Type().NumField(); i++ {
			t := reflectAssetClaimSignature.Type().Field(i)

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

type NullableAssetParentConnectionSignature struct {
	value *AssetParentConnectionSignature
	isSet bool
}

func (v NullableAssetParentConnectionSignature) Get() *AssetParentConnectionSignature {
	return v.value
}

func (v *NullableAssetParentConnectionSignature) Set(val *AssetParentConnectionSignature) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetParentConnectionSignature) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetParentConnectionSignature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetParentConnectionSignature(val *AssetParentConnectionSignature) *NullableAssetParentConnectionSignature {
	return &NullableAssetParentConnectionSignature{value: val, isSet: true}
}

func (v NullableAssetParentConnectionSignature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetParentConnectionSignature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
