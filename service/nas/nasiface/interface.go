// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package nasiface provides an interface to enable mocking the NIFCLOUD NAS service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package nasiface

import (
	"github.com/kzmake/nifcloud-sdk-go/nifcloud"
	"github.com/kzmake/nifcloud-sdk-go/nifcloud/request"
	"github.com/kzmake/nifcloud-sdk-go/service/nas"
)

// NasAPI provides an interface to enable mocking the
// nas.Nas service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // NIFCLOUD NAS.
//    func myFunc(svc nasiface.NasAPI) bool {
//        // Make svc.AuthorizeNASSecurityGroupIngress request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := nas.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockNasClient struct {
//        nasiface.NasAPI
//    }
//    func (m *mockNasClient) AuthorizeNASSecurityGroupIngress(input *nas.AuthorizeNASSecurityGroupIngressInput) (*nas.AuthorizeNASSecurityGroupIngressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockNasClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type NasAPI interface {
	AuthorizeNASSecurityGroupIngress(*nas.AuthorizeNASSecurityGroupIngressInput) (*nas.AuthorizeNASSecurityGroupIngressOutput, error)
	AuthorizeNASSecurityGroupIngressWithContext(nifcloud.Context, *nas.AuthorizeNASSecurityGroupIngressInput, ...request.Option) (*nas.AuthorizeNASSecurityGroupIngressOutput, error)
	AuthorizeNASSecurityGroupIngressRequest(*nas.AuthorizeNASSecurityGroupIngressInput) (*request.Request, *nas.AuthorizeNASSecurityGroupIngressOutput)

	CreateNASInstance(*nas.CreateNASInstanceInput) (*nas.CreateNASInstanceOutput, error)
	CreateNASInstanceWithContext(nifcloud.Context, *nas.CreateNASInstanceInput, ...request.Option) (*nas.CreateNASInstanceOutput, error)
	CreateNASInstanceRequest(*nas.CreateNASInstanceInput) (*request.Request, *nas.CreateNASInstanceOutput)

	CreateNASSecurityGroup(*nas.CreateNASSecurityGroupInput) (*nas.CreateNASSecurityGroupOutput, error)
	CreateNASSecurityGroupWithContext(nifcloud.Context, *nas.CreateNASSecurityGroupInput, ...request.Option) (*nas.CreateNASSecurityGroupOutput, error)
	CreateNASSecurityGroupRequest(*nas.CreateNASSecurityGroupInput) (*request.Request, *nas.CreateNASSecurityGroupOutput)

	DeleteNASInstance(*nas.DeleteNASInstanceInput) (*nas.DeleteNASInstanceOutput, error)
	DeleteNASInstanceWithContext(nifcloud.Context, *nas.DeleteNASInstanceInput, ...request.Option) (*nas.DeleteNASInstanceOutput, error)
	DeleteNASInstanceRequest(*nas.DeleteNASInstanceInput) (*request.Request, *nas.DeleteNASInstanceOutput)

	DeleteNASSecurityGroup(*nas.DeleteNASSecurityGroupInput) (*nas.DeleteNASSecurityGroupOutput, error)
	DeleteNASSecurityGroupWithContext(nifcloud.Context, *nas.DeleteNASSecurityGroupInput, ...request.Option) (*nas.DeleteNASSecurityGroupOutput, error)
	DeleteNASSecurityGroupRequest(*nas.DeleteNASSecurityGroupInput) (*request.Request, *nas.DeleteNASSecurityGroupOutput)

	DescribeNASInstances(*nas.DescribeNASInstancesInput) (*nas.DescribeNASInstancesOutput, error)
	DescribeNASInstancesWithContext(nifcloud.Context, *nas.DescribeNASInstancesInput, ...request.Option) (*nas.DescribeNASInstancesOutput, error)
	DescribeNASInstancesRequest(*nas.DescribeNASInstancesInput) (*request.Request, *nas.DescribeNASInstancesOutput)

	DescribeNASSecurityGroups(*nas.DescribeNASSecurityGroupsInput) (*nas.DescribeNASSecurityGroupsOutput, error)
	DescribeNASSecurityGroupsWithContext(nifcloud.Context, *nas.DescribeNASSecurityGroupsInput, ...request.Option) (*nas.DescribeNASSecurityGroupsOutput, error)
	DescribeNASSecurityGroupsRequest(*nas.DescribeNASSecurityGroupsInput) (*request.Request, *nas.DescribeNASSecurityGroupsOutput)

	GetMetricStatistics(*nas.GetMetricStatisticsInput) (*nas.GetMetricStatisticsOutput, error)
	GetMetricStatisticsWithContext(nifcloud.Context, *nas.GetMetricStatisticsInput, ...request.Option) (*nas.GetMetricStatisticsOutput, error)
	GetMetricStatisticsRequest(*nas.GetMetricStatisticsInput) (*request.Request, *nas.GetMetricStatisticsOutput)

	ModifyNASInstance(*nas.ModifyNASInstanceInput) (*nas.ModifyNASInstanceOutput, error)
	ModifyNASInstanceWithContext(nifcloud.Context, *nas.ModifyNASInstanceInput, ...request.Option) (*nas.ModifyNASInstanceOutput, error)
	ModifyNASInstanceRequest(*nas.ModifyNASInstanceInput) (*request.Request, *nas.ModifyNASInstanceOutput)

	ModifyNASSecurityGroup(*nas.ModifyNASSecurityGroupInput) (*nas.ModifyNASSecurityGroupOutput, error)
	ModifyNASSecurityGroupWithContext(nifcloud.Context, *nas.ModifyNASSecurityGroupInput, ...request.Option) (*nas.ModifyNASSecurityGroupOutput, error)
	ModifyNASSecurityGroupRequest(*nas.ModifyNASSecurityGroupInput) (*request.Request, *nas.ModifyNASSecurityGroupOutput)

	RevokeNASSecurityGroupIngress(*nas.RevokeNASSecurityGroupIngressInput) (*nas.RevokeNASSecurityGroupIngressOutput, error)
	RevokeNASSecurityGroupIngressWithContext(nifcloud.Context, *nas.RevokeNASSecurityGroupIngressInput, ...request.Option) (*nas.RevokeNASSecurityGroupIngressOutput, error)
	RevokeNASSecurityGroupIngressRequest(*nas.RevokeNASSecurityGroupIngressInput) (*request.Request, *nas.RevokeNASSecurityGroupIngressOutput)
}

var _ NasAPI = (*nas.Nas)(nil)
