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


CREATE TABLE `CONT_MAIL_SEND_MAIL_ADDR` (
  `MAIL_SEND_ID` int(11) NOT NULL COMMENT 'メール送信ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `MAIL_SEQ` int(11) NOT NULL COMMENT 'メール連番',
  `SUB_NO` int(11) NOT NULL DEFAULT '1' COMMENT '注意報の市区町村ごとの採番(コンテンツ情報注警報の枝番と同一)',
  `MAIL_ADDR` varchar(256) NOT NULL COMMENT 'メールアドレス',
  `MAIL_SEND_FLG` char(1) DEFAULT NULL COMMENT 'メール配信フラグ',
  `MAIL_SEND_DTE` datetime DEFAULT NULL COMMENT '配信日時',
  `MAIL_SEND_TXT` text COMMENT '配信結果メッセージ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`MAIL_SEND_ID`,`USER_ID`,`MAIL_SEQ`,`SUB_NO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツメール配信メールアドレス'

JSON Sample
-------------------------------------
{    "mail_send_id": 7,    "user_id": 8,    "mail_seq": 9,    "sub_no": 15,    "mail_addr": "nelDVwwfYuxFCgfsmFPZRXNhX",    "mail_send_flg": "QWbWMJKlPXZoaZrtTrnDLnFNC",    "mail_send_dte": "2041-09-27T01:12:00.115778404+09:00",    "mail_send_txt": "PqTsnwmdJcCucljwgxfixdcPX",    "upda_dte": "2040-07-27T12:41:47.088399312+09:00",    "upda_user_id": 93,    "crea_dte": "2229-11-02T23:57:43.359627472+09:00",    "crea_user_id": 12}



*/

// CONT_MAIL_SEND_MAIL_ADDR_ struct is a row record of the CONT_MAIL_SEND_MAIL_ADDR table in the anpidb database
type CONT_MAIL_SEND_MAIL_ADDR_ struct {
	//[ 0] MAIL_SEND_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	MAILSENDID int32 `gorm:"primary_key;column:MAIL_SEND_ID;type:int;"` // メール送信ID
	//[ 1] USER_ID                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"primary_key;column:USER_ID;type:int;"` // ユーザID
	//[ 2] MAIL_SEQ                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	MAILSEQ int32 `gorm:"primary_key;column:MAIL_SEQ;type:int;"` // メール連番
	//[ 3] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [1]
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;default:1;"` // 注意報の市区町村ごとの採番(コンテンツ情報注警報の枝番と同一)
	//[ 4] MAIL_ADDR                                      varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILADDR string `gorm:"column:MAIL_ADDR;type:varchar;size:256;"` // メールアドレス
	//[ 5] MAIL_SEND_FLG                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSENDFLG sql.NullString `gorm:"column:MAIL_SEND_FLG;type:char;size:1;"` // メール配信フラグ
	//[ 6] MAIL_SEND_DTE                                  datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	MAILSENDDTE time.Time `gorm:"column:MAIL_SEND_DTE;type:datetime;"` // 配信日時
	//[ 7] MAIL_SEND_TXT                                  text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	MAILSENDTXT sql.NullString `gorm:"column:MAIL_SEND_TXT;type:text;size:65535;"` // 配信結果メッセージ
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var CONT_MAIL_SEND_MAIL_ADDRTableInfo = &TableInfo{
	Name: "CONT_MAIL_SEND_MAIL_ADDR",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MAIL_SEND_ID",
			Comment:            `メール送信ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
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
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "MAIL_SEQ",
			Comment:            `メール連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MAILSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_seq",
			ProtobufFieldName:  "mail_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "SUB_NO",
			Comment:            `注意報の市区町村ごとの採番(コンテンツ情報注警報の枝番と同一)`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "sub_no",
			ProtobufFieldName:  "sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "MAIL_SEND_FLG",
			Comment:            `メール配信フラグ`,
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
			Index:              6,
			Name:               "MAIL_SEND_DTE",
			Comment:            `配信日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "MAILSENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "mail_send_dte",
			ProtobufFieldName:  "mail_send_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "MAIL_SEND_TXT",
			Comment:            `配信結果メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "MAILSENDTXT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_send_txt",
			ProtobufFieldName:  "mail_send_txt",
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
func (c *CONT_MAIL_SEND_MAIL_ADDR_) TableName() string {
	return "CONT_MAIL_SEND_MAIL_ADDR"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_MAIL_SEND_MAIL_ADDR_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_MAIL_SEND_MAIL_ADDR_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_MAIL_SEND_MAIL_ADDR_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_MAIL_SEND_MAIL_ADDR_) TableInfo() *TableInfo {
	return CONT_MAIL_SEND_MAIL_ADDRTableInfo
}
