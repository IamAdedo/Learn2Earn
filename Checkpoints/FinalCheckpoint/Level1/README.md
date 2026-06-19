Here is  grouping all the exercises by **percentage (XP weight)** for the **Checkpoint** —

| Question # | % Weight | Exercises  |
|------------|----------|-------------------------------|
| **Level 1** | **5%** | `only1`, `onlya`, `onlyb`, `onlyf`, `onlyz`|

**Total: 5 exercises** (all in Checkpoint01)



## ✅ **Checkpoint Level 1 | in this category you might be given either onlyf, onlyz,onlyb and onlya**

### 🧠 **Goal**

Print the character `1` and **nothing else** — no newline, no space, no fmt.

---

### ✅ **Solution Code**

```go
package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune('1')
}
```

---

---

### ⚠️ **Important Notes**

* ❌ Do **not** use `fmt.Print` — unless your question allows it
* ❌ No spaces, no newline (`\n`) but if checker complains, then add it to your code.
* ✅ The output must be exactly:

  ```
  1
  ```
---

## ✅ **Question category1: `onlyz`**

---

### ✅ **Solution**

```go
package main

import "github.com/01-edu/z01"

func main() {
	z01.PrintRune('z')
}
```

the output must be exactly this

```
z
```


### ⚠️ **Important Notes**

When it comes to alphabets, the expected output is case sensitive. if the expected output is upper case, you should print upper and if it is lower, you should print lower.