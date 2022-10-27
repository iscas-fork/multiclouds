package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/v3/volumes"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateBlockstorageV3Volumes
type CreateBlockstorageV3VolumesRequest struct{
    Opts volumes.CreateOptsBuilder
}

func NewCreateBlockstorageV3VolumesRequest()*CreateBlockstorageV3VolumesRequest{
    return &CreateBlockstorageV3VolumesRequest{}
}

//response struct for the CreateBlockstorageV3Volumes
type CreateBlockstorageV3VolumesResponse struct{
    CreateResult volumes.CreateResult
}

func NewCreateBlockstorageV3VolumesResponse(createResult volumes.CreateResult,)*CreateBlockstorageV3VolumesResponse {
    return &CreateBlockstorageV3VolumesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateBlockstorageV3Volumes(request *CreateBlockstorageV3VolumesRequest)(*CreateBlockstorageV3VolumesResponse){
    return NewCreateBlockstorageV3VolumesResponse(volumes.Create(oc.client,request.Opts, ))

}
//request struct for the DeleteBlockstorageV3Volumes
type DeleteBlockstorageV3VolumesRequest struct{
    Id string
    Opts volumes.DeleteOptsBuilder
}

func NewDeleteBlockstorageV3VolumesRequest()*DeleteBlockstorageV3VolumesRequest{
    return &DeleteBlockstorageV3VolumesRequest{}
}

//response struct for the DeleteBlockstorageV3Volumes
type DeleteBlockstorageV3VolumesResponse struct{
    DeleteResult volumes.DeleteResult
}

func NewDeleteBlockstorageV3VolumesResponse(deleteResult volumes.DeleteResult,)*DeleteBlockstorageV3VolumesResponse {
    return &DeleteBlockstorageV3VolumesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteBlockstorageV3Volumes(request *DeleteBlockstorageV3VolumesRequest)(*DeleteBlockstorageV3VolumesResponse){
    return NewDeleteBlockstorageV3VolumesResponse(volumes.Delete(oc.client,request.Id,request.Opts, ))

}
//request struct for the GetBlockstorageV3Volumes
type GetBlockstorageV3VolumesRequest struct{
    Id string
}

func NewGetBlockstorageV3VolumesRequest()*GetBlockstorageV3VolumesRequest{
    return &GetBlockstorageV3VolumesRequest{}
}

//response struct for the GetBlockstorageV3Volumes
type GetBlockstorageV3VolumesResponse struct{
    GetResult volumes.GetResult
}

func NewGetBlockstorageV3VolumesResponse(getResult volumes.GetResult,)*GetBlockstorageV3VolumesResponse {
    return &GetBlockstorageV3VolumesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetBlockstorageV3Volumes(request *GetBlockstorageV3VolumesRequest)(*GetBlockstorageV3VolumesResponse){
    return NewGetBlockstorageV3VolumesResponse(volumes.Get(oc.client,request.Id, ))

}
//request struct for the ListBlockstorageV3Volumes
type ListBlockstorageV3VolumesRequest struct{
    Opts volumes.ListOptsBuilder
}

func NewListBlockstorageV3VolumesRequest()*ListBlockstorageV3VolumesRequest{
    return &ListBlockstorageV3VolumesRequest{}
}

//response struct for the ListBlockstorageV3Volumes
type ListBlockstorageV3VolumesResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageV3VolumesResponse(pager pagination.Pager,)*ListBlockstorageV3VolumesResponse {
    return &ListBlockstorageV3VolumesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageV3Volumes(request *ListBlockstorageV3VolumesRequest)(*ListBlockstorageV3VolumesResponse){
    return NewListBlockstorageV3VolumesResponse(volumes.List(oc.client,request.Opts, ))

}
//request struct for the UpdateBlockstorageV3Volumes
type UpdateBlockstorageV3VolumesRequest struct{
    Id string
    Opts volumes.UpdateOptsBuilder
}

func NewUpdateBlockstorageV3VolumesRequest()*UpdateBlockstorageV3VolumesRequest{
    return &UpdateBlockstorageV3VolumesRequest{}
}

//response struct for the UpdateBlockstorageV3Volumes
type UpdateBlockstorageV3VolumesResponse struct{
    UpdateResult volumes.UpdateResult
}

func NewUpdateBlockstorageV3VolumesResponse(updateResult volumes.UpdateResult,)*UpdateBlockstorageV3VolumesResponse {
    return &UpdateBlockstorageV3VolumesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateBlockstorageV3Volumes(request *UpdateBlockstorageV3VolumesRequest)(*UpdateBlockstorageV3VolumesResponse){
    return NewUpdateBlockstorageV3VolumesResponse(volumes.Update(oc.client,request.Id,request.Opts, ))

}