//line parser.go.y:2
package filter

import __yyfmt__ "fmt"

//line parser.go.y:2
import "fmt"

type Token struct {
	tok int
	lit interface{}
	pos Position
}

//line parser.go.y:14
type yySymType struct {
	yys        int
	statements []Statement
	statement  Statement
	expr       Expression
	tok        Token
}

const IDENT = 57346
const TRUE = 57347
const FALSE = 57348
const NULL = 57349
const VALUE = 57350
const NUMBER = 57351
const PR = 57352
const EQ = 57353
const NE = 57354
const CO = 57355
const SW = 57356
const EW = 57357
const GT = 57358
const GE = 57359
const LT = 57360
const LE = 57361
const AND = 57362
const OR = 57363
const NOT = 57364
const LPAREN = 57365
const RPAREN = 57366
const LBOXP = 57367
const RBOXP = 57368
const SP = 57369

var yyToknames = []string{
	"IDENT",
	"TRUE",
	"FALSE",
	"NULL",
	"VALUE",
	"NUMBER",
	"PR",
	"EQ",
	"NE",
	"CO",
	"SW",
	"EW",
	"GT",
	"GE",
	"LT",
	"LE",
	"AND",
	"OR",
	"NOT",
	"LPAREN",
	"RPAREN",
	"LBOXP",
	"RBOXP",
	"SP",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line parser.go.y:140

type Lexer struct {
	s          *Scanner
	recentLit  interface{}
	recentPos  Position
	statements []Statement
}

func (l *Lexer) Lex(lval *yySymType) int {
	tok, lit, pos := l.s.Scan()
	if tok == EOF {
		return 0
	}
	lval.tok = Token{tok: tok, lit: lit, pos: pos}
	l.recentLit = lit
	l.recentPos = pos
	return tok
}

func (l *Lexer) Error(e string) {
	panic("parse error, " + fmt.Sprintf("Line %d, Column %d: %q %s",
		l.recentPos.Line, l.recentPos.Column, l.recentLit, e))
}

func Parse(s *Scanner) []Statement {
	l := Lexer{s: s}
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.statements
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 24
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 56

var yyAct = []int{

	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	6, 8, 9, 2, 3, 20, 22, 42, 21, 6,
	8, 9, 23, 24, 43, 9, 8, 9, 5, 4,
	25, 8, 9, 34, 39, 40, 41, 5, 4, 33,
	32, 1, 0, 31, 7, 0, 0, 35, 36, 37,
	38, 27, 28, 29, 30, 26,
}
var yyPact = []int{

	15, -1000, 6, -10, 15, -7, -1000, -1000, 15, 15,
	-1000, 46, 46, 32, 31, 25, 46, 46, 46, 46,
	15, 11, 15, 4, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -9,
	-1000, 0, -1000, -1000,
}
var yyPgo = []int{

	0, 41, 13, 14, 30,
}
var yyR1 = []int{

	0, 1, 1, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 3, 4,
	4, 4, 4, 4,
}
var yyR2 = []int{

	0, 0, 2, 2, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 4, 4, 1, 1,
	1, 1, 1, 1,
}
var yyChk = []int{

	-1000, -1, -2, -3, 23, 22, 4, -1, 20, 21,
	10, 11, 12, 13, 14, 15, 16, 17, 18, 19,
	25, -2, 23, -2, -2, -4, 9, 5, 6, 7,
	8, -4, 8, 8, 8, -4, -4, -4, -4, -2,
	24, -2, 26, 24,
}
var yyDef = []int{

	1, -2, 1, 0, 0, 0, 18, 2, 0, 0,
	3, 0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 13, 14, 4, 19, 20, 21, 22,
	23, 5, 6, 7, 8, 9, 10, 11, 12, 0,
	15, 0, 17, 16,
}
var yyTok1 = []int{

	1,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line parser.go.y:36
		{
			yyVAL.statements = nil
			if l, isLexer := yylex.(*Lexer); isLexer {
				l.statements = yyVAL.statements
			}
		}
	case 2:
		//line parser.go.y:43
		{
			yyVAL.statements = append([]Statement{yyS[yypt-1].statement}, yyS[yypt-0].statements...)
			if l, isLexer := yylex.(*Lexer); isLexer {
				l.statements = yyVAL.statements
			}
		}
	case 3:
		//line parser.go.y:52
		{
			yyVAL.statement = &AttrStatement{Attr: yyS[yypt-1].expr, Operator: yyS[yypt-0].tok.lit.(string)}
		}
	case 4:
		//line parser.go.y:56
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 5:
		//line parser.go.y:60
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 6:
		//line parser.go.y:64
		{
			yyVAL.statement = &RegexStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), Value: yyS[yypt-0].tok.lit}
		}
	case 7:
		//line parser.go.y:68
		{
			yyVAL.statement = &RegexStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), Value: yyS[yypt-0].tok.lit}
		}
	case 8:
		//line parser.go.y:72
		{
			yyVAL.statement = &RegexStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), Value: yyS[yypt-0].tok.lit}
		}
	case 9:
		//line parser.go.y:76
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 10:
		//line parser.go.y:80
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 11:
		//line parser.go.y:84
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 12:
		//line parser.go.y:88
		{
			yyVAL.statement = &CompStatement{LHE: yyS[yypt-2].expr, Operator: yyS[yypt-1].tok.lit.(string), RHE: yyS[yypt-0].expr}
		}
	case 13:
		//line parser.go.y:92
		{
			yyVAL.statement = &LogStatement{LHS: yyS[yypt-2].statement, Operator: yyS[yypt-1].tok.lit.(string), RHS: yyS[yypt-0].statement}
		}
	case 14:
		//line parser.go.y:96
		{
			yyVAL.statement = &LogStatement{LHS: yyS[yypt-2].statement, Operator: yyS[yypt-1].tok.lit.(string), RHS: yyS[yypt-0].statement}
		}
	case 15:
		//line parser.go.y:100
		{
			yyVAL.statement = &ParenStatement{Operator: "", SubStatement: yyS[yypt-1].statement}
		}
	case 16:
		//line parser.go.y:104
		{
			yyVAL.statement = &ParenStatement{Operator: yyS[yypt-3].tok.lit.(string), SubStatement: yyS[yypt-1].statement}
		}
	case 17:
		//line parser.go.y:108
		{
			yyVAL.statement = &GroupStatement{ParentAttr: yyS[yypt-3].expr, SubStatement: yyS[yypt-1].statement}
		}
	case 18:
		//line parser.go.y:114
		{
			yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit.(string)}
		}
	case 19:
		//line parser.go.y:120
		{
			yyVAL.expr = &NumberExpression{Lit: yyS[yypt-0].tok.lit.(int)}
		}
	case 20:
		//line parser.go.y:124
		{
			yyVAL.expr = &BoolExpression{Lit: true}
		}
	case 21:
		//line parser.go.y:128
		{
			yyVAL.expr = &BoolExpression{Lit: false}
		}
	case 22:
		//line parser.go.y:132
		{
			yyVAL.expr = &IdentifierExpression{Lit: yyS[yypt-0].tok.lit.(string)}
		}
	case 23:
		//line parser.go.y:136
		{
			yyVAL.expr = &AttrValueExpression{Lit: yyS[yypt-0].tok.lit.(string)}
		}
	}
	goto yystack /* stack new state and value */
}
