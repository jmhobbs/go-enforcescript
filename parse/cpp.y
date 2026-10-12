%{
package parse

import "github.com/jmhobbs/go-enforcescript/ast"
%}

%token INTEGER FLOAT BOOL STRING VECTOR VOID 
%token IDENTIFIER
%token TYPE_VOID TYPE_BOOL TYPE_INTEGER TYPE_FLOAT TYPE_STRING TYPE_VECTOR
%token TOK_CLASS TOK_PAREN_OPEN TOK_PAREN_CLOSE TOK_BLOCK_OPEN TOK_BLOCK_CLOSE TOK_SEMICOLON TOK_ASSIGN TOK_QUOTE TOK_COMMA TOK_COLON 
%token COMMENT INLINE_COMMENT

%union{
  identifier string

  variableType uint8

  integerValue int
  floatValue float64
  stringValue string
  boolValue bool
  vectorValue [3]float64

  cOrF ast.ClassOrFunction
  csOrFs []ast.ClassOrFunction

  args []ast.Argument

  block ast.Block
}

%token <identifier> CLASS
%token <identifier> IDENTIFIER

%token <stringValue> STRING COMMENT INLINE_COMMENT
%token <integerValue> INTEGER
%token <floatValue> FLOAT

%type <cOrF> classOrFunction 
%type <csOrFs> classesOrFunctions

%type <cOrF> classDeclaration, functionDeclaration

%type <args> arguments

%type <variableType> type

%type <block> block

%%

file
  : classesOrFunctions {
    file = ast.File{
      Entries: $1,
    }
  }
  ;

classesOrFunctions
  : classesOrFunctions classOrFunction {
    $$ = append($1, $2)
  }
  | classOrFunction {
    $$ = []ast.ClassOrFunction{$1}
  }
  ;

classOrFunction
  : classDeclaration {
    $$ = $1
  }
  | functionDeclaration {
    $$ = $1
  }
  ;

classDeclaration
  : TOK_CLASS IDENTIFIER block {
    $$ = ast.Class{
      Identifier: $2,
      Body: $3,
    }
  }
  ;

functionDeclaration
  : type IDENTIFIER arguments block {
    $$ = ast.Function{
      ReturnType: $1,
      Identifier: $2,
      Arguments: $3,
      Body: $4,
    }
  }
  ;

type
  : TYPE_VOID {
    $$ = ast.TypeVoid
  }
  | TYPE_BOOL {
    $$ = ast.TypeBool
  }
  | TYPE_INTEGER {
    $$ = ast.TypeInteger
  }
  | TYPE_FLOAT {
    $$ = ast.TypeFloat
  }
  | TYPE_STRING {
    $$ = ast.TypeString
  }
  | TYPE_VECTOR {
    $$ = ast.TypeVector
  }
  ;

arguments
  : TOK_PAREN_OPEN TOK_PAREN_CLOSE {
    $$ = []ast.Argument{}
  }
  ;

block
  : TOK_BLOCK_OPEN TOK_BLOCK_CLOSE {
    $$ = ast.Block{
      Statements: []ast.Statement{},
    }
  }
  ;

%%
