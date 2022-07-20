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


CREATE TABLE `ADDR_PREF_MST` (
  `PREF_CD` char(2) NOT NULL COMMENT '都道府県コード',
  `PREF_NM` varchar(10) NOT NULL COMMENT '都道府県名',
  `PREF_NMC` varchar(20) DEFAULT NULL,
  `ADDR_AREA_CD` char(2) NOT NULL COMMENT '住所エリアコード',
  `RAIL_AREA_CD` char(2) NOT NULL COMMENT '鉄道エリアコード',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`PREF_CD`),
  KEY `ADDR_PREF_MST_IDX_1` (`ADDR_AREA_CD`,`PREF_CD`) USING BTREE,
  KEY `ADDR_PREF_MST_IDX_2` (`RAIL_AREA_CD`,`PREF_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "pref_cd": "AReUZMBiqeuGpfSloCwYdHcfu",    "pref_nm": "qgVHwmcgUDdyAMlvtSWrjShro",    "pref_nmc": "pMDEYhOwKSEKaeGshFFJFcQGW",    "addr_area_cd": "VgeUPrvHYQMHFOuNQdQxOQfPM",    "rail_area_cd": "yKOSDuyXtAcNGYtssGfKSSWjL",    "dele_flg": "GdFowfviNKXTVyiLvodrcpXfP",    "upda_dte": "2063-12-24T15:44:24.315175801+09:00",    "upda_user_cd": "ehKMIHdOSHbaQHNkVyBQMYVuY",    "crea_dte": "2244-01-12T16:36:31.094440808+09:00",    "crea_user_cd": "UTWSuucoTAxvcxarcghkUNKtl"}



*/

// ADDR_PREF_MST_ struct is a row record of the ADDR_PREF_MST table in the anpidb database
type ADDR_PREF_MST_ struct {
	//[ 0] PREF_CD                                        char(2)              null: false  primary: true   isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD string `gorm:"primary_key;column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 1] PREF_NM                                        varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	PREFNM string `gorm:"column:PREF_NM;type:varchar;size:10;"` // 都道府県名
	//[ 2] PREF_NMC                                       varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	PREFNMC sql.NullString `gorm:"column:PREF_NMC;type:varchar;size:20;"`
	//[ 3] ADDR_AREA_CD                                   char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	ADDRAREACD string `gorm:"column:ADDR_AREA_CD;type:char;size:2;"` // 住所エリアコード
	//[ 4] RAIL_AREA_CD                                   char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	RAILAREACD string `gorm:"column:RAIL_AREA_CD;type:char;size:2;"` // 鉄道エリアコード
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

var ADDR_PREF_MSTTableInfo = &TableInfo{
	Name: "ADDR_PREF_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "string",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "PREF_NM",
			Comment:            `都道府県名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "PREFNM",
			GoFieldType:        "string",
			JSONFieldName:      "pref_nm",
			ProtobufFieldName:  "pref_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "PREF_NMC",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "PREFNMC",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pref_nmc",
			ProtobufFieldName:  "pref_nmc",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ADDR_AREA_CD",
			Comment:            `住所エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "ADDRAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "addr_area_cd",
			ProtobufFieldName:  "addr_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "RAIL_AREA_CD",
			Comment:            `鉄道エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "RAILAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "rail_area_cd",
			ProtobufFieldName:  "rail_area_cd",
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
func (a *ADDR_PREF_MST_) TableName() string {
	return "ADDR_PREF_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADDR_PREF_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADDR_PREF_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADDR_PREF_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADDR_PREF_MST_) TableInfo() *TableInfo {
	return ADDR_PREF_MSTTableInfo
}
