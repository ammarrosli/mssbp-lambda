package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func InitEnv() {
	_ = os.Setenv("DDBTABLE_MSSBP", "develop-mssbp")
	_ = os.Setenv("MG_DOMAIN", "mailgun.digitalsymphony.it")
	_ = os.Setenv("MG_ADMIN_EMAIL", "dsadmin@maildrop.cc")
	_ = os.Setenv("MG_SENDER_EMAIL", "noreply@mssbp.com.my")
	_ = os.Setenv("SENTINO_ENDPOINT", "https://www.sentinocrm.com/service2/register?ws=1")
	_ = os.Setenv("SENTINO_PROJECT_ID", "Jtd9tL1820180423115313")
	_ = os.Setenv("SENTINO_SOURCE_ID", "14")
	_ = os.Setenv("TELEGRAM_CHAT_ID", "-1001267327365")
	_ = os.Setenv("FUNCTION_NAME", "mssbp")
	_ = os.Setenv("TEST", "false")
}

func TestEDM(t *testing.T) {
	InitEnv()

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	request := events.APIGatewayProxyRequest{
		Headers: headers,
		Body: `{
			"name": "Tester",
			"email": "dstester@maildrop.cc",
			"phone": "0123456789",
			"location": "",
			"type": "",
			"utm_sources": "",
			"source": "SearchOP1"
		}`,
	}

	response, _ := StartHandler(request)
	fmt.Println(response.Body)
	assert.Equal(t, 200, response.StatusCode)
}
