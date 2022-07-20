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


CREATE TABLE `ANKEN_TARGET_DEPT` (
  `ANKEN_TARGET_ID` int(11) NOT NULL COMMENT '送信先対象選択ID',
  `DEPT_SEQ` int(11) NOT NULL COMMENT '部署連番',
  `DEPT_COND_KBN` char(1) NOT NULL COMMENT '部署検索区分（1:のみ 2:すべて）',
  `DEPT_POS` int(11) NOT NULL COMMENT '部署階層',
  `DEPT_NM` varchar(2000) NOT NULL COMMENT '部署名',
  `DEPT_ID` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID',
  `DEPT_ID1` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID1',
  `DEPT_ID2` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID2',
  `DEPT_ID3` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID3',
  `DEPT_ID4` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID4',
  `DEPT_ID5` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID5',
  `DEPT_ID6` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID6',
  `DEPT_ID7` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID7',
  `DEPT_ID8` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID8',
  `DEPT_ID9` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID9',
  `DEPT_ID10` int(11) NOT NULL DEFAULT '0' COMMENT '部署ID10',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`,`DEPT_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 30,    "dept_seq": 54,    "dept_cond_kbn": "ZkyMcbhiHZHtfZeasvJNprBxn",    "dept_pos": 38,    "dept_nm": "bupauRKKMurdnfQiQLkCWkhLA",    "dept_id": 11,    "dept_id_1": 90,    "dept_id_2": 67,    "dept_id_3": 88,    "dept_id_4": 38,    "dept_id_5": 50,    "dept_id_6": 98,    "dept_id_7": 56,    "dept_id_8": 14,    "dept_id_9": 98,    "dept_id_10": 53,    "upda_dte": "2138-03-15T09:58:18.32480506+09:00",    "upda_user_id": 61,    "crea_dte": "2301-06-12T07:26:28.133356027+09:00",    "crea_user_id": 95}



*/

// ANKEN_TARGET_DEPT_ struct is a row record of the ANKEN_TARGET_DEPT table in the anpidb database
type ANKEN_TARGET_DEPT_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] DEPT_SEQ                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	DEPTSEQ int32 `gorm:"primary_key;column:DEPT_SEQ;type:int;"` // 部署連番
	//[ 2] DEPT_COND_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DEPTCONDKBN string `gorm:"column:DEPT_COND_KBN;type:char;size:1;"` // 部署検索区分（1:のみ 2:すべて）
	//[ 3] DEPT_POS                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTPOS int32 `gorm:"column:DEPT_POS;type:int;"` // 部署階層
	//[ 4] DEPT_NM                                        varchar(2000)        null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2000    default: []
	DEPTNM string `gorm:"column:DEPT_NM;type:varchar;size:2000;"` // 部署名
	//[ 5] DEPT_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID int32 `gorm:"column:DEPT_ID;type:int;default:0;"` // 部署ID
	//[ 6] DEPT_ID1                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID1 int32 `gorm:"column:DEPT_ID1;type:int;default:0;"` // 部署ID1
	//[ 7] DEPT_ID2                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID2 int32 `gorm:"column:DEPT_ID2;type:int;default:0;"` // 部署ID2
	//[ 8] DEPT_ID3                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID3 int32 `gorm:"column:DEPT_ID3;type:int;default:0;"` // 部署ID3
	//[ 9] DEPT_ID4                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID4 int32 `gorm:"column:DEPT_ID4;type:int;default:0;"` // 部署ID4
	//[10] DEPT_ID5                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID5 int32 `gorm:"column:DEPT_ID5;type:int;default:0;"` // 部署ID5
	//[11] DEPT_ID6                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID6 int32 `gorm:"column:DEPT_ID6;type:int;default:0;"` // 部署ID6
	//[12] DEPT_ID7                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID7 int32 `gorm:"column:DEPT_ID7;type:int;default:0;"` // 部署ID7
	//[13] DEPT_ID8                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID8 int32 `gorm:"column:DEPT_ID8;type:int;default:0;"` // 部署ID8
	//[14] DEPT_ID9                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID9 int32 `gorm:"column:DEPT_ID9;type:int;default:0;"` // 部署ID9
	//[15] DEPT_ID10                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID10 int32 `gorm:"column:DEPT_ID10;type:int;default:0;"` // 部署ID10
	//[16] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[17] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[18] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[19] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_TARGET_DEPTTableInfo = &TableInfo{
	Name: "ANKEN_TARGET_DEPT",
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
			Name:               "DEPT_SEQ",
			Comment:            `部署連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_seq",
			ProtobufFieldName:  "dept_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DEPT_COND_KBN",
			Comment:            `部署検索区分（1:のみ 2:すべて）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEPTCONDKBN",
			GoFieldType:        "string",
			JSONFieldName:      "dept_cond_kbn",
			ProtobufFieldName:  "dept_cond_kbn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DEPT_POS",
			Comment:            `部署階層`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "DEPT_NM",
			Comment:            `部署名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2000,
			GoFieldName:        "DEPTNM",
			GoFieldType:        "string",
			JSONFieldName:      "dept_nm",
			ProtobufFieldName:  "dept_nm",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DEPT_ID",
			Comment:            `部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DEPT_ID1",
			Comment:            `部署ID1`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID1",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_1",
			ProtobufFieldName:  "dept_id_1",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "DEPT_ID2",
			Comment:            `部署ID2`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID2",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_2",
			ProtobufFieldName:  "dept_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "DEPT_ID3",
			Comment:            `部署ID3`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID3",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_3",
			ProtobufFieldName:  "dept_id_3",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "DEPT_ID4",
			Comment:            `部署ID4`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID4",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_4",
			ProtobufFieldName:  "dept_id_4",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "DEPT_ID5",
			Comment:            `部署ID5`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID5",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_5",
			ProtobufFieldName:  "dept_id_5",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "DEPT_ID6",
			Comment:            `部署ID6`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID6",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_6",
			ProtobufFieldName:  "dept_id_6",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "DEPT_ID7",
			Comment:            `部署ID7`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID7",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_7",
			ProtobufFieldName:  "dept_id_7",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "DEPT_ID8",
			Comment:            `部署ID8`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID8",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_8",
			ProtobufFieldName:  "dept_id_8",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "DEPT_ID9",
			Comment:            `部署ID9`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID9",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_9",
			ProtobufFieldName:  "dept_id_9",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "DEPT_ID10",
			Comment:            `部署ID10`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID10",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_10",
			ProtobufFieldName:  "dept_id_10",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_TARGET_DEPT_) TableName() string {
	return "ANKEN_TARGET_DEPT"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_DEPT_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_DEPT_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_DEPT_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_DEPT_) TableInfo() *TableInfo {
	return ANKEN_TARGET_DEPTTableInfo
}
