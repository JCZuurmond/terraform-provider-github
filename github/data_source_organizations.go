package github

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceOrganizations() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceOrganizationsRead,
		Schema: map[string]*schema.Schema{
			"organizations": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"login": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"node_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"repos_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"events_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"hooks_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"issues_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"members_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"public_members_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"avatar_url": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceOrganizationsRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/organizations", "https://api.github.com"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	organizations := make([]map[string]interface{}, 0)
	err = json.NewDecoder(r.Body).Decode(&organizations)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("organizations", organizations); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
