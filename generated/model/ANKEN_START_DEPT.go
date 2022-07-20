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


CREATE TABLE `ANKEN_START_DEPT` (
  `ANKEN_START_DEPT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '案件起動部署ID',
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `DEPT_POS` int(11) DEFAULT '0' COMMENT '先頭部署POS',
  `DEPT_ID` int(11) DEFAULT '0' COMMENT '部署ID1',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_START_DEPT_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=683 DEFAULT CHARSET=utf8 COMMENT='案件起動部署'

JSON Sample
-------------------------------------
{    "anken_start_dept_id": 47,    "anken_id": 22,    "dept_pos": 23,    "dept_id": 37,    "crea_dte": "2121-07-15T01:51:07.226953635+09:00",    "crea_user_id": 15}



*/

// ANKEN_START_DEPT_ struct is a row record of the ANKEN_START_DEPT table in the anpidb database
type ANKEN_START_DEPT_ struct {
	//[ 0] ANKEN_START_DEPT_ID                            int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENSTARTDEPTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_START_DEPT_ID;type:int;"` // 案件起動部署ID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;"` // 案件ID
	//[ 2] DEPT_POS                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTPOS sql.NullInt64 `gorm:"column:DEPT_POS;type:int;default:0;"` // 先頭部署POS
	//[ 3] DEPT_ID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID sql.NullInt64 `gorm:"column:DEPT_ID;type:int;default:0;"` // 部署ID1
	//[ 4] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 5] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_START_DEPTTableInfo = &TableInfo{
	Name: "ANKEN_START_DEPT",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_START_DEPT_ID",
			Comment:            `案件起動部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENSTARTDEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_start_dept_id",
			ProtobufFieldName:  "anken_start_dept_id",
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
			Name:               "DEPT_POS",
			Comment:            `先頭部署POS`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTPOS",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_pos",
			ProtobufFieldName:  "dept_pos",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DEPT_ID",
			Comment:            `部署ID1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_START_DEPT_) TableName() string {
	return "ANKEN_START_DEPT"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_START_DEPT_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_START_DEPT_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_START_DEPT_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_START_DEPT_) TableInfo() *TableInfo {
	return ANKEN_START_DEPTTableInfo
}
