package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/services"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListComputeV2ExtensionsServices
type ListComputeV2ExtensionsServicesRequest struct{
    Opts services.ListOptsBuilder
}

func NewListComputeV2ExtensionsServicesRequest()*ListComputeV2ExtensionsServicesRequest{
    return &ListComputeV2ExtensionsServicesRequest{}
}

//response struct for the ListComputeV2ExtensionsServices
type ListComputeV2ExtensionsServicesResponse struct{
    Pager pagination.Pager
}

func NewListComputeV2ExtensionsServicesResponse(pager pagination.Pager,)*ListComputeV2ExtensionsServicesResponse {
    return &ListComputeV2ExtensionsServicesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListComputeV2ExtensionsServices(request *ListComputeV2ExtensionsServicesRequest)(*ListComputeV2ExtensionsServicesResponse){
    return NewListComputeV2ExtensionsServicesResponse(services.List(oc.client,request.Opts, ))

}
//request struct for the UpdateComputeV2ExtensionsServices
type UpdateComputeV2ExtensionsServicesRequest struct{
    Id string
    Opts services.UpdateOpts
}

func NewUpdateComputeV2ExtensionsServicesRequest()*UpdateComputeV2ExtensionsServicesRequest{
    return &UpdateComputeV2ExtensionsServicesRequest{}
}

//response struct for the UpdateComputeV2ExtensionsServices
type UpdateComputeV2ExtensionsServicesResponse struct{
    UpdateResult services.UpdateResult
}

func NewUpdateComputeV2ExtensionsServicesResponse(updateResult services.UpdateResult,)*UpdateComputeV2ExtensionsServicesResponse {
    return &UpdateComputeV2ExtensionsServicesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateComputeV2ExtensionsServices(request *UpdateComputeV2ExtensionsServicesRequest)(*UpdateComputeV2ExtensionsServicesResponse){
    return NewUpdateComputeV2ExtensionsServicesResponse(services.Update(oc.client,request.Id,request.Opts, ))

}
//request struct for the DeleteComputeV2ExtensionsServices
type DeleteComputeV2ExtensionsServicesRequest struct{
    Id string
}

func NewDeleteComputeV2ExtensionsServicesRequest()*DeleteComputeV2ExtensionsServicesRequest{
    return &DeleteComputeV2ExtensionsServicesRequest{}
}

//response struct for the DeleteComputeV2ExtensionsServices
type DeleteComputeV2ExtensionsServicesResponse struct{
    DeleteResult services.DeleteResult
}

func NewDeleteComputeV2ExtensionsServicesResponse(deleteResult services.DeleteResult,)*DeleteComputeV2ExtensionsServicesResponse {
    return &DeleteComputeV2ExtensionsServicesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteComputeV2ExtensionsServices(request *DeleteComputeV2ExtensionsServicesRequest)(*DeleteComputeV2ExtensionsServicesResponse){
    return NewDeleteComputeV2ExtensionsServicesResponse(services.Delete(oc.client,request.Id, ))

}