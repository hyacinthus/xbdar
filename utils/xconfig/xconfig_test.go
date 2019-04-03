package xconfig

import (
	"testing"
)

func TestResetValues(t *testing.T) {
	rv := new(SyncValues)
	var (
		b1, b2 bool
		s1, s2 string
		i1, i2 int
	)

	// add
	rv.AddBool(&b1, &b2)
	rv.AddString(&s1, &s2)
	rv.AddInt(&i1, &i2)

	// change
	b1 = true
	s1 = "a"
	i1 = 123

	// sync
	rv.Sync()

	if b1 != b2 {
		t.Errorf("want %v, got %v", b1, b2)
	}
	if s1 != s2 {
		t.Errorf("want %v, got %v", s1, s2)
	}
	if i1 != i2 {
		t.Errorf("want %v, got %v", i1, i2)
	}
}
