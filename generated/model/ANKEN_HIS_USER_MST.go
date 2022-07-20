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


CREATE TABLE `ANKEN_HIS_USER_MST` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `USER_NO` varchar(20) NOT NULL COMMENT 'ユーザ番号',
  `USER_NM` varchar(60) NOT NULL COMMENT 'ユーザ名',
  `USER_NMC` varchar(60) DEFAULT NULL COMMENT 'ユーザ名カナ',
  `LOGIN_ID` varchar(50) DEFAULT NULL COMMENT 'ログインID',
  `PASSWD` varchar(70) DEFAULT NULL COMMENT 'パスワード',
  `CLIENT_AUTH_KBN` int(11) DEFAULT NULL COMMENT '全社権限',
  `START_MAIL_FLG` char(1) DEFAULT '0' COMMENT '起動通知フラグ	 0:通知しない 1:通知する',
  `IKKATU_MAIL_FLG` char(1) DEFAULT '0' COMMENT '一括処理通知フラグ	 0:通知しない 1:通知する',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県	 居住地都道府県',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村	 居住地市区町村',
  `MAIL_STS` char(1) DEFAULT NULL COMMENT 'メールステータス',
  `CUR_RESP_KBN` char(1) DEFAULT '0' COMMENT '最新応答区分 0:案件起動 1:自主応答',
  `TOKEN` varchar(255) DEFAULT NULL COMMENT 'トークン',
  `DEVICE_TYPE` char(1) DEFAULT NULL COMMENT 'デバイスタイプ 1:iOS, 2:Android',
  `PARTNER_FLG` char(1) NOT NULL DEFAULT '0' COMMENT 'パートナーフラグ 0:社員 1:パートナー（集計から除外する）',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`USER_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件履歴ユーザマスタ'

JSON Sample
-------------------------------------
{    "anken_id": 81,    "user_id": 13,    "client_id": 4,    "user_no": "FWmxcYHqJcgTMTBVXRRSbnIhc",    "user_nm": "NWBIuAOHNpgqIQrNFZMMYmJrY",    "user_nmc": "xEQBYimFVOZuoWGGyIgvCkJnH",    "login_id": "EmCDaYVJPwOPrUioxrWkhyfsN",    "passwd": "ktnMZVZNqVyfnhqprWFxxbutH",    "client_auth_kbn": 10,    "start_mail_flg": "FitIJPnpKqkytWqZjNBviMcHL",    "ikkatu_mail_flg": "tteiYOMIFLaLDMjYaBUjmrqxi",    "pref_cd": "NktYMamALNxMEWvrhRODCNqip",    "city_cd": "OewKTDcOXpCHbHmKpDAlcMeuB",    "mail_sts": "IANqZFcMxODDmScGaYKgvIBLA",    "cur_resp_kbn": "axIIIWZudrBVGJjhLesFbveAO",    "token": "kPRYxsvSWVgyOEHRaHNQMWhZL",    "device_type": "RjOmVhCTQrISBXvvjYFaIZyRS",    "partner_flg": "PaMvUAIEZPdMmiGFSVemvZGFY",    "upda_dte": "2276-06-27T01:59:14.687079984+09:00",    "upda_user_id": 80,    "crea_dte": "2105-03-16T01:54:49.034730654+09:00",    "crea_user_id": 2}



*/

// ANKEN_HIS_USER_MST_ struct is a row record of the ANKEN_HIS_USER_MST table in the anpidb database
type ANKEN_HIS_USER_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] USER_ID                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"primary_key;column:USER_ID;type:int;"` // ユーザID
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] USER_NO                                        varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	USERNO string `gorm:"column:USER_NO;type:varchar;size:20;"` // ユーザ番号
	//[ 4] USER_NM                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	USERNM string `gorm:"column:USER_NM;type:varchar;size:60;"` // ユーザ名
	//[ 5] USER_NMC                                       varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	USERNMC sql.NullString `gorm:"column:USER_NMC;type:varchar;size:60;"` // ユーザ名カナ
	//[ 6] LOGIN_ID                                       varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LOGINID sql.NullString `gorm:"column:LOGIN_ID;type:varchar;size:50;"` // ログインID
	//[ 7] PASSWD                                         varchar(70)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 70      default: []
	PASSWD sql.NullString `gorm:"column:PASSWD;type:varchar;size:70;"` // パスワード
	//[ 8] CLIENT_AUTH_KBN                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTAUTHKBN sql.NullInt64 `gorm:"column:CLIENT_AUTH_KBN;type:int;"` // 全社権限
	//[ 9] START_MAIL_FLG                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	STARTMAILFLG sql.NullString `gorm:"column:START_MAIL_FLG;type:char;size:1;default:0;"` // 起動通知フラグ	 0:通知しない 1:通知する
	//[10] IKKATU_MAIL_FLG                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	IKKATUMAILFLG sql.NullString `gorm:"column:IKKATU_MAIL_FLG;type:char;size:1;default:0;"` // 一括処理通知フラグ	 0:通知しない 1:通知する
	//[11] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県	 居住地都道府県
	//[12] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村	 居住地市区町村
	//[13] MAIL_STS                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSTS sql.NullString `gorm:"column:MAIL_STS;type:char;size:1;"` // メールステータス
	//[14] CUR_RESP_KBN                                   char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	CURRESPKBN sql.NullString `gorm:"column:CUR_RESP_KBN;type:char;size:1;default:0;"` // 最新応答区分 0:案件起動 1:自主応答
	//[15] TOKEN                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TOKEN sql.NullString `gorm:"column:TOKEN;type:varchar;size:255;"` // トークン
	//[16] DEVICE_TYPE                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DEVICETYPE sql.NullString `gorm:"column:DEVICE_TYPE;type:char;size:1;"` // デバイスタイプ 1:iOS, 2:Android
	//[17] PARTNER_FLG                                    char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	PARTNERFLG string `gorm:"column:PARTNER_FLG;type:char;size:1;default:0;"` // パートナーフラグ 0:社員 1:パートナー（集計から除外する）
	//[18] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[19] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[20] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[21] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_USER_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_USER_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
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
			Index:              3,
			Name:               "USER_NO",
			Comment:            `ユーザ番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "USERNO",
			GoFieldType:        "string",
			JSONFieldName:      "user_no",
			ProtobufFieldName:  "user_no",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "USER_NM",
			Comment:            `ユーザ名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "USERNM",
			GoFieldType:        "string",
			JSONFieldName:      "user_nm",
			ProtobufFieldName:  "user_nm",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "USER_NMC",
			Comment:            `ユーザ名カナ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "USERNMC",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_nmc",
			ProtobufFieldName:  "user_nmc",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "LOGIN_ID",
			Comment:            `ログインID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "LOGINID",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "login_id",
			ProtobufFieldName:  "login_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "PASSWD",
			Comment:            `パスワード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(70)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       70,
			GoFieldName:        "PASSWD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "passwd",
			ProtobufFieldName:  "passwd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CLIENT_AUTH_KBN",
			Comment:            `全社権限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTAUTHKBN",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "client_auth_kbn",
			ProtobufFieldName:  "client_auth_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
			Name:  "START_MAIL_FLG",
			Comment: `起動通知フラグ	 0:通知しない 1:通知する`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STARTMAILFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "start_mail_flg",
			ProtobufFieldName:  "start_mail_flg",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index: 10,
			Name:  "IKKATU_MAIL_FLG",
			Comment: `一括処理通知フラグ	 0:通知しない 1:通知する`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "IKKATUMAILFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "ikkatu_mail_flg",
			ProtobufFieldName:  "ikkatu_mail_flg",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index: 11,
			Name:  "PREF_CD",
			Comment: `都道府県	 居住地都道府県`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index: 12,
			Name:  "CITY_CD",
			Comment: `市区町村	 居住地市区町村`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "MAIL_STS",
			Comment:            `メールステータス`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "CUR_RESP_KBN",
			Comment:            `最新応答区分 0:案件起動 1:自主応答`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURRESPKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_resp_kbn",
			ProtobufFieldName:  "cur_resp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "TOKEN",
			Comment:            `トークン`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "TOKEN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "token",
			ProtobufFieldName:  "token",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "DEVICE_TYPE",
			Comment:            `デバイスタイプ 1:iOS, 2:Android`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEVICETYPE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "device_type",
			ProtobufFieldName:  "device_type",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "PARTNER_FLG",
			Comment:            `パートナーフラグ 0:社員 1:パートナー（集計から除外する）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PARTNERFLG",
			GoFieldType:        "string",
			JSONFieldName:      "partner_flg",
			ProtobufFieldName:  "partner_flg",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_HIS_USER_MST_) TableName() string {
	return "ANKEN_HIS_USER_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_USER_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_USER_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_USER_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_USER_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_USER_MSTTableInfo
}
