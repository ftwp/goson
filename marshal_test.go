package goson

import (
	"reflect"
	"testing"
)

func TestMarshalJSON(t *testing.T) {
	type args struct {
		v interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "",
			args: args{
				v: &TestJSON{
					version: "v0.0.1",
					message: TestMessage{
						Last:            1647586620828811896,
						Version:         1,
						this_is_version: 76,
					},
					MessageV2: TestMessage{
						Last:            1647586620828811896,
						Version:         1,
						this_is_version: 5,
					},
					//this_is_message: &TestMessage{
					//	Last:            0,
					//	Version:         0,
					//	this_is_version: 0,
					//},
					_this_is_message_too: &TestMessage{
						Last:            222,
						Version:         222,
						this_is_version: 222,
					},
					Type:   7,
					TypeV2: 5,
				},
			},
			want:    `{"version":"0.0.1","message":{"last":1647586620828811896,"version":1},"type":7}`,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MarshalJSON(tt.args.v)
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), tt.want)
			}
		})
	}
}
