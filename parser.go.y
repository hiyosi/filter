%{
package filter

import "fmt"

type Token struct {
	tok int
	lit interface{}
	pos Position
}

%}

%union{
	statements []Statement
	statement  Statement
	expr	   Expression
	tok	   Token
}

%type<statements> statements
%type<statement> statement
%type<expr> attrName
%type<expr> attrValue

%token<tok> IDENT TRUE FALSE NULL VALUE NUMBER PR EQ NE CO SW EW GT GE LT LE AND OR NOT LPAREN RPAREN LBOXP RBOXP SP

%left AND
%left OR
%right NOT

%%

statements
	:
	{
		$$ = nil
		if l, isLexer := yylex.(*Lexer); isLexer {
			l.statements = $$
		}
	}
	| statement statements
	{
		$$ = append([]Statement{$1}, $2...)
		if l, isLexer := yylex.(*Lexer); isLexer {
			l.statements = $$
		}
	}

statement //TODO: gt,ge,lt,le operators should take string and boolean
        : attrName PR
        {
                $$ = &AttrStatement{Attr: $1, Operator: $2.lit.(string) }
        }
        | attrName EQ attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | attrName NE attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | attrName CO VALUE
        {
                $$ = &RegexStatement{LHE: $1, Operator: $2.lit.(string), Value: $3.lit}
        }
        | attrName SW VALUE
        {
                $$ = &RegexStatement{LHE: $1, Operator: $2.lit.(string), Value: $3.lit}
        }
        | attrName EW VALUE
        {
                $$ = &RegexStatement{LHE: $1, Operator: $2.lit.(string), Value: $3.lit}
        }
        | attrName GT attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | attrName GE attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | attrName LT attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | attrName LE attrValue
        {
                $$ = &CompStatement{LHE: $1, Operator: $2.lit.(string), RHE: $3}
        }
        | statement AND statement
        {
                $$ = &LogStatement{LHS: $1, Operator: $2.lit.(string), RHS: $3}
        }
        | statement OR statement
        {
                $$ = &LogStatement{LHS: $1, Operator: $2.lit.(string), RHS: $3}        
        }
        | LPAREN statement RPAREN
        {
                $$ = &ParenStatement{Operator: "", SubStatement: $2}
        }
        | NOT LPAREN statement RPAREN
        {
                $$ = &ParenStatement{Operator: $1.lit.(string), SubStatement: $3}
        }
        | attrName LBOXP statement RBOXP
        {
                $$ = &GroupStatement{ParentAttr: $1, SubStatement: $3}
        }

attrName
        : IDENT 
	{
		$$ = &IdentifierExpression{Lit: $1.lit.(string)}
        }

attrValue
        : NUMBER
	{
		$$ = &NumberExpression{Lit: $1.lit.(int)}
	}
        | TRUE
        {
       		$$ = &BoolExpression{Lit: true}  
        }
        | FALSE
        {
       		$$ = &BoolExpression{Lit: false}
        }
        | NULL
        {
       		$$ = &IdentifierExpression{Lit: $1.lit.(string)}
        }        
        | VALUE
	{
		$$ = &AttrValueExpression{Lit: $1.lit.(string)}
        }        

%%

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
