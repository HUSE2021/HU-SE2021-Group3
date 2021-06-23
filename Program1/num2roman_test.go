package Program1
import (
"github.com/stretchr/testify/assert"
"testing"
)

type question struct {
	para
	ans
}
// para 是参数
// one 代表第一个参数
type para struct {
	one int
}

// ans 是答案
// one 代表第一个答案
type ans struct {
	one string
}

func Test_Problem0013(t *testing.T) {
	ast := assert.New(t)

	qs := []question{
		question{
			para{1},
			ans{"I"},
		},
		question{
			para{1990},
			ans{"MCMXC"},
		},
		question{
			para{2008},
			ans{"MMVIII"},
		},
		question{
			para{99},
			ans{"XCIX"},
		},
		question{
			para{47},
			ans{"XLVII"},
		},

	}
	for _, q := range qs {
		a, p := q.ans, q.para
		ast.Equal(a.one, num2roman(p.one), "输入:%v", p)
	}
}