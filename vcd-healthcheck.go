package main

import (
    "bufio"
    "fmt"
    "net/url"
    "os"
    "strings"

    "github.com/vmware/govcloudair"
    "github.com/howeyc/gopass"
    "github.com/olekukonko/tablewriter"
    "github.com/UKCloud/vcd-healthcheck/healthcheck"

    types "github.com/vmware/govcloudair/types/v56"
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
func (c *Config) Client() (*govcloudair.VCDClient, error) {
    u, err := url.ParseRequestURI(c.Href)
    if err != nil {
        return nil, fmt.Errorf("Unable to pass url: %s", err)
    }

    vcdclient := govcloudair.NewVCDClient(*u, c.Insecure)
    org, vcd, err := vcdclient.Authenticate(c.User, c.Password, c.Org, c.VDC)
    if err != nil {
        return nil, fmt.Errorf("Unable to authenticate: %s", err)
    }
    vcdclient.Org = org
    vcdclient.OrgVdc = vcd
    return vcdclient, nil
}

// CheckVM is called for each search result 
func CheckVM(client *govcloudair.VCDClient, s types.QueryResultVMRecordType) ([]string, error) {
  if s.VAppTemplate == true {
    return nil, nil
  }

  ReturnRow := false
  // red := color.New(color.FgRed).SprintFunc()

  VM, err := client.FindVMByHREF(s.HREF)
  if err != nil {
      return nil, fmt.Errorf("Unable to load VM: %s", err)
  }

  HWVersion, err := healthcheck.HardwareVersion(s, VM.VM)
  if err != nil {
    ReturnRow = true
  }

  NetworkDevice, err := healthcheck.NetworkDevice(s, VM.VM)
  if err != nil {
    ReturnRow = true
  }

  SnapshotString, err := healthcheck.VMSnapshots(s, VM.VM)
  if err != nil {
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

