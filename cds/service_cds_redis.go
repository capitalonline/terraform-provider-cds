package cds

import (
	"context"
	"fmt"
	"log"
	"terraform-provider-cds/cds/connectivity"

	"github.com/capitalonline/cds-gic-sdk-go/redis"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
)

type RedisService struct {
	client *connectivity.CdsClient
}

//create Redis
func (me *RedisService) CreateRedis(ctx context.Context, request *redis.CreateDBInstanceRequest) (*redis.CreateDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().CreateDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}

//update Redis
func (me *RedisService) UpdateRedis(ctx context.Context) {

}

//describeDBInstance
func (me *RedisService) DescribeRedis(ctx context.Context, request *redis.DescribeDBInstancesRequest) (*redis.DescribeDBInstancesResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DescribeDBInstances(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}

//delete Redis
func (me *RedisService) DeleteRedis(ctx context.Context, request *redis.DeleteDBInstanceRequest) (*redis.DeleteDBInstanceResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisClient().DeleteDBInstance(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}

//describe regions
func (me *RedisService) DescribeRegions(ctx context.Context, request *redis.DescribeRegionsRequest) (*redis.DescribeRegionsResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())
	response, err := me.client.UseRedisGetClient().DescribeRegins(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}

//describe available db config
func (me *RedisService) DescribeAvailableDBConfig(ctx context.Context, request *redis.DescribeAvailableDBConfigRequest) (*redis.DescribeAvailableDBConfigResponse, error) {
	logId := getLogId(ctx)
	ratelimit.Check(request.GetAction())

	response, err := me.client.UseRedisGetClient().DescribeAvailableDBConfig(request)
	log.Println(fmt.Sprintf("[DEBUG]%s api[%s] , request body [%s], response body [%s]", logId, request.GetAction(), request.ToJsonString(), response.ToJsonString()))

	return response, err
}
