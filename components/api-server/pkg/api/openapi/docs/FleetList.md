# FleetList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | Pointer to **string** |  | [optional] 
**Page** | Pointer to **int32** |  | [optional] 
**Size** | Pointer to **int32** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Href** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] 
**Items** | Pointer to [**[]Fleet**](Fleet.md) |  | [optional] 

## Methods

### NewFleetList

`func NewFleetList() *FleetList`

NewFleetList instantiates a new FleetList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFleetListWithDefaults

`func NewFleetListWithDefaults() *FleetList`

NewFleetListWithDefaults instantiates a new FleetList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *FleetList) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *FleetList) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *FleetList) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *FleetList) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPage

`func (o *FleetList) GetPage() int32`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *FleetList) GetPageOk() (*int32, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *FleetList) SetPage(v int32)`

SetPage sets Page field to given value.

### HasPage

`func (o *FleetList) HasPage() bool`

HasPage returns a boolean if a field has been set.

### GetSize

`func (o *FleetList) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *FleetList) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *FleetList) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *FleetList) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTotal

`func (o *FleetList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FleetList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FleetList) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FleetList) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetId

`func (o *FleetList) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FleetList) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FleetList) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FleetList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetHref

`func (o *FleetList) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *FleetList) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *FleetList) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *FleetList) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetCreatedAt

`func (o *FleetList) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *FleetList) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *FleetList) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *FleetList) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *FleetList) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *FleetList) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *FleetList) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *FleetList) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetItems

`func (o *FleetList) GetItems() []Fleet`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *FleetList) GetItemsOk() (*[]Fleet, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *FleetList) SetItems(v []Fleet)`

SetItems sets Items field to given value.

### HasItems

`func (o *FleetList) HasItems() bool`

HasItems returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


