// Author: xufei
// Date: 2020/3/11

package main

import (
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", "config/rbac_policy.csv")
	if err != nil {
		panic(err)
	}

	if ok, err := enforcer.Enforce("admin", "data1", "read"); err != nil {
		log.Printf("admin policy failed: %v", err)
	} else if !ok {
		log.Printf("admin policy data1 wrong!")
	}

	if ok, err := enforcer.Enforce("admin", "data2", "read"); err != nil {
		log.Printf("admin policy failed: %v", err)
	} else if !ok {
		log.Printf("admin policy data2 wrong!")
	}
}
