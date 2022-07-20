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


CREATE TABLE `HOLIDAY_MST` (
  `HOLIDAY_MST_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '祝日マスタID',
  `HOLIDAY_DTE` date NOT NULL COMMENT '祝日',
  `HOLIDAY_TXT` varchar(100) NOT NULL COMMENT '祝日テキスト',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`HOLIDAY_MST_ID`),
  KEY `HOLIDAY_MST_IDX_1` (`HOLIDAY_DTE`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=146 DEFAULT CHARSET=utf8 COMMENT='祝日マスタ'

JSON Sample
-------------------------------------
{    "holiday_mst_id": 45,    "holiday_dte": "2150-08-04T20:52:02.089504498+09:00",    "holiday_txt": "bJUMoCEpCXLZIIKElmxGIUQvH",    "crea_dte": "2055-02-23T15:56:46.18390373+09:00",    "crea_user_id": 23}



*/

// HOLIDAY_MST_ struct is a row record of the HOLIDAY_MST table in the anpidb database
type HOLIDAY_MST_ struct {
	//[ 0] HOLIDAY_MST_ID                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	HOLIDAYMSTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:HOLIDAY_MST_ID;type:int;"` // 祝日マスタID
	//[ 1] HOLIDAY_DTE                                    date                 null: false  primary: false  isArray: false  auto: false  col: date            len: -1      default: []
	HOLIDAYDTE time.Time `gorm:"column:HOLIDAY_DTE;type:date;"` // 祝日
	//[ 2] HOLIDAY_TXT                                    varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	HOLIDAYTXT string `gorm:"column:HOLIDAY_TXT;type:varchar;size:100;"` // 祝日テキスト
	//[ 3] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 4] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var HOLIDAY_MSTTableInfo = &TableInfo{
	Name: "HOLIDAY_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "HOLIDAY_MST_ID",
			Comment:            `祝日マスタID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "HOLIDAYMSTID",
			GoFieldType:        "int32",
			JSONFieldName:      "holiday_mst_id",
			ProtobufFieldName:  "holiday_mst_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "HOLIDAY_DTE",
			Comment:            `祝日`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "date",
			DatabaseTypePretty: "date",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "date",
			ColumnLength:       -1,
			GoFieldName:        "HOLIDAYDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "holiday_dte",
			ProtobufFieldName:  "holiday_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "HOLIDAY_TXT",
			Comment:            `祝日テキスト`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "HOLIDAYTXT",
			GoFieldType:        "string",
			JSONFieldName:      "holiday_txt",
			ProtobufFieldName:  "holiday_txt",
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
func (h *HOLIDAY_MST_) TableName() string {
	return "HOLIDAY_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HOLIDAY_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HOLIDAY_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HOLIDAY_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *HOLIDAY_MST_) TableInfo() *TableInfo {
	return HOLIDAY_MSTTableInfo
}
