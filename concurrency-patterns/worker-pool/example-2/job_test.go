package main

import (
	"context"
	"errors"
	"reflect"
	"testing"
)

var (
	errDefault = errors.New("wrong argument type")
	descriptor = JobDescriptor{
		ID:    JobID("1"),
		JType: jobType("anyType"),
		Metadata: jobMetadata{
			"foo": "foo",
			"bar": "bar",
		},
	}
	execFn = func(ctx context.Context, args interface{}) (interface{}, error) {
		argVal, ok := args.(int)
		if !ok {
			return nil, errDefault
		}

		return argVal * 2, nil
	}
)

func TestJobExecute(t *testing.T) {
	ctx := context.TODO()

	type fields struct {
		descriptor JobDescriptor
		execFn     ExecutionFn
		args       interface{}
	}

	tests := []struct {
		name   string
		fields fields
		want   Result
	}{
		{
			name: "job execution success",
			fields: fields{
				descriptor: descriptor,
				execFn:     execFn,
				args:       10,
			},
		},
		{
			name: "job execution failure",
			fields: fields{
				descriptor: descriptor,
				execFn:     execFn,
				args:       "10",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			job := Job{
				Descriptor: test.fields.descriptor,
				ExecFn:     test.fields.execFn,
				Args:       test.fields.args,
			}

			got := job.execute(ctx)
			if test.want.Err != nil {
				if !reflect.DeepEqual(got.Err, test.want.Err) {
					t.Errorf("Job.execute() = %v, wantError %v", got.Err, test.want.Err)
				}
				return
			}

			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("Job.execute() = %v, want %v", got, test.want)
			}
		})
	}
}
