package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/objectstorage/v1/objects"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListObjectstorageV1Objects
type ListObjectstorageV1ObjectsRequest struct{
    ContainerName string
    Opts objects.ListOptsBuilder
}

func NewListObjectstorageV1ObjectsRequest()*ListObjectstorageV1ObjectsRequest{
    return &ListObjectstorageV1ObjectsRequest{}
}

//response struct for the ListObjectstorageV1Objects
type ListObjectstorageV1ObjectsResponse struct{
    Pager pagination.Pager
}

func NewListObjectstorageV1ObjectsResponse(pager pagination.Pager,)*ListObjectstorageV1ObjectsResponse {
    return &ListObjectstorageV1ObjectsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListObjectstorageV1Objects(request *ListObjectstorageV1ObjectsRequest)(*ListObjectstorageV1ObjectsResponse){
    return NewListObjectstorageV1ObjectsResponse(objects.List(oc.client,request.ContainerName,request.Opts, ))

}
//request struct for the DownloadObjectstorageV1Objects
type DownloadObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.DownloadOptsBuilder
}

func NewDownloadObjectstorageV1ObjectsRequest()*DownloadObjectstorageV1ObjectsRequest{
    return &DownloadObjectstorageV1ObjectsRequest{}
}

//response struct for the DownloadObjectstorageV1Objects
type DownloadObjectstorageV1ObjectsResponse struct{
    DownloadResult objects.DownloadResult
}

func NewDownloadObjectstorageV1ObjectsResponse(downloadResult objects.DownloadResult,)*DownloadObjectstorageV1ObjectsResponse {
    return &DownloadObjectstorageV1ObjectsResponse{
            DownloadResult:downloadResult,
    }
}

// action function
func (oc *OpenstackClient) DownloadObjectstorageV1Objects(request *DownloadObjectstorageV1ObjectsRequest)(*DownloadObjectstorageV1ObjectsResponse){
    return NewDownloadObjectstorageV1ObjectsResponse(objects.Download(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the CreateObjectstorageV1Objects
type CreateObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.CreateOptsBuilder
}

func NewCreateObjectstorageV1ObjectsRequest()*CreateObjectstorageV1ObjectsRequest{
    return &CreateObjectstorageV1ObjectsRequest{}
}

//response struct for the CreateObjectstorageV1Objects
type CreateObjectstorageV1ObjectsResponse struct{
    CreateResult objects.CreateResult
}

func NewCreateObjectstorageV1ObjectsResponse(createResult objects.CreateResult,)*CreateObjectstorageV1ObjectsResponse {
    return &CreateObjectstorageV1ObjectsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateObjectstorageV1Objects(request *CreateObjectstorageV1ObjectsRequest)(*CreateObjectstorageV1ObjectsResponse){
    return NewCreateObjectstorageV1ObjectsResponse(objects.Create(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the CopyObjectstorageV1Objects
type CopyObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.CopyOptsBuilder
}

func NewCopyObjectstorageV1ObjectsRequest()*CopyObjectstorageV1ObjectsRequest{
    return &CopyObjectstorageV1ObjectsRequest{}
}

//response struct for the CopyObjectstorageV1Objects
type CopyObjectstorageV1ObjectsResponse struct{
    CopyResult objects.CopyResult
}

func NewCopyObjectstorageV1ObjectsResponse(copyResult objects.CopyResult,)*CopyObjectstorageV1ObjectsResponse {
    return &CopyObjectstorageV1ObjectsResponse{
            CopyResult:copyResult,
    }
}

// action function
func (oc *OpenstackClient) CopyObjectstorageV1Objects(request *CopyObjectstorageV1ObjectsRequest)(*CopyObjectstorageV1ObjectsResponse){
    return NewCopyObjectstorageV1ObjectsResponse(objects.Copy(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the DeleteObjectstorageV1Objects
type DeleteObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.DeleteOptsBuilder
}

func NewDeleteObjectstorageV1ObjectsRequest()*DeleteObjectstorageV1ObjectsRequest{
    return &DeleteObjectstorageV1ObjectsRequest{}
}

//response struct for the DeleteObjectstorageV1Objects
type DeleteObjectstorageV1ObjectsResponse struct{
    DeleteResult objects.DeleteResult
}

func NewDeleteObjectstorageV1ObjectsResponse(deleteResult objects.DeleteResult,)*DeleteObjectstorageV1ObjectsResponse {
    return &DeleteObjectstorageV1ObjectsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteObjectstorageV1Objects(request *DeleteObjectstorageV1ObjectsRequest)(*DeleteObjectstorageV1ObjectsResponse){
    return NewDeleteObjectstorageV1ObjectsResponse(objects.Delete(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the GetObjectstorageV1Objects
type GetObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.GetOptsBuilder
}

func NewGetObjectstorageV1ObjectsRequest()*GetObjectstorageV1ObjectsRequest{
    return &GetObjectstorageV1ObjectsRequest{}
}

//response struct for the GetObjectstorageV1Objects
type GetObjectstorageV1ObjectsResponse struct{
    GetResult objects.GetResult
}

func NewGetObjectstorageV1ObjectsResponse(getResult objects.GetResult,)*GetObjectstorageV1ObjectsResponse {
    return &GetObjectstorageV1ObjectsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetObjectstorageV1Objects(request *GetObjectstorageV1ObjectsRequest)(*GetObjectstorageV1ObjectsResponse){
    return NewGetObjectstorageV1ObjectsResponse(objects.Get(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the UpdateObjectstorageV1Objects
type UpdateObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.UpdateOptsBuilder
}

func NewUpdateObjectstorageV1ObjectsRequest()*UpdateObjectstorageV1ObjectsRequest{
    return &UpdateObjectstorageV1ObjectsRequest{}
}

//response struct for the UpdateObjectstorageV1Objects
type UpdateObjectstorageV1ObjectsResponse struct{
    UpdateResult objects.UpdateResult
}

func NewUpdateObjectstorageV1ObjectsResponse(updateResult objects.UpdateResult,)*UpdateObjectstorageV1ObjectsResponse {
    return &UpdateObjectstorageV1ObjectsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateObjectstorageV1Objects(request *UpdateObjectstorageV1ObjectsRequest)(*UpdateObjectstorageV1ObjectsResponse){
    return NewUpdateObjectstorageV1ObjectsResponse(objects.Update(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the CreateTempURLObjectstorageV1Objects
type CreateTempURLObjectstorageV1ObjectsRequest struct{
    ContainerName string
    ObjectName string
    Opts objects.CreateTempURLOpts
}

func NewCreateTempURLObjectstorageV1ObjectsRequest()*CreateTempURLObjectstorageV1ObjectsRequest{
    return &CreateTempURLObjectstorageV1ObjectsRequest{}
}

//response struct for the CreateTempURLObjectstorageV1Objects
type CreateTempURLObjectstorageV1ObjectsResponse struct{
    String string
    Error error
}

func NewCreateTempURLObjectstorageV1ObjectsResponse(string string,error error,)*CreateTempURLObjectstorageV1ObjectsResponse {
    return &CreateTempURLObjectstorageV1ObjectsResponse{
            String:string,
            Error:error,
    }
}

// action function
func (oc *OpenstackClient) CreateTempURLObjectstorageV1Objects(request *CreateTempURLObjectstorageV1ObjectsRequest)(*CreateTempURLObjectstorageV1ObjectsResponse){
    return NewCreateTempURLObjectstorageV1ObjectsResponse(objects.CreateTempURL(oc.client,request.ContainerName,request.ObjectName,request.Opts, ))

}
//request struct for the BulkDeleteObjectstorageV1Objects
type BulkDeleteObjectstorageV1ObjectsRequest struct{
    Container string
    Objects 
}

func NewBulkDeleteObjectstorageV1ObjectsRequest()*BulkDeleteObjectstorageV1ObjectsRequest{
    return &BulkDeleteObjectstorageV1ObjectsRequest{}
}

//response struct for the BulkDeleteObjectstorageV1Objects
type BulkDeleteObjectstorageV1ObjectsResponse struct{
    BulkDeleteResult objects.BulkDeleteResult
}

func NewBulkDeleteObjectstorageV1ObjectsResponse(bulkDeleteResult objects.BulkDeleteResult,)*BulkDeleteObjectstorageV1ObjectsResponse {
    return &BulkDeleteObjectstorageV1ObjectsResponse{
            BulkDeleteResult:bulkDeleteResult,
    }
}

// action function
func (oc *OpenstackClient) BulkDeleteObjectstorageV1Objects(request *BulkDeleteObjectstorageV1ObjectsRequest)(*BulkDeleteObjectstorageV1ObjectsResponse){
    return NewBulkDeleteObjectstorageV1ObjectsResponse(objects.BulkDelete(oc.client,request.Container,request.Objects, ))

}