package casts_test

import (
	"github.com/Sunset-/sunset-go/utils/casts"
	"testing"
)

func TestStrToFloat64(t *testing.T){
	str := "123"
	expected := float64(123)

	res := casts.StrToFloat64(str)

	if res!=expected{
		t.Errorf("casts.StrToFloat64(%s) = %f , not %f",str,res,expected)
	}
}