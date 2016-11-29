// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package applicationdiscoveryservice

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol/jsonrpc"
)

// The AWS Application Discovery Service helps Systems Integrators quickly and
// reliably plan application migration projects by automatically identifying
// applications running in on-premises data centers, their associated dependencies,
// and their performance profile.
//
// Planning data center migrations can involve thousands of workloads that are
// often deeply interdependent. Application discovery and dependency mapping
// are important early first steps in the migration process, but difficult to
// perform at scale due to the lack of automated tools.
//
// The AWS Application Discovery Service automatically collects configuration
// and usage data from servers to develop a list of applications, how they perform,
// and how they are interdependent. This information is securely retained in
// an AWS Application Discovery Service database which you can export as a CSV
// file into your preferred visualization tool or cloud migration solution to
// help reduce the complexity and time in planning your cloud migration.
//
// The Application Discovery Service is currently available for preview. Only
// customers who are engaged with AWS Professional Services (https://aws.amazon.com/professional-services/)
// or a certified AWS partner can use the service. To see the list of certified
// partners and request access to the Application Discovery Service, complete
// the following preview form (http://aws.amazon.com/application-discovery/preview/).
//
// This API reference provides descriptions, syntax, and usage examples for
// each of the actions and data types for the Discovery Service. The topic for
// each action shows the API request parameters and the response. Alternatively,
// you can use one of the AWS SDKs to access an API that is tailored to the
// programming language or platform that you're using. For more information,
// see AWS SDKs (http://aws.amazon.com/tools/#SDKs).
//
// This guide is intended for use with the AWS Discovery Service User Guide
// (http://docs.aws.amazon.com/application-discovery/latest/userguide/what-is-appdiscovery.html).
//
// The following are short descriptions of each API action, organized by function.
//
// Managing AWS Agents Using the Application Discovery Service
//
// An AWS agent is software that you install on on-premises servers and virtual
// machines that are targeted for discovery and migration. Agents run on Linux
// and Windows Server and collect server configuration and activity information
// about your applications and infrastructure. Specifically, agents collect
// the following information and send it to the Application Discovery Service
// using Secure Sockets Layer (SSL) encryption:
//
//    * User information (user name, home directory)
//
//    * Group information (name)
//
//    * List of installed packages
//
//    * List of kernel modules
//
//    * All create and stop process events
//
//    * DNS queries
//
//    * NIC information
//
//    * TCP/UDP process listening ports
//
//    * TCPV4/V6 connections
//
//    * Operating system information
//
//    * System performance
//
//    * Process performance
//
// The Application Discovery Service API includes the following actions to manage
// AWS agents:
//
//    * StartDataCollectionByAgentIds: Instructs the specified agents to start
//    collecting data. The Application Discovery Service takes several minutes
//    to receive and process data after you initiate data collection.
//
//    * StopDataCollectionByAgentIds: Instructs the specified agents to stop
//    collecting data.
//
//    * DescribeAgents: Lists AWS agents by ID or lists all agents associated
//    with your user account if you did not specify an agent ID. The output
//    includes agent IDs, IP addresses, media access control (MAC) addresses,
//    agent health, host name where the agent resides, and the version number
//    of each agent.
//
// Querying Configuration Items
//
// A configuration item is an IT asset that was discovered in your data center
// by an AWS agent. When you use the Application Discovery Service, you can
// specify filters and query specific configuration items. The service supports
// Server, Process, and Connection configuration items. This means you can specify
// a value for the following keys and query your IT assets:
//
// Server
//
//    * server.HostName
//
//    * server.osName
//
//    * server.osVersion
//
//    * server.configurationId
//
//    * server.agentId
//
// Process
//
//    * process.name
//
//    * process.CommandLine
//
//    * process.configurationId
//
//    * server.hostName
//
//    * server.osName
//
//    * server.osVersion
//
//    * server.configurationId
//
//    * server.agentId
//
// Connection
//
//    * connection.sourceIp
//
//    * connection.sourcePort
//
//    * connection.destinationIp
//
//    * connection.destinationPort
//
//    * sourceProcess.configurationId
//
//    * sourceProcess.commandLine
//
//    * sourceProcess.name
//
//    * destinationProcessId.configurationId
//
//    * destinationProcess.commandLine
//
//    * destinationProcess.name
//
//    * sourceServer.configurationId
//
//    * sourceServer.hostName
//
//    * sourceServer.osName
//
//    * sourceServer.osVersion
//
//    * destinationServer.configurationId
//
//    * destinationServer.hostName
//
//    * destinationServer.osName
//
//    * destinationServer.osVersion
//
// The Application Discovery Service includes the following actions for querying
// configuration items.
//
//    * DescribeConfigurations: Retrieves a list of attributes for a specific
//    configuration ID. For example, the output for a server configuration item
//    includes a list of attributes about the server, including host name, operating
//    system, number of network cards, etc.
//
//    * ListConfigurations: Retrieves a list of configuration items according
//    to the criteria you specify in a filter. The filter criteria identify
//    relationship requirements. For example, you can specify filter criteria
//    of process.name with values of nginx and apache.
//
// Tagging Discovered Configuration Items
//
// You can tag discovered configuration items. Tags are metadata that help you
// categorize IT assets in your data center. Tags use a key-value format. For
// example, {"key": "serverType", "value": "webServer"}.
//
//    * CreateTags: Creates one or more tags for a configuration items.
//
//    * DescribeTags: Retrieves a list of configuration items that are tagged
//    with a specific tag. Or, retrieves a list of all tags assigned to a specific
//    configuration item.
//
//    * DeleteTags: Deletes the association between a configuration item and
//    one or more tags.
//
// Exporting Data
//
// You can export data as a CSV file to an Amazon S3 bucket or into your preferred
// visualization tool or cloud migration solution to help reduce the complexity
// and time in planning your cloud migration.
//
//    * ExportConfigurations: Exports all discovered configuration data to an
//    Amazon S3 bucket. Data includes tags and tag associations, processes,
//    connections, servers, and system performance. This API returns an export
//    ID which you can query using the GetExportStatus API.
//
//    * DescribeExportConfigurations: Gets the status of the data export. When
//    the export is complete, the service returns an Amazon S3 URL where you
//    can download CSV files that include the data.
//The service client's operations are safe to be used concurrently.
// It is not safe to mutate any of the client's properties though.
type ApplicationDiscoveryService struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// A ServiceName is the name of the service the client will make API calls to.
const ServiceName = "discovery"

// New creates a new instance of the ApplicationDiscoveryService client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     // Create a ApplicationDiscoveryService client from just a session.
//     svc := applicationdiscoveryservice.New(mySession)
//
//     // Create a ApplicationDiscoveryService client with additional configuration
//     svc := applicationdiscoveryservice.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *ApplicationDiscoveryService {
	c := p.ClientConfig(ServiceName, cfgs...)
	return newClient(*c.Config, c.Handlers, c.Endpoint, c.SigningRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, endpoint, signingRegion string) *ApplicationDiscoveryService {
	svc := &ApplicationDiscoveryService{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				SigningRegion: signingRegion,
				Endpoint:      endpoint,
				APIVersion:    "2015-11-01",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSPoseidonService_V2015_11_01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(jsonrpc.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a ApplicationDiscoveryService operation and runs any
// custom request initialization.
func (c *ApplicationDiscoveryService) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
