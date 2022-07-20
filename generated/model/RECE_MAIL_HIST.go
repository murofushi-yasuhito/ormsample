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


CREATE TABLE `RECE_MAIL_HIST` (
  `RECE_MAIL_HIST_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'メール受信履歴ID',
  `MAIL_FROM` varchar(256) DEFAULT NULL COMMENT '差出人',
  `MAIL_TO` varchar(256) DEFAULT NULL COMMENT '送信先',
  `SUBJECT` varchar(256) DEFAULT NULL COMMENT '件名',
  `DATE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '送信日時',
  `REPLAY_TO` varchar(256) DEFAULT NULL COMMENT '返信先',
  `BODY` text COMMENT '本文',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`RECE_MAIL_HIST_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='メール受信履歴'

JSON Sample
-------------------------------------
{    "rece_mail_hist_id": 46,    "mail_from": "eFCVIfDXRSInvVWcDEluuIMix",    "mail_to": "lXkELLNpMHxOZLdRYWrIqfDTl",    "subject": "BkjpafXEqYmlBnFNdFYKrTiFS",    "date": "2031-08-05T15:14:09.386359924+09:00",    "replay_to": "erhIPCFHSQjQYmfnQMxeioVOo",    "body": "CxhCLDPRdDxwHYbQBsbhOpieH",    "upda_dte": "2220-05-27T12:30:07.148853803+09:00",    "upda_user_id": 3,    "crea_dte": "2252-09-06T17:22:30.180849154+09:00",    "crea_user_id": 99}



*/

// RECE_MAIL_HIST_ struct is a row record of the RECE_MAIL_HIST table in the anpidb database
type RECE_MAIL_HIST_ struct {
	//[ 0] RECE_MAIL_HIST_ID                              int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	RECEMAILHISTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:RECE_MAIL_HIST_ID;type:int;"` // メール受信履歴ID
	//[ 1] MAIL_FROM                                      varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILFROM sql.NullString `gorm:"column:MAIL_FROM;type:varchar;size:256;"` // 差出人
	//[ 2] MAIL_TO                                        varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILTO sql.NullString `gorm:"column:MAIL_TO;type:varchar;size:256;"` // 送信先
	//[ 3] SUBJECT                                        varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	SUBJECT sql.NullString `gorm:"column:SUBJECT;type:varchar;size:256;"` // 件名
	//[ 4] DATE                                           timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	DATE time.Time `gorm:"column:DATE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 送信日時
	//[ 5] REPLAY_TO                                      varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	REPLAYTO sql.NullString `gorm:"column:REPLAY_TO;type:varchar;size:256;"` // 返信先
	//[ 6] BODY                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	BODY sql.NullString `gorm:"column:BODY;type:text;size:65535;"` // 本文
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[10] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var RECE_MAIL_HISTTableInfo = &TableInfo{
	Name: "RECE_MAIL_HIST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "RECE_MAIL_HIST_ID",
			Comment:            `メール受信履歴ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RECEMAILHISTID",
			GoFieldType:        "int32",
			JSONFieldName:      "rece_mail_hist_id",
			ProtobufFieldName:  "rece_mail_hist_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "MAIL_FROM",
			Comment:            `差出人`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "MAILFROM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_from",
			ProtobufFieldName:  "mail_from",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "MAIL_TO",
			Comment:            `送信先`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "MAILTO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_to",
			ProtobufFieldName:  "mail_to",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "SUBJECT",
			Comment:            `件名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "SUBJECT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "subject",
			ProtobufFieldName:  "subject",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "DATE",
			Comment:            `送信日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "DATE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date",
			ProtobufFieldName:  "date",
			ProtobufType:       "uint64",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "REPLAY_TO",
			Comment:            `返信先`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "REPLAYTO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "replay_to",
			ProtobufFieldName:  "replay_to",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "BODY",
			Comment:            `本文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "BODY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "body",
			ProtobufFieldName:  "body",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RECE_MAIL_HIST_) TableName() string {
	return "RECE_MAIL_HIST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RECE_MAIL_HIST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RECE_MAIL_HIST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RECE_MAIL_HIST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RECE_MAIL_HIST_) TableInfo() *TableInfo {
	return RECE_MAIL_HISTTableInfo
}
