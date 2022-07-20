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


CREATE TABLE `RAIL_AREA_MST` (
  `RAIL_AREA_CD` char(2) NOT NULL COMMENT '鉄道エリアコード',
  `RAIL_AREA_NM` varchar(40) NOT NULL COMMENT '鉄道エリア名',
  `ABRO_KBN_CD` char(1) NOT NULL COMMENT '海外区分コード',
  `DISP_DSQ` int(10) NOT NULL COMMENT '表示順',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_CD` char(8) NOT NULL COMMENT '作成者コード',
  `UPDA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '更新者コード',
  `UPDA_USER_CD` char(8) NOT NULL COMMENT '更新者コード',
  PRIMARY KEY (`RAIL_AREA_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "rail_area_cd": "vEfKKcLrbkSXHESIEfjaaQStH",    "rail_area_nm": "yTSZZTtwyTLpMplCaooKKUgJH",    "abro_kbn_cd": "yKRLRxTCaTrxPDFWhCUjHmMMj",    "disp_dsq": 56,    "dele_flg": "pbruruspCXHVtRRZdtVTbwEpT",    "crea_dte": "2284-07-01T03:51:04.376200411+09:00",    "crea_user_cd": "MiCfJDfMBbbhVRdkVaLSreynx",    "upda_dte": "2167-10-16T04:26:40.355406014+09:00",    "upda_user_cd": "vwMUaiBnvGaZfAKwPRLRiMKgl"}



*/

// RAIL_AREA_MST_ struct is a row record of the RAIL_AREA_MST table in the anpidb database
type RAIL_AREA_MST_ struct {
	//[ 0] RAIL_AREA_CD                                   char(2)              null: false  primary: true   isArray: false  auto: false  col: char            len: 2       default: []
	RAILAREACD string `gorm:"primary_key;column:RAIL_AREA_CD;type:char;size:2;"` // 鉄道エリアコード
	//[ 1] RAIL_AREA_NM                                   varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	RAILAREANM string `gorm:"column:RAIL_AREA_NM;type:varchar;size:40;"` // 鉄道エリア名
	//[ 2] ABRO_KBN_CD                                    char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	ABROKBNCD string `gorm:"column:ABRO_KBN_CD;type:char;size:1;"` // 海外区分コード
	//[ 3] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 4] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 6] CREA_USER_CD                                   char(8)              null: false  primary: false  isArray: false  auto: false  col: char            len: 8       default: []
	CREAUSERCD string `gorm:"column:CREA_USER_CD;type:char;size:8;"` // 作成者コード
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 更新者コード
	//[ 8] UPDA_USER_CD                                   char(8)              null: false  primary: false  isArray: false  auto: false  col: char            len: 8       default: []
	UPDAUSERCD string `gorm:"column:UPDA_USER_CD;type:char;size:8;"` // 更新者コード

}

var RAIL_AREA_MSTTableInfo = &TableInfo{
	Name: "RAIL_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "RAIL_AREA_CD",
			Comment:            `鉄道エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "RAILAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "rail_area_cd",
			ProtobufFieldName:  "rail_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "RAIL_AREA_NM",
			Comment:            `鉄道エリア名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "RAILAREANM",
			GoFieldType:        "string",
			JSONFieldName:      "rail_area_nm",
			ProtobufFieldName:  "rail_area_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ABRO_KBN_CD",
			Comment:            `海外区分コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ABROKBNCD",
			GoFieldType:        "string",
			JSONFieldName:      "abro_kbn_cd",
			ProtobufFieldName:  "abro_kbn_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "int32",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "UPDA_DTE",
			Comment:            `更新者コード`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RAIL_AREA_MST_) TableName() string {
	return "RAIL_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RAIL_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RAIL_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RAIL_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RAIL_AREA_MST_) TableInfo() *TableInfo {
	return RAIL_AREA_MSTTableInfo
}
