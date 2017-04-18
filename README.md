# vCloud Director Healthcheck

This is a healthcheck script intended to be run against Skyscape's vCloud Director and provide feedback on any mis-configured or non-optimal configurations.

[![Latest Version](http://img.shields.io/github/release/UKCloud/vcd-healthcheck.svg?style=flat-square)](https://github.com/UKCloud/vcd-healthcheck/releases)
[![Build Status](https://travis-ci.org/UKCloud/vcd-healthcheck.svg?branch=master)](https://travis-ci.org/UKCloud/vcd-healthcheck)
[![GoDoc](https://godoc.org/github.com/UKCloud/vcd-healthcheck?status.svg)](https://godoc.org/github.com/UKCloud/vcd-healthcheck)
[![ZenHub](https://raw.githubusercontent.com/ZenHubIO/support/master/zenhub-badge.png)](https://zenhub.com)

## Installation
Download the [latest release](https://github.com/UKCloud/vcd-healthcheck/releases) of the healthcheck from GitHub. Release binaries are provided for you to download for both Windows and Linux. If you require other platforms, you can retrieve the source and compile for yourself.

## Usage
Run the command:
```
vcd-healthcheck
```
You will be prompted to enter your Username, Password and Organisation ID. 

Note: Your Username is not your email address used to login to the UKCloud Portal. You must retrieve the Username and Organisation ID to use from the [UKCloud Portal API Page](https://portal.ukcloud.com/user/api).

Optionally, you can set your user credentials as the following environment variables to prevent being prompted.
```
VCLOUD_USERNAME=1111.1.111111
VCLOUD_PASSWORD=VerySecret
VCLOUD_ORG=1-1-11-111111
```
## About the Checks
The healthcheck script will search for all VMs accessible to the user account you specify. For each VM found by the search, the following checks are performed. If any VMs do not meet the recommendations, its details will be listed. If all of the VMs meet the recommendations, nothing will be output.
* Check that the VM's hardware version is 9.
* Check that the VM's Network Device is VMXNET3.
* Check that there are no VM Snapshots older than 7 days.

License and Authors
-------------------
Authors:
  * Rob Coward (rcoward@ukcloud.com)

Copyright 2016 UKCloud

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

test
