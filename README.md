golang-sql-builder-benchmark
====================

A collection of benchmarks for popular Go database/SQL builders. All builders construct the same SQL queries.

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. squirrel lite (forked squirrel): https://github.com/elgris/squirrel

# Benchmarks

* **BenchmarkBuilderDbrSimple** - Simple SQL query with dbr
* **BenchmarkBuilderDbrComplex** - Complex SQL query with dbr
* **BenchmarkDbrBuilderSubquery** - SQL query with dbr that uses subquery
* **BenchmarkBuilderSquirrelSimple** - Simple SQL query with squirrel
* **BenchmarkBuilderSquirrelComplex** - Complex SQL query with squirrel
* **BenchmarkBuilderSquirrelComplex** - SQL query with squirrel that uses subquery
* **BenchmarkBuilderSquirrelLiteSimple** - Simple SQL query with squirrel-lite
* **BenchmarkBuilderSquirrelLiteComplex** - Complex SQL query with squirrel-lite
* **BenchmarkBuilderSquirrelLiteComplex** - SQL query with squirrel-lite that uses subquery

# Output

`go test -bench=. -benchmem | column -t` on 2.6 GHz i5 Macbook Pro:

```
BenchmarkDbrBuilderSimple             500000         3968     ns/op  864    B/op  14   allocs/op
BenchmarkDbrBuilderComplex            50000          29207    ns/op  2215   B/op  42   allocs/op
BenchmarkDbrBuilderSubquery           100000         10601    ns/op  2140   B/op  33   allocs/op

BenchmarkSquirrelBuilderSimple        100000         19738    ns/op  2737   B/op  52   allocs/op
BenchmarkSquirrelBuilderComplex       20000          66745    ns/op  11629  B/op  271  allocs/op
BenchmarkSquirrelBuilderSubquery      20000          57307    ns/op  9246   B/op  199  allocs/op

BenchmarkSquirrelLiteBuilderSimple    500000         3822     ns/op  952    B/op  15   allocs/op
BenchmarkSquirrelLiteBuilderComplex   100000         22955    ns/op  3634   B/op  89   allocs/op
BenchmarkSquirrelLiteBuilderSubquery  100000         15193    ns/op  2848   B/op  60   allocs/op
```
