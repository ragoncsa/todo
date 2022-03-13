package authz

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/ragoncsa/todo/config"
)

type authzDecisionResult struct {
	Allow bool `json:"allow"`
}

type authzDecision struct {
	Result     *authzDecisionResult `json:"result"`
	DecisionId string               `json:"decision_id"`
}

type DecisionRequest struct {
	Method string   `json:"method"`
	Path   []string `json:"path"`
	Owner  string   `json:"owner"`
	User   string   `json:"user"`
}

type decisionReqInternal struct {
	Input *DecisionRequest
}

type Client interface {
	IsAllowed(dreq *DecisionRequest) (bool, error)
}

type client struct {
	restClient    *resty.Client
	authzEndpoint string
}

func New(conf *config.Config) Client {
	return &client{
		restClient:    resty.New(),
		authzEndpoint: conf.Authz.Endpoint,
	}
}

func (c *client) IsAllowed(dreq *DecisionRequest) (bool, error) {
	dreqStr, err := json.Marshal(&decisionReqInternal{Input: dreq})
	if err != nil {
		return false, err
	}
	resp, err := c.restClient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(dreqStr).
		SetResult(&authzDecision{}).
		Post(c.authzEndpoint)
	// rstr := string(resp.Body())
	// _ = rstr
	if err != nil || resp.IsError() {
		return false, err
	}
	return resp.Result().(*authzDecision).Result.Allow, nil
}
