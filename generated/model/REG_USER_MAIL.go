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


CREATE TABLE `REG_USER_MAIL` (
  `REG_USER_MAIL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザ初期登録メールID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'ユーザID',
  `USER_MAIL_ADDR` varchar(256) NOT NULL COMMENT 'ユーザメールアドレス',
  `ACCS_KEY` varchar(32) DEFAULT NULL,
  `ACCS_KEY_EXPIRY` timestamp NULL DEFAULT NULL COMMENT 'アクセスキー有効期限',
  `REG_MAIL_SEND_FLG` char(1) DEFAULT '0' COMMENT '登録確認配信フラグ	 0:未配信,1:配信済,2:メール送信不可,3:メール送信エラー',
  `REG_MAIL_SEND_MSG` text COMMENT '登録確認配信メッセージ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`REG_USER_MAIL_ID`),
  UNIQUE KEY `REG\USER_MAIL_IDX_1` (`CLIENT_ID`,`ACCS_KEY`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=305 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "reg_user_mail_id": 25,    "client_id": 95,    "user_id": 78,    "user_mail_addr": "WhXpLdHwquXwDZQaPCCKITDZV",    "accs_key": "AiVoAXBunIhlPjsiwRhsdSHyp",    "accs_key_expiry": "2257-09-01T16:15:33.613825821+09:00",    "reg_mail_send_flg": "nAZmtOrOwQKysGrufHeZghaeG",    "reg_mail_send_msg": "vjplmTCVBqellHPOqOVRhYOFn",    "upda_dte": "2262-03-03T18:44:03.915081765+09:00",    "upda_user_id": 80,    "crea_dte": "2185-05-08T03:56:26.451490028+09:00",    "crea_user_id": 71}



*/

// REG_USER_MAIL_ struct is a row record of the REG_USER_MAIL table in the anpidb database
type REG_USER_MAIL_ struct {
	//[ 0] REG_USER_MAIL_ID                               int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	REGUSERMAILID int32 `gorm:"primary_key;AUTO_INCREMENT;column:REG_USER_MAIL_ID;type:int;"` // ユーザ初期登録メールID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	USERID int32 `gorm:"column:USER_ID;type:int;default:0;"` // ユーザID
	//[ 3] USER_MAIL_ADDR                                 varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	USERMAILADDR string `gorm:"column:USER_MAIL_ADDR;type:varchar;size:256;"` // ユーザメールアドレス
	//[ 4] ACCS_KEY                                       varchar(32)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	ACCSKEY sql.NullString `gorm:"column:ACCS_KEY;type:varchar;size:32;"`
	//[ 5] ACCS_KEY_EXPIRY                                timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	ACCSKEYEXPIRY time.Time `gorm:"column:ACCS_KEY_EXPIRY;type:timestamp;"` // アクセスキー有効期限
	//[ 6] REG_MAIL_SEND_FLG                              char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	REGMAILSENDFLG sql.NullString `gorm:"column:REG_MAIL_SEND_FLG;type:char;size:1;default:0;"` // 登録確認配信フラグ	 0:未配信,1:配信済,2:メール送信不可,3:メール送信エラー
	//[ 7] REG_MAIL_SEND_MSG                              text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	REGMAILSENDMSG sql.NullString `gorm:"column:REG_MAIL_SEND_MSG;type:text;size:65535;"` // 登録確認配信メッセージ
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var REG_USER_MAILTableInfo = &TableInfo{
	Name: "REG_USER_MAIL",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "REG_USER_MAIL_ID",
			Comment:            `ユーザ初期登録メールID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "REGUSERMAILID",
			GoFieldType:        "int32",
			JSONFieldName:      "reg_user_mail_id",
			ProtobufFieldName:  "reg_user_mail_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			Name:               "USER_MAIL_ADDR",
			Comment:            `ユーザメールアドレス`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "USERMAILADDR",
			GoFieldType:        "string",
			JSONFieldName:      "user_mail_addr",
			ProtobufFieldName:  "user_mail_addr",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ACCS_KEY",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "ACCSKEY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "accs_key",
			ProtobufFieldName:  "accs_key",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "ACCS_KEY_EXPIRY",
			Comment:            `アクセスキー有効期限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "ACCSKEYEXPIRY",
			GoFieldType:        "time.Time",
			JSONFieldName:      "accs_key_expiry",
			ProtobufFieldName:  "accs_key_expiry",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index: 6,
			Name:  "REG_MAIL_SEND_FLG",
			Comment: `登録確認配信フラグ	 0:未配信,1:配信済,2:メール送信不可,3:メール送信エラー`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "REGMAILSENDFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "reg_mail_send_flg",
			ProtobufFieldName:  "reg_mail_send_flg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "REG_MAIL_SEND_MSG",
			Comment:            `登録確認配信メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "REGMAILSENDMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "reg_mail_send_msg",
			ProtobufFieldName:  "reg_mail_send_msg",
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
func (r *REG_USER_MAIL_) TableName() string {
	return "REG_USER_MAIL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *REG_USER_MAIL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *REG_USER_MAIL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *REG_USER_MAIL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *REG_USER_MAIL_) TableInfo() *TableInfo {
	return REG_USER_MAILTableInfo
}
