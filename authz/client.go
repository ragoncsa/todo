package authz

import (
	"encoding/json"

	"github.com/go-resty/resty/v2"
	"github.com/ragoncsa/todo/config"
)

type decisionResult struct {
	Allow bool `json:"allow"`
}

type decision struct {
	Result     *decisionResult `json:"result"`
	DecisionId string          `json:"decision_id"`
}

type DecisionRequest struct {
	Method string   `json:"method"`
	Path   []string `json:"path"`
	Owner  string   `json:"owner"`
	User   string   `json:"user"`
	TaskID string   `json:"taskid"`
}

type decisionReqInternal struct {
	Input *DecisionRequest
}

type Client interface {
	IsAllowed(dreq *DecisionRequest) (bool, error)
}

type client struct {
	restClient *resty.Client
	endpoint   string
}

func New(conf *config.Config) Client {
	return &client{
		restClient: resty.New(),
		endpoint:   conf.Authz.Endpoint,
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
		SetResult(&decision{}).
		Post(c.endpoint)
	if err != nil || resp.IsError() {
		return false, err
	}
	return resp.Result().(*decision).Result.Allow, nil
}
