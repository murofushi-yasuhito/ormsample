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


CREATE TABLE `CLIENT_TMPL_ADD_QES_MST` (
  `CLIENT_TMPL_ID` int(11) NOT NULL COMMENT 'クライアントテンプレートID',
  `ADD_QES_TMPL_ID` int(11) NOT NULL COMMENT '使用する設問のCLIENT_TMPL_ID',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`CLIENT_TMPL_ID`,`ADD_QES_TMPL_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='クライアントテンプレート選択肢マスタ'

JSON Sample
-------------------------------------
{    "client_tmpl_id": 5,    "add_qes_tmpl_id": 24,    "upda_dte": "2147-02-03T15:18:04.869493587+09:00",    "upda_user_id": 70,    "crea_dte": "2298-11-07T13:54:05.741764347+09:00",    "crea_user_id": 34}



*/

// CLIENT_TMPL_ADD_QES_MST_ struct is a row record of the CLIENT_TMPL_ADD_QES_MST table in the anpidb database
type CLIENT_TMPL_ADD_QES_MST_ struct {
	//[ 0] CLIENT_TMPL_ID                                 int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTTMPLID int32 `gorm:"primary_key;column:CLIENT_TMPL_ID;type:int;"` // クライアントテンプレートID
	//[ 1] ADD_QES_TMPL_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ADDQESTMPLID int32 `gorm:"primary_key;column:ADD_QES_TMPL_ID;type:int;"` // 使用する設問のCLIENT_TMPL_ID
	//[ 2] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 3] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 4] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 5] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CLIENT_TMPL_ADD_QES_MSTTableInfo = &TableInfo{
	Name: "CLIENT_TMPL_ADD_QES_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CLIENT_TMPL_ID",
			Comment:            `クライアントテンプレートID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTTMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_tmpl_id",
			ProtobufFieldName:  "client_tmpl_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ADD_QES_TMPL_ID",
			Comment:            `使用する設問のCLIENT_TMPL_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ADDQESTMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "add_qes_tmpl_id",
			ProtobufFieldName:  "add_qes_tmpl_id",
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
func (c *CLIENT_TMPL_ADD_QES_MST_) TableName() string {
	return "CLIENT_TMPL_ADD_QES_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CLIENT_TMPL_ADD_QES_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CLIENT_TMPL_ADD_QES_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CLIENT_TMPL_ADD_QES_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CLIENT_TMPL_ADD_QES_MST_) TableInfo() *TableInfo {
	return CLIENT_TMPL_ADD_QES_MSTTableInfo
}
