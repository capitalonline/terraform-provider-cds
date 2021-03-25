# terraform-provider-cds

Terraform CDS provider

## Building

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.
```sh
# make sure you have installed the terraform.
## if not, please refer to this website.
## https://learn.hashicorp.com/terraform/getting-started/install.html

# Download project source code, move it to $GOPATH/src

$ cd terraform-provider-cds

# if this command fails to execute, please config GOPROXY, refer to : https://goproxy.io/
$ go get

# complie
$ go build -o terraform-provider-cds

$ mkdir ~/.terraform.d/plugins/

$ mv terraform-provider-cds ~/.terraform.d/plugins/

#this is for linux, if you use other system, please move the compiled file to right path.
```
If you already have the compiled file, just put it to your terraform.d/plugins path.

## Building with Version

```sh
# make sure you have installed the terraform.
# Download project source code

$ cd terraform-provider-cds

# Complie for Linux,
$ make linux-with-version

# For Mac
$ make mac-with-version

# For Windows 
$ make windows-with-version

```

## Using
All this resource will happen in our test account so that this is the fastest way to let you test the terrform, and need the Access Key ID, Access Secret Key.
You can test it directly.


### Features
we provide the examples of .tf file.
#### Preparation
```bash
# Configure ak, sk
# and select the region id where the resource will be created in sn.sh .
$ source sn.sh
```
#### VM Create
```bash
$ cd examples/cds-instance/
$ terraform init
$ terraform plan
$ terraform apply
```
#### VM Destroy
```sh
$ terraform destroy
```

**Note:** Acceptance tests create real resources in CDS which often cost money to run.

### Appendix
region_id, public_network.type, image_id, instance_type, refer to [CDS OpenApi](https://github.com/capitalonline/openapi/blob/master/%E9%A6%96%E4%BA%91OpenAPI(v1.2).md#6describesecuritygroups)


## Upgrade
| terraform version | terraform-provider-cds tag | need upgrade |
| :----: | :----: | :----: |
| 0.13.* or higher  | *                          | True         |
| 0.12.* or lower   | *                          | False        |

You have to refer to it https://www.terraform.io/upgrade-guides/index.html. And Step by step upgrade to the corresponding version

The easiest way to upgrade:
1. modify versions.tf file
```
terraform {
    required_providers {
        cds = {
            source = "terraform.capitalonline.net/capitalonline/cds"
            version = "1.0.0"
        }
    }
}
```
2. change compiled file path
- Implied Local Mirror Directories
  * Windows: %APPDATA%/terraform.d/plugins and %APPDATA%/HashiCorp/Terraform/plugins
  * Mac OS X: $HOME/.terraform.d/plugins/, ~/Library/Application Support/io.terraform/plugins, and /Library/Application Support/io.terraform/plugins
  * Linux and other Unix-like systems: $HOME/.terraform.d/plugins/, and XDG Base Directory data directories as configured, after appending terraform/plugins. Without any XDG environment variables set, Terraform will use ~/.local/share/terraform/plugins, /usr/local/share/terraform/plugins, and /usr/share/terraform/plugins.
- Flatform
  * Windows: windows_amd64
  * Mac OS X: darwin_amd64
  * Linux and other Unix-like systems: linux_amd64
- File path format
  * [local mirror dir]/[hostname]/[namespace]/[type]/[tag version]/[platform]/[provider]
```bash
$ go build -o ~/.terraform.d/plugins/terraform.capitalonline.net/capitalonline/cds/1.0.0/darwin_amd64/terraform-provider-cds
$ cd example/cds_instance
$ tree -a
.
├── main.tf
├── variables.tf
└── versions.tf
$ terraform init
$ tree -a
.
├── .terraform
│   └── providers
│       └── terraform.capitalonline.net
│           └── capitalonline
│               └── cds
│                   └── 1.0.0
│                       └── darwin_amd64 -> /Users/XXX/.terraform.d/plugins/terraform.capitalonline.net/capitalonline/cds/1.0.0/darwin_amd64
├── .terraform.lock.hcl
├── main.tf
├── variables.tf
└── versions.tf
```
now you can continue your option.