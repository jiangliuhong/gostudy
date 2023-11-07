package study

import (
	"context"
	"encoding/json"
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
			"test_array": {
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

type TestArrayItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func resourceServerGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	log.Printf("test read")
	fmt.Sprintln(data.Get("name"))
	data.SetId(data.Get("name").(string))
	return nil
}
func resourceServerGroupCreate(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	data.Set("ips", []string{})
	tra := []*TestArrayItem{
		{Id: "1", Name: "2"},
	}
	//转为map数组
	b, _ := json.Marshal(tra)
	var testArray []map[string]interface{}
	_ = json.Unmarshal(b, &testArray)
	err := data.Set("test_array", &testArray)
	if err != nil {
		return diag.FromErr(err)
	}
	return resourceServerGroupRead(ctx, data, meta)
}

func resourceServerGroupDelete(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}
