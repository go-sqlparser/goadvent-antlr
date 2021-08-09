# goadvent-antlr

Updated examples from [goadvent blog post](https://blog.gopheracademy.com/advent-2017/parsing-with-antlr4-and-go/)

The `parser` folder has already been updated by:

```shell
wget http://www.antlr.org/download/antlr-4.9.2-complete.jar
alias antlr='java -jar antlr-4.9.2-complete.jar'

antlr -Dlanguage=Go -o parser Calc.g4
```
