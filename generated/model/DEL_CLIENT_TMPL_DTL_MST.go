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


CREATE TABLE `DEL_CLIENT_TMPL_DTL_MST` (
  `CLIENT_TMPL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'クライアントテンプレートID',
  `SEQ` int(11) NOT NULL COMMENT 'クライアントID',
  `CLIENT_TMPL_OPT2_MNG_ID` int(11) NOT NULL COMMENT 'テンプレート種別連番',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`CLIENT_TMPL_ID`),
  KEY `CLIENT_TMPL_MST_IDX_1` (`SEQ`)
) ENGINE=InnoDB AUTO_INCREMENT=3085 DEFAULT CHARSET=utf8 COMMENT='クライアントテンプレートマスタ'

JSON Sample
-------------------------------------
{    "client_tmpl_id": 39,    "seq": 48,    "client_tmpl_opt_2_mng_id": 99,    "upda_dte": "2247-10-29T23:19:01.001319993+09:00",    "upda_user_id": 50,    "crea_dte": "2241-07-20T00:44:28.57589949+09:00",    "crea_user_id": 56}



*/

// DEL_CLIENT_TMPL_DTL_MST struct is a row record of the DEL_CLIENT_TMPL_DTL_MST table in the anpidb database
type DEL_CLIENT_TMPL_DTL_MST struct {
	//[ 0] CLIENT_TMPL_ID                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CLIENTTMPLID int32 `gorm:"primary_key;AUTO_INCREMENT;column:CLIENT_TMPL_ID;type:int;"` // クライアントテンプレートID
	//[ 1] SEQ                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"column:SEQ;type:int;"` // クライアントID
	//[ 2] CLIENT_TMPL_OPT2_MNG_ID                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTTMPLOPT2MNGID int32 `gorm:"column:CLIENT_TMPL_OPT2_MNG_ID;type:int;"` // テンプレート種別連番
	//[ 3] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 4] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 6] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var DEL_CLIENT_TMPL_DTL_MSTTableInfo = &TableInfo{
	Name: "DEL_CLIENT_TMPL_DTL_MST",
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
			IsAutoIncrement:    true,
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
			Name:               "SEQ",
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
			GoFieldName:        "SEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CLIENT_TMPL_OPT2_MNG_ID",
			Comment:            `テンプレート種別連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTTMPLOPT2MNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_tmpl_opt_2_mng_id",
			ProtobufFieldName:  "client_tmpl_opt_2_mng_id",
			ProtobufType:       "int32",
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
func (d *DEL_CLIENT_TMPL_DTL_MST) TableName() string {
	return "DEL_CLIENT_TMPL_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEL_CLIENT_TMPL_DTL_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEL_CLIENT_TMPL_DTL_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEL_CLIENT_TMPL_DTL_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DEL_CLIENT_TMPL_DTL_MST) TableInfo() *TableInfo {
	return DEL_CLIENT_TMPL_DTL_MSTTableInfo
}
