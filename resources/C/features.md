**Values:**
- [X] Ints (signed and unsigned)
  - [X] Unsigned 16-64
  - [X] Signed 16-64
- [-] Strings (with basic operations like concatenation and length checks)
- [X] Chars (useful for lightweight string fuzzing)
- [X] Floats (basic arithmetic, though fuzzing them adds complexity)
  - [X] Float
  - [X] Double
- [X] Boolean
- [X] Arrays (1D only for simplicity)

**Variables:**
- [X] Local Variables
- [X] Scoped Shadowing (e.g., re-declaring a variable in a narrower scope)
- [X] Global Variables
- [X] Constants (const keyword)

**Expressions:**
- [X] Comparison (==, !=, <=, >=, <, >)
- [X] Arithmetic (+, -, /, *)
- [X] Logical (&&, ||, !)

**Statements**
- [X] if
- [X] if-else
- [X] While-loop
- [X] switch (minimal, with a default case)
- [X] Do-While
- [X] For loop
- [X] break and continue

**Functions**
- [X] Function Declarations
- [X] Function Calls
- [X] Recursive Functions (optional for added complexity)