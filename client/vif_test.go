package client

import (
	"os"
	"testing"
)

func TestGetVIFs(t *testing.T) {

	c, err := NewClient(GetConfigFromEnv())

	if err != nil {
		t.Fatalf("failed to create client with error: %v", err)
	}

	vm := findVmForTest(c, t)

	vifs, err := c.GetVIFs(&vm)

	for _, vif := range vifs {
		if vif.Device == "" {
			t.Errorf("expecting `Device` field to be set on VIF")
		}

		if vif.MacAddress == "" {
			t.Errorf("expecting `MacAddress` field to be set on VIF")
		}

		if vif.Network == "" {
			t.Errorf("expecting `Network` field to be set on VIF")
		}

		if vif.VmId != vm.Id {
			t.Errorf("VIF's VmId `%s` should have matched: %v", vif.VmId, vm)
		}

		if len(vif.Device) == 0 {
			t.Errorf("expecting `Device` field to be set on VIF instead received: %s", vif.Device)
		}

		// 		if !vif.Attached {
		// 			t.Errorf("expecting `Attached` field to be true on VIF instead received: %t", vif.Attached)
		// 		}
	}
}

func TestGetVIF(t *testing.T) {

	c, err := NewClient(GetConfigFromEnv())

	if err != nil {
		t.Fatalf("failed to create client with error: %v", err)
	}

	vm := findVmForTest(c, t)

	vifs, err := c.GetVIFs(&vm)

	expectedVIF := vifs[0]

	vif, err := c.GetVIF(&VIF{
		MacAddress: expectedVIF.MacAddress,
	})

	if err != nil {
		t.Fatalf("failed to get VIF with error: %v", err)
	}

	if vif.MacAddress != expectedVIF.MacAddress {
		t.Errorf("expected VIF: %v does not match the VIF we received %v", expectedVIF, vif)
	}
}

func findVmForTest(c *Client, t *testing.T) Vm {

	vms, err := c.GetVms()

	if err != nil {
		t.Fatalf("failed to get VMs with error: %v", err)
	}

	if len(vms) < 1 {
		t.Fatalf("request returned %d vm results, expected atleast 1", len(vms))
	}
	return vms[0]
}

func TestCreateVIF_DeleteVIF(t *testing.T) {
	if p := os.Getenv("XOA_POOL"); p != "xenserver-ddelnano" {
		t.Skip("This test isn't safe to run outside of my personal deployment yet!")
	}

	c, err := NewClient(GetConfigFromEnv())

	if err != nil {
		t.Fatalf("failed to create client with error: %v", err)
	}

	vmName := "XOA"
	vm, err := c.GetVm(Vm{NameLabel: vmName})

	if err != nil {
		t.Fatalf("failed to get VM with error: %v", err)
	}

	pifs, err := c.GetPIFByDevice("eth1", -1)

	if err != nil {
		t.Fatalf("failed to get PIF with error: %v", err)
	}

	pif := pifs[0]

	vif, err := c.CreateVIF(vm, &VIF{Network: pif.Network})

	if err != nil {
		t.Fatalf("failed to create VIF with error: %v", err)
	}

	err = c.DeleteVIF(vif)

	if err != nil {
		t.Errorf("failed to delete the VIF with error: %v", err)
	}
}
