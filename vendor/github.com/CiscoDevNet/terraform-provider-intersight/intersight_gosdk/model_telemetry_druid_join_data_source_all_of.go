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

// TelemetryDruidJoinDataSourceAllOf struct for TelemetryDruidJoinDataSourceAllOf
type TelemetryDruidJoinDataSourceAllOf struct {
	// Left-hand datasource. Must be of type table, join, lookup, query, or inline. Placing another join as the left datasource allows you to join arbitrarily many datasources.
	Left string `json:"left"`
	// Right-hand datasource. Must be of type lookup, query, or inline.
	Right string `json:"right"`
	// String prefix that will be applied to all columns from the right-hand datasource, to prevent them from colliding with columns from the left-hand datasource. Can be any string, so long as it is nonempty and is not be a prefix of the string __time. Any columns from the left-hand side that start with your rightPrefix will be shadowed. It is up to you to provide a prefix that will not shadow any important columns from the left side.
	RightPrefix string `json:"rightPrefix"`
	// Expression that must be an equality where one side is an expression of the left-hand side, and the other side is a simple column reference to the right-hand side. The right-hand reference must be a simple column reference.
	Condition            string `json:"condition"`
	JoinType             string `json:"joinType"`
	AdditionalProperties map[string]interface{}
}

type _TelemetryDruidJoinDataSourceAllOf TelemetryDruidJoinDataSourceAllOf

// NewTelemetryDruidJoinDataSourceAllOf instantiates a new TelemetryDruidJoinDataSourceAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTelemetryDruidJoinDataSourceAllOf(left string, right string, rightPrefix string, condition string, joinType string) *TelemetryDruidJoinDataSourceAllOf {
	this := TelemetryDruidJoinDataSourceAllOf{}
	this.Left = left
	this.Right = right
	this.RightPrefix = rightPrefix
	this.Condition = condition
	this.JoinType = joinType
	return &this
}

// NewTelemetryDruidJoinDataSourceAllOfWithDefaults instantiates a new TelemetryDruidJoinDataSourceAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTelemetryDruidJoinDataSourceAllOfWithDefaults() *TelemetryDruidJoinDataSourceAllOf {
	this := TelemetryDruidJoinDataSourceAllOf{}
	return &this
}

// GetLeft returns the Left field value
func (o *TelemetryDruidJoinDataSourceAllOf) GetLeft() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Left
}

// GetLeftOk returns a tuple with the Left field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSourceAllOf) GetLeftOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Left, true
}

// SetLeft sets field value
func (o *TelemetryDruidJoinDataSourceAllOf) SetLeft(v string) {
	o.Left = v
}

// GetRight returns the Right field value
func (o *TelemetryDruidJoinDataSourceAllOf) GetRight() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Right
}

// GetRightOk returns a tuple with the Right field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSourceAllOf) GetRightOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Right, true
}

// SetRight sets field value
func (o *TelemetryDruidJoinDataSourceAllOf) SetRight(v string) {
	o.Right = v
}

// GetRightPrefix returns the RightPrefix field value
func (o *TelemetryDruidJoinDataSourceAllOf) GetRightPrefix() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RightPrefix
}

// GetRightPrefixOk returns a tuple with the RightPrefix field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSourceAllOf) GetRightPrefixOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RightPrefix, true
}

// SetRightPrefix sets field value
func (o *TelemetryDruidJoinDataSourceAllOf) SetRightPrefix(v string) {
	o.RightPrefix = v
}

// GetCondition returns the Condition field value
func (o *TelemetryDruidJoinDataSourceAllOf) GetCondition() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Condition
}

// GetConditionOk returns a tuple with the Condition field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSourceAllOf) GetConditionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Condition, true
}

// SetCondition sets field value
func (o *TelemetryDruidJoinDataSourceAllOf) SetCondition(v string) {
	o.Condition = v
}

// GetJoinType returns the JoinType field value
func (o *TelemetryDruidJoinDataSourceAllOf) GetJoinType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.JoinType
}

// GetJoinTypeOk returns a tuple with the JoinType field value
// and a boolean to check if the value has been set.
func (o *TelemetryDruidJoinDataSourceAllOf) GetJoinTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.JoinType, true
}

// SetJoinType sets field value
func (o *TelemetryDruidJoinDataSourceAllOf) SetJoinType(v string) {
	o.JoinType = v
}

func (o TelemetryDruidJoinDataSourceAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["left"] = o.Left
	}
	if true {
		toSerialize["right"] = o.Right
	}
	if true {
		toSerialize["rightPrefix"] = o.RightPrefix
	}
	if true {
		toSerialize["condition"] = o.Condition
	}
	if true {
		toSerialize["joinType"] = o.JoinType
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TelemetryDruidJoinDataSourceAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varTelemetryDruidJoinDataSourceAllOf := _TelemetryDruidJoinDataSourceAllOf{}

	if err = json.Unmarshal(bytes, &varTelemetryDruidJoinDataSourceAllOf); err == nil {
		*o = TelemetryDruidJoinDataSourceAllOf(varTelemetryDruidJoinDataSourceAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "left")
		delete(additionalProperties, "right")
		delete(additionalProperties, "rightPrefix")
		delete(additionalProperties, "condition")
		delete(additionalProperties, "joinType")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTelemetryDruidJoinDataSourceAllOf struct {
	value *TelemetryDruidJoinDataSourceAllOf
	isSet bool
}

func (v NullableTelemetryDruidJoinDataSourceAllOf) Get() *TelemetryDruidJoinDataSourceAllOf {
	return v.value
}

func (v *NullableTelemetryDruidJoinDataSourceAllOf) Set(val *TelemetryDruidJoinDataSourceAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTelemetryDruidJoinDataSourceAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTelemetryDruidJoinDataSourceAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTelemetryDruidJoinDataSourceAllOf(val *TelemetryDruidJoinDataSourceAllOf) *NullableTelemetryDruidJoinDataSourceAllOf {
	return &NullableTelemetryDruidJoinDataSourceAllOf{value: val, isSet: true}
}

func (v NullableTelemetryDruidJoinDataSourceAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTelemetryDruidJoinDataSourceAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
