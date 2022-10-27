package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/tags"
)
//request struct for the ListTags
type ListTagsRequest struct{
    serverID string
}

func NewListTagsRequest()*ListTagsRequest{
    return &ListTagsRequest{}
}

//response struct for the ListTags
type ListTagsResponse struct{
    ListResult tags.ListResult
}

func NewListTagsResponse(listResult tags.ListResult,)*ListTagsResponse {
    return &ListTagsResponse{
            ListResult:listResult,
    }
}

// action function
func (oc *OpenstackClient) ListTags(request *ListTagsRequest)(*ListTagsResponse){
    return NewListTagsResponse(tags.List(oc.client,request.serverID, ))

}
//request struct for the CheckTags
type CheckTagsRequest struct{
    serverID string
    tag string
}

func NewCheckTagsRequest()*CheckTagsRequest{
    return &CheckTagsRequest{}
}

//response struct for the CheckTags
type CheckTagsResponse struct{
    CheckResult tags.CheckResult
}

func NewCheckTagsResponse(checkResult tags.CheckResult,)*CheckTagsResponse {
    return &CheckTagsResponse{
            CheckResult:checkResult,
    }
}

// action function
func (oc *OpenstackClient) CheckTags(request *CheckTagsRequest)(*CheckTagsResponse){
    return NewCheckTagsResponse(tags.Check(oc.client,request.serverID,request.tag, ))

}
//request struct for the ReplaceAllTags
type ReplaceAllTagsRequest struct{
    serverID string
    opts tags.ReplaceAllOptsBuilder
}

func NewReplaceAllTagsRequest()*ReplaceAllTagsRequest{
    return &ReplaceAllTagsRequest{}
}

//response struct for the ReplaceAllTags
type ReplaceAllTagsResponse struct{
    ReplaceAllResult tags.ReplaceAllResult
}

func NewReplaceAllTagsResponse(replaceAllResult tags.ReplaceAllResult,)*ReplaceAllTagsResponse {
    return &ReplaceAllTagsResponse{
            ReplaceAllResult:replaceAllResult,
    }
}

// action function
func (oc *OpenstackClient) ReplaceAllTags(request *ReplaceAllTagsRequest)(*ReplaceAllTagsResponse){
    return NewReplaceAllTagsResponse(tags.ReplaceAll(oc.client,request.serverID,request.opts, ))

}
//request struct for the AddTags
type AddTagsRequest struct{
    serverID string
    tag string
}

func NewAddTagsRequest()*AddTagsRequest{
    return &AddTagsRequest{}
}

//response struct for the AddTags
type AddTagsResponse struct{
    AddResult tags.AddResult
}

func NewAddTagsResponse(addResult tags.AddResult,)*AddTagsResponse {
    return &AddTagsResponse{
            AddResult:addResult,
    }
}

// action function
func (oc *OpenstackClient) AddTags(request *AddTagsRequest)(*AddTagsResponse){
    return NewAddTagsResponse(tags.Add(oc.client,request.serverID,request.tag, ))

}
//request struct for the DeleteTags
type DeleteTagsRequest struct{
    serverID string
    tag string
}

func NewDeleteTagsRequest()*DeleteTagsRequest{
    return &DeleteTagsRequest{}
}

//response struct for the DeleteTags
type DeleteTagsResponse struct{
    DeleteResult tags.DeleteResult
}

func NewDeleteTagsResponse(deleteResult tags.DeleteResult,)*DeleteTagsResponse {
    return &DeleteTagsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteTags(request *DeleteTagsRequest)(*DeleteTagsResponse){
    return NewDeleteTagsResponse(tags.Delete(oc.client,request.serverID,request.tag, ))

}
//request struct for the DeleteAllTags
type DeleteAllTagsRequest struct{
    serverID string
}

func NewDeleteAllTagsRequest()*DeleteAllTagsRequest{
    return &DeleteAllTagsRequest{}
}

//response struct for the DeleteAllTags
type DeleteAllTagsResponse struct{
    DeleteResult tags.DeleteResult
}

func NewDeleteAllTagsResponse(deleteResult tags.DeleteResult,)*DeleteAllTagsResponse {
    return &DeleteAllTagsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteAllTags(request *DeleteAllTagsRequest)(*DeleteAllTagsResponse){
    return NewDeleteAllTagsResponse(tags.DeleteAll(oc.client,request.serverID, ))

}