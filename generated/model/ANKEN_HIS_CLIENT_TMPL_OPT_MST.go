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


CREATE TABLE `ANKEN_HIS_CLIENT_TMPL_OPT_MST` (
  `ANKEN_ID` int(11) NOT NULL,
  `CLIENT_TMPL_ID` int(11) NOT NULL COMMENT 'クライアントテンプレートID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `STS_CD` char(1) NOT NULL COMMENT '状態コード',
  `STS_MSG` varchar(200) DEFAULT NULL,
  `STS_MSG_E` varchar(200) DEFAULT NULL,
  `RES_MSG` varchar(200) DEFAULT NULL,
  `RES_MSG_E` varchar(200) DEFAULT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  KEY `CLIENT_TMPL_OPT_MST_IDX_1` (`ANKEN_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='クライアントテンプレート選択肢マスタ'

JSON Sample
-------------------------------------
{    "anken_id": 71,    "client_tmpl_id": 10,    "seq": 0,    "sts_cd": "DGlmKyIrRIHghrQouHviCeAJd",    "sts_msg": "tteJSkBiKPOXlVddMeRHNmmyu",    "sts_msg_e": "OPXHnQENUTvPlNpExShPxflSY",    "res_msg": "DkKfgVvNmMtuXoYPXsxYkmRbx",    "res_msg_e": "OIGMuwCNLlNqkSiJZgsgLAaKT",    "upda_dte": "2058-10-20T02:37:34.925432768+09:00",    "upda_user_id": 48,    "crea_dte": "2089-05-10T08:03:29.817118095+09:00",    "crea_user_id": 81}


Comments
-------------------------------------
[ 0] Warning table: ANKEN_HIS_CLIENT_TMPL_OPT_MST does not have a primary key defined, setting col position 1 ANKEN_ID as primary key




*/

// ANKEN_HIS_CLIENT_TMPL_OPT_MST_ struct is a row record of the ANKEN_HIS_CLIENT_TMPL_OPT_MST table in the anpidb database
type ANKEN_HIS_CLIENT_TMPL_OPT_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"`
	//[ 1] CLIENT_TMPL_ID                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTTMPLID int32 `gorm:"column:CLIENT_TMPL_ID;type:int;"` // クライアントテンプレートID
	//[ 2] SEQ                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"column:SEQ;type:int;"` // 連番
	//[ 3] STS_CD                                         char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STSCD string `gorm:"column:STS_CD;type:char;size:1;"` // 状態コード
	//[ 4] STS_MSG                                        varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	STSMSG sql.NullString `gorm:"column:STS_MSG;type:varchar;size:200;"`
	//[ 5] STS_MSG_E                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	STSMSGE sql.NullString `gorm:"column:STS_MSG_E;type:varchar;size:200;"`
	//[ 6] RES_MSG                                        varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	RESMSG sql.NullString `gorm:"column:RES_MSG;type:varchar;size:200;"`
	//[ 7] RES_MSG_E                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	RESMSGE sql.NullString `gorm:"column:RES_MSG_E;type:varchar;size:200;"`
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_CLIENT_TMPL_OPT_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_CLIENT_TMPL_OPT_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "ANKEN_ID",
			Comment: ``,
			Notes: `Warning table: ANKEN_HIS_CLIENT_TMPL_OPT_MST does not have a primary key defined, setting col position 1 ANKEN_ID as primary key
`,
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
			Name:               "CLIENT_TMPL_ID",
			Comment:            `クライアントテンプレートID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTTMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_tmpl_id",
			ProtobufFieldName:  "client_tmpl_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "SEQ",
			Comment:            `連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "STS_CD",
			Comment:            `状態コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STSCD",
			GoFieldType:        "string",
			JSONFieldName:      "sts_cd",
			ProtobufFieldName:  "sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "STS_MSG",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "STSMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_msg",
			ProtobufFieldName:  "sts_msg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "STS_MSG_E",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "STSMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_msg_e",
			ProtobufFieldName:  "sts_msg_e",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "RES_MSG",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "RESMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg",
			ProtobufFieldName:  "res_msg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "RES_MSG_E",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "RESMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg_e",
			ProtobufFieldName:  "res_msg_e",
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
func (a *ANKEN_HIS_CLIENT_TMPL_OPT_MST_) TableName() string {
	return "ANKEN_HIS_CLIENT_TMPL_OPT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_CLIENT_TMPL_OPT_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_CLIENT_TMPL_OPT_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_CLIENT_TMPL_OPT_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_CLIENT_TMPL_OPT_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_CLIENT_TMPL_OPT_MSTTableInfo
}
