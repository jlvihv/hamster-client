package resource

import (
	"fmt"
	"hamster-client/utils"
	"testing"
)

func TestHttpService_GetResourceList(t *testing.T) {
	service := NewServiceImpl(nil, utils.NewHttp())
	publicKey := "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDHmVkaKCU9lkeYB4bScoyrjGFk0M9wYFwfFEI1lXsWYCjuFWGtIwd8wHT+ApGsUY83y/5AWjxRhO6yNjK+9mDE+UxLYNvaDhWEsw9UlXCKLg7gcDgZmdLkr5ellJjuW5crsMWMwCg7CNigL6vxCTLu5lzgbeqJT09W7/hhyn7pcFw7ggul5O530hhvv+rvhZtkurHVOUHp8ofiSZdOtmp8ZgDtsP6CzsAUgiyBSswaP4TFyrn+USbr9Vt9A+qhCBa2RT2ADPY3jEPvslf9wnzYsiROdzneb9AB+HCEsr53CeGHAym1DcW0wZ2Nwf2k0rXVtyCUt1O2zptkbhl5Gav5PYYhsC0TYUJdUDT1ZELv4Uqbs7Nw8ILOk7PY/R6wcu3ZxBqOdbGc7FKRbaPkPzSC3ozt8yQ2S7WLDTTNgHICKXBYTCDFFCEcV9d0MxMwvIGZ67MMKb71Udmh1BdlBODbNQdEtXea07GkFHZ6bjbR2oOOBEuoXwWzYx9JtJiolrM= lzw@MacBook-Air.local"
	list, err := service.GetResourceList(publicKey)
	if err != nil {
		fmt.Println("err", err)
	}
	fmt.Println(list)
}
