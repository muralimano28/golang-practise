## quick tips
1. Go lang doesnt allow to declare variables that are not used. But this restriction is only function level not package level. It means you can declare a variable at package level and dont have to use it.
2. Go lang follows strict data typing. So int variables cant be used to calculate float32 output. It will throw an error.
3. When creating our own package, those packages are not compiled and stored in bin directory. Instead while importing those packages, only necessary things are compiled with the code where we imported custom package and that binary comes inside bin directory.
4. "go install" command compiles only main package.
5. functions starting with small letters are local functions and functions starting with capital letters are exported functions. So, while importing a package, you can use only exported functions.
6. you can create a function with name "Main" because that's not same as function main. And function main in package main is the entry point for go code.
