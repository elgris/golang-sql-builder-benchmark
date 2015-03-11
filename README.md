golang-sql-builder-benchmark
====================

A collection of benchmarks for popular Go database/SQL builders. All builders construct the same SQL queries.

# Selects

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. squirrel lite (forked squirrel): https://github.com/elgris/squirrel

# Benchmarks

* **BenchmarkSelectDbrSimple** - Simple SQL query with dbr
* **BenchmarkSelectDbrComplex** - Complex SQL query with dbr
* **BenchmarkDbrSelectSubquery** - SQL query with dbr that uses subquery
* **BenchmarkSelectSquirrelSimple** - Simple SQL query with squirrel
* **BenchmarkSelectSquirrelComplex** - Complex SQL query with squirrel
* **BenchmarkSelectSquirrelComplex** - SQL query with squirrel that uses subquery
* **BenchmarkSelectSquirrelLiteSimple** - Simple SQL query with squirrel-lite
* **BenchmarkSelectSquirrelLiteComplex** - Complex SQL query with squirrel-lite
* **BenchmarkSelectSquirrelLiteComplex** - SQL query with squirrel-lite that uses subquery

# Output

`go test -bench=. -benchmem | column -t` on 2.6 GHz i5 Macbook Pro:

```
BenchmarkDbrSelectSimple             500000         3968     ns/op  864    B/op  14   allocs/op
BenchmarkDbrSelectComplex            50000          29207    ns/op  2215   B/op  42   allocs/op
BenchmarkDbrSelectSubquery           100000         10601    ns/op  2140   B/op  33   allocs/op

BenchmarkSquirrelSelectSimple        100000         19738    ns/op  2737   B/op  52   allocs/op
BenchmarkSquirrelSelectComplex       20000          66745    ns/op  11629  B/op  271  allocs/op
BenchmarkSquirrelSelectSubquery      20000          57307    ns/op  9246   B/op  199  allocs/op

BenchmarkSquirrelLiteSelectSimple    500000         3822     ns/op  952    B/op  15   allocs/op
BenchmarkSquirrelLiteSelectComplex   100000         22955    ns/op  3634   B/op  89   allocs/op
BenchmarkSquirrelLiteSelectSubquery  100000         15193    ns/op  2848   B/op  60   allocs/op
```
