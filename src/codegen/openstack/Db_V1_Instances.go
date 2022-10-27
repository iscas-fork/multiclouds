package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/db/v1/instances"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the CreateDbV1Instances
type CreateDbV1InstancesRequest struct{
    Opts instances.CreateOptsBuilder
}

func NewCreateDbV1InstancesRequest()*CreateDbV1InstancesRequest{
    return &CreateDbV1InstancesRequest{}
}

//response struct for the CreateDbV1Instances
type CreateDbV1InstancesResponse struct{
    CreateResult instances.CreateResult
}

func NewCreateDbV1InstancesResponse(createResult instances.CreateResult,)*CreateDbV1InstancesResponse {
    return &CreateDbV1InstancesResponse{
            CreateResult:createResult,
    }
}

// action function
func (oc *OpenstackClient) CreateDbV1Instances(request *CreateDbV1InstancesRequest)(*CreateDbV1InstancesResponse){
    return NewCreateDbV1InstancesResponse(instances.Create(oc.client,request.Opts, ))

}
//request struct for the ListDbV1Instances
type ListDbV1InstancesRequest struct{
}

func NewListDbV1InstancesRequest()*ListDbV1InstancesRequest{
    return &ListDbV1InstancesRequest{}
}

//response struct for the ListDbV1Instances
type ListDbV1InstancesResponse struct{
    Pager pagination.Pager
}

func NewListDbV1InstancesResponse(pager pagination.Pager,)*ListDbV1InstancesResponse {
    return &ListDbV1InstancesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListDbV1Instances(request *ListDbV1InstancesRequest)(*ListDbV1InstancesResponse){
    return NewListDbV1InstancesResponse(instances.List(oc.client, ))

}
//request struct for the GetDbV1Instances
type GetDbV1InstancesRequest struct{
    Id string
}

func NewGetDbV1InstancesRequest()*GetDbV1InstancesRequest{
    return &GetDbV1InstancesRequest{}
}

//response struct for the GetDbV1Instances
type GetDbV1InstancesResponse struct{
    GetResult instances.GetResult
}

func NewGetDbV1InstancesResponse(getResult instances.GetResult,)*GetDbV1InstancesResponse {
    return &GetDbV1InstancesResponse{
            GetResult:getResult,
    }
}

// action function
func (oc *OpenstackClient) GetDbV1Instances(request *GetDbV1InstancesRequest)(*GetDbV1InstancesResponse){
    return NewGetDbV1InstancesResponse(instances.Get(oc.client,request.Id, ))

}
//request struct for the DeleteDbV1Instances
type DeleteDbV1InstancesRequest struct{
    Id string
}

func NewDeleteDbV1InstancesRequest()*DeleteDbV1InstancesRequest{
    return &DeleteDbV1InstancesRequest{}
}

//response struct for the DeleteDbV1Instances
type DeleteDbV1InstancesResponse struct{
    DeleteResult instances.DeleteResult
}

func NewDeleteDbV1InstancesResponse(deleteResult instances.DeleteResult,)*DeleteDbV1InstancesResponse {
    return &DeleteDbV1InstancesResponse{
            DeleteResult:deleteResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteDbV1Instances(request *DeleteDbV1InstancesRequest)(*DeleteDbV1InstancesResponse){
    return NewDeleteDbV1InstancesResponse(instances.Delete(oc.client,request.Id, ))

}
//request struct for the EnableRootUserDbV1Instances
type EnableRootUserDbV1InstancesRequest struct{
    Id string
}

func NewEnableRootUserDbV1InstancesRequest()*EnableRootUserDbV1InstancesRequest{
    return &EnableRootUserDbV1InstancesRequest{}
}

//response struct for the EnableRootUserDbV1Instances
type EnableRootUserDbV1InstancesResponse struct{
    EnableRootUserResult instances.EnableRootUserResult
}

func NewEnableRootUserDbV1InstancesResponse(enableRootUserResult instances.EnableRootUserResult,)*EnableRootUserDbV1InstancesResponse {
    return &EnableRootUserDbV1InstancesResponse{
            EnableRootUserResult:enableRootUserResult,
    }
}

// action function
func (oc *OpenstackClient) EnableRootUserDbV1Instances(request *EnableRootUserDbV1InstancesRequest)(*EnableRootUserDbV1InstancesResponse){
    return NewEnableRootUserDbV1InstancesResponse(instances.EnableRootUser(oc.client,request.Id, ))

}
//request struct for the IsRootEnabledDbV1Instances
type IsRootEnabledDbV1InstancesRequest struct{
    Id string
}

func NewIsRootEnabledDbV1InstancesRequest()*IsRootEnabledDbV1InstancesRequest{
    return &IsRootEnabledDbV1InstancesRequest{}
}

//response struct for the IsRootEnabledDbV1Instances
type IsRootEnabledDbV1InstancesResponse struct{
    IsRootEnabledResult instances.IsRootEnabledResult
}

func NewIsRootEnabledDbV1InstancesResponse(isRootEnabledResult instances.IsRootEnabledResult,)*IsRootEnabledDbV1InstancesResponse {
    return &IsRootEnabledDbV1InstancesResponse{
            IsRootEnabledResult:isRootEnabledResult,
    }
}

// action function
func (oc *OpenstackClient) IsRootEnabledDbV1Instances(request *IsRootEnabledDbV1InstancesRequest)(*IsRootEnabledDbV1InstancesResponse){
    return NewIsRootEnabledDbV1InstancesResponse(instances.IsRootEnabled(oc.client,request.Id, ))

}
//request struct for the RestartDbV1Instances
type RestartDbV1InstancesRequest struct{
    Id string
}

func NewRestartDbV1InstancesRequest()*RestartDbV1InstancesRequest{
    return &RestartDbV1InstancesRequest{}
}

//response struct for the RestartDbV1Instances
type RestartDbV1InstancesResponse struct{
    ActionResult instances.ActionResult
}

func NewRestartDbV1InstancesResponse(actionResult instances.ActionResult,)*RestartDbV1InstancesResponse {
    return &RestartDbV1InstancesResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) RestartDbV1Instances(request *RestartDbV1InstancesRequest)(*RestartDbV1InstancesResponse){
    return NewRestartDbV1InstancesResponse(instances.Restart(oc.client,request.Id, ))

}
//request struct for the ResizeDbV1Instances
type ResizeDbV1InstancesRequest struct{
    Id string
    FlavorRef string
}

func NewResizeDbV1InstancesRequest()*ResizeDbV1InstancesRequest{
    return &ResizeDbV1InstancesRequest{}
}

//response struct for the ResizeDbV1Instances
type ResizeDbV1InstancesResponse struct{
    ActionResult instances.ActionResult
}

func NewResizeDbV1InstancesResponse(actionResult instances.ActionResult,)*ResizeDbV1InstancesResponse {
    return &ResizeDbV1InstancesResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ResizeDbV1Instances(request *ResizeDbV1InstancesRequest)(*ResizeDbV1InstancesResponse){
    return NewResizeDbV1InstancesResponse(instances.Resize(oc.client,request.Id,request.FlavorRef, ))

}
//request struct for the ResizeVolumeDbV1Instances
type ResizeVolumeDbV1InstancesRequest struct{
    Id string
    Size int
}

func NewResizeVolumeDbV1InstancesRequest()*ResizeVolumeDbV1InstancesRequest{
    return &ResizeVolumeDbV1InstancesRequest{}
}

//response struct for the ResizeVolumeDbV1Instances
type ResizeVolumeDbV1InstancesResponse struct{
    ActionResult instances.ActionResult
}

func NewResizeVolumeDbV1InstancesResponse(actionResult instances.ActionResult,)*ResizeVolumeDbV1InstancesResponse {
    return &ResizeVolumeDbV1InstancesResponse{
            ActionResult:actionResult,
    }
}

// action function
func (oc *OpenstackClient) ResizeVolumeDbV1Instances(request *ResizeVolumeDbV1InstancesRequest)(*ResizeVolumeDbV1InstancesResponse){
    return NewResizeVolumeDbV1InstancesResponse(instances.ResizeVolume(oc.client,request.Id,request.Size, ))

}
//request struct for the AttachConfigurationGroupDbV1Instances
type AttachConfigurationGroupDbV1InstancesRequest struct{
    InstanceID string
    ConfigID string
}

func NewAttachConfigurationGroupDbV1InstancesRequest()*AttachConfigurationGroupDbV1InstancesRequest{
    return &AttachConfigurationGroupDbV1InstancesRequest{}
}

//response struct for the AttachConfigurationGroupDbV1Instances
type AttachConfigurationGroupDbV1InstancesResponse struct{
    ConfigurationResult instances.ConfigurationResult
}

func NewAttachConfigurationGroupDbV1InstancesResponse(configurationResult instances.ConfigurationResult,)*AttachConfigurationGroupDbV1InstancesResponse {
    return &AttachConfigurationGroupDbV1InstancesResponse{
            ConfigurationResult:configurationResult,
    }
}

// action function
func (oc *OpenstackClient) AttachConfigurationGroupDbV1Instances(request *AttachConfigurationGroupDbV1InstancesRequest)(*AttachConfigurationGroupDbV1InstancesResponse){
    return NewAttachConfigurationGroupDbV1InstancesResponse(instances.AttachConfigurationGroup(oc.client,request.InstanceID,request.ConfigID, ))

}
//request struct for the DetachConfigurationGroupDbV1Instances
type DetachConfigurationGroupDbV1InstancesRequest struct{
    InstanceID string
}

func NewDetachConfigurationGroupDbV1InstancesRequest()*DetachConfigurationGroupDbV1InstancesRequest{
    return &DetachConfigurationGroupDbV1InstancesRequest{}
}

//response struct for the DetachConfigurationGroupDbV1Instances
type DetachConfigurationGroupDbV1InstancesResponse struct{
    ConfigurationResult instances.ConfigurationResult
}

func NewDetachConfigurationGroupDbV1InstancesResponse(configurationResult instances.ConfigurationResult,)*DetachConfigurationGroupDbV1InstancesResponse {
    return &DetachConfigurationGroupDbV1InstancesResponse{
            ConfigurationResult:configurationResult,
    }
}

// action function
func (oc *OpenstackClient) DetachConfigurationGroupDbV1Instances(request *DetachConfigurationGroupDbV1InstancesRequest)(*DetachConfigurationGroupDbV1InstancesResponse){
    return NewDetachConfigurationGroupDbV1InstancesResponse(instances.DetachConfigurationGroup(oc.client,request.InstanceID, ))

}