# Quadchecker 

A Go program that reads a shape from stdin and identifies which quad function produced it, along with its dimensions.

---

## Usage

```bash
./quadA 3 3 | ./quadchecker
```

Output

\# [quadA] [3] [3]

```bash
./quadC 1 1 | ./quadchecker
```

Output

\# [quadC] [1] [1] || [quadD] [1] [1] || [quadE] [1] [1]

```bash
echo "hello" | ./quadchecker
```
Output

\# Not a quad function


---


## How It Works — Step by Step

### Step 1 — Read Input from Stdin

```go
data, err := io.ReadAll(os.Stdin)
input := string(data)
```

The program does not take command-line arguments. Instead it reads everything piped into it via stdin — the entire shape as a string, newlines included.

If reading fails, it prints `Not a quad function` and exits.

---

### Step 2 — Get the Dimensions of the Shape

```go
x, y := getDimensions(input)
```

Before comparing the input against any known quad, the program needs to know the width (`x`) and height (`y`) of the shape.

#### Inside `getDimensions`:

**2a. Split into lines**
```go
lines := strings.Split(s, "\n")
```
The input is split on newline characters. For example:
```
"o-o\n| |\no-o\n"  →  ["o-o", "| |", "o-o", ""]
```

**2b. Strip the trailing empty string**
```go
if len(lines) > 0 && lines[len(lines)-1] == "" {
    lines = lines[:len(lines)-1]
}
```
Every quad shape ends with `\n`. When you split `"a\nb\n"` on `\n`, you get `["a", "b", ""]` — a trailing empty element. This removes it so only real lines remain.

**2c. Height = number of lines**
```go
height := len(lines)
```

**2d. Width = length of the first line**
```go
width := len(lines[0])
```

**2e. Validate — all lines must be the same width**
```go
for _, line := range lines {
    if len(line) != width {
        return 0, 0
    }
}
```
If any line has a different length, the shape is malformed. Return `0, 0` to signal invalid input.

**2f. Return dimensions**
```go
return width, height
```

If dimensions come back as `0, 0`, the program prints `Not a quad function` and exits.

---

### Step 3 — Reconstruct Each Known Quad and Compare

```go
if input == makeQuad(x, y, 'o', 'o', 'o', 'o', '-', '|') {
    matches = append(matches, fmt.Sprintf("[quadA] [%d] [%d]", x, y))
}
```

For each of the 5 known quad types, the program **generates the expected shape** using the same dimensions as the input, then does a **direct string comparison**.

If they match, it records the result. This is the core logic — no pattern recognition, just exact equality.

#### Inside `makeQuad`:

```go
func makeQuad(x, y int, tl, tr, bl, br, edge, vert rune) string
```

Parameters:
| Parameter | Meaning |
|-----------|---------|
| `x` | Width (number of columns) |
| `y` | Height (number of rows) |
| `tl` | Top-left corner character |
| `tr` | Top-right corner character |
| `bl` | Bottom-left corner character |
| `br` | Bottom-right corner character |
| `edge` | Top and bottom edge character (between corners) |
| `vert` | Left and right vertical edge character (between corners) |

**The drawing logic — row by row, column by column:**

```go
for row := 0; row < y; row++ {
    for col := 0; col < x; col++ {
        switch {
        case row == 0 && col == 0:         // top-left corner
        case row == 0 && col == x-1:       // top-right corner
        case row == y-1 && col == 0:       // bottom-left corner
        case row == y-1 && col == x-1:     // bottom-right corner
        case row == 0 || row == y-1:       // top or bottom edge
        case col == 0 || col == x-1:       // left or right vertical edge
        default:                           // interior space
        }
    }
    sb.WriteByte('\n') // end of each row
}
```

Corners are checked **first** (most specific), edges second, interior last. Order matters here — if corners were checked after edges they would be drawn with the wrong character.

Uses `strings.Builder` for efficient string construction instead of repeated `+=` concatenation.

---

### Step 4 — The 5 Quad Definitions

Each quad is defined by its corner and edge characters:

| Quad | TL | TR | BL | BR | Edge | Vert |
|------|----|----|----|----|------|------|
| A | `o` | `o` | `o` | `o` | `-` | `\|` |
| B | `/` | `\` | `\` | `/` | `*` | `*` |
| C | `A` | `A` | `C` | `C` | `B` | `B` |
| D | `A` | `C` | `A` | `C` | `B` | `B` |
| E | `A` | `C` | `C` | `A` | `B` | `B` |

**Visual examples at 3x3:**

```
quadA        quadB        quadC        quadD        quadE
o-o           /*\          AAA          ABA          ABA
| |           * *          B B          B B          B B
o-o           \*/          CCC          ACA          CAC
```

Note: quadC, D, and E share the same characters (`A`, `B`, `C`) but arrange the corners differently. This is why a 1x1 input of `A\n` matches all three — at size 1x1 every position is a corner, and all three quads put `A` in the top-left.

---

### Step 5 — Collect and Output Matches

```go
matches := []string{}
// ... append to matches for each quad that matches

fmt.Println(strings.Join(matches, " || "))
```

All matches are collected into a slice first, then joined with ` || ` as separator. Since the checks are written in alphabetical order (A → B → C → D → E), the output is always alphabetically sorted without any sorting step needed.

If no matches are found:

```go
if len(matches) == 0 {
    fmt.Println("Not a quad function")
}
```

---

## Edge Cases Handled

| Input | Behaviour |
|-------|------------------|
| Empty input | `getDimensions` returns `0,0` → `Not a quad function` |
| No trailing `\n` | Lines won't match any `makeQuad` output → `Not a quad function` |
| Inconsistent line widths | `getDimensions` returns `0,0` → `Not a quad function` |
| 1x1 input matching multiple quads | All matching quads printed, separated by ` \|\| ` |
| Valid shape but unknown pattern | `matches` stays empty → `Not a quad function` |

---

## File Structure

```
quadchecker/
├── go.mod       ← module definition (no external dependencies)
└── main.go      ← entire program lives here
```