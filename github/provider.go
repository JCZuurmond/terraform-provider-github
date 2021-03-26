package github

import (
	"context"

	"golang.org/x/oauth2"

	"github.com/google/go-github/v34/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    false,
				Sensitive:   true,
				DefaultFunc: schema.EnvDefaultFunc("GITHUB_TOKEN", nil),
			},
		},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	token := d.Get("token").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (token != "") {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client, err := github.NewClient(tc)

		if err != nil {
			return nil, diag.FromErr(err)
		}

		// list all repositories for the authenticated user
		repos, _, err := client.Repositories.List(ctx, "", nil)

		return client, diags
	}

	return nil, diags
}
