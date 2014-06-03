package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/igm/vim25"
)

func init() {
	commands["vmPowerOff"] = vmPowerOff
}

func vmPowerOff() {
	vmName := flag.Arg(1)
	if vmName == "" {
		log.Fatal("Missing Virtual Machine Obj reference value as param, i.e. vm-123")
	}
	service := vim25.VimService{URL: VSPHERE_URL}
	response := new(vim25.RetrieveServiceContentResponse)
	if err := service.Invoke(vim25.RetrieveServiceContent{This: vim25.ServiceInstanceRef}, response); err != nil {
		fmt.Println(err)
	}
	sc := response.ServiceContent

	if err := service.Login(sc.SessionManager, VSPHERE_LOGIN, VSPHERE_PASS); err != nil {
		log.Fatal(err)
	}

	var vmRef vim25.VirtualMachine = &vim25.ManagedObjectReference{"VirtualMachine", vmName}
	request := &vim25.PowerOffVM_Task{
		This: vmRef,
	}
	resp := new(vim25.PowerOffVm_TaskResponse)
	if err := service.Invoke(request, resp); err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Task)

}