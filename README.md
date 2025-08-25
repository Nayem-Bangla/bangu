## Bangu — a tiny interpreted language in Go

Bangu is a small interpreter written in Go featuring a lexer, Pratt parser, AST, evaluator with environments and closures, arrays and hashes, strings, and a simple REPL.

### Highlights
- **Types**: integers, booleans, strings, null
- **Operators**: `+ - * / < > == !=` and prefix `- !`
- **Bindings**: `let x = 5;`
- **Control flow**: `if (cond) { ... } else { ... }`
- **Functions & closures**: `fn(x, y) { x + y; }`
- **Collections**: arrays `[1,2,3]`, hashes `{ "k": 1, 2: 4, true: 5 }`
- **Builtins**: `len`, `first`, `last`, `rest`, `push`, `puts`
- **REPL** with persistent environment

### Quick start

Requirements: Go 1.22+

Run the REPL:

```bash
go run ./main
```

Example session:

```text
>> let add = fn(x, y) { x + y; };
>> add(5, 7)
12
>> [1, 2, 3][1]
2
>> {"greet": "Hello, " + "World!"}["greet"]
Hello, World!
```

### Examples

Copy and paste these directly into the REPL (`go run main.go`).

#### Functions and calls
```text;
>> len("Hello World!")
12
>> len("")
0
>> len("Hey Bob, how ya doin?")
21
>> len("1234")
4
```

#### Arrays and builtins
```text
>> let a = [1, 2 * 2, 10 - 5, 8 / 2];
>> a[0]
1
>> a[1]
4
>> a[5 - 3]
5
>> a[99]
null
>> let a = [1, 2, 3, 4];
>> let b = push(a, 5);
>> a
[1, 2, 3, 4]
>> b
[1, 2, 3, 4, 5]
```

#### Hashes (maps)
```text
>> let people = [{"name": "Alice", "age": 24}, {"name": "Anna", "age": 28}];
>> people[0]["name"];
Alice
>> people[1]["age"];
28
>> people[1]["age"] + people[0]["age"];
52
>> let getName = fn(person) { person["name"]; };
>> getName(people[0]);
Alice
>> getName(people[1]);
Anna
```

#### Control flow
```text
if (1 < 2) { 10 } else { 20 }
```

#### Strings
```text
>> let firstName = "Arafat";
>> let lastName = "Hasan";
>> let fullName = fn(first, last) { first + " " + last };
>> fullName(firstName, lastName);
Arafat Hasan
```

### Project layout
- `token/` — token types and keyword lookup
- `lexer/` — input to tokens
- `ast/` — AST nodes and stringification
- `parser/` — Pratt parser with precedence and associativity
- `object/` — runtime objects, arrays, hashes, functions, builtins
- `evaluator/` — tree-walking evaluator, environments and closures
- `repl/` — interactive prompt
- `main/` — REPL entrypoint

### Language cheatsheet
- Bindings: `let x = 5;`
- Functions: `let add = fn(x, y) { x + y; }; add(2, 3)`
- If: `if (1 < 2) { 10 } else { 20 }`
- Arrays: `[1,2,3][0]` → 1, `push([1,2], 3)` → `[1, 2, 3]`
- Hashes: `{ "one": 1, 2: 4, true: 5 }["one"]` → 1
- Strings: `"Hello, " + "World!"` → `Hello, World!`
- Builtins: `len`, `first`, `last`, `rest`, `push`, `puts`

### Tests

```bash
go test ./...
```

The suite covers lexing, parsing, evaluation, collections, builtins, and errors.

### Implementation notes
- Pratt parser with precedence: sum, product, call, index
- Environments provide lexical scoping; functions close over `Env`
- Hash keys: string, integer, boolean (types implementing `Hashable`)
- Strings support `+` concatenation

### Roadmap
- Execute `.bangu` files from CLI
- Additional numeric types
- Loops and assignments
- Standard library modules
- Richer error messages with source positions

### License
MIT

