package cmd

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	kld "github.com/kaleido-io/kaleido-sdk-go/kaleido"
	"github.com/spf13/cobra"
	resty "gopkg.in/resty.v1"
)

func TestAppCreds_CreateCmd(t *testing.T) {

	called := false
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()
	fakeAPIURL := "https://fake-console.kaleido.io/api/v1"
	consortiumID := "u0m7e3itin"
	environmentID := "e3itinu0m7"
	membershipID := "u0ypwvdcan"

	fakeAppcredsAPIURL := fmt.Sprintf("%v/consortia/%v/environments/%v/appcreds", fakeAPIURL,
		consortiumID, environmentID)
	httpmock.RegisterResponder("POST", fakeAppcredsAPIURL, func(req *http.Request) (*http.Response, error) {
		responseString := fmt.Sprintf(`{"membership_id":"%v","auth_type":"basic_auth",
			"_id":"u0k719wsku","_revision":"0","username":"u0k719wsku",
			"password":"uaL7VZi0j4VYvxEhzNr-A2T5QdaPFt8YltEU_RgB_ac"}
			{"membership_id":"u0ypwvdcan","auth_type":"basic_auth",
				"_id":"u0nqupl03f","_revision":"0","username":"u0nqupl03f",
				"password":"pga09rvrQPDsLs7Xd7P0zQZ7w6wjSKDxZR_9HFZg8CQ"}`, membershipID)
		reqPayload, _ := ioutil.ReadAll(req.Body)

		payload := fmt.Sprintf(`{"membership_id":"%v"}`, membershipID)
		if string(reqPayload) != payload {
			t.Errorf("The request should have the payload %q, but got %q", payload, string(reqPayload))
		}
		resp := httpmock.NewStringResponse(201, responseString)
		called = true
		return resp, nil
	},
	)

	getNewClient = func() kld.KaleidoClient {
		r := resty.New().SetHostURL(fakeAPIURL).SetAuthToken("APIKey")
		return kld.KaleidoClient{r}
	}

	appCredsCreateCmd.ParseFlags([]string{"-c", consortiumID, "-e", environmentID, "-m", membershipID})
	appCredsCreateCmd.Run(&cobra.Command{}, []string{})

	if !called {
		t.Errorf("CreateAppCreds should have been called")
	}
}
