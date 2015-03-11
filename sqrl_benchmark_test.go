package db_sql_benchmark

import (
	"testing"

	"github.com/elgris/sqrl"
	_ "github.com/go-sql-driver/mysql"
)

//
// Select benchmarks
//
func BenchmarkSqrlSelectSimple(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkSqrlSelectConditional(b *testing.B) {
	for n := 0; n < b.N; n++ {
		qb := sqrl.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		if n%2 == 0 {
			qb.GroupBy("subdomain_id").
				Having("number = ?", 1).
				OrderBy("state").
				Limit(7).
				Offset(8)
		}

		qb.ToSql()
	}
}

func BenchmarkSqrlSelectComplex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			Where(sqrl.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			Where(sqrl.Eq{"h": []int{1, 2, 3}}).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having("j = k").
			Having("jj = ?", 1).
			Having("jjj = ?", 2).
			OrderBy("l").
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()
	}
}

func BenchmarkSqrlSelectSubquery(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subSelect := sqrl.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		sqrl.Select("a", "b").
			From("c").
			Distinct().
			Column(sqrl.Alias(subSelect, "subq")).
			Where(sqrl.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()

	}
}

func BenchmarkSqrlSelectMoreComplex(b *testing.B) {
	for i := 0; i < b.N; i++ {

		sqrl.Select("a", "b").
			Prefix("WITH prefix AS ?", 0).
			Distinct().
			Columns("c").
			Column("IF(d IN ("+sqrl.Placeholders(3)+"), 1, 0) as stat_column", 1, 2, 3).
			Column(sqrl.Expr("a > ?", 100)).
			Column(sqrl.Eq{"b": []int{101, 102, 103}}).
			From("e").
			JoinClause("CROSS JOIN j1").
			Join("j2").
			LeftJoin("j3").
			RightJoin("j4").
			Where("f = ?", 4).
			Where(sqrl.Eq{"g": 5}).
			Where(map[string]interface{}{"h": 6}).
			Where(sqrl.Eq{"i": []int{7, 8, 9}}).
			Where(sqrl.Or{sqrl.Expr("j = ?", 10), sqrl.And{sqrl.Eq{"k": 11}, sqrl.Expr("true")}}).
			GroupBy("l").
			Having("m = n").
			OrderBy("o ASC", "p DESC").
			Limit(12).
			Offset(13).
			Suffix("FETCH FIRST ? ROWS ONLY", 14).
			ToSql()
	}
}

//
// Insert benchmark
//
func BenchmarkSqrlInsert(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Insert("mytable").
			Columns("id", "a", "b", "price", "created", "updated").
			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05").
			ToSql()
	}
}

//
// Update benchmark
//
func BenchmarkSqrlUpdateSetColumns(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Update("mytable").
			Set("foo", 1).
			Set("bar", sqrl.Expr("COALESCE(bar, 0) + 1")).
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10).
			Offset(20).
			ToSql()
	}
}

func BenchmarkSqrlUpdateSetMap(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Update("mytable").
			SetMap(map[string]interface{}{"b": 1, "c": 2, "bar": sqrl.Expr("COALESCE(bar, 0) + 1")}).
			Where("id = ?", 9).
			Limit(10).
			Offset(20).
			ToSql()
	}
}

//
// Delete benchmark
//
func BenchmarkSqrlDelete(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sqrl.Delete("test_table").
			Where("b = ?", 1).
			OrderBy("c").
			Limit(2).
			Offset(3).
			ToSql()
	}
}
