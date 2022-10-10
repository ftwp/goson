package goson

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestGSON_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		g       goson
		want    []byte
		wantErr bool
	}{
		{
			name: "",
			g: goson{
				v: map[string]json.RawMessage{
					//"test": json.Marshal("test value 1"),
				},
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.g.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_goson_UnmarshalJSON(t *testing.T) {
	type fields struct {
		v map[string]json.RawMessage
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "",
			fields: fields{
				v: make(map[string]json.RawMessage),
			},
			args: args{
				data: []byte(`{"version":"0.0.1","message":{"last":1647586620828811896,"version":1},"type":7}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &goson{
				v: tt.fields.v,
			}
			if err := g.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			g.Value()
			marshalJSON, err := g.MarshalJSON()
			if err != nil {
				return
			}
			t.Log("marshal:", string(marshalJSON))
		})
	}
}

type TestMessage struct {
	Last            int64 `json:"last"`
	Version         int   `json:"version"`
	this_is_version int   `json:"this_is_version" json-getter:"MyVersion"`
}

func (t *TestMessage) ThisIsVersion() int {
	return t.this_is_version
}

func (t *TestMessage) MyVersion() int {
	return t.this_is_version
}

func (t *TestMessage) SetThisIsVersion(this_is_version int) {
	t.this_is_version = this_is_version
}

type TestJSON struct {
	version              string       `json:"version" json-getter:"Version" json-setter:"SetMyVersion"`
	message              TestMessage  `json:"message"`
	MessageV2            TestMessage  `json:"message_v2"`
	this_is_message      *TestMessage `json:"this_is_message"`
	_this_is_message_too *TestMessage `json:"_this_is_message_too" json-setter:"SetThisMessageToo"`
	Type                 int          `json:"type"`
	TypeV2               int          `json:"-"`
}

func (t *TestJSON) SetMyVersion(version string) {
	t.version = version
}

func (t *TestJSON) ThisIsMessageToo() *TestMessage {
	return t._this_is_message_too
}

func (t *TestJSON) Set_this_is_message_too(_this_is_message_too *TestMessage) {
	t._this_is_message_too = _this_is_message_too
}

func (t *TestJSON) SetThisMessageToo(_this_is_message_too *TestMessage) {
	t._this_is_message_too = _this_is_message_too
}

func (t *TestJSON) ThisIsMessage() *TestMessage {
	return t.this_is_message
}

func (t *TestJSON) SetThisIsMessage(this_is_message *TestMessage) {
	t.this_is_message = this_is_message
}

func (t *TestJSON) Message() TestMessage {
	return t.message
}

func (t *TestJSON) SetMessage(message TestMessage) {
	t.message = message
}

func (t TestJSON) Version() string {
	return t.version
}
