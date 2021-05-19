# WorkflowTemplateEvaluationAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClassId** | **string** | The fully-qualified name of the instantiated, concrete type. This property is used as a discriminator to identify the type of the payload when marshaling and unmarshaling data. | [default to "workflow.TemplateEvaluation"]
**ObjectType** | **string** | The fully-qualified name of the instantiated, concrete type. The value should be the same as the &#39;ClassId&#39; property. | [default to "workflow.TemplateEvaluation"]
**ExpectedOutputDataType** | Pointer to [**NullableWorkflowBaseDataType**](workflow.BaseDataType.md) |  | [optional] 
**Output** | Pointer to **interface{}** | The output generated by the template execution. | [optional] [readonly] 
**PreviewTemplate** | Pointer to **string** | The generated template based on the stages provided in the request body. | [optional] [readonly] 
**Stages** | Pointer to [**[]TemplateTransformationStage**](TemplateTransformationStage.md) |  | [optional] 
**Values** | Pointer to **interface{}** | Values to be fed to the template for execution. | [optional] 

## Methods

### NewWorkflowTemplateEvaluationAllOf

`func NewWorkflowTemplateEvaluationAllOf(classId string, objectType string, ) *WorkflowTemplateEvaluationAllOf`

NewWorkflowTemplateEvaluationAllOf instantiates a new WorkflowTemplateEvaluationAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTemplateEvaluationAllOfWithDefaults

`func NewWorkflowTemplateEvaluationAllOfWithDefaults() *WorkflowTemplateEvaluationAllOf`

NewWorkflowTemplateEvaluationAllOfWithDefaults instantiates a new WorkflowTemplateEvaluationAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClassId

`func (o *WorkflowTemplateEvaluationAllOf) GetClassId() string`

GetClassId returns the ClassId field if non-nil, zero value otherwise.

### GetClassIdOk

`func (o *WorkflowTemplateEvaluationAllOf) GetClassIdOk() (*string, bool)`

GetClassIdOk returns a tuple with the ClassId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClassId

`func (o *WorkflowTemplateEvaluationAllOf) SetClassId(v string)`

SetClassId sets ClassId field to given value.


### GetObjectType

`func (o *WorkflowTemplateEvaluationAllOf) GetObjectType() string`

GetObjectType returns the ObjectType field if non-nil, zero value otherwise.

### GetObjectTypeOk

`func (o *WorkflowTemplateEvaluationAllOf) GetObjectTypeOk() (*string, bool)`

GetObjectTypeOk returns a tuple with the ObjectType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectType

`func (o *WorkflowTemplateEvaluationAllOf) SetObjectType(v string)`

SetObjectType sets ObjectType field to given value.


### GetExpectedOutputDataType

`func (o *WorkflowTemplateEvaluationAllOf) GetExpectedOutputDataType() WorkflowBaseDataType`

GetExpectedOutputDataType returns the ExpectedOutputDataType field if non-nil, zero value otherwise.

### GetExpectedOutputDataTypeOk

`func (o *WorkflowTemplateEvaluationAllOf) GetExpectedOutputDataTypeOk() (*WorkflowBaseDataType, bool)`

GetExpectedOutputDataTypeOk returns a tuple with the ExpectedOutputDataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedOutputDataType

`func (o *WorkflowTemplateEvaluationAllOf) SetExpectedOutputDataType(v WorkflowBaseDataType)`

SetExpectedOutputDataType sets ExpectedOutputDataType field to given value.

### HasExpectedOutputDataType

`func (o *WorkflowTemplateEvaluationAllOf) HasExpectedOutputDataType() bool`

HasExpectedOutputDataType returns a boolean if a field has been set.

### SetExpectedOutputDataTypeNil

`func (o *WorkflowTemplateEvaluationAllOf) SetExpectedOutputDataTypeNil(b bool)`

 SetExpectedOutputDataTypeNil sets the value for ExpectedOutputDataType to be an explicit nil

### UnsetExpectedOutputDataType
`func (o *WorkflowTemplateEvaluationAllOf) UnsetExpectedOutputDataType()`

UnsetExpectedOutputDataType ensures that no value is present for ExpectedOutputDataType, not even an explicit nil
### GetOutput

`func (o *WorkflowTemplateEvaluationAllOf) GetOutput() interface{}`

GetOutput returns the Output field if non-nil, zero value otherwise.

### GetOutputOk

`func (o *WorkflowTemplateEvaluationAllOf) GetOutputOk() (*interface{}, bool)`

GetOutputOk returns a tuple with the Output field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutput

`func (o *WorkflowTemplateEvaluationAllOf) SetOutput(v interface{})`

SetOutput sets Output field to given value.

### HasOutput

`func (o *WorkflowTemplateEvaluationAllOf) HasOutput() bool`

HasOutput returns a boolean if a field has been set.

### SetOutputNil

`func (o *WorkflowTemplateEvaluationAllOf) SetOutputNil(b bool)`

 SetOutputNil sets the value for Output to be an explicit nil

### UnsetOutput
`func (o *WorkflowTemplateEvaluationAllOf) UnsetOutput()`

UnsetOutput ensures that no value is present for Output, not even an explicit nil
### GetPreviewTemplate

`func (o *WorkflowTemplateEvaluationAllOf) GetPreviewTemplate() string`

GetPreviewTemplate returns the PreviewTemplate field if non-nil, zero value otherwise.

### GetPreviewTemplateOk

`func (o *WorkflowTemplateEvaluationAllOf) GetPreviewTemplateOk() (*string, bool)`

GetPreviewTemplateOk returns a tuple with the PreviewTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewTemplate

`func (o *WorkflowTemplateEvaluationAllOf) SetPreviewTemplate(v string)`

SetPreviewTemplate sets PreviewTemplate field to given value.

### HasPreviewTemplate

`func (o *WorkflowTemplateEvaluationAllOf) HasPreviewTemplate() bool`

HasPreviewTemplate returns a boolean if a field has been set.

### GetStages

`func (o *WorkflowTemplateEvaluationAllOf) GetStages() []TemplateTransformationStage`

GetStages returns the Stages field if non-nil, zero value otherwise.

### GetStagesOk

`func (o *WorkflowTemplateEvaluationAllOf) GetStagesOk() (*[]TemplateTransformationStage, bool)`

GetStagesOk returns a tuple with the Stages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStages

`func (o *WorkflowTemplateEvaluationAllOf) SetStages(v []TemplateTransformationStage)`

SetStages sets Stages field to given value.

### HasStages

`func (o *WorkflowTemplateEvaluationAllOf) HasStages() bool`

HasStages returns a boolean if a field has been set.

### SetStagesNil

`func (o *WorkflowTemplateEvaluationAllOf) SetStagesNil(b bool)`

 SetStagesNil sets the value for Stages to be an explicit nil

### UnsetStages
`func (o *WorkflowTemplateEvaluationAllOf) UnsetStages()`

UnsetStages ensures that no value is present for Stages, not even an explicit nil
### GetValues

`func (o *WorkflowTemplateEvaluationAllOf) GetValues() interface{}`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkflowTemplateEvaluationAllOf) GetValuesOk() (*interface{}, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkflowTemplateEvaluationAllOf) SetValues(v interface{})`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkflowTemplateEvaluationAllOf) HasValues() bool`

HasValues returns a boolean if a field has been set.

### SetValuesNil

`func (o *WorkflowTemplateEvaluationAllOf) SetValuesNil(b bool)`

 SetValuesNil sets the value for Values to be an explicit nil

### UnsetValues
`func (o *WorkflowTemplateEvaluationAllOf) UnsetValues()`

UnsetValues ensures that no value is present for Values, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

