# Chapter: Writing Step Definitions (1 hour)

## 📌 Chapter Objectives

By the end of this chapter, you will be able to:

- Write BDD step definition functions in Go
- Use regular expressions to match Gherkin steps
- Understand the role of `godog`'s `*ScenarioContext`
- Handle structured data inputs using **Gherkin Data Tables**

---

## 🧠 Theoretical Concepts

### 🧩 What are Step Definitions?

Step definitions are Go functions that bind your natural language steps (written in `.feature` files using Gherkin) to executable test logic.

They allow you to go from:
```gherkin
Given I enter username "john"
```
To executing:
```go
func (s *Suite) iEnterUsername(username string) error {
    // store or verify "john"
}
```

---

### 🎯 Gherkin Steps + Regex + Go Functions = BDD Automation

Every Gherkin step is matched to a Go function using a regular expression:

```gherkin
Given I enter username "john"
```

is matched to:

```go
ctx.Step(`^I enter username "([^"]*)"$`, s.iEnterUsername)
```

---

## 🏗 Real-World Analogy

> Think of a voice assistant like Alexa.  
> You say: “Turn on the bedroom light.”  
> It matches this to a known voice command and runs the actual logic to toggle the switch.

The same applies here:
- Gherkin step = voice command
- Regex = command parser
- Go function = logic to execute the command

---

## 🛠️ Hands-On Examples

Let’s walk through progressively advanced examples.

---

### ✅ Basic Step Definition

#### Gherkin
```gherkin
When I enter username "john"
```

#### Go
```go
func (s *Suite) iEnterUsername(username string) error {
    s.enteredUsername = username
    return nil
}
```

#### Registration
```go
ctx.Step(`^I enter username "([^"]*)"$`, s.iEnterUsername)
```

📌 **Explanation**:
- The `([^"]*)` captures `"john"` and passes it as a `string` parameter.
- The method is defined on a struct `Suite` that holds shared state across steps.

---

### 🧠 Step Matching with Multiple Parameters

#### Gherkin
```gherkin
And I enter password "secret" for username "john"
```

#### Go
```go
func (s *Suite) iEnterPasswordForUsername(password, username string) error {
    s.credentials[username] = password
    return nil
}
```

#### Regex
```go
ctx.Step(`^I enter password "([^"]*)" for username "([^"]*)"$`, s.iEnterPasswordForUsername)
```

---

### 📚 Using ScenarioContext

In your step registration block:

```go
func InitializeScenario(ctx *godog.ScenarioContext) {
    s := &Suite{
        credentials: make(map[string]string),
    }

    ctx.Step(`^I enter username "([^"]*)"$`, s.iEnterUsername)
    ctx.Step(`^I enter password "([^"]*)" for username "([^"]*)"$`, s.iEnterPasswordForUsername)
}
```

The `ScenarioContext` ensures that each test scenario gets a clean state.

---

## 📋 Using Gherkin Data Tables

Gherkin lets you define structured inputs using tables:

### Gherkin Table Example
```gherkin
Given the following users exist:
  | username | password |
  | alice    | 12345    |
  | bob      | pass123  |
```

### Go Step Function
```go
func (s *Suite) theFollowingUsersExist(table *godog.Table) error {
    for _, row := range table.Rows[1:] { // skip header
        username := row.Cells[0].Value
        password := row.Cells[1].Value
        s.credentials[username] = password
    }
    return nil
}
```

### Registration
```go
ctx.Step(`^the following users exist:$`, s.theFollowingUsersExist)
```

---

## 🚀 Advanced: Using Structs for Table Parsing

```go
type User struct {
    Username string
    Password string
}

func (s *Suite) theUsersExist(table *godog.Table) error {
    headers := table.Rows[0].Cells
    for _, row := range table.Rows[1:] {
        u := User{}
        for i, cell := range row.Cells {
            switch headers[i].Value {
            case "username":
                u.Username = cell.Value
            case "password":
                u.Password = cell.Value
            }
        }
        s.credentials[u.Username] = u.Password
    }
    return nil
}
```

---

## 📑 Summary Checklist

| Concept                   | Key Point |
|--------------------------|-----------|
| Step Definition          | Go function tied to a Gherkin step |
| Regex Matching           | Extracts step values as parameters |
| ScenarioContext          | Registers and resets state per scenario |
| Data Tables              | Used to pass structured inputs to a step |

---

## 🎯 Interview Questions

1. What are step definitions in BDD and how are they used in Go?
2. How does regex matching work in godog?
3. What is `ScenarioContext` and why is it important?
4. How do you handle tabular data in Gherkin and map it to Go code?
5. Can a single step definition be reused across multiple features?

---

## 📺 Curated YouTube Videos

1. [Using godog and Go for BDD Testing](https://www.youtube.com/watch?v=46tFrpI0TSM)
2. [Gherkin Data Tables Tutorial](https://www.youtube.com/watch?v=pXB4Lbyf4mg)
3. [BDD with Cucumber and Structuring Step Definitions](https://www.youtube.com/watch?v=5SbtXnD1Lx0)
4. [Regex Crash Course for Testers](https://www.youtube.com/watch?v=rhzKDrUiJVk)

---

## 📚 Additional Resources

- [godog GitHub Repository](https://github.com/cucumber/godog)
- [Gherkin Syntax Reference](https://cucumber.io/docs/gherkin/)
- [Go Regex Reference](https://pkg.go.dev/regexp)

---

## ✅ Summary

- You learned how to write effective BDD step definitions in Go
- You matched Gherkin phrases to Go using regular expressions
- You handled complex structured data using tables
- You understood how to manage state across steps with `ScenarioContext`

> In the next chapter, we’ll look at scenario outlines, background steps, and organizing step definitions across packages.
