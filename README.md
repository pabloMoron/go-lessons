# Acerca de mi 
[Perfil](https://github.com/pabloMoron/profile)

# Contenido
Codigo con la sintaxis y fundamentos del lenguaje GOLANG, basado en el contenido de este [curso](https://www.udemy.com/course/lenguaje-go) y algunas especificaciones que agregan valor bajo mi punto de vista.

# GOLANG
Go is a general-purpose language designed with systems programming in mind. It is strongly typed and garbage-collected and has explicit support for concurrent programming. Programs are constructed from packages, whose properties allow efficient management of dependencies.

The syntax is compact and simple to parse, allowing for easy analysis by automatic tools such as integrated development environments.

# Notation
The syntax is specified using a [variant](https://en.wikipedia.org/wiki/Wirth_syntax_notation) of Extended Backus-Naur Form (EBNF).

WSN defined in itself:
```
 SYNTAX     = { PRODUCTION } .
 PRODUCTION = IDENTIFIER "=" EXPRESSION "." .
 EXPRESSION = TERM { "|" TERM } .
 TERM       = FACTOR { FACTOR } .
 FACTOR     = IDENTIFIER
            | LITERAL
            | "[" EXPRESSION "]"
            | "(" EXPRESSION ")"
            | "{" EXPRESSION "}" .
 IDENTIFIER = letter { letter } .
 LITERAL    = """" character { character } """" .
```

 ## GO syntax:
```
Syntax      = { Production } .
Production  = production_name "=" [ Expression ] "." .
Expression  = Term { "|" Term } .
Term        = Factor { Factor } .
Factor      = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```
```
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

## Characters
```
newline        = /* the Unicode code point U+000A */ .
unicode_char   = /* an arbitrary Unicode code point except newline */ .
unicode_letter = /* a Unicode code point categorized as "Letter" */ .
unicode_digit  = /* a Unicode code point categorized as "Number, decimal digit" */ .
```

## Letters and digits

The underscore character _ (U+005F) is considered a lowercase letter.

```
letter        = unicode_letter | "_" .
decimal_digit = "0" … "9" .
binary_digit  = "0" | "1" .
octal_digit   = "0" … "7" .
hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
```

## Identifiers
Identifiers name program entities such as variables and types. An identifier is a sequence of one or more letters and digits. The first character in an identifier must be a letter.

```
identifier = letter { letter | unicode_digit } .

eg:

- foo1      👍
- foo       👍
- _         👍
- _foo2     👍

- 1         👎
- 1_foo     👎
- 1foo      👎
- +         👎
```

## Keywords
The following keywords are reserved and may not be used as identifiers.

```
break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var
```

## Operators and punctuation
The following character sequences represent operators (including assignment operators) and punctuation:


```
+    &     +=    &=     &&    ==    !=    (    )
-    |     -=    |=     ||    <     <=    [    ]
*    ^     *=    ^=     <-    >     >=    {    }
/    <<    /=    <<=    ++    =     :=    ,    ;
%    >>    %=    >>=    --    !     ...   .    :
     &^          &^=          ~
```

## Operators
Operators combine operands into expressions.

```
Expression = UnaryExpr | Expression binary_op Expression .
UnaryExpr  = PrimaryExpr | unary_op UnaryExpr .

binary_op  = "||" | "&&" | rel_op | add_op | mul_op .
rel_op     = "==" | "!=" | "<" | "<=" | ">" | ">=" .
add_op     = "+" | "-" | "|" | "^" .
mul_op     = "*" | "/" | "%" | "<<" | ">>" | "&" | "&^" .

unary_op   = "+" | "-" | "!" | "^" | "*" | "&" | "<-" .
```
 
[Full Specs](https://go.dev/ref/spec)