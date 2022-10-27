package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/objectstorage/v1/containers"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListObjectstorageV1Containers
type ListObjectstorageV1ContainersRequest struct{
    Opts containers.ListOptsBuilder
}

func NewListObjectstorageV1ContainersRequest()*ListObjectstorageV1ContainersRequest{
    return &ListObjectstorageV1ContainersRequest{}
}

//response struct for the ListObjectstorageV1Containers
type ListObjectstorageV1ContainersResponse struct{
    Pager pagination.Pager
}

func NewListObjectstorageV1ContainersResponse(pager pagination.Pager,)*ListObjectstorageV1ContainersResponse {
    return &ListObjectstorageV1ContainersResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListObjectstorageV1Containers(request *ListObjectstorageV1ContainersRequest)(*ListObjectstorageV1ContainersResponse){
    return NewListObjectstorageV1ContainersResponse(containers.List(oc.client,request.Opts, ))

}
//request struct for the CreateObjectstorageV1Containers
type CreateObjectstorageV1ContainersRequest struct{
    ContainerName string
    Opts containers.CreateOptsBuilder
}

func NewCreateObjectstorageV1ContainersRequest()*CreateObjectstorageV1ContainersRequest{
    return &CreateObjectstorageV1ContainersRequest{}
}

//response struct for the CreateObjectstorageV1Containers
type CreateObjectstorageV1ContainersResponse struct{
    CreateResult containers.CreateResult
}

func NewCreateObjectstorageV1ContainersResponse(createResult containers.CreateResult,)*CreateObjectstorageV1ContainersResponse {
    return &CreateObjectstorageV1ContainersResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateObjectstorageV1Containers(request *CreateObjectstorageV1ContainersRequest)(*CreateObjectstorageV1ContainersResponse){
    return NewCreateObjectstorageV1ContainersResponse(containers.Create(oc.client,request.ContainerName,request.Opts, ))

}
//request struct for the BulkDeleteObjectstorageV1Containers
type BulkDeleteObjectstorageV1ContainersRequest struct{
    Containers 
}

func NewBulkDeleteObjectstorageV1ContainersRequest()*BulkDeleteObjectstorageV1ContainersRequest{
    return &BulkDeleteObjectstorageV1ContainersRequest{}
}

//response struct for the BulkDeleteObjectstorageV1Containers
type BulkDeleteObjectstorageV1ContainersResponse struct{
    BulkDeleteResult containers.BulkDeleteResult
}

func NewBulkDeleteObjectstorageV1ContainersResponse(bulkDeleteResult containers.BulkDeleteResult,)*BulkDeleteObjectstorageV1ContainersResponse {
    return &BulkDeleteObjectstorageV1ContainersResponse{
            BulkDeleteResult:bulkDeleteResult,
    }
}

// action function
func (oc *OpenstackClient) BulkDeleteObjectstorageV1Containers(request *BulkDeleteObjectstorageV1ContainersRequest)(*BulkDeleteObjectstorageV1ContainersResponse){
    return NewBulkDeleteObjectstorageV1ContainersResponse(containers.BulkDelete(oc.client,request.Containers, ))

}
//request struct for the DeleteObjectstorageV1Containers
type DeleteObjectstorageV1ContainersRequest struct{
    ContainerName string
}

func NewDeleteObjectstorageV1ContainersRequest()*DeleteObjectstorageV1ContainersRequest{
    return &DeleteObjectstorageV1ContainersRequest{}
}

//response struct for the DeleteObjectstorageV1Containers
type DeleteObjectstorageV1ContainersResponse struct{
    DeleteResult containers.DeleteResult
}

func NewDeleteObjectstorageV1ContainersResponse(deleteResult containers.DeleteResult,)*DeleteObjectstorageV1ContainersResponse {
    return &DeleteObjectstorageV1ContainersResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteObjectstorageV1Containers(request *DeleteObjectstorageV1ContainersRequest)(*DeleteObjectstorageV1ContainersResponse){
    return NewDeleteObjectstorageV1ContainersResponse(containers.Delete(oc.client,request.ContainerName, ))

}
//request struct for the UpdateObjectstorageV1Containers
type UpdateObjectstorageV1ContainersRequest struct{
    ContainerName string
    Opts containers.UpdateOptsBuilder
}

func NewUpdateObjectstorageV1ContainersRequest()*UpdateObjectstorageV1ContainersRequest{
    return &UpdateObjectstorageV1ContainersRequest{}
}

//response struct for the UpdateObjectstorageV1Containers
type UpdateObjectstorageV1ContainersResponse struct{
    UpdateResult containers.UpdateResult
}

func NewUpdateObjectstorageV1ContainersResponse(updateResult containers.UpdateResult,)*UpdateObjectstorageV1ContainersResponse {
    return &UpdateObjectstorageV1ContainersResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateObjectstorageV1Containers(request *UpdateObjectstorageV1ContainersRequest)(*UpdateObjectstorageV1ContainersResponse){
    return NewUpdateObjectstorageV1ContainersResponse(containers.Update(oc.client,request.ContainerName,request.Opts, ))

}
//request struct for the GetObjectstorageV1Containers
type GetObjectstorageV1ContainersRequest struct{
    ContainerName string
    Opts containers.GetOptsBuilder
}

func NewGetObjectstorageV1ContainersRequest()*GetObjectstorageV1ContainersRequest{
    return &GetObjectstorageV1ContainersRequest{}
}

//response struct for the GetObjectstorageV1Containers
type GetObjectstorageV1ContainersResponse struct{
    GetResult containers.GetResult
}

func NewGetObjectstorageV1ContainersResponse(getResult containers.GetResult,)*GetObjectstorageV1ContainersResponse {
    return &GetObjectstorageV1ContainersResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetObjectstorageV1Containers(request *GetObjectstorageV1ContainersRequest)(*GetObjectstorageV1ContainersResponse){
    return NewGetObjectstorageV1ContainersResponse(containers.Get(oc.client,request.ContainerName,request.Opts, ))

}