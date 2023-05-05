package aws_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tae2089/bungee/internal/di"
)

func Test_awsServiceImpl_GetEc2List(t *testing.T) {
	awsService := di.InitAwsService("", "ap-northeast-2")
	t.Run("get aws ec2 list", func(t *testing.T) {
		err := awsService.GetEc2List()
		assert.Nil(t, err)
	})
}
