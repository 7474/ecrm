package ecrm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	ecsTypes "github.com/aws/aws-sdk-go-v2/service/ecs/types"
)

func TestIsTaskDefinitionUnavailableError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "client exception unable to describe",
			err: &ecsTypes.ClientException{
				Message: aws.String("Unable to describe task definition."),
			},
			want: true,
		},
		{
			name: "wrapped client exception unable to describe",
			err: fmt.Errorf("wrapped: %w", &ecsTypes.ClientException{
				Message: aws.String("unable to describe task definition"),
			}),
			want: true,
		},
		{
			name: "other client exception",
			err: &ecsTypes.ClientException{
				Message: aws.String("some other client error"),
			},
			want: false,
		},
		{
			name: "non client exception",
			err:  errors.New("boom"),
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isTaskDefinitionUnavailableError(tt.err); got != tt.want {
				t.Errorf("isTaskDefinitionUnavailableError()=%v want=%v", got, tt.want)
			}
		})
	}
}
