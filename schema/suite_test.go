package schema_test

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"

	"github.com/phogolabs/gom/schema"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSchema(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Schema Suite")
}

func CreateTable(reader *bytes.Buffer) string {
	query := &bytes.Buffer{}
	fmt.Fprintln(query, "CREATE TABLE test (")
	fmt.Fprintln(query, " char_field_null                      char(20) NULL,")
	fmt.Fprintln(query, " char_field_not_null                  char(20) NOT NULL,")
	fmt.Fprintln(query, " character_field_null                 character(20) NULL,")
	fmt.Fprintln(query, " character_field_not_null             character(20) NOT NULL,")
	fmt.Fprintln(query, " varchar_field_null                   varchar(20) NULL,")
	fmt.Fprintln(query, " varchar_field_not_null               varchar(20) NOT NULL,")
	fmt.Fprintln(query, " character_varying_field_null         character varying(20) NULL,")
	fmt.Fprintln(query, " character_varying_field_not_null     character varying(20) NOT NULL,")
	fmt.Fprintln(query, " text_field_null                      text NULL,")
	fmt.Fprintln(query, " text_field_not_null                  text NOT NULL,")
	fmt.Fprintln(query, " bit_field_null                       bit(20) NULL,")
	fmt.Fprintln(query, " bit_field_not_null                   bit(20) NOT NULL,")
	fmt.Fprintln(query, " smallint_field_null                  smallint NULL,")
	fmt.Fprintln(query, " smallint_field_not_null              smallint NOT NULL,")
	fmt.Fprintln(query, " int_field_null                       int NULL,")
	fmt.Fprintln(query, " int_field_not_null                   int NOT NULL,")
	fmt.Fprintln(query, " integer_field_null                   integer NULL,")
	fmt.Fprintln(query, " integer_field_not_null               integer NOT NULL,")
	fmt.Fprintln(query, " bigint_field_null                    bigint NULL,")
	fmt.Fprintln(query, " bigint_field_not_null                bigint NOT NULL,")
	fmt.Fprintln(query, " serial_field_not_null                serial NOT NULL,")
	fmt.Fprintln(query, " numeric_field_null                   numeric(20,20) NULL,")
	fmt.Fprintln(query, " numeric_field_not_null               numeric(20,20) NOT NULL,")
	fmt.Fprintln(query, " double_precision_field_null          double precision NULL,")
	fmt.Fprintln(query, " double_precision_field_not_null      double precision NOT NULL,")
	fmt.Fprintln(query, " real_field_null                      real NULL,")
	fmt.Fprintln(query, " real_field_not_null                  real NOT NULL,")
	fmt.Fprintln(query, " bool_field_null                      bool NULL,")
	fmt.Fprintln(query, " bool_field_not_null                  bool NOT NULL,")
	fmt.Fprintln(query, " boolean_field_null                   boolean NULL,")
	fmt.Fprintln(query, " boolean_field_not_null               boolean NOT NULL,")
	fmt.Fprintln(query, " date_field_null                      date NULL,")
	fmt.Fprintln(query, " date_field_not_null                  date NOT NULL,")
	fmt.Fprintln(query, " timestamp_field_null                 timestamp NULL,")
	fmt.Fprintln(query, " time_field_null                      time NULL,")
	fmt.Fprint(query, " time_field_not_null                  time NOT NULL")

	if reader.Len() > 0 {
		fmt.Fprintln(query, ",")
	} else {
		fmt.Fprintln(query)
	}

	io.Copy(query, reader)
	fmt.Fprintln(query, ")")

	return query.String()
}

func ExpectColumnsForPostgreSQL(columns []schema.Column) {
	column := columns[0]
	Expect(column.Name).To(Equal("char_field_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("bpchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[1]
	Expect(column.Name).To(Equal("char_field_not_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("bpchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[2]
	Expect(column.Name).To(Equal("character_field_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("bpchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[3]
	Expect(column.Name).To(Equal("character_field_not_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("bpchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[4]
	Expect(column.Name).To(Equal("varchar_field_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[5]
	Expect(column.Name).To(Equal("varchar_field_not_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[6]
	Expect(column.Name).To(Equal("character_varying_field_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[7]
	Expect(column.Name).To(Equal("character_varying_field_not_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[8]
	Expect(column.Name).To(Equal("text_field_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[9]
	Expect(column.Name).To(Equal("text_field_not_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[10]
	Expect(column.Name).To(Equal("bit_field_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[11]
	Expect(column.Name).To(Equal("bit_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[12]
	Expect(column.Name).To(Equal("smallint_field_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("int2"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(16))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int16"))

	column = columns[13]
	Expect(column.Name).To(Equal("smallint_field_not_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("int2"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(16))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int16"))

	column = columns[14]
	Expect(column.Name).To(Equal("int_field_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("int4"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(32))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[15]
	Expect(column.Name).To(Equal("int_field_not_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("int4"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(32))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[16]
	Expect(column.Name).To(Equal("integer_field_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("int4"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(32))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[17]
	Expect(column.Name).To(Equal("integer_field_not_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("int4"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(32))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[18]
	Expect(column.Name).To(Equal("bigint_field_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("int8"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(64))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int64"))

	column = columns[19]
	Expect(column.Name).To(Equal("bigint_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("int8"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(64))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int64"))

	column = columns[20]
	Expect(column.Name).To(Equal("serial_field_not_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("int4"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(32))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[21]
	Expect(column.Name).To(Equal("numeric_field_null"))
	Expect(column.Type.Name).To(Equal("numeric"))
	Expect(column.Type.Underlying).To(Equal("numeric"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("null.Float64"))

	column = columns[22]
	Expect(column.Name).To(Equal("numeric_field_not_null"))
	Expect(column.Type.Name).To(Equal("numeric"))
	Expect(column.Type.Underlying).To(Equal("numeric"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("float64"))

	column = columns[23]
	Expect(column.Name).To(Equal("double_precision_field_null"))
	Expect(column.Type.Name).To(Equal("double precision"))
	Expect(column.Type.Underlying).To(Equal("float8"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(53))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Float64"))

	column = columns[24]
	Expect(column.Name).To(Equal("double_precision_field_not_null"))
	Expect(column.Type.Name).To(Equal("double precision"))
	Expect(column.Type.Underlying).To(Equal("float8"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(53))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("float64"))

	column = columns[25]
	Expect(column.Name).To(Equal("real_field_null"))
	Expect(column.Type.Name).To(Equal("real"))
	Expect(column.Type.Underlying).To(Equal("float4"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(24))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Float32"))

	column = columns[26]
	Expect(column.Name).To(Equal("real_field_not_null"))
	Expect(column.Type.Name).To(Equal("real"))
	Expect(column.Type.Underlying).To(Equal("float4"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(24))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("float32"))

	column = columns[27]
	Expect(column.Name).To(Equal("bool_field_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Bool"))

	column = columns[28]
	Expect(column.Name).To(Equal("bool_field_not_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("bool"))

	column = columns[29]
	Expect(column.Name).To(Equal("boolean_field_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Bool"))

	column = columns[30]
	Expect(column.Name).To(Equal("boolean_field_not_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("bool"))

	column = columns[31]
	Expect(column.Name).To(Equal("date_field_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[32]
	Expect(column.Name).To(Equal("date_field_not_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[33]
	Expect(column.Name).To(Equal("timestamp_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[34]
	Expect(column.Name).To(Equal("time_field_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[35]
	Expect(column.Name).To(Equal("time_field_not_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[36]
	Expect(column.Name).To(Equal("varbit_field_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[37]
	Expect(column.Name).To(Equal("varbit_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[38]
	Expect(column.Name).To(Equal("bit_varying_field_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[39]
	Expect(column.Name).To(Equal("bit_varying_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[40]
	Expect(column.Name).To(Equal("smallserial_field_not_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("int2"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(16))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int16"))

	column = columns[41]
	Expect(column.Name).To(Equal("bigserial_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("int8"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(64))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int64"))

	column = columns[42]
	Expect(column.Name).To(Equal("money_field_null"))
	Expect(column.Type.Name).To(Equal("money"))
	Expect(column.Type.Underlying).To(Equal("money"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[43]
	Expect(column.Name).To(Equal("money_field_not_null"))
	Expect(column.Type.Name).To(Equal("money"))
	Expect(column.Type.Underlying).To(Equal("money"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[44]
	Expect(column.Name).To(Equal("timestamp_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[45]
	Expect(column.Name).To(Equal("timestamp_without_tz_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[46]
	Expect(column.Name).To(Equal("timestamp_without_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[47]
	Expect(column.Name).To(Equal("timestamp_with_tz_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp with time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamptz"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[48]
	Expect(column.Name).To(Equal("timestamp_with_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp with time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamptz"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[49]
	Expect(column.Name).To(Equal("time_without_tz_field_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[50]
	Expect(column.Name).To(Equal("time_without_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[51]
	Expect(column.Name).To(Equal("time_with_tz_field_null"))
	Expect(column.Type.Name).To(Equal("time with time zone"))
	Expect(column.Type.Underlying).To(Equal("timetz"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[52]
	Expect(column.Name).To(Equal("time_with_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("time with time zone"))
	Expect(column.Type.Underlying).To(Equal("timetz"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[53]
	Expect(column.Name).To(Equal("bytea_field_null"))
	Expect(column.Type.Name).To(Equal("bytea"))
	Expect(column.Type.Underlying).To(Equal("bytea"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Bytes"))

	column = columns[54]
	Expect(column.Name).To(Equal("bytea_field_not_null"))
	Expect(column.Type.Name).To(Equal("bytea"))
	Expect(column.Type.Underlying).To(Equal("bytea"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("[]byte"))

	column = columns[55]
	Expect(column.Name).To(Equal("jsonb_field_null"))
	Expect(column.Type.Name).To(Equal("jsonb"))
	Expect(column.Type.Underlying).To(Equal("jsonb"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.JSON"))

	column = columns[56]
	Expect(column.Name).To(Equal("jsonb_field_not_null"))
	Expect(column.Type.Name).To(Equal("jsonb"))
	Expect(column.Type.Underlying).To(Equal("jsonb"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("gom.JSON"))

	column = columns[57]
	Expect(column.Name).To(Equal("xml_field_null"))
	Expect(column.Type.Name).To(Equal("xml"))
	Expect(column.Type.Underlying).To(Equal("xml"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[58]
	Expect(column.Name).To(Equal("xml_field_not_null"))
	Expect(column.Type.Name).To(Equal("xml"))
	Expect(column.Type.Underlying).To(Equal("xml"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[59]
	Expect(column.Name).To(Equal("uuid_field_null"))
	Expect(column.Type.Name).To(Equal("uuid"))
	Expect(column.Type.Underlying).To(Equal("uuid"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("uuid.NullUUID"))

	column = columns[60]
	Expect(column.Name).To(Equal("uuid_field_not_null"))
	Expect(column.Type.Name).To(Equal("uuid"))
	Expect(column.Type.Underlying).To(Equal("uuid"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("uuid.UUID"))

	column = columns[61]
	Expect(column.Name).To(Equal("hstore_field_null"))
	Expect(column.Type.Name).To(Equal("USER-DEFINED"))
	Expect(column.Type.Underlying).To(Equal("hstore"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[62]
	Expect(column.Name).To(Equal("hstore_field_not_null"))
	Expect(column.Type.Name).To(Equal("USER-DEFINED"))
	Expect(column.Type.Underlying).To(Equal("hstore"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))
}

func ExpectColumnsForMySQL(columns []schema.Column) {
	column := columns[0]
	Expect(column.Name).To(Equal("char_field_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Byte"))

	column = columns[1]
	Expect(column.Name).To(Equal("char_field_not_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("byte"))

	column = columns[2]
	Expect(column.Name).To(Equal("character_field_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Byte"))

	column = columns[3]
	Expect(column.Name).To(Equal("character_field_not_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("byte"))

	column = columns[4]
	Expect(column.Name).To(Equal("varchar_field_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[5]
	Expect(column.Name).To(Equal("varchar_field_not_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[6]
	Expect(column.Name).To(Equal("character_varying_field_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[7]
	Expect(column.Name).To(Equal("character_varying_field_not_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(20))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[8]
	Expect(column.Name).To(Equal("text_field_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(65535))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[9]
	Expect(column.Name).To(Equal("text_field_not_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(65535))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[10]
	Expect(column.Name).To(Equal("bit_field_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[11]
	Expect(column.Name).To(Equal("bit_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[12]
	Expect(column.Name).To(Equal("smallint_field_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("smallint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(5))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int16"))

	column = columns[13]
	Expect(column.Name).To(Equal("smallint_field_not_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("smallint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(5))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int16"))

	column = columns[14]
	Expect(column.Name).To(Equal("int_field_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(10))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[15]
	Expect(column.Name).To(Equal("int_field_not_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(10))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[16]
	Expect(column.Name).To(Equal("integer_field_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(10))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[17]
	Expect(column.Name).To(Equal("integer_field_not_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(10))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[18]
	Expect(column.Name).To(Equal("bigint_field_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("bigint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(19))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int64"))

	column = columns[19]
	Expect(column.Name).To(Equal("bigint_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("bigint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(19))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int64"))

	column = columns[20]
	Expect(column.Name).To(Equal("serial_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("bigint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("uint64"))

	column = columns[21]
	Expect(column.Name).To(Equal("numeric_field_null"))
	Expect(column.Type.Name).To(Equal("decimal"))
	Expect(column.Type.Underlying).To(Equal("decimal"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("null.Float64"))

	column = columns[22]
	Expect(column.Name).To(Equal("numeric_field_not_null"))
	Expect(column.Type.Name).To(Equal("decimal"))
	Expect(column.Type.Underlying).To(Equal("decimal"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("float64"))

	column = columns[23]
	Expect(column.Name).To(Equal("double_precision_field_null"))
	Expect(column.Type.Name).To(Equal("double"))
	Expect(column.Type.Underlying).To(Equal("double"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(22))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[24]
	Expect(column.Name).To(Equal("double_precision_field_not_null"))
	Expect(column.Type.Name).To(Equal("double"))
	Expect(column.Type.Underlying).To(Equal("double"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(22))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[25]
	Expect(column.Name).To(Equal("real_field_null"))
	Expect(column.Type.Name).To(Equal("double"))
	Expect(column.Type.Underlying).To(Equal("double"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(22))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[26]
	Expect(column.Name).To(Equal("real_field_not_null"))
	Expect(column.Type.Name).To(Equal("double"))
	Expect(column.Type.Underlying).To(Equal("double"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(22))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[27]
	Expect(column.Name).To(Equal("bool_field_null"))
	Expect(column.Type.Name).To(Equal("tinyint"))
	Expect(column.Type.Underlying).To(Equal("tinyint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(3))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int8"))

	column = columns[28]
	Expect(column.Name).To(Equal("bool_field_not_null"))
	Expect(column.Type.Name).To(Equal("tinyint"))
	Expect(column.Type.Underlying).To(Equal("tinyint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(3))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int8"))

	column = columns[29]
	Expect(column.Name).To(Equal("boolean_field_null"))
	Expect(column.Type.Name).To(Equal("tinyint"))
	Expect(column.Type.Underlying).To(Equal("tinyint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(3))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int8"))

	column = columns[30]
	Expect(column.Name).To(Equal("boolean_field_not_null"))
	Expect(column.Type.Name).To(Equal("tinyint"))
	Expect(column.Type.Underlying).To(Equal("tinyint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(3))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int8"))

	column = columns[31]
	Expect(column.Name).To(Equal("date_field_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[32]
	Expect(column.Name).To(Equal("date_field_not_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[33]
	Expect(column.Name).To(Equal("timestamp_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[34]
	Expect(column.Name).To(Equal("time_field_null"))
	Expect(column.Type.Name).To(Equal("time"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[35]
	Expect(column.Name).To(Equal("time_field_not_null"))
	Expect(column.Type.Name).To(Equal("time"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))
}

func ExpectColumnsForSQLite(columns []schema.Column) {
	column := columns[0]
	Expect(column.Name).To(Equal("char_field_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Byte"))

	column = columns[1]
	Expect(column.Name).To(Equal("char_field_not_null"))
	Expect(column.Type.Name).To(Equal("char"))
	Expect(column.Type.Underlying).To(Equal("char"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("byte"))

	column = columns[2]
	Expect(column.Name).To(Equal("character_field_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("character"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[3]
	Expect(column.Name).To(Equal("character_field_not_null"))
	Expect(column.Type.Name).To(Equal("character"))
	Expect(column.Type.Underlying).To(Equal("character"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[4]
	Expect(column.Name).To(Equal("varchar_field_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[5]
	Expect(column.Name).To(Equal("varchar_field_not_null"))
	Expect(column.Type.Name).To(Equal("varchar"))
	Expect(column.Type.Underlying).To(Equal("varchar"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[6]
	Expect(column.Name).To(Equal("character_varying_field_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("character varying"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[7]
	Expect(column.Name).To(Equal("character_varying_field_not_null"))
	Expect(column.Type.Name).To(Equal("character varying"))
	Expect(column.Type.Underlying).To(Equal("character varying"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[8]
	Expect(column.Name).To(Equal("text_field_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[9]
	Expect(column.Name).To(Equal("text_field_not_null"))
	Expect(column.Type.Name).To(Equal("text"))
	Expect(column.Type.Underlying).To(Equal("text"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[10]
	Expect(column.Name).To(Equal("bit_field_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[11]
	Expect(column.Name).To(Equal("bit_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit"))
	Expect(column.Type.Underlying).To(Equal("bit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[12]
	Expect(column.Name).To(Equal("smallint_field_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("smallint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int16"))

	column = columns[13]
	Expect(column.Name).To(Equal("smallint_field_not_null"))
	Expect(column.Type.Name).To(Equal("smallint"))
	Expect(column.Type.Underlying).To(Equal("smallint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int16"))

	column = columns[14]
	Expect(column.Name).To(Equal("int_field_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[15]
	Expect(column.Name).To(Equal("int_field_not_null"))
	Expect(column.Type.Name).To(Equal("int"))
	Expect(column.Type.Underlying).To(Equal("int"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[16]
	Expect(column.Name).To(Equal("integer_field_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("integer"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int"))

	column = columns[17]
	Expect(column.Name).To(Equal("integer_field_not_null"))
	Expect(column.Type.Name).To(Equal("integer"))
	Expect(column.Type.Underlying).To(Equal("integer"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[18]
	Expect(column.Name).To(Equal("bigint_field_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("bigint"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Int64"))

	column = columns[19]
	Expect(column.Name).To(Equal("bigint_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigint"))
	Expect(column.Type.Underlying).To(Equal("bigint"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int64"))

	column = columns[20]
	Expect(column.Name).To(Equal("serial_field_not_null"))
	Expect(column.Type.Name).To(Equal("serial"))
	Expect(column.Type.Underlying).To(Equal("serial"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int"))

	column = columns[21]
	Expect(column.Name).To(Equal("numeric_field_null"))
	Expect(column.Type.Name).To(Equal("numeric"))
	Expect(column.Type.Underlying).To(Equal("numeric"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("null.Float64"))

	column = columns[22]
	Expect(column.Name).To(Equal("numeric_field_not_null"))
	Expect(column.Type.Name).To(Equal("numeric"))
	Expect(column.Type.Underlying).To(Equal("numeric"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(20))
	Expect(column.ScanType).To(Equal("float64"))

	column = columns[23]
	Expect(column.Name).To(Equal("double_precision_field_null"))
	Expect(column.Type.Name).To(Equal("double precision"))
	Expect(column.Type.Underlying).To(Equal("double precision"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Float64"))

	column = columns[24]
	Expect(column.Name).To(Equal("double_precision_field_not_null"))
	Expect(column.Type.Name).To(Equal("double precision"))
	Expect(column.Type.Underlying).To(Equal("double precision"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("float64"))

	column = columns[25]
	Expect(column.Name).To(Equal("real_field_null"))
	Expect(column.Type.Name).To(Equal("real"))
	Expect(column.Type.Underlying).To(Equal("real"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Float32"))

	column = columns[26]
	Expect(column.Name).To(Equal("real_field_not_null"))
	Expect(column.Type.Name).To(Equal("real"))
	Expect(column.Type.Underlying).To(Equal("real"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("float32"))

	column = columns[27]
	Expect(column.Name).To(Equal("bool_field_null"))
	Expect(column.Type.Name).To(Equal("bool"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[28]
	Expect(column.Name).To(Equal("bool_field_not_null"))
	Expect(column.Type.Name).To(Equal("bool"))
	Expect(column.Type.Underlying).To(Equal("bool"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[29]
	Expect(column.Name).To(Equal("boolean_field_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("boolean"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Bool"))

	column = columns[30]
	Expect(column.Name).To(Equal("boolean_field_not_null"))
	Expect(column.Type.Name).To(Equal("boolean"))
	Expect(column.Type.Underlying).To(Equal("boolean"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("bool"))

	column = columns[31]
	Expect(column.Name).To(Equal("date_field_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[32]
	Expect(column.Name).To(Equal("date_field_not_null"))
	Expect(column.Type.Name).To(Equal("date"))
	Expect(column.Type.Underlying).To(Equal("date"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[33]
	Expect(column.Name).To(Equal("timestamp_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[34]
	Expect(column.Name).To(Equal("time_field_null"))
	Expect(column.Type.Name).To(Equal("time"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[35]
	Expect(column.Name).To(Equal("time_field_not_null"))
	Expect(column.Type.Name).To(Equal("time"))
	Expect(column.Type.Underlying).To(Equal("time"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[36]
	Expect(column.Name).To(Equal("varbit_field_null"))
	Expect(column.Type.Name).To(Equal("varbit"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[37]
	Expect(column.Name).To(Equal("varbit_field_not_null"))
	Expect(column.Type.Name).To(Equal("varbit"))
	Expect(column.Type.Underlying).To(Equal("varbit"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[38]
	Expect(column.Name).To(Equal("bit_varying_field_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("bit varying"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[39]
	Expect(column.Name).To(Equal("bit_varying_field_not_null"))
	Expect(column.Type.Name).To(Equal("bit varying"))
	Expect(column.Type.Underlying).To(Equal("bit varying"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(20))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[40]
	Expect(column.Name).To(Equal("smallserial_field_not_null"))
	Expect(column.Type.Name).To(Equal("smallserial"))
	Expect(column.Type.Underlying).To(Equal("smallserial"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int16"))

	column = columns[41]
	Expect(column.Name).To(Equal("bigserial_field_not_null"))
	Expect(column.Type.Name).To(Equal("bigserial"))
	Expect(column.Type.Underlying).To(Equal("bigserial"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("int64"))

	column = columns[42]
	Expect(column.Name).To(Equal("money_field_null"))
	Expect(column.Type.Name).To(Equal("money"))
	Expect(column.Type.Underlying).To(Equal("money"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[43]
	Expect(column.Name).To(Equal("money_field_not_null"))
	Expect(column.Type.Name).To(Equal("money"))
	Expect(column.Type.Underlying).To(Equal("money"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[44]
	Expect(column.Name).To(Equal("timestamp_without_tz_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp without time zone"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[45]
	Expect(column.Name).To(Equal("timestamp_without_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp without time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp without time zone"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[46]
	Expect(column.Name).To(Equal("timestamp_with_tz_field_null"))
	Expect(column.Type.Name).To(Equal("timestamp with time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp with time zone"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[47]
	Expect(column.Name).To(Equal("timestamp_with_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp with time zone"))
	Expect(column.Type.Underlying).To(Equal("timestamp with time zone"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[48]
	Expect(column.Name).To(Equal("time_without_tz_field_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time without time zone"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[49]
	Expect(column.Name).To(Equal("time_without_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("time without time zone"))
	Expect(column.Type.Underlying).To(Equal("time without time zone"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[50]
	Expect(column.Name).To(Equal("time_with_tz_field_null"))
	Expect(column.Type.Name).To(Equal("time with time zone"))
	Expect(column.Type.Underlying).To(Equal("time with time zone"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Time"))

	column = columns[51]
	Expect(column.Name).To(Equal("time_with_tz_field_not_null"))
	Expect(column.Type.Name).To(Equal("time with time zone"))
	Expect(column.Type.Underlying).To(Equal("time with time zone"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))

	column = columns[52]
	Expect(column.Name).To(Equal("bytea_field_null"))
	Expect(column.Type.Name).To(Equal("bytea"))
	Expect(column.Type.Underlying).To(Equal("bytea"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.Bytes"))

	column = columns[53]
	Expect(column.Name).To(Equal("bytea_field_not_null"))
	Expect(column.Type.Name).To(Equal("bytea"))
	Expect(column.Type.Underlying).To(Equal("bytea"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("[]byte"))

	column = columns[54]
	Expect(column.Name).To(Equal("jsonb_field_null"))
	Expect(column.Type.Name).To(Equal("jsonb"))
	Expect(column.Type.Underlying).To(Equal("jsonb"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.JSON"))

	column = columns[55]
	Expect(column.Name).To(Equal("jsonb_field_not_null"))
	Expect(column.Type.Name).To(Equal("jsonb"))
	Expect(column.Type.Underlying).To(Equal("jsonb"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("gom.JSON"))

	column = columns[56]
	Expect(column.Name).To(Equal("xml_field_null"))
	Expect(column.Type.Name).To(Equal("xml"))
	Expect(column.Type.Underlying).To(Equal("xml"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("null.String"))

	column = columns[57]
	Expect(column.Name).To(Equal("xml_field_not_null"))
	Expect(column.Type.Name).To(Equal("xml"))
	Expect(column.Type.Underlying).To(Equal("xml"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("string"))

	column = columns[58]
	Expect(column.Name).To(Equal("uuid_field_null"))
	Expect(column.Type.Name).To(Equal("uuid"))
	Expect(column.Type.Underlying).To(Equal("uuid"))
	Expect(column.Type.IsNullable).To(Equal(true))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("uuid.NullUUID"))

	column = columns[59]
	Expect(column.Name).To(Equal("uuid_field_not_null"))
	Expect(column.Type.Name).To(Equal("uuid"))
	Expect(column.Type.Underlying).To(Equal("uuid"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("uuid.UUID"))

	column = columns[60]
	Expect(column.Name).To(Equal("timestamp_field_not_null"))
	Expect(column.Type.Name).To(Equal("timestamp"))
	Expect(column.Type.Underlying).To(Equal("timestamp"))
	Expect(column.Type.IsNullable).To(Equal(false))
	Expect(column.Type.CharMaxLength).To(Equal(0))
	Expect(column.Type.Precision).To(Equal(0))
	Expect(column.Type.PrecisionScale).To(Equal(0))
	Expect(column.ScanType).To(Equal("time.Time"))
}
