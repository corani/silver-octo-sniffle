# Oberon 07 EBNF

```ebnf
ident = letter {letter | digit}.
letter = "A" | "B" | ... | "Z" | "a" | "b" | ... | "z".
digit = "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9".
number = integer | real.
integer = digit {digit} | digit {hexDigit} "H".
real = digit {digit} "." {digit} [ScaleFactor].
ScaleFactor = "E" ["+" | "-"] digit {digit}.
hexDigit = digit | "A" | "B" | "C" | "D" | "E" | "F".
string = """ {character} """ | digit {hexDigit} "X".

module = MODULE ident ";" [ImportList] DeclarationSequence
    [BEGIN StatementSequence] END ident ".".
ImportList = IMPORT import {"," import} ";".
import = ident [":=" ident].
DeclarationSequence = [CONST {ConstDeclaration ";"}]
    [TYPE {TypeDeclaration ";"}] [VAR {VariableDeclaration ";"}]
    {ProcedureDeclaration ";"}.
ConstDeclaration = identdef "=" ConstExpression.
ConstExpression = expression.
TypeDeclaration = identdef "=" StrucType.
StrucType = ArrayType | RecordType | PointerType | ProcedureType.
ArrayType = ARRAY length {"," length} OF type.
length = ConstExpression.
RecordType = RECORD ["(" BaseType ")"] [FieldListSequence] END.
BaseType = qualident.
FieldListSequence = FieldList {";" FieldList}.
FieldList = IdentList ":" type.
IdentList = identdef {"," identdef}.
PointerType = POINTER TO type.
ProcedureType = PROCEDURE [FormalParameters].
FormalParameters = "(" [FPSection {";" FPSection}] ")" [":" qualident].
FPSection = [VAR] ident {"," ident} ":" FormalType.
FormalType = {ARRAY OF} qualident.
qualident = [ident "."] ident.
identdef = ident ["*"].
VariableDeclaration = IdentList ":" type.
type = qualident | StrucType.
ProcedureDeclaration = ProcedureHeading ";" ProcedureBody ident.
ProcedureHeading = PROCEDURE identdef [FormalParameters].
ProcedureBody = DeclarationSequence [BEGIN StatementSequence]
    [RETURN expression] END.
expression = SimpleExpression [relation SimpleExpression].
relation = "=" | "#" | "<" | "<=" | ">" | ">=" | IN | IS.
SimpleExpression = ["+" | "-"] term {AddOperator term}.
AddOperator = "+" | "-" | OR.
term = factor {MulOperator factor}.
MulOperator = "*" | "/" | DIV | MOD | "&".
factor = number | string | NIL | TRUE | FALSE |
    set | designator [ActualParameters] | "(" expression ")" | "~" factor.
designator = qualident {selector}.
selector = "." ident | "[" ExpList "]" | "^" | "(" qualident ")".
set = "{" [element {"," element}] "}".
element = expression [".." expression].
ExpList = expression {"," expression}.
ActualParameters = "(" [ExpList] ")".
statement = [assignment | ProcedureCall | IfStatement | CaseStatement |
    WhileStatement | RepeatStatement | ForStatement].
assignment = designator ":=" expression.
ProcedureCall = designator [ActualParameters].
StatementSequence = statement {";" statement}.
IfStatement = IF expression THEN StatementSequence
    {ELSIF expression THEN StatementSequence}
    [ELSE StatementSequence] END.
CaseStatement = CASE expression OF case {"|" case} END.
case = [CaseLabelList ":" StatementSequence].
CaseLabelList = LabelRange {"," LabelRange}.
LabelRange = label [".." label].
label = integer | string | qualident.
WhileStatement = WHILE expression DO StatementSequence
    {ELSIF expression DO StatementSequence} END.
RepeatStatement = REPEAT StatementSequence UNTIL expression.
ForStatement = FOR ident ":=" expression TO expression [BY ConstExpression]
    DO StatementSequence END.
```

Source: http://oberon07.com/o7EBNF.xhtml
