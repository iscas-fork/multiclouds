package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/remoteconsoles"
)
//request struct for the CreateComputeV2ExtensionsRemoteconsoles
type CreateComputeV2ExtensionsRemoteconsolesRequest struct{
    ServerID string
    Opts remoteconsoles.CreateOptsBuilder
}

func NewCreateComputeV2ExtensionsRemoteconsolesRequest()*CreateComputeV2ExtensionsRemoteconsolesRequest{
    return &CreateComputeV2ExtensionsRemoteconsolesRequest{}
}

//response struct for the CreateComputeV2ExtensionsRemoteconsoles
type CreateComputeV2ExtensionsRemoteconsolesResponse struct{
    CreateResult remoteconsoles.CreateResult
}

func NewCreateComputeV2ExtensionsRemoteconsolesResponse(createResult remoteconsoles.CreateResult,)*CreateComputeV2ExtensionsRemoteconsolesResponse {
    return &CreateComputeV2ExtensionsRemoteconsolesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateComputeV2ExtensionsRemoteconsoles(request *CreateComputeV2ExtensionsRemoteconsolesRequest)(*CreateComputeV2ExtensionsRemoteconsolesResponse){
    return NewCreateComputeV2ExtensionsRemoteconsolesResponse(remoteconsoles.Create(oc.client,request.ServerID,request.Opts, ))

}