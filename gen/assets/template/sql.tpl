
{{$tableName := LowerCase .Name}}
DROP TABLE IF EXISTS "public"."tbl_{{$tableName}}s";
CREATE TABLE "public"."tbl_{{$tableName}}s" (
	"id" uuid NOT NULL,
	"data" jsonb NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."tbl_{{$tableName}}s" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table tbl_{{$tableName}}s
-- ----------------------------
ALTER TABLE "public"."tbl_{{$tableName}}s" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;