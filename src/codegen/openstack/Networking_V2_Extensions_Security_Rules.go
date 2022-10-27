package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/security/rules"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListNetworkingV2ExtensionsSecurityRules
type ListNetworkingV2ExtensionsSecurityRulesRequest struct{
    Opts rules.ListOpts
}

func NewListNetworkingV2ExtensionsSecurityRulesRequest()*ListNetworkingV2ExtensionsSecurityRulesRequest{
    return &ListNetworkingV2ExtensionsSecurityRulesRequest{}
}

//response struct for the ListNetworkingV2ExtensionsSecurityRules
type ListNetworkingV2ExtensionsSecurityRulesResponse struct{
    Pager pagination.Pager
}

func NewListNetworkingV2ExtensionsSecurityRulesResponse(pager pagination.Pager,)*ListNetworkingV2ExtensionsSecurityRulesResponse {
    return &ListNetworkingV2ExtensionsSecurityRulesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListNetworkingV2ExtensionsSecurityRules(request *ListNetworkingV2ExtensionsSecurityRulesRequest)(*ListNetworkingV2ExtensionsSecurityRulesResponse){
    return NewListNetworkingV2ExtensionsSecurityRulesResponse(rules.List(oc.client,request.Opts, ))

}
//request struct for the CreateNetworkingV2ExtensionsSecurityRules
type CreateNetworkingV2ExtensionsSecurityRulesRequest struct{
    Opts rules.CreateOptsBuilder
}

func NewCreateNetworkingV2ExtensionsSecurityRulesRequest()*CreateNetworkingV2ExtensionsSecurityRulesRequest{
    return &CreateNetworkingV2ExtensionsSecurityRulesRequest{}
}

//response struct for the CreateNetworkingV2ExtensionsSecurityRules
type CreateNetworkingV2ExtensionsSecurityRulesResponse struct{
    CreateResult rules.CreateResult
}

func NewCreateNetworkingV2ExtensionsSecurityRulesResponse(createResult rules.CreateResult,)*CreateNetworkingV2ExtensionsSecurityRulesResponse {
    return &CreateNetworkingV2ExtensionsSecurityRulesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateNetworkingV2ExtensionsSecurityRules(request *CreateNetworkingV2ExtensionsSecurityRulesRequest)(*CreateNetworkingV2ExtensionsSecurityRulesResponse){
    return NewCreateNetworkingV2ExtensionsSecurityRulesResponse(rules.Create(oc.client,request.Opts, ))

}
//request struct for the GetNetworkingV2ExtensionsSecurityRules
type GetNetworkingV2ExtensionsSecurityRulesRequest struct{
    Id string
}

func NewGetNetworkingV2ExtensionsSecurityRulesRequest()*GetNetworkingV2ExtensionsSecurityRulesRequest{
    return &GetNetworkingV2ExtensionsSecurityRulesRequest{}
}

//response struct for the GetNetworkingV2ExtensionsSecurityRules
type GetNetworkingV2ExtensionsSecurityRulesResponse struct{
    GetResult rules.GetResult
}

func NewGetNetworkingV2ExtensionsSecurityRulesResponse(getResult rules.GetResult,)*GetNetworkingV2ExtensionsSecurityRulesResponse {
    return &GetNetworkingV2ExtensionsSecurityRulesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetNetworkingV2ExtensionsSecurityRules(request *GetNetworkingV2ExtensionsSecurityRulesRequest)(*GetNetworkingV2ExtensionsSecurityRulesResponse){
    return NewGetNetworkingV2ExtensionsSecurityRulesResponse(rules.Get(oc.client,request.Id, ))

}
//request struct for the DeleteNetworkingV2ExtensionsSecurityRules
type DeleteNetworkingV2ExtensionsSecurityRulesRequest struct{
    Id string
}

func NewDeleteNetworkingV2ExtensionsSecurityRulesRequest()*DeleteNetworkingV2ExtensionsSecurityRulesRequest{
    return &DeleteNetworkingV2ExtensionsSecurityRulesRequest{}
}

//response struct for the DeleteNetworkingV2ExtensionsSecurityRules
type DeleteNetworkingV2ExtensionsSecurityRulesResponse struct{
    DeleteResult rules.DeleteResult
}

func NewDeleteNetworkingV2ExtensionsSecurityRulesResponse(deleteResult rules.DeleteResult,)*DeleteNetworkingV2ExtensionsSecurityRulesResponse {
    return &DeleteNetworkingV2ExtensionsSecurityRulesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteNetworkingV2ExtensionsSecurityRules(request *DeleteNetworkingV2ExtensionsSecurityRulesRequest)(*DeleteNetworkingV2ExtensionsSecurityRulesResponse){
    return NewDeleteNetworkingV2ExtensionsSecurityRulesResponse(rules.Delete(oc.client,request.Id, ))

}