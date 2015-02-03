package filter

type (
	Statement interface {
		statement()
	}

	Expression interface {
		expression()
	}
)

type (
	ExpressionStatement struct {
		Expr Expression
	}

	AttrStatement struct {
		Attr     Expression
		Operator string
	}

	CompStatement struct {
		LHE      Expression
		Operator string
		RHE      Expression
	}

	RegexStatement struct {
		LHE      Expression
		Operator string
		Value    interface{}
	}

	ParenStatement struct {
		Operator     string
		SubStatement Statement
	}

	LogStatement struct {
		LHS      Statement
		Operator string
		RHS      Statement
	}

	GroupStatement struct {
		ParentAttr   Expression
		SubStatement Statement
	}
)

func (x *ExpressionStatement) statement() {}
func (x *AttrStatement) statement()       {}
func (x *CompStatement) statement()       {}
func (x *RegexStatement) statement()      {}
func (x *ParenStatement) statement()      {}
func (x *LogStatement) statement()        {}
func (x *GroupStatement) statement()      {}

type (
	NumberExpression struct {
		Lit int
	}

	IdentifierExpression struct {
		Lit string
	}

	AttrValueExpression struct {
		Lit string
	}

	BoolExpression struct {
		Lit bool
	}
)

func (x *NumberExpression) expression()     {}
func (x *IdentifierExpression) expression() {}
func (x *AttrValueExpression) expression()  {}
func (x *BoolExpression) expression()       {}
