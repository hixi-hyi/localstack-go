package localstack

import "fmt"

type Service struct {
	Package string
	Name    string
	Port    int64
}
type ServiceList []*Service

var APIGateway = &Service{
	Package: "apigateway",
	Name:    "APIGateway",
	Port:    4567,
}
var Kinesis = &Service{
	Package: "kinesis",
	Name:    "Kinesis",
	Port:    4568,
}
var DynamoDB = &Service{
	Package: "dynamodb",
	Name:    "DynamoDB",
	Port:    4569,
}
var DynamoDBStreams = &Service{
	Package: "dynamodbstreams",
	Name:    "DynamoDBStreams",
	Port:    4570,
}
var S3 = &Service{
	Package: "s3",
	Name:    "S3",
	Port:    4572,
}
var Firehose = &Service{
	Package: "firehose",
	Name:    "Firehose",
	Port:    4573,
}
var Lambda = &Service{
	Package: "lambda",
	Name:    "Lambda",
	Port:    4574,
}
var SNS = &Service{
	Package: "sns",
	Name:    "SNS",
	Port:    4575,
}
var SQS = &Service{
	Package: "sqs",
	Name:    "SQS",
	Port:    4576,
}
var Redshift = &Service{
	Package: "redshift",
	Name:    "Redshift",
	Port:    4577,
}
var ElasticsearchService = &Service{
	Package: "elasticsearchservice",
	Name:    "ElasticsearchService",
	Port:    4578,
}
var SES = &Service{
	Package: "ses",
	Name:    "SES",
	Port:    4579,
}
var Route53 = &Service{
	Package: "route53",
	Name:    "Route53",
	Port:    4580,
}
var CloudFormation = &Service{
	Package: "cloudformation",
	Name:    "CloudFormation",
	Port:    4581,
}
var CloudWatch = &Service{
	Package: "cloudwatch",
	Name:    "CloudWatch",
	Port:    4582,
}
var SMS = &Service{
	Package: "sms",
	Name:    "SMS",
	Port:    4583,
}
var SecretsManager = &Service{
	Package: "secretsmanager",
	Name:    "SecretsManager",
	Port:    4584,
}
var SFN = &Service{
	Package: "sfn",
	Name:    "SFN",
	Port:    4585,
}
var CloudWatchLogs = &Service{
	Package: "cloudwatchlogs",
	Name:    "CloudWatchLogs",
	Port:    4586,
}
var STS = &Service{
	Package: "sts",
	Name:    "STS",
	Port:    4592,
}
var IAM = &Service{
	Package: "iam",
	Name:    "IAM",
	Port:    4593,
}

var Services = ServiceList{
	APIGateway,
	Kinesis,
	DynamoDB,
	DynamoDBStreams,
	S3,
	Firehose,
	Lambda,
	SNS,
	SQS,
	Redshift,
	ElasticsearchService,
	SES,
	Route53,
	CloudFormation,
	CloudWatch,
	SMS,
	SecretsManager,
	SFN,
	CloudWatchLogs,
	STS,
	IAM,
}

func (sl ServiceList) Get(serviceName string) (*Service, error) {
	for _, s := range Services {
		if s.Name == serviceName || s.Package == serviceName {
			return s, nil
		}
	}
	return nil, fmt.Errorf("Service not found: %s", serviceName)
}
