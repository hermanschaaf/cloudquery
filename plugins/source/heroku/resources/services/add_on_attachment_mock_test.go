// Code generated by codegen; DO NOT EDIT.

package services

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/heroku/client"
	faker "github.com/cloudquery/faker/v3"
	heroku "github.com/heroku/heroku-go/v5"
	"github.com/julienschmidt/httprouter"
)

func createAddOnAttachments() (client.HerokuService, error) {
	var item heroku.AddOnAttachment
	if err := faker.FakeData(&item); err != nil {
		return nil, err
	}
	mux := httprouter.New()
	mux.GET("/*any", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		resp := heroku.AddOnAttachmentListResult{
			item,
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
	s := heroku.NewService(ts.Client())
	s.URL = ts.URL
	return s, nil
}

func TestAddOnAttachment(t *testing.T) {
	client.HerokuMockTestHelper(t, AddOnAttachments(), createAddOnAttachments, client.TestOptions{})
}
