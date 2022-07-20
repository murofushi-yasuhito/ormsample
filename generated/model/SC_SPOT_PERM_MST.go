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


CREATE TABLE `SC_SPOT_PERM_MST` (
  `USER_ID` int(11) unsigned zerofill NOT NULL COMMENT '安否のユーザID',
  `RW_SPOT_ID` varchar(32) NOT NULL COMMENT 'レスキューWEBの拠点ID',
  `SC_SPOT_ID` varchar(32) NOT NULL COMMENT 'ステータスCheckerの拠点ID',
  `PERMISSION` char(1) NOT NULL COMMENT 'ユーザが拠点に持つ権限 1:回答権限 2:閲覧権限 3:管理権限',
  `BATCH_FLG` char(1) DEFAULT NULL COMMENT '一括処理中フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_ID`,`RW_SPOT_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ステータスChecker拠点権限マスタ'

JSON Sample
-------------------------------------
{    "user_id": 76,    "rw_spot_id": "dNNDJoyUQUkwXZnKQKQELfbwq",    "sc_spot_id": "ouIlwKrufAjCInEKWmFslqvtA",    "permission": "BaZJHDhaYRDnESrLhjiZWjPaC",    "batch_flg": "WXRwKEUfmunsPvHqanpwQAQxW",    "upda_dte": "2147-05-14T04:33:30.165593487+09:00",    "upda_user_id": 92,    "crea_dte": "2152-12-25T05:16:20.846116519+09:00",    "crea_user_id": 90}


Comments
-------------------------------------
[ 0] column is set for unsigned



*/

// SC_SPOT_PERM_MST_ struct is a row record of the SC_SPOT_PERM_MST table in the anpidb database
type SC_SPOT_PERM_MST_ struct {
	//[ 0] USER_ID                                        uint                 null: false  primary: true   isArray: false  auto: false  col: uint            len: -1      default: []
	USERID uint32 `gorm:"primary_key;column:USER_ID;type:uint;"` // 安否のユーザID
	//[ 1] RW_SPOT_ID                                     varchar(32)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 32      default: []
	RWSPOTID string `gorm:"primary_key;column:RW_SPOT_ID;type:varchar;size:32;"` // レスキューWEBの拠点ID
	//[ 2] SC_SPOT_ID                                     varchar(32)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 32      default: []
	SCSPOTID string `gorm:"column:SC_SPOT_ID;type:varchar;size:32;"` // ステータスCheckerの拠点ID
	//[ 3] PERMISSION                                     char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	PERMISSION string `gorm:"column:PERMISSION;type:char;size:1;"` // ユーザが拠点に持つ権限 1:回答権限 2:閲覧権限 3:管理権限
	//[ 4] BATCH_FLG                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	BATCHFLG sql.NullString `gorm:"column:BATCH_FLG;type:char;size:1;"` // 一括処理中フラグ
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var SC_SPOT_PERM_MSTTableInfo = &TableInfo{
	Name: "SC_SPOT_PERM_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_ID",
			Comment:            `安否のユーザID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "uint32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "uint32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "RW_SPOT_ID",
			Comment:            `レスキューWEBの拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "RWSPOTID",
			GoFieldType:        "string",
			JSONFieldName:      "rw_spot_id",
			ProtobufFieldName:  "rw_spot_id",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "SC_SPOT_ID",
			Comment:            `ステータスCheckerの拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(32)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       32,
			GoFieldName:        "SCSPOTID",
			GoFieldType:        "string",
			JSONFieldName:      "sc_spot_id",
			ProtobufFieldName:  "sc_spot_id",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "PERMISSION",
			Comment:            `ユーザが拠点に持つ権限 1:回答権限 2:閲覧権限 3:管理権限`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PERMISSION",
			GoFieldType:        "string",
			JSONFieldName:      "permission",
			ProtobufFieldName:  "permission",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "BATCH_FLG",
			Comment:            `一括処理中フラグ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "BATCHFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "batch_flg",
			ProtobufFieldName:  "batch_flg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (s *SC_SPOT_PERM_MST_) TableName() string {
	return "SC_SPOT_PERM_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SC_SPOT_PERM_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SC_SPOT_PERM_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SC_SPOT_PERM_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SC_SPOT_PERM_MST_) TableInfo() *TableInfo {
	return SC_SPOT_PERM_MSTTableInfo
}
