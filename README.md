# localstack-go

## Synopips
```
package main
import (
    "github.com/hixi-hyi/localstack-go/localstack"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

var AWSSession = session.Must(session.NewSession())
var awsSns = sns.New(AWSSession)

func init() {
    l := localstack.NewLocalStack()
    # l := localstack.New(&localstack.Config{ Domain: "localstack" })
    awsSns = l.SNS(AWSSession)
}
```
