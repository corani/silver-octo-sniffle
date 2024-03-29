# TODO

The goal is to implement a working Oberon compiler. The checklist below will grow as I go along.

- [x] (temporary) built-in `print`.
- [x] support simple integers.
- [x] print tokens, AST and IR.
- [x] "golden tests" runner.
- [x] support binary `+` / `-`.
- [x] support unary `+` / `-`.
- [x] support binary `*` / `/`.
- [x] support grouping `(...)`.
- [x] wrap statements in dummy module.
- [x] separate statements and expressions.
- [x] basic type checking.
- [x] support `REAL` outcome of divisions.
- [x] replace `print` statement with call expression (still built-in).
- [x] run "golden tests" as unit tests.
- [x] support integer `DIV` / `MOD`.
- [x] skip comments in lexer.
- [x] support string literals.
- [x] support boolean literals `TRUE` / `FALSE`.
- [x] support logical "not" `~`.
- [x] support relational operators `=` / `#` / `<` / `<=` / `>=` / `>`.
- [x] support REAL literals.
- [x] support hex integer literals.
- [x] support hex string literals.
- [x] support scale factor in REAL literals.
- [x] better printing of boolean values.
- [x] support statement sequences (separated by `;`).
- [x] get "golden tests" input from the code block in the markdown file.
- [x] remove dedicated "golden test" runner and rely on unit-tests.
- [x] reorganize "golden tests" to combine multiple tiny tests.
- [x] make statement sequence a separate node.
- [x] support basic `MODULE` structure (excluding imports and declarations) .
- [x] don't panic on compilation errors.
- [x] support `IF` / `THEN` / `ELSIF` / `ELSE` with constant expressions.
- [x] basic global variable and assignment support.
- [x] support global constant literals.
- [x] support global constant expressions.
- [x] require a `MODULE` instead of allowing naked statements.
- [X] support built-in `INC` and `DEC` procedures. 
- [X] support built-in `FLT` and `FLOOR` procedures.
- [X] add designator to CallExpr.
- [X] no automatic type conversion.
- [X] support `REPEAT` / `UNTIL` statement.
- [X] support `WHILE` / `DO` / `ELSIF` statement.
- [X] support `FOR` / `BY` / `DO` statement.
- [x] support built-in `ORD` procedure for booleans. 
- [X] support `CHAR` types. 
- [X] support built-in `ORD` and `CHR` procedures.
- [x] support relational operators for `CHAR` types.
- [X] support `SET` types.
- [X] support built-in `INCL` and `EXCL` procedures. 
- [x] verify `SET` operations `+` / `-` / `*` / `/`.
- [x] improve test coverage.
- [x] factor out builtin functions.
- [x] use visitor pattern for builtin functions.
- [x] split code into packages.
- [x] better error reporting.
- [x] continue parsing if errors are found.
- [x] combine constants and vars into symbol-table.
- [x] add types to symbol-table.
- [x] use a type node instead of a token.
- [x] solidify types during type checking.
- [x] support `POINTER` types.
- [x] support built-in `NEW` procedure.
- [x] do we need to distinguish lhs / rhs designators?
- [x] support pointer dereferences.
- [x] support temporary built-in `DELETE` procedure.
- [x] support qualified-identifiers.
- [x] replace `print` with builtin `Texts` procedures.
- [x] move generated code to `oberonMain` to allow for preamble.
- [x] capture command line arguments.
- [x] capture environment variables.
- [ ] basic local variable and assignment support.
- [ ] generate call-graph in markdown.
- [ ] support multiple modules in checker (and generator?).
- [ ] better sync points during parser error recovery.
- [ ] add "golden tests" for error conditions (e.g. invalid code).
- [ ] improve AST and type information for variables.
- [ ] support procedure definitions.
- [ ] add paramDecl to support by-value vs by-reference parameters.
- [ ] add frames in checker (and generator?).
- [ ] do we need a new AST for checked nodes?
- [ ] support local procedure calls (i.e. in the same module).
- [ ] support `ARRAY` types.
- [ ] support built-in `LEN` procedure.
- [ ] support `RECORD` types.
- [ ] support `PROCEDURE` types.
- [ ] support `CASE` / `OF` statement.
- [ ] make builtin `Texts` procedures real procedures.
- [ ] better command line arguments.
- [ ] update "golden test" files in-place, so any additional content doesn't get overwritten.
- [ ] support reference counting or garbage collection.
- [ ] a magic "C" call built-in module?
- [ ] remove builtin `print` procedure. 
- [ ] remove builtin `DELETE` procedure.
- [ ] a standard library.
- [ ] ...
