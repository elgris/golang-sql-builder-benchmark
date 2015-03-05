package db_sql_benchmark

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	squirlite "github.com/elgris/squirrel"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
	"github.com/lann/squirrel"
)

const MysqlDSN = "root:111@tcp(docker:3306)/test"

//
// dbr
//

func mysqlConn() *sql.DB {
	db, err := sql.Open("mysql", MysqlDSN)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func dbrSess() *dbr.Session {
	return dbr.NewConnection(mysqlConn(), nil).NewSession(nil)
}

func BenchmarkDbrBuilderSimple(b *testing.B) {
	sess := dbrSess()

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sess.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkDbrBuilderComplex(b *testing.B) {
	sess := dbrSess()

	arg_eq1 := dbr.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}
	arg_eq3 := dbr.Eq{"h": []int{1, 2, 3}}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sess.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			Where(arg_eq1).
			Where(arg_eq2).
			Where(arg_eq3).
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

func BenchmarkDbrBuilderSubquery(b *testing.B) {
	sess := dbrSess()

	arg_eq1 := dbr.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subQuery, _ := sess.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()

		sess.Select("a", "b", fmt.Sprintf("(%s) AS subq", subQuery)).
			From("c").
			Distinct().
			Where(arg_eq1).
			Where(arg_eq2).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()

	}
}

//
// Squirrel
//

func BenchmarkSquirrelBuilderSimple(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		squirrel.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkSquirrelBuilderComplex(b *testing.B) {
	arg_eq1 := squirrel.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}
	arg_eq3 := squirrel.Eq{"h": []int{1, 2, 3}}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		squirrel.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			Where(arg_eq1).
			Where(arg_eq2).
			Where(arg_eq3).
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

func BenchmarkSquirrelBuilderSubquery(b *testing.B) {
	arg_eq1 := squirrel.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subBuilder := squirrel.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		squirrel.Select("a", "b").
			From("c").
			Distinct().
			Column(squirrel.Alias(subBuilder, "subq")).
			Where(arg_eq1).
			Where(arg_eq2).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()

	}
}

//
// squirrel lite
//

func BenchmarkSquirrelLiteBuilderSimple(b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		squirlite.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam").
			ToSql()
	}
}

func BenchmarkSquirrelLiteBuilderComplex(b *testing.B) {
	arg_eq1 := squirlite.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}
	arg_eq3 := squirlite.Eq{"h": []int{1, 2, 3}}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		squirlite.Select("a", "b", "z", "y", "x").
			Distinct().
			From("c").
			Where("d = ? OR e = ?", 1, "wat").
			Where(arg_eq1).
			Where(arg_eq2).
			Where(arg_eq3).
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

func BenchmarkSquirrelLiteBuilderSubquery(b *testing.B) {
	arg_eq1 := squirlite.Eq{"f": 2, "x": "hi"}
	arg_eq2 := map[string]interface{}{"g": 3}

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		subBuilder := squirlite.Select("id").
			From("tickets").
			Where("subdomain_id = ? and (state = ? or state = ?)", 1, "open", "spam")

		squirlite.Select("a", "b").
			From("c").
			Distinct().
			Column(squirlite.Alias(subBuilder, "subq")).
			Where(arg_eq1).
			Where(arg_eq2).
			OrderBy("l").
			OrderBy("l").
			Limit(7).
			Offset(8).
			ToSql()

	}
}
