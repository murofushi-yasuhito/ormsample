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


CREATE TABLE `ADMIN_GRP_DTL_MST` (
  `ADMIN_GRP_ID` int(11) NOT NULL COMMENT '管理グループID',
  `ADMIN_GRP_SEQ` int(11) NOT NULL COMMENT '管理グループ連番',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `DEPT_AUTH_KBN` int(11) NOT NULL COMMENT '部署権限区分',
  `DEPT_START_MAIL_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '部署起動通知フラグ	 0:通知しない 1:通知する',
  `CLIENT_ID` int(11) NOT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ADMIN_GRP_ID`,`ADMIN_GRP_SEQ`),
  UNIQUE KEY `ADMIN_GRP_DTL_MST_PKI` (`ADMIN_GRP_ID`,`ADMIN_GRP_SEQ`) USING BTREE,
  KEY `CLIENT_ID` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "admin_grp_id": 40,    "admin_grp_seq": 66,    "user_id": 7,    "dept_auth_kbn": 38,    "dept_start_mail_flg": "KlqrvbgQeTOmxtLvdCNNQfxSO",    "client_id": 2,    "upda_dte": "2096-03-17T18:25:57.357804643+09:00",    "upda_user_id": 20,    "crea_dte": "2142-04-25T09:19:59.115264597+09:00",    "crea_user_id": 41}



*/

// ADMIN_GRP_DTL_MST struct is a row record of the ADMIN_GRP_DTL_MST table in the anpidb database
type ADMIN_GRP_DTL_MST struct {
	//[ 0] ADMIN_GRP_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ADMINGRPID int32 `gorm:"primary_key;column:ADMIN_GRP_ID;type:int;"` // 管理グループID
	//[ 1] ADMIN_GRP_SEQ                                  int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ADMINGRPSEQ int32 `gorm:"primary_key;column:ADMIN_GRP_SEQ;type:int;"` // 管理グループ連番
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 3] DEPT_AUTH_KBN                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTAUTHKBN int32 `gorm:"column:DEPT_AUTH_KBN;type:int;"` // 部署権限区分
	//[ 4] DEPT_START_MAIL_FLG                            char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DEPTSTARTMAILFLG string `gorm:"column:DEPT_START_MAIL_FLG;type:char;size:1;default:0;"` // 部署起動通知フラグ	 0:通知しない 1:通知する
	//[ 5] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"`
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ADMIN_GRP_DTL_MSTTableInfo = &TableInfo{
	Name: "ADMIN_GRP_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ADMIN_GRP_ID",
			Comment:            `管理グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ADMINGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "admin_grp_id",
			ProtobufFieldName:  "admin_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ADMIN_GRP_SEQ",
			Comment:            `管理グループ連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ADMINGRPSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "admin_grp_seq",
			ProtobufFieldName:  "admin_grp_seq",
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
			Name:               "DEPT_AUTH_KBN",
			Comment:            `部署権限区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTAUTHKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_auth_kbn",
			ProtobufFieldName:  "dept_auth_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "DEPT_START_MAIL_FLG",
			Comment: `部署起動通知フラグ	 0:通知しない 1:通知する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEPTSTARTMAILFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dept_start_mail_flg",
			ProtobufFieldName:  "dept_start_mail_flg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ADMIN_GRP_DTL_MST) TableName() string {
	return "ADMIN_GRP_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADMIN_GRP_DTL_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADMIN_GRP_DTL_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADMIN_GRP_DTL_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADMIN_GRP_DTL_MST) TableInfo() *TableInfo {
	return ADMIN_GRP_DTL_MSTTableInfo
}
