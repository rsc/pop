package translators_test

import (
	"testing"

	"github.com/markbates/pop/fizz"
	"github.com/markbates/pop/fizz/translators"
	"github.com/stretchr/testify/require"
)

var _ fizz.Translator = (*translators.Postgres)(nil)
var pgt = translators.Postgres{}

func Test_Postgres_CreateTable(t *testing.T) {
	r := require.New(t)
	ddl := `CREATE TABLE IF NOT EXISTS "users" (
"id" SERIAL PRIMARY KEY,
"created_at" timestamp NOT NULL,
"updated_at" timestamp NOT NULL,
"first_name" VARCHAR (255) NOT NULL,
"last_name" VARCHAR (255) NOT NULL,
"email" VARCHAR (20) NOT NULL,
"permissions" jsonb,
"age" integer DEFAULT '40'
);`

	res, _ := fizz.AString(`
	create_table("users", func(t) {
		t.Column("first_name", "string", {})
		t.Column("last_name", "string", {})
		t.Column("email", "string", {"size":20})
		t.Column("permissions", "jsonb", {"null": true})
		t.Column("age", "integer", {"null": true, "default": 40})
	})
	`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_DropTable(t *testing.T) {
	r := require.New(t)

	ddl := `DROP TABLE IF EXISTS "users";`

	res, _ := fizz.AString(`drop_table("users")`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_RenameTable(t *testing.T) {
	r := require.New(t)

	ddl := `ALTER TABLE "users" RENAME TO "people";`

	res, _ := fizz.AString(`rename_table("users", "people")`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_RenameTable_NotEnoughValues(t *testing.T) {
	r := require.New(t)

	_, err := pgt.RenameTable([]fizz.Table{})
	r.Error(err)
}

func Test_Postgres_AddColumn(t *testing.T) {
	r := require.New(t)
	ddl := `ALTER TABLE "mytable" ADD COLUMN "mycolumn" VARCHAR (50) NOT NULL DEFAULT 'foo';`

	res, _ := fizz.AString(`add_column("mytable", "mycolumn", "string", {"default": "foo", "size": 50})`, pgt)

	r.Equal(ddl, res)
}

func Test_Postgres_DropColumn(t *testing.T) {
	r := require.New(t)
	ddl := `ALTER TABLE "table_name" DROP COLUMN "column_name";`

	res, _ := fizz.AString(`drop_column("table_name", "column_name")`, pgt)

	r.Equal(ddl, res)
}

func Test_Postgres_RenameColumn(t *testing.T) {
	r := require.New(t)
	ddl := `ALTER TABLE "table_name" RENAME COLUMN "old_column" TO "new_column";`

	res, _ := fizz.AString(`rename_column("table_name", "old_column", "new_column")`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_AddIndex(t *testing.T) {
	r := require.New(t)
	ddl := `CREATE INDEX "table_name_column_name_idx" ON "table_name" (column_name);`

	res, _ := fizz.AString(`add_index("table_name", "column_name", {})`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_AddIndex_Unique(t *testing.T) {
	r := require.New(t)
	ddl := `CREATE UNIQUE INDEX "table_name_column_name_idx" ON "table_name" (column_name);`

	res, _ := fizz.AString(`add_index("table_name", "column_name", {"unique": true})`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_AddIndex_MultiColumn(t *testing.T) {
	r := require.New(t)
	ddl := `CREATE INDEX "table_name_col1_col2_col3_idx" ON "table_name" (col1, col2, col3);`

	res, _ := fizz.AString(`add_index("table_name", ["col1", "col2", "col3"], {})`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_AddIndex_CustomName(t *testing.T) {
	r := require.New(t)
	ddl := `CREATE INDEX "custom_name" ON "table_name" (column_name);`

	res, _ := fizz.AString(`add_index("table_name", "column_name", {"name": "custom_name"})`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_DropIndex(t *testing.T) {
	r := require.New(t)
	ddl := `DROP INDEX IF EXISTS "my_idx";`

	res, _ := fizz.AString(`drop_index("my_idx")`, pgt)
	r.Equal(ddl, res)
}

func Test_Postgres_RenameIndex(t *testing.T) {
	r := require.New(t)

	ddl := `ALTER INDEX "old_ix" RENAME TO "new_ix";`

	res, _ := fizz.AString(`rename_index("old_ix", "new_ix")`, pgt)
	r.Equal(ddl, res)
}
