package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/layer3/portforwarding"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsLayer3Portforwarding
type ListNetworkingV2ExtensionsLayer3PortforwardingRequest struct{
    Opts portforwarding.ListOptsBuilder
    Id string
}

func NewListNetworkingV2ExtensionsLayer3PortforwardingRequest()*ListNetworkingV2ExtensionsLayer3PortforwardingRequest{
    return &ListNetworkingV2ExtensionsLayer3PortforwardingRequest{}
}

//response struct for the ListNetworkingV2ExtensionsLayer3Portforwarding
type ListNetworkingV2ExtensionsLayer3PortforwardingResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsLayer3PortforwardingResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsLayer3PortforwardingResponse {
    return &ListNetworkingV2ExtensionsLayer3PortforwardingResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsLayer3Portforwarding(request *ListNetworkingV2ExtensionsLayer3PortforwardingRequest)(*ListNetworkingV2ExtensionsLayer3PortforwardingResponse){
    return NewListNetworkingV2ExtensionsLayer3PortforwardingResponse(portforwarding.List(oc.client,request.Opts,request.Id, ))

}
//request struct for the GetNetworkingV2ExtensionsLayer3Portforwarding
type GetNetworkingV2ExtensionsLayer3PortforwardingRequest struct{
    FloatingIpId string
    PfId string
}

func NewGetNetworkingV2ExtensionsLayer3PortforwardingRequest()*GetNetworkingV2ExtensionsLayer3PortforwardingRequest{
    return &GetNetworkingV2ExtensionsLayer3PortforwardingRequest{}
}

//response struct for the GetNetworkingV2ExtensionsLayer3Portforwarding
type GetNetworkingV2ExtensionsLayer3PortforwardingResponse struct{
    GetResult portforwarding.GetResult
}

func NewGetNetworkingV2ExtensionsLayer3PortforwardingResponse(getResult portforwarding.GetResult,)*GetNetworkingV2ExtensionsLayer3PortforwardingResponse {
    return &GetNetworkingV2ExtensionsLayer3PortforwardingResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsLayer3Portforwarding(request *GetNetworkingV2ExtensionsLayer3PortforwardingRequest)(*GetNetworkingV2ExtensionsLayer3PortforwardingResponse){
    return NewGetNetworkingV2ExtensionsLayer3PortforwardingResponse(portforwarding.Get(oc.client,request.FloatingIpId,request.PfId, ))

}
//request struct for the CreateNetworkingV2ExtensionsLayer3Portforwarding
type CreateNetworkingV2ExtensionsLayer3PortforwardingRequest struct{
    FloatingIpId string
    Opts portforwarding.CreateOptsBuilder
}

func NewCreateNetworkingV2ExtensionsLayer3PortforwardingRequest()*CreateNetworkingV2ExtensionsLayer3PortforwardingRequest{
    return &CreateNetworkingV2ExtensionsLayer3PortforwardingRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsLayer3Portforwarding
type CreateNetworkingV2ExtensionsLayer3PortforwardingResponse struct{
    CreateResult portforwarding.CreateResult
}

func NewCreateNetworkingV2ExtensionsLayer3PortforwardingResponse(createResult portforwarding.CreateResult,)*CreateNetworkingV2ExtensionsLayer3PortforwardingResponse {
    return &CreateNetworkingV2ExtensionsLayer3PortforwardingResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsLayer3Portforwarding(request *CreateNetworkingV2ExtensionsLayer3PortforwardingRequest)(*CreateNetworkingV2ExtensionsLayer3PortforwardingResponse){
    return NewCreateNetworkingV2ExtensionsLayer3PortforwardingResponse(portforwarding.Create(oc.client,request.FloatingIpId,request.Opts, ))

}
//request struct for the UpdateNetworkingV2ExtensionsLayer3Portforwarding
type UpdateNetworkingV2ExtensionsLayer3PortforwardingRequest struct{
    FipID string
    PfID string
    Opts portforwarding.UpdateOptsBuilder
}

func NewUpdateNetworkingV2ExtensionsLayer3PortforwardingRequest()*UpdateNetworkingV2ExtensionsLayer3PortforwardingRequest{
    return &UpdateNetworkingV2ExtensionsLayer3PortforwardingRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsLayer3Portforwarding
type UpdateNetworkingV2ExtensionsLayer3PortforwardingResponse struct{
    UpdateResult portforwarding.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsLayer3PortforwardingResponse(updateResult portforwarding.UpdateResult,)*UpdateNetworkingV2ExtensionsLayer3PortforwardingResponse {
    return &UpdateNetworkingV2ExtensionsLayer3PortforwardingResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsLayer3Portforwarding(request *UpdateNetworkingV2ExtensionsLayer3PortforwardingRequest)(*UpdateNetworkingV2ExtensionsLayer3PortforwardingResponse){
    return NewUpdateNetworkingV2ExtensionsLayer3PortforwardingResponse(portforwarding.Update(oc.client,request.FipID,request.PfID,request.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsLayer3Portforwarding
type DeleteNetworkingV2ExtensionsLayer3PortforwardingRequest struct{
    FloatingIpId string
    PfId string
}

func NewDeleteNetworkingV2ExtensionsLayer3PortforwardingRequest()*DeleteNetworkingV2ExtensionsLayer3PortforwardingRequest{
    return &DeleteNetworkingV2ExtensionsLayer3PortforwardingRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsLayer3Portforwarding
type DeleteNetworkingV2ExtensionsLayer3PortforwardingResponse struct{
    DeleteResult portforwarding.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsLayer3PortforwardingResponse(deleteResult portforwarding.DeleteResult,)*DeleteNetworkingV2ExtensionsLayer3PortforwardingResponse {
    return &DeleteNetworkingV2ExtensionsLayer3PortforwardingResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsLayer3Portforwarding(request *DeleteNetworkingV2ExtensionsLayer3PortforwardingRequest)(*DeleteNetworkingV2ExtensionsLayer3PortforwardingResponse){
    return NewDeleteNetworkingV2ExtensionsLayer3PortforwardingResponse(portforwarding.Delete(oc.client,request.FloatingIpId,request.PfId, ))

}