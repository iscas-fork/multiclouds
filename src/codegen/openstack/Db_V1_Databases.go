package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/databases"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateDbV1Databases
type CreateDbV1DatabasesRequest struct{
    InstanceID string
    Opts databases.CreateOptsBuilder
}

func NewCreateDbV1DatabasesRequest()*CreateDbV1DatabasesRequest{
    return &CreateDbV1DatabasesRequest{}
}

//response struct for the CreateDbV1Databases
type CreateDbV1DatabasesResponse struct{
    CreateResult databases.CreateResult
}

func NewCreateDbV1DatabasesResponse(createResult databases.CreateResult,)*CreateDbV1DatabasesResponse {
    return &CreateDbV1DatabasesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateDbV1Databases(request *CreateDbV1DatabasesRequest)(*CreateDbV1DatabasesResponse){
    return NewCreateDbV1DatabasesResponse(databases.Create(oc.client,request.InstanceID,request.Opts, ))

}
//request struct for the ListDbV1Databases
type ListDbV1DatabasesRequest struct{
    InstanceID string
}

func NewListDbV1DatabasesRequest()*ListDbV1DatabasesRequest{
    return &ListDbV1DatabasesRequest{}
}

//response struct for the ListDbV1Databases
type ListDbV1DatabasesResponse struct{
    Pager pagination.Pager
}

func NewListDbV1DatabasesResponse(pager pagination.Pager,)*ListDbV1DatabasesResponse {
    return &ListDbV1DatabasesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDbV1Databases(request *ListDbV1DatabasesRequest)(*ListDbV1DatabasesResponse){
    return NewListDbV1DatabasesResponse(databases.List(oc.client,request.InstanceID, ))

}
//request struct for the DeleteDbV1Databases
type DeleteDbV1DatabasesRequest struct{
    InstanceID string
    DbName string
}

func NewDeleteDbV1DatabasesRequest()*DeleteDbV1DatabasesRequest{
    return &DeleteDbV1DatabasesRequest{}
}

//response struct for the DeleteDbV1Databases
type DeleteDbV1DatabasesResponse struct{
    DeleteResult databases.DeleteResult
}

func NewDeleteDbV1DatabasesResponse(deleteResult databases.DeleteResult,)*DeleteDbV1DatabasesResponse {
    return &DeleteDbV1DatabasesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteDbV1Databases(request *DeleteDbV1DatabasesRequest)(*DeleteDbV1DatabasesResponse){
    return NewDeleteDbV1DatabasesResponse(databases.Delete(oc.client,request.InstanceID,request.DbName, ))

}