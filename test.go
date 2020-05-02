package github-api
import "github.com/paladium/go-github/v31/github"

func TestCanListOrganizationsForUser(t *testing.T){
	client := github.NewClient(nil)
	orgs, _, err := client.Organizations.List(context.Background(), "paladium", nil)
	assert.Nil(t, err)

}