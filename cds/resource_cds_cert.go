package cds

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/common"
	"github.com/capitalonline/cds-gic-sdk-go/haproxy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceCdsCert() *schema.Resource {
	return &schema.Resource{
		Create: createResourceCdsCert,
		Read:   readResourceCdsCert,
		Update: updateResourceCdsCert,
		Delete: deleteResourceCdsCert,
		Schema: map[string]*schema.Schema{
			"certificate_id": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate id.",
			},
			"brand": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate brand.",
			},
			"certificate": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate Infomation.",
			},
			"certificate_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate name.",
			},
			"certificate_type": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate type.",
			},
			"domain": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate domain.",
			},
			"start_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate start time.",
			},
			"end_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate end time.",
			},
			"organization": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate Organization.",
			},
			"private_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Certificate private key.",
			},
			"public_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Certificate public key.",
			},
			"valid": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Certificate vaild.",
			},
		},
		Description: "Upload CA certificate for haproxy\n\n" +
			"## Example usage\n\n" +
			"```hcl\n" +
			`
resource cds_certificate my_cds_certificate {
  certificate_name  = "my_cert"
  certificate       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  private_key       = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
}
` +
			"\n```",
	}
}

func readResourceCdsCert(data *schema.ResourceData, meta interface{}) error {
	log.Println("read certificate")
	defer logElapsed("resource.certificate.read")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	certs, err := haproxyService.DescribeCACertificates(ctx, haproxy.NewDescribeCACertificatesRequest())
	if err != nil {
		return err
	}
	if *certs.Code != "Success" {
		return fmt.Errorf("read cert list failed, error: %s", *certs.Message)
	}

	for _, entry := range certs.Data {
		if *entry.CertificateName == data.Get("certificate_name").(string) {
			request := haproxy.NewDescribeCACertificateRequest()

			request.CertificateId = entry.CertificateId
			response, err := haproxyService.DescribeCACertificate(ctx, request)
			if err != nil {
				return err
			}

			if *response.Code != "Success" {
				return fmt.Errorf("read cert error: %s", *response.Message)
			}

			data.SetId(*response.Data.CertificateId)
			data.Set("certificate_id", *response.Data.CertificateId)
			data.Set("brand", response.Data.Brand)
			data.Set("certificate_type", response.Data.CertificateType)
			data.Set("domain", response.Data.Domain)
			data.Set("start_time", response.Data.StartTime)
			data.Set("end_time", response.Data.EndTime)
			data.Set("organization", response.Data.Organization)
			data.Set("private_key", response.Data.PrivateKey)
			data.Set("public_key", response.Data.PublicKey)
			data.Set("valid", response.Data.Valid)
			return nil
		}

	}

	return errors.New("read cert info failed")
}

func createResourceCdsCert(data *schema.ResourceData, meta interface{}) error {
	log.Println("create certificate")
	defer logElapsed("resource.certificate.create")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewUploadCACertificateRequest()

	request.Certificate = common.StringPtr(data.Get("certificate").(string))
	request.PrivateKey = common.StringPtr(data.Get("private_key").(string))
	request.CertificateName = common.StringPtr(data.Get("certificate_name").(string))

	response, err := haproxyService.UploadCACertificate(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("Haproxy modify haproxy strategy with error: %s", err.Error())
	}

	data.SetId(time.Now().String())

	return readResourceCdsCert(data, meta)
}

func updateResourceCdsCert(data *schema.ResourceData, meta interface{}) error {
	return nil
}

func deleteResourceCdsCert(data *schema.ResourceData, meta interface{}) error {
	log.Println("delete certificate")
	defer logElapsed("resource.certificate.delete")()
	logId := getLogId(contextNil)
	ctx := context.WithValue(context.TODO(), "logId", logId)

	haproxyService := HaproxyService{client: meta.(*CdsClient).apiConn}

	request := haproxy.NewDeleteCACertificateRequest()
	request.CertificateId = common.StringPtr(data.Get("certificate_id").(string))

	response, err := haproxyService.DeleteCACertificate(ctx, request)
	if err != nil {
		return err
	}
	if *response.Code != "Success" {
		return fmt.Errorf("delete certificate failed with %s", *response.Message)
	}

	time.Sleep(time.Second * 20)

	return nil
}
