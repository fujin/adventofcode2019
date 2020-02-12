package main

import (
	"reflect"
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func TestNewComputer(t *testing.T) {
	type args struct {
		intcodes []int
	}
	tests := []struct {
		name string
		args args
		want *Computer
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewComputer(tt.args.intcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewComputer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComputer_run(t *testing.T) {
	type fields struct {
		intcodes []int
	}
	type args struct {
		oc opcode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				intcodes: tt.fields.intcodes,
			}
			c.run(tt.args.oc)
		})
	}
}

func TestComputer_parse(t *testing.T) {
	type fields struct {
		intcodes []int
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Computer{
				intcodes: tt.fields.intcodes,
			}
			c.parse()
		})
	}
}
