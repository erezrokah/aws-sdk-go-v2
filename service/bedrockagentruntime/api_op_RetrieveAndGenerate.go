// Code generated by smithy-go-codegen DO NOT EDIT.

package bedrockagentruntime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/bedrockagentruntime/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Queries a knowledge base and generates responses based on the retrieved
// results. The response only cites sources that are relevant to the query.
func (c *Client) RetrieveAndGenerate(ctx context.Context, params *RetrieveAndGenerateInput, optFns ...func(*Options)) (*RetrieveAndGenerateOutput, error) {
	if params == nil {
		params = &RetrieveAndGenerateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RetrieveAndGenerate", params, optFns, c.addOperationRetrieveAndGenerateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RetrieveAndGenerateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RetrieveAndGenerateInput struct {

	// Contains the query to be made to the knowledge base.
	//
	// This member is required.
	Input *types.RetrieveAndGenerateInput

	// Contains configurations for the knowledge base query and retrieval process. For
	// more information, see [Query configurations].
	//
	// [Query configurations]: https://docs.aws.amazon.com/bedrock/latest/userguide/kb-test-config.html
	RetrieveAndGenerateConfiguration *types.RetrieveAndGenerateConfiguration

	// Contains details about the session with the knowledge base.
	SessionConfiguration *types.RetrieveAndGenerateSessionConfiguration

	// The unique identifier of the session. Reuse the same value to continue the same
	// session with the knowledge base.
	SessionId *string

	noSmithyDocumentSerde
}

type RetrieveAndGenerateOutput struct {

	// Contains the response generated from querying the knowledge base.
	//
	// This member is required.
	Output *types.RetrieveAndGenerateOutput

	// The unique identifier of the session. Reuse the same value to continue the same
	// session with the knowledge base.
	//
	// This member is required.
	SessionId *string

	// A list of segments of the generated response that are based on sources in the
	// knowledge base, alongside information about the sources.
	Citations []types.Citation

	// Specifies if there is a guardrail intervention in the response.
	GuardrailAction types.GuadrailAction

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRetrieveAndGenerateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsRestjson1_serializeOpRetrieveAndGenerate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpRetrieveAndGenerate{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "RetrieveAndGenerate"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = addClientRequestID(stack); err != nil {
		return err
	}
	if err = addComputeContentLength(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addComputePayloadSHA256(stack); err != nil {
		return err
	}
	if err = addRetry(stack, options); err != nil {
		return err
	}
	if err = addRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = addRecordResponseTiming(stack); err != nil {
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
	if err = addTimeOffsetBuild(stack, c); err != nil {
		return err
	}
	if err = addUserAgentRetryMode(stack, options); err != nil {
		return err
	}
	if err = addOpRetrieveAndGenerateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRetrieveAndGenerate(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addRecursionDetection(stack); err != nil {
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

func newServiceMetadataMiddleware_opRetrieveAndGenerate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "RetrieveAndGenerate",
	}
}
