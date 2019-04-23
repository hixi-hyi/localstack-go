package localstack

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/apigateway"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodbstreams"
	"github.com/aws/aws-sdk-go/service/elasticsearchservice"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/iam"
	"github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/redshift"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/aws/aws-sdk-go/service/sfn"
	"github.com/aws/aws-sdk-go/service/sms"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sts"
)

type Config struct {
	Domain string
}
type LocalStack struct {
	Config *Config
}

func NewLocalStack() *LocalStack {
	cfg := &Config{
		Domain: "localstack",
	}
	return New(cfg)
}

func New(cfg *Config) *LocalStack {
	return &LocalStack{Config: cfg}
}

func (l *LocalStack) Session(s *session.Session, port int64) client.ConfigProvider {
	c := &aws.Config{
		Endpoint: aws.String(fmt.Sprintf("http://%s:%d", l.Config.Domain, port)),
	}
	return s.Copy(c)
}

func (l *LocalStack) APIGateway(s *session.Session) *apigateway.APIGateway {
	return apigateway.New(l.Session(s, 4567))
}

func (l *LocalStack) Kinesis(s *session.Session) *kinesis.Kinesis {
	return kinesis.New(l.Session(s, 4568))
}

func (l *LocalStack) DynamoDB(s *session.Session) *dynamodb.DynamoDB {
	return dynamodb.New(l.Session(s, 4569))
}

func (l *LocalStack) DynamoDBStreams(s *session.Session) *dynamodbstreams.DynamoDBStreams {
	return dynamodbstreams.New(l.Session(s, 4570))
}

func (l *LocalStack) S3(s *session.Session) *s3.S3 {
	return s3.New(l.Session(s, 4572))
}

func (l *LocalStack) Firehose(s *session.Session) *firehose.Firehose {
	return firehose.New(l.Session(s, 4573))
}

func (l *LocalStack) Lambda(s *session.Session) *lambda.Lambda {
	return lambda.New(l.Session(s, 4574))
}

func (l *LocalStack) SNS(s *session.Session) *sns.SNS {
	return sns.New(l.Session(s, 4575))
}

func (l *LocalStack) SQS(s *session.Session) *sqs.SQS {
	return sqs.New(l.Session(s, 4576))
}

func (l *LocalStack) Redshift(s *session.Session) *redshift.Redshift {
	return redshift.New(l.Session(s, 4577))
}

func (l *LocalStack) ElasticsearchService(s *session.Session) *elasticsearchservice.ElasticsearchService {
	return elasticsearchservice.New(l.Session(s, 4578))
}
func (l *LocalStack) SES(s *session.Session) *ses.SES {
	return ses.New(l.Session(s, 4579))
}

func (l *LocalStack) Route53(s *session.Session) *route53.Route53 {
	return route53.New(l.Session(s, 4580))
}

func (l *LocalStack) CloudFormation(s *session.Session) *cloudformation.CloudFormation {
	return cloudformation.New(l.Session(s, 4581))
}

func (l *LocalStack) CloudWatch(s *session.Session) *cloudwatch.CloudWatch {
	return cloudwatch.New(l.Session(s, 4582))
}

func (l *LocalStack) SMS(s *session.Session) *sms.SMS {
	return sms.New(l.Session(s, 4583))
}

func (l *LocalStack) SecretsManager(s *session.Session) *secretsmanager.SecretsManager {
	return secretsmanager.New(l.Session(s, 4584))
}

func (l *LocalStack) SFN(s *session.Session) *sfn.SFN {
	return sfn.New(l.Session(s, 4585))
}

func (l *LocalStack) CloudWatchLogs(s *session.Session) *cloudwatchlogs.CloudWatchLogs {
	return cloudwatchlogs.New(l.Session(s, 4586))
}

func (l *LocalStack) STS(s *session.Session) *sts.STS {
	return sts.New(l.Session(s, 4592))
}

func (l *LocalStack) IAM(s *session.Session) *iam.IAM {
	return iam.New(l.Session(s, 4593))
}
