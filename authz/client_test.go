package authz

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/jarcoal/httpmock"
)

func TestIsAllowed(t *testing.T) {
	const url = "http://myurl"
	client := &client{
		restClient:    resty.New(),
		authzEndpoint: url,
	}
	httpmock.ActivateNonDefault(client.restClient.GetClient())
	defer httpmock.DeactivateAndReset()
	httpmock.RegisterResponder("POST", url,
		func(req *http.Request) (*http.Response, error) {
			var drIn decisionReqInternal
			if err := json.NewDecoder(req.Body).Decode(&drIn); err != nil {
				t.Fatalf("Parsing incoming request to authz endpoint failed: %v", err)
			}
			var resp *http.Response
			var err error
			if toJson(t, drIn.Input) == toJson(t, &DecisionRequest{
				Method: "POST",
				Owner:  "johndoe",
				Path:   []string{"tasks"},
				User:   "johndoe",
			}) {
				resp, err = httpmock.NewJsonResponse(200, &authzDecision{
					DecisionId: "798d0f89-3ba0-4662-87b1-c2a8adb8f62a",
					Result: &authzDecisionResult{
						Allow: true,
					},
				})
			} else {
				// just returning error for everything else for simplicity
				resp, err = httpmock.NewJsonResponse(400, struct{}{})
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			return resp, nil
		},
	)

	allowed, err := client.IsAllowed(&DecisionRequest{
		Method: "POST",
		Owner:  "johndoe",
		Path:   []string{"tasks"},
		User:   "johndoe",
	})
	if err != nil {
		t.Fatalf("isAllowed err: %v", err)
	}

	if got, want := allowed, true; got != want {
		t.Fatalf("isAllowed got %v want: %v", got, want)
	}

}

func toJson(t *testing.T, d interface{}) string {
	jsonStr, err := json.Marshal(d)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	return string(jsonStr)
}
