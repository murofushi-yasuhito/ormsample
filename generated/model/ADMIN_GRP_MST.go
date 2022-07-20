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


CREATE TABLE `ADMIN_GRP_MST` (
  `ADMIN_GRP_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '管理グループID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `DEPT_ID` int(11) NOT NULL,
  `BATCH_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '荳諡ｬ蜃ｦ逅・ｸｭ繝輔Λ繧ｰ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ADMIN_GRP_ID`),
  KEY `ADMIN_GRP_MST_IDX_1` (`CLIENT_ID`) USING BTREE,
  KEY `BATCH_FLG` (`BATCH_FLG`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=36331 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "admin_grp_id": 62,    "client_id": 65,    "dept_id": 1,    "batch_flg": "SnFlINsvrSmPrXSmRmunoFWtx",    "upda_dte": "2248-11-10T03:46:42.088187513+09:00",    "upda_user_id": 20,    "crea_dte": "2045-09-03T20:44:51.973031238+09:00",    "crea_user_id": 4}



*/

// ADMIN_GRP_MST_ struct is a row record of the ADMIN_GRP_MST table in the anpidb database
type ADMIN_GRP_MST_ struct {
	//[ 0] ADMIN_GRP_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ADMINGRPID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ADMIN_GRP_ID;type:int;"` // 管理グループID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] DEPT_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID int32 `gorm:"column:DEPT_ID;type:int;"`
	//[ 3] BATCH_FLG                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	BATCHFLG string `gorm:"column:BATCH_FLG;type:char;size:1;default:0;"` // 荳諡ｬ蜃ｦ逅・ｸｭ繝輔Λ繧ｰ
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ADMIN_GRP_MSTTableInfo = &TableInfo{
	Name: "ADMIN_GRP_MST",
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
			IsAutoIncrement:    true,
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
			Name:               "DEPT_ID",
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
			GoFieldName:        "DEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "BATCH_FLG",
			Comment:            `荳諡ｬ蜃ｦ逅・ｸｭ繝輔Λ繧ｰ`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ADMIN_GRP_MST_) TableName() string {
	return "ADMIN_GRP_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADMIN_GRP_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADMIN_GRP_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADMIN_GRP_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADMIN_GRP_MST_) TableInfo() *TableInfo {
	return ADMIN_GRP_MSTTableInfo
}
