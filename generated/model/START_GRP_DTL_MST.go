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


CREATE TABLE `START_GRP_DTL_MST` (
  `START_GRP_ID` int(11) NOT NULL COMMENT '起動グループID',
  `START_GRP_SEQ` int(11) NOT NULL COMMENT '起動グループ連番',
  `PREF_CD` char(2) NOT NULL COMMENT '都道府県コード',
  `CITY_CD` varchar(5) DEFAULT NULL COMMENT '市区町村コード',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`START_GRP_ID`,`START_GRP_SEQ`),
  UNIQUE KEY `START_GRP_DTL_MST_PKI` (`START_GRP_ID`,`START_GRP_SEQ`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "start_grp_id": 42,    "start_grp_seq": 92,    "pref_cd": "VFlJEgVnrOsouBxFIMqmWyDro",    "city_cd": "bUGiYbyrpqeTTdYVvaNitrdmY",    "upda_dte": "2312-11-15T06:22:01.631973866+09:00",    "upda_user_id": 21,    "crea_dte": "2245-08-28T12:36:03.882020562+09:00",    "crea_user_id": 19}



*/

// START_GRP_DTL_MST_ struct is a row record of the START_GRP_DTL_MST table in the anpidb database
type START_GRP_DTL_MST_ struct {
	//[ 0] START_GRP_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	STARTGRPID int32 `gorm:"primary_key;column:START_GRP_ID;type:int;"` // 起動グループID
	//[ 1] START_GRP_SEQ                                  int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	STARTGRPSEQ int32 `gorm:"primary_key;column:START_GRP_SEQ;type:int;"` // 起動グループ連番
	//[ 2] PREF_CD                                        char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD string `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 3] CITY_CD                                        varchar(5)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:varchar;size:5;"` // 市区町村コード
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var START_GRP_DTL_MSTTableInfo = &TableInfo{
	Name: "START_GRP_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "START_GRP_ID",
			Comment:            `起動グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_id",
			ProtobufFieldName:  "start_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "START_GRP_SEQ",
			Comment:            `起動グループ連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_seq",
			ProtobufFieldName:  "start_grp_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "string",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
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
func (s *START_GRP_DTL_MST_) TableName() string {
	return "START_GRP_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *START_GRP_DTL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *START_GRP_DTL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *START_GRP_DTL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *START_GRP_DTL_MST_) TableInfo() *TableInfo {
	return START_GRP_DTL_MSTTableInfo
}
