# golang-exercises
A repository holding all the exercises I go through when learning golang

## Adding an exercise
1. Create a directory under src in the format ```<xxx>_<exercise name>.go // ex. 001_hello_world.go``` 
2. Make it a module

## Creating a package
1. ```package <package name> //at the top of the file```

## Creating a module
1. Make it a package
2. ```go mod init <module name>```

## Running a file
```go run <file_name>```


# Variables

## Declaring
```
    var myVariable int = 24
	myVariable = 42

	i := 52
```
## Redeclaration & Shadowing ( local variables overwrite global ones )
You can redeclare variables in inner scopes but not in the same scope
## Visibility
Scopes
    package level
        lower case - visible only in package
        upper case - visible to other packages
    block
        visible in block only
## Naming conventions
camelCase

Accronyms are all upper case - 
```var theURL```

instead of 

```var theUrl```

## Type conversion
Always explicit
```<type>()```

strconv for string conversions since casting will just point to ascii char

## More
All variables must be used ( strict linter )


# Variable types
## Boolean
```var b bool = true```
## Numbers
``` var i int = 12```
### Boolean operators are supported ( and, or, xor )
### e notation is supported
### complex numbers are supported
## String ( utf8[] )
## Runes ( char, 32bit, represented as a int32)

## constants

## arrays
```
var <name> = [<len>]<type>
var a [5]int
b := [5]int{1, 2, 3, 4, 5}
var twoDimensional [2][3]int
```

### Arrays are a value type by default
### Pointers exist 
```
a := 5
b := &a
```

1:58.40