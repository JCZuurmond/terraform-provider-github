package github

import(
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"

	"github.com/google/go-github/v34/github"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRepository() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRepositoryCreate,
		ReadContext: resourceRepositoryRead,
		DeleteContext: resourceRepositoryDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type: schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"organization": {
				Type: schema.TypeString,
				ForceNew: true,
				Default: "",
				Optional: true,
			},
			"private": {
				Type: schema.TypeBool,
				Required: true,
				ForceNew: true,
			},
		},
	}
}

func resourceRepositoryCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*github.Client) // TODO: what's happening here?

	repo := &github.Repository{
		Name: github.String(d.Get("name").(string)),
		Private: github.Bool(d.Get("private").(bool)),
	}

	d.SetId(repo.GetName())

	repo, _, _ = client.Repositories.Create(ctx, d.Get("organization").(string), repo)

	resourceRepositoryRead(ctx, d, m)

	return diags
}

func resourceRepositoryRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	client := m.(*github.Client) // TODO: what's happening here?

	repo, _, _ := client.Repositories.Get(ctx, d.Get("organization").(string), d.Id())
	d.Set("name", d.Id())
	d.Set("private", repo.GetPrivate())
	d.Set("organization", d.Get("organization").(string))

	return diags
}

func resourceRepositoryDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	return diags
}
