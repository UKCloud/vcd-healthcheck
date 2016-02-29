package main

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "strings"
    "time"

    "github.com/hmrc/vmware-govcd"
    "github.com/howeyc/gopass"
    "github.com/olekukonko/tablewriter"
    // "github.com/fatih/color"
    types "github.com/hmrc/vmware-govcd/types/v56"
)

// VERSION is set at build time by using the following: 
// go build -ldflags "-X main.VERSION=$(git describe --tags)"
var VERSION string

// Config details for connecting to vCloud Director
type Config struct {
    User     string
    Password string
    Org      string
    Href     string
    VDC      string
    Insecure bool
}

// Client connection using the govcd library
func (c *Config) Client() (*govcd.VCDClient, error) {
    u, err := url.ParseRequestURI(c.Href)
    if err != nil {
        return nil, fmt.Errorf("Unable to pass url: %s", err)
    }

    vcdclient := govcd.NewVCDClient(*u, c.Insecure)
    org, vcd, err := vcdclient.Authenticate(c.User, c.Password, c.Org, c.VDC)
    if err != nil {
        return nil, fmt.Errorf("Unable to authenticate: %s", err)
    }
    vcdclient.Org = org
    vcdclient.OrgVdc = vcd
    return vcdclient, nil
}

// CheckVM is called for each search result 
func CheckVM(client *govcd.VCDClient, s types.QueryResultVMRecordType) ([]string, error) {
  if s.VAppTemplate == true {
    return nil, nil
  }

  ReturnRow := false
  // red := color.New(color.FgRed).SprintFunc()

  VM, err := client.FindVMByHREF(s.HREF)
  if err != nil {
      return nil, fmt.Errorf("Unable to load VM: %s", err)
  }

  HWVersion := fmt.Sprintf("%d", s.HardwareVersion)
  if s.HardwareVersion != 9 { 
    // HWVersion = red(HWVersion)
    ReturnRow = true
  }

  NetworkDevice := "Unknown"
  for _,v := range VM.VM.VirtualHardwareSection.Item {
    if v.ResourceType == 10 {
      NetworkDevice = v.ResourceSubType
    }
  }
  if NetworkDevice != "VMXNET3" {
    // NetworkDevice = red(NetworkDevice)
    ReturnRow = true
  }

  SnapshotCount := 0
  OldSnapshots := 0
  // CurrentTime := time.now()
  for _, snapshot := range VM.VM.Snapshots.Snapshot {
    SnapshotCount++
    Created, _ :=time.Parse("RFC3339", snapshot.Created)
    if time.Now().Sub(Created).Hours() > (7 * 24) {
      OldSnapshots++
    }
  }
  SnapshotString := fmt.Sprintf("%d", OldSnapshots)

  if OldSnapshots > 0 {
    // SnapshotString = red(SnapshotString)
    ReturnRow = true
  }

  if ReturnRow == true {
    return []string{s.Name, HWVersion, NetworkDevice, SnapshotString}, nil
  } 

  return nil, nil
}


func main() {

  var User string
  var maskedPassword []byte 
  var Org string

  reader := bufio.NewReader(os.Stdin)
  if os.Getenv("VCLOUD_USERNAME") == "" {
    fmt.Print("Enter your Username: ")
    User, _ = reader.ReadString('\n')
  }

  if os.Getenv("VCLOUD_PASSWORD") == "" {
    fmt.Print("Enter your Password: ")
    maskedPassword, _ = gopass.GetPasswdMasked()
  }

  if os.Getenv("VCLOUD_ORG") == "" {
    fmt.Print("Enter your Organisation ID: ")
    Org, _ = reader.ReadString('\n')
  }

  fmt.Printf("Skyscape Cloud Service vCloud Healthcheck (%s)\n", VERSION)

  config := Config{
        User:     strings.TrimSpace(User),
        Password: strings.TrimSpace(string(maskedPassword)),
        Org:      strings.TrimSpace(Org),
        Href:     "https://api.vcd.portal.skyscapecloud.com/api",
        VDC:      "",
    }

  client, err := config.Client() // We now have a client
  if err != nil {
      fmt.Println(err)
      os.Exit(1)
  }

  results, err := client.Query(map[string]string{"type": "vm"})
  fmt.Printf("Found %d VMs ... processing\n", int(results.Results.Total))

  table := tablewriter.NewWriter(os.Stdout)
  table.SetHeader([]string{"VM", "H/W Version", "Network Device", "Old Snapshots"})
  table.SetBorder(false)

  TableRows := 0
  for _, s := range results.Results.VMRecord {

    row, err := CheckVM(client, *s)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }

    if row != nil {
      table.Append(row)
      TableRows++
    }
  }

  if TableRows > 0 {
    table.Render()
  } else {
    fmt.Printf("No problems found.")
  }
}

