package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/containerinfra/v1/quotas"
)
//request struct for the CreateContainerinfraV1Quotas
type CreateContainerinfraV1QuotasRequest struct{
    Opts quotas.CreateOptsBuilder
}

func NewCreateContainerinfraV1QuotasRequest()*CreateContainerinfraV1QuotasRequest{
    return &CreateContainerinfraV1QuotasRequest{}
}

//response struct for the CreateContainerinfraV1Quotas
type CreateContainerinfraV1QuotasResponse struct{
    CreateResult quotas.CreateResult
}

func NewCreateContainerinfraV1QuotasResponse(createResult quotas.CreateResult,)*CreateContainerinfraV1QuotasResponse {
    return &CreateContainerinfraV1QuotasResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateContainerinfraV1Quotas(request *CreateContainerinfraV1QuotasRequest)(*CreateContainerinfraV1QuotasResponse){
    return NewCreateContainerinfraV1QuotasResponse(quotas.Create(oc.client,request.Opts, ))

}