package study

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func datasourceServerGroup() *schema.Resource {
	return &schema.Resource{
		ReadContext: datasourceServerGroupRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"policys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"exists": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func datasourceServerGroupRead(ctx context.Context, data *schema.ResourceData, meta interface{}) diag.Diagnostics {
	name := data.Get("name").(string)
	sg := &ServerGroup{}
	f := serverGroupsGroupEachPageHandleToFindByName(name, sg)
	b, err := f()
	if err != nil {
		return diag.FromErr(err)
	}
	log.Println("b:", b)
	log.Println("sg:", sg)
	if b {
		data.SetId(sg.Name)
		err := data.Set("policys", sg.Policys)
		if err != nil {
			return diag.Errorf("Error set policies value: %s", err)
		}
		data.Set("exists", true)
	} else {
		data.Set("exists", false)
	}

	return nil
}

type ServerGroup struct {
	Name    string
	Policys []string
}

func serverGroupsGroupEachPageHandleToFindByName(name string, sg *ServerGroup) func() (bool, error) {
	return func() (bool, error) {
		group := ServerGroup{
			Name:    name,
			Policys: []string{"test1", "test2"},
		}
		sg.Name = group.Name
		sg.Policys = group.Policys
		return true, nil
	}
}
