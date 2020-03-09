package service

import (
	"bytes"
	"io"
	"reflect"
	"testing"
)

func Test_parseHTML(t *testing.T) {
	type args struct {
		data io.Reader
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "test with links",
			args: args{data: bytes.NewBufferString(`<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`)},
			want: []string{"foo", "/bar/baz"},
		},
		{
			name: "empty list",
			args: args{data: bytes.NewBufferString(`<h1>Hello</h1>`)},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseHTML(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseHTML() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseHTML() got = %v, want %v", got, tt.want)
			}
		})
	}
}
