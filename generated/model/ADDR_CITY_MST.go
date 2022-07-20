package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `ADDR_CITY_MST` (
  `CITY_CD` char(5) NOT NULL COMMENT '市区町村コード',
  `CITY_NMJ` varchar(40) NOT NULL COMMENT '市区町村名',
  `CITY_NMC` varchar(40) NOT NULL COMMENT '市区町村名カナ',
  `PREF_CD` char(2) NOT NULL COMMENT '都道府県コード',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` char(8) NOT NULL COMMENT '更新者コード',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` char(8) NOT NULL COMMENT '作成者コード',
  PRIMARY KEY (`CITY_CD`),
  KEY `ADDR_CITY_MST_IDX_1` (`PREF_CD`,`CITY_CD`) USING BTREE,
  KEY `ADDR_CITY_MST_IDX_2` (`PREF_CD`,`CITY_NMJ`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "city_cd": "ykuXyJeBHKrCiBQKBjrvRqBhe",    "city_nmj": "wwNmgbCKDXWJGuUeWpuPYYjqB",    "city_nmc": "JsxAPYqIoqifnAUaXwBhbKGeY",    "pref_cd": "ovOSYFQXbolxpLuFYnlyvXNxu",    "dele_flg": "qxskoYfeMaUHncxugxDZMlutB",    "upda_dte": "2059-01-23T18:01:49.084715025+09:00",    "upda_user_cd": "arhJehhhmmktnvsJZWsJwqAaI",    "crea_dte": "2085-01-10T19:54:01.551031415+09:00",    "crea_user_cd": "mGOPbOAAwhlnXUTTTpSOGtPuD"}



*/

// ADDR_CITY_MST_ struct is a row record of the ADDR_CITY_MST table in the anpidb database
type ADDR_CITY_MST_ struct {
	//[ 0] CITY_CD                                        char(5)              null: false  primary: true   isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD string `gorm:"primary_key;column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[ 1] CITY_NMJ                                       varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	CITYNMJ string `gorm:"column:CITY_NMJ;type:varchar;size:40;"` // 市区町村名
	//[ 2] CITY_NMC                                       varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	CITYNMC string `gorm:"column:CITY_NMC;type:varchar;size:40;"` // 市区町村名カナ
	//[ 3] PREF_CD                                        char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD string `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 4] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_CD                                   char(8)              null: false  primary: false  isArray: false  auto: false  col: char            len: 8       default: []
	UPDAUSERCD string `gorm:"column:UPDA_USER_CD;type:char;size:8;"` // 更新者コード
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_CD                                   char(8)              null: false  primary: false  isArray: false  auto: false  col: char            len: 8       default: []
	CREAUSERCD string `gorm:"column:CREA_USER_CD;type:char;size:8;"` // 作成者コード

}

var ADDR_CITY_MSTTableInfo = &TableInfo{
	Name: "ADDR_CITY_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "string",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CITY_NMJ",
			Comment:            `市区町村名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "CITYNMJ",
			GoFieldType:        "string",
			JSONFieldName:      "city_nmj",
			ProtobufFieldName:  "city_nmj",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CITY_NMC",
			Comment:            `市区町村名カナ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "CITYNMC",
			GoFieldType:        "string",
			JSONFieldName:      "city_nmc",
			ProtobufFieldName:  "city_nmc",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "string",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "UPDA_DTE",
			Comment:            `更新日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UPDADTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "upda_dte",
			ProtobufFieldName:  "upda_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "UPDA_USER_CD",
			Comment:            `更新者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(8)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       8,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CREA_DTE",
			Comment:            `作成日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CREADTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "crea_dte",
			ProtobufFieldName:  "crea_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CREA_USER_CD",
			Comment:            `作成者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(8)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       8,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ADDR_CITY_MST_) TableName() string {
	return "ADDR_CITY_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADDR_CITY_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADDR_CITY_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADDR_CITY_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADDR_CITY_MST_) TableInfo() *TableInfo {
	return ADDR_CITY_MSTTableInfo
}
