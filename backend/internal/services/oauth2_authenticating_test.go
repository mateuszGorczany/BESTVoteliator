package services

import (
	"strings"
	"testing"

	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"github.com/spf13/viper"
)

func init() {
	common.LoadConfig()
}

func TestValidateGoogleJWT(t *testing.T) {
	type args struct {
		tokenString string
	}
	type email string
	tests := []struct {
		name    string
		args    args
		want    email
		wantErr bool
	}{
		{
			name:    "Test token",
			args:    args{tokenString: viper.GetString("test.JWT")},
			want:    email(strings.ToLower(viper.GetString("test.UserEmail"))),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateGoogleJWT(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateGoogleJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			email := email(strings.ToLower(got.Email))
			if email != tt.want {
				t.Errorf("ValidateGoogleJWT() = %v, want %v", email, tt.want)
			}
		})
	}
}
