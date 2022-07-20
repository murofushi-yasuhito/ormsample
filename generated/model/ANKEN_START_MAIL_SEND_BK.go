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


CREATE TABLE `ANKEN_START_MAIL_SEND_BK` (
  `ANKEN_START_MAIL_SEND_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '案件起動通知送信ID',
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID	 通知専用は0',
  `MAIL_ONLY_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '通知専用区分	 1:通知専用 0:ユーザ',
  `MAIL_ADDR` varchar(256) NOT NULL COMMENT 'メールアドレス',
  `ACCS_KEY` varchar(20) NOT NULL COMMENT 'アクセスキー',
  `MAIL_SEND_FLG` char(1) NOT NULL DEFAULT '0' COMMENT 'メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_START_MAIL_SEND_ID`),
  KEY `ANKEN_START_MAIL_SEND_IDX1` (`ANKEN_ID`),
  KEY `test32` (`MAIL_ADDR`(255))
) ENGINE=InnoDB AUTO_INCREMENT=711066 DEFAULT CHARSET=utf8 COMMENT='案件起動通知配信'

JSON Sample
-------------------------------------
{    "anken_start_mail_send_id": 0,    "anken_id": 92,    "client_id": 78,    "user_id": 92,    "mail_only_kbn": "ZoZmCQATvRVrfYEdHiGNSwROH",    "mail_addr": "rveHifvCVxTAGchQUJBHLugYr",    "accs_key": "vZlXWmhsgcNJYFleONKMsEidu",    "mail_send_flg": "wyaVeHcwAGeMRPjwkVFwqFUCY",    "upda_dte": "2067-02-03T02:06:22.267165988+09:00",    "upda_user_id": 90,    "crea_dte": "2183-02-21T11:27:11.85980476+09:00",    "crea_user_id": 58}



*/

// ANKEN_START_MAIL_SEND_BK struct is a row record of the ANKEN_START_MAIL_SEND_BK table in the anpidb database
type ANKEN_START_MAIL_SEND_BK struct {
	//[ 0] ANKEN_START_MAIL_SEND_ID                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENSTARTMAILSENDID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_START_MAIL_SEND_ID;type:int;"` // 案件起動通知送信ID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;"` // 案件ID
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID	 通知専用は0
	//[ 4] MAIL_ONLY_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILONLYKBN string `gorm:"column:MAIL_ONLY_KBN;type:char;size:1;default:0;"` // 通知専用区分	 1:通知専用 0:ユーザ
	//[ 5] MAIL_ADDR                                      varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILADDR string `gorm:"column:MAIL_ADDR;type:varchar;size:256;"` // メールアドレス
	//[ 6] ACCS_KEY                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	ACCSKEY string `gorm:"column:ACCS_KEY;type:varchar;size:20;"` // アクセスキー
	//[ 7] MAIL_SEND_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSENDFLG string `gorm:"column:MAIL_SEND_FLG;type:char;size:1;default:0;"` // メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_START_MAIL_SEND_BKTableInfo = &TableInfo{
	Name: "ANKEN_START_MAIL_SEND_BK",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_START_MAIL_SEND_ID",
			Comment:            `案件起動通知送信ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENSTARTMAILSENDID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_start_mail_send_id",
			ProtobufFieldName:  "anken_start_mail_send_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
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
			Index: 3,
			Name:  "USER_ID",
			Comment: `ユーザID	 通知専用は0`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "MAIL_ONLY_KBN",
			Comment: `通知専用区分	 1:通知専用 0:ユーザ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILONLYKBN",
			GoFieldType:        "string",
			JSONFieldName:      "mail_only_kbn",
			ProtobufFieldName:  "mail_only_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "MAIL_ADDR",
			Comment:            `メールアドレス`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "MAILADDR",
			GoFieldType:        "string",
			JSONFieldName:      "mail_addr",
			ProtobufFieldName:  "mail_addr",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ACCS_KEY",
			Comment:            `アクセスキー`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "ACCSKEY",
			GoFieldType:        "string",
			JSONFieldName:      "accs_key",
			ProtobufFieldName:  "accs_key",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index: 7,
			Name:  "MAIL_SEND_FLG",
			Comment: `メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSENDFLG",
			GoFieldType:        "string",
			JSONFieldName:      "mail_send_flg",
			ProtobufFieldName:  "mail_send_flg",
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
func (a *ANKEN_START_MAIL_SEND_BK) TableName() string {
	return "ANKEN_START_MAIL_SEND_BK"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_START_MAIL_SEND_BK) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_START_MAIL_SEND_BK) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_START_MAIL_SEND_BK) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_START_MAIL_SEND_BK) TableInfo() *TableInfo {
	return ANKEN_START_MAIL_SEND_BKTableInfo
}
