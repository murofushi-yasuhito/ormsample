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


CREATE TABLE `ANKEN_HIS_GRP_DTL_MST` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件IDを',
  `GRP_ID` int(11) NOT NULL COMMENT 'グループID',
  `GRP_SEND_SEQ` int(11) NOT NULL COMMENT 'グループ送信先連番',
  `POST_ID` int(11) DEFAULT NULL COMMENT '役職ID',
  `POST_COND_KBN` char(1) DEFAULT '0' COMMENT '役職条件区分	 0:指定なし 1:以上 2:以下',
  `POST_LEVEL` int(11) DEFAULT NULL COMMENT '職位順',
  `HUB_ID` int(11) DEFAULT NULL COMMENT '拠点ID',
  `DEPT_ID` int(11) DEFAULT NULL COMMENT '末端の部署ID',
  `DEPT_ID1` int(11) DEFAULT NULL COMMENT '部署ID1',
  `DEPT_ID2` int(11) DEFAULT NULL COMMENT '部署ID2',
  `DEPT_ID3` int(11) DEFAULT NULL COMMENT '部署ID3',
  `DEPT_ID4` int(11) DEFAULT NULL COMMENT '部署ID4',
  `DEPT_ID5` int(11) DEFAULT NULL COMMENT '部署ID5',
  `DEPT_ID6` int(11) DEFAULT NULL COMMENT '部署ID6',
  `DEPT_ID7` int(11) DEFAULT NULL COMMENT '部署ID7',
  `DEPT_ID8` int(11) DEFAULT NULL COMMENT '部署ID8',
  `DEPT_ID9` int(11) DEFAULT NULL COMMENT '部署ID9',
  `DEPT_ID10` int(11) DEFAULT NULL COMMENT '部署ID10',
  `DEPT_COND_KBN` char(1) DEFAULT '0' COMMENT '部署条件区分	 0:なし 1:のみ 2:すべて',
  `HOME_PREF_CD` char(2) DEFAULT NULL COMMENT '居住地都道府県',
  `HOME_CITY_CD` char(5) DEFAULT NULL COMMENT '居住地市区町村',
  `AREA_PREF_CD` char(2) DEFAULT NULL COMMENT '地域都道府県',
  `AREA_CITY_CD` char(5) DEFAULT NULL COMMENT '地域市区町村',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`GRP_ID`,`GRP_SEND_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "anken_id": 92,    "grp_id": 35,    "grp_send_seq": 7,    "post_id": 16,    "post_cond_kbn": "AMIQSnbPibDbSCsMuRiIUsQWv",    "post_level": 51,    "hub_id": 30,    "dept_id": 86,    "dept_id_1": 64,    "dept_id_2": 73,    "dept_id_3": 21,    "dept_id_4": 28,    "dept_id_5": 17,    "dept_id_6": 90,    "dept_id_7": 2,    "dept_id_8": 64,    "dept_id_9": 86,    "dept_id_10": 65,    "dept_cond_kbn": "EtImnaWqEyPBxttWiJrbOcbbv",    "home_pref_cd": "LnqtuanJilqHDLbmAlyLajkIj",    "home_city_cd": "xWpTDhLJZtmquMSMHjGxynJis",    "area_pref_cd": "BbAqZDkUCkgddfhFcOYaxROtg",    "area_city_cd": "pEWBiLqZTDtBwcMlJlJuOITPC",    "upda_dte": "2054-11-21T22:34:42.911700931+09:00",    "upda_user_id": 78,    "crea_dte": "2088-06-21T17:01:54.12430829+09:00",    "crea_user_id": 24}



*/

// ANKEN_HIS_GRP_DTL_MST_ struct is a row record of the ANKEN_HIS_GRP_DTL_MST table in the anpidb database
type ANKEN_HIS_GRP_DTL_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件IDを
	//[ 1] GRP_ID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	GRPID int32 `gorm:"primary_key;column:GRP_ID;type:int;"` // グループID
	//[ 2] GRP_SEND_SEQ                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	GRPSENDSEQ int32 `gorm:"primary_key;column:GRP_SEND_SEQ;type:int;"` // グループ送信先連番
	//[ 3] POST_ID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	POSTID sql.NullInt64 `gorm:"column:POST_ID;type:int;"` // 役職ID
	//[ 4] POST_COND_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	POSTCONDKBN sql.NullString `gorm:"column:POST_COND_KBN;type:char;size:1;default:0;"` // 役職条件区分	 0:指定なし 1:以上 2:以下
	//[ 5] POST_LEVEL                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	POSTLEVEL sql.NullInt64 `gorm:"column:POST_LEVEL;type:int;"` // 職位順
	//[ 6] HUB_ID                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	HUBID sql.NullInt64 `gorm:"column:HUB_ID;type:int;"` // 拠点ID
	//[ 7] DEPT_ID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID sql.NullInt64 `gorm:"column:DEPT_ID;type:int;"` // 末端の部署ID
	//[ 8] DEPT_ID1                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID1 sql.NullInt64 `gorm:"column:DEPT_ID1;type:int;"` // 部署ID1
	//[ 9] DEPT_ID2                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID2 sql.NullInt64 `gorm:"column:DEPT_ID2;type:int;"` // 部署ID2
	//[10] DEPT_ID3                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID3 sql.NullInt64 `gorm:"column:DEPT_ID3;type:int;"` // 部署ID3
	//[11] DEPT_ID4                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID4 sql.NullInt64 `gorm:"column:DEPT_ID4;type:int;"` // 部署ID4
	//[12] DEPT_ID5                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID5 sql.NullInt64 `gorm:"column:DEPT_ID5;type:int;"` // 部署ID5
	//[13] DEPT_ID6                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID6 sql.NullInt64 `gorm:"column:DEPT_ID6;type:int;"` // 部署ID6
	//[14] DEPT_ID7                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID7 sql.NullInt64 `gorm:"column:DEPT_ID7;type:int;"` // 部署ID7
	//[15] DEPT_ID8                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID8 sql.NullInt64 `gorm:"column:DEPT_ID8;type:int;"` // 部署ID8
	//[16] DEPT_ID9                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID9 sql.NullInt64 `gorm:"column:DEPT_ID9;type:int;"` // 部署ID9
	//[17] DEPT_ID10                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID10 sql.NullInt64 `gorm:"column:DEPT_ID10;type:int;"` // 部署ID10
	//[18] DEPT_COND_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DEPTCONDKBN sql.NullString `gorm:"column:DEPT_COND_KBN;type:char;size:1;default:0;"` // 部署条件区分	 0:なし 1:のみ 2:すべて
	//[19] HOME_PREF_CD                                   char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	HOMEPREFCD sql.NullString `gorm:"column:HOME_PREF_CD;type:char;size:2;"` // 居住地都道府県
	//[20] HOME_CITY_CD                                   char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	HOMECITYCD sql.NullString `gorm:"column:HOME_CITY_CD;type:char;size:5;"` // 居住地市区町村
	//[21] AREA_PREF_CD                                   char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	AREAPREFCD sql.NullString `gorm:"column:AREA_PREF_CD;type:char;size:2;"` // 地域都道府県
	//[22] AREA_CITY_CD                                   char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	AREACITYCD sql.NullString `gorm:"column:AREA_CITY_CD;type:char;size:5;"` // 地域市区町村
	//[23] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[24] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[25] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[26] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_GRP_DTL_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_GRP_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_ID",
			Comment:            `案件IDを`,
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
			Name:               "GRP_ID",
			Comment:            `グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_id",
			ProtobufFieldName:  "grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "GRP_SEND_SEQ",
			Comment:            `グループ送信先連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPSENDSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_send_seq",
			ProtobufFieldName:  "grp_send_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "POST_ID",
			Comment:            `役職ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "POSTID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "post_id",
			ProtobufFieldName:  "post_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "POST_COND_KBN",
			Comment: `役職条件区分	 0:指定なし 1:以上 2:以下`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "POSTCONDKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "post_cond_kbn",
			ProtobufFieldName:  "post_cond_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "POST_LEVEL",
			Comment:            `職位順`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "POSTLEVEL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "post_level",
			ProtobufFieldName:  "post_level",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "HUB_ID",
			Comment:            `拠点ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "HUBID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "hub_id",
			ProtobufFieldName:  "hub_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "DEPT_ID",
			Comment:            `末端の部署ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "DEPT_ID1",
			Comment:            `部署ID1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID1",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_1",
			ProtobufFieldName:  "dept_id_1",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "DEPT_ID2",
			Comment:            `部署ID2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID2",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_2",
			ProtobufFieldName:  "dept_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "DEPT_ID3",
			Comment:            `部署ID3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID3",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_3",
			ProtobufFieldName:  "dept_id_3",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "DEPT_ID4",
			Comment:            `部署ID4`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID4",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_4",
			ProtobufFieldName:  "dept_id_4",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "DEPT_ID5",
			Comment:            `部署ID5`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID5",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_5",
			ProtobufFieldName:  "dept_id_5",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "DEPT_ID6",
			Comment:            `部署ID6`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID6",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_6",
			ProtobufFieldName:  "dept_id_6",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "DEPT_ID7",
			Comment:            `部署ID7`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID7",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_7",
			ProtobufFieldName:  "dept_id_7",
			ProtobufType:       "int32",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "DEPT_ID8",
			Comment:            `部署ID8`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID8",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_8",
			ProtobufFieldName:  "dept_id_8",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "DEPT_ID9",
			Comment:            `部署ID9`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID9",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_9",
			ProtobufFieldName:  "dept_id_9",
			ProtobufType:       "int32",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "DEPT_ID10",
			Comment:            `部署ID10`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID10",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_10",
			ProtobufFieldName:  "dept_id_10",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index: 18,
			Name:  "DEPT_COND_KBN",
			Comment: `部署条件区分	 0:なし 1:のみ 2:すべて`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEPTCONDKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_cond_kbn",
			ProtobufFieldName:  "dept_cond_kbn",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "HOME_PREF_CD",
			Comment:            `居住地都道府県`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "HOMEPREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "home_pref_cd",
			ProtobufFieldName:  "home_pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "HOME_CITY_CD",
			Comment:            `居住地市区町村`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "HOMECITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "home_city_cd",
			ProtobufFieldName:  "home_city_cd",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "AREA_PREF_CD",
			Comment:            `地域都道府県`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "AREAPREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_pref_cd",
			ProtobufFieldName:  "area_pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "AREA_CITY_CD",
			Comment:            `地域市区町村`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "AREACITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_city_cd",
			ProtobufFieldName:  "area_city_cd",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
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
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
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
			ProtobufPos:        27,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_HIS_GRP_DTL_MST_) TableName() string {
	return "ANKEN_HIS_GRP_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_GRP_DTL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_GRP_DTL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_GRP_DTL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_GRP_DTL_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_GRP_DTL_MSTTableInfo
}
