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


CREATE TABLE `DELI_SET_MENU_CATE_MST` (
  `DELI_SET_MENU_CATE_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '配信設定メニューカテゴリID',
  `DELI_SET_MENU_DTL_ID` int(11) NOT NULL DEFAULT '0' COMMENT '配信設定メニュー詳細ID',
  `CATE1_CD` varchar(20) NOT NULL DEFAULT '' COMMENT 'カテゴリ1コード',
  `CATE2_CD` varchar(20) NOT NULL DEFAULT '' COMMENT 'カテゴリ2コード',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`DELI_SET_MENU_CATE_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=43 DEFAULT CHARSET=utf8 COMMENT='配信設定メニューカテゴリマスタ'

JSON Sample
-------------------------------------
{    "deli_set_menu_cate_id": 30,    "deli_set_menu_dtl_id": 86,    "cate_1_cd": "UBBwNtHtVjHXxqdoEhUkjZVqK",    "cate_2_cd": "qvInIMIweLktxMcTOJHmuaDYW",    "upda_dte": "2133-03-05T12:59:27.884358693+09:00",    "upda_user_id": 68,    "crea_dte": "2076-11-13T05:01:57.075331332+09:00",    "crea_user_id": 35}



*/

// DELI_SET_MENU_CATE_MST_ struct is a row record of the DELI_SET_MENU_CATE_MST table in the anpidb database
type DELI_SET_MENU_CATE_MST_ struct {
	//[ 0] DELI_SET_MENU_CATE_ID                          int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	DELISETMENUCATEID int32 `gorm:"primary_key;AUTO_INCREMENT;column:DELI_SET_MENU_CATE_ID;type:int;"` // 配信設定メニューカテゴリID
	//[ 1] DELI_SET_MENU_DTL_ID                           int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DELISETMENUDTLID int32 `gorm:"column:DELI_SET_MENU_DTL_ID;type:int;default:0;"` // 配信設定メニュー詳細ID
	//[ 2] CATE1_CD                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:20;"` // カテゴリ1コード
	//[ 3] CATE2_CD                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	CATE2CD string `gorm:"column:CATE2_CD;type:varchar;size:20;"` // カテゴリ2コード
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var DELI_SET_MENU_CATE_MSTTableInfo = &TableInfo{
	Name: "DELI_SET_MENU_CATE_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "DELI_SET_MENU_CATE_ID",
			Comment:            `配信設定メニューカテゴリID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DELISETMENUCATEID",
			GoFieldType:        "int32",
			JSONFieldName:      "deli_set_menu_cate_id",
			ProtobufFieldName:  "deli_set_menu_cate_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "DELI_SET_MENU_DTL_ID",
			Comment:            `配信設定メニュー詳細ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DELISETMENUDTLID",
			GoFieldType:        "int32",
			JSONFieldName:      "deli_set_menu_dtl_id",
			ProtobufFieldName:  "deli_set_menu_dtl_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "CATE2_CD",
			Comment:            `カテゴリ2コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DELI_SET_MENU_CATE_MST_) TableName() string {
	return "DELI_SET_MENU_CATE_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DELI_SET_MENU_CATE_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DELI_SET_MENU_CATE_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DELI_SET_MENU_CATE_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DELI_SET_MENU_CATE_MST_) TableInfo() *TableInfo {
	return DELI_SET_MENU_CATE_MSTTableInfo
}
