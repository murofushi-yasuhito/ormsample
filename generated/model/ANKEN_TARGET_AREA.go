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


CREATE TABLE `ANKEN_TARGET_AREA` (
  `ANKEN_TARGET_ID` int(11) NOT NULL COMMENT '送信先対象選択ID',
  `AREA_SEQ` int(11) NOT NULL COMMENT '地域連番',
  `AREA_DEPT_KBN` char(1) NOT NULL COMMENT '地域住所区分',
  `AREA_HUB_KBN` char(1) NOT NULL COMMENT '拠点住所区分',
  `AREA_HOME_KBN` char(1) NOT NULL DEFAULT '' COMMENT '居住地住所区分',
  `PREF_CD` char(2) NOT NULL COMMENT '都道府県CD',
  `CITY_CD` char(5) NOT NULL COMMENT '市区町村CD',
  `PREF_NM` varchar(10) NOT NULL DEFAULT '' COMMENT '都道府県名',
  `CITY_NM` varchar(40) NOT NULL DEFAULT '' COMMENT '市区町村名',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`,`AREA_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 48,    "area_seq": 47,    "area_dept_kbn": "IjpROgjANqUaMGTudPGhDObHC",    "area_hub_kbn": "yDnKDqrsVnhdTXvbJqNeliYMV",    "area_home_kbn": "iJqfjYSnlPMYWZvvBcEFVNeET",    "pref_cd": "MGtkAnrFuXYgGcMaZFhmBuUJI",    "city_cd": "FyIOuPOTdcssRnpQbOPqVZwrD",    "pref_nm": "GjTBJjoHbKvuNqxTNGOWnTcIN",    "city_nm": "eevNUrMoFGjBfeJZFjOvtPXKn",    "upda_dte": "2284-02-04T18:13:40.444075554+09:00",    "upda_user_id": 56,    "crea_dte": "2218-07-28T01:31:43.062470049+09:00",    "crea_user_id": 44}



*/

// ANKEN_TARGET_AREA_ struct is a row record of the ANKEN_TARGET_AREA table in the anpidb database
type ANKEN_TARGET_AREA_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] AREA_SEQ                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	AREASEQ int32 `gorm:"primary_key;column:AREA_SEQ;type:int;"` // 地域連番
	//[ 2] AREA_DEPT_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	AREADEPTKBN string `gorm:"column:AREA_DEPT_KBN;type:char;size:1;"` // 地域住所区分
	//[ 3] AREA_HUB_KBN                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	AREAHUBKBN string `gorm:"column:AREA_HUB_KBN;type:char;size:1;"` // 拠点住所区分
	//[ 4] AREA_HOME_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	AREAHOMEKBN string `gorm:"column:AREA_HOME_KBN;type:char;size:1;"` // 居住地住所区分
	//[ 5] PREF_CD                                        char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD string `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県CD
	//[ 6] CITY_CD                                        char(5)              null: false  primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD string `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村CD
	//[ 7] PREF_NM                                        varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	PREFNM string `gorm:"column:PREF_NM;type:varchar;size:10;"` // 都道府県名
	//[ 8] CITY_NM                                        varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	CITYNM string `gorm:"column:CITY_NM;type:varchar;size:40;"` // 市区町村名
	//[ 9] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[10] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[11] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[12] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_TARGET_AREATableInfo = &TableInfo{
	Name: "ANKEN_TARGET_AREA",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_TARGET_ID",
			Comment:            `送信先対象選択ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENTARGETID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_target_id",
			ProtobufFieldName:  "anken_target_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "AREA_SEQ",
			Comment:            `地域連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AREASEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "area_seq",
			ProtobufFieldName:  "area_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "AREA_DEPT_KBN",
			Comment:            `地域住所区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AREADEPTKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_dept_kbn",
			ProtobufFieldName:  "area_dept_kbn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "AREA_HUB_KBN",
			Comment:            `拠点住所区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AREAHUBKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_hub_kbn",
			ProtobufFieldName:  "area_hub_kbn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "AREA_HOME_KBN",
			Comment:            `居住地住所区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AREAHOMEKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_home_kbn",
			ProtobufFieldName:  "area_home_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "PREF_CD",
			Comment:            `都道府県CD`,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CITY_CD",
			Comment:            `市区町村CD`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "string",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "PREF_NM",
			Comment:            `都道府県名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "PREFNM",
			GoFieldType:        "string",
			JSONFieldName:      "pref_nm",
			ProtobufFieldName:  "pref_nm",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CITY_NM",
			Comment:            `市区町村名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "CITYNM",
			GoFieldType:        "string",
			JSONFieldName:      "city_nm",
			ProtobufFieldName:  "city_nm",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_TARGET_AREA_) TableName() string {
	return "ANKEN_TARGET_AREA"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_AREA_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_AREA_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_AREA_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_AREA_) TableInfo() *TableInfo {
	return ANKEN_TARGET_AREATableInfo
}
