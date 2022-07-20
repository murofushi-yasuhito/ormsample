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


CREATE TABLE `ANKEN_MSG` (
  `MSG_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'メッセージID',
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `MSG_SEND_KND` char(1) NOT NULL COMMENT 'メッセージ送信種別	 1:案件 2:再送 3:指示',
  `MSG_TTL` varchar(200) DEFAULT NULL,
  `MSG_BODY` varchar(12000) DEFAULT NULL,
  `SEND_USER_ID` int(11) DEFAULT NULL COMMENT '送信者ID',
  `MAIL_SEND_FLG` char(1) DEFAULT '0' COMMENT 'メール配信フラグ     0:未配信 1:配信済 2:配信不可',
  `AUTO_RESEND_NUM` int(11) DEFAULT NULL COMMENT '自動再送回数',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MSG_ID`),
  KEY `ANKEN_MSG_IDX1` (`ANKEN_ID`),
  KEY `ANKEN_MSG_IDX2` (`MAIL_SEND_FLG`)
) ENGINE=InnoDB AUTO_INCREMENT=460251 DEFAULT CHARSET=utf8 COMMENT='案件メッセージ'

JSON Sample
-------------------------------------
{    "msg_id": 12,    "anken_id": 66,    "msg_send_knd": "DJUEFtDHYfrHbwUXgnlHoLmbq",    "msg_ttl": "LVESZiVsByiDWPnYAwjwLMndC",    "msg_body": "hobAdXuhtPGjElBOWZtoErtfe",    "send_user_id": 9,    "mail_send_flg": "gooSAXpmiqdsHIuYtDtnEPXEj",    "auto_resend_num": 34,    "upda_dte": "2246-01-21T15:35:06.681305802+09:00",    "upda_user_id": 52,    "crea_dte": "2211-01-14T05:20:30.157093668+09:00",    "crea_user_id": 54}



*/

// ANKEN_MSG_ struct is a row record of the ANKEN_MSG table in the anpidb database
type ANKEN_MSG_ struct {
	//[ 0] MSG_ID                                         int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MSGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MSG_ID;type:int;"` // メッセージID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;"` // 案件ID
	//[ 2] MSG_SEND_KND                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MSGSENDKND string `gorm:"column:MSG_SEND_KND;type:char;size:1;"` // メッセージ送信種別	 1:案件 2:再送 3:指示
	//[ 3] MSG_TTL                                        varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	MSGTTL sql.NullString `gorm:"column:MSG_TTL;type:varchar;size:200;"`
	//[ 4] MSG_BODY                                       varchar(12000)       null: true   primary: false  isArray: false  auto: false  col: varchar         len: 12000   default: []
	MSGBODY sql.NullString `gorm:"column:MSG_BODY;type:varchar;size:12000;"`
	//[ 5] SEND_USER_ID                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SENDUSERID sql.NullInt64 `gorm:"column:SEND_USER_ID;type:int;"` // 送信者ID
	//[ 6] MAIL_SEND_FLG                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSENDFLG sql.NullString `gorm:"column:MAIL_SEND_FLG;type:char;size:1;default:0;"` // メール配信フラグ     0:未配信 1:配信済 2:配信不可
	//[ 7] AUTO_RESEND_NUM                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AUTORESENDNUM sql.NullInt64 `gorm:"column:AUTO_RESEND_NUM;type:int;"` // 自動再送回数
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_MSGTableInfo = &TableInfo{
	Name: "ANKEN_MSG",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MSG_ID",
			Comment:            `メッセージID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MSGID",
			GoFieldType:        "int32",
			JSONFieldName:      "msg_id",
			ProtobufFieldName:  "msg_id",
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
			Index: 2,
			Name:  "MSG_SEND_KND",
			Comment: `メッセージ送信種別	 1:案件 2:再送 3:指示`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MSGSENDKND",
			GoFieldType:        "string",
			JSONFieldName:      "msg_send_knd",
			ProtobufFieldName:  "msg_send_knd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "MSG_TTL",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "MSGTTL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "msg_ttl",
			ProtobufFieldName:  "msg_ttl",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "MSG_BODY",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12000,
			GoFieldName:        "MSGBODY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "msg_body",
			ProtobufFieldName:  "msg_body",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "SEND_USER_ID",
			Comment:            `送信者ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SENDUSERID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "send_user_id",
			ProtobufFieldName:  "send_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "MAIL_SEND_FLG",
			Comment:            `メール配信フラグ     0:未配信 1:配信済 2:配信不可`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "AUTO_RESEND_NUM",
			Comment:            `自動再送回数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AUTORESENDNUM",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "auto_resend_num",
			ProtobufFieldName:  "auto_resend_num",
			ProtobufType:       "int32",
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
func (a *ANKEN_MSG_) TableName() string {
	return "ANKEN_MSG"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_MSG_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_MSG_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_MSG_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_MSG_) TableInfo() *TableInfo {
	return ANKEN_MSGTableInfo
}
