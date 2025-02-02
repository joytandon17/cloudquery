// Code generated by codegen; DO NOT EDIT.
package services

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/eks"
)

//go:generate mockgen -package=mocks -destination=../mocks/eks.go -source=eks.go EksClient
type EksClient interface {
	DescribeAddon(context.Context, *eks.DescribeAddonInput, ...func(*eks.Options)) (*eks.DescribeAddonOutput, error)
	DescribeAddonConfiguration(context.Context, *eks.DescribeAddonConfigurationInput, ...func(*eks.Options)) (*eks.DescribeAddonConfigurationOutput, error)
	DescribeAddonVersions(context.Context, *eks.DescribeAddonVersionsInput, ...func(*eks.Options)) (*eks.DescribeAddonVersionsOutput, error)
	DescribeCluster(context.Context, *eks.DescribeClusterInput, ...func(*eks.Options)) (*eks.DescribeClusterOutput, error)
	DescribeEksAnywhereSubscription(context.Context, *eks.DescribeEksAnywhereSubscriptionInput, ...func(*eks.Options)) (*eks.DescribeEksAnywhereSubscriptionOutput, error)
	DescribeFargateProfile(context.Context, *eks.DescribeFargateProfileInput, ...func(*eks.Options)) (*eks.DescribeFargateProfileOutput, error)
	DescribeIdentityProviderConfig(context.Context, *eks.DescribeIdentityProviderConfigInput, ...func(*eks.Options)) (*eks.DescribeIdentityProviderConfigOutput, error)
	DescribeNodegroup(context.Context, *eks.DescribeNodegroupInput, ...func(*eks.Options)) (*eks.DescribeNodegroupOutput, error)
	DescribeUpdate(context.Context, *eks.DescribeUpdateInput, ...func(*eks.Options)) (*eks.DescribeUpdateOutput, error)
	ListAddons(context.Context, *eks.ListAddonsInput, ...func(*eks.Options)) (*eks.ListAddonsOutput, error)
	ListClusters(context.Context, *eks.ListClustersInput, ...func(*eks.Options)) (*eks.ListClustersOutput, error)
	ListEksAnywhereSubscriptions(context.Context, *eks.ListEksAnywhereSubscriptionsInput, ...func(*eks.Options)) (*eks.ListEksAnywhereSubscriptionsOutput, error)
	ListFargateProfiles(context.Context, *eks.ListFargateProfilesInput, ...func(*eks.Options)) (*eks.ListFargateProfilesOutput, error)
	ListIdentityProviderConfigs(context.Context, *eks.ListIdentityProviderConfigsInput, ...func(*eks.Options)) (*eks.ListIdentityProviderConfigsOutput, error)
	ListNodegroups(context.Context, *eks.ListNodegroupsInput, ...func(*eks.Options)) (*eks.ListNodegroupsOutput, error)
	ListTagsForResource(context.Context, *eks.ListTagsForResourceInput, ...func(*eks.Options)) (*eks.ListTagsForResourceOutput, error)
	ListUpdates(context.Context, *eks.ListUpdatesInput, ...func(*eks.Options)) (*eks.ListUpdatesOutput, error)
}
