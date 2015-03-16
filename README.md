golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders. Provides feature list and benchmarks

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. sqrl: https://github.com/elgris/sqrl


# Feature list

| feature                    | dbr | squirrel | sqrl |
|----------------------------|-----|----------|------|
| SelectBuilder              | +   | +        | +    |
| InsertBuilder              | +   | +        | +    |
| UpdateBuilder              | +   | +        | +    |
| DeleteBuilder              | +   | +        | +    |
| PostgreSQL support         |     | +        | +    |
| Custom placeholders        |     | +        | +    |
| JOINs support              |     | +        | +    |
| Subquery in query builder  |     | +        | +    |
| Aliases for columns        |     | +        | +    |
| CASE expression            |     | +        | +    |

Some explanations here:
- `Custom placeholders` - ability to use not only `?` placeholders, Useful for PostgreSQL
- `JOINs support` - ability to build JOINs in SELECT queries like `Select("*").From("a").Join("b")`
- `Subquery in query builder` - when you prepare a subquery with one builder and then pass it to another. Something like:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(subQ).From("a")
```
- `Aliases for columns` - easy way to alias a column, especially if column is specified by subquery:
```go
subQ := Select("aa", "bb").From("dd")
qb := Select().Column(Alias(subQ, "alias")).From("a")
```
- `CASE expression` - syntactic sugar for [CASE expressions](http://dev.mysql.com/doc/refman/5.7/en/case.html)

# Benchmarks

`go test -bench=. -benchmem | column -t` on 2.6 GHz i5 Macbook Pro:

```
BenchmarkDbrSelectSimple            1000000                                         2698     ns/op  864    B/op  14   allocs/op
BenchmarkDbrSelectConditional       300000                                          4067     ns/op  1031   B/op  19   allocs/op
BenchmarkDbrSelectComplex           100000                                          12274    ns/op  3325   B/op  53   allocs/op
BenchmarkDbrSelectSubquery          200000                                          9904     ns/op  2853   B/op  40   allocs/op
BenchmarkDbrInsert                  500000                                          3945     ns/op  1136   B/op  19   allocs/op
BenchmarkDbrUpdateSetColumns        300000                                          4187     ns/op  1039   B/op  24   allocs/op
BenchmarkDbrUpdateSetMap            300000                                          5341     ns/op  1389   B/op  26   allocs/op
BenchmarkDbrDelete                  1000000                                         2262     ns/op  483    B/op  13   allocs/op
BenchmarkSqrlSelectSimple           500000                                          3159     ns/op  952    B/op  15   allocs/op
BenchmarkSqrlSelectConditional      300000                                          4816     ns/op  1112   B/op  20   allocs/op
BenchmarkSqrlSelectComplex          100000                                          20034    ns/op  4750   B/op  100  allocs/op
BenchmarkSqrlSelectSubquery         100000                                          13888    ns/op  3560   B/op  67   allocs/op
BenchmarkSqrlSelectMoreComplex      50000                                           30291    ns/op  7264   B/op  150  allocs/op
BenchmarkSqrlInsert                 300000                                          4854     ns/op  1304   B/op  25   allocs/op
BenchmarkSqrlUpdateSetColumns       300000                                          5709     ns/op  1369   B/op  32   allocs/op
BenchmarkSqrlUpdateSetMap           200000                                          9736     ns/op  1788   B/op  36   allocs/op
BenchmarkSqrlDelete                 1000000                                         1930     ns/op  496    B/op  12   allocs/op
BenchmarkSquirrelSelectSimple       200000                                          10824    ns/op  2737   B/op  52   allocs/op
BenchmarkSquirrelSelectConditional  100000                                          16815    ns/op  4025   B/op  84   allocs/op
BenchmarkSquirrelSelectComplex      30000                                           58980    ns/op  12743  B/op  283  allocs/op
BenchmarkSquirrelSelectSubquery     30000                                           45759    ns/op  9960   B/op  206  allocs/op
BenchmarkSquirrelSelectMoreComplex  20000                                           80773    ns/op  17158  B/op  386  allocs/op
BenchmarkSquirrelInsert             100000                                          15453    ns/op  3361   B/op  75   allocs/op
BenchmarkSquirrelUpdateSetColumns   50000                                           22643    ns/op  4786   B/op  108  allocs/op
BenchmarkSquirrelUpdateSetMap       50000                                           24155    ns/op  5204   B/op  112  allocs/op
BenchmarkSquirrelDelete             100000                                          13370    ns/op  2817   B/op  67   allocs/op
```

# Conclusion

If your queries are very simple, pick `dbr`, the fastest one.

If really need immutability of query builder and you're ready to sacrifice extra memory, use `squirrel`, the slowest but most reliable one.

If you like those sweet helpers that `squirrel` provides to ease your query building or if you plan to use the same builder for `PostgreSQL`, take `sqrl` as it's balanced between performance and features.