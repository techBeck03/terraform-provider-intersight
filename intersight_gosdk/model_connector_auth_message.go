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

// ConnectorAuthMessage A base abstract message for connector messages that require authentication to be passed from the Intersight services.
type ConnectorAuthMessage struct {
	ConnectorBaseMessage
	// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ClassId string `json:"ClassId"`
	// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
	ObjectType string `json:"ObjectType"`
	// The platform locale to assign user. A locale defines one or more organizations (domains) the user is allowed access, and access is limited to the organizations specified in the locale.
	RemoteUserLocale *string `json:"RemoteUserLocale,omitempty"`
	// The user name passed to the platform for use in platform audit logs.
	RemoteUserName *string `json:"RemoteUserName,omitempty"`
	// The list of roles to pass to the platform to validate the action against.
	RemoteUserRoles *string `json:"RemoteUserRoles,omitempty"`
	// The session Id passed to the platform for use in platforms auditing.
	RemoteUserSessionId  *string `json:"RemoteUserSessionId,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ConnectorAuthMessage ConnectorAuthMessage

// NewConnectorAuthMessage instantiates a new ConnectorAuthMessage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectorAuthMessage(classId string, objectType string) *ConnectorAuthMessage {
	this := ConnectorAuthMessage{}
	this.ClassId = classId
	this.ObjectType = objectType
	return &this
}

// NewConnectorAuthMessageWithDefaults instantiates a new ConnectorAuthMessage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectorAuthMessageWithDefaults() *ConnectorAuthMessage {
	this := ConnectorAuthMessage{}
	return &this
}

// GetClassId returns the ClassId field value
func (o *ConnectorAuthMessage) GetClassId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClassId
}

// GetClassIdOk returns a tuple with the ClassId field value
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetClassIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClassId, true
}

// SetClassId sets field value
func (o *ConnectorAuthMessage) SetClassId(v string) {
	o.ClassId = v
}

// GetObjectType returns the ObjectType field value
func (o *ConnectorAuthMessage) GetObjectType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetObjectTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectType, true
}

// SetObjectType sets field value
func (o *ConnectorAuthMessage) SetObjectType(v string) {
	o.ObjectType = v
}

// GetRemoteUserLocale returns the RemoteUserLocale field value if set, zero value otherwise.
func (o *ConnectorAuthMessage) GetRemoteUserLocale() string {
	if o == nil || o.RemoteUserLocale == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserLocale
}

// GetRemoteUserLocaleOk returns a tuple with the RemoteUserLocale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetRemoteUserLocaleOk() (*string, bool) {
	if o == nil || o.RemoteUserLocale == nil {
		return nil, false
	}
	return o.RemoteUserLocale, true
}

// HasRemoteUserLocale returns a boolean if a field has been set.
func (o *ConnectorAuthMessage) HasRemoteUserLocale() bool {
	if o != nil && o.RemoteUserLocale != nil {
		return true
	}

	return false
}

// SetRemoteUserLocale gets a reference to the given string and assigns it to the RemoteUserLocale field.
func (o *ConnectorAuthMessage) SetRemoteUserLocale(v string) {
	o.RemoteUserLocale = &v
}

// GetRemoteUserName returns the RemoteUserName field value if set, zero value otherwise.
func (o *ConnectorAuthMessage) GetRemoteUserName() string {
	if o == nil || o.RemoteUserName == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserName
}

// GetRemoteUserNameOk returns a tuple with the RemoteUserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetRemoteUserNameOk() (*string, bool) {
	if o == nil || o.RemoteUserName == nil {
		return nil, false
	}
	return o.RemoteUserName, true
}

// HasRemoteUserName returns a boolean if a field has been set.
func (o *ConnectorAuthMessage) HasRemoteUserName() bool {
	if o != nil && o.RemoteUserName != nil {
		return true
	}

	return false
}

// SetRemoteUserName gets a reference to the given string and assigns it to the RemoteUserName field.
func (o *ConnectorAuthMessage) SetRemoteUserName(v string) {
	o.RemoteUserName = &v
}

// GetRemoteUserRoles returns the RemoteUserRoles field value if set, zero value otherwise.
func (o *ConnectorAuthMessage) GetRemoteUserRoles() string {
	if o == nil || o.RemoteUserRoles == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserRoles
}

// GetRemoteUserRolesOk returns a tuple with the RemoteUserRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetRemoteUserRolesOk() (*string, bool) {
	if o == nil || o.RemoteUserRoles == nil {
		return nil, false
	}
	return o.RemoteUserRoles, true
}

// HasRemoteUserRoles returns a boolean if a field has been set.
func (o *ConnectorAuthMessage) HasRemoteUserRoles() bool {
	if o != nil && o.RemoteUserRoles != nil {
		return true
	}

	return false
}

// SetRemoteUserRoles gets a reference to the given string and assigns it to the RemoteUserRoles field.
func (o *ConnectorAuthMessage) SetRemoteUserRoles(v string) {
	o.RemoteUserRoles = &v
}

// GetRemoteUserSessionId returns the RemoteUserSessionId field value if set, zero value otherwise.
func (o *ConnectorAuthMessage) GetRemoteUserSessionId() string {
	if o == nil || o.RemoteUserSessionId == nil {
		var ret string
		return ret
	}
	return *o.RemoteUserSessionId
}

// GetRemoteUserSessionIdOk returns a tuple with the RemoteUserSessionId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectorAuthMessage) GetRemoteUserSessionIdOk() (*string, bool) {
	if o == nil || o.RemoteUserSessionId == nil {
		return nil, false
	}
	return o.RemoteUserSessionId, true
}

// HasRemoteUserSessionId returns a boolean if a field has been set.
func (o *ConnectorAuthMessage) HasRemoteUserSessionId() bool {
	if o != nil && o.RemoteUserSessionId != nil {
		return true
	}

	return false
}

// SetRemoteUserSessionId gets a reference to the given string and assigns it to the RemoteUserSessionId field.
func (o *ConnectorAuthMessage) SetRemoteUserSessionId(v string) {
	o.RemoteUserSessionId = &v
}

func (o ConnectorAuthMessage) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedConnectorBaseMessage, errConnectorBaseMessage := json.Marshal(o.ConnectorBaseMessage)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	errConnectorBaseMessage = json.Unmarshal([]byte(serializedConnectorBaseMessage), &toSerialize)
	if errConnectorBaseMessage != nil {
		return []byte{}, errConnectorBaseMessage
	}
	if true {
		toSerialize["ClassId"] = o.ClassId
	}
	if true {
		toSerialize["ObjectType"] = o.ObjectType
	}
	if o.RemoteUserLocale != nil {
		toSerialize["RemoteUserLocale"] = o.RemoteUserLocale
	}
	if o.RemoteUserName != nil {
		toSerialize["RemoteUserName"] = o.RemoteUserName
	}
	if o.RemoteUserRoles != nil {
		toSerialize["RemoteUserRoles"] = o.RemoteUserRoles
	}
	if o.RemoteUserSessionId != nil {
		toSerialize["RemoteUserSessionId"] = o.RemoteUserSessionId
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ConnectorAuthMessage) UnmarshalJSON(bytes []byte) (err error) {
	type ConnectorAuthMessageWithoutEmbeddedStruct struct {
		// The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ClassId string `json:"ClassId"`
		// The fully-qualified name of the instantiated, concrete type. The value should be the same as the 'ClassId' property. The enum values provides the list of concrete types that can be instantiated from this abstract type.
		ObjectType string `json:"ObjectType"`
		// The platform locale to assign user. A locale defines one or more organizations (domains) the user is allowed access, and access is limited to the organizations specified in the locale.
		RemoteUserLocale *string `json:"RemoteUserLocale,omitempty"`
		// The user name passed to the platform for use in platform audit logs.
		RemoteUserName *string `json:"RemoteUserName,omitempty"`
		// The list of roles to pass to the platform to validate the action against.
		RemoteUserRoles *string `json:"RemoteUserRoles,omitempty"`
		// The session Id passed to the platform for use in platforms auditing.
		RemoteUserSessionId *string `json:"RemoteUserSessionId,omitempty"`
	}

	varConnectorAuthMessageWithoutEmbeddedStruct := ConnectorAuthMessageWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varConnectorAuthMessageWithoutEmbeddedStruct)
	if err == nil {
		varConnectorAuthMessage := _ConnectorAuthMessage{}
		varConnectorAuthMessage.ClassId = varConnectorAuthMessageWithoutEmbeddedStruct.ClassId
		varConnectorAuthMessage.ObjectType = varConnectorAuthMessageWithoutEmbeddedStruct.ObjectType
		varConnectorAuthMessage.RemoteUserLocale = varConnectorAuthMessageWithoutEmbeddedStruct.RemoteUserLocale
		varConnectorAuthMessage.RemoteUserName = varConnectorAuthMessageWithoutEmbeddedStruct.RemoteUserName
		varConnectorAuthMessage.RemoteUserRoles = varConnectorAuthMessageWithoutEmbeddedStruct.RemoteUserRoles
		varConnectorAuthMessage.RemoteUserSessionId = varConnectorAuthMessageWithoutEmbeddedStruct.RemoteUserSessionId
		*o = ConnectorAuthMessage(varConnectorAuthMessage)
	} else {
		return err
	}

	varConnectorAuthMessage := _ConnectorAuthMessage{}

	err = json.Unmarshal(bytes, &varConnectorAuthMessage)
	if err == nil {
		o.ConnectorBaseMessage = varConnectorAuthMessage.ConnectorBaseMessage
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "ClassId")
		delete(additionalProperties, "ObjectType")
		delete(additionalProperties, "RemoteUserLocale")
		delete(additionalProperties, "RemoteUserName")
		delete(additionalProperties, "RemoteUserRoles")
		delete(additionalProperties, "RemoteUserSessionId")

		// remove fields from embedded structs
		reflectConnectorBaseMessage := reflect.ValueOf(o.ConnectorBaseMessage)
		for i := 0; i < reflectConnectorBaseMessage.Type().NumField(); i++ {
			t := reflectConnectorBaseMessage.Type().Field(i)

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

type NullableConnectorAuthMessage struct {
	value *ConnectorAuthMessage
	isSet bool
}

func (v NullableConnectorAuthMessage) Get() *ConnectorAuthMessage {
	return v.value
}

func (v *NullableConnectorAuthMessage) Set(val *ConnectorAuthMessage) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectorAuthMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectorAuthMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectorAuthMessage(val *ConnectorAuthMessage) *NullableConnectorAuthMessage {
	return &NullableConnectorAuthMessage{value: val, isSet: true}
}

func (v NullableConnectorAuthMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectorAuthMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
