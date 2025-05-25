# A Repo to Work On Data Structures and Alorithms in Go
This repository was borne out of working on the data structures and algorithms course on [boot.dev](https://www.boot.dev/).

The original course was in python, and while I found the information to be great and presented well, I found that for myself I needed more practice to really cement the concepts of some of these algorithms.

I was also working on the data structures and algorithms course after I had learnt some GO and I could not find any comprehensive course on data structures and algorithms that were suited for Go, so i thought I would just implement the algorithms from scratch and also write tests against those algorithms, with the idea that if I ever wanted to practice these algorithms again I could use the tests I had written to test my implementation.

Algorithms that have been worked on in this repo this far include:


**The sorting algorithms are in their own sorting directory**
* Bubble Sort
* Insertion Sort
* Merge Sort
* Quick Sort 
* Selection Sort
* Power Sets
* Linked Lists 

Each implementation (apart from the sorting solutions) are in a specific directory thta also comes with a test suite. 

* sorting
    * sorting.go
    * sorting_test.go
* power_sets
    * power_sets.go
    * power_sets_test.go

and so on and so forth.

The main files in each directory have been left blank except for the signatures of the functions and methods, because these will be required to tie up with the tests that have been defined. Where required, some helper methods have been included which can help with debugging.

 For each of the algorithms, the solution file can be found in the relevant folder with the **_solution** suffix.

## Using The Test Files
From within the folder that you are working on , to run all files that have been defined in the specific folder's test file, just run 
```
go test 
```
To run a specific test 
```
go test -run <TestExampleName>
```

For more information regarding working with go's built in test suite, i highly recommend working on [Learn Go With Tests](https://quii.gitbook.io/learn-go-with-tests)

