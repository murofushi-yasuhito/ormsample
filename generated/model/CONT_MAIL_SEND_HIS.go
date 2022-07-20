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


CREATE TABLE `CONT_MAIL_SEND_HIS` (
  `MAIL_SEND_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'メール送信ID',
  `THRE_SID` int(11) NOT NULL COMMENT 'スレッドSID',
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `TOP_SID` int(11) NOT NULL COMMENT '親情報番号',
  `BEF_SID` int(11) NOT NULL COMMENT '継続情報番号',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `CATE1_CD` varchar(10) DEFAULT NULL COMMENT 'カテゴリ1コード',
  `CATE2_CD` varchar(10) DEFAULT NULL COMMENT 'カテゴリ2コード',
  `ELMNT_LVL` int(11) DEFAULT NULL COMMENT '影響レベル',
  `MAIL_SEND_FLG` char(1) DEFAULT '0' COMMENT 'メール配信フラグ	 1:配信済 0:未配信 2:メール送信不可 3:メール送信エラー 4:安否起動',
  `MAIL_SEND_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '配信日時',
  `MAIL_SEND_MSG` text COMMENT '配信結果メッセージ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MAIL_SEND_ID`),
  KEY `CONT_MAIL_SEND_HIS_IDX_1` (`CATE1_CD`,`USER_ID`) USING BTREE,
  KEY `CONT_MAIL_SEND_HIS_IDX_3` (`BEF_SID`) USING BTREE,
  KEY `CONT_MAIL_SEND_HIS_IDX_4` (`SID`,`MAIL_SEND_FLG`,`MAIL_SEND_DTE`) USING BTREE,
  KEY `CONT_MAIL_SEND_HIS_IDX_2` (`SID`) USING BTREE,
  KEY `CONT_MAIL_SEND_HIS_IDX_5` (`THRE_SID`,`USER_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14122722 DEFAULT CHARSET=utf8 COMMENT='コンテンツメール配信履歴'

JSON Sample
-------------------------------------
{    "mail_send_id": 17,    "thre_sid": 94,    "sid": 13,    "top_sid": 35,    "bef_sid": 31,    "user_id": 13,    "cate_1_cd": "CAfCmXbFoBqMgvjIEJQXYskqj",    "cate_2_cd": "tFkoEjcsJAYoWApfZQxBsVWtI",    "elmnt_lvl": 38,    "mail_send_flg": "yUITKBlqljuphqplEOqtwAiSR",    "mail_send_dte": "2155-02-18T16:20:53.326680822+09:00",    "mail_send_msg": "ZEjVIMGoJyIAlIKOMDmKvpHbv",    "upda_dte": "2163-10-15T23:30:30.495205532+09:00",    "upda_user_id": 76,    "crea_dte": "2250-12-07T12:36:08.860546197+09:00",    "crea_user_id": 77}



*/

// CONT_MAIL_SEND_HIS_ struct is a row record of the CONT_MAIL_SEND_HIS table in the anpidb database
type CONT_MAIL_SEND_HIS_ struct {
	//[ 0] MAIL_SEND_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MAILSENDID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MAIL_SEND_ID;type:int;"` // メール送信ID
	//[ 1] THRE_SID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	THRESID int32 `gorm:"column:THRE_SID;type:int;"` // スレッドSID
	//[ 2] SID                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"column:SID;type:int;"` // 情報番号
	//[ 3] TOP_SID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TOPSID int32 `gorm:"column:TOP_SID;type:int;"` // 親情報番号
	//[ 4] BEF_SID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BEFSID int32 `gorm:"column:BEF_SID;type:int;"` // 継続情報番号
	//[ 5] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 6] CATE1_CD                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD sql.NullString `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード
	//[ 7] CATE2_CD                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE2CD sql.NullString `gorm:"column:CATE2_CD;type:varchar;size:10;"` // カテゴリ2コード
	//[ 8] ELMNT_LVL                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ELMNTLVL sql.NullInt64 `gorm:"column:ELMNT_LVL;type:int;"` // 影響レベル
	//[ 9] MAIL_SEND_FLG                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSENDFLG sql.NullString `gorm:"column:MAIL_SEND_FLG;type:char;size:1;default:0;"` // メール配信フラグ	 1:配信済 0:未配信 2:メール送信不可 3:メール送信エラー 4:安否起動
	//[10] MAIL_SEND_DTE                                  timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	MAILSENDDTE time.Time `gorm:"column:MAIL_SEND_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 配信日時
	//[11] MAIL_SEND_MSG                                  text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	MAILSENDMSG sql.NullString `gorm:"column:MAIL_SEND_MSG;type:text;size:65535;"` // 配信結果メッセージ
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_MAIL_SEND_HISTableInfo = &TableInfo{
	Name: "CONT_MAIL_SEND_HIS",
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
			Name:               "THRE_SID",
			Comment:            `スレッドSID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "THRESID",
			GoFieldType:        "int32",
			JSONFieldName:      "thre_sid",
			ProtobufFieldName:  "thre_sid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "SID",
			Comment:            `情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SID",
			GoFieldType:        "int32",
			JSONFieldName:      "sid",
			ProtobufFieldName:  "sid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "TOP_SID",
			Comment:            `親情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPSID",
			GoFieldType:        "int32",
			JSONFieldName:      "top_sid",
			ProtobufFieldName:  "top_sid",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "BEF_SID",
			Comment:            `継続情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BEFSID",
			GoFieldType:        "int32",
			JSONFieldName:      "bef_sid",
			ProtobufFieldName:  "bef_sid",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CATE2_CD",
			Comment:            `カテゴリ2コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ELMNT_LVL",
			Comment:            `影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "elmnt_lvl",
			ProtobufFieldName:  "elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
			Name:  "MAIL_SEND_FLG",
			Comment: `メール配信フラグ	 1:配信済 0:未配信 2:メール送信不可 3:メール送信エラー 4:安否起動`,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "MAIL_SEND_DTE",
			Comment:            `配信日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "MAILSENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "mail_send_dte",
			ProtobufFieldName:  "mail_send_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "MAIL_SEND_MSG",
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
			GoFieldName:        "MAILSENDMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_send_msg",
			ProtobufFieldName:  "mail_send_msg",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_MAIL_SEND_HIS_) TableName() string {
	return "CONT_MAIL_SEND_HIS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_MAIL_SEND_HIS_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_MAIL_SEND_HIS_) TableInfo() *TableInfo {
	return CONT_MAIL_SEND_HISTableInfo
}
