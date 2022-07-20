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


CREATE TABLE `ANKEN_STATUS_CUR_BK` (
  `ANKEN_STS_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '案件状態ID',
  `ANKEN_ID` int(11) NOT NULL DEFAULT '0' COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'ユーザID',
  `ACCS_KEY` varchar(20) NOT NULL COMMENT 'アクセスキー',
  `CUR_STS_CD` char(2) DEFAULT '99' COMMENT '最新状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C',
  `CUR_STS_CMNT` varchar(2000) DEFAULT NULL COMMENT '最新状態コメント',
  `CUR_ANS_DTE` datetime DEFAULT NULL COMMENT '最新回答日時',
  `CUR_FAMI_STS_CD` char(1) DEFAULT NULL COMMENT '家族安否状況コード',
  `CUR_QES1_STS_CD` char(1) DEFAULT NULL COMMENT '設問1ステータス',
  `CUR_QES2_STS_CD` char(1) DEFAULT NULL COMMENT '設問2ステータス',
  `CUR_QES3_STS_CD` char(1) DEFAULT NULL COMMENT '設問3ステータス',
  `CUR_QES4_STS_CD` char(1) DEFAULT NULL COMMENT '設問4ステータス',
  `CUR_QES5_STS_CD` char(1) DEFAULT NULL COMMENT '設問5ステータス',
  `AGENT_USER_ID` int(11) DEFAULT NULL COMMENT '代理応答ユーザID',
  `AGENT_USER_NO` varchar(20) DEFAULT NULL COMMENT '代理応答ユーザ番号',
  `AGENT_USER_NM` varchar(60) DEFAULT NULL COMMENT '代理応答ユーザ名',
  `RESP_KBN` char(1) DEFAULT '0' COMMENT '応答区分 0:案件起動 1:自主応答',
  `MAIL_STS` char(1) DEFAULT '0' COMMENT 'メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_STS_ID`),
  UNIQUE KEY `ANKEN_STATUS_CUR_IDX1` (`ANKEN_ID`,`USER_ID`),
  KEY `ANKEN_STATUS_CUR_IDX2` (`ACCS_KEY`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=39341486 DEFAULT CHARSET=utf8 COMMENT='案件状態最新	 案件状態最新情報'

JSON Sample
-------------------------------------
{    "anken_sts_id": 78,    "anken_id": 66,    "user_id": 95,    "accs_key": "YYlhPAgisxpjCnkBIRvPYrLNi",    "cur_sts_cd": "EFKdXwNBEjLRhonYfLpBnOxGc",    "cur_sts_cmnt": "tOkTbJXUWiQbDankpcsYFRxOx",    "cur_ans_dte": "2240-03-05T05:51:16.464662964+09:00",    "cur_fami_sts_cd": "yICJHFIOCfhpybNtnGtHVrYTs",    "cur_qes_1_sts_cd": "qRKZJohqmscqGwbhURpdRnHLB",    "cur_qes_2_sts_cd": "LIRVAnpUcxZGwqBTciLJAspJs",    "cur_qes_3_sts_cd": "oCrmgWiyVRBMJGwWwgaZMUSNo",    "cur_qes_4_sts_cd": "wXhZJOfskrqjhiiOWnFVFNajn",    "cur_qes_5_sts_cd": "wFaWRCAlkBxSHkbrDoBGySaAy",    "agent_user_id": 93,    "agent_user_no": "aygfFenKuntuMFIfSoaELDESi",    "agent_user_nm": "CplKFZYHxTjKSeUZGPLetMiup",    "resp_kbn": "iXepckthVIlFBCFoTQdxAxZlh",    "mail_sts": "RNpqlPCMavRCkQjdTEVKElXXG",    "upda_dte": "2092-08-23T13:00:37.532452238+09:00",    "upda_user_id": 99,    "crea_dte": "2217-10-28T12:04:18.496879316+09:00",    "crea_user_id": 70}



*/

// ANKEN_STATUS_CUR_BK struct is a row record of the ANKEN_STATUS_CUR_BK table in the anpidb database
type ANKEN_STATUS_CUR_BK struct {
	//[ 0] ANKEN_STS_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENSTSID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_STS_ID;type:int;"` // 案件状態ID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;default:0;"` // 案件ID
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	USERID int32 `gorm:"column:USER_ID;type:int;default:0;"` // ユーザID
	//[ 3] ACCS_KEY                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	ACCSKEY string `gorm:"column:ACCS_KEY;type:varchar;size:20;"` // アクセスキー
	//[ 4] CUR_STS_CD                                     char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: [99]
	CURSTSCD sql.NullString `gorm:"column:CUR_STS_CD;type:char;size:2;default:99;"` // 最新状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C
	//[ 5] CUR_STS_CMNT                                   varchar(2000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2000    default: []
	CURSTSCMNT sql.NullString `gorm:"column:CUR_STS_CMNT;type:varchar;size:2000;"` // 最新状態コメント
	//[ 6] CUR_ANS_DTE                                    datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CURANSDTE time.Time `gorm:"column:CUR_ANS_DTE;type:datetime;"` // 最新回答日時
	//[ 7] CUR_FAMI_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURFAMISTSCD sql.NullString `gorm:"column:CUR_FAMI_STS_CD;type:char;size:1;"` // 家族安否状況コード
	//[ 8] CUR_QES1_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURQES1STSCD sql.NullString `gorm:"column:CUR_QES1_STS_CD;type:char;size:1;"` // 設問1ステータス
	//[ 9] CUR_QES2_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURQES2STSCD sql.NullString `gorm:"column:CUR_QES2_STS_CD;type:char;size:1;"` // 設問2ステータス
	//[10] CUR_QES3_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURQES3STSCD sql.NullString `gorm:"column:CUR_QES3_STS_CD;type:char;size:1;"` // 設問3ステータス
	//[11] CUR_QES4_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURQES4STSCD sql.NullString `gorm:"column:CUR_QES4_STS_CD;type:char;size:1;"` // 設問4ステータス
	//[12] CUR_QES5_STS_CD                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CURQES5STSCD sql.NullString `gorm:"column:CUR_QES5_STS_CD;type:char;size:1;"` // 設問5ステータス
	//[13] AGENT_USER_ID                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AGENTUSERID sql.NullInt64 `gorm:"column:AGENT_USER_ID;type:int;"` // 代理応答ユーザID
	//[14] AGENT_USER_NO                                  varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AGENTUSERNO sql.NullString `gorm:"column:AGENT_USER_NO;type:varchar;size:20;"` // 代理応答ユーザ番号
	//[15] AGENT_USER_NM                                  varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AGENTUSERNM sql.NullString `gorm:"column:AGENT_USER_NM;type:varchar;size:60;"` // 代理応答ユーザ名
	//[16] RESP_KBN                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	RESPKBN sql.NullString `gorm:"column:RESP_KBN;type:char;size:1;default:0;"` // 応答区分 0:案件起動 1:自主応答
	//[17] MAIL_STS                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	MAILSTS sql.NullString `gorm:"column:MAIL_STS;type:char;size:1;default:0;"` // メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG
	//[18] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[19] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[20] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[21] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_STATUS_CUR_BKTableInfo = &TableInfo{
	Name: "ANKEN_STATUS_CUR_BK",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_STS_ID",
			Comment:            `案件状態ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENSTSID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_sts_id",
			ProtobufFieldName:  "anken_sts_id",
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
			Name:               "ACCS_KEY",
			Comment:            `アクセスキー`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "ACCSKEY",
			GoFieldType:        "string",
			JSONFieldName:      "accs_key",
			ProtobufFieldName:  "accs_key",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "CUR_STS_CD",
			Comment: `最新状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "CURSTSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_sts_cd",
			ProtobufFieldName:  "cur_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CUR_STS_CMNT",
			Comment:            `最新状態コメント`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2000,
			GoFieldName:        "CURSTSCMNT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_sts_cmnt",
			ProtobufFieldName:  "cur_sts_cmnt",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CUR_ANS_DTE",
			Comment:            `最新回答日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CURANSDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "cur_ans_dte",
			ProtobufFieldName:  "cur_ans_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CUR_FAMI_STS_CD",
			Comment:            `家族安否状況コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURFAMISTSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_fami_sts_cd",
			ProtobufFieldName:  "cur_fami_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CUR_QES1_STS_CD",
			Comment:            `設問1ステータス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURQES1STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_qes_1_sts_cd",
			ProtobufFieldName:  "cur_qes_1_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CUR_QES2_STS_CD",
			Comment:            `設問2ステータス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURQES2STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_qes_2_sts_cd",
			ProtobufFieldName:  "cur_qes_2_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "CUR_QES3_STS_CD",
			Comment:            `設問3ステータス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURQES3STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_qes_3_sts_cd",
			ProtobufFieldName:  "cur_qes_3_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "CUR_QES4_STS_CD",
			Comment:            `設問4ステータス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURQES4STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_qes_4_sts_cd",
			ProtobufFieldName:  "cur_qes_4_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "CUR_QES5_STS_CD",
			Comment:            `設問5ステータス`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CURQES5STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cur_qes_5_sts_cd",
			ProtobufFieldName:  "cur_qes_5_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "AGENT_USER_ID",
			Comment:            `代理応答ユーザID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AGENTUSERID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "agent_user_id",
			ProtobufFieldName:  "agent_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "AGENT_USER_NO",
			Comment:            `代理応答ユーザ番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "AGENTUSERNO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "agent_user_no",
			ProtobufFieldName:  "agent_user_no",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "AGENT_USER_NM",
			Comment:            `代理応答ユーザ名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "AGENTUSERNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "agent_user_nm",
			ProtobufFieldName:  "agent_user_nm",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "RESP_KBN",
			Comment:            `応答区分 0:案件起動 1:自主応答`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "RESPKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "resp_kbn",
			ProtobufFieldName:  "resp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index: 17,
			Name:  "MAIL_STS",
			Comment: `メールステータス	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG`,
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
func (a *ANKEN_STATUS_CUR_BK) TableName() string {
	return "ANKEN_STATUS_CUR_BK"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_STATUS_CUR_BK) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_STATUS_CUR_BK) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_STATUS_CUR_BK) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_STATUS_CUR_BK) TableInfo() *TableInfo {
	return ANKEN_STATUS_CUR_BKTableInfo
}
