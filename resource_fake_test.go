package main

import (
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccFake__empty(t *testing.T) {
	config := `
resource "fake" "test" {
  value = ""
}
`
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("fake.test", "id", "toto"),
					resource.TestCheckResourceAttr("fake.test", "value", ""),
				),
			},
		},
	})
}

func TestAccFake__default_value(t *testing.T) {
	config := `
resource "fake" "test" {
  value = ""
}
`
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("fake.test", "id", "toto"),
					resource.TestCheckResourceAttr("fake.test", "value", "plop"),
				),
			},
		},
	})
}
