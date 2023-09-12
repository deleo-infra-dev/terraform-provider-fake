package main

import (
  "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var testAccProvider *schema.Provider
var testAccProviders map[string]*schema.Provider

func init() {
  testAccProvider = Provider()
  testAccProviders = map[string]*schema.Provider{
    "fake": testAccProvider,
  }
}