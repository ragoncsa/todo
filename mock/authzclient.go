package mock

import "github.com/ragoncsa/todo/authz"

type AlwaysAllow struct{}

func (c *AlwaysAllow) IsAllowed(dreq *authz.DecisionRequest) (bool, error) {
	return true, nil
}

type AlwaysDeny struct{}

func (c *AlwaysDeny) IsAllowed(dreq *authz.DecisionRequest) (bool, error) {
	return false, nil
}
