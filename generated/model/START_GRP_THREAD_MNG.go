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


CREATE TABLE `START_GRP_THREAD_MNG` (
  `START_GRP_THREAD_MNG_ID` int(11) NOT NULL AUTO_INCREMENT,
  `PROC_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '0:実行待ち 1:実行中 2:完了 9:エラー',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `START_GRP_ID` int(11) NOT NULL COMMENT '起動グループID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID（0の場合はグループ対象ユーザ全て）',
  `LOGIN_USER_ID` int(11) NOT NULL COMMENT '更新ユーザID',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`START_GRP_THREAD_MNG_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=24995 DEFAULT CHARSET=utf8 ROW_FORMAT=COMPACT

JSON Sample
-------------------------------------
{    "start_grp_thread_mng_id": 37,    "proc_flg": "LOkbUoVElRttKPuugcryxPxgo",    "client_id": 5,    "start_grp_id": 91,    "user_id": 33,    "login_user_id": 50,    "upda_dte": "2292-11-20T09:15:12.275758268+09:00",    "upda_user_id": 25,    "crea_dte": "2071-01-29T23:52:21.239431011+09:00",    "crea_user_id": 82}



*/

// START_GRP_THREAD_MNG_ struct is a row record of the START_GRP_THREAD_MNG table in the anpidb database
type START_GRP_THREAD_MNG_ struct {
	//[ 0] START_GRP_THREAD_MNG_ID                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	STARTGRPTHREADMNGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:START_GRP_THREAD_MNG_ID;type:int;"`
	//[ 1] PROC_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	PROCFLG string `gorm:"column:PROC_FLG;type:char;size:1;default:0;"` // 0:実行待ち 1:実行中 2:完了 9:エラー
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] START_GRP_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	STARTGRPID int32 `gorm:"column:START_GRP_ID;type:int;"` // 起動グループID
	//[ 4] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID（0の場合はグループ対象ユーザ全て）
	//[ 5] LOGIN_USER_ID                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LOGINUSERID int32 `gorm:"column:LOGIN_USER_ID;type:int;"` // 更新ユーザID
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var START_GRP_THREAD_MNGTableInfo = &TableInfo{
	Name: "START_GRP_THREAD_MNG",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "START_GRP_THREAD_MNG_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPTHREADMNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_thread_mng_id",
			ProtobufFieldName:  "start_grp_thread_mng_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "PROC_FLG",
			Comment:            `0:実行待ち 1:実行中 2:完了 9:エラー`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PROCFLG",
			GoFieldType:        "string",
			JSONFieldName:      "proc_flg",
			ProtobufFieldName:  "proc_flg",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CLIENT_ID",
			Comment:            `クライアントID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "START_GRP_ID",
			Comment:            `起動グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_id",
			ProtobufFieldName:  "start_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "USER_ID",
			Comment:            `ユーザID（0の場合はグループ対象ユーザ全て）`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "LOGIN_USER_ID",
			Comment:            `更新ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LOGINUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "login_user_id",
			ProtobufFieldName:  "login_user_id",
			ProtobufType:       "int32",
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *START_GRP_THREAD_MNG_) TableName() string {
	return "START_GRP_THREAD_MNG"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *START_GRP_THREAD_MNG_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *START_GRP_THREAD_MNG_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *START_GRP_THREAD_MNG_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *START_GRP_THREAD_MNG_) TableInfo() *TableInfo {
	return START_GRP_THREAD_MNGTableInfo
}
