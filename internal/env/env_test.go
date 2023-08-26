package env

import (
	"net/http"
	"reflect"
	"testing"
)

func TestNewEnv(t *testing.T) {
	type args struct {
		contextFile string
	}
	tests := []struct {
		name string
		args args
		want Environment
	}{
		{
			name: "hello",
			args: args{
				contextFile: "./test-ace-context.yaml",
			},
			want: Environment{
				ContextFile: "./test-ace-context.yaml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewEnv(tt.args.contextFile); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEnvironment_loadAceContext(t *testing.T) {
	tests := []struct {
		name    string
		env     *Environment
		wantErr bool
	}{
		{
			name: "load context from file",
			env: &Environment{
				ContextFile: "./internal/env/test-ace-context.yaml",
				AceContext: aceContex{
					AceEnvVersion: "7.0.20",
					Domain:        "test.act3-ace.ai",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.env.loadAceContext(); (err != nil) != tt.wantErr {
				t.Errorf("Environment.loadAceContext() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestEnvironment_EnvContextHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		env  *Environment
		args args
	}{
		{
			name: "env context handler",
			env:  &Environment{},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.env.EnvContextHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestEnvironment_EnvVersionHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		env  *Environment
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.env.EnvVersionHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestEnvironment_EnvVersionBadgeHandler(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		env  *Environment
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.env.EnvVersionBadgeHandler(tt.args.w, tt.args.r)
		})
	}
}
