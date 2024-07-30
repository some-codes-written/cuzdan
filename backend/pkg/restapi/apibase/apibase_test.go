// path: backend\pkg\restapi\apibase\apibase.go

package apibase

import (
	"reflect"
	"testing"
)

func TestNewResultInfo(t *testing.T) {
	type args struct {
		code         int
		message      string
		content_type string
		data         interface{}
	}
	tests := []struct {
		name string
		args args
		want *ResultInfo
	}{
		{
			name: "TestNewResultInfo",
			args: args{
				code:         200,
				message:      "OK",
				content_type: "application/json",
				data:         nil,
			},
			want: &ResultInfo{
				StatusCode:  200,
				StatusText:  "OK",
				Message:     "OK",
				Data:        nil,
				ContentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewResultInfo(tt.args.code, tt.args.message, tt.args.content_type, tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResultInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
