
# Chapter: Introduction to `godog` (1 hour)

## 📌 Chapter Objectives

By the end of this chapter, you will be able to:

- Understand what `godog` is and how it relates to BDD and Cucumber
- Set up and install `godog` in a Go project
- Create a basic directory structure for a BDD-driven Go application
- Write and execute your first Gherkin-based feature file with `godog` using Go's `go test` command

---

## 🧠 1. What is `godog`?

`godog` is a Behavior-Driven Development (BDD) framework for Go that works similarly to **Cucumber**. It allows you to write your software’s expected behavior in **plain English** using the [Gherkin syntax](https://cucumber.io/docs/gherkin/), and then **map those steps to actual Go functions** that implement the behavior.

### 🔄 Analogy (Real-world example)

> **Imagine you're ordering pizza over the phone.**
> 
> - You say: "I want a large Margherita with extra cheese."
> - The pizza shop interprets this plain sentence and executes a process internally to deliver your pizza.

BDD works in the same way:
- **Feature file** = Your plain-language order (written in Gherkin)
- **Step definitions** = The kitchen (Go functions)
- **godog** = The waiter who reads the order and ensures it gets prepared

---

## 💡 Why `godog`?

- Helps **non-technical stakeholders** understand software behavior
- Encourages **collaboration** between devs, testers, and product owners
- Enables **automated testing** based on business-readable specifications

---

## 🔧 2. Installing `godog`

Make sure you have Go 1.19+ installed.

### 🛠 Installation command

```bash
go install github.com/cucumber/godog/cmd/godog@latest
```

This will install `godog` binary in your `$GOPATH/bin`.

### ✅ Verify Installation

```bash
godog --version
```

You should see output like:

```
Godog version: v0.15.0
```

---

## 📁 3. Basic Project Structure

Here’s a minimal example of a godog-based BDD project using `go test`:

```
.
├── go.mod
├── godog_test.go
├── features/
│   └── addition.feature
└── stepdefs/
    └── addition_steps.go
```

### 🔍 Description of Files

- `go.mod` — Your Go module
- `godog_test.go` — Test entry point using `*testing.T`
- `features/` — Directory to hold `.feature` files written in Gherkin
- `stepdefs/` — Directory for Go code implementing step definitions

---

## 🧪 4. Hands-On Example

### ✅ Scenario

Let's write a test scenario for a simple calculator app that adds two numbers.

#### 📄 `features/addition.feature`

```gherkin
Feature: Calculator

  Scenario: Adding two numbers
    Given I have entered 50 into the calculator
    And I have entered 70 into the calculator
    When I press add
    Then the result should be 120 on the screen
```

#### 🧩 `stepdefs/addition_steps.go`

```go
package stepdefs

import (
    "fmt"
    "github.com/cucumber/godog"
)

type calculator struct {
    inputs []int
    result int
}

func (c *calculator) iHaveEnteredIntoTheCalculator(arg1 int) error {
    c.inputs = append(c.inputs, arg1)
    return nil
}

func (c *calculator) iPressAdd() error {
    c.result = 0
    for _, v := range c.inputs {
        c.result += v
    }
    return nil
}

func (c *calculator) theResultShouldBeOnTheScreen(expected int) error {
    if c.result != expected {
        return fmt.Errorf("expected %d, but got %d", expected, c.result)
    }
    return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
    calc := &calculator{}

    ctx.Step(`^I have entered (\d+) into the calculator$`, calc.iHaveEnteredIntoTheCalculator)
    ctx.Step(`^I press add$`, calc.iPressAdd)
    ctx.Step(`^the result should be (\d+) on the screen$`, calc.theResultShouldBeOnTheScreen)
}
```

---

#### 🧪 `godog_test.go`

```go
package main

import (
    "os"
    "testing"

    "github.com/cucumber/godog"
    "github.com/cucumber/godog/colors"

    "your_module_path/stepdefs"
)

func TestFeatures(t *testing.T) {
    opts := godog.Options{
        Format:        "pretty",
        Paths:         []string{"features"},
        Output:        colors.Colored(os.Stdout),
        Strict:        true,
        TestingT:      t,
        StopOnFailure: true,
    }

    suite := godog.TestSuite{
        ScenarioInitializer: stepdefs.InitializeScenario,
        Options:             &opts,
    }

    if suite.Run() != 0 {
        t.Fatal("test suite failed")
    }
}
```

> 🔁 Replace `your_module_path/stepdefs` with the correct import path from your `go.mod`.

---

### ▶️ Running the test

```bash
go test -v
```

Expected output:

```
=== RUN   TestFeatures
Feature: Calculator
=== RUN   TestFeatures/Adding_two_numbers
...
PASS
```

---

## 🎯 Interview Questions

1. What is the difference between TDD and BDD?
2. What are the benefits of using `godog` with Go for BDD testing?
3. What is the role of `.feature` files in the BDD workflow?
4. How does `godog` connect a Gherkin step to a Go function?
5. Why is `godog run` deprecated, and what is the new recommended method?

---

## 📺 Curated YouTube Videos

1. [BDD with Go using Godog (Intro) - Kevin Gillette](https://www.youtube.com/watch?v=aL5G_AiRIzA)
2. [Gherkin and Feature Files for Beginners](https://www.youtube.com/watch?v=cr5ZqQeY55k)
3. [Test Your Code Using Godog (DevSecCon)](https://www.youtube.com/watch?v=46tFrpI0TSM)
4. [Intro to Cucumber BDD Testing](https://www.youtube.com/watch?v=5SbtXnD1Lx0)

---

## 📚 Additional Resources

- Godog GitHub: [https://github.com/cucumber/godog](https://github.com/cucumber/godog)
- Gherkin Reference: [https://cucumber.io/docs/gherkin/](https://cucumber.io/docs/gherkin/)
- Go Modules Guide: [https://go.dev/doc/modules](https://go.dev/doc/modules)

---

## ✅ Summary

- `godog` enables writing automated acceptance tests in plain English
- It bridges the gap between non-tech stakeholders and developers
- Use `go test` instead of `godog run` for long-term compatibility

> In the next chapter, we'll focus on **writing more complex scenarios**, **scenario outlines**, and **integration with CI tools like GitHub Actions**.
