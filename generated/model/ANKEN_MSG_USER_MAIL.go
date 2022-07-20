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


CREATE TABLE `ANKEN_MSG_USER_MAIL` (
  `MAIL_SEND_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '送信メッセージユーザID',
  `MSG_ID` int(11) NOT NULL COMMENT 'メッセージID',
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `MAIL_ADDR` varchar(256) DEFAULT NULL COMMENT 'メールアドレス',
  `MAIL_SEND_FLG` char(1) DEFAULT NULL COMMENT 'メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー',
  `MAIL_STS` char(1) DEFAULT '0' COMMENT 'メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG',
  `MAIL_STS_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT 'メールステータス更新日時',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MAIL_SEND_ID`),
  KEY `ANKEN_MSG_USER_MAIL_IDX_1` (`MSG_ID`,`ANKEN_ID`,`USER_ID`) USING BTREE,
  KEY `ANKEN_MSG_USER_MAIL_IDX_2` (`MAIL_ADDR`(255)) USING BTREE,
  KEY `ANKEN_MSG_USER_MAIL_IDX_4` (`ANKEN_ID`) USING BTREE,
  KEY `ANKEN_MSG_USER_MAIL_IDX_3` (`MAIL_STS_DTE`) USING BTREE,
  KEY `ANKEN_MSG_USER_MAIL_IDX_5` (`ANKEN_ID`,`USER_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=49985646 DEFAULT CHARSET=utf8 COMMENT='案件メッセージユーザメール'

JSON Sample
-------------------------------------
{    "mail_send_id": 49,    "msg_id": 48,    "anken_id": 15,    "user_id": 6,    "mail_addr": "WZVreUAlpdlOasyREctqqVSfA",    "mail_send_flg": "dfsYBDbSCNoNDPLtSQRZEdsUB",    "mail_sts": "BBDYGFIKnixjKbetMQZWIxuNJ",    "mail_sts_dte": "2210-03-06T06:10:35.735275305+09:00",    "upda_dte": "2252-06-29T19:42:10.341108195+09:00",    "upda_user_id": 1,    "crea_dte": "2216-04-04T01:25:56.387108292+09:00",    "crea_user_id": 14}



*/

// ANKEN_MSG_USER_MAIL_ struct is a row record of the ANKEN_MSG_USER_MAIL table in the anpidb database
type ANKEN_MSG_USER_MAIL_ struct {
	//[ 0] MAIL_SEND_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MAILSENDID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MAIL_SEND_ID;type:int;"` // 送信メッセージユーザID
	//[ 1] MSG_ID                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MSGID int32 `gorm:"column:MSG_ID;type:int;"` // メッセージID
	//[ 2] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;"` // 案件ID
	//[ 3] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 4] MAIL_ADDR                                      varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILADDR sql.NullString `gorm:"column:MAIL_ADDR;type:varchar;size:256;"` // メールアドレス
	//[ 5] MAIL_SEND_FLG                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSENDFLG sql.NullString `gorm:"column:MAIL_SEND_FLG;type:char;size:1;"` // メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー
	//[ 6] MAIL_STS                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSTS sql.NullString `gorm:"column:MAIL_STS;type:char;size:1;default:0;"` // メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG
	//[ 7] MAIL_STS_DTE                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	MAILSTSDTE time.Time `gorm:"column:MAIL_STS_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // メールステータス更新日時
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_MSG_USER_MAILTableInfo = &TableInfo{
	Name: "ANKEN_MSG_USER_MAIL",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MAIL_SEND_ID",
			Comment:            `送信メッセージユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MAILSENDID",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_send_id",
			ProtobufFieldName:  "mail_send_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "MSG_ID",
			Comment:            `メッセージID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MSGID",
			GoFieldType:        "int32",
			JSONFieldName:      "msg_id",
			ProtobufFieldName:  "msg_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "MAIL_ADDR",
			Comment:            `メールアドレス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "MAILADDR",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_addr",
			ProtobufFieldName:  "mail_addr",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "MAIL_SEND_FLG",
			Comment: `メール配信フラグ	 0:未配信 1:配信済 2:配信不可 3:配信エラー`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSENDFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_send_flg",
			ProtobufFieldName:  "mail_send_flg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index: 6,
			Name:  "MAIL_STS",
			Comment: `メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSTS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_sts",
			ProtobufFieldName:  "mail_sts",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "MAIL_STS_DTE",
			Comment:            `メールステータス更新日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "MAILSTSDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "mail_sts_dte",
			ProtobufFieldName:  "mail_sts_dte",
			ProtobufType:       "uint64",
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
func (a *ANKEN_MSG_USER_MAIL_) TableName() string {
	return "ANKEN_MSG_USER_MAIL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_MSG_USER_MAIL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_MSG_USER_MAIL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_MSG_USER_MAIL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_MSG_USER_MAIL_) TableInfo() *TableInfo {
	return ANKEN_MSG_USER_MAILTableInfo
}
