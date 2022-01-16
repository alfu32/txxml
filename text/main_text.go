package text

import (
	"regexp"
	"strings"
)

const SEP_PHRASE=`\s|`

type Text struct {
	t string
	seps string
	tokens []string
}

func (t Text) Init (s string) Text{
	t.t=s
	return t
}

func (t Text) Separators(regexString string) Text{
	t.seps=regexString
	if tNormalized,err:=regexp.Compile(t.seps);
	  err==nil {
		  t.tokens=tNormalized.Split(t.t,-1)
	  }
	return t
}

func (t Text) ToLowerSnake() string{
	return strings.Join(t.tokens,"_")
}