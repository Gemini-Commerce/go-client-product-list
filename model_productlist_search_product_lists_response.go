/*
Collection Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistSearchProductListsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistSearchProductListsResponse{}

// ProductlistSearchProductListsResponse struct for ProductlistSearchProductListsResponse
type ProductlistSearchProductListsResponse struct {
	Results []ProductlistProductListEntity `json:"results,omitempty"`
	TotalSize *int64 `json:"totalSize,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistSearchProductListsResponse ProductlistSearchProductListsResponse

// NewProductlistSearchProductListsResponse instantiates a new ProductlistSearchProductListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistSearchProductListsResponse() *ProductlistSearchProductListsResponse {
	this := ProductlistSearchProductListsResponse{}
	return &this
}

// NewProductlistSearchProductListsResponseWithDefaults instantiates a new ProductlistSearchProductListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistSearchProductListsResponseWithDefaults() *ProductlistSearchProductListsResponse {
	this := ProductlistSearchProductListsResponse{}
	return &this
}

// GetResults returns the Results field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsResponse) GetResults() []ProductlistProductListEntity {
	if o == nil || IsNil(o.Results) {
		var ret []ProductlistProductListEntity
		return ret
	}
	return o.Results
}

// GetResultsOk returns a tuple with the Results field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsResponse) GetResultsOk() ([]ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.Results) {
		return nil, false
	}
	return o.Results, true
}

// HasResults returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsResponse) HasResults() bool {
	if o != nil && !IsNil(o.Results) {
		return true
	}

	return false
}

// SetResults gets a reference to the given []ProductlistProductListEntity and assigns it to the Results field.
func (o *ProductlistSearchProductListsResponse) SetResults(v []ProductlistProductListEntity) {
	o.Results = v
}

// GetTotalSize returns the TotalSize field value if set, zero value otherwise.
func (o *ProductlistSearchProductListsResponse) GetTotalSize() int64 {
	if o == nil || IsNil(o.TotalSize) {
		var ret int64
		return ret
	}
	return *o.TotalSize
}

// GetTotalSizeOk returns a tuple with the TotalSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistSearchProductListsResponse) GetTotalSizeOk() (*int64, bool) {
	if o == nil || IsNil(o.TotalSize) {
		return nil, false
	}
	return o.TotalSize, true
}

// HasTotalSize returns a boolean if a field has been set.
func (o *ProductlistSearchProductListsResponse) HasTotalSize() bool {
	if o != nil && !IsNil(o.TotalSize) {
		return true
	}

	return false
}

// SetTotalSize gets a reference to the given int64 and assigns it to the TotalSize field.
func (o *ProductlistSearchProductListsResponse) SetTotalSize(v int64) {
	o.TotalSize = &v
}

func (o ProductlistSearchProductListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistSearchProductListsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Results) {
		toSerialize["results"] = o.Results
	}
	if !IsNil(o.TotalSize) {
		toSerialize["totalSize"] = o.TotalSize
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistSearchProductListsResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistSearchProductListsResponse := _ProductlistSearchProductListsResponse{}

	err = json.Unmarshal(data, &varProductlistSearchProductListsResponse)

	if err != nil {
		return err
	}

	*o = ProductlistSearchProductListsResponse(varProductlistSearchProductListsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "results")
		delete(additionalProperties, "totalSize")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistSearchProductListsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistSearchProductListsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistSearchProductListsResponse struct {
	value *ProductlistSearchProductListsResponse
	isSet bool
}

func (v NullableProductlistSearchProductListsResponse) Get() *ProductlistSearchProductListsResponse {
	return v.value
}

func (v *NullableProductlistSearchProductListsResponse) Set(val *ProductlistSearchProductListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistSearchProductListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistSearchProductListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistSearchProductListsResponse(val *ProductlistSearchProductListsResponse) *NullableProductlistSearchProductListsResponse {
	return &NullableProductlistSearchProductListsResponse{value: val, isSet: true}
}

func (v NullableProductlistSearchProductListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistSearchProductListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


