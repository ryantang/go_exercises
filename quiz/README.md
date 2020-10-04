# Exercise #1: Quiz Game

## Usage

This program allows the user to run small quiz on the command line. To build and run this program, run the following commands:

```
$ go build quiz.go #compiles the quiz binary
$ ./quiz # this will start he quiz
```

The user can also create their own quiz questions and answers in CSV format. See "problems.csv" and "pokemon.csv" as examples. To choose different quiz questions and change the timer setting, use the command line flags. 

```
$  ./quiz -h
Usage of ./quiz:
  -f string
    	Quiz CSV file in 'question,answer' format (default "problems.csv")
  -s string
    	Number of seconds alloted in this quiz (default "30")
```

## Exercise Assignment

This assignment, from [Gophercises](github.com/gophercises/quiz), provides a nice introduction to Go for a programmer with experience in a different language.

### Part 1 (Implemented)

Create a program that will read in a quiz provided via a CSV file (more details below) and will then give the quiz to a user keeping track of how many questions they get right and how many they get incorrect. Regardless of whether the answer is correct or wrong the next question should be asked immediately afterwards.

The CSV file should default to `problems.csv` (example shown below), but the user should be able to customize the filename via a flag.

The CSV file will be in a format like below, where the first column is a question and the second column in the same row is the answer to that question.

```
5+5,10
7+3,10
1+1,2
8+3,11
1+2,3
8+6,14
3+1,4
1+4,5
5+1,6
2+3,5
3+3,6
2+4,6
5+2,7
```

You can assume that quizzes will be relatively short (< 100 questions) and will have single word/number answers.

At the end of the quiz the program should output the total number of questions correct and how many questions there were in total. Questions given invalid answers are considered incorrect.

**NOTE:** *CSV files may have questions with commas in them. Eg: `"what 2+2, sir?",4` is a valid row in a CSV. I suggest you look into the CSV package in Go and don't try to write your own CSV parser.*

### Part 2 (Implemented)

Adapt your program from part 1 to add a timer. The default time limit should be 30 seconds, but should also be customizable via a flag.

Your quiz should stop as soon as the time limit has exceeded. That is, you shouldn't wait for the user to answer one final questions but should ideally stop the quiz entirely even if you are currently waiting on an answer from the end user.

Users should be asked to press enter (or some other key) before the timer starts, and then the questions should be printed out to the screen one at a time until the user provides an answer. Regardless of whether the answer is correct or wrong the next question should be asked.

At the end of the quiz the program should still output the total number of questions correct and how many questions there were in total. Questions given invalid answers or unanswered are considered incorrect.

### Bonus (Not Implemented)

As a bonus exercises you can also...

1. Add string trimming and cleanup to help ensure that correct answers with extra whitespace, capitalization, etc are not considered incorrect. *Hint: Check out the [strings](https://golang.org/pkg/strings/) package.*
2. Add an option (a new flag) to shuffle the quiz order each time it is run.

## Reflections

Things that I learned from this exercise.

### In this exercise I got to work with

* Flags
* Strings
* Opening Files
* Go's CSV Package
* A simple usage of goroutines and channels
* Timer
* Pointers (I mainly coded in Ruby before and didn't use pointers much.)

### Flags

I was a bit disappointed that the flags library didn't seem to offer both the short form (e.g. `-f filename`) and the long form (e.g. `--file filename`) option. It seems like you can use any option you want, but the flags only use a single dash (e.g. `-f` or `-file`).

I'd also forget to call `flag.Parse()` and wonder why it wasn't working.

### Goroutines, functions, and channels

Concurrency in Go is as simple as running `go <function>`, which immediately kicks off the function in parallel. Examples on the Internet show functions inline, which looks something like this:

```
func main(){
  //a bunch of code here
  go func(waitTime time.Duration, correctAnswers *int, totalQuestions int) {
    timer := time.NewTimer(waitTime)

    <-timer.C
    fmt.Println("\nRan out of time. Corect answers: ", *correctAnswers, "/", totalQuestions)
    os.Exit(0)
  }
  // the rest of main's code
}
```

After a bunch of reading about functions, I learned that

1. This is called an "anonymous function", because the function is not named.
1. Functions are a first class value in Go. In this case (i.e. `go func(waitTime ...) {...}`, we are passing a function (i.e. `func (waitTime...) {...}` to another function (i.e. `go`)
1. The syntax `<-timer.C` creates a blocking channel. This means that the rest of this goroutine waits for the timer to finish before moving on to the `fmt.Println()` statement.

I thought this inlined anonymous function made my main function look bloated. So, I pulled out a separate (named) function:

```
func main(){
// a bunch of code
	go runTimer(seconds, &correctAnswers, totalQuestions)
// the rest of the code
}

func runTimer(waitTime time.Duration, correctAnswers *int, totalQuestions int) {
	timer := time.NewTimer(waitTime)

	<-timer.C
	fmt.Println("\nRan out of time. Corect answers: ", *correctAnswers, "/", totalQuestions)
	os.Exit(0)
}
```

From a readability standpoint, I prefer the named function (`runTimer`) as opposed to the anonymous function we had above.

### Working with files

Go has a concept called `io.Reader` that I'm still trying to understand. Documentation says that it's an interface that has a `Read` method. 

The relevant thing is that the CSV package takes input in the type of `io.Reader`. My early implementation had me reading in data from the file using `ioutil.Readfile(csvfile)`, and then converting it back into `io.Reader` using `strings.NewReader(string(data))`. When I learned that `os.Open(csvfile)` returned an `io.Reader`, I was able to save a bunch of unnecessary steps.

As I tried to wrap my head around this concept of an `io.Reader`, my coworker described it as "a thin wrapper around a file descriptor". I thought that was interesting.
