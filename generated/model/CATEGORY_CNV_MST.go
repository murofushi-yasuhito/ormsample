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


CREATE TABLE `CATEGORY_CNV_MST` (
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード',
  `CATE2_CD` varchar(10) NOT NULL COMMENT 'カテゴリ2コード',
  `CATE_CD` varchar(12) NOT NULL COMMENT 'カテゴリコード(RSQ3)',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`CATE1_CD`,`CATE2_CD`,`CATE_CD`),
  UNIQUE KEY `CATEGORY_CNV_MST_PKI` (`CATE1_CD`,`CATE2_CD`,`CATE_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='カテゴリ変換マスタ'

JSON Sample
-------------------------------------
{    "cate_1_cd": "ygLBhNTUKrIFWiCFfKgNtKCTQ",    "cate_2_cd": "BpbrAsTJfTeSiBSAqlfMtICBi",    "cate_cd": "FwjHTxCjyxRuSPCrHpCEFAnUg",    "upda_dte": "2238-03-06T19:19:25.772402136+09:00",    "upda_user_id": 98,    "crea_dte": "2148-03-05T12:09:33.091829319+09:00",    "crea_user_id": 58}



*/

// CATEGORY_CNV_MST_ struct is a row record of the CATEGORY_CNV_MST table in the anpidb database
type CATEGORY_CNV_MST_ struct {
	//[ 0] CATE1_CD                                       varchar(10)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"primary_key;column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード
	//[ 1] CATE2_CD                                       varchar(10)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE2CD string `gorm:"primary_key;column:CATE2_CD;type:varchar;size:10;"` // カテゴリ2コード
	//[ 2] CATE_CD                                        varchar(12)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD string `gorm:"primary_key;column:CATE_CD;type:varchar;size:12;"` // カテゴリコード(RSQ3)
	//[ 3] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 4] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 6] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var CATEGORY_CNV_MSTTableInfo = &TableInfo{
	Name: "CATEGORY_CNV_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CATE_CD",
			Comment:            `カテゴリコード(RSQ3)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_cd",
			ProtobufFieldName:  "cate_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CATEGORY_CNV_MST_) TableName() string {
	return "CATEGORY_CNV_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CATEGORY_CNV_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CATEGORY_CNV_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CATEGORY_CNV_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CATEGORY_CNV_MST_) TableInfo() *TableInfo {
	return CATEGORY_CNV_MSTTableInfo
}
