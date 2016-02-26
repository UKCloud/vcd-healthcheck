# vCloud Director Healthcheck
This is a collection of scripts intended to be run again a vCloud Director VDC and provide feedback and recommendations for any mis-configured or non-optimal configurations.

[![Build Status](https://travis-ci.org/skyscape-cloud-services/vcd-healthcheck.svg?branch=master)](https://travis-ci.org/skyscape-cloud-services/vcd-healthcheck)

## Installation
Download the latest release of the healthcheck from GitHub.

## Usage
Run the command:
```
vcd-healthcheck
```
You will be prompted to enter your Username, Password and Organisation ID. 

Note: Your Username is not your email address used to login to the Skyscape Portal. You must retrieve the Username and Organisation ID to use from the [Skyscape Portal API Page](https://portal.skyscapecloud.com/user/api).

Optionally, you can set your user credentials as the following environment variables to prevent being prompted.
```
VCLOUD_USERNAME=1111.1.111111
VCLOUD_PASSWORD=VerySecret
VCLOUD_ORG=1-1-11-111111
```

License and Authors
-------------------
Authors:
  * Rob Coward (rcoward@skyscapecloud.com)

Copyright 2016 Skyscape Cloud Services

Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance with the License. You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.

test
