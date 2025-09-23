1. How do we run a GO project

We need to install go and run the file using `go run "file name"` or build it using go build and then execute the file

2. Some important commands in GO CLI

-> go build: compiles a bunch of go source files
-> go run: compiles and executes one or two files
-> go fmt: format all the code in the current directory
-> go install: compiles and installs a package
-> go get: downloads the raw source code of someone else package
-> go test: runs the test associated with the current project

3. What does package main mean?

A package is a collection of go source files, The first line package main basically means that this particular file belongs to package main.

2 types of packages in go

-> executable: Generates a file that we can run
-> reusable/dependency: Used to write common code or helper functions.

If the first line package main means its a executable type package if its anything other then main package than its a reusable or runnable file.

Anytime we create a executable package it should have a main function.

4. What does import fmt mean?

-> This means that give my package main the access to all the functionality inside fmt. 
-> FMT is part of a standard lib package.
-> It basically means formatting.
-> It is an input output package





