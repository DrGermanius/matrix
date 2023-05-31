package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

const (
	metaTemplate = "testData %s #%d; when %s, then %s"
	errTemplate  = "%s \n got = %v \n want = %v \n"
)

var (
	validIntMatrix    = [][]string{{"1", "2", "3"}, {"4", "5", "6"}, {"7", "8", "9"}}
	matrixWithStrings = [][]string{{"a", "b", "c"}, {"4", "5", "6"}, {"7", "8", "9"}}
)

func Test_invertMatrix(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args args
		want [][]string
	}{
		{
			name: "invert matrix happy path",
			when: "everything is OK",
			then: "the columns and rows are inverted in matrix should be inverted",

			args: args{matrix: validIntMatrix},
			want: [][]string{{"1", "4", "7"}, {"2", "5", "8"}, {"3", "6", "9"}},
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			if got := invertMatrix(tt.args.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}

func Test_matrixToFlatString(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args args
		want string
	}{
		{
			name: "matrix to flat string happy path",
			when: "everything is OK",
			then: "matrix presented as flat string",

			args: args{matrix: validIntMatrix},
			want: "1,2,3,4,5,6,7,8,9",
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			if got := matrixToFlatString(tt.args.matrix); got != tt.want {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}

func Test_matrixToString(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args args
		want string
	}{
		{
			name: "matrix to string happy path",
			when: "everything is OK",
			then: "matrix presented as flat string",

			args: args{matrix: validIntMatrix},
			want: "1,2,3\n4,5,6\n7,8,9\n",
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			if got := matrixToString(tt.args.matrix); got != tt.want {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}

func Test_multiplyIntMatrix(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args    args
		want    int
		wantErr error
	}{
		{
			name: "matrix multiply happy path",
			when: "everything is OK",
			then: "the product of matrix elements should be returned",

			args: args{matrix: validIntMatrix},
			want: 362880,
		},
		{
			name: "matrix to string unhappy path",
			when: "matrix consists the non-interger elements",
			then: "error should be returned",

			args:    args{matrix: matrixWithStrings},
			wantErr: errMatrixConsistsNonIntegerElems,
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			got, err := multiplyIntMatrix(tt.args.matrix)
			if tt.wantErr != nil {
				if !errors.Is(tt.wantErr, err) {
					t.Errorf(errTemplate, meta, err, tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}

func Test_stringMatrixToInt(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args    args
		want    [][]int
		wantErr error
	}{
		{
			name: "string matrix to int happy path",
			when: "everything is OK",
			then: "the int matrix should be returned",

			args: args{matrix: validIntMatrix},
			want: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
		},
		{
			name: "string matrix to int unhappy path",
			when: "the initial matrix consists strings",
			then: "the error should be returned",

			args:    args{matrix: matrixWithStrings},
			wantErr: errMatrixConsistsNonIntegerElems,
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			got, err := stringMatrixToInt(tt.args.matrix)
			if tt.wantErr != nil {
				if !errors.Is(tt.wantErr, err) {
					t.Errorf(errTemplate, meta, err, tt.wantErr)
				}
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf(errTemplate, meta, err, tt.wantErr)
			}
		})
	}
}

func Test_sumIntMatrix(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args    args
		want    int
		wantErr error
	}{
		{
			name: "matrix to string happy path",
			when: "everything is OK",
			then: "the sum os matrix elements should be returned",

			args: args{matrix: validIntMatrix},
			want: 45,
		},
		{
			name: "matrix to string unhappy path",
			when: "matrix consists the non-interger elements",
			then: "error should be returned",

			args:    args{matrix: matrixWithStrings},
			wantErr: errMatrixConsistsNonIntegerElems,
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			got, err := sumIntMatrix(tt.args.matrix)
			if tt.wantErr != nil {
				if !errors.Is(tt.wantErr, err) {
					t.Errorf(errTemplate, meta, err, tt.wantErr)
				}
				return
			}
			if got != tt.want {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}

func Test_isMatrixSquare(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name string
		when string // description of testData conditions
		then string // description of expected result

		args args
		want bool
	}{
		{
			name: "is matrix square happy path",
			when: "the matrix is in square format",
			then: "true should be returned",

			args: args{matrix: validIntMatrix},
			want: true,
		},
		{
			name: "is matrix square happy path",
			when: "the matrix is not in square format",
			then: "false should be returned",

			args: args{matrix: [][]string{{"1"}, {"2", "3"}, {"4", "5", "6"}}},
			want: false,
		},
	}
	for i, tt := range tests {
		meta := fmt.Sprintf(metaTemplate, tt.name, i, tt.when, tt.then)

		t.Run(tt.name, func(t *testing.T) {
			if got := isMatrixSquare(tt.args.matrix); got != tt.want {
				t.Errorf(errTemplate, meta, got, tt.want)
			}
		})
	}
}
