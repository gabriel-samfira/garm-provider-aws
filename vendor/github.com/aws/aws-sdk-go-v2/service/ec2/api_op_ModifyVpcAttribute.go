// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the specified attribute of the specified VPC.
func (c *Client) ModifyVpcAttribute(ctx context.Context, params *ModifyVpcAttributeInput, optFns ...func(*Options)) (*ModifyVpcAttributeOutput, error) {
	if params == nil {
		params = &ModifyVpcAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyVpcAttribute", params, optFns, c.addOperationModifyVpcAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyVpcAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyVpcAttributeInput struct {

	// The ID of the VPC.
	//
	// This member is required.
	VpcId *string

	// Indicates whether the instances launched in the VPC get DNS hostnames. If
	// enabled, instances in the VPC get DNS hostnames; otherwise, they do not. You
	// cannot modify the DNS resolution and DNS hostnames attributes in the same
	// request. Use separate requests for each attribute. You can only enable DNS
	// hostnames if you've enabled DNS support.
	EnableDnsHostnames *types.AttributeBooleanValue

	// Indicates whether the DNS resolution is supported for the VPC. If enabled,
	// queries to the Amazon provided DNS server at the 169.254.169.253 IP address, or
	// the reserved IP address at the base of the VPC network range "plus two" succeed.
	// If disabled, the Amazon provided DNS service in the VPC that resolves public DNS
	// hostnames to IP addresses is not enabled. You cannot modify the DNS resolution
	// and DNS hostnames attributes in the same request. Use separate requests for each
	// attribute.
	EnableDnsSupport *types.AttributeBooleanValue

	// Indicates whether Network Address Usage metrics are enabled for your VPC.
	EnableNetworkAddressUsageMetrics *types.AttributeBooleanValue

	noSmithyDocumentSerde
}

type ModifyVpcAttributeOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyVpcAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyVpcAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyVpcAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyVpcAttribute"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyVpcAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyVpcAttribute(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyVpcAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyVpcAttribute",
	}
}
