package openstack

// Code generated by cloud manager.





import (
    "github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/lockunlock"
)
//request struct for the LockLockunlock
type LockLockunlockRequest struct{
    id string
}

func NewLockLockunlockRequest()*LockLockunlockRequest{
    return &LockLockunlockRequest{}
}

//response struct for the LockLockunlock
type LockLockunlockResponse struct{
    LockResult lockunlock.LockResult
}

func NewLockLockunlockResponse(lockResult lockunlock.LockResult,)*LockLockunlockResponse {
    return &LockLockunlockResponse{
            LockResult:lockResult,
    }
}

// action function
func (oc *OpenstackClient) LockLockunlock(request *LockLockunlockRequest)(*LockLockunlockResponse){
    return NewLockLockunlockResponse(lockunlock.Lock(oc.client,request.id, ))

}
//request struct for the UnlockLockunlock
type UnlockLockunlockRequest struct{
    id string
}

func NewUnlockLockunlockRequest()*UnlockLockunlockRequest{
    return &UnlockLockunlockRequest{}
}

//response struct for the UnlockLockunlock
type UnlockLockunlockResponse struct{
    UnlockResult lockunlock.UnlockResult
}

func NewUnlockLockunlockResponse(unlockResult lockunlock.UnlockResult,)*UnlockLockunlockResponse {
    return &UnlockLockunlockResponse{
            UnlockResult:unlockResult,
    }
}

// action function
func (oc *OpenstackClient) UnlockLockunlock(request *UnlockLockunlockRequest)(*UnlockLockunlockResponse){
    return NewUnlockLockunlockResponse(lockunlock.Unlock(oc.client,request.id, ))

}