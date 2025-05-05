# Chapter: Introduction to Gherkin & Cucumber

---

## 🧠 Objective

This chapter introduces you to **Gherkin**, a human-readable language used to define **behavior-driven development (BDD)** scenarios, and **Cucumber**, a tool that executes those scenarios. You’ll learn the Gherkin syntax, structure of a feature file, and how Cucumber connects the dots between feature files and your Go (or other language) code using `godog` in the Go ecosystem.

---

## 📘 What is Gherkin?

**Gherkin** is a domain-specific language for writing BDD-style acceptance tests in a natural language format.

- It is structured, readable, and understandable by both technical and non-technical stakeholders.
- Gherkin supports **over 70 spoken languages**.

### 🛠️ Real-World Analogy

Imagine you run a bakery.

Instead of writing test cases in a programming language, you define scenarios like:

```gherkin
Feature: Customer places an order

  Scenario: Customer orders chocolate cake
    Given the customer has selected a chocolate cake
    When they proceed to checkout
    Then the order should be confirmed
```

This format helps everyone — even your bakery manager or delivery partner — understand what the system is expected to do.

---

## 🧩 Gherkin Syntax Essentials

Here are the main keywords in Gherkin:

| Keyword  | Description |
|----------|-------------|
| `Feature` | The business functionality being described |
| `Scenario` | A specific example or test case |
| `Given` | Pre-conditions or starting context |
| `When` | The action or event the user performs |
| `Then` | The expected outcome or result |
| `And` / `But` | Used to chain conditions for better readability |

### 🧪 Example

```gherkin
Feature: User Login

  Scenario: Successful login with valid credentials
    Given the user is on the login page
    When the user enters valid credentials
    Then the user should be redirected to the dashboard

  Scenario: Failed login with invalid credentials
    Given the user is on the login page
    When the user enters incorrect password
    Then an error message should be shown
```

---

## 🛠️ Hands-On Example in Go (with `godog`)

Let’s say you’re building a simple CLI tool for login. Here’s how the Gherkin file looks:

### 📄 `features/login.feature`

```gherkin
Feature: User Login

  Scenario: Successful login
    Given a user with username "rahul" and password "secure123"
    When the user logs in with username "rahul" and password "secure123"
    Then the login should be successful
```

### 🧑‍💻 Go Step Definitions

```go
package main

import (
	"fmt"
)

var storedUser = struct {
	username string
	password string
}{}

func aUserWithUsernameAndPassword(username, password string) error {
	storedUser.username = username
	storedUser.password = password
	return nil
}

func theUserLogsInWithUsernameAndPassword(username, password string) error {
	if storedUser.username != username || storedUser.password != password {
		return fmt.Errorf("invalid credentials")
	}
	return nil
}

func theLoginShouldBeSuccessful() error {
	// You could set flags, mock state, etc.
	return nil
}
```

In `godog`, you would register these step functions with the actual regex matching the Gherkin steps.

---

## 🧪 How Does Cucumber Work?

Cucumber acts as the **glue** between your feature files and the implementation code.

### 🧬 Execution Flow:

```
[Feature File in Gherkin] → [Step Definitions] → [Application Code]
```

- **Cucumber (or godog)** reads `.feature` files.
- It **matches each step** with a corresponding function.
- The matched function is executed and the result is displayed.

---

## ❓ Common Interview Questions

1. **What is Gherkin in BDD?**
   - Gherkin is a plain-text language used to define behavior in a structured format using `Given/When/Then`.

2. **What is the difference between a Feature and a Scenario?**
   - A Feature defines a high-level functionality; a Scenario is a specific test case under that feature.

3. **Can non-programmers write Gherkin?**
   - Yes, Gherkin is designed to be understood by non-developers such as product owners or QA analysts.

4. **What does Cucumber do?**
   - Cucumber reads Gherkin-based `.feature` files and maps them to automation code to validate software behavior.

5. **How does Cucumber handle step definitions?**
   - Each step in the Gherkin file is matched to a regular expression function in a supported language.

---

## 🎥 Recommended YouTube Tutorials

Here are some well-curated tutorials to deepen your understanding:

- 📺 [BDD with Cucumber, Gherkin and Godog in Go (Golang)](https://www.youtube.com/watch?v=DJtaPaC5s4w)
- 📺 [Cucumber for Beginners | Full Course](https://www.youtube.com/watch?v=H9jD7aFeBHU)
- 📺 [Gherkin Language Basics](https://www.youtube.com/watch?v=hdcoT2j4KFc)
- 📺 [BDD with godog in GoLang – Short Intro](https://www.youtube.com/watch?v=tPUfU9RkEXg)

---

## 🧰 Tools You Should Try

- `godog` – Go-based implementation of Cucumber
- VS Code with Cucumber plugin – syntax highlighting
- [https://cucumber.io/docs](https://cucumber.io/docs) – official documentation

---

## ✅ Summary

- Gherkin helps bridge the gap between business users and developers.
- Cucumber and its variants (like `godog`) allow you to **execute** these descriptions as automated tests.
- Practicing writing `.feature` files and implementing them in Go using `godog` is the best way to master this skill.

---

Next Chapter 👉 **Step Definitions and Integrating with `godog`**