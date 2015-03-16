golang-sql-builder-benchmark
====================

A comparison of popular Go SQL query builders. Provides feature list and benchmarks

# Builders

1. dbr: https://github.com/gocraft/dbr
2. squirrel: https://github.com/lann/squirrel
3. sqrl: https://github.com/elgris/sqrl
4. gocu: github.com/doug-martin/goqu - just for SELECT query


# Feature list

| feature                    | dbr | squirrel | sqrl | goqu |
|----------------------------|-----|----------|------|------|
| SelectBuilder              | +   | +        | +    | +    |
| InsertBuilder              | +   | +        | +    | +    |
| UpdateBuilder              | +   | +        | +    | +    |
| DeleteBuilder              | +   | +        | +    | +    |
| PostgreSQL support         |     | +        | +    | +    |
| Custom placeholders        |     | +        | +    | +    |
| JOINs support              |     | +        | +    | +    |
| Subquery in query builder  |     | +        | +    | +    |
| Aliases for columns        |     | +        | +    | +    |
| CASE expression            |     | +        | +    | +    |

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
BenchmarkDbrSelectSimple            500000       2610     ns/op  864    B/op  14   allocs/op
BenchmarkDbrSelectConditional       500000       3808     ns/op  1031   B/op  19   allocs/op
BenchmarkDbrSelectComplex           200000       11585    ns/op  3323   B/op  53   allocs/op
BenchmarkDbrSelectSubquery          200000       10025    ns/op  2851   B/op  40   allocs/op
BenchmarkDbrInsert                  500000       3717     ns/op  1136   B/op  19   allocs/op
BenchmarkDbrUpdateSetColumns        300000       4106     ns/op  1038   B/op  24   allocs/op
BenchmarkDbrUpdateSetMap            300000       5396     ns/op  1388   B/op  26   allocs/op
BenchmarkDbrDelete                  1000000      2150     ns/op  482    B/op  13   allocs/op


BenchmarkGoquSelectSimple           100000       15180    ns/op  3282   B/op  46   allocs/op
BenchmarkGoquSelectConditional      100000       19655    ns/op  4258   B/op  61   allocs/op
BenchmarkGoquSelectComplex          30000        50628    ns/op  11414  B/op  215  allocs/op


BenchmarkSqrlSelectSimple           500000       3555     ns/op  952    B/op  15   allocs/op
BenchmarkSqrlSelectConditional      300000       4377     ns/op  1112   B/op  20   allocs/op
BenchmarkSqrlSelectComplex          100000       24040    ns/op  4751   B/op  100  allocs/op
BenchmarkSqrlSelectSubquery         100000       26203    ns/op  3560   B/op  67   allocs/op
BenchmarkSqrlSelectMoreComplex      30000        47018    ns/op  7256   B/op  150  allocs/op
BenchmarkSqrlInsert                 200000       7773     ns/op  1304   B/op  25   allocs/op
BenchmarkSqrlUpdateSetColumns       200000       8633     ns/op  1369   B/op  32   allocs/op
BenchmarkSqrlUpdateSetMap           200000       15786    ns/op  1788   B/op  36   allocs/op
BenchmarkSqrlDelete                 500000       3669     ns/op  496    B/op  12   allocs/op


BenchmarkSquirrelSelectSimple       100000       14934    ns/op  2737   B/op  52   allocs/op
BenchmarkSquirrelSelectConditional  100000       18034    ns/op  4023   B/op  84   allocs/op
BenchmarkSquirrelSelectComplex      20000        63096    ns/op  12742  B/op  283  allocs/op
BenchmarkSquirrelSelectSubquery     30000        48956    ns/op  9954   B/op  206  allocs/op
BenchmarkSquirrelSelectMoreComplex  20000        83842    ns/op  17153  B/op  386  allocs/op
BenchmarkSquirrelInsert             100000       14517    ns/op  3356   B/op  75   allocs/op
BenchmarkSquirrelUpdateSetColumns   100000       23995    ns/op  4787   B/op  108  allocs/op
BenchmarkSquirrelUpdateSetMap       50000        27141    ns/op  5203   B/op  112  allocs/op
BenchmarkSquirrelDelete             100000       16728    ns/op  2815   B/op  67   allocs/op
```

# Conclusion

If your queries are very simple, pick `dbr`, the fastest one.

If really need immutability of query builder and you're ready to sacrifice extra memory, use `squirrel`, the slowest but most reliable one.

If you like those sweet helpers that `squirrel` provides to ease your query building or if you plan to use the same builder for `PostgreSQL`, take `sqrl` as it's balanced between performance and features.

`goqu` has LOTS of features and ways to build queries. Although it requires stubbing sql connection if you need just to build a query. It can be done with [sqlmock](http://github.com/DATA-DOG/go-sqlmock). Disadvantage: the builder is slow and has TOO MANY features, so building a query may become a nightmare. But if you need total control on everything - this is your choice.