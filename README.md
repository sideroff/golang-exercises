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
```go <file_name>```


# Variables

## Declaring
```
    var myVariable int = 24
	myVariable = 42

	i := 52
```
## Redeclaration
## Shadowing
You can redeclare variables in inner scopes but not in the same scope
## Visibility
Scopes
    package level
        lower case - visible only in package
        upper case - visible to other packages
    block
        visible in block only
## Naming conventions
Accronyms are all upper case - 

```var theURL```

instead of 

```var theUrl```
## Type conversion


# Variable types
## Boolean
## Numbers
## Text