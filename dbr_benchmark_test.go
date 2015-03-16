package db_sql_benchmark

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

//
// Select benchmarks
//

func dbrSess() *dbr.Session {
	return dbr.NewConnection(nil, nil).NewSession(nil)
}

func BenchmarkDbrSelectSimple(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sess.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkDbrSelectConditional(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		qb := sess.Select("id").
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

func BenchmarkDbrSelectComplex(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sess.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			Where(dbr.Eq{"h": []int{1, 2, 3}}).
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

func BenchmarkDbrSelectSubquery(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subQuery, _ := sess.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()

		sess.Select("a", "b", fmt.Sprintf("(%s) AS subq", subQuery)).
			From("c").
			Distinct().
			Where(dbr.Eq{"f": 2, "x": "hi"}).
			Where(map[string]interface{}{"g": 3}).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()

	}

}

//
// Insert benchmark
//
func BenchmarkDbrInsert(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sess.InsertInto("mytable").
			Columns("id", "a", "b", "price", "created", "updated").
			Values(1, "test_a", "test_b", 100.05, "2014-01-05", "2015-01-05").
			ToSql()
	}
}

//
// Update benchmark
//
func BenchmarkDbrUpdateSetColumns(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sess.Update("mytable").
			Set("foo", 1).
			Set("bar", dbr.Expr("COALESCE(bar, 0) + 1")).
			Set("c", 2).
			Where("id = ?", 9).
			Limit(10).
			Offset(20).
			ToSql()
	}
}

func BenchmarkDbrUpdateSetMap(b *testing.B) {
	sess := dbrSess()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sess.Update("mytable").
			SetMap(map[string]interface{}{"b": 1, "c": 2, "bar": dbr.Expr("COALESCE(bar, 0) + 1")}).
			Where("id = ?", 9).
			Limit(10).
			Offset(20).
			ToSql()
	}
}

//
// Delete benchmark
//
func BenchmarkDbrDelete(b *testing.B) {
	sess := dbrSess()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		sess.DeleteFrom("test_table").
			Where("b = ?", 1).
			OrderBy("c").
			Limit(2).
			Offset(3).
			ToSql()
	}
}
