package winterfx

import "strings"

type CheckResult interface {
	Collect(name string, err error)
	Result() (s string, failed bool)
}

type checkResult struct {
	failed bool
	sb     *strings.Builder
}

func NewCheckResult() CheckResult {
	return &checkResult{sb: &strings.Builder{}}
}

func (c *checkResult) Collect(name string, err error) {
	if c.sb.Len() > 0 {
		c.sb.WriteString("\n")
	}
	c.sb.WriteString(name)
	if err == nil {
		c.sb.WriteString(": OK")
	} else {
		c.failed = true
		c.sb.WriteString(": ")
		c.sb.WriteString(err.Error())
	}
}

func (c *checkResult) Result() (s string, failed bool) {
	if c.sb.Len() == 0 {
		s = "OK"
	} else {
		s = c.sb.String()
	}
	failed = c.failed
	return
}
