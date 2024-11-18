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

// checks if the ProductlistListProductListsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistListProductListsResponse{}

// ProductlistListProductListsResponse struct for ProductlistListProductListsResponse
type ProductlistListProductListsResponse struct {
	Lists                []ProductlistProductListEntity `json:"lists,omitempty"`
	NextPageToken        *string                        `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistListProductListsResponse ProductlistListProductListsResponse

// NewProductlistListProductListsResponse instantiates a new ProductlistListProductListsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistListProductListsResponse() *ProductlistListProductListsResponse {
	this := ProductlistListProductListsResponse{}
	return &this
}

// NewProductlistListProductListsResponseWithDefaults instantiates a new ProductlistListProductListsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistListProductListsResponseWithDefaults() *ProductlistListProductListsResponse {
	this := ProductlistListProductListsResponse{}
	return &this
}

// GetLists returns the Lists field value if set, zero value otherwise.
func (o *ProductlistListProductListsResponse) GetLists() []ProductlistProductListEntity {
	if o == nil || IsNil(o.Lists) {
		var ret []ProductlistProductListEntity
		return ret
	}
	return o.Lists
}

// GetListsOk returns a tuple with the Lists field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListsResponse) GetListsOk() ([]ProductlistProductListEntity, bool) {
	if o == nil || IsNil(o.Lists) {
		return nil, false
	}
	return o.Lists, true
}

// HasLists returns a boolean if a field has been set.
func (o *ProductlistListProductListsResponse) HasLists() bool {
	if o != nil && !IsNil(o.Lists) {
		return true
	}

	return false
}

// SetLists gets a reference to the given []ProductlistProductListEntity and assigns it to the Lists field.
func (o *ProductlistListProductListsResponse) SetLists(v []ProductlistProductListEntity) {
	o.Lists = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ProductlistListProductListsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ProductlistListProductListsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ProductlistListProductListsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ProductlistListProductListsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistListProductListsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Lists) {
		toSerialize["lists"] = o.Lists
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistListProductListsResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistListProductListsResponse := _ProductlistListProductListsResponse{}

	err = json.Unmarshal(data, &varProductlistListProductListsResponse)

	if err != nil {
		return err
	}

	*o = ProductlistListProductListsResponse(varProductlistListProductListsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "lists")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistListProductListsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistListProductListsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistListProductListsResponse struct {
	value *ProductlistListProductListsResponse
	isSet bool
}

func (v NullableProductlistListProductListsResponse) Get() *ProductlistListProductListsResponse {
	return v.value
}

func (v *NullableProductlistListProductListsResponse) Set(val *ProductlistListProductListsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistListProductListsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistListProductListsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistListProductListsResponse(val *ProductlistListProductListsResponse) *NullableProductlistListProductListsResponse {
	return &NullableProductlistListProductListsResponse{value: val, isSet: true}
}

func (v NullableProductlistListProductListsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistListProductListsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
