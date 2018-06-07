package wif

import "testing"

const privateKey = "65ce06a2274323d0f6d83c661aea13b7fd6c2259b9aa7d7e7483bd8babe4293201"
const privateKeyWif = "QsmQyTSvpDrwo8Ma4KxKcGYq5Q5Gaoa6HadbNCGReyscffZk1D2F"

func TestDecode(t *testing.T) {
	r, e := Decode(privateKeyWif)

	if e != nil {
		t.Errorf("Decode threw an error: %s.", e)
	}

	if r.hexString != privateKey {
		t.Errorf("Decode failed, got: %s, want: %s.", r.hexString, privateKey)
	}
}
