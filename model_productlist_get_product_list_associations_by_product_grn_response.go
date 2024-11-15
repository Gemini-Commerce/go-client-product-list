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

// checks if the ProductlistGetProductListAssociationsByProductGrnResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistGetProductListAssociationsByProductGrnResponse{}

// ProductlistGetProductListAssociationsByProductGrnResponse struct for ProductlistGetProductListAssociationsByProductGrnResponse
type ProductlistGetProductListAssociationsByProductGrnResponse struct {
	Associations []ProductlistProductListAssociation `json:"associations,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistGetProductListAssociationsByProductGrnResponse ProductlistGetProductListAssociationsByProductGrnResponse

// NewProductlistGetProductListAssociationsByProductGrnResponse instantiates a new ProductlistGetProductListAssociationsByProductGrnResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistGetProductListAssociationsByProductGrnResponse() *ProductlistGetProductListAssociationsByProductGrnResponse {
	this := ProductlistGetProductListAssociationsByProductGrnResponse{}
	return &this
}

// NewProductlistGetProductListAssociationsByProductGrnResponseWithDefaults instantiates a new ProductlistGetProductListAssociationsByProductGrnResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistGetProductListAssociationsByProductGrnResponseWithDefaults() *ProductlistGetProductListAssociationsByProductGrnResponse {
	this := ProductlistGetProductListAssociationsByProductGrnResponse{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) GetAssociations() []ProductlistProductListAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []ProductlistProductListAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) GetAssociationsOk() ([]ProductlistProductListAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) HasAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []ProductlistProductListAssociation and assigns it to the Associations field.
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) SetAssociations(v []ProductlistProductListAssociation) {
	o.Associations = v
}

func (o ProductlistGetProductListAssociationsByProductGrnResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistGetProductListAssociationsByProductGrnResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistGetProductListAssociationsByProductGrnResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistGetProductListAssociationsByProductGrnResponse := _ProductlistGetProductListAssociationsByProductGrnResponse{}

	err = json.Unmarshal(data, &varProductlistGetProductListAssociationsByProductGrnResponse)

	if err != nil {
		return err
	}

	*o = ProductlistGetProductListAssociationsByProductGrnResponse(varProductlistGetProductListAssociationsByProductGrnResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associations")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *ProductlistGetProductListAssociationsByProductGrnResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableProductlistGetProductListAssociationsByProductGrnResponse struct {
	value *ProductlistGetProductListAssociationsByProductGrnResponse
	isSet bool
}

func (v NullableProductlistGetProductListAssociationsByProductGrnResponse) Get() *ProductlistGetProductListAssociationsByProductGrnResponse {
	return v.value
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnResponse) Set(val *ProductlistGetProductListAssociationsByProductGrnResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistGetProductListAssociationsByProductGrnResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistGetProductListAssociationsByProductGrnResponse(val *ProductlistGetProductListAssociationsByProductGrnResponse) *NullableProductlistGetProductListAssociationsByProductGrnResponse {
	return &NullableProductlistGetProductListAssociationsByProductGrnResponse{value: val, isSet: true}
}

func (v NullableProductlistGetProductListAssociationsByProductGrnResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistGetProductListAssociationsByProductGrnResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


