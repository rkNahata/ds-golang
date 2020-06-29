package dll

import "testing"

func TestLRUCacheAll(t *testing.T) {
	c := Constructor(4)
	c.Set(1, 111)
	if c.list.Size() != 1 {
		t.Errorf("expected 1 but got %d", c.list.Size())
	}
	c.Set(2, 222)
	c.Set(3, 333)
	c.Set(4, 444)
	if c.list.Size() != 4 {
		t.Errorf("expected 4 but got %d", c.list.Size())
	}
	c.Set(5, 555)
	if c.list.Size() != 4 {
		t.Errorf("expected 4 but got %d", c.list.Size())
	}
	if c.Get(1) != -1 {
		t.Errorf("expected -1")
	}
	if c.Get(4) != 444 {
		t.Errorf("expected 444")
	}
	c.Set(4,1234)
	if c.Get(4) != 1234 {
		t.Errorf("expected 1234")
	}

}
