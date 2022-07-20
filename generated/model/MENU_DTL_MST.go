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


CREATE TABLE `MENU_DTL_MST` (
  `MENU_DTL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'メニュー詳細ID',
  `MENU_ID` int(11) NOT NULL COMMENT 'メニューID',
  `CATE1_CD` varchar(12) NOT NULL COMMENT 'カテゴリ1コード',
  `CATE2_CD` varchar(12) DEFAULT NULL COMMENT 'カテゴリ2コード',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MENU_DTL_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "menu_dtl_id": 6,    "menu_id": 74,    "cate_1_cd": "XIJiRVRXEaZYmBAyrSrDtjpoL",    "cate_2_cd": "ICPIVwLlxFdIGFZfQAFuYqvFV",    "disp_dsq": 4,    "upda_dte": "2225-03-18T00:27:03.086915661+09:00",    "upda_user_id": 62,    "crea_dte": "2184-04-30T05:36:21.566982018+09:00",    "crea_user_id": 85}



*/

// MENU_DTL_MST_ struct is a row record of the MENU_DTL_MST table in the anpidb database
type MENU_DTL_MST_ struct {
	//[ 0] MENU_DTL_ID                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MENUDTLID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MENU_DTL_ID;type:int;"` // メニュー詳細ID
	//[ 1] MENU_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MENUID int32 `gorm:"column:MENU_ID;type:int;"` // メニューID
	//[ 2] CATE1_CD                                       varchar(12)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:12;"` // カテゴリ1コード
	//[ 3] CATE2_CD                                       varchar(12)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATE2CD sql.NullString `gorm:"column:CATE2_CD;type:varchar;size:12;"` // カテゴリ2コード
	//[ 4] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var MENU_DTL_MSTTableInfo = &TableInfo{
	Name: "MENU_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MENU_DTL_ID",
			Comment:            `メニュー詳細ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MENUDTLID",
			GoFieldType:        "int32",
			JSONFieldName:      "menu_dtl_id",
			ProtobufFieldName:  "menu_dtl_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "MENU_ID",
			Comment:            `メニューID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MENUID",
			GoFieldType:        "int32",
			JSONFieldName:      "menu_id",
			ProtobufFieldName:  "menu_id",
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
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
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
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MENU_DTL_MST_) TableName() string {
	return "MENU_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MENU_DTL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MENU_DTL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MENU_DTL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MENU_DTL_MST_) TableInfo() *TableInfo {
	return MENU_DTL_MSTTableInfo
}
