package text

import (
	"regexp"
	"strings"
	"fmt"
)

const SEP_PHRASE=`\s`

type Text struct {
	t string
	seps string
	tokens []string
}

func (t Text) Init (s string) Text{
	t.t=s
	fmt.Println(s)
	return t
}

func (t Text) Separators(regexString string) Text{
	t.seps=regexString
	fmt.Println(regexString)
	if tNormalized,err:=regexp.Compile(t.seps);
	  err==nil {
		  t.tokens=tNormalized.Split(t.t,-1)
	  } else {
		fmt.Println(err.Error())
	  }
	  fmt.Println(t.tokens)
	return t
}
func (t Text) JustText() string {
	return strings.Join(t.tokens," ")
}
func (t Text) ToLowerSnake() string{
	return strings.Join(t.tokens,"_")
}
func (t Text) ToUpperSnake() string{
	return strings.ToUpper(
		strings.Join(t.tokens,"_"),
	)
}
func (t Text) ToLowerKebap() string{
	return strings.Join(t.tokens,"-")
}
func (t Text) ToUpperKebap() string{
	return strings.ToUpper(
		strings.Join(t.tokens,"-"),
	)
}