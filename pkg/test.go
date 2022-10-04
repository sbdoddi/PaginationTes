
package pkg
/*

import(
	//"golangPagination/pkg/validation"
	//"encoding/json"
	//"errors"
	"fmt"
	//"log"

	"os"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/awserr"
	//"github.com/aws/aws-lambda-go/events"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
    "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
//	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
//	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

/*var(
	ErrorFailedToUnmarshalRecord = "failed to unmarshal records"
	ErrorFailedToFetchRecords = "Failed to fetch records"
)*/
/*
type Pagination struct{
	Email string `dynamodbav:"email" json:"email"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	RollNo int `json:"rollNo"`
}

var(
	dynaClient dynamodbiface.DynamoDBAPI
)
//const tableName = "PaginationTable"
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
const tableName = "PaginationTable"
var limit int = 1;
/*
func GetPaginator() () {
	// Create the Expression to fill the input struct with.
	// Get all movies in that year; we'll pull out those with a higher rating later
	filt := expression.Name("FirstName").Equal(expression.Value("Somashankar"))

	// Or we could get by ratings and pull out those with the right year later
	//    filt := expression.Name("info.rating").GreaterThan(expression.Value(min_rating))

	// Get back the title, year, and rating

	minRollNo := 20
	proj := expression.NamesList(expression.Name("Email"), expression.Name("LastName"), expression.Name("RollNo"))

	expr, err := expression.NewBuilder().WithFilter(filt).WithProjection(proj).Build()
	if err != nil {
		log.Fatalf("Got error building expression: %s", err)
	}
	// Build the query input parameters
	params := &dynamodb.ScanInput{
		ExpressionAttributeNames:  expr.Names(),
		ExpressionAttributeValues: expr.Values(),
		FilterExpression:          expr.Filter(),
		ProjectionExpression:      expr.Projection(),
		TableName:                 aws.String(tableName),
	}

	// Make the DynamoDB Query API call
	result, err := dynaClient.Scan(params)
	if err != nil {
		log.Fatalf("Query API call failed: %s", err)
	}
	numItems := 0

	for _, i := range result.Items {
		item := Pagination{}

		err = dynamodbattribute.UnmarshalMap(i, &item)

		if err != nil {
			log.Fatalf("Got error unmarshalling: %s", err)
		}

		// Which ones had a higher rating than minimum?
		if item.RollNo > minRollNo {
			// Or it we had filtered by rating previously:
			//   if item.Year == year {
			numItems++

			fmt.Println("Email: ", item.Email)
			fmt.Println("RollNo:", item.RollNo)
			fmt.Println()
		}
	}

	fmt.Println("Found", numItems, "movie(s) with a rollNo above", minRollNo, "in")



}


*/
/*
func handlers()(){

	//var firstName string = "Somashankar"

	// Construct the Key condition builder
/*	keyCond := expression.Key("FirstName-index").Equal(expression.Value(firstName))

	// Create the project expression builder with a names list.
	proj := expression.NamesList(expression.Name("RollNo"), expression.Name("LastName"))

	// Combine the key condition, and projection together as a DynamoDB expression
	// builder.
	expr, err := expression.NewBuilder().
		WithKeyCondition(keyCond).
		WithProjection(proj).
		Build()
	if err != nil {
		fmt.Println(err)
	} */

	// Use the built expression to populate the DynamoDB Query's API input
/*	// parameters.
	input := &dynamodb.QueryInput{
		TableName: aws.String(tableName),
        IndexName: aws.String("countryId"),
		KeyConditionExpression: aws.String("gsi1pk = :gsi1pk and gsi1sk > :gsi1sk"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
            ":gsi1pk": &types.AttributeValueMemberS{Value: "123"},
            ":gsi1sk": &types.AttributeValueMemberN{Value: "20150101"},
        },
	}

	result, err := dynaClient.Query(input)
	if err != nil {
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

	items := new([]Pagination)

	err1 := dynamodbattribute.UnmarshalListOfMaps(result.Items, &items)

	if err1!= nil{
		fmt.Println("Unable to un marshal Map")
	}
	

	//return items, nil
	fmt.Println(result)
}
/*	expr, err := expression.NewBuilder().WithFilter(
        expression.And(
            expression.AttributeNotExists(expression.Name("deletedAt")),
            expression.Contains(expression.Name("FirstName"), "Somashankar"),
        ),
    ).Build()
    if err != nil {
        panic(err)
    }

    out, err := dynaClient.Scan(&dynamodb.ScanInput{
        TableName:                 aws.String("PaginationTable"),
        FilterExpression:          expr.Filter(),
        ExpressionAttributeNames:  expr.Names(),
        ExpressionAttributeValues: expr.Values(),
    })
    if err != nil {
        panic(err)
    }

    fmt.Println(out.Items)

}*/