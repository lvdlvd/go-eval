%{

package eval

%}

%union {
	val Evaluator
}

%type	<val>	expr term factor power

%token '+' '-' '*' '/' '(' ')'

%token	<val> CONST VAR

%%

expr:
	term
|	'+' expr
	{
		$$ = $2
	}
|	'-' expr
	{
		$$ = neg{$2}
	}

term:
	factor
|	term '+' factor
	{
		$$ = add{$1, $3}
	}
|	term '-' factor
	{
		$$ = add{$1, neg{$3}}
	}

factor:
	power
|	factor '*' power
	{
		$$ = mul{$1, $3}
	}
|	factor '/' power
	{
		$$ = mul{$1, inv{$3}}
	}

power:
	CONST
|   VAR
|	'(' expr ')'
	{
		$$ = $2
	}


%%
