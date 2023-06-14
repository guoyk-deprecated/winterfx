package probefx

import "strings"

type Result interface {
	Collect(name string, err error)

	Result() (s string, failed bool)
}

type result struct {
	failed bool
	sb     *strings.Builder
}

func NewResult() Result {
	return &result{sb: &strings.Builder{}}
}

func (c *result) Collect(name string, err error) {
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

func (c *result) Result() (s string, failed bool) {
	if c.sb.Len() == 0 {
		s = "OK"
	} else {
		s = c.sb.String()
	}
	failed = c.failed
	return
}
