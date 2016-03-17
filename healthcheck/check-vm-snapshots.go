package healthcheck

import (
	"fmt"
	"time"
	types "github.com/skyscape-cloud-services/vmware-govcd/types/v56"
)

func VMSnapshots(s types.QueryResultVMRecordType, vm *types.VM) (string, error) {

	SnapshotCount := 0
	OldSnapshots := 0
	MaxSnapshotAge := 7

	for _, snapshot := range vm.Snapshots.Snapshot {
		SnapshotCount++
		Created, _ := time.Parse("RFC3339", snapshot.Created)
		if time.Now().Sub(Created).Hours() > (float64(MaxSnapshotAge) * 24) {
			OldSnapshots++
		}
	}
	SnapshotString := fmt.Sprintf("%d", OldSnapshots)

	if OldSnapshots > 0 {
		return SnapshotString, fmt.Errorf("%d / %d Snapshots older than %d days", OldSnapshots, SnapshotCount, MaxSnapshotAge)
	}

	return SnapshotString, nil
}
