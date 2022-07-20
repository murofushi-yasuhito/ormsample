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


CREATE TABLE `CATEGORY2_MST` (
  `CATE2_CD` varchar(10) NOT NULL COMMENT 'カテゴリ2コード',
  `CATE2_NM` varchar(20) NOT NULL COMMENT 'カテゴリ2名称',
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード',
  `DISP_DSQ` int(11) DEFAULT NULL COMMENT '表示順',
  `LEVEL_ID` int(11) DEFAULT NULL COMMENT 'レベルID',
  `LEVEL_KBN` char(1) DEFAULT NULL COMMENT 'レベル区分	 C:コンテンツレベル A:エリアレベル',
  `CATE_NOTE` varchar(100) DEFAULT NULL COMMENT 'カテゴリ説明',
  `VALID_KBN` char(1) DEFAULT '1' COMMENT '有効区分',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`CATE2_CD`),
  UNIQUE KEY `CATEGORY2_MST_PKI` (`CATE2_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='カテゴリ2マスタ'

JSON Sample
-------------------------------------
{    "cate_2_cd": "ijYIVEHwLUKbNKcvHNTHmCPnL",    "cate_2_nm": "rrTVCwWpMveudDMbMGZWoGbYw",    "cate_1_cd": "lUQxcvmEFlLZJAOEoSTXEDqEO",    "disp_dsq": 29,    "level_id": 12,    "level_kbn": "jviPrPBBoxaSrnQoCyNRtZych",    "cate_note": "oJIZTeYUsctLbtXfWViiPAsOK",    "valid_kbn": "IfhiHsxdQcxakmqOHsnEKuyTd",    "upda_dte": "2147-04-02T21:36:44.441862096+09:00",    "upda_user_id": 54,    "crea_dte": "2309-09-06T11:58:41.984449372+09:00",    "crea_user_id": 21}



*/

// CATEGORY2_MST_ struct is a row record of the CATEGORY2_MST table in the anpidb database
type CATEGORY2_MST_ struct {
	//[ 0] CATE2_CD                                       varchar(10)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE2CD string `gorm:"primary_key;column:CATE2_CD;type:varchar;size:10;"` // カテゴリ2コード
	//[ 1] CATE2_NM                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	CATE2NM string `gorm:"column:CATE2_NM;type:varchar;size:20;"` // カテゴリ2名称
	//[ 2] CATE1_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード
	//[ 3] DISP_DSQ                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ sql.NullInt64 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 4] LEVEL_ID                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LEVELID sql.NullInt64 `gorm:"column:LEVEL_ID;type:int;"` // レベルID
	//[ 5] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;"` // レベル区分	 C:コンテンツレベル A:エリアレベル
	//[ 6] CATE_NOTE                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	CATENOTE sql.NullString `gorm:"column:CATE_NOTE;type:varchar;size:100;"` // カテゴリ説明
	//[ 7] VALID_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	VALIDKBN sql.NullString `gorm:"column:VALID_KBN;type:char;size:1;default:1;"` // 有効区分
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var CATEGORY2_MSTTableInfo = &TableInfo{
	Name: "CATEGORY2_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CATE2_CD",
			Comment:            `カテゴリ2コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CATE2_NM",
			Comment:            `カテゴリ2名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "CATE2NM",
			GoFieldType:        "string",
			JSONFieldName:      "cate_2_nm",
			ProtobufFieldName:  "cate_2_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "LEVEL_ID",
			Comment:            `レベルID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "level_id",
			ProtobufFieldName:  "level_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "LEVEL_KBN",
			Comment: `レベル区分	 C:コンテンツレベル A:エリアレベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LEVELKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "level_kbn",
			ProtobufFieldName:  "level_kbn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CATE_NOTE",
			Comment:            `カテゴリ説明`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "CATENOTE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_note",
			ProtobufFieldName:  "cate_note",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "VALID_KBN",
			Comment:            `有効区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "VALIDKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "valid_kbn",
			ProtobufFieldName:  "valid_kbn",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "UPDA_USER_ID",
			Comment:            `更新者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDAUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "upda_user_id",
			ProtobufFieldName:  "upda_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "CREA_USER_ID",
			Comment:            `作成者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CREAUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "crea_user_id",
			ProtobufFieldName:  "crea_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CATEGORY2_MST_) TableName() string {
	return "CATEGORY2_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CATEGORY2_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CATEGORY2_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CATEGORY2_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CATEGORY2_MST_) TableInfo() *TableInfo {
	return CATEGORY2_MSTTableInfo
}
