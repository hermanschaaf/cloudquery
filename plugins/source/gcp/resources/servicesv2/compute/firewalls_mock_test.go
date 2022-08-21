// Code generated by codegen; DO NOT EDIT.

package compute

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	faker "github.com/cloudquery/faker/v3"
	"github.com/cloudquery/plugins/source/gcp/client"
	"github.com/julienschmidt/httprouter"

	"google.golang.org/api/compute/v1"

	"google.golang.org/api/option"
)

func createComputeFirewalls() (*client.Services, error) {
	var item compute.Firewall
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/projects/testProject/global/firewalls", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := &compute.FirewallList{
			Items: []*compute.Firewall{&item},
		}
		b, err := json.Marshal(resp)
		if err != nil {
			http.Error(w, "unable to marshal request: "+err.Error(), http.StatusBadRequest)
			return
		}
		if _, err := w.Write(b); err != nil {
			http.Error(w, "failed to write", http.StatusBadRequest)
			return
		}
	})
	ts := httptest.NewServer(mux)
	svc, err := compute.NewService(context.Background(), option.WithoutAuthentication(), option.WithEndpoint(ts.URL))
	if err != nil {
		return nil, err
	}
	return &client.Services{
		Compute: svc,
	}, nil
}

func TestComputeFirewalls(t *testing.T) {
	client.GcpMockTestHelper(t, ComputeFirewalls(), createComputeFirewalls, client.TestOptions{})
}
