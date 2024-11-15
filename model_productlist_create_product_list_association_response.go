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

// checks if the ProductlistCreateProductListAssociationResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistCreateProductListAssociationResponse{}

// ProductlistCreateProductListAssociationResponse struct for ProductlistCreateProductListAssociationResponse
type ProductlistCreateProductListAssociationResponse struct {
	Association *ProductlistProductListAssociation `json:"association,omitempty"`
	Errors []ProductlistProductListAssociationError `json:"errors,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistCreateProductListAssociationResponse ProductlistCreateProductListAssociationResponse

// NewProductlistCreateProductListAssociationResponse instantiates a new ProductlistCreateProductListAssociationResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistCreateProductListAssociationResponse() *ProductlistCreateProductListAssociationResponse {
	this := ProductlistCreateProductListAssociationResponse{}
	return &this
}

// NewProductlistCreateProductListAssociationResponseWithDefaults instantiates a new ProductlistCreateProductListAssociationResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistCreateProductListAssociationResponseWithDefaults() *ProductlistCreateProductListAssociationResponse {
	this := ProductlistCreateProductListAssociationResponse{}
	return &this
}

// GetAssociation returns the Association field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationResponse) GetAssociation() ProductlistProductListAssociation {
	if o == nil || IsNil(o.Association) {
		var ret ProductlistProductListAssociation
		return ret
	}
	return *o.Association
}

// GetAssociationOk returns a tuple with the Association field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationResponse) GetAssociationOk() (*ProductlistProductListAssociation, bool) {
	if o == nil || IsNil(o.Association) {
		return nil, false
	}
	return o.Association, true
}

// HasAssociation returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationResponse) HasAssociation() bool {
	if o != nil && !IsNil(o.Association) {
		return true
	}

	return false
}

// SetAssociation gets a reference to the given ProductlistProductListAssociation and assigns it to the Association field.
func (o *ProductlistCreateProductListAssociationResponse) SetAssociation(v ProductlistProductListAssociation) {
	o.Association = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *ProductlistCreateProductListAssociationResponse) GetErrors() []ProductlistProductListAssociationError {
	if o == nil || IsNil(o.Errors) {
		var ret []ProductlistProductListAssociationError
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistCreateProductListAssociationResponse) GetErrorsOk() ([]ProductlistProductListAssociationError, bool) {
	if o == nil || IsNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *ProductlistCreateProductListAssociationResponse) HasErrors() bool {
	if o != nil && !IsNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []ProductlistProductListAssociationError and assigns it to the Errors field.
func (o *ProductlistCreateProductListAssociationResponse) SetErrors(v []ProductlistProductListAssociationError) {
	o.Errors = v
}

func (o ProductlistCreateProductListAssociationResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistCreateProductListAssociationResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Association) {
		toSerialize["association"] = o.Association
	}
	if !IsNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistCreateProductListAssociationResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistCreateProductListAssociationResponse := _ProductlistCreateProductListAssociationResponse{}

	err = json.Unmarshal(data, &varProductlistCreateProductListAssociationResponse)

	if err != nil {
		return err
	}

	*o = ProductlistCreateProductListAssociationResponse(varProductlistCreateProductListAssociationResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "association")
		delete(additionalProperties, "errors")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistCreateProductListAssociationResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistCreateProductListAssociationResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistCreateProductListAssociationResponse struct {
	value *ProductlistCreateProductListAssociationResponse
	isSet bool
}

func (v NullableProductlistCreateProductListAssociationResponse) Get() *ProductlistCreateProductListAssociationResponse {
	return v.value
}

func (v *NullableProductlistCreateProductListAssociationResponse) Set(val *ProductlistCreateProductListAssociationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistCreateProductListAssociationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistCreateProductListAssociationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistCreateProductListAssociationResponse(val *ProductlistCreateProductListAssociationResponse) *NullableProductlistCreateProductListAssociationResponse {
	return &NullableProductlistCreateProductListAssociationResponse{value: val, isSet: true}
}

func (v NullableProductlistCreateProductListAssociationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistCreateProductListAssociationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


