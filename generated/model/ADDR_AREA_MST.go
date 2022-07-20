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


CREATE TABLE `ADDR_AREA_MST` (
  `ADDR_AREA_CD` char(2) NOT NULL COMMENT '住所エリアコード',
  `ADDR_AREA_NM` varchar(40) NOT NULL COMMENT '住所エリア名',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`ADDR_AREA_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "addr_area_cd": "FsJclrDgXDhOyjrkRvuWnMGxs",    "addr_area_nm": "kktwxwpLKUwNldfKLRYmVxqDl",    "disp_dsq": 16,    "dele_flg": "DRbcqPlcvUVXEmrTyctTDKTFr",    "upda_dte": "2036-10-03T08:29:11.211468144+09:00",    "upda_user_cd": "LbcbTTfaVDqQYnjsRfldWohnj",    "crea_dte": "2303-10-21T09:55:43.133606305+09:00",    "crea_user_cd": "fTCvgXpVSdRQVCWBXVgVPKSod"}



*/

// ADDR_AREA_MST_ struct is a row record of the ADDR_AREA_MST table in the anpidb database
type ADDR_AREA_MST_ struct {
	//[ 0] ADDR_AREA_CD                                   char(2)              null: false  primary: true   isArray: false  auto: false  col: char            len: 2       default: []
	ADDRAREACD string `gorm:"primary_key;column:ADDR_AREA_CD;type:char;size:2;"` // 住所エリアコード
	//[ 1] ADDR_AREA_NM                                   varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	ADDRAREANM string `gorm:"column:ADDR_AREA_NM;type:varchar;size:40;"` // 住所エリア名
	//[ 2] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 3] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"`
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"`
}

var ADDR_AREA_MSTTableInfo = &TableInfo{
	Name: "ADDR_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ADDR_AREA_CD",
			Comment:            `住所エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "ADDRAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "addr_area_cd",
			ProtobufFieldName:  "addr_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ADDR_AREA_NM",
			Comment:            `住所エリア名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "ADDRAREANM",
			GoFieldType:        "string",
			JSONFieldName:      "addr_area_nm",
			ProtobufFieldName:  "addr_area_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "int32",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			Name:               "UPDA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
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
			Name:               "CREA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ADDR_AREA_MST_) TableName() string {
	return "ADDR_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADDR_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADDR_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADDR_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADDR_AREA_MST_) TableInfo() *TableInfo {
	return ADDR_AREA_MSTTableInfo
}
