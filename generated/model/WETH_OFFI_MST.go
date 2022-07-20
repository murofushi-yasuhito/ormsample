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


CREATE TABLE `WETH_OFFI_MST` (
  `WETH_OFFI_CD` char(4) NOT NULL COMMENT '気象官署コード',
  `WETH_OFFI_NM` varchar(40) NOT NULL COMMENT '気象官署名',
  `WETH_OBSE_CD` char(3) NOT NULL COMMENT '気象台コード（管区気象台）',
  `PREF_CD` char(2) NOT NULL COMMENT '都道府県コード',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`WETH_OFFI_CD`),
  KEY `WETH_OFFI_MST_IDX_1` (`WETH_OBSE_CD`,`WETH_OFFI_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "weth_offi_cd": "ftwtHDaeEqCdeRuHhZlAmMjYs",    "weth_offi_nm": "dWmNZvpolNCejOoUVjinBWqUJ",    "weth_obse_cd": "TEGskBDdKbhDLQwTnOBsBlGTQ",    "pref_cd": "PZVajutCUupOudBTtetkUcZSg",    "dele_flg": "AxqqmcTXbrnyWvtcbegWkCdaD",    "upda_dte": "2167-01-19T15:39:55.963208193+09:00",    "upda_user_cd": "oxMVDwhYEBLbPWdbyIWoxwDJL",    "crea_dte": "2164-05-12T20:08:30.851360871+09:00",    "crea_user_cd": "oHMXJerwBSOyktrVnlUVmfBoE"}



*/

// WETH_OFFI_MST_ struct is a row record of the WETH_OFFI_MST table in the anpidb database
type WETH_OFFI_MST_ struct {
	//[ 0] WETH_OFFI_CD                                   char(4)              null: false  primary: true   isArray: false  auto: false  col: char            len: 4       default: []
	WETHOFFICD string `gorm:"primary_key;column:WETH_OFFI_CD;type:char;size:4;"` // 気象官署コード
	//[ 1] WETH_OFFI_NM                                   varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WETHOFFINM string `gorm:"column:WETH_OFFI_NM;type:varchar;size:40;"` // 気象官署名
	//[ 2] WETH_OBSE_CD                                   char(3)              null: false  primary: false  isArray: false  auto: false  col: char            len: 3       default: []
	WETHOBSECD string `gorm:"column:WETH_OBSE_CD;type:char;size:3;"` // 気象台コード（管区気象台）
	//[ 3] PREF_CD                                        char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD string `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 4] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"`
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"`
}

var WETH_OFFI_MSTTableInfo = &TableInfo{
	Name: "WETH_OFFI_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "WETH_OFFI_CD",
			Comment:            `気象官署コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "WETHOFFICD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_offi_cd",
			ProtobufFieldName:  "weth_offi_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WETH_OFFI_NM",
			Comment:            `気象官署名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WETHOFFINM",
			GoFieldType:        "string",
			JSONFieldName:      "weth_offi_nm",
			ProtobufFieldName:  "weth_offi_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "WETH_OBSE_CD",
			Comment:            `気象台コード（管区気象台）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHOBSECD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_obse_cd",
			ProtobufFieldName:  "weth_obse_cd",
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_OFFI_MST_) TableName() string {
	return "WETH_OFFI_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_OFFI_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_OFFI_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_OFFI_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_OFFI_MST_) TableInfo() *TableInfo {
	return WETH_OFFI_MSTTableInfo
}
