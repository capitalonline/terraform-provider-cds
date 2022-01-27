package cds

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-cds/cds/connectivity"

	"github.com/capitalonline/cds-gic-sdk-go/mongodb"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type MongodbService struct {
	client *connectivity.CdsClient
}

//create mongodb
func (me *MongodbService) CreateMongodb(ctx context.Context, request *mongodb.CreateDBInstanceRequest) (*mongodb.CreateDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseMongodbClient().CreateDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}

//update mongodb
func (me *MongodbService) UpdateMongodb(ctx context.Context) {

}

//DescribeZones
func (me *MongodbService) DescribeZones(ctx context.Context, request *mongodb.DescribeZonesRequest) (*mongodb.DescribeZonesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseMongodbGetClient().DescribeZones(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

//DescribeSpecInfo
func (me *MongodbService) DescribeSpecInfo(ctx context.Context, request *mongodb.DescribeSpecInfoRequest) (*mongodb.DescribeSpecInfoResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseMongodbGetClient().DescribeSpecInfo(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

//DescribeDBInstances
func (me *MongodbService) DescribeDBInstances(ctx context.Context, request *mongodb.DescribeDBInstancesRequest) (*mongodb.DescribeDBInstancesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseMongodbGetClient().DescribeDBInstances(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))
	return response, err
}

//DeleteDBInstance
func (me *MongodbService) DeleteDBInstance(ctx context.Context, request *mongodb.DeleteDBInstanceRequest) (*mongodb.DeleteDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseMongodbClient().DeleteDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}
