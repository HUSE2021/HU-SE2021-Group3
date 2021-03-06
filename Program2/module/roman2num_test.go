package mymain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type question struct {
	para
	ans
}


type para struct {
	one string
}

type ans struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{"MCMXC"},
			ans{1990},
		},
		question{
			para{"MMVIII"},
			ans{2008},
		},
		question{
			para{"XCIX"},
			ans{99},
		},
		question{
			para{"XLVII"},
			ans{47},
		},

	}

	for _, q := range qs {
		a, p := q.ans, q.para

		ast.Equal(a.one, romanToInt(p.one), "输入:%v", p)
	}
}
