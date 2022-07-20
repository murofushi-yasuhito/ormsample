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


CREATE TABLE `USER_MST` (
  `USER_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `USER_NO` varchar(20) NOT NULL COMMENT 'ユーザ番号',
  `USER_NM` varchar(60) NOT NULL COMMENT 'ユーザ名',
  `USER_NMC` varchar(60) DEFAULT NULL COMMENT 'ユーザ名カナ',
  `LOGIN_ID` varchar(50) DEFAULT NULL COMMENT 'ログインID',
  `PASSWD` varbinary(140) DEFAULT NULL,
  `CLIENT_AUTH_KBN` int(11) DEFAULT NULL COMMENT '全社権限',
  `START_MAIL_FLG` char(1) DEFAULT '0' COMMENT '起動通知フラグ	 0:通知しない 1:通知する',
  `IKKATU_MAIL_FLG` char(1) DEFAULT '0' COMMENT '一括処理通知フラグ	 0:通知しない 1:通知する',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県コード	 居住地都道府県',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村コード	 居住地市区町村',
  `MAIL_STS` char(1) DEFAULT NULL COMMENT 'メールステータス',
  `BATCH_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '一括処理中フラグ',
  `TOKEN` varchar(255) DEFAULT '' COMMENT 'トークン',
  `DEVICE_TYPE` char(1) DEFAULT '' COMMENT 'デバイスタイプ',
  `MEMO` varchar(50) DEFAULT NULL,
  `PARTNER_FLG` char(1) NOT NULL DEFAULT '0' COMMENT 'パートナーフラグ 0:社員 1:パートナー（集計から除外する）',
  `SC_USER_LINK_FLG` char(1) DEFAULT '0' COMMENT 'ステータスChecker連携ユーザフラグ',
  `SC_MNG_AUTH` char(1) DEFAULT '0' COMMENT 'ステータスChecker全拠点権限',
  `SC_TEL_NUM` varchar(13) DEFAULT NULL COMMENT 'ステータスChecker電話番号',
  `DESIG_DEPT_AUTH_KBN` int(1) NOT NULL DEFAULT '0' COMMENT '指定部署権限',
  `INIT_PASSWD_FLG` char(1) NOT NULL DEFAULT '0' COMMENT 'ユーザ非作成パスワード（0:ユーザにて変更済み 1:初期または管理者作成のパスワード）',
  `ONE_TIME_PASSWD` char(16) DEFAULT NULL COMMENT 'ワンタイムパスワード',
  `ONE_TIME_EXPIRY` timestamp NULL DEFAULT NULL COMMENT 'ワンタイムパスワード有効期限',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  `START_GRP_ID` int(11) DEFAULT NULL,
  PRIMARY KEY (`USER_ID`),
  KEY `USER_MST_IDX_2` (`USER_NO`) USING BTREE,
  KEY `USER_MST_IDX_3` (`BATCH_FLG`) USING BTREE,
  KEY `USER_MST_IDX_4` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3512633 DEFAULT CHARSET=utf8 COMMENT='ユーザマスタ'

JSON Sample
-------------------------------------
{    "user_id": 89,    "client_id": 7,    "user_no": "IjEjPeKJboUHYDBTiWIEOsbSB",    "user_nm": "MBZIftrQqrtZbeVcwcHnDlUVO",    "user_nmc": "jwkQZxbrosxqvXqfsDMaMHmld",    "login_id": "VfXicQACAvvjiBRGjrRtSAatm",    "passwd": "KkYBFi1GLSogMj49OVYeFy5QSA8BA1lFLyMcVhNYFBJiAQIEKjA/PlYeVmJDCQs2QzAFJ0ENGUwPPVwtPQ==",    "client_auth_kbn": 98,    "start_mail_flg": "LkZKhLdMqmcWUwiDOHpLNTWQU",    "ikkatu_mail_flg": "ZEcSkiJXJndTOpcjSLePjFhyj",    "pref_cd": "nYEAwgQLgKeOakOUuLXRpGqWT",    "city_cd": "EOskUslKYhdNTfVkjjBCnQRrD",    "mail_sts": "WiQuMnCgBoTuWpcDiLkBMZPfL",    "batch_flg": "wnvesKGZEqFGuHoQCfdytDKZu",    "token": "pbFoBDdDqanYGhagIxUIYXtjK",    "device_type": "nEyuwSjRTbmgoBKCcvBbYidii",    "memo": "VSlpOVcJAwbTwYBeEVLshTWJo",    "partner_flg": "UBxMDLksJQxqLdqSVSFcidbQd",    "sc_user_link_flg": "ujcKYQsYsSKwvbhuGKdBnjMBM",    "sc_mng_auth": "OMkLHrRKkCNbRYLQBKAeJCaAD",    "sc_tel_num": "HPRsYIcXrvLmUwaprHIgliCRl",    "desig_dept_auth_kbn": 34,    "init_passwd_flg": "UIgBLtNkCoVBYmYmCNkttTVWm",    "one_time_passwd": "LxlcLlxhkxVYISlMOflwndjOR",    "one_time_expiry": "2034-04-14T18:43:52.162168258+09:00",    "upda_dte": "2115-03-27T03:15:02.459100639+09:00",    "upda_user_id": 36,    "crea_dte": "2065-12-05T17:26:07.831148998+09:00",    "crea_user_id": 72,    "start_grp_id": 95}



*/

// USER_MST struct is a row record of the USER_MST table in the anpidb database
type USER_MST struct {
	//[ 0] USER_ID                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	USERID int32 `gorm:"primary_key;AUTO_INCREMENT;column:USER_ID;type:int;"` // ユーザID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] USER_NO                                        varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	USERNO string `gorm:"column:USER_NO;type:varchar;size:20;"` // ユーザ番号
	//[ 3] USER_NM                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	USERNM string `gorm:"column:USER_NM;type:varchar;size:60;"` // ユーザ名
	//[ 4] USER_NMC                                       varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	USERNMC sql.NullString `gorm:"column:USER_NMC;type:varchar;size:60;"` // ユーザ名カナ
	//[ 5] LOGIN_ID                                       varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LOGINID sql.NullString `gorm:"column:LOGIN_ID;type:varchar;size:50;"` // ログインID
	//[ 6] PASSWD                                         varbinary            null: true   primary: false  isArray: false  auto: false  col: varbinary       len: -1      default: []
	PASSWD []byte `gorm:"column:PASSWD;type:varbinary;"`
	//[ 7] CLIENT_AUTH_KBN                                int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTAUTHKBN sql.NullInt64 `gorm:"column:CLIENT_AUTH_KBN;type:int;"` // 全社権限
	//[ 8] START_MAIL_FLG                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	STARTMAILFLG sql.NullString `gorm:"column:START_MAIL_FLG;type:char;size:1;default:0;"` // 起動通知フラグ	 0:通知しない 1:通知する
	//[ 9] IKKATU_MAIL_FLG                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	IKKATUMAILFLG sql.NullString `gorm:"column:IKKATU_MAIL_FLG;type:char;size:1;default:0;"` // 一括処理通知フラグ	 0:通知しない 1:通知する
	//[10] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード	 居住地都道府県
	//[11] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード	 居住地市区町村
	//[12] MAIL_STS                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSTS sql.NullString `gorm:"column:MAIL_STS;type:char;size:1;"` // メールステータス
	//[13] BATCH_FLG                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	BATCHFLG string `gorm:"column:BATCH_FLG;type:char;size:1;default:0;"` // 一括処理中フラグ
	//[14] TOKEN                                          varchar(255)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	TOKEN sql.NullString `gorm:"column:TOKEN;type:varchar;size:255;"` // トークン
	//[15] DEVICE_TYPE                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DEVICETYPE sql.NullString `gorm:"column:DEVICE_TYPE;type:char;size:1;"` // デバイスタイプ
	//[16] MEMO                                           varchar(50)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	MEMO sql.NullString `gorm:"column:MEMO;type:varchar;size:50;"`
	//[17] PARTNER_FLG                                    char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	PARTNERFLG string `gorm:"column:PARTNER_FLG;type:char;size:1;default:0;"` // パートナーフラグ 0:社員 1:パートナー（集計から除外する）
	//[18] SC_USER_LINK_FLG                               char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	SCUSERLINKFLG sql.NullString `gorm:"column:SC_USER_LINK_FLG;type:char;size:1;default:0;"` // ステータスChecker連携ユーザフラグ
	//[19] SC_MNG_AUTH                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	SCMNGAUTH sql.NullString `gorm:"column:SC_MNG_AUTH;type:char;size:1;default:0;"` // ステータスChecker全拠点権限
	//[20] SC_TEL_NUM                                     varchar(13)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 13      default: []
	SCTELNUM sql.NullString `gorm:"column:SC_TEL_NUM;type:varchar;size:13;"` // ステータスChecker電話番号
	//[21] DESIG_DEPT_AUTH_KBN                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DESIGDEPTAUTHKBN int32 `gorm:"column:DESIG_DEPT_AUTH_KBN;type:int;default:0;"` // 指定部署権限
	//[22] INIT_PASSWD_FLG                                char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	INITPASSWDFLG string `gorm:"column:INIT_PASSWD_FLG;type:char;size:1;default:0;"` // ユーザ非作成パスワード（0:ユーザにて変更済み 1:初期または管理者作成のパスワード）
	//[23] ONE_TIME_PASSWD                                char(16)             null: true   primary: false  isArray: false  auto: false  col: char            len: 16      default: []
	ONETIMEPASSWD sql.NullString `gorm:"column:ONE_TIME_PASSWD;type:char;size:16;"` // ワンタイムパスワード
	//[24] ONE_TIME_EXPIRY                                timestamp            null: true   primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: []
	ONETIMEEXPIRY time.Time `gorm:"column:ONE_TIME_EXPIRY;type:timestamp;"` // ワンタイムパスワード有効期限
	//[25] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[26] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[27] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[28] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID
	//[29] START_GRP_ID                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	STARTGRPID sql.NullInt64 `gorm:"column:START_GRP_ID;type:int;"`
}

var USER_MSTTableInfo = &TableInfo{
	Name: "USER_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "PASSWD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varbinary",
			DatabaseTypePretty: "varbinary",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varbinary",
			ColumnLength:       -1,
			GoFieldName:        "PASSWD",
			GoFieldType:        "[]byte",
			JSONFieldName:      "passwd",
			ProtobufFieldName:  "passwd",
			ProtobufType:       "bytes",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index: 8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index: 10,
			Name:  "PREF_CD",
			Comment: `都道府県コード	 居住地都道府県`,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index: 11,
			Name:  "CITY_CD",
			Comment: `市区町村コード	 居住地市区町村`,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "BATCH_FLG",
			Comment:            `一括処理中フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "BATCHFLG",
			GoFieldType:        "string",
			JSONFieldName:      "batch_flg",
			ProtobufFieldName:  "batch_flg",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "DEVICE_TYPE",
			Comment:            `デバイスタイプ`,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "MEMO",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "MEMO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "memo",
			ProtobufFieldName:  "memo",
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
			Name:               "SC_USER_LINK_FLG",
			Comment:            `ステータスChecker連携ユーザフラグ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "SCUSERLINKFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sc_user_link_flg",
			ProtobufFieldName:  "sc_user_link_flg",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "SC_MNG_AUTH",
			Comment:            `ステータスChecker全拠点権限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "SCMNGAUTH",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sc_mng_auth",
			ProtobufFieldName:  "sc_mng_auth",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "SC_TEL_NUM",
			Comment:            `ステータスChecker電話番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(13)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       13,
			GoFieldName:        "SCTELNUM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sc_tel_num",
			ProtobufFieldName:  "sc_tel_num",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "DESIG_DEPT_AUTH_KBN",
			Comment:            `指定部署権限`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DESIGDEPTAUTHKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "desig_dept_auth_kbn",
			ProtobufFieldName:  "desig_dept_auth_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "INIT_PASSWD_FLG",
			Comment:            `ユーザ非作成パスワード（0:ユーザにて変更済み 1:初期または管理者作成のパスワード）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "INITPASSWDFLG",
			GoFieldType:        "string",
			JSONFieldName:      "init_passwd_flg",
			ProtobufFieldName:  "init_passwd_flg",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "ONE_TIME_PASSWD",
			Comment:            `ワンタイムパスワード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(16)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       16,
			GoFieldName:        "ONETIMEPASSWD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "one_time_passwd",
			ProtobufFieldName:  "one_time_passwd",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "ONE_TIME_EXPIRY",
			Comment:            `ワンタイムパスワード有効期限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "ONETIMEEXPIRY",
			GoFieldType:        "time.Time",
			JSONFieldName:      "one_time_expiry",
			ProtobufFieldName:  "one_time_expiry",
			ProtobufType:       "uint64",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
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
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
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
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
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
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "START_GRP_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "start_grp_id",
			ProtobufFieldName:  "start_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        30,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *USER_MST) TableName() string {
	return "USER_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_MST) TableInfo() *TableInfo {
	return USER_MSTTableInfo
}
