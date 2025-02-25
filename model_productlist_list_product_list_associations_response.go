/*
Collection Service

The Collection Service provides a comprehensive API for managing product collections and categories. It supports the creation, retrieval, updating, and deletion of product collections, as well as managing associations between collections and products. Key functionalities include: - **Collection Management**: Create, update, and delete collections of products, enabling businesses to organize products into meaningful groups for easier browsing and marketing. - **Retrieval Operations**: Retrieve collections using various identifiers such as ID, URL key, or code, and list all existing collections with optional search and filter capabilities. - **Association Management**: Establish and manage associations between collections and products, allowing products to belong to one or multiple collections. - **Bulk Operations**: Perform bulk updates on product collection associations, optimizing the management of large-scale collections. - **Count and Search**: Get counts of products in a collection and search collections using filters or product identifiers for efficient management. - **Non-Hierarchical Structure**: While the service supports flat collections and categories, it does not yet include hierarchical relationships between categories. Designed for flexibility and scalability, the Collection Service empowers businesses to organize and present their product offerings effectively while maintaining a streamlined management process.

API version: v1
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package productlist

import (
	"encoding/json"
)

// checks if the ProductlistListProductListAssociationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProductlistListProductListAssociationsResponse{}

// ProductlistListProductListAssociationsResponse Defines the structure for ListProductListAssociationsResponse.
type ProductlistListProductListAssociationsResponse struct {
	// Field associations of type repeated ProductListAssociation.
	Associations []ProductlistProductListAssociation `json:"associations,omitempty"`
	// Field next_page_token of type string.
	NextPageToken        *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ProductlistListProductListAssociationsResponse ProductlistListProductListAssociationsResponse

// NewProductlistListProductListAssociationsResponse instantiates a new ProductlistListProductListAssociationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProductlistListProductListAssociationsResponse() *ProductlistListProductListAssociationsResponse {
	this := ProductlistListProductListAssociationsResponse{}
	return &this
}

// NewProductlistListProductListAssociationsResponseWithDefaults instantiates a new ProductlistListProductListAssociationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProductlistListProductListAssociationsResponseWithDefaults() *ProductlistListProductListAssociationsResponse {
	this := ProductlistListProductListAssociationsResponse{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsResponse) GetAssociations() []ProductlistProductListAssociation {
	if o == nil || IsNil(o.Associations) {
		var ret []ProductlistProductListAssociation
		return ret
	}
	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsResponse) GetAssociationsOk() ([]ProductlistProductListAssociation, bool) {
	if o == nil || IsNil(o.Associations) {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsResponse) IsSetAssociations() bool {
	if o != nil && !IsNil(o.Associations) {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given []ProductlistProductListAssociation and assigns it to the Associations field.
func (o *ProductlistListProductListAssociationsResponse) SetAssociations(v []ProductlistProductListAssociation) {
	o.Associations = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *ProductlistListProductListAssociationsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProductlistListProductListAssociationsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *ProductlistListProductListAssociationsResponse) IsSetNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *ProductlistListProductListAssociationsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o ProductlistListProductListAssociationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProductlistListProductListAssociationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Associations) {
		toSerialize["associations"] = o.Associations
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *ProductlistListProductListAssociationsResponse) UnmarshalJSON(data []byte) (err error) {
	varProductlistListProductListAssociationsResponse := _ProductlistListProductListAssociationsResponse{}

	err = json.Unmarshal(data, &varProductlistListProductListAssociationsResponse)

	if err != nil {
		return err
	}

	*o = ProductlistListProductListAssociationsResponse(varProductlistListProductListAssociationsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "associations")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *ProductlistListProductListAssociationsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *ProductlistListProductListAssociationsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableProductlistListProductListAssociationsResponse struct {
	value *ProductlistListProductListAssociationsResponse
	isSet bool
}

func (v NullableProductlistListProductListAssociationsResponse) Get() *ProductlistListProductListAssociationsResponse {
	return v.value
}

func (v *NullableProductlistListProductListAssociationsResponse) Set(val *ProductlistListProductListAssociationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableProductlistListProductListAssociationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableProductlistListProductListAssociationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProductlistListProductListAssociationsResponse(val *ProductlistListProductListAssociationsResponse) *NullableProductlistListProductListAssociationsResponse {
	return &NullableProductlistListProductListAssociationsResponse{value: val, isSet: true}
}

func (v NullableProductlistListProductListAssociationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProductlistListProductListAssociationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
