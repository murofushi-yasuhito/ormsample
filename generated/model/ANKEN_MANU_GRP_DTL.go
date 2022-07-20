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


CREATE TABLE `ANKEN_MANU_GRP_DTL` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `GRP_ID` int(11) NOT NULL COMMENT 'グループID',
  `GRP_NM` varchar(60) DEFAULT NULL COMMENT 'グループ名',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`GRP_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動グループ指定詳細'

JSON Sample
-------------------------------------
{    "anken_id": 96,    "grp_id": 93,    "grp_nm": "wsvCIXPgPZJOKCLqcVmFPWBuB",    "crea_dte": "2199-07-06T11:12:10.01165922+09:00",    "crea_user_id": 39}



*/

// ANKEN_MANU_GRP_DTL_ struct is a row record of the ANKEN_MANU_GRP_DTL table in the anpidb database
type ANKEN_MANU_GRP_DTL_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] GRP_ID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	GRPID int32 `gorm:"primary_key;column:GRP_ID;type:int;"` // グループID
	//[ 2] GRP_NM                                         varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	GRPNM sql.NullString `gorm:"column:GRP_NM;type:varchar;size:60;"` // グループ名
	//[ 3] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 4] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_MANU_GRP_DTLTableInfo = &TableInfo{
	Name: "ANKEN_MANU_GRP_DTL",
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
			Name:               "GRP_ID",
			Comment:            `グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_id",
			ProtobufFieldName:  "grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "GRP_NM",
			Comment:            `グループ名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "GRPNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "grp_nm",
			ProtobufFieldName:  "grp_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_MANU_GRP_DTL_) TableName() string {
	return "ANKEN_MANU_GRP_DTL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_MANU_GRP_DTL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_MANU_GRP_DTL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_MANU_GRP_DTL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_MANU_GRP_DTL_) TableInfo() *TableInfo {
	return ANKEN_MANU_GRP_DTLTableInfo
}
