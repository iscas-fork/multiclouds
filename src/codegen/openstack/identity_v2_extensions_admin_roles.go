package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/identity/v2/extensions/admin/roles"
    "github.com/gophercloud/gophercloud/pagination"
)
//request struct for the ListIdentityV2ExtensionsAdminRoles
type ListIdentityV2ExtensionsAdminRolesRequest struct{
}

func NewListIdentityV2ExtensionsAdminRolesRequest()*ListIdentityV2ExtensionsAdminRolesRequest{
    return &ListIdentityV2ExtensionsAdminRolesRequest{}
}

//response struct for the ListIdentityV2ExtensionsAdminRoles
type ListIdentityV2ExtensionsAdminRolesResponse struct{
    Pager pagination.Pager
}

func NewListIdentityV2ExtensionsAdminRolesResponse(pager pagination.Pager,)*ListIdentityV2ExtensionsAdminRolesResponse {
    return &ListIdentityV2ExtensionsAdminRolesResponse{
            Pager:pager,
    }
}

// action function
func (oc *OpenstackClient) ListIdentityV2ExtensionsAdminRoles(request *ListIdentityV2ExtensionsAdminRolesRequest)(*ListIdentityV2ExtensionsAdminRolesResponse){
    return NewListIdentityV2ExtensionsAdminRolesResponse(roles.List(oc.client, ))

}
//request struct for the AddUserIdentityV2ExtensionsAdminRoles
type AddUserIdentityV2ExtensionsAdminRolesRequest struct{
    TenantID string
    UserID string
    RoleID string
}

func NewAddUserIdentityV2ExtensionsAdminRolesRequest()*AddUserIdentityV2ExtensionsAdminRolesRequest{
    return &AddUserIdentityV2ExtensionsAdminRolesRequest{}
}

//response struct for the AddUserIdentityV2ExtensionsAdminRoles
type AddUserIdentityV2ExtensionsAdminRolesResponse struct{
    UserRoleResult roles.UserRoleResult
}

func NewAddUserIdentityV2ExtensionsAdminRolesResponse(userRoleResult roles.UserRoleResult,)*AddUserIdentityV2ExtensionsAdminRolesResponse {
    return &AddUserIdentityV2ExtensionsAdminRolesResponse{
            UserRoleResult:userRoleResult,
    }
}

// action function
func (oc *OpenstackClient) AddUserIdentityV2ExtensionsAdminRoles(request *AddUserIdentityV2ExtensionsAdminRolesRequest)(*AddUserIdentityV2ExtensionsAdminRolesResponse){
    return NewAddUserIdentityV2ExtensionsAdminRolesResponse(roles.AddUser(oc.client,request.TenantID,request.UserID,request.RoleID, ))

}
//request struct for the DeleteUserIdentityV2ExtensionsAdminRoles
type DeleteUserIdentityV2ExtensionsAdminRolesRequest struct{
    TenantID string
    UserID string
    RoleID string
}

func NewDeleteUserIdentityV2ExtensionsAdminRolesRequest()*DeleteUserIdentityV2ExtensionsAdminRolesRequest{
    return &DeleteUserIdentityV2ExtensionsAdminRolesRequest{}
}

//response struct for the DeleteUserIdentityV2ExtensionsAdminRoles
type DeleteUserIdentityV2ExtensionsAdminRolesResponse struct{
    UserRoleResult roles.UserRoleResult
}

func NewDeleteUserIdentityV2ExtensionsAdminRolesResponse(userRoleResult roles.UserRoleResult,)*DeleteUserIdentityV2ExtensionsAdminRolesResponse {
    return &DeleteUserIdentityV2ExtensionsAdminRolesResponse{
            UserRoleResult:userRoleResult,
    }
}

// action function
func (oc *OpenstackClient) DeleteUserIdentityV2ExtensionsAdminRoles(request *DeleteUserIdentityV2ExtensionsAdminRolesRequest)(*DeleteUserIdentityV2ExtensionsAdminRolesResponse){
    return NewDeleteUserIdentityV2ExtensionsAdminRolesResponse(roles.DeleteUser(oc.client,request.TenantID,request.UserID,request.RoleID, ))

}