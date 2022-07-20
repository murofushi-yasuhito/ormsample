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


CREATE TABLE `ANKEN_TARGET_USER` (
  `ANKEN_TARGET_ID` int(11) NOT NULL COMMENT '送信先対象選択ID',
  `USER_SEQ` int(11) NOT NULL COMMENT 'ユーザ連番',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `USER_NO` varchar(20) NOT NULL COMMENT 'ユーザ番号',
  `USER_NM` varchar(60) NOT NULL DEFAULT '' COMMENT '役職名',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`,`USER_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 63,    "user_seq": 57,    "user_id": 62,    "user_no": "OldXoeSvqebqEjQMLyVHwMnbw",    "user_nm": "xDbNMxedXvTZVZTLOkWFqnFmP",    "upda_dte": "2096-05-14T11:45:48.907567113+09:00",    "upda_user_id": 72,    "crea_dte": "2309-08-03T19:47:50.053255843+09:00",    "crea_user_id": 91}



*/

// ANKEN_TARGET_USER_ struct is a row record of the ANKEN_TARGET_USER table in the anpidb database
type ANKEN_TARGET_USER_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] USER_SEQ                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERSEQ int32 `gorm:"primary_key;column:USER_SEQ;type:int;"` // ユーザ連番
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 3] USER_NO                                        varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	USERNO string `gorm:"column:USER_NO;type:varchar;size:20;"` // ユーザ番号
	//[ 4] USER_NM                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	USERNM string `gorm:"column:USER_NM;type:varchar;size:60;"` // 役職名
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_TARGET_USERTableInfo = &TableInfo{
	Name: "ANKEN_TARGET_USER",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_TARGET_ID",
			Comment:            `送信先対象選択ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENTARGETID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_target_id",
			ProtobufFieldName:  "anken_target_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "USER_SEQ",
			Comment:            `ユーザ連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "user_seq",
			ProtobufFieldName:  "user_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "USER_NO",
			Comment:            `ユーザ番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "USERNO",
			GoFieldType:        "string",
			JSONFieldName:      "user_no",
			ProtobufFieldName:  "user_no",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "USER_NM",
			Comment:            `役職名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "USERNM",
			GoFieldType:        "string",
			JSONFieldName:      "user_nm",
			ProtobufFieldName:  "user_nm",
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
func (a *ANKEN_TARGET_USER_) TableName() string {
	return "ANKEN_TARGET_USER"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_USER_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_USER_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_USER_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_USER_) TableInfo() *TableInfo {
	return ANKEN_TARGET_USERTableInfo
}
