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


CREATE TABLE `SM_SENDER_LOG` (
  `SM_SENDER_LOG_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '配信ログ管理ID',
  `DATE_TIME` timestamp NULL DEFAULT NULL COMMENT 'ログ出力日時',
  `STATUS` varchar(255) DEFAULT NULL COMMENT '配信状態',
  `RELAY_FROM_MAIL_ADDR` varchar(256) DEFAULT NULL COMMENT '送信元メールアドレス',
  `RELAY_TO_MAIL_ADDR` varchar(256) DEFAULT NULL COMMENT '送信先メールアドレス',
  `SEND_ID` varchar(255) DEFAULT NULL COMMENT '配信ID',
  PRIMARY KEY (`SM_SENDER_LOG_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "sm_sender_log_id": 41,    "date_time": "2026-08-08T15:22:07.492873651+09:00",    "status": "uWNYYYJbYPoDBuHufBGJOmSOg",    "relay_from_mail_addr": "poFJlrucbFLyfntvOtnOXsTee",    "relay_to_mail_addr": "mYRQCKURKZtspmIaNxKNLwQCC",    "send_id": "vtvsxVICvWKtBZxQWAsLExPkm"}



*/

// SM_SENDER_LOG_ struct is a row record of the SM_SENDER_LOG table in the anpidb database
type SM_SENDER_LOG_ struct {
	//[ 0] SM_SENDER_LOG_ID                               int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	SMSENDERLOGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:SM_SENDER_LOG_ID;type:int;"` // 配信ログ管理ID
	//[ 1] DATE_TIME                                      timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	DATETIME time.Time `gorm:"column:DATE_TIME;type:timestamp;"` // ログ出力日時
	//[ 2] STATUS                                         varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	STATUS sql.NullString `gorm:"column:STATUS;type:varchar;size:255;"` // 配信状態
	//[ 3] RELAY_FROM_MAIL_ADDR                           varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	RELAYFROMMAILADDR sql.NullString `gorm:"column:RELAY_FROM_MAIL_ADDR;type:varchar;size:256;"` // 送信元メールアドレス
	//[ 4] RELAY_TO_MAIL_ADDR                             varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	RELAYTOMAILADDR sql.NullString `gorm:"column:RELAY_TO_MAIL_ADDR;type:varchar;size:256;"` // 送信先メールアドレス
	//[ 5] SEND_ID                                        varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	SENDID sql.NullString `gorm:"column:SEND_ID;type:varchar;size:255;"` // 配信ID

}

var SM_SENDER_LOGTableInfo = &TableInfo{
	Name: "SM_SENDER_LOG",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "SM_SENDER_LOG_ID",
			Comment:            `配信ログ管理ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SMSENDERLOGID",
			GoFieldType:        "int32",
			JSONFieldName:      "sm_sender_log_id",
			ProtobufFieldName:  "sm_sender_log_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "DATE_TIME",
			Comment:            `ログ出力日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "DATETIME",
			GoFieldType:        "time.Time",
			JSONFieldName:      "date_time",
			ProtobufFieldName:  "date_time",
			ProtobufType:       "uint64",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "STATUS",
			Comment:            `配信状態`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "STATUS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "status",
			ProtobufFieldName:  "status",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "RELAY_FROM_MAIL_ADDR",
			Comment:            `送信元メールアドレス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "RELAYFROMMAILADDR",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "relay_from_mail_addr",
			ProtobufFieldName:  "relay_from_mail_addr",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "RELAY_TO_MAIL_ADDR",
			Comment:            `送信先メールアドレス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "RELAYTOMAILADDR",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "relay_to_mail_addr",
			ProtobufFieldName:  "relay_to_mail_addr",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "SEND_ID",
			Comment:            `配信ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "SENDID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "send_id",
			ProtobufFieldName:  "send_id",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SM_SENDER_LOG_) TableName() string {
	return "SM_SENDER_LOG"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SM_SENDER_LOG_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SM_SENDER_LOG_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SM_SENDER_LOG_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SM_SENDER_LOG_) TableInfo() *TableInfo {
	return SM_SENDER_LOGTableInfo
}
