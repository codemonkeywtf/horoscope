package main

import (
	//	"bytes"
	"context"
	// "encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"io/ioutil"
	"net/http"
	"regexp"
)

func getHoroscope(sign string) string {
	var horoscope = ""
	re := regexp.MustCompile(`horoscope\s=?\s(.*)?;`)
	response, err := http.Get("http://astrology.kudosmedia.net/index.php/m/" + sign + "?day=today")

	if err != nil {
		fmt.Printf("The request failed, error: %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		horoscope = re.FindStringSubmatch(string(data))[1]
	}
	return horoscope
}

func handler(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	res := getHoroscope(request.QueryStringParameters["sign"])
	// horoscope := &res
	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       res,
	}, nil
}

func main() {
	lambda.Start(handler)

}
