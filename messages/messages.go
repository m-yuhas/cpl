/*
messages.go - Entry Point for Code Execution
(C) 2017 Michael Yuhas
*/

/*
The messages package stores, warinings, info, and other user facing text
*/
package messages

const CLIHeaderText string = "欢迎中华电脑语言第0.2版本!\n©2017 － 迈克尔 余哈斯"
const CLIInputError string = "错误：输入错误发生了、请查下你刚输入的句"

const CommentEndWithoutStart string = "错误：结束解没有链接的‘解：’"
const LineEndsWithPoundSign string = "错误：一句不能被#终结"

const InvalidFunctionDeclaration string = "错误：不当函数声明"
const EndFunctionNotFound string = "错误：找不到对应的结束函数"
const FunctionWithinFunction string = "错误：不能创造函数在另外一个函数内面"
const DuplicateName string = "错误：名号已经用的"

const InvalidClassDeclaration string = "错误：不当类声明"
const ClassWithinClass string = "错误：不能创造类在另外一个类内面"
const EndClassNotFound string = "错误：找不到对应的结束类"

const OutputCommandSyntaxError string = "错误：输出命令需要括弧"
