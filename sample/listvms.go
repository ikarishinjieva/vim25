package main

import (
	"fmt"
	"log"

	"github.com/igm/vim25"
)

func init() {
	commands["listvms"] = listVMs
}

func listVMs() {
	si := &vim25.ManagedObjectReference{"ServiceInstance", "ServiceInstance"}
	service := vim25.VimService{URL: VSPHERE_URL}

	response := new(vim25.RetrieveServiceContentResponse)
	if err := service.Invoke(vim25.RetrieveServiceContent{This: si}, response); err != nil {
		log.Fatal(err)
	}

	sc := response.ServiceContent

	loginResponse := new(vim25.LoginResponse)
	login := vim25.Login{
		This:     sc.SessionManager,
		Username: VSPHERE_LOGIN,
		Password: VSPHERE_PASS,
	}

	if err := service.Invoke(login, loginResponse); err != nil {
		log.Fatal(err)
	}

	ccv := &vim25.CreateContainerView{
		This:      sc.ViewManager,
		Container: sc.RootFolder,
		Type:      []string{"VirtualMachine"},
		Recursive: true,
	}

	containerViewResponse := new(vim25.CreateContainerViewResponse)
	if err := service.Invoke(ccv, containerViewResponse); err != nil {
		log.Fatal(err)
	}

	oSpec := &vim25.ObjectSpec{
		Obj:  containerViewResponse.ContainerView,
		Skip: true,
	}

	tSpec := &vim25.TraversalSpec{
		SelectionSpec: vim25.SelectionSpec{Name: "traverseEntities"},
		Path:          "view",
		Skip:          false,
		Type:          "ContainerView",
		TypeAttr:      "TraversalSpec",
	}
	oSpec.SelectSet = append(oSpec.SelectSet, tSpec)

	pSpec := &vim25.PropertySpec{
		Type:    "VirtualMachine",
		PathSet: []string{"name"},
	}

	fSpec := &vim25.PropertyFilterSpec{
		ObjectSet: []*vim25.ObjectSpec{oSpec},
		PropSet:   []*vim25.PropertySpec{pSpec},
	}

	rpse := vim25.RetrievePropertiesEx{
		This:    sc.PropertyCollector,
		SpecSet: []*vim25.PropertyFilterSpec{fSpec},
		Options: vim25.RetrieveOptions{},
	}

	rep := new(vim25.RetrievePropertiesExResponse)
	if err := service.Invoke(rpse, rep); err != nil {
		log.Fatal(err)
	}
	for _, rep := range rep.RetrieveResult.Objects {
		fmt.Println(rep.Obj.Type, rep.Obj.Value)
		for _, prop := range rep.PropSet {
			fmt.Printf("\t %s %s\n", prop.Name, prop.Val.Value)
		}
	}
}
