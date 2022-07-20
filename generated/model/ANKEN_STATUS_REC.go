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


CREATE TABLE `ANKEN_STATUS_REC` (
  `ANKEN_STS_REC_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '案件状況登録ID',
  `ANKEN_ID` int(11) NOT NULL DEFAULT '0' COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'ユーザID',
  `MSG_KND` char(1) NOT NULL COMMENT 'メッセージ種別	 M:メール C:コメント',
  `STS_CD` char(2) DEFAULT '99' COMMENT '状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C',
  `STS_CMNT` varchar(2000) DEFAULT NULL COMMENT '状態コメント',
  `MSG_TTL` varchar(100) DEFAULT NULL COMMENT 'メッセージ件名',
  `MSG_BODY` varchar(6000) DEFAULT NULL COMMENT 'メッセージ本文',
  `ANS_DTE` datetime DEFAULT NULL COMMENT '回答日時',
  `FAMI_STS_CD` char(1) DEFAULT NULL COMMENT '家族状況コード',
  `QES1_STS_CD` char(1) DEFAULT NULL COMMENT '設問1ステータス',
  `QES2_STS_CD` char(1) DEFAULT NULL COMMENT '設問2ステータス',
  `QES3_STS_CD` char(1) DEFAULT NULL COMMENT '設問3ステータス',
  `QES4_STS_CD` char(1) DEFAULT NULL COMMENT '設問4ステータス',
  `QES5_STS_CD` char(1) DEFAULT NULL COMMENT '設問5ステータス',
  `AGENT_USER_ID` int(11) DEFAULT NULL COMMENT '代理応答ユーザID',
  `AGENT_USER_NO` varchar(20) DEFAULT NULL COMMENT '代理応答ユーザNO',
  `AGENT_USER_NM` varchar(60) DEFAULT NULL COMMENT '代理応答ユーザNM',
  `RESP_KBN` char(1) DEFAULT '0' COMMENT '応答区分 0:案件起動 1:自主応答',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_STS_REC_ID`),
  KEY `ANKEN_STATUS_REC_IDX1` (`ANKEN_ID`,`USER_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=843870 DEFAULT CHARSET=utf8 COMMENT='案件状態登録'

JSON Sample
-------------------------------------
{    "anken_sts_rec_id": 56,    "anken_id": 42,    "user_id": 43,    "msg_knd": "SJSQpXYQgRdUskvGEbSQtHgno",    "sts_cd": "XIivqfMDbpLULnwjxajRyQmCH",    "sts_cmnt": "UtKOvmMUFdXBdkxRthgddDimh",    "msg_ttl": "uJbKBgDfSjeXhEtDwamRLLblM",    "msg_body": "wtwMBOHSMdHNtoWPxLjeVCVsI",    "ans_dte": "2310-12-04T12:24:47.426873878+09:00",    "fami_sts_cd": "TGGMnOKuqhcvnoWvIdCSYGhee",    "qes_1_sts_cd": "sAclQvXFigAxPJYgNhumfGaIl",    "qes_2_sts_cd": "tImqcnigBkUJbIZlsSbGpNVHx",    "qes_3_sts_cd": "AhYXuMnljIfbkNZCIKLISuVqs",    "qes_4_sts_cd": "jAHuZpimfnoZTrqwGJkceskfu",    "qes_5_sts_cd": "FCZLekskoZEqUYnqAjPQQPYPC",    "agent_user_id": 84,    "agent_user_no": "uSqcPcIYxCeLgOdyNbCCWfuRJ",    "agent_user_nm": "kRhvIwsHbUgrokqMynuaMSohs",    "resp_kbn": "PaBLQQuLeSeGXqRYWypbgewmp",    "crea_dte": "2191-06-27T11:31:40.642217754+09:00",    "crea_user_id": 37}



*/

// ANKEN_STATUS_REC_ struct is a row record of the ANKEN_STATUS_REC table in the anpidb database
type ANKEN_STATUS_REC_ struct {
	//[ 0] ANKEN_STS_REC_ID                               int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENSTSRECID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_STS_REC_ID;type:int;"` // 案件状況登録ID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;default:0;"` // 案件ID
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	USERID int32 `gorm:"column:USER_ID;type:int;default:0;"` // ユーザID
	//[ 3] MSG_KND                                        char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MSGKND string `gorm:"column:MSG_KND;type:char;size:1;"` // メッセージ種別	 M:メール C:コメント
	//[ 4] STS_CD                                         char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: [99]
	STSCD sql.NullString `gorm:"column:STS_CD;type:char;size:2;default:99;"` // 状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C
	//[ 5] STS_CMNT                                       varchar(2000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2000    default: []
	STSCMNT sql.NullString `gorm:"column:STS_CMNT;type:varchar;size:2000;"` // 状態コメント
	//[ 6] MSG_TTL                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	MSGTTL sql.NullString `gorm:"column:MSG_TTL;type:varchar;size:100;"` // メッセージ件名
	//[ 7] MSG_BODY                                       varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	MSGBODY sql.NullString `gorm:"column:MSG_BODY;type:varchar;size:6000;"` // メッセージ本文
	//[ 8] ANS_DTE                                        datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ANSDTE time.Time `gorm:"column:ANS_DTE;type:datetime;"` // 回答日時
	//[ 9] FAMI_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	FAMISTSCD sql.NullString `gorm:"column:FAMI_STS_CD;type:char;size:1;"` // 家族状況コード
	//[10] QES1_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QES1STSCD sql.NullString `gorm:"column:QES1_STS_CD;type:char;size:1;"` // 設問1ステータス
	//[11] QES2_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QES2STSCD sql.NullString `gorm:"column:QES2_STS_CD;type:char;size:1;"` // 設問2ステータス
	//[12] QES3_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QES3STSCD sql.NullString `gorm:"column:QES3_STS_CD;type:char;size:1;"` // 設問3ステータス
	//[13] QES4_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QES4STSCD sql.NullString `gorm:"column:QES4_STS_CD;type:char;size:1;"` // 設問4ステータス
	//[14] QES5_STS_CD                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QES5STSCD sql.NullString `gorm:"column:QES5_STS_CD;type:char;size:1;"` // 設問5ステータス
	//[15] AGENT_USER_ID                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	AGENTUSERID sql.NullInt64 `gorm:"column:AGENT_USER_ID;type:int;"` // 代理応答ユーザID
	//[16] AGENT_USER_NO                                  varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	AGENTUSERNO sql.NullString `gorm:"column:AGENT_USER_NO;type:varchar;size:20;"` // 代理応答ユーザNO
	//[17] AGENT_USER_NM                                  varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	AGENTUSERNM sql.NullString `gorm:"column:AGENT_USER_NM;type:varchar;size:60;"` // 代理応答ユーザNM
	//[18] RESP_KBN                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	RESPKBN sql.NullString `gorm:"column:RESP_KBN;type:char;size:1;default:0;"` // 応答区分 0:案件起動 1:自主応答
	//[19] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[20] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_STATUS_RECTableInfo = &TableInfo{
	Name: "ANKEN_STATUS_REC",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_STS_REC_ID",
			Comment:            `案件状況登録ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENSTSRECID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_sts_rec_id",
			ProtobufFieldName:  "anken_sts_rec_id",
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
			Index: 3,
			Name:  "MSG_KND",
			Comment: `メッセージ種別	 M:メール C:コメント`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MSGKND",
			GoFieldType:        "string",
			JSONFieldName:      "msg_knd",
			ProtobufFieldName:  "msg_knd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "STS_CD",
			Comment: `状態コード	 99:応答なし 90:メール返信 91:応答のみ A,B,C`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_cd",
			ProtobufFieldName:  "sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "STS_CMNT",
			Comment:            `状態コメント`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2000,
			GoFieldName:        "STSCMNT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_cmnt",
			ProtobufFieldName:  "sts_cmnt",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "MSG_TTL",
			Comment:            `メッセージ件名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "MSGTTL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "msg_ttl",
			ProtobufFieldName:  "msg_ttl",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "MSG_BODY",
			Comment:            `メッセージ本文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(6000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       6000,
			GoFieldName:        "MSGBODY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "msg_body",
			ProtobufFieldName:  "msg_body",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ANS_DTE",
			Comment:            `回答日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ANSDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "ans_dte",
			ProtobufFieldName:  "ans_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "FAMI_STS_CD",
			Comment:            `家族状況コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "FAMISTSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "fami_sts_cd",
			ProtobufFieldName:  "fami_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "QES1_STS_CD",
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
			GoFieldName:        "QES1STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_1_sts_cd",
			ProtobufFieldName:  "qes_1_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "QES2_STS_CD",
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
			GoFieldName:        "QES2STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_2_sts_cd",
			ProtobufFieldName:  "qes_2_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "QES3_STS_CD",
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
			GoFieldName:        "QES3STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_3_sts_cd",
			ProtobufFieldName:  "qes_3_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "QES4_STS_CD",
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
			GoFieldName:        "QES4STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_4_sts_cd",
			ProtobufFieldName:  "qes_4_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "QES5_STS_CD",
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
			GoFieldName:        "QES5STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_5_sts_cd",
			ProtobufFieldName:  "qes_5_sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "AGENT_USER_NO",
			Comment:            `代理応答ユーザNO`,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "AGENT_USER_NM",
			Comment:            `代理応答ユーザNM`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_STATUS_REC_) TableName() string {
	return "ANKEN_STATUS_REC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_STATUS_REC_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_STATUS_REC_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_STATUS_REC_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_STATUS_REC_) TableInfo() *TableInfo {
	return ANKEN_STATUS_RECTableInfo
}
