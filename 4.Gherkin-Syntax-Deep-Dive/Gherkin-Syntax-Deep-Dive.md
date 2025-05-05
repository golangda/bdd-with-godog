
# Chapter: Gherkin Syntax Deep Dive

## ⏱ Duration: 1 hour

---

## 🎯 Learning Objectives

By the end of this chapter, you will be able to:
- Understand and use advanced Gherkin keywords: `And`, `But`, `Background`, and `Scenario Outline`
- Write expressive, readable, and maintainable `.feature` files
- Apply best practices to structure your BDD scenarios effectively

---

## 📖 What is Gherkin?

**Gherkin** is a domain-specific language used to write executable specifications in BDD. It is human-readable and structured to define system behavior without focusing on implementation details.

> Gherkin makes it easy for non-technical stakeholders to read and validate software requirements.

---

## 🧱 Gherkin Building Blocks (Recap)

- `Feature`: Describes the functionality under test.
- `Scenario`: A specific example or case under that feature.
- `Given / When / Then`: Define the preconditions, actions, and expected outcomes.

---

## 🔑 Deep Dive into Keywords

### 1. `And`, `But` – Chaining Steps

These keywords are used to improve the readability of `Given`, `When`, and `Then` steps.

    Scenario: Login to the app
      Given the user navigates to the login page
      And the user has a valid username and password
      When the user submits the login form
      Then the user should be redirected to the dashboard
      But the user should not see any error messages

📝 **Note:** `And` and `But` can follow `Given`, `When`, or `Then`.

---

### 2. `Background` – Shared Setup for All Scenarios

Use `Background` to define common steps that apply to every scenario in a feature.

    Feature: Manage To-Do List

    Background:
      Given the user has launched the To-Do app
      And the user is logged in

    Scenario: Add a new task
      When the user adds a task "Buy groceries"
      Then the task "Buy groceries" should be visible in the list

    Scenario: Delete a task
      Given the user has a task "Pay bills"
      When the user deletes the task "Pay bills"
      Then the task should be removed from the list

💡 **Best Practice:** Don’t overload `Background` with too many steps. Keep it relevant.

---

### 3. `Scenario Outline` – Parameterized Scenarios

`Scenario Outline` allows running the same scenario with different input data.

    Scenario Outline: Add multiple tasks
      Given the user has launched the To-Do app
      When the user adds a task "<task_name>"
      Then the task "<task_name>" should be visible in the list

    Examples:
      | task_name       |
      | Buy groceries   |
      | Pay bills       |
      | Call mom        |

✅ **When to use:** Useful for data-driven testing with multiple combinations.

---

## ✍️ Hands-on Exercise: Gherkin for To-Do List Feature

Write `.feature` file for the following user stories:

1. As a user, I want to add tasks so I can keep track of my to-dos.
2. As a user, I want to mark a task as complete so I know what I’ve finished.
3. As a user, I want to delete a task if it’s no longer needed.

    Feature: To-Do List Management

    Background:
      Given the user is logged in to the To-Do app

    Scenario: Add a new task
      When the user adds a task "Finish assignment"
      Then the task "Finish assignment" should be visible in the list

    Scenario: Mark a task as completed
      Given the user has a task "Read book"
      When the user marks the task "Read book" as completed
      Then the task "Read book" should appear as completed

    Scenario Outline: Delete a task
      Given the user has a task "<task_name>"
      When the user deletes the task "<task_name>"
      Then the task "<task_name>" should be removed from the list

    Examples:
      | task_name       |
      | Watch video     |
      | Pay rent        |

---

## ✅ Best Practices

| Practice | Description |
|---------|-------------|
| **Use plain language** | Avoid technical jargon. Keep scenarios readable by all team members |
| **One scenario, one behavior** | Each scenario should represent one atomic behavior |
| **Use meaningful examples** | Replace placeholder values with realistic, relatable data |
| **Reuse steps** | Avoid duplicating similar steps across features |

---

## 🧠 Interview Questions

1. What is the difference between `Scenario` and `Scenario Outline` in Gherkin?
2. How does `Background` improve scenario readability?
3. Why should we prefer `And`/`But` over repeating `Given`/`Then`?
4. Can you give an example of a poorly written scenario and how you would fix it?
5. How do you decide when to use `Background` vs setup inside a `Scenario`?

---

## 🎥 Curated YouTube Videos

Here are some high-quality beginner-friendly video tutorials:

1. **Cucumber Gherkin Tutorial for Beginners**  
   🔗 https://www.youtube.com/watch?v=U5X3vD4rCFE

2. **How to Write Gherkin Feature Files**  
   🔗 https://www.youtube.com/watch?v=n1tEcDgUe7E

3. **BDD with Cucumber, Gherkin and Examples**  
   🔗 https://www.youtube.com/watch?v=Em63T9Ecdjo

4. **Cucumber Tutorial: Scenario Outline vs Examples**  
   🔗 https://www.youtube.com/watch?v=yzIxf6VhAqA

---

## 📌 Summary

- Use Gherkin to define behavior in a human-readable way.
- Keywords like `And`, `But`, `Background`, and `Scenario Outline` improve structure and reusability.
- Good Gherkin scenarios act as both **living documentation** and **executable test cases**.

---

## 📚 References

- [Cucumber Gherkin Reference](https://cucumber.io/docs/gherkin/reference/)
- [godog Documentation](https://github.com/cucumber/godog)
