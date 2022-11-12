package main

import (
	"fmt"
	"sort"
)

type Tenant struct {
	Name     string
	id       string
	clientId string
}

type Tenants []Tenant

func (a Tenants) Len() int {
	return len(a)
}

func (a Tenants) Less(i, j int) bool {
	if a[i].Name == a[j].Name {
		return a[i].id < a[j].id
	}
	return a[i].Name < a[j].Name
}

func (a Tenants) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	var tenants = Tenants{}
	tenants = append(tenants, Tenant{
		Name: "Jose Manuel",
		id:   "abc",
	})
	tenants = append(tenants, Tenant{
		Name: "Jose Manuel",
		id:   "dae",
	})
	tenants = append(tenants, Tenant{
		Name: "Joseph Manel",
		id:   "aba",
	})
	tenants = append(tenants, Tenant{
		Name: "Ioseph Manel",
		id:   "ddd",
	})
	sort.Sort(Tenants(tenants))
	for _, v := range tenants {
		fmt.Println(v)
	}
}
