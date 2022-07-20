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


CREATE TABLE `ANPI_EXCEPT_AREA_MST` (
  `ANPI_EXCEPT_AREA_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '安否除外エリアID',
  `AREA_KBN` char(4) NOT NULL COMMENT 'エリア区分',
  `AREA_CD` varchar(10) NOT NULL COMMENT 'エリアコード',
  `AREA_NM` varchar(100) DEFAULT NULL COMMENT 'エリア名',
  `BIKO` varchar(300) DEFAULT NULL COMMENT '備考',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANPI_EXCEPT_AREA_ID`),
  KEY `ANPI_EXCEPT_AREA_MST_IDX1` (`AREA_KBN`,`AREA_CD`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8 COMMENT='安否起動除外エリアマスタ'

JSON Sample
-------------------------------------
{    "anpi_except_area_id": 62,    "area_kbn": "qfGIEYpxiNCYmWIVTdnkoaGav",    "area_cd": "wcKWNtUrMLrqCZMFsFBIKbEla",    "area_nm": "LZiXLaATEMcoLCeKcDswORPdT",    "biko": "fYNdlBxRCqQcXyRMgXGYgCAOH",    "dele_flg": "kxFnJcPeiCBIdcRYjRuxXCEIQ",    "upda_dte": "2229-01-27T07:15:35.544698846+09:00",    "upda_user_id": 19,    "crea_dte": "2122-11-21T19:05:28.483128062+09:00",    "crea_user_id": 58}



*/

// ANPI_EXCEPT_AREA_MST_ struct is a row record of the ANPI_EXCEPT_AREA_MST table in the anpidb database
type ANPI_EXCEPT_AREA_MST_ struct {
	//[ 0] ANPI_EXCEPT_AREA_ID                            int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANPIEXCEPTAREAID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANPI_EXCEPT_AREA_ID;type:int;"` // 安否除外エリアID
	//[ 1] AREA_KBN                                       char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:char;size:4;"` // エリア区分
	//[ 2] AREA_CD                                        varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AREACD string `gorm:"column:AREA_CD;type:varchar;size:10;"` // エリアコード
	//[ 3] AREA_NM                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	AREANM sql.NullString `gorm:"column:AREA_NM;type:varchar;size:100;"` // エリア名
	//[ 4] BIKO                                           varchar(300)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 300     default: []
	BIKO sql.NullString `gorm:"column:BIKO;type:varchar;size:300;"` // 備考
	//[ 5] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANPI_EXCEPT_AREA_MSTTableInfo = &TableInfo{
	Name: "ANPI_EXCEPT_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANPI_EXCEPT_AREA_ID",
			Comment:            `安否除外エリアID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPIEXCEPTAREAID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_except_area_id",
			ProtobufFieldName:  "anpi_except_area_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "AREA_KBN",
			Comment:            `エリア区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "AREAKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_kbn",
			ProtobufFieldName:  "area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "AREA_CD",
			Comment:            `エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD",
			GoFieldType:        "string",
			JSONFieldName:      "area_cd",
			ProtobufFieldName:  "area_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "AREA_NM",
			Comment:            `エリア名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "AREANM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_nm",
			ProtobufFieldName:  "area_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "BIKO",
			Comment:            `備考`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(300)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       300,
			GoFieldName:        "BIKO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "biko",
			ProtobufFieldName:  "biko",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANPI_EXCEPT_AREA_MST_) TableName() string {
	return "ANPI_EXCEPT_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_EXCEPT_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_EXCEPT_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_EXCEPT_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_EXCEPT_AREA_MST_) TableInfo() *TableInfo {
	return ANPI_EXCEPT_AREA_MSTTableInfo
}
