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


CREATE TABLE `AUTH_DTL_MST` (
  `AUTH_KBN` int(11) NOT NULL COMMENT '権限区分',
  `INN_AUTH_KBN` int(11) NOT NULL COMMENT '内部権限区分',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) DEFAULT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) DEFAULT NULL COMMENT '作成者ID',
  PRIMARY KEY (`AUTH_KBN`,`INN_AUTH_KBN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='権限詳細マスタ'

JSON Sample
-------------------------------------
{    "auth_kbn": 87,    "inn_auth_kbn": 48,    "upda_dte": "2301-05-22T15:28:02.19534922+09:00",    "upda_user_id": 82,    "crea_dte": "2288-10-26T09:49:20.955853763+09:00",    "crea_user_id": 17}



*/

// AUTH_DTL_MST_ struct is a row record of the AUTH_DTL_MST table in the anpidb database
type AUTH_DTL_MST_ struct {
	//[ 0] AUTH_KBN                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	AUTHKBN int32 `gorm:"primary_key;column:AUTH_KBN;type:int;"` // 権限区分
	//[ 1] INN_AUTH_KBN                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	INNAUTHKBN int32 `gorm:"primary_key;column:INN_AUTH_KBN;type:int;"` // 内部権限区分
	//[ 2] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 3] UPDA_USER_ID                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID sql.NullInt64 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 4] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 5] CREA_USER_ID                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID sql.NullInt64 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var AUTH_DTL_MSTTableInfo = &TableInfo{
	Name: "AUTH_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "AUTH_KBN",
			Comment:            `権限区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AUTHKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "auth_kbn",
			ProtobufFieldName:  "auth_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "INN_AUTH_KBN",
			Comment:            `内部権限区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "INNAUTHKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "inn_auth_kbn",
			ProtobufFieldName:  "inn_auth_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "UPDA_USER_ID",
			Comment:            `更新者ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDAUSERID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "upda_user_id",
			ProtobufFieldName:  "upda_user_id",
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
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CREAUSERID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "crea_user_id",
			ProtobufFieldName:  "crea_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *AUTH_DTL_MST_) TableName() string {
	return "AUTH_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *AUTH_DTL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *AUTH_DTL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *AUTH_DTL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *AUTH_DTL_MST_) TableInfo() *TableInfo {
	return AUTH_DTL_MSTTableInfo
}
