package healthcheck

import (
	"fmt"
	types "github.com/skyscape-cloud-services/vmware-govcd/types/v56"
)

func NetworkDevice(s types.QueryResultVMRecordType, vm *types.VM) (string, error) {

	device_type := "unknown"

	for _, v := range vm.VirtualHardwareSection.Item {
		if v.ResourceType == 10 {
			device_type = v.ResourceSubType
		}
	}
	if device_type != "VMXNET3" {
		return device_type, fmt.Errorf("VMs NIC is: %s", device_type)
	}

	return device_type, nil
}
