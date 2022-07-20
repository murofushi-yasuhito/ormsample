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


CREATE TABLE `QES_TMPL_TYPE_MST` (
  `QES_TMPL_TYPE_ID` int(11) NOT NULL COMMENT 'テンプレート種別ID',
  `QES_TMPL_TYPE_NM` varchar(60) NOT NULL COMMENT 'テンプレート種別名',
  `QES_TMPL_TYPE_NM_E` varchar(60) NOT NULL COMMENT 'テンプレート種別名（英）',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`QES_TMPL_TYPE_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='テンプレートマスタ'

JSON Sample
-------------------------------------
{    "qes_tmpl_type_id": 86,    "qes_tmpl_type_nm": "CTIUoYkAUpGyrMWRfLQjbhpjD",    "qes_tmpl_type_nm_e": "xfwNlEfMKwpOhemnBniQWeCPD",    "upda_dte": "2054-06-13T23:35:47.756035966+09:00",    "upda_user_id": 67,    "crea_dte": "2084-09-02T07:38:31.313713852+09:00",    "crea_user_id": 34}



*/

// QES_TMPL_TYPE_MST_ struct is a row record of the QES_TMPL_TYPE_MST table in the anpidb database
type QES_TMPL_TYPE_MST_ struct {
	//[ 0] QES_TMPL_TYPE_ID                               int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	QESTMPLTYPEID int32 `gorm:"primary_key;column:QES_TMPL_TYPE_ID;type:int;"` // テンプレート種別ID
	//[ 1] QES_TMPL_TYPE_NM                               varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	QESTMPLTYPENM string `gorm:"column:QES_TMPL_TYPE_NM;type:varchar;size:60;"` // テンプレート種別名
	//[ 2] QES_TMPL_TYPE_NM_E                             varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	QESTMPLTYPENME string `gorm:"column:QES_TMPL_TYPE_NM_E;type:varchar;size:60;"` // テンプレート種別名（英）
	//[ 3] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 4] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 6] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var QES_TMPL_TYPE_MSTTableInfo = &TableInfo{
	Name: "QES_TMPL_TYPE_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "QES_TMPL_TYPE_ID",
			Comment:            `テンプレート種別ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "QESTMPLTYPEID",
			GoFieldType:        "int32",
			JSONFieldName:      "qes_tmpl_type_id",
			ProtobufFieldName:  "qes_tmpl_type_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "QES_TMPL_TYPE_NM",
			Comment:            `テンプレート種別名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "QESTMPLTYPENM",
			GoFieldType:        "string",
			JSONFieldName:      "qes_tmpl_type_nm",
			ProtobufFieldName:  "qes_tmpl_type_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "QES_TMPL_TYPE_NM_E",
			Comment:            `テンプレート種別名（英）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "QESTMPLTYPENME",
			GoFieldType:        "string",
			JSONFieldName:      "qes_tmpl_type_nm_e",
			ProtobufFieldName:  "qes_tmpl_type_nm_e",
			ProtobufType:       "string",
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
func (q *QES_TMPL_TYPE_MST_) TableName() string {
	return "QES_TMPL_TYPE_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (q *QES_TMPL_TYPE_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (q *QES_TMPL_TYPE_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (q *QES_TMPL_TYPE_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (q *QES_TMPL_TYPE_MST_) TableInfo() *TableInfo {
	return QES_TMPL_TYPE_MSTTableInfo
}
