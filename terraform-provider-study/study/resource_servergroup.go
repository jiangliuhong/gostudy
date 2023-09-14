package study

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
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
			"ips": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceServerGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("test read")
	fmt.Sprintln(data.Get("name"))
	data.SetId(data.Get("name").(string))
	return nil
}
func resourceServerGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	data.Set("ips", []string{})
	return resourceServerGroupRead(ctx, data, meta)
}

func resourceServerGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
