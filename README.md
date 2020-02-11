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

# This is for linux and MacOS.
## If you use other system, please move the compiled file to right path. Please refer to this website.
## https://www.terraform.io/docs/configuration/providers.html#third-party-plugins
$ mkdir ~/.terraform.d/plugins/

$ mv terraform-provider-cds ~/.terraform.d/plugins/

```
If you already have the compiled file, just put it to your terraform.d/plugins path.

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
