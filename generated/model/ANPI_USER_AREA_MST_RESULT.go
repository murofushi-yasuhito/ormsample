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


CREATE TABLE `ANPI_USER_AREA_MST_RESULT` (
  `CLIENT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'クライアントID',
  `ERROR_CNT` int(11) NOT NULL DEFAULT '0' COMMENT 'エラーカウンター',
  `LAST_UPDATE_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '最終チェック日時',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`CLIENT_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=499 DEFAULT CHARSET=utf8 COMMENT='安否起動DLCTチェック結果'

JSON Sample
-------------------------------------
{    "client_id": 92,    "error_cnt": 88,    "last_update_dte": "2238-05-19T14:17:14.839728773+09:00",    "upda_dte": "2092-01-26T00:57:45.000482385+09:00",    "upda_user_id": 62,    "crea_dte": "2302-11-05T22:21:00.57986371+09:00",    "crea_user_id": 52}



*/

// ANPI_USER_AREA_MST_RESULT struct is a row record of the ANPI_USER_AREA_MST_RESULT table in the anpidb database
type ANPI_USER_AREA_MST_RESULT struct {
	//[ 0] CLIENT_ID                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CLIENTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:CLIENT_ID;type:int;"` // クライアントID
	//[ 1] ERROR_CNT                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	ERRORCNT int32 `gorm:"column:ERROR_CNT;type:int;default:0;"` // エラーカウンター
	//[ 2] LAST_UPDATE_DTE                                timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	LASTUPDATEDTE time.Time `gorm:"column:LAST_UPDATE_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 最終チェック日時
	//[ 3] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 4] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 6] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var ANPI_USER_AREA_MST_RESULTTableInfo = &TableInfo{
	Name: "ANPI_USER_AREA_MST_RESULT",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CLIENT_ID",
			Comment:            `クライアントID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ERROR_CNT",
			Comment:            `エラーカウンター`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ERRORCNT",
			GoFieldType:        "int32",
			JSONFieldName:      "error_cnt",
			ProtobufFieldName:  "error_cnt",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "LAST_UPDATE_DTE",
			Comment:            `最終チェック日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "LASTUPDATEDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "last_update_dte",
			ProtobufFieldName:  "last_update_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANPI_USER_AREA_MST_RESULT) TableName() string {
	return "ANPI_USER_AREA_MST_RESULT"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_USER_AREA_MST_RESULT) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_USER_AREA_MST_RESULT) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_USER_AREA_MST_RESULT) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_USER_AREA_MST_RESULT) TableInfo() *TableInfo {
	return ANPI_USER_AREA_MST_RESULTTableInfo
}
