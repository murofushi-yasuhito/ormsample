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


CREATE TABLE `ANKEN_HIS_DEPT_MST` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `DEPT_ID` int(11) NOT NULL COMMENT '部署ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `TOP_DEPT_ID` int(11) NOT NULL COMMENT '先頭部署ID',
  `UP_DEPT_ID` int(11) DEFAULT NULL COMMENT '上部署ID',
  `DEPT_POS` int(11) NOT NULL COMMENT '部署ポジション	 1～10',
  `DEPT_NM` varchar(200) NOT NULL COMMENT '部署名',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県コード',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村コード',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DISP_DSQ_SUB_NO` int(11) NOT NULL COMMENT 'ポジション別表示順	 同階層内での表示順',
  `MISS_TEETH_FLG` char(1) DEFAULT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`DEPT_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件履歴部署マスタ'

JSON Sample
-------------------------------------
{    "anken_id": 33,    "dept_id": 76,    "client_id": 97,    "top_dept_id": 26,    "up_dept_id": 42,    "dept_pos": 91,    "dept_nm": "swtDnvCioXBxTXIPuDSUByHkd",    "pref_cd": "IppfowjlysqqVRsSDQuQhLvev",    "city_cd": "rxyOMyeGAiFgYUhpYVPdijxHj",    "disp_dsq": 19,    "disp_dsq_sub_no": 24,    "miss_teeth_flg": "nshldgHsGkSTPZubiRaXJAurH",    "upda_dte": "2180-08-06T22:34:51.308174822+09:00",    "upda_user_id": 30,    "crea_dte": "2085-01-21T18:43:51.499301057+09:00",    "crea_user_id": 42}



*/

// ANKEN_HIS_DEPT_MST_ struct is a row record of the ANKEN_HIS_DEPT_MST table in the anpidb database
type ANKEN_HIS_DEPT_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] DEPT_ID                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID int32 `gorm:"primary_key;column:DEPT_ID;type:int;"` // 部署ID
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] TOP_DEPT_ID                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TOPDEPTID int32 `gorm:"column:TOP_DEPT_ID;type:int;"` // 先頭部署ID
	//[ 4] UP_DEPT_ID                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDEPTID sql.NullInt64 `gorm:"column:UP_DEPT_ID;type:int;"` // 上部署ID
	//[ 5] DEPT_POS                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTPOS int32 `gorm:"column:DEPT_POS;type:int;"` // 部署ポジション	 1～10
	//[ 6] DEPT_NM                                        varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM string `gorm:"column:DEPT_NM;type:varchar;size:200;"` // 部署名
	//[ 7] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 8] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[ 9] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[10] DISP_DSQ_SUB_NO                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQSUBNO int32 `gorm:"column:DISP_DSQ_SUB_NO;type:int;"` // ポジション別表示順	 同階層内での表示順
	//[11] MISS_TEETH_FLG                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MISSTEETHFLG sql.NullString `gorm:"column:MISS_TEETH_FLG;type:char;size:1;"`
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_DEPT_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_DEPT_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "DEPT_ID",
			Comment:            `部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CLIENT_ID",
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
			GoFieldName:        "CLIENTID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "TOP_DEPT_ID",
			Comment:            `先頭部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPDEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "top_dept_id",
			ProtobufFieldName:  "top_dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "UP_DEPT_ID",
			Comment:            `上部署ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDEPTID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "up_dept_id",
			ProtobufFieldName:  "up_dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "DEPT_POS",
			Comment: `部署ポジション	 1～10`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTPOS",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_pos",
			ProtobufFieldName:  "dept_pos",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DEPT_NM",
			Comment:            `部署名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM",
			GoFieldType:        "string",
			JSONFieldName:      "dept_nm",
			ProtobufFieldName:  "dept_nm",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index: 10,
			Name:  "DISP_DSQ_SUB_NO",
			Comment: `ポジション別表示順	 同階層内での表示順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQSUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "disp_dsq_sub_no",
			ProtobufFieldName:  "disp_dsq_sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "MISS_TEETH_FLG",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MISSTEETHFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "miss_teeth_flg",
			ProtobufFieldName:  "miss_teeth_flg",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_HIS_DEPT_MST_) TableName() string {
	return "ANKEN_HIS_DEPT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_DEPT_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_DEPT_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_DEPT_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_DEPT_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_DEPT_MSTTableInfo
}
