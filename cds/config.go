package cds

import (
	"github.com/capitalonline/cds-gic-sdk-go/common"
	"terraform-provider-cds/cds/connectivity"
)

type Config struct {
	SecretId   string
	SecretKey  string
	Region     string
	commonConn *common.Client
	CdsConn    *connectivity.CdsClient
}

type CdsClient struct {
	apiConn *connectivity.CdsClient
}

func (c *Config) Client() (interface{}, error) {
	var cdsClient CdsClient
	cdsClient.apiConn = connectivity.NewCdsClient(c.SecretId, c.SecretKey, c.Region)
	return &cdsClient, nil
}
