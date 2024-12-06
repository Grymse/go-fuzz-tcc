# Version 0.1
- Create a TinyC grammar-based fuzzer
- Add probabilistic choice of grammar rule
- Add variable-system

# Version 0.2
- Add SmallC to grammars
  - int
  - functions, while, do-while, loop
  - expressions
- All code in main function
- Add max depth system to allow termination

# Version 1.0 (Full grammar)
- Full grammar with features
  - uint, int, float, double, char, string, arrays
  - functions, while, do-while, loop
  - expressions
- No use of probabilistic
- Max depth of 200 => Then always take fastest route to terminal

# Version 1.1
- Full grammar with features
  - uint, int, float, double, char, string, arrays
  - functions, while, do-while, loop
  - expressions
- No use of probabilistic
- Dynamic max-depth
  - 1: 100 waves, peak 50 depth, valley 4
  - 2: 100 waves, peak 50 depth, valley 4-10

1 - peak = 100, waves = 100
2 - peak = 300, waves = 100
3 - peak = 300, waves = 10
4 - valleyMax=4, valleyMin=4
5 - valleyMax=5, valleyMin=4
6 - waves = 100
7 - waves = 10, peak = 200


- Find bounds of values
- Sample within this
