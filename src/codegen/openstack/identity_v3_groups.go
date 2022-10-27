package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v3/groups"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV3Groups
type ListIdentityV3GroupsRequest struct{
    Opts groups.ListOptsBuilder
}

func NewListIdentityV3GroupsRequest()*ListIdentityV3GroupsRequest{
    return &ListIdentityV3GroupsRequest{}
}

//response struct for the ListIdentityV3Groups
type ListIdentityV3GroupsResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV3GroupsResponse(pager pagination.Pager,)*ListIdentityV3GroupsResponse {
    return &ListIdentityV3GroupsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV3Groups(request *ListIdentityV3GroupsRequest)(*ListIdentityV3GroupsResponse){
    return NewListIdentityV3GroupsResponse(groups.List(oc.client,request.Opts, ))

}
//request struct for the GetIdentityV3Groups
type GetIdentityV3GroupsRequest struct{
    Id string
}

func NewGetIdentityV3GroupsRequest()*GetIdentityV3GroupsRequest{
    return &GetIdentityV3GroupsRequest{}
}

//response struct for the GetIdentityV3Groups
type GetIdentityV3GroupsResponse struct{
    GetResult groups.GetResult
}

func NewGetIdentityV3GroupsResponse(getResult groups.GetResult,)*GetIdentityV3GroupsResponse {
    return &GetIdentityV3GroupsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetIdentityV3Groups(request *GetIdentityV3GroupsRequest)(*GetIdentityV3GroupsResponse){
    return NewGetIdentityV3GroupsResponse(groups.Get(oc.client,request.Id, ))

}
//request struct for the CreateIdentityV3Groups
type CreateIdentityV3GroupsRequest struct{
    Opts groups.CreateOptsBuilder
}

func NewCreateIdentityV3GroupsRequest()*CreateIdentityV3GroupsRequest{
    return &CreateIdentityV3GroupsRequest{}
}

//response struct for the CreateIdentityV3Groups
type CreateIdentityV3GroupsResponse struct{
    CreateResult groups.CreateResult
}

func NewCreateIdentityV3GroupsResponse(createResult groups.CreateResult,)*CreateIdentityV3GroupsResponse {
    return &CreateIdentityV3GroupsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateIdentityV3Groups(request *CreateIdentityV3GroupsRequest)(*CreateIdentityV3GroupsResponse){
    return NewCreateIdentityV3GroupsResponse(groups.Create(oc.client,request.Opts, ))

}
//request struct for the UpdateIdentityV3Groups
type UpdateIdentityV3GroupsRequest struct{
    GroupID string
    Opts groups.UpdateOptsBuilder
}

func NewUpdateIdentityV3GroupsRequest()*UpdateIdentityV3GroupsRequest{
    return &UpdateIdentityV3GroupsRequest{}
}

//response struct for the UpdateIdentityV3Groups
type UpdateIdentityV3GroupsResponse struct{
    UpdateResult groups.UpdateResult
}

func NewUpdateIdentityV3GroupsResponse(updateResult groups.UpdateResult,)*UpdateIdentityV3GroupsResponse {
    return &UpdateIdentityV3GroupsResponse{
            UpdateResult:updateResult,
    }
}

// action function
func (oc *OpenstackClient) UpdateIdentityV3Groups(request *UpdateIdentityV3GroupsRequest)(*UpdateIdentityV3GroupsResponse){
    return NewUpdateIdentityV3GroupsResponse(groups.Update(oc.client,request.GroupID,request.Opts, ))

}
//request struct for the DeleteIdentityV3Groups
type DeleteIdentityV3GroupsRequest struct{
    GroupID string
}

func NewDeleteIdentityV3GroupsRequest()*DeleteIdentityV3GroupsRequest{
    return &DeleteIdentityV3GroupsRequest{}
}

//response struct for the DeleteIdentityV3Groups
type DeleteIdentityV3GroupsResponse struct{
    DeleteResult groups.DeleteResult
}

func NewDeleteIdentityV3GroupsResponse(deleteResult groups.DeleteResult,)*DeleteIdentityV3GroupsResponse {
    return &DeleteIdentityV3GroupsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteIdentityV3Groups(request *DeleteIdentityV3GroupsRequest)(*DeleteIdentityV3GroupsResponse){
    return NewDeleteIdentityV3GroupsResponse(groups.Delete(oc.client,request.GroupID, ))

}