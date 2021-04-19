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
)

// IamDomainGroupAllOf Definition of the list of properties defined in 'iam.DomainGroup', excluding properties defined in parent classes.
type IamDomainGroupAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// The name of the domain-group.
	Name *string `json:"Name,omitempty"`
	// The partition number domain group related messages are produced for 'Partition1' category service topics.
	Partition1 *int64 `json:"Partition1,omitempty"`
	// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition2' category service topics.
	Partition2 *int64 `json:"Partition2,omitempty"`
	// In a cloud environment this parameter will indicate to which partition number domain group related messages are produced for 'Partition3' category service topics.
	Partition3 *int64 `json:"Partition3,omitempty"`
	// Partition key used for producing messages to Kafka partitions. By default Domain-group id will be used as partition key. For Domain-groups belonging to Early adopters domain-group id will be prefixed with 'H' and used as partition key, such partition key will be treated differently and messages will always be produced to partition 0.
	PartitionKey *string `json:"PartitionKey,omitempty"`
	// The number of devices in the domain-group. Device registration notifications are processed to update the usage of the domain-group. The on-boarding account will have multiple domain-groups, and during the device registration least used domain-group will be selected for the device.
	Usage                *int64                  `json:"Usage,omitempty"`
	Account              *IamAccountRelationship `json:"Account,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IamDomainGroupAllOf IamDomainGroupAllOf

// NewIamDomainGroupAllOf instantiates a new IamDomainGroupAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIamDomainGroupAllOf(classId string, objectType string) *IamDomainGroupAllOf {
	this := IamDomainGroupAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewIamDomainGroupAllOfWithDefaults instantiates a new IamDomainGroupAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIamDomainGroupAllOfWithDefaults() *IamDomainGroupAllOf {
	this := IamDomainGroupAllOf{}
	var classId string = "iam.DomainGroup"
	this.ClassId = classId
	var objectType string = "iam.DomainGroup"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *IamDomainGroupAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *IamDomainGroupAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *IamDomainGroupAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *IamDomainGroupAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *IamDomainGroupAllOf) SetName(v string) {
	o.Name = &v
}

// GetPartition1 returns the Partition1 field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetPartition1() int64 {
	if o == nil || o.Partition1 == nil {
		var ret int64
		return ret
	}
	return *o.Partition1
}

// GetPartition1Ok returns a tuple with the Partition1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetPartition1Ok() (*int64, bool) {
	if o == nil || o.Partition1 == nil {
		return nil, false
	}
	return o.Partition1, true
}

// HasPartition1 returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasPartition1() bool {
	if o != nil && o.Partition1 != nil {
		return true
	}

	return false
}

// SetPartition1 gets a reference to the given int64 and assigns it to the Partition1 field.
func (o *IamDomainGroupAllOf) SetPartition1(v int64) {
	o.Partition1 = &v
}

// GetPartition2 returns the Partition2 field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetPartition2() int64 {
	if o == nil || o.Partition2 == nil {
		var ret int64
		return ret
	}
	return *o.Partition2
}

// GetPartition2Ok returns a tuple with the Partition2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetPartition2Ok() (*int64, bool) {
	if o == nil || o.Partition2 == nil {
		return nil, false
	}
	return o.Partition2, true
}

// HasPartition2 returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasPartition2() bool {
	if o != nil && o.Partition2 != nil {
		return true
	}

	return false
}

// SetPartition2 gets a reference to the given int64 and assigns it to the Partition2 field.
func (o *IamDomainGroupAllOf) SetPartition2(v int64) {
	o.Partition2 = &v
}

// GetPartition3 returns the Partition3 field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetPartition3() int64 {
	if o == nil || o.Partition3 == nil {
		var ret int64
		return ret
	}
	return *o.Partition3
}

// GetPartition3Ok returns a tuple with the Partition3 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetPartition3Ok() (*int64, bool) {
	if o == nil || o.Partition3 == nil {
		return nil, false
	}
	return o.Partition3, true
}

// HasPartition3 returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasPartition3() bool {
	if o != nil && o.Partition3 != nil {
		return true
	}

	return false
}

// SetPartition3 gets a reference to the given int64 and assigns it to the Partition3 field.
func (o *IamDomainGroupAllOf) SetPartition3(v int64) {
	o.Partition3 = &v
}

// GetPartitionKey returns the PartitionKey field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetPartitionKey() string {
	if o == nil || o.PartitionKey == nil {
		var ret string
		return ret
	}
	return *o.PartitionKey
}

// GetPartitionKeyOk returns a tuple with the PartitionKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetPartitionKeyOk() (*string, bool) {
	if o == nil || o.PartitionKey == nil {
		return nil, false
	}
	return o.PartitionKey, true
}

// HasPartitionKey returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasPartitionKey() bool {
	if o != nil && o.PartitionKey != nil {
		return true
	}

	return false
}

// SetPartitionKey gets a reference to the given string and assigns it to the PartitionKey field.
func (o *IamDomainGroupAllOf) SetPartitionKey(v string) {
	o.PartitionKey = &v
}

// GetUsage returns the Usage field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetUsage() int64 {
	if o == nil || o.Usage == nil {
		var ret int64
		return ret
	}
	return *o.Usage
}

// GetUsageOk returns a tuple with the Usage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetUsageOk() (*int64, bool) {
	if o == nil || o.Usage == nil {
		return nil, false
	}
	return o.Usage, true
}

// HasUsage returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasUsage() bool {
	if o != nil && o.Usage != nil {
		return true
	}

	return false
}

// SetUsage gets a reference to the given int64 and assigns it to the Usage field.
func (o *IamDomainGroupAllOf) SetUsage(v int64) {
	o.Usage = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *IamDomainGroupAllOf) GetAccount() IamAccountRelationship {
	if o == nil || o.Account == nil {
		var ret IamAccountRelationship
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IamDomainGroupAllOf) GetAccountOk() (*IamAccountRelationship, bool) {
	if o == nil || o.Account == nil {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *IamDomainGroupAllOf) HasAccount() bool {
	if o != nil && o.Account != nil {
		return true
	}

	return false
}

// SetAccount gets a reference to the given IamAccountRelationship and assigns it to the Account field.
func (o *IamDomainGroupAllOf) SetAccount(v IamAccountRelationship) {
	o.Account = &v
}

func (o IamDomainGroupAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.Name != nil {
		toSerialize["Name"] = o.Name
	}
	if o.Partition1 != nil {
		toSerialize["Partition1"] = o.Partition1
	}
	if o.Partition2 != nil {
		toSerialize["Partition2"] = o.Partition2
	}
	if o.Partition3 != nil {
		toSerialize["Partition3"] = o.Partition3
	}
	if o.PartitionKey != nil {
		toSerialize["PartitionKey"] = o.PartitionKey
	}
	if o.Usage != nil {
		toSerialize["Usage"] = o.Usage
	}
	if o.Account != nil {
		toSerialize["Account"] = o.Account
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IamDomainGroupAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varIamDomainGroupAllOf := _IamDomainGroupAllOf{}

	if err = json.Unmarshal(bytes, &varIamDomainGroupAllOf); err == nil {
		*o = IamDomainGroupAllOf(varIamDomainGroupAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "Name")
		delete(additionalProperties, "Partition1")
		delete(additionalProperties, "Partition2")
		delete(additionalProperties, "Partition3")
		delete(additionalProperties, "PartitionKey")
		delete(additionalProperties, "Usage")
		delete(additionalProperties, "Account")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIamDomainGroupAllOf struct {
	value *IamDomainGroupAllOf
	isSet bool
}

func (v NullableIamDomainGroupAllOf) Get() *IamDomainGroupAllOf {
	return v.value
}

func (v *NullableIamDomainGroupAllOf) Set(val *IamDomainGroupAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIamDomainGroupAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIamDomainGroupAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIamDomainGroupAllOf(val *IamDomainGroupAllOf) *NullableIamDomainGroupAllOf {
	return &NullableIamDomainGroupAllOf{value: val, isSet: true}
}

func (v NullableIamDomainGroupAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIamDomainGroupAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
