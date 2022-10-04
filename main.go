package main

import(
	"os"
	"fmt"
//"strings"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"log"
	//"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	
)

type Order struct{
	OrderId string `dynamodbav:"orederId" json:"orderId"`
	CountryId string `json:"countryId"`
	CustName string `json:"custName"`
	OrderNumber string `json:"orderNumber"`
	OrderStatus string `json:"orderStatus"`
	OrderType string `json:"orderType"`
}
type Orders struct{
	OrderId string `json:"orderId"`
	CountryId string `json:"countryId"`
}

var(
	dynaClient dynamodbiface.DynamoDBAPI
)

func main()  {
	region := os.Getenv("AWS_REGION")
	awsSession, err := session.NewSession(
		&aws.Config{
			Region : aws.String(region)},)
	if err != nil{
		return
	}
	dynaClient = dynamodb.New(awsSession)
	
	lambda.Start(handlers)
}
const tableName = "Order"
func handlers()  {
		params := &dynamodb.QueryInput{
			TableName:              aws.String(tableName),
			IndexName:              aws.String("countryId-orderId-index"),
			KeyConditionExpression: aws.String("countryId = :countryId"),
			ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
				":countryId": {
					S: aws.String("IND"),
				},
				/*":orderId":{
					S:aws.String("4"),
				},*/
			},
			ScanIndexForward: aws.Bool(false), // sort by sort key in ASCENDING order
			Limit:            aws.Int64(3),
		}
		
	
		result, err := dynaClient.Query(params)
		if err!= nil{
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeProvisionedThroughputExceededException:
					fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
				case dynamodb.ErrCodeResourceNotFoundException:
					fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
				case dynamodb.ErrCodeInternalServerError:
					fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
				default:
					fmt.Println(aerr.Error())
				}
			} else {
				// Print the error, cast err to awserr.Error to get the Code and
				// Message from an error.
				fmt.Println(err.Error())
			}
			//return nil, errors.New(ErrorFailedToUnmarshalRecord)
		}
		fmt.Println(result)
	    item := Orders{}
		err = dynamodbattribute.UnmarshalMap(result.LastEvaluatedKey, &item)
		if err!= nil{
			fmt.Println(" Not Able to unmarshal objects")
		}
		fmt.Println("value..",string(item.OrderId))

		for len(result.LastEvaluatedKey) != 00 {
			var temp string = string(item.OrderId)
			//var tt string = "3"
			log.Print(temp)
			fmt.Println("----------------------------")
			params1 := &dynamodb.QueryInput{
				TableName:              aws.String(tableName),
				IndexName:              aws.String("countryId-orderId-index"),
				KeyConditionExpression: aws.String("countryId = :countryId AND orderId < :orderId"),
				ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
					":countryId": {
						S: aws.String("IND"),
					},
					":orderId":{
						S: aws.String(temp),
					},
				},
				ScanIndexForward: aws.Bool(false), // sort by sort key in ASCENDING order
				Limit:            aws.Int64(3),
			}
			result1, err := dynaClient.Query(params1)
			if err!= nil{
				if aerr, ok := err.(awserr.Error); ok {
					switch aerr.Code() {
					case dynamodb.ErrCodeProvisionedThroughputExceededException:
						fmt.Println(dynamodb.ErrCodeProvisionedThroughputExceededException, aerr.Error())
					case dynamodb.ErrCodeResourceNotFoundException:
						fmt.Println(dynamodb.ErrCodeResourceNotFoundException, aerr.Error())
					case dynamodb.ErrCodeInternalServerError:
						fmt.Println(dynamodb.ErrCodeInternalServerError, aerr.Error())
					default:
						fmt.Println(aerr.Error())
					}
				} else {
					// Print the error, cast err to awserr.Error to get the Code and
					// Message from an error.
					fmt.Println(err.Error())
				}
				//return nil, errors.New(ErrorFailedToUnmarshalRecord)
			}
			fmt.Println(result1)
			if len(result.LastEvaluatedKey) != 00{
				break;
			}
		}
		
}