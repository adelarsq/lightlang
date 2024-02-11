# Design Principles

Main principles:

- Performance
- Simple syntax - as minimal as possible
- Simple but robust toolbox

Based on Fennel:

- No early returns

Based on F#:

- Functional first
- Indentation based
- File ordering 
- Units of measure

Based on V:

- No null
- No global variables
- No undefined values
- No undefined behavior
- No variable shadowing
- Bounds checking
- Immutable variables by default
- Pure functions by default
- Immutable structs by default
- Option/Result and mandatory error checks
- Sum types
- Generics

Based on Zig:

- No hidden control flow
- No hidden memory allocations
- No preprocessor, no macros ???

New features:

- Custom keywords
- Any chars can be used (usefull for Math and Science)

## CLI Commands

Any CLI command can be called on any folder.

- light or light help <topic> - help about CLI commands
- light install - install packages (including packages from the host language)
- light run - execute current project
- light watch - watch changes on the code
- light test - executes tests and shows coverage
- light deps - shows dependency information (inclusing packages from the host language)
- light fmt - format code
- light bug - starts a bug report
- light doc - shows documentation for the current project or package
- light version - shows version


