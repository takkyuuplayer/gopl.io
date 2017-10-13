package main

import "testing"

func TestPackageLevelDeclaration(t *testing.T) {
	// boiling_test.go
	if boilingF != 212 {
		t.Errorf(`boilingF = %#v, want %#v`, boilingF, 212)
	}
}

func TestFunctionDecralation(t *testing.T) {
	const freezingF, boilingF = 32.0, 212.0

	if fToC(freezingF) != 0 {
		t.Errorf(`fToC(freezingF) = %#v, want %#v`, fToC(freezingF), 0)
	}

	if fToC(boilingF) != 100 {
		t.Errorf(`fToC(boilingF) = %#v, want %#v`, fToC(boilingF), 100)
	}

}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
