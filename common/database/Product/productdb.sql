
DROP TABLE IF EXISTS "public"."tbl_product";
CREATE TABLE "public"."tbl_product" (
	"id" uuid NOT NULL,
	"data" jsonb NOT NULL
)
WITH (OIDS=FALSE);
ALTER TABLE "public"."tbl_product" OWNER TO "postgres";

-- ----------------------------
--  Primary key structure for table tbl_product
-- ----------------------------
ALTER TABLE "public"."tbl_product" ADD PRIMARY KEY ("id") NOT DEFERRABLE INITIALLY IMMEDIATE;
