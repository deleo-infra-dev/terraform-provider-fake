package main

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func resourceFake() *schema.Resource {
	return &schema.Resource{
		Create: createFake,
		Read:   readFake,
		Update: createFake,
		Delete: deleteFake,

		Schema: map[string]*schema.Schema{
			"value": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "plop",
			},
		},
	}
}

func createFake(d *schema.ResourceData, m interface{}) error {
	d.SetId("toto")
	return nil
}

func readFake(d *schema.ResourceData, m interface{}) error {
	return nil
}

func deleteFake(d *schema.ResourceData, m interface{}) error {
	return nil
}
