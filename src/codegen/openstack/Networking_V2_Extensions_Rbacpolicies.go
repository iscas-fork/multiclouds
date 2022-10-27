package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/rbacpolicies"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsRbacpolicies
type ListNetworkingV2ExtensionsRbacpoliciesRequest struct{
    Opts rbacpolicies.ListOptsBuilder
}

func NewListNetworkingV2ExtensionsRbacpoliciesRequest()*ListNetworkingV2ExtensionsRbacpoliciesRequest{
    return &ListNetworkingV2ExtensionsRbacpoliciesRequest{}
}

//response struct for the ListNetworkingV2ExtensionsRbacpolicies
type ListNetworkingV2ExtensionsRbacpoliciesResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsRbacpoliciesResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsRbacpoliciesResponse {
    return &ListNetworkingV2ExtensionsRbacpoliciesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsRbacpolicies(request *ListNetworkingV2ExtensionsRbacpoliciesRequest)(*ListNetworkingV2ExtensionsRbacpoliciesResponse){
    return NewListNetworkingV2ExtensionsRbacpoliciesResponse(rbacpolicies.List(oc.client,request.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsRbacpolicies
type GetNetworkingV2ExtensionsRbacpoliciesRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsRbacpoliciesRequest()*GetNetworkingV2ExtensionsRbacpoliciesRequest{
    return &GetNetworkingV2ExtensionsRbacpoliciesRequest{}
}

//response struct for the GetNetworkingV2ExtensionsRbacpolicies
type GetNetworkingV2ExtensionsRbacpoliciesResponse struct{
    GetResult rbacpolicies.GetResult
}

func NewGetNetworkingV2ExtensionsRbacpoliciesResponse(getResult rbacpolicies.GetResult,)*GetNetworkingV2ExtensionsRbacpoliciesResponse {
    return &GetNetworkingV2ExtensionsRbacpoliciesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsRbacpolicies(request *GetNetworkingV2ExtensionsRbacpoliciesRequest)(*GetNetworkingV2ExtensionsRbacpoliciesResponse){
    return NewGetNetworkingV2ExtensionsRbacpoliciesResponse(rbacpolicies.Get(oc.client,request.Id, ))

}
//request struct for the CreateNetworkingV2ExtensionsRbacpolicies
type CreateNetworkingV2ExtensionsRbacpoliciesRequest struct{
    Opts rbacpolicies.CreateOptsBuilder
}

func NewCreateNetworkingV2ExtensionsRbacpoliciesRequest()*CreateNetworkingV2ExtensionsRbacpoliciesRequest{
    return &CreateNetworkingV2ExtensionsRbacpoliciesRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsRbacpolicies
type CreateNetworkingV2ExtensionsRbacpoliciesResponse struct{
    CreateResult rbacpolicies.CreateResult
}

func NewCreateNetworkingV2ExtensionsRbacpoliciesResponse(createResult rbacpolicies.CreateResult,)*CreateNetworkingV2ExtensionsRbacpoliciesResponse {
    return &CreateNetworkingV2ExtensionsRbacpoliciesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsRbacpolicies(request *CreateNetworkingV2ExtensionsRbacpoliciesRequest)(*CreateNetworkingV2ExtensionsRbacpoliciesResponse){
    return NewCreateNetworkingV2ExtensionsRbacpoliciesResponse(rbacpolicies.Create(oc.client,request.Opts, ))

}
//request struct for the DeleteNetworkingV2ExtensionsRbacpolicies
type DeleteNetworkingV2ExtensionsRbacpoliciesRequest struct{
    RbacPolicyID string
}

func NewDeleteNetworkingV2ExtensionsRbacpoliciesRequest()*DeleteNetworkingV2ExtensionsRbacpoliciesRequest{
    return &DeleteNetworkingV2ExtensionsRbacpoliciesRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsRbacpolicies
type DeleteNetworkingV2ExtensionsRbacpoliciesResponse struct{
    DeleteResult rbacpolicies.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsRbacpoliciesResponse(deleteResult rbacpolicies.DeleteResult,)*DeleteNetworkingV2ExtensionsRbacpoliciesResponse {
    return &DeleteNetworkingV2ExtensionsRbacpoliciesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsRbacpolicies(request *DeleteNetworkingV2ExtensionsRbacpoliciesRequest)(*DeleteNetworkingV2ExtensionsRbacpoliciesResponse){
    return NewDeleteNetworkingV2ExtensionsRbacpoliciesResponse(rbacpolicies.Delete(oc.client,request.RbacPolicyID, ))

}
//request struct for the UpdateNetworkingV2ExtensionsRbacpolicies
type UpdateNetworkingV2ExtensionsRbacpoliciesRequest struct{
    RbacPolicyID string
    Opts rbacpolicies.UpdateOptsBuilder
}

func NewUpdateNetworkingV2ExtensionsRbacpoliciesRequest()*UpdateNetworkingV2ExtensionsRbacpoliciesRequest{
    return &UpdateNetworkingV2ExtensionsRbacpoliciesRequest{}
}

//response struct for the UpdateNetworkingV2ExtensionsRbacpolicies
type UpdateNetworkingV2ExtensionsRbacpoliciesResponse struct{
    UpdateResult rbacpolicies.UpdateResult
}

func NewUpdateNetworkingV2ExtensionsRbacpoliciesResponse(updateResult rbacpolicies.UpdateResult,)*UpdateNetworkingV2ExtensionsRbacpoliciesResponse {
    return &UpdateNetworkingV2ExtensionsRbacpoliciesResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateNetworkingV2ExtensionsRbacpolicies(request *UpdateNetworkingV2ExtensionsRbacpoliciesRequest)(*UpdateNetworkingV2ExtensionsRbacpoliciesResponse){
    return NewUpdateNetworkingV2ExtensionsRbacpoliciesResponse(rbacpolicies.Update(oc.client,request.RbacPolicyID,request.Opts, ))

}