package cds

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"terraform-provider-cds/cds/connectivity"
	"time"

	"github.com/capitalonline/cds-gic-sdk-go/mysql"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type MySQLService struct {
	client *connectivity.CdsClient
}

// Create mysql
func (me *MySQLService) CreateMySQL(ctx context.Context, request *mysql.CreateDBInstanceRequest) (*mysql.CreateDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLClient().CreateDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) UpdateMySQL(ctx context.Context, request *mysql.ModifyDBInstanceSpecRequest) (*mysql.ModifyDBInstanceSpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLClient().ModifyDBInstanceSpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) GetMySQLList(ctx context.Context, request *mysql.DescribeDBInstancesRequest) (*mysql.DescribeDBInstancesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeDBInstances(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) DeleteMySQL(ctx context.Context, request *mysql.DeleteDBInstanceRequest) (*mysql.DeleteDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLClient().DeleteDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) CreateReadOnlyMySQL(ctx context.Context, request *mysql.CreateReadOnlyDBInstanceRequest) (*mysql.CreateReadOnlyDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLClient().CreateReadOnlyDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) GetAvailableReadOnlyConfig(ctx context.Context, request *mysql.DescribeAvailableReadOnlyConfigRequest) (*mysql.DescribeAvailableReadOnlyConfigResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeAvailableReadOnlyConfig(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) DescribeRegions(ctx context.Context, request *mysql.DescribeRegionsRequest) (*mysql.DescribeRegionsResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeRegions(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) DescribeAvailableDBConfig(ctx context.Context, request *mysql.DescribeAvailableDBConfigRequest) (*mysql.DescribeAvailableDBConfigResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeAvailableDBConfig(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) DescribeModifiableDBSpec(ctx context.Context, request *mysql.DescribeModifiableDBSpecRequest) (*mysql.DescribeModifiableDBSpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeModifiableDBSpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) ModifyDBInstanceSpec(ctx context.Context, request *mysql.ModifyDBInstanceSpecRequest) (*mysql.ModifyDBInstanceSpecResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLClient().ModifyDBInstanceSpec(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

func (me *MySQLService) DescribeDBInstances(ctx context.Context, request *mysql.DescribeDBInstancesRequest) (*mysql.DescribeDBInstancesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)

	response, err := me.client.UseMySQLGetClient().DescribeDBInstances(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}
