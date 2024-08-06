package _request

import "golang.org/x/net/context"

type SessionContext struct {
	UserId    uint64
	ProjectId uint64
	TenantId  uint64
	Context   context.Context
}
