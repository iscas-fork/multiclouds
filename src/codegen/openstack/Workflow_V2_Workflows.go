package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/workflow/v2/workflows"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateWorkflowV2Workflows
type CreateWorkflowV2WorkflowsRequest struct{
    Opts workflows.CreateOptsBuilder
}

func NewCreateWorkflowV2WorkflowsRequest()*CreateWorkflowV2WorkflowsRequest{
    return &CreateWorkflowV2WorkflowsRequest{}
}

//response struct for the CreateWorkflowV2Workflows
type CreateWorkflowV2WorkflowsResponse struct{
    CreateResult workflows.CreateResult
}

func NewCreateWorkflowV2WorkflowsResponse(createResult workflows.CreateResult,)*CreateWorkflowV2WorkflowsResponse {
    return &CreateWorkflowV2WorkflowsResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateWorkflowV2Workflows(request *CreateWorkflowV2WorkflowsRequest)(*CreateWorkflowV2WorkflowsResponse){
    return NewCreateWorkflowV2WorkflowsResponse(workflows.Create(oc.client,request.Opts, ))

}
//request struct for the DeleteWorkflowV2Workflows
type DeleteWorkflowV2WorkflowsRequest struct{
    Id string
}

func NewDeleteWorkflowV2WorkflowsRequest()*DeleteWorkflowV2WorkflowsRequest{
    return &DeleteWorkflowV2WorkflowsRequest{}
}

//response struct for the DeleteWorkflowV2Workflows
type DeleteWorkflowV2WorkflowsResponse struct{
    DeleteResult workflows.DeleteResult
}

func NewDeleteWorkflowV2WorkflowsResponse(deleteResult workflows.DeleteResult,)*DeleteWorkflowV2WorkflowsResponse {
    return &DeleteWorkflowV2WorkflowsResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteWorkflowV2Workflows(request *DeleteWorkflowV2WorkflowsRequest)(*DeleteWorkflowV2WorkflowsResponse){
    return NewDeleteWorkflowV2WorkflowsResponse(workflows.Delete(oc.client,request.Id, ))

}
//request struct for the GetWorkflowV2Workflows
type GetWorkflowV2WorkflowsRequest struct{
    Id string
}

func NewGetWorkflowV2WorkflowsRequest()*GetWorkflowV2WorkflowsRequest{
    return &GetWorkflowV2WorkflowsRequest{}
}

//response struct for the GetWorkflowV2Workflows
type GetWorkflowV2WorkflowsResponse struct{
    GetResult workflows.GetResult
}

func NewGetWorkflowV2WorkflowsResponse(getResult workflows.GetResult,)*GetWorkflowV2WorkflowsResponse {
    return &GetWorkflowV2WorkflowsResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetWorkflowV2Workflows(request *GetWorkflowV2WorkflowsRequest)(*GetWorkflowV2WorkflowsResponse){
    return NewGetWorkflowV2WorkflowsResponse(workflows.Get(oc.client,request.Id, ))

}
//request struct for the ListWorkflowV2Workflows
type ListWorkflowV2WorkflowsRequest struct{
    Opts workflows.ListOptsBuilder
}

func NewListWorkflowV2WorkflowsRequest()*ListWorkflowV2WorkflowsRequest{
    return &ListWorkflowV2WorkflowsRequest{}
}

//response struct for the ListWorkflowV2Workflows
type ListWorkflowV2WorkflowsResponse struct{
    Pager pagination.Pager
}

func NewListWorkflowV2WorkflowsResponse(pager pagination.Pager,)*ListWorkflowV2WorkflowsResponse {
    return &ListWorkflowV2WorkflowsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListWorkflowV2Workflows(request *ListWorkflowV2WorkflowsRequest)(*ListWorkflowV2WorkflowsResponse){
    return NewListWorkflowV2WorkflowsResponse(workflows.List(oc.client,request.Opts, ))

}