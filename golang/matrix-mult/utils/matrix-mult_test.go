package utils

import (
	"errors"
	"reflect"
	"testing"
)

func TestMatrixMultWrongDimension(t *testing.T){
	wrongA := [][]float32{
					{1,3}, // one less element
					{4,5,6},
			}

	_,_,err := getMatrixDimension(&wrongA)
	if !errors.Is(err,ErrMalformedMatrix){
		t.Fail()
	}

}

func TestMatrixMult(t *testing.T){
	A := [][]float32{
		{1,2,3},
		{4,5,6},
	}

	B := [][]float32{
		{7,8},
		{9,10},
		{11,12},
	}

	var C [][]float32
	
	matMult(&A,&B,&C)

	expectedC1 := [][]float32{
		{58 , 64},
		{139,  154},	
	}

	if reflect.DeepEqual(C, expectedC1) {
		t.Fail()
	}

}