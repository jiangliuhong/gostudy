package study

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceServerGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext:   resourceServerGroupRead,
		CreateContext: resourceServerGroupCreate,
		DeleteContext: resourceServerGroupDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
			},
		},
	}
}

func resourceServerGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	fmt.Sprintln(data.Get("name"))
	data.SetId(data.Get("name").(string))
	return nil
}
func resourceServerGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return resourceServerGroupRead(ctx, data, meta)
}

func resourceServerGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
