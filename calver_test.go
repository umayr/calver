package calver

import (
	"os"
	"testing"
	"time"
)

func mockNowFunc(fn func() time.Time) func() {
	now = fn
	return func() {
		now = time.Now
	}
}

func TestMain(m *testing.M) {
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007, 2, 5, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	os.Exit(m.Run())
}

func TestNew_YYYYMMDD(t *testing.T) {
	c, _ := New("YYYY.MM.DD", "")
	if c.String() != "YYYY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.2.5"
		v1 = "2007.2.5-1"
		v2 = "2007.2.5-dev.2"
		v3 = "2007.2.5-2"
		v4 = "2007.2.5-dev.3"
		v5 = "2007.2.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YYMMDD(t *testing.T) {
	c, _ := New("YY.MM.DD", "")
	if c.String() != "YY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.2.5"
		v1 = "7.2.5-1"
		v2 = "7.2.5-dev.2"
		v3 = "7.2.5-2"
		v4 = "7.2.5-dev.3"
		v5 = "7.2.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YY0M0D(t *testing.T) {
	c, _ := New("YY.0M.0D", "")
	if c.String() != "YY.0M.0D" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.02.05"
		v1 = "7.02.05-1"
		v2 = "7.02.05-dev.2"
		v3 = "7.02.05-2"
		v4 = "7.02.05-dev.3"
		v5 = "7.02.05-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_0YMMDD(t *testing.T) {
	c, _ := New("0Y.MM.DD", "")
	if c.String() != "0Y.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "07.2.5"
		v1 = "07.2.5-1"
		v2 = "07.2.5-dev.2"
		v3 = "07.2.5-2"
		v4 = "07.2.5-dev.3"
		v5 = "07.2.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_0Y0MDD(t *testing.T) {
	c, _ := New("0Y.0M.DD", "")
	if c.String() != "0Y.0M.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "07.02.5"
		v1 = "07.02.5-1"
		v2 = "07.02.5-dev.2"
		v3 = "07.02.5-2"
		v4 = "07.02.5-dev.3"
		v5 = "07.02.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_0Y0M0D(t *testing.T) {
	c, _ := New("0Y.0M.0D", "")
	if c.String() != "0Y.0M.0D" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "07.02.05"
		v1 = "07.02.05-1"
		v2 = "07.02.05-dev.2"
		v3 = "07.02.05-2"
		v4 = "07.02.05-dev.3"
		v5 = "07.02.05-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YYWWDD(t *testing.T) {
	c, _ := New("YY.WW.DD", "")
	if c.String() != "YY.WW.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.6.5"
		v1 = "7.6.5-1"
		v2 = "7.6.5-dev.2"
		v3 = "7.6.5-2"
		v4 = "7.6.5-dev.3"
		v5 = "7.6.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YY0WDD(t *testing.T) {
	c, _ := New("YY.0W.DD", "")
	if c.String() != "YY.0W.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.06.5"
		v1 = "7.06.5-1"
		v2 = "7.06.5-dev.2"
		v3 = "7.06.5-2"
		v4 = "7.06.5-dev.3"
		v5 = "7.06.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YYWW(t *testing.T) {
	c, _ := New("YY.WW", "")
	if c.String() != "YY.WW" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.6"
		v1 = "7.6-1"
		v2 = "7.6-dev.2"
		v3 = "7.6-2"
		v4 = "7.6-dev.3"
		v5 = "7.6-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_YYMM(t *testing.T) {
	c, _ := New("YY.MM", "")
	if c.String() != "YY.MM" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "7.2"
		v1 = "7.2-1"
		v2 = "7.2-dev.2"
		v3 = "7.2-2"
		v4 = "7.2-dev.3"
		v5 = "7.2-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_0Y0M(t *testing.T) {
	c, _ := New("0Y.0M", "")
	if c.String() != "0Y.0M" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "07.02"
		v1 = "07.02-1"
		v2 = "07.02-dev.2"
		v3 = "07.02-2"
		v4 = "07.02-dev.3"
		v5 = "07.02-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_WWDD(t *testing.T) {
	c, _ := New("WW.DD", "")
	if c.String() != "WW.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "6.5"
		v1 = "6.5-1"
		v2 = "6.5-dev.2"
		v3 = "6.5-2"
		v4 = "6.5-dev.3"
		v5 = "6.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_0W0D(t *testing.T) {
	c, _ := New("0W.0D", "")
	if c.String() != "0W.0D" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "06.05"
		v1 = "06.05-1"
		v2 = "06.05-dev.2"
		v3 = "06.05-2"
		v4 = "06.05-dev.3"
		v5 = "06.05-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_DifferentModifier(t *testing.T) {
	c, _ := New("YYYY.MM.DD", "alpha")
	if c.String() != "YYYY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.2.5"
		v1 = "2007.2.5-1"
		v2 = "2007.2.5-alpha.2"
		v3 = "2007.2.5-2"
		v4 = "2007.2.5-alpha.3"
		v5 = "2007.2.5-3"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}

	if c.String() != v5 {
		t.Errorf("release version should be %s but it was %s", v5, c.String())
	}
}

func TestNew_DifferentDay(t *testing.T) {
	n := 0
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007, 2, 5+n, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	c, _ := New("YYYY.MM.DD", "")
	if c.String() != "YYYY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.2.5"
		v1 = "2007.2.6"
		v2 = "2007.2.7-dev"
		v3 = "2007.2.8"
		v4 = "2007.2.9-dev"
		v5 = "2007.2.10"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	n++

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	n++

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	n++

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	n++

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	n++

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}
}

func TestNew_DifferentMonth(t *testing.T) {
	m := time.January
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007, m, 5, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	c, _ := New("YYYY.MM.DD", "")
	if c.String() != "YYYY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.1.5"
		v1 = "2007.2.5"
		v2 = "2007.3.5-dev"
		v3 = "2007.4.5"
		v4 = "2007.5.5-dev"
		v5 = "2007.6.5"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	m = time.February

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	m = time.March

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	m = time.April

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	m = time.May

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	m = time.June

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}
}

func TestNew_DifferentYear(t *testing.T) {
	n := 0
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007+n, 2, 5, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	c, _ := New("YYYY.MM.DD", "")
	if c.String() != "YYYY.MM.DD" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.2.5"
		v1 = "2008.2.5"
		v2 = "2009.2.5-dev"
		v3 = "2010.2.5"
		v4 = "2011.2.5-dev"
		v5 = "2012.2.5"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	n++

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	n++

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	n++

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	n++

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	n++

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}
}

func TestNew_DifferentWeek(t *testing.T) {
	m := time.January
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007, m, 5, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	c, _ := New("YYYY.WW", "")
	if c.String() != "YYYY.WW" {
		t.Error("empty version doesn't return the format")
	}

	const (
		v0 = "2007.1"
		v1 = "2007.6"
		v2 = "2007.10-dev"
		v3 = "2007.14"
		v4 = "2007.18-dev"
		v5 = "2007.23"
	)

	r0 := c.Release()
	if r0 != v0 {
		t.Errorf("release version should be %s but it was %s", v0, r0)
	}

	m = time.February

	r1 := c.Release()
	if r1 != v1 {
		t.Errorf("release version should be %s but it was %s", v1, r1)
	}

	m = time.March

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("release version should be %s but it was %s", v2, r2)
	}

	m = time.April

	r3 := c.Release()
	if r3 != v3 {
		t.Errorf("release version should be %s but it was %s", v3, r3)
	}

	m = time.May

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("release version should be %s but it was %s", v4, r4)
	}

	m = time.June

	r5 := c.Release()
	if r5 != v5 {
		t.Errorf("release version should be %s but it was %s", v5, r5)
	}
}

func TestNew_UnsupportedFormat(t *testing.T) {
	_, err := New("YYYY.XX.HH", "")
	if err == nil || err.Error() != "unsupported format: YYYY.XX.HH" {
		t.Error("invalid format should not supported")
	}
}

func TestParse(t *testing.T) {
	c0, _ := Parse("2007.1.1", "YYYY.MM.DD", "")
	if c0.String() != "2007.1.1" {
		t.Errorf("failed to parse the version, expected 2007.1.1 but got %s", c0.String())
	}

	c1, _ := Parse("2007.1.1-dev", "YYYY.MM.DD", "")
	if c1.String() != "2007.1.1-dev" {
		t.Errorf("failed to parse the version, expected 2007.1.1-dev but got %s", c0.String())
	}

	c2, _ := Parse("2007.1.1-1000", "YYYY.MM.DD", "")
	if c2.String() != "2007.1.1-1000" {
		t.Errorf("failed to parse the version, expected 2007.1.1-1000 but got %s", c0.String())
	}

	c3, _ := Parse("2007.1.1-dev.99", "YYYY.MM.DD", "")
	if c3.String() != "2007.1.1-dev.99" {
		t.Errorf("failed to parse the version, expected 2007.1.1-dev.99 but got %s", c0.String())
	}
}

func TestCalVer_PreRelease(t *testing.T) {
	reset := mockNowFunc(func() time.Time {
		return time.Date(2007, 2, 5, 0, 0, 0, 0, time.UTC)
	})
	defer reset()

	c, _ := New("YYYY.MM.DD", "")

	const (
		v0 = "2007.2.5-dev"
		v1 = "2007.2.5-dev.1"
		v2 = "2007.2.5-dev.2"
		v3 = "2007.2.5-dev.3"
		v4 = "2007.2.5-dev.4"
	)
	r0 := c.PreRelease()
	if r0 != v0 {
		t.Errorf("prerelease version should be %s but it was %s", v0, r0)
	}

	r1 := c.PreRelease()
	if r1 != v1 {
		t.Errorf("prerelease version should be %s but it was %s", v1, r1)
	}

	r2 := c.PreRelease()
	if r2 != v2 {
		t.Errorf("prerelease version should be %s but it was %s", v2, r2)
	}

	r3 := c.PreRelease()
	if r3 != v3 {
		t.Errorf("prerelease version should be %s but it was %s", v3, r3)
	}

	r4 := c.PreRelease()
	if r4 != v4 {
		t.Errorf("prerelease version should be %s but it was %s", v4, r4)
	}

	p, _ := Parse(v0, "YYYY.MM.DD", "")
	if p.String() != v0 {
		t.Errorf("prerelease version should be %s but it was %s", v0, p.String())
	}

	r1 = p.PreRelease()
	if r1 != v1 {
		t.Errorf("prerelease version should be %s but it was %s", v1, r1)
	}

	r2 = p.PreRelease()
	if r2 != v2 {
		t.Errorf("prerelease version should be %s but it was %s", v2, r2)
	}

	r3 = p.PreRelease()
	if r3 != v3 {
		t.Errorf("prerelease version should be %s but it was %s", v3, r3)
	}

	r4 = p.PreRelease()
	if r4 != v4 {
		t.Errorf("prerelease version should be %s but it was %s", v4, r4)
	}

	// older date
	p1, _ := Parse("2007.2.4-dev.4", "YYYY.MM.DD", "")

	r5 := p1.PreRelease()
	if r5 != v0 {
		t.Errorf("prerelease version should be %s but it was %s", v0, r5)
	}
}
