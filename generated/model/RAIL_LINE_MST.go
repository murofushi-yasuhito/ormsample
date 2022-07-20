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


CREATE TABLE `RAIL_LINE_MST` (
  `RAIL_LINE_ID` int(11) NOT NULL COMMENT '鉄道路線・区間ID',
  `RAIL_LINE_CD` varchar(8) NOT NULL COMMENT '鉄道路線・区間コード',
  `RAIL_LINE_FLG` char(1) NOT NULL COMMENT '鉄道路線・区間フラグ',
  `RAIL_LINE_NM` varchar(60) NOT NULL COMMENT '鉄道路線・区間名',
  `RAIL_LINE_NMC` varchar(60) DEFAULT NULL COMMENT '鉄道路線・区間名(かな)',
  `RAIL_AREA_CD` char(2) NOT NULL COMMENT '鉄道エリアコード',
  `RAIL_COMP_CD` char(4) NOT NULL COMMENT '鉄道会社コード',
  `RAIL_KIND` char(2) NOT NULL COMMENT '路線種別',
  `DELI_RAIL_FLG` char(1) NOT NULL COMMENT '配信対象路線フラグ',
  `LOOP_LINE_FLG` char(1) NOT NULL COMMENT '環状線フラグ',
  `HAVE_P4_FLG` char(1) NOT NULL DEFAULT '0' COMMENT 'P4保持フラグ',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `STAR_DTE` datetime NOT NULL COMMENT '開業日',
  `END_DTE` datetime NOT NULL COMMENT '閉業日',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) NOT NULL COMMENT '更新者コード',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) NOT NULL COMMENT '作成者コード',
  PRIMARY KEY (`RAIL_LINE_ID`),
  KEY `RAIL_LINE_MST_TMP_IDX_1` (`RAIL_LINE_CD`),
  KEY `RAIL_LINE_MST_TMP_IDX_2` (`RAIL_AREA_CD`,`RAIL_COMP_CD`,`RAIL_LINE_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "rail_line_id": 62,    "rail_line_cd": "EIvqAkbPCyCPYvuslRiOoTiOy",    "rail_line_flg": "IoMqOGFelRmyqrNXAKGFubnjb",    "rail_line_nm": "ZhvmYkkgTGixbgJfrljfBbctd",    "rail_line_nmc": "BRAUhyRmwwgqqkwgIctKDXRiH",    "rail_area_cd": "EpUiODJPeLqWElJKKpfkEhJMW",    "rail_comp_cd": "ECmoJXrGwLtQoLpGsnONKlnTb",    "rail_kind": "DOQnSuyCNIYkMdWEmtHYvhBqT",    "deli_rail_flg": "gqcoIcKeFxpVxMKliEgnWQMQQ",    "loop_line_flg": "hURPISQHxGNvfHyQYDGpcSISo",    "have_p_4_flg": "TEUwhGwbIZmlpKseJmBboBTGk",    "disp_dsq": 40,    "dele_flg": "cJAfXKDWRHNiGVdSJWXZbxkAb",    "star_dte": "2212-05-19T00:56:24.850438361+09:00",    "end_dte": "2054-07-28T07:24:48.332066259+09:00",    "upda_dte": "2058-10-07T22:12:04.549870323+09:00",    "upda_user_cd": "RyFXeQZFoyaPjwtjiHUeTyBgx",    "crea_dte": "2057-01-26T05:48:31.753943517+09:00",    "crea_user_cd": "GpMBVYMpVVZUxodxHKTvtWsII"}



*/

// RAIL_LINE_MST_ struct is a row record of the RAIL_LINE_MST table in the anpidb database
type RAIL_LINE_MST_ struct {
	//[ 0] RAIL_LINE_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	RAILLINEID int32 `gorm:"primary_key;column:RAIL_LINE_ID;type:int;"` // 鉄道路線・区間ID
	//[ 1] RAIL_LINE_CD                                   varchar(8)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 8       default: []
	RAILLINECD string `gorm:"column:RAIL_LINE_CD;type:varchar;size:8;"` // 鉄道路線・区間コード
	//[ 2] RAIL_LINE_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	RAILLINEFLG string `gorm:"column:RAIL_LINE_FLG;type:char;size:1;"` // 鉄道路線・区間フラグ
	//[ 3] RAIL_LINE_NM                                   varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	RAILLINENM string `gorm:"column:RAIL_LINE_NM;type:varchar;size:60;"` // 鉄道路線・区間名
	//[ 4] RAIL_LINE_NMC                                  varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	RAILLINENMC sql.NullString `gorm:"column:RAIL_LINE_NMC;type:varchar;size:60;"` // 鉄道路線・区間名(かな)
	//[ 5] RAIL_AREA_CD                                   char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	RAILAREACD string `gorm:"column:RAIL_AREA_CD;type:char;size:2;"` // 鉄道エリアコード
	//[ 6] RAIL_COMP_CD                                   char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	RAILCOMPCD string `gorm:"column:RAIL_COMP_CD;type:char;size:4;"` // 鉄道会社コード
	//[ 7] RAIL_KIND                                      char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	RAILKIND string `gorm:"column:RAIL_KIND;type:char;size:2;"` // 路線種別
	//[ 8] DELI_RAIL_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELIRAILFLG string `gorm:"column:DELI_RAIL_FLG;type:char;size:1;"` // 配信対象路線フラグ
	//[ 9] LOOP_LINE_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	LOOPLINEFLG string `gorm:"column:LOOP_LINE_FLG;type:char;size:1;"` // 環状線フラグ
	//[10] HAVE_P4_FLG                                    char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	HAVEP4FLG string `gorm:"column:HAVE_P4_FLG;type:char;size:1;default:0;"` // P4保持フラグ
	//[11] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[12] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[13] STAR_DTE                                       datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	STARDTE time.Time `gorm:"column:STAR_DTE;type:datetime;"` // 開業日
	//[14] END_DTE                                        datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ENDDTE time.Time `gorm:"column:END_DTE;type:datetime;"` // 閉業日
	//[15] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[16] UPDA_USER_CD                                   varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD string `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"` // 更新者コード
	//[17] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[18] CREA_USER_CD                                   varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD string `gorm:"column:CREA_USER_CD;type:varchar;size:10;"` // 作成者コード

}

var RAIL_LINE_MSTTableInfo = &TableInfo{
	Name: "RAIL_LINE_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "RAIL_LINE_ID",
			Comment:            `鉄道路線・区間ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RAILLINEID",
			GoFieldType:        "int32",
			JSONFieldName:      "rail_line_id",
			ProtobufFieldName:  "rail_line_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "RAIL_LINE_CD",
			Comment:            `鉄道路線・区間コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(8)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       8,
			GoFieldName:        "RAILLINECD",
			GoFieldType:        "string",
			JSONFieldName:      "rail_line_cd",
			ProtobufFieldName:  "rail_line_cd",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "RAIL_LINE_FLG",
			Comment:            `鉄道路線・区間フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "RAILLINEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "rail_line_flg",
			ProtobufFieldName:  "rail_line_flg",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "RAIL_LINE_NM",
			Comment:            `鉄道路線・区間名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "RAILLINENM",
			GoFieldType:        "string",
			JSONFieldName:      "rail_line_nm",
			ProtobufFieldName:  "rail_line_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "RAIL_LINE_NMC",
			Comment:            `鉄道路線・区間名(かな)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "RAILLINENMC",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "rail_line_nmc",
			ProtobufFieldName:  "rail_line_nmc",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "RAIL_AREA_CD",
			Comment:            `鉄道エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "RAILAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "rail_area_cd",
			ProtobufFieldName:  "rail_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "RAIL_COMP_CD",
			Comment:            `鉄道会社コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "RAILCOMPCD",
			GoFieldType:        "string",
			JSONFieldName:      "rail_comp_cd",
			ProtobufFieldName:  "rail_comp_cd",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "RAIL_KIND",
			Comment:            `路線種別`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "RAILKIND",
			GoFieldType:        "string",
			JSONFieldName:      "rail_kind",
			ProtobufFieldName:  "rail_kind",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "DELI_RAIL_FLG",
			Comment:            `配信対象路線フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELIRAILFLG",
			GoFieldType:        "string",
			JSONFieldName:      "deli_rail_flg",
			ProtobufFieldName:  "deli_rail_flg",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "LOOP_LINE_FLG",
			Comment:            `環状線フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LOOPLINEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "loop_line_flg",
			ProtobufFieldName:  "loop_line_flg",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "HAVE_P4_FLG",
			Comment:            `P4保持フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "HAVEP4FLG",
			GoFieldType:        "string",
			JSONFieldName:      "have_p_4_flg",
			ProtobufFieldName:  "have_p_4_flg",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "STAR_DTE",
			Comment:            `開業日`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "STARDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "star_dte",
			ProtobufFieldName:  "star_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "END_DTE",
			Comment:            `閉業日`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "end_dte",
			ProtobufFieldName:  "end_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "UPDA_USER_CD",
			Comment:            `更新者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "CREA_USER_CD",
			Comment:            `作成者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (r *RAIL_LINE_MST_) TableName() string {
	return "RAIL_LINE_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (r *RAIL_LINE_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (r *RAIL_LINE_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (r *RAIL_LINE_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (r *RAIL_LINE_MST_) TableInfo() *TableInfo {
	return RAIL_LINE_MSTTableInfo
}
