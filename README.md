# ktcloudvpc-sdk-for-drv

This repository is Go SDK for KT Cloud VPC connection driver.

## How to install

Reference ktcloudvpc-sdk-for-drv package in your code:

```Go
import "github.com/cloud-barista/ktcloudvpc-sdk-for-drv"
```

Then update your `go.mod`:

```shell
go mod tidy
```

## Modified from original code (Gophercloud)
The followings are modified from original code.
- Connection client codes
  - Modified : v3auth(), NewIdentityV3(), initClientOpts(), NewImageServiceV2()  
- Compute v2 > Image list request code
  - Added : List()
- Compute v2 > Image URL code
  - Added : listURL()
- CommonFunction
  - Added : InitLog()


## Original source code : Gophercloud
Gophercloud is an OpenStack Go SDK.
[https://github.com/gophercloud/gophercloud](https://github.com/gophercloud/gophercloud)

Gophercloud is licensed under the Apache License, Version 2.0