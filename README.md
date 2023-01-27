# <span style="color:#C0BFEC">**🦔 Simple calculator**</span>

## <span style="color:#C0BFEC">***Enter to run:*** </span>

```shell
go run calc.go "EXPRESSION"
```

## <span style="color:#C0BFEC">***Usage:***</span>

In place of `EXPRESSION` should be your expression with the following operations:
* `+` - addition
* `-` - subtraction
* `*` - multiplication
* `/` - division

You can also use brackets (`(` and `)`) to prioritize operations.

## <span style="color:#C0BFEC">***Examples:***</span>

```shell
$ go run calc.go "(1+2)-3"
Expression value: 0

$ go run calc.go "(1+2)*3"
Expression value: 9

$ go run calc.go "(((((1)+1)+1)+1)-1)"
Expression value:  3

$ go run calc.go "124-521*255/1*4*12"
Expression value:  -6376916

$ go run calc.go "2197/169"
Expression value:  13

$ go run calc.go "(((((((17)))))))"
Expression value: 17

$ go run calc.go "2/2*2/2*2/2*2/2"
Expression value: 2

$ go run calc.go "2164927*0"
Expression value: 0

```

