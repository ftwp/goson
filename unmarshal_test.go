package goson

import (
	"testing"
)

func TestUnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
		v    interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				data: []byte(`{"_this_is_message_too":{"last":222,"this_is_version":222,"version":222},"message":{"last":1647586620828811896,"version":1},"message_v2":{"last":1647586620828811896,"version":1},"this_is_message":null,"type":7,"version":"v0.0.1"}`),
				v:    &TestJSON{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UnmarshalJSON(tt.args.data, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			json, err := MarshalJSON(tt.args.v)
			if err != nil {
				return
			}
			t.Log(string(json))
		})
	}
}
