// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Synchronizes a job to the source control repository. This operation takes the
// job artifacts from the Glue internal stores and makes a commit to the remote
// repository that is configured on the job. This API supports optional parameters
// which take in the repository information.
func (c *Client) UpdateSourceControlFromJob(ctx context.Context, params *UpdateSourceControlFromJobInput, optFns ...func(*Options)) (*UpdateSourceControlFromJobOutput, error) {
	if params == nil {
		params = &UpdateSourceControlFromJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateSourceControlFromJob", params, optFns, c.addOperationUpdateSourceControlFromJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateSourceControlFromJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateSourceControlFromJobInput struct {

	// The type of authentication, which can be an authentication token stored in
	// Amazon Web Services Secrets Manager, or a personal access token.
	AuthStrategy types.SourceControlAuthStrategy

	// The value of the authorization token.
	AuthToken *string

	// An optional branch in the remote repository.
	BranchName *string

	// A commit ID for a commit in the remote repository.
	CommitId *string

	// An optional folder in the remote repository.
	Folder *string

	// The name of the Glue job to be synchronized to or from the remote repository.
	JobName *string

	// The provider for the remote repository.
	Provider types.SourceControlProvider

	// The name of the remote repository that contains the job artifacts.
	RepositoryName *string

	// The owner of the remote repository that contains the job artifacts.
	RepositoryOwner *string

	noSmithyDocumentSerde
}

type UpdateSourceControlFromJobOutput struct {

	// The name of the Glue job.
	JobName *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateSourceControlFromJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateSourceControlFromJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateSourceControlFromJob{}, middleware.After)
	if err != nil {
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
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateSourceControlFromJob(options.Region), middleware.Before); err != nil {
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
	return nil
}

func newServiceMetadataMiddleware_opUpdateSourceControlFromJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "UpdateSourceControlFromJob",
	}
}