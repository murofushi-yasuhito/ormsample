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


CREATE TABLE `USER_MAIL_MST` (
  `USER_MAIL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザメールID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `MAIL_ADDR_SEQ` int(11) NOT NULL COMMENT 'メールアドレス連番',
  `MAIL_ADDR` varchar(256) NOT NULL COMMENT 'メールアドレス',
  `MAIL_STS` char(1) NOT NULL DEFAULT '0' COMMENT 'メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG',
  `MAIL_STS_DTE` timestamp NOT NULL DEFAULT '2000-01-01 00:00:00' COMMENT 'メールステータス更新日時',
  `CARR_KBN` char(2) DEFAULT NULL COMMENT 'キャリア区分',
  `CLIENT_ID` int(11) NOT NULL,
  `USER_OPE_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '0:メールアドレス表示, 1:メールアドレス記号表示\r\n0:運用管理者操作, 1:ユーザ操作',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_MAIL_ID`),
  KEY `USER_MAIL_MST_IDX_1` (`USER_ID`) USING BTREE,
  KEY `USER_MAIL_MST_IDX_2` (`CLIENT_ID`,`MAIL_ADDR`(255)) USING BTREE,
  KEY `USER_MAIL_MST_IDX_3` (`MAIL_STS_DTE`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1875532 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "user_mail_id": 3,    "user_id": 23,    "mail_addr_seq": 12,    "mail_addr": "IXfCnPLHspxIdjjIkcnuVZDun",    "mail_sts": "fCTmcvCiUXOPwlIpafBgvPwwY",    "mail_sts_dte": "2248-09-23T18:47:20.885259043+09:00",    "carr_kbn": "loVFqJshAjTWcesxiMJVWoLfy",    "client_id": 37,    "user_ope_flg": "jxcHeQGwwxmukFWbDOPxmZQWQ",    "upda_dte": "2066-06-23T07:56:38.028664371+09:00",    "upda_user_id": 70,    "crea_dte": "2305-02-06T17:02:46.116601867+09:00",    "crea_user_id": 59}



*/

// USER_MAIL_MST_ struct is a row record of the USER_MAIL_MST table in the anpidb database
type USER_MAIL_MST_ struct {
	//[ 0] USER_MAIL_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	USERMAILID int32 `gorm:"primary_key;AUTO_INCREMENT;column:USER_MAIL_ID;type:int;"` // ユーザメールID
	//[ 1] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 2] MAIL_ADDR_SEQ                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MAILADDRSEQ int32 `gorm:"column:MAIL_ADDR_SEQ;type:int;"` // メールアドレス連番
	//[ 3] MAIL_ADDR                                      varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	MAILADDR string `gorm:"column:MAIL_ADDR;type:varchar;size:256;"` // メールアドレス
	//[ 4] MAIL_STS                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSTS string `gorm:"column:MAIL_STS;type:char;size:1;default:0;"` // メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG
	//[ 5] MAIL_STS_DTE                                   timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [2000-01-01 00:00:00]
	MAILSTSDTE time.Time `gorm:"column:MAIL_STS_DTE;type:timestamp;default:2000-01-01 00:00:00;"` // メールステータス更新日時
	//[ 6] CARR_KBN                                       char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	CARRKBN sql.NullString `gorm:"column:CARR_KBN;type:char;size:2;"` // キャリア区分
	//[ 7] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"`
	//[ 8] USER_OPE_FLG                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	USEROPEFLG string `gorm:"column:USER_OPE_FLG;type:char;size:1;default:0;"` // 0:メールアドレス表示, 1:メールアドレス記号表示\r\n0:運用管理者操作, 1:ユーザ操作
	//[ 9] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[10] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[11] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[12] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_MAIL_MSTTableInfo = &TableInfo{
	Name: "USER_MAIL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_MAIL_ID",
			Comment:            `ユーザメールID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERMAILID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_mail_id",
			ProtobufFieldName:  "user_mail_id",
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "MAIL_ADDR_SEQ",
			Comment:            `メールアドレス連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MAILADDRSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_addr_seq",
			ProtobufFieldName:  "mail_addr_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "MAIL_STS",
			Comment: `メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSTS",
			GoFieldType:        "string",
			JSONFieldName:      "mail_sts",
			ProtobufFieldName:  "mail_sts",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CARR_KBN",
			Comment:            `キャリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "CARRKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "carr_kbn",
			ProtobufFieldName:  "carr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CLIENT_ID",
			Comment:            ``,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "USER_OPE_FLG",
			Comment:            `0:メールアドレス表示, 1:メールアドレス記号表示\r\n0:運用管理者操作, 1:ユーザ操作`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "USEROPEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "user_ope_flg",
			ProtobufFieldName:  "user_ope_flg",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *USER_MAIL_MST_) TableName() string {
	return "USER_MAIL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_MAIL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_MAIL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_MAIL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_MAIL_MST_) TableInfo() *TableInfo {
	return USER_MAIL_MSTTableInfo
}
