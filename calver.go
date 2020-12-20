package calver

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const (
	// FullYear notation for CalVer which is YYYY
	FullYear = "YYYY"
	// ShortYear notation for CalVer which is YY
	ShortYear = "YY"
	// PaddedYear notation for CalVer which is 0Y
	PaddedYear = "0Y"
	// ShortMonth notation for CalVer which is MM
	ShortMonth = "MM"
	// PaddedMonth notation for CalVer which is 0M
	PaddedMonth = "0M"
	// ShortWeek notation for CalVer which is WW
	ShortWeek = "WW"
	// PaddedWeek notation for CalVer which is 0W
	PaddedWeek = "0W"
	// ShortDay notation for CalVer which is DD
	ShortDay = "DD"
	// PaddedDay notation for CalVer which is 0D
	PaddedDay = "0D"
)

type segment int

const (
	segmentEmpty segment = iota
	segmentFullYear
	segmentShortYear
	segmentPaddedYear
	segmentShortMonth
	segmentPaddedMonth
	segmentShortWeek
	segmentPaddedWeek
	segmentShortDay
	segmentPaddedDay
)

func (s segment) String() string {
	switch s {
	case segmentFullYear:
		return FullYear
	case segmentShortYear:
		return ShortYear
	case segmentPaddedYear:
		return PaddedYear
	case segmentShortMonth:
		return ShortMonth
	case segmentPaddedMonth:
		return PaddedMonth
	case segmentShortWeek:
		return ShortWeek
	case segmentPaddedWeek:
		return PaddedWeek
	case segmentShortDay:
		return ShortDay
	case segmentPaddedDay:
		return PaddedDay
	case segmentEmpty:
		return ""
	default:
		panic("invalid format segment")
	}
}

func (s segment) pattern() string {
	switch s {
	case segmentFullYear:
		return "2006"
	case segmentPaddedYear:
		return "06"
	case segmentShortMonth:
		return "1"
	case segmentPaddedMonth:
		return "01"
	case segmentShortDay:
		return "2"
	case segmentPaddedDay:
		return "02"
	default:
		panic("unsupported format segment")
	}
}

func (s segment) conv(t time.Time) string {
	switch s {
	case segmentEmpty:
		return ""
	case segmentShortWeek:
		_, w := t.ISOWeek()
		return fmt.Sprintf("%d", w)
	case segmentPaddedWeek:
		_, w := t.ISOWeek()
		return fmt.Sprintf("%02d", w)
	case segmentShortYear:
		y := t.Format("06")
		if strings.HasPrefix(y, "0") {
			return strings.TrimPrefix(y, "0")
		}
		return y
	}

	return t.Format(s.pattern())
}

func (s segment) parse(raw string) (string, error) {
	t, err := time.Parse(s.pattern(), raw)
	if err != nil {
		return "", fmt.Errorf("provided string doesn't match the format segment: %s", s.String())
	}
	return t.Format(s.pattern()), nil
}

func newSegment(s string) (segment, error) {
	switch s {
	case FullYear:
		return segmentFullYear, nil
	case ShortYear:
		return segmentShortYear, nil
	case PaddedYear:
		return segmentPaddedYear, nil
	case ShortMonth:
		return segmentShortMonth, nil
	case PaddedMonth:
		return segmentPaddedMonth, nil
	case ShortWeek:
		return segmentShortWeek, nil
	case PaddedWeek:
		return segmentPaddedWeek, nil
	case ShortDay:
		return segmentShortDay, nil
	case PaddedDay:
		return segmentPaddedDay, nil
	default:
		return segment(0), fmt.Errorf("invalid format segment: %s", s)
	}
}

type format struct {
	major segment
	minor segment
	micro segment
}

func (f format) parse(raw string) (string, string, string, error) {
	segs := strings.Split(raw, ".")

	major, err := f.major.parse(segs[0])
	if err != nil {
		return "", "", "", err
	}

	minor, err := f.minor.parse(segs[1])
	if err != nil {
		return "", "", "", err
	}

	var micro string
	if len(segs) > 2 {
		micro, err = f.micro.parse(segs[2])
		if err != nil {
			return "", "", "", err
		}
	}

	return major, minor, micro, nil
}

func (f format) String() string {
	v := ""

	if f.major != segmentEmpty {
		v += f.major.String()
	}

	if f.minor != segmentEmpty {
		v += fmt.Sprintf(".%s", f.minor.String())
	}

	if f.micro != segmentEmpty {
		v += fmt.Sprintf(".%s", f.micro.String())
	}

	return v
}

func newFormat(raw string) (*format, error) {
	parts := strings.Split(raw, ".")

	if len(parts) < 2 {
		return nil, fmt.Errorf("major, minor and micro are all required for a valid format")
	}

	if len(parts) > 3 {
		return nil, fmt.Errorf("format could only consist three parts: major, minor and micro")
	}

	major, err := newSegment(parts[0])
	if err != nil {
		return nil, fmt.Errorf("invalid format segment: %s", parts[0])
	}

	minor, err := newSegment(parts[1])
	if err != nil {
		return nil, fmt.Errorf("invalid format segment: %s", parts[1])
	}

	if len(parts) == 2 {
		return &format{major, minor, segmentEmpty}, nil
	}

	micro, err := newSegment(parts[2])
	if err != nil {
		return nil, fmt.Errorf("invalid format segment: %s", parts[2])
	}

	return &format{major, minor, micro}, nil
}

// CalVer is the type to contain all information regarding current version
type CalVer struct {
	major     string
	minor     string
	micro     string
	increment uint64
	modifier  string
	pre       bool
	format    *format
	date      time.Time
}

// this is for testing purpose only
var now = time.Now

func (c *CalVer) next(pre bool) (string, string, string, uint64) {
	t := now()

	date := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())

	if date.Equal(c.date) {
		if !pre && c.pre {
			return c.major, c.minor, c.micro, c.increment
		}

		return c.major, c.minor, c.micro, c.increment + 1
	}

	c.date = date

	return c.format.major.conv(t), c.format.minor.conv(t), c.format.micro.conv(t), 0
}

// Release generates new release version and returns the string.
// It calculates the next version using time.Now function, and in case of
// multiple releases on the same day, it will bump up the `Iterations` for
// instance:
//		2020.12.12		->	2020.12.12-1
//		2020.12.12-1	->	2020.12.12-2
//		[..]
//		2020.12.12-999	->	2020.12.12-1000
// Furthermore, if the previous version was a prerelease with an iteration
// then it will remove the prerelease modifier and keep the same version
func (c *CalVer) Release() string {
	c.major, c.minor, c.micro, c.increment = c.next(false)

	c.pre = false

	return c.String()
}

// PreRelease generates new prerelease version and returns the string.
// It works same as Release but it suffixes each version with the provided
// `modifier`
func (c *CalVer) PreRelease() string {
	c.major, c.minor, c.micro, c.increment = c.next(true)

	c.pre = true

	return c.String()
}

func (c *CalVer) String() string {
	v := ""

	if c.major == "" && c.minor == "" {
		// in case both `major` and `minor` are empty then it means there hasn't been any release yet
		// so we can just show the format as the version
		v += c.format.String()
	}

	if c.major != "" {
		v += c.major
	}

	if c.minor != "" {
		v += fmt.Sprintf(".%s", c.minor)
	}

	if c.micro != "" {
		v += fmt.Sprintf(".%s", c.micro)
	}

	if c.pre {
		v += fmt.Sprintf("-%s", c.modifier)
	}

	if c.increment > 0 {
		if c.pre {
			v += fmt.Sprintf(".%d", c.increment)
		} else {
			v += fmt.Sprintf("-%d", c.increment)
		}
	}

	return v
}

// New creates a new instance of CalVer using the provided format and modifier
// which defaults to `dev`
func New(format, modifier string) (*CalVer, error) {
	if modifier == "" {
		modifier = "dev"
	}

	f, err := newFormat(format)
	if err != nil {
		return nil, err
	}

	return &CalVer{modifier: modifier, format: f}, nil
}

// Parse takes raw version, tries to parse it into provided format and returns
// the CalVer instance. It takes a modifier as well which default to `dev`
func Parse(raw, format, modifier string) (*CalVer, error) {
	c, err := New(format, modifier)
	if err != nil {
		return nil, err
	}

	parts := strings.Split(raw, "-")
	if len(parts) > 1 {
		// meaning that it could either be an iterative build or a prerelease
		var i string
		if strings.Contains(raw, c.modifier) {
			c.pre = true
			if strings.Contains(parts[1], ".") {
				i = strings.Split(parts[1], ".")[1]
			} else {
				i = "0"
			}
		} else {
			i = parts[1]
		}

		inc, err := strconv.ParseUint(i, 0, 64)
		if err != nil {
			return nil, fmt.Errorf("provided string doesn't match the format: %s", c.format)
		}

		c.increment = inc
	}

	// for now I'm only checking for `.` to verify if the format is valid which barely
	// tells anything so this would be something I need to address later
	count := strings.Count(c.format.String(), ".")
	if strings.Count(parts[0], ".") != count {
		return nil, fmt.Errorf("provided string doesn't match the format: %s", c.format)
	}

	major, minor, micro, err := c.format.parse(parts[0])
	if err != nil {
		return nil, err
	}

	c.major = major
	c.minor = minor
	c.micro = micro

	return c, nil
}
