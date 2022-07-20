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


CREATE TABLE `USER_DLCT_DAY_MST` (
  `USER_DLCT_ID` int(11) NOT NULL COMMENT 'ユーザDLCT_ID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `DELI_DAY_KND` char(1) NOT NULL COMMENT '配信日種別',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_DLCT_ID`,`SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザDLCT_曜日マスタ	 天気予報の配信日'

JSON Sample
-------------------------------------
{    "user_dlct_id": 91,    "seq": 8,    "deli_day_knd": "DeKPCyrdOhyDfxmDtbtIFccwp",    "crea_dte": "2155-11-25T19:46:15.626670208+09:00",    "crea_user_id": 65}



*/

// USER_DLCT_DAY_MST_ struct is a row record of the USER_DLCT_DAY_MST table in the anpidb database
type USER_DLCT_DAY_MST_ struct {
	//[ 0] USER_DLCT_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERDLCTID int32 `gorm:"primary_key;column:USER_DLCT_ID;type:int;"` // ユーザDLCT_ID
	//[ 1] SEQ                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"primary_key;column:SEQ;type:int;"` // 連番
	//[ 2] DELI_DAY_KND                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELIDAYKND string `gorm:"column:DELI_DAY_KND;type:char;size:1;"` // 配信日種別
	//[ 3] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 4] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_DLCT_DAY_MSTTableInfo = &TableInfo{
	Name: "USER_DLCT_DAY_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_DLCT_ID",
			Comment:            `ユーザDLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_dlct_id",
			ProtobufFieldName:  "user_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "SEQ",
			Comment:            `連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DELI_DAY_KND",
			Comment:            `配信日種別`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELIDAYKND",
			GoFieldType:        "string",
			JSONFieldName:      "deli_day_knd",
			ProtobufFieldName:  "deli_day_knd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *USER_DLCT_DAY_MST_) TableName() string {
	return "USER_DLCT_DAY_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_DLCT_DAY_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_DLCT_DAY_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_DLCT_DAY_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_DLCT_DAY_MST_) TableInfo() *TableInfo {
	return USER_DLCT_DAY_MSTTableInfo
}
