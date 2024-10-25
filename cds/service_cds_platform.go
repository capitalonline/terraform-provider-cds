package cds

import (
	"context"
	"github.com/capitalonline/cds-gic-sdk-go/platform"
	"github.com/terraform-providers/terraform-provider-tencentcloud/tencentcloud/ratelimit"
	"log"
	"math/rand"
	"terraform-provider-cds/cds/connectivity"
	"time"
)

type PlatformService struct {
	client *connectivity.CdsClient
}

func (me *PlatformService) DescribeSubjects(ctx context.Context, request *platform.DescribeSubjectsRequest) (response *platform.DescribeSubjectsResponse, errRet error) {

	logId := getLogId(ctx)
	defer func() {
		if errRet != nil {
			log.Printf("[CRITAL]%s api[%s] fail, request body [%s], reason[%s]\n",
				logId, request.GetAction(), request.ToJsonString(), errRet.Error())
		}
	}()

	ratelimit.Check(request.GetAction())
	// add a random delay to avoid concurrency with Terraform "count" way
	minSleepMs, maxSleepMs := 2000, 10000
	sleepMs := minSleepMs + rand.Intn(maxSleepMs)
	time.Sleep(time.Duration(sleepMs) * time.Millisecond)
	response, err := me.client.UsePlatformGetClient().DescribeSubjects(request)
	if err == nil {
		log.Printf("[DEBUG]%s api[%s] , request body [%s], response body[%s]\n",
			logId, request.GetAction(), request.ToJsonString(), response.ToJsonString())
		return
	}

	errRet = err
	return
}
