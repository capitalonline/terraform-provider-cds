# link the new provider repo
pushd "$GOPATH/src/github.com/hashicorp/terraform-website/ext/providers"
ln -sf "$GOPATH/src/terraform-provider-cds" "cds"
popd

# link the layout file
pushd "$GOPATH/src/github.com/hashicorp/terraform-website/content/source/layouts"
ln -sf "../../../ext/providers/cds/website/cds.erb" "cds.erb"
popd

# link the content
pushd "$GOPATH/src/github.com/hashicorp/terraform-website/content/source/docs/providers"
ln -sf "../../../../ext/providers/cds/website/docs" "cds"
popd

## start middleman
#cd "$GOPATH/src/github.com/terraform-provider-cds"
#make website