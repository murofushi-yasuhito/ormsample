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


CREATE TABLE `BATCH_STATUS` (
  `STATUS_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '一括処理ステータス管理ID	 オートインクリメント',
  `TOP_STATUS_ID` int(11) NOT NULL COMMENT '親ステータスID	 親STATUS_IDを格納する。親レコードは0固定。',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID	 クライアントID',
  `PROC_DIV` int(11) DEFAULT NULL COMMENT '一括処理区分	 10:ユーザー 11:ユーザー差分 20:拠点 30:部署 40:役職 50:グループ 60:起動グループ 70:管理グループ 親レコード以外は0固定',
  `PROC_MODE` int(11) DEFAULT NULL COMMENT '一括処理モード	 0:チェック＆更新 1:チェックのみ 親レコード以外はNULL。',
  `PROC_STATUS` int(11) DEFAULT NULL COMMENT '処理ステータス	 1:一括処理中 2:一括処理完了 3:処理結果表示完了 親レコード以外は0固定',
  `CHECK_RESULT` int(11) unsigned NOT NULL DEFAULT '0' COMMENT 'チェック結果 1:エラーなし 2:エラーあり',
  `FILENAME` varchar(256) DEFAULT NULL COMMENT '処理ファイル名	 アップロードファイル名。親レコード以外はNULL。',
  `USER_FILENAME` varchar(256) DEFAULT NULL COMMENT 'ユーザアップロードファイル名',
  `LANGUAGE` int(11) DEFAULT NULL COMMENT '言語種別	 メッセージの言語種別 1:日本語 2:英語 親レコードはNULL',
  `MESSAGE` varchar(512) DEFAULT NULL COMMENT 'メッセージ	 メッセージ',
  `HOSTNAME` varchar(256) DEFAULT NULL COMMENT '一括処理実行APサーバホスト名',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時	 更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID	 更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時	 作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID	 作成者ID',
  PRIMARY KEY (`STATUS_ID`),
  UNIQUE KEY `STATUS_ID` (`STATUS_ID`) USING BTREE,
  KEY `BATCH_STATUS_IDX_1` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=742960 DEFAULT CHARSET=utf8 COMMENT='一括処理ステータス管理'

JSON Sample
-------------------------------------
{    "status_id": 23,    "top_status_id": 32,    "client_id": 95,    "proc_div": 76,    "proc_mode": 28,    "proc_status": 73,    "check_result": 72,    "filename": "hrcoZBlNVnyosiijcJTIAtgXn",    "user_filename": "sTiRgCyZVmFUyucYZmFcOlUIo",    "language": 61,    "message": "eGoXbdbenhgJpvrydgXlAQsVn",    "hostname": "aPrZnCLHkHVQRUPcojfodfhyH",    "upda_dte": "2250-01-07T00:56:03.644031498+09:00",    "upda_user_id": 27,    "crea_dte": "2266-06-30T19:45:00.137919533+09:00",    "crea_user_id": 73}


Comments
-------------------------------------
[ 6] column is set for unsigned



*/

// BATCH_STATUS_ struct is a row record of the BATCH_STATUS table in the anpidb database
type BATCH_STATUS_ struct {
	//[ 0] STATUS_ID                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	STATUSID int32 `gorm:"primary_key;AUTO_INCREMENT;column:STATUS_ID;type:int;"` // 一括処理ステータス管理ID	 オートインクリメント
	//[ 1] TOP_STATUS_ID                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TOPSTATUSID int32 `gorm:"column:TOP_STATUS_ID;type:int;"` // 親ステータスID	 親STATUS_IDを格納する。親レコードは0固定。
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID	 クライアントID
	//[ 3] PROC_DIV                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PROCDIV sql.NullInt64 `gorm:"column:PROC_DIV;type:int;"` // 一括処理区分	 10:ユーザー 11:ユーザー差分 20:拠点 30:部署 40:役職 50:グループ 60:起動グループ 70:管理グループ 親レコード以外は0固定
	//[ 4] PROC_MODE                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PROCMODE sql.NullInt64 `gorm:"column:PROC_MODE;type:int;"` // 一括処理モード	 0:チェック＆更新 1:チェックのみ 親レコード以外はNULL。
	//[ 5] PROC_STATUS                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PROCSTATUS sql.NullInt64 `gorm:"column:PROC_STATUS;type:int;"` // 処理ステータス	 1:一括処理中 2:一括処理完了 3:処理結果表示完了 親レコード以外は0固定
	//[ 6] CHECK_RESULT                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [0]
	CHECKRESULT uint32 `gorm:"column:CHECK_RESULT;type:uint;default:0;"` // チェック結果 1:エラーなし 2:エラーあり
	//[ 7] FILENAME                                       varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	FILENAME sql.NullString `gorm:"column:FILENAME;type:varchar;size:256;"` // 処理ファイル名	 アップロードファイル名。親レコード以外はNULL。
	//[ 8] USER_FILENAME                                  varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	USERFILENAME sql.NullString `gorm:"column:USER_FILENAME;type:varchar;size:256;"` // ユーザアップロードファイル名
	//[ 9] LANGUAGE                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LANGUAGE sql.NullInt64 `gorm:"column:LANGUAGE;type:int;"` // 言語種別	 メッセージの言語種別 1:日本語 2:英語 親レコードはNULL
	//[10] MESSAGE                                        varchar(512)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	MESSAGE sql.NullString `gorm:"column:MESSAGE;type:varchar;size:512;"` // メッセージ	 メッセージ
	//[11] HOSTNAME                                       varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	HOSTNAME sql.NullString `gorm:"column:HOSTNAME;type:varchar;size:256;"` // 一括処理実行APサーバホスト名
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時	 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID	 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時	 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID	 作成者ID

}

var BATCH_STATUSTableInfo = &TableInfo{
	Name: "BATCH_STATUS",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index: 0,
			Name:  "STATUS_ID",
			Comment: `一括処理ステータス管理ID	 オートインクリメント`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STATUSID",
			GoFieldType:        "int32",
			JSONFieldName:      "status_id",
			ProtobufFieldName:  "status_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index: 1,
			Name:  "TOP_STATUS_ID",
			Comment: `親ステータスID	 親STATUS_IDを格納する。親レコードは0固定。`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPSTATUSID",
			GoFieldType:        "int32",
			JSONFieldName:      "top_status_id",
			ProtobufFieldName:  "top_status_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index: 2,
			Name:  "CLIENT_ID",
			Comment: `クライアントID	 クライアントID`,
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
			Name:  "PROC_DIV",
			Comment: `一括処理区分	 10:ユーザー 11:ユーザー差分 20:拠点 30:部署 40:役職 50:グループ 60:起動グループ 70:管理グループ 親レコード以外は0固定`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PROCDIV",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "proc_div",
			ProtobufFieldName:  "proc_div",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "PROC_MODE",
			Comment: `一括処理モード	 0:チェック＆更新 1:チェックのみ 親レコード以外はNULL。`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PROCMODE",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "proc_mode",
			ProtobufFieldName:  "proc_mode",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "PROC_STATUS",
			Comment: `処理ステータス	 1:一括処理中 2:一括処理完了 3:処理結果表示完了 親レコード以外は0固定`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "PROCSTATUS",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "proc_status",
			ProtobufFieldName:  "proc_status",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CHECK_RESULT",
			Comment:            `チェック結果 1:エラーなし 2:エラーあり`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "CHECKRESULT",
			GoFieldType:        "uint32",
			JSONFieldName:      "check_result",
			ProtobufFieldName:  "check_result",
			ProtobufType:       "uint32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index: 7,
			Name:  "FILENAME",
			Comment: `処理ファイル名	 アップロードファイル名。親レコード以外はNULL。`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "FILENAME",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "filename",
			ProtobufFieldName:  "filename",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "USER_FILENAME",
			Comment:            `ユーザアップロードファイル名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "USERFILENAME",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_filename",
			ProtobufFieldName:  "user_filename",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
			Name:  "LANGUAGE",
			Comment: `言語種別	 メッセージの言語種別 1:日本語 2:英語 親レコードはNULL`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LANGUAGE",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "language",
			ProtobufFieldName:  "language",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index: 10,
			Name:  "MESSAGE",
			Comment: `メッセージ	 メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "MESSAGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "HOSTNAME",
			Comment:            `一括処理実行APサーバホスト名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "HOSTNAME",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "hostname",
			ProtobufFieldName:  "hostname",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index: 12,
			Name:  "UPDA_DTE",
			Comment: `更新日時	 更新日時`,
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
			Index: 13,
			Name:  "UPDA_USER_ID",
			Comment: `更新者ID	 更新者ID`,
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
			Index: 14,
			Name:  "CREA_DTE",
			Comment: `作成日時	 作成日時`,
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
			Index: 15,
			Name:  "CREA_USER_ID",
			Comment: `作成者ID	 作成者ID`,
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
func (b *BATCH_STATUS_) TableName() string {
	return "BATCH_STATUS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (b *BATCH_STATUS_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (b *BATCH_STATUS_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (b *BATCH_STATUS_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (b *BATCH_STATUS_) TableInfo() *TableInfo {
	return BATCH_STATUSTableInfo
}
