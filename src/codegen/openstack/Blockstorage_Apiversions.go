package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/blockstorage/apiversions"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListBlockstorageApiversions
type ListBlockstorageApiversionsRequest struct{
}

func NewListBlockstorageApiversionsRequest()*ListBlockstorageApiversionsRequest{
    return &ListBlockstorageApiversionsRequest{}
}

//response struct for the ListBlockstorageApiversions
type ListBlockstorageApiversionsResponse struct{
    Pager pagination.Pager
}

func NewListBlockstorageApiversionsResponse(pager pagination.Pager,)*ListBlockstorageApiversionsResponse {
    return &ListBlockstorageApiversionsResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListBlockstorageApiversions(request *ListBlockstorageApiversionsRequest)(*ListBlockstorageApiversionsResponse){
    return NewListBlockstorageApiversionsResponse(apiversions.List(oc.client, ))

}