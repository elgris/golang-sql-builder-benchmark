package db_sql_benchmark

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doug-martin/goqu"
	_ "github.com/go-sql-driver/mysql"
)

var driver *sql.DB

func init() {
	db, _ := sqlmock.New()
	driver = db
}

//
// Select benchmarks
//
func BenchmarkGoquSelectSimple(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		db.From("tickets").
			Where(
			goqu.And(
				goqu.I("subdomain_id").Eq(1),
				goqu.Or(
					goqu.I("state").Eq("open"),
					goqu.I("state").Eq("spam"),
				),
			),
		).
			ToSql()
	}
}

func BenchmarkGoquSelectConditional(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		qb := db.From("tickets").
			Where(
			goqu.And(
				goqu.I("subdomain_id").Eq(1),
				goqu.Or(
					goqu.I("state").Eq("open"),
					goqu.I("state").Eq("spam"),
				),
			),
		)
		if n%2 == 0 {
			qb.GroupBy("subdomain_id").
				Having(goqu.I("number").Eq(1)).
				Order(goqu.I("state").Asc()).
				Limit(7).
				Offset(8)
		}

		qb.ToSql()
	}
}

func BenchmarkGoquSelectComplex(b *testing.B) {
	db := goqu.New("default", driver)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		db.From("a", "b", "z", "y").
			Select(goqu.DISTINCT("x")).
			Where(
			goqu.Or(
				goqu.I("d").Eq(1),
				goqu.I("e").Eq("wat"),
			)).
			Where(
			goqu.And(
				goqu.I("f").Eq(2),
				goqu.I("x").Eq("hi"),
			)).
			Where(
			goqu.And(
				goqu.I("g").Eq(3),
			)).
			Where(
			goqu.And(
				goqu.I("h").Eq([]int{1, 2, 3}),
			)).
			GroupBy("i").
			GroupBy("ii").
			GroupBy("iii").
			Having(goqu.I("j = k")).
			Having(goqu.I("jj").Eq(1)).
			Having(goqu.I("jjj").Eq(2)).
			Order(goqu.I("l").Asc()).
			Order(goqu.I("l").Asc()).
			Order(goqu.I("l").Asc()).
			Limit(7).
			Offset(8).
			ToSql()
	}
}
