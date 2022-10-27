package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v1/volumetypes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV1Volumetypes
type CreateBlockstorageV1VolumetypesRequest struct{
    Opts volumetypes.CreateOptsBuilder
}

func NewCreateBlockstorageV1VolumetypesRequest()*CreateBlockstorageV1VolumetypesRequest{
    return &CreateBlockstorageV1VolumetypesRequest{}
}

//response struct for the CreateBlockstorageV1Volumetypes
type CreateBlockstorageV1VolumetypesResponse struct{
    CreateResult volumetypes.CreateResult
}

func NewCreateBlockstorageV1VolumetypesResponse(createResult volumetypes.CreateResult,)*CreateBlockstorageV1VolumetypesResponse {
    return &CreateBlockstorageV1VolumetypesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV1Volumetypes(request *CreateBlockstorageV1VolumetypesRequest)(*CreateBlockstorageV1VolumetypesResponse){
    return NewCreateBlockstorageV1VolumetypesResponse(volumetypes.Create(oc.client,request.Opts, ))

}
//request struct for the DeleteBlockstorageV1Volumetypes
type DeleteBlockstorageV1VolumetypesRequest struct{
    Id string
}

func NewDeleteBlockstorageV1VolumetypesRequest()*DeleteBlockstorageV1VolumetypesRequest{
    return &DeleteBlockstorageV1VolumetypesRequest{}
}

//response struct for the DeleteBlockstorageV1Volumetypes
type DeleteBlockstorageV1VolumetypesResponse struct{
    DeleteResult volumetypes.DeleteResult
}

func NewDeleteBlockstorageV1VolumetypesResponse(deleteResult volumetypes.DeleteResult,)*DeleteBlockstorageV1VolumetypesResponse {
    return &DeleteBlockstorageV1VolumetypesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV1Volumetypes(request *DeleteBlockstorageV1VolumetypesRequest)(*DeleteBlockstorageV1VolumetypesResponse){
    return NewDeleteBlockstorageV1VolumetypesResponse(volumetypes.Delete(oc.client,request.Id, ))

}
//request struct for the GetBlockstorageV1Volumetypes
type GetBlockstorageV1VolumetypesRequest struct{
    Id string
}

func NewGetBlockstorageV1VolumetypesRequest()*GetBlockstorageV1VolumetypesRequest{
    return &GetBlockstorageV1VolumetypesRequest{}
}

//response struct for the GetBlockstorageV1Volumetypes
type GetBlockstorageV1VolumetypesResponse struct{
    GetResult volumetypes.GetResult
}

func NewGetBlockstorageV1VolumetypesResponse(getResult volumetypes.GetResult,)*GetBlockstorageV1VolumetypesResponse {
    return &GetBlockstorageV1VolumetypesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV1Volumetypes(request *GetBlockstorageV1VolumetypesRequest)(*GetBlockstorageV1VolumetypesResponse){
    return NewGetBlockstorageV1VolumetypesResponse(volumetypes.Get(oc.client,request.Id, ))

}
//request struct for the ListBlockstorageV1Volumetypes
type ListBlockstorageV1VolumetypesRequest struct{
}

func NewListBlockstorageV1VolumetypesRequest()*ListBlockstorageV1VolumetypesRequest{
    return &ListBlockstorageV1VolumetypesRequest{}
}

//response struct for the ListBlockstorageV1Volumetypes
type ListBlockstorageV1VolumetypesResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV1VolumetypesResponse(pager pagination.Pager,)*ListBlockstorageV1VolumetypesResponse {
    return &ListBlockstorageV1VolumetypesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV1Volumetypes(request *ListBlockstorageV1VolumetypesRequest)(*ListBlockstorageV1VolumetypesResponse){
    return NewListBlockstorageV1VolumetypesResponse(volumetypes.List(oc.client, ))

}