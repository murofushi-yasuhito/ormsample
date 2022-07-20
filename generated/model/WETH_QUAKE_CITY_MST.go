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


CREATE TABLE `WETH_QUAKE_CITY_MST` (
  `WETH_QUAKE_CITY_CD` char(7) NOT NULL COMMENT '震度観測点市区町村コード	 2013/01桁数変更',
  `WETH_QUAKE_CITY_NMJ` varchar(40) NOT NULL COMMENT '震度観測点市区町村名',
  `WETH_QUAKE_CITY_NMC` varchar(40) DEFAULT NULL COMMENT '震度観測点市区町村名カナ',
  `WETH_QUAKE_AREA_CD` char(3) NOT NULL COMMENT '震度観測点地域コード	 2013/01桁数変更',
  `CITY_CD` char(5) NOT NULL COMMENT '市区町村コード',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`WETH_QUAKE_CITY_CD`),
  KEY `WETH_QUAKE_CITY_MST_IDX_1` (`WETH_QUAKE_AREA_CD`,`WETH_QUAKE_CITY_CD`) USING BTREE,
  KEY `WETH_QUAKE_CITY_MST_IDX_2` (`CITY_CD`,`WETH_QUAKE_CITY_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "weth_quake_city_cd": "ZHHwgjFsddAyymZixlcMovFZh",    "weth_quake_city_nmj": "wCgYqsHhmualZSCgWlKbGxooW",    "weth_quake_city_nmc": "OnYZMbFrcFuhicIyVGkgISiWI",    "weth_quake_area_cd": "lPnfcDhQBtBNvjEAauQUEwxoL",    "city_cd": "hTaaiZYXDXXIFTANhYmaOEpZu",    "dele_flg": "tfPCWCKBPtSJyfCVcXLHrIrrG",    "upda_dte": "2213-07-10T13:05:39.018505908+09:00",    "upda_user_cd": "pHEDpkYpBFIkEYCQpHYWMAKQP",    "crea_dte": "2169-10-27T22:28:32.072537994+09:00",    "crea_user_cd": "UHEZFfEiXijOtaBCNQpMcCHVD"}



*/

// WETH_QUAKE_CITY_MST_ struct is a row record of the WETH_QUAKE_CITY_MST table in the anpidb database
type WETH_QUAKE_CITY_MST_ struct {
	//[ 0] WETH_QUAKE_CITY_CD                             char(7)              null: false  primary: true   isArray: false  auto: false  col: char            len: 7       default: []
	WETHQUAKECITYCD string `gorm:"primary_key;column:WETH_QUAKE_CITY_CD;type:char;size:7;"` // 震度観測点市区町村コード	 2013/01桁数変更
	//[ 1] WETH_QUAKE_CITY_NMJ                            varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WETHQUAKECITYNMJ string `gorm:"column:WETH_QUAKE_CITY_NMJ;type:varchar;size:40;"` // 震度観測点市区町村名
	//[ 2] WETH_QUAKE_CITY_NMC                            varchar(40)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WETHQUAKECITYNMC sql.NullString `gorm:"column:WETH_QUAKE_CITY_NMC;type:varchar;size:40;"` // 震度観測点市区町村名カナ
	//[ 3] WETH_QUAKE_AREA_CD                             char(3)              null: false  primary: false  isArray: false  auto: false  col: char            len: 3       default: []
	WETHQUAKEAREACD string `gorm:"column:WETH_QUAKE_AREA_CD;type:char;size:3;"` // 震度観測点地域コード	 2013/01桁数変更
	//[ 4] CITY_CD                                        char(5)              null: false  primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD string `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[ 5] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 7] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"`
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"`
}

var WETH_QUAKE_CITY_MSTTableInfo = &TableInfo{
	Name: "WETH_QUAKE_CITY_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index: 0,
			Name:  "WETH_QUAKE_CITY_CD",
			Comment: `震度観測点市区町村コード	 2013/01桁数変更`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "WETHQUAKECITYCD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_quake_city_cd",
			ProtobufFieldName:  "weth_quake_city_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WETH_QUAKE_CITY_NMJ",
			Comment:            `震度観測点市区町村名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WETHQUAKECITYNMJ",
			GoFieldType:        "string",
			JSONFieldName:      "weth_quake_city_nmj",
			ProtobufFieldName:  "weth_quake_city_nmj",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "WETH_QUAKE_CITY_NMC",
			Comment:            `震度観測点市区町村名カナ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WETHQUAKECITYNMC",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_quake_city_nmc",
			ProtobufFieldName:  "weth_quake_city_nmc",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "WETH_QUAKE_AREA_CD",
			Comment: `震度観測点地域コード	 2013/01桁数変更`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHQUAKEAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_quake_area_cd",
			ProtobufFieldName:  "weth_quake_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "string",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "UPDA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CREA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_QUAKE_CITY_MST_) TableName() string {
	return "WETH_QUAKE_CITY_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_QUAKE_CITY_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_QUAKE_CITY_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_QUAKE_CITY_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_QUAKE_CITY_MST_) TableInfo() *TableInfo {
	return WETH_QUAKE_CITY_MSTTableInfo
}
