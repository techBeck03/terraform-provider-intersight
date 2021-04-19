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

// KubernetesStatefulSetStatusAllOf Definition of the list of properties defined in 'kubernetes.StatefulSetStatus', excluding properties defined in parent classes.
type KubernetesStatefulSetStatusAllOf struct {
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property.
	ObjectType string `json:"ObjectType"`
	// AvailableReplicas indicates the current avaliable replicas running.
	AvailableReplicas *int64 `json:"AvailableReplicas,omitempty"`
	// CollisionCount is the count of hash collisions for the StatefulSet. The StatefulSet controller uses this field as a collision avoidance mechanism when it needs to create the name for the newest ControllerRevision.
	CollisionCount *int64 `json:"CollisionCount,omitempty"`
	// CurrentRevision, if not empty, indicates the version of the StatefulSet used to generate Pods.
	CurrentRevision *string `json:"CurrentRevision,omitempty"`
	// ObservedGeneration is the most recent generation observed for this StatefulSet. It corresponds to the StatefulSet's generation, which is updated on mutation by the API Server.
	ObservedGeneration *int64 `json:"ObservedGeneration,omitempty"`
	// ReadyReplicas is the number of Pods created by the StatefulSet controller that have a Ready Condition.
	ReadyReplicas *int64 `json:"ReadyReplicas,omitempty"`
	// Number of replicas the statefulset desired to have.
	Replicas *int64 `json:"Replicas,omitempty"`
	// UpdateRevision, if not empty, indicates the version of the StatefulSet used to generate the pods that have yet to be updated, i.e. pod number <replicas> - <updatedReplicas>, until pod number <replicas>.
	UpdateRevision *string `json:"UpdateRevision,omitempty"`
	// UpdatedReplicas is the number of Pods created by the StatefulSet controller from the StatefulSet version indicated by updateRevision.
	UpdatedReplicas      *int64 `json:"UpdatedReplicas,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _KubernetesStatefulSetStatusAllOf KubernetesStatefulSetStatusAllOf

// NewKubernetesStatefulSetStatusAllOf instantiates a new KubernetesStatefulSetStatusAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKubernetesStatefulSetStatusAllOf(classId string, objectType string) *KubernetesStatefulSetStatusAllOf {
	this := KubernetesStatefulSetStatusAllOf{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewKubernetesStatefulSetStatusAllOfWithDefaults instantiates a new KubernetesStatefulSetStatusAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKubernetesStatefulSetStatusAllOfWithDefaults() *KubernetesStatefulSetStatusAllOf {
	this := KubernetesStatefulSetStatusAllOf{}
	var classId string = "kubernetes.StatefulSetStatus"
	this.ClassId = classId
	var objectType string = "kubernetes.StatefulSetStatus"
	this.ObjectType = objectType
	return &this
}

// GetClassId returns the ClassId field value
func (o *KubernetesStatefulSetStatusAllOf) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *KubernetesStatefulSetStatusAllOf) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *KubernetesStatefulSetStatusAllOf) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *KubernetesStatefulSetStatusAllOf) SetObjectType(v string) {
	o.ObjectType = v
}

// GetAvailableReplicas returns the AvailableReplicas field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetAvailableReplicas() int64 {
	if o == nil || o.AvailableReplicas == nil {
		var ret int64
		return ret
	}
	return *o.AvailableReplicas
}

// GetAvailableReplicasOk returns a tuple with the AvailableReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetAvailableReplicasOk() (*int64, bool) {
	if o == nil || o.AvailableReplicas == nil {
		return nil, false
	}
	return o.AvailableReplicas, true
}

// HasAvailableReplicas returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasAvailableReplicas() bool {
	if o != nil && o.AvailableReplicas != nil {
		return true
	}

	return false
}

// SetAvailableReplicas gets a reference to the given int64 and assigns it to the AvailableReplicas field.
func (o *KubernetesStatefulSetStatusAllOf) SetAvailableReplicas(v int64) {
	o.AvailableReplicas = &v
}

// GetCollisionCount returns the CollisionCount field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetCollisionCount() int64 {
	if o == nil || o.CollisionCount == nil {
		var ret int64
		return ret
	}
	return *o.CollisionCount
}

// GetCollisionCountOk returns a tuple with the CollisionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetCollisionCountOk() (*int64, bool) {
	if o == nil || o.CollisionCount == nil {
		return nil, false
	}
	return o.CollisionCount, true
}

// HasCollisionCount returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasCollisionCount() bool {
	if o != nil && o.CollisionCount != nil {
		return true
	}

	return false
}

// SetCollisionCount gets a reference to the given int64 and assigns it to the CollisionCount field.
func (o *KubernetesStatefulSetStatusAllOf) SetCollisionCount(v int64) {
	o.CollisionCount = &v
}

// GetCurrentRevision returns the CurrentRevision field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetCurrentRevision() string {
	if o == nil || o.CurrentRevision == nil {
		var ret string
		return ret
	}
	return *o.CurrentRevision
}

// GetCurrentRevisionOk returns a tuple with the CurrentRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetCurrentRevisionOk() (*string, bool) {
	if o == nil || o.CurrentRevision == nil {
		return nil, false
	}
	return o.CurrentRevision, true
}

// HasCurrentRevision returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasCurrentRevision() bool {
	if o != nil && o.CurrentRevision != nil {
		return true
	}

	return false
}

// SetCurrentRevision gets a reference to the given string and assigns it to the CurrentRevision field.
func (o *KubernetesStatefulSetStatusAllOf) SetCurrentRevision(v string) {
	o.CurrentRevision = &v
}

// GetObservedGeneration returns the ObservedGeneration field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetObservedGeneration() int64 {
	if o == nil || o.ObservedGeneration == nil {
		var ret int64
		return ret
	}
	return *o.ObservedGeneration
}

// GetObservedGenerationOk returns a tuple with the ObservedGeneration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetObservedGenerationOk() (*int64, bool) {
	if o == nil || o.ObservedGeneration == nil {
		return nil, false
	}
	return o.ObservedGeneration, true
}

// HasObservedGeneration returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasObservedGeneration() bool {
	if o != nil && o.ObservedGeneration != nil {
		return true
	}

	return false
}

// SetObservedGeneration gets a reference to the given int64 and assigns it to the ObservedGeneration field.
func (o *KubernetesStatefulSetStatusAllOf) SetObservedGeneration(v int64) {
	o.ObservedGeneration = &v
}

// GetReadyReplicas returns the ReadyReplicas field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetReadyReplicas() int64 {
	if o == nil || o.ReadyReplicas == nil {
		var ret int64
		return ret
	}
	return *o.ReadyReplicas
}

// GetReadyReplicasOk returns a tuple with the ReadyReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetReadyReplicasOk() (*int64, bool) {
	if o == nil || o.ReadyReplicas == nil {
		return nil, false
	}
	return o.ReadyReplicas, true
}

// HasReadyReplicas returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasReadyReplicas() bool {
	if o != nil && o.ReadyReplicas != nil {
		return true
	}

	return false
}

// SetReadyReplicas gets a reference to the given int64 and assigns it to the ReadyReplicas field.
func (o *KubernetesStatefulSetStatusAllOf) SetReadyReplicas(v int64) {
	o.ReadyReplicas = &v
}

// GetReplicas returns the Replicas field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetReplicas() int64 {
	if o == nil || o.Replicas == nil {
		var ret int64
		return ret
	}
	return *o.Replicas
}

// GetReplicasOk returns a tuple with the Replicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetReplicasOk() (*int64, bool) {
	if o == nil || o.Replicas == nil {
		return nil, false
	}
	return o.Replicas, true
}

// HasReplicas returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasReplicas() bool {
	if o != nil && o.Replicas != nil {
		return true
	}

	return false
}

// SetReplicas gets a reference to the given int64 and assigns it to the Replicas field.
func (o *KubernetesStatefulSetStatusAllOf) SetReplicas(v int64) {
	o.Replicas = &v
}

// GetUpdateRevision returns the UpdateRevision field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetUpdateRevision() string {
	if o == nil || o.UpdateRevision == nil {
		var ret string
		return ret
	}
	return *o.UpdateRevision
}

// GetUpdateRevisionOk returns a tuple with the UpdateRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetUpdateRevisionOk() (*string, bool) {
	if o == nil || o.UpdateRevision == nil {
		return nil, false
	}
	return o.UpdateRevision, true
}

// HasUpdateRevision returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasUpdateRevision() bool {
	if o != nil && o.UpdateRevision != nil {
		return true
	}

	return false
}

// SetUpdateRevision gets a reference to the given string and assigns it to the UpdateRevision field.
func (o *KubernetesStatefulSetStatusAllOf) SetUpdateRevision(v string) {
	o.UpdateRevision = &v
}

// GetUpdatedReplicas returns the UpdatedReplicas field value if set, zero value otherwise.
func (o *KubernetesStatefulSetStatusAllOf) GetUpdatedReplicas() int64 {
	if o == nil || o.UpdatedReplicas == nil {
		var ret int64
		return ret
	}
	return *o.UpdatedReplicas
}

// GetUpdatedReplicasOk returns a tuple with the UpdatedReplicas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *KubernetesStatefulSetStatusAllOf) GetUpdatedReplicasOk() (*int64, bool) {
	if o == nil || o.UpdatedReplicas == nil {
		return nil, false
	}
	return o.UpdatedReplicas, true
}

// HasUpdatedReplicas returns a boolean if a field has been set.
func (o *KubernetesStatefulSetStatusAllOf) HasUpdatedReplicas() bool {
	if o != nil && o.UpdatedReplicas != nil {
		return true
	}

	return false
}

// SetUpdatedReplicas gets a reference to the given int64 and assigns it to the UpdatedReplicas field.
func (o *KubernetesStatefulSetStatusAllOf) SetUpdatedReplicas(v int64) {
	o.UpdatedReplicas = &v
}

func (o KubernetesStatefulSetStatusAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.AvailableReplicas != nil {
		toSerialize["AvailableReplicas"] = o.AvailableReplicas
	}
	if o.CollisionCount != nil {
		toSerialize["CollisionCount"] = o.CollisionCount
	}
	if o.CurrentRevision != nil {
		toSerialize["CurrentRevision"] = o.CurrentRevision
	}
	if o.ObservedGeneration != nil {
		toSerialize["ObservedGeneration"] = o.ObservedGeneration
	}
	if o.ReadyReplicas != nil {
		toSerialize["ReadyReplicas"] = o.ReadyReplicas
	}
	if o.Replicas != nil {
		toSerialize["Replicas"] = o.Replicas
	}
	if o.UpdateRevision != nil {
		toSerialize["UpdateRevision"] = o.UpdateRevision
	}
	if o.UpdatedReplicas != nil {
		toSerialize["UpdatedReplicas"] = o.UpdatedReplicas
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *KubernetesStatefulSetStatusAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varKubernetesStatefulSetStatusAllOf := _KubernetesStatefulSetStatusAllOf{}

	if err = json.Unmarshal(bytes, &varKubernetesStatefulSetStatusAllOf); err == nil {
		*o = KubernetesStatefulSetStatusAllOf(varKubernetesStatefulSetStatusAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "AvailableReplicas")
		delete(additionalProperties, "CollisionCount")
		delete(additionalProperties, "CurrentRevision")
		delete(additionalProperties, "ObservedGeneration")
		delete(additionalProperties, "ReadyReplicas")
		delete(additionalProperties, "Replicas")
		delete(additionalProperties, "UpdateRevision")
		delete(additionalProperties, "UpdatedReplicas")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableKubernetesStatefulSetStatusAllOf struct {
	value *KubernetesStatefulSetStatusAllOf
	isSet bool
}

func (v NullableKubernetesStatefulSetStatusAllOf) Get() *KubernetesStatefulSetStatusAllOf {
	return v.value
}

func (v *NullableKubernetesStatefulSetStatusAllOf) Set(val *KubernetesStatefulSetStatusAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableKubernetesStatefulSetStatusAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableKubernetesStatefulSetStatusAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKubernetesStatefulSetStatusAllOf(val *KubernetesStatefulSetStatusAllOf) *NullableKubernetesStatefulSetStatusAllOf {
	return &NullableKubernetesStatefulSetStatusAllOf{value: val, isSet: true}
}

func (v NullableKubernetesStatefulSetStatusAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKubernetesStatefulSetStatusAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
