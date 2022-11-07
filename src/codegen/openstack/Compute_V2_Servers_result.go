package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
)


//extract response info from pager for ListComputeV2Servers
func ExtractListComputeV2ServersResponse(response *ListComputeV2ServersResponse)([]servers.Server,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return servers.ExtractServers(page)
}


// call result's extract function
func ExtractCreateComputeV2ServersResponse(response *CreateComputeV2ServersResponse)(interface{}){
    return response.CreateResult.Body
}


// call result's extract function
func ExtractDeleteComputeV2ServersResponse(response *DeleteComputeV2ServersResponse)(interface{}){
    return response.DeleteResult.Body
}


// call result's extract function
func ExtractForceDeleteComputeV2ServersResponse(response *ForceDeleteComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractGetComputeV2ServersResponse(response *GetComputeV2ServersResponse)(interface{}){
    return response.GetResult.Body
}


// call result's extract function
func ExtractUpdateComputeV2ServersResponse(response *UpdateComputeV2ServersResponse)(interface{}){
    return response.UpdateResult.Body
}


// call result's extract function
func ExtractChangeAdminPasswordComputeV2ServersResponse(response *ChangeAdminPasswordComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractRebootComputeV2ServersResponse(response *RebootComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractRebuildComputeV2ServersResponse(response *RebuildComputeV2ServersResponse)(interface{}){
    return response.RebuildResult.Body
}


// call result's extract function
func ExtractResizeComputeV2ServersResponse(response *ResizeComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractConfirmResizeComputeV2ServersResponse(response *ConfirmResizeComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractRevertResizeComputeV2ServersResponse(response *RevertResizeComputeV2ServersResponse)(interface{}){
    return response.ActionResult.Body
}


// call result's extract function
func ExtractResetMetadataComputeV2ServersResponse(response *ResetMetadataComputeV2ServersResponse)(interface{}){
    return response.ResetMetadataResult.Body
}


// call result's extract function
func ExtractMetadataComputeV2ServersResponse(response *MetadataComputeV2ServersResponse)(interface{}){
    return response.GetMetadataResult.Body
}


// call result's extract function
func ExtractUpdateMetadataComputeV2ServersResponse(response *UpdateMetadataComputeV2ServersResponse)(interface{}){
    return response.UpdateMetadataResult.Body
}


// call result's extract function
func ExtractCreateMetadatumComputeV2ServersResponse(response *CreateMetadatumComputeV2ServersResponse)(interface{}){
    return response.CreateMetadatumResult.Body
}


// call result's extract function
func ExtractMetadatumComputeV2ServersResponse(response *MetadatumComputeV2ServersResponse)(interface{}){
    return response.GetMetadatumResult.Body
}


// call result's extract function
func ExtractDeleteMetadatumComputeV2ServersResponse(response *DeleteMetadatumComputeV2ServersResponse)(interface{}){
    return response.DeleteMetadatumResult.Body
}


//extract response info from pager for ListAddressesComputeV2Servers
func ExtractListAddressesComputeV2ServersResponse(response *ListAddressesComputeV2ServersResponse)(map[string][]servers.Address,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return servers.ExtractAddresses(page)
}


//extract response info from pager for ListAddressesByNetworkComputeV2Servers
func ExtractListAddressesByNetworkComputeV2ServersResponse(response *ListAddressesByNetworkComputeV2ServersResponse)([]servers.Address,error){
	page, err := response.Pager.AllPages()
	if err != nil {
		return nil, err
	}
	return servers.ExtractNetworkAddresses(page)
}


// call result's extract function
func ExtractCreateImageComputeV2ServersResponse(response *CreateImageComputeV2ServersResponse)(interface{}){
    return response.CreateImageResult.Body
}


// call result's extract function
func ExtractGetPasswordComputeV2ServersResponse(response *GetPasswordComputeV2ServersResponse)(interface{}){
    return response.GetPasswordResult.Body
}


// call result's extract function
func ExtractShowConsoleOutputComputeV2ServersResponse(response *ShowConsoleOutputComputeV2ServersResponse)(interface{}){
    return response.ShowConsoleOutputResult.Body
}