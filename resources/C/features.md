**Values:**
- [X] Ints (signed and unsigned)
  - [ ] Unsigned 16-64
  - [ ] Signed 16-64
- [-] Strings (with basic operations like concatenation and length checks)
- [-] Chars (useful for lightweight string fuzzing)
- [-] Floats (basic arithmetic, though fuzzing them adds complexity)
  - [ ] Float
  - [ ] Double
- [ ] Boolean

**Variables:**
- [X] Local Variables
- [X] Scoped Shadowing (e.g., re-declaring a variable in a narrower scope)
- [X] Global Variables
- [X] Constants (const keyword)

**Expressions:**
- [X] Comparison (==, !=, <=, >=, <, >)
- [X] Arithmetic (+, -, /, *, %)
- [X] Logical (&&, ||, !)
- [ ] Bitwise (&, |, ^, ~)
- [ ] Shifting (<<, >>)
- [ ] Increment/Decrement (++, --)

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

**Memory and Pointers**
- [-] Pointers (basic dereferencing, * and &)
- [-] Arrays (1D only for simplicity)
- [-] malloc and free (manual memory management)
- [-] sizeof operator

**I/O**
- [-] Basic printf for debugging output
- [-] File I/O (optional, e.g., reading/writing small files)




**TODO**
- Add shadowing variables
- Add types