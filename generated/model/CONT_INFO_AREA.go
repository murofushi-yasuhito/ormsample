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


CREATE TABLE `CONT_INFO_AREA` (
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `SUB_NO` int(11) NOT NULL COMMENT '枝番',
  `ELMNT_LVL` int(11) NOT NULL DEFAULT '99' COMMENT '影響レベル',
  `LIFT_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '解除フラグ',
  `AREA_KBN` char(4) NOT NULL COMMENT 'エリア区分',
  `AREA_CD1` varchar(2) NOT NULL DEFAULT '0' COMMENT 'エリアコード1',
  `AREA_CD2` varchar(10) DEFAULT '0' COMMENT 'エリアコード2',
  `AREA_CD3` varchar(10) DEFAULT '0' COMMENT 'エリアコード3',
  `AREA_CD4` varchar(10) DEFAULT '0' COMMENT 'エリアコード4',
  `AREA_CD5` varchar(10) DEFAULT '0' COMMENT 'エリアコード5',
  `WETH_AREA_CD1` varchar(2) DEFAULT '0' COMMENT '気象予報区コード1',
  `WETH_AREA_CD2` varchar(7) DEFAULT '0' COMMENT '気象予報区コード2',
  `WETH_AREA_CD3` varchar(7) DEFAULT '0' COMMENT '気象予報区コード3',
  `BEF_ELMNT_LVL` int(11) DEFAULT '0' COMMENT '前回影響レベル',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SID`,`SUB_NO`),
  KEY `CONT_INFO_AREA_IDX_1` (`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`WETH_AREA_CD1`,`WETH_AREA_CD2`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツ情報地域'

JSON Sample
-------------------------------------
{    "sid": 8,    "sub_no": 38,    "elmnt_lvl": 13,    "lift_flg": "eCdjgDsjYtVRlLdGkkbuuejEk",    "area_kbn": "XkIbZDmEvRGMHLDtSFjZQAbVO",    "area_cd_1": "chnxLXiFyNPLpEdyBKMvCbmqm",    "area_cd_2": "jGUtOnBuGMktSHGQEsbXePYAK",    "area_cd_3": "HICLNrJHuHmSSZJVCHQroquPm",    "area_cd_4": "qerDETJLqdZtMHvbMyBIBiuPl",    "area_cd_5": "iYXrDPUXebvRtFisOkryrxfib",    "weth_area_cd_1": "bbZPmeiyOZXQpBrZYcESUIIiq",    "weth_area_cd_2": "MwrhGbUasHVAPMoHLrlKECcuw",    "weth_area_cd_3": "VxnBHjLcMunBjaQmpQNgNGBMa",    "bef_elmnt_lvl": 93,    "crea_dte": "2269-02-10T22:35:56.148732452+09:00",    "crea_user_id": 51}



*/

// CONT_INFO_AREA_ struct is a row record of the CONT_INFO_AREA table in the anpidb database
type CONT_INFO_AREA_ struct {
	//[ 0] SID                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"primary_key;column:SID;type:int;"` // 情報番号
	//[ 1] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;"` // 枝番
	//[ 2] ELMNT_LVL                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [99]
	ELMNTLVL int32 `gorm:"column:ELMNT_LVL;type:int;default:99;"` // 影響レベル
	//[ 3] LIFT_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	LIFTFLG string `gorm:"column:LIFT_FLG;type:char;size:1;default:0;"` // 解除フラグ
	//[ 4] AREA_KBN                                       char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:char;size:4;"` // エリア区分
	//[ 5] AREA_CD1                                       varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	AREACD1 string `gorm:"column:AREA_CD1;type:varchar;size:2;default:0;"` // エリアコード1
	//[ 6] AREA_CD2                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD2 sql.NullString `gorm:"column:AREA_CD2;type:varchar;size:10;default:0;"` // エリアコード2
	//[ 7] AREA_CD3                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD3 sql.NullString `gorm:"column:AREA_CD3;type:varchar;size:10;default:0;"` // エリアコード3
	//[ 8] AREA_CD4                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD4 sql.NullString `gorm:"column:AREA_CD4;type:varchar;size:10;default:0;"` // エリアコード4
	//[ 9] AREA_CD5                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD5 sql.NullString `gorm:"column:AREA_CD5;type:varchar;size:10;default:0;"` // エリアコード5
	//[10] WETH_AREA_CD1                                  varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	WETHAREACD1 sql.NullString `gorm:"column:WETH_AREA_CD1;type:varchar;size:2;default:0;"` // 気象予報区コード1
	//[11] WETH_AREA_CD2                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD2 sql.NullString `gorm:"column:WETH_AREA_CD2;type:varchar;size:7;default:0;"` // 気象予報区コード2
	//[12] WETH_AREA_CD3                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD3 sql.NullString `gorm:"column:WETH_AREA_CD3;type:varchar;size:7;default:0;"` // 気象予報区コード3
	//[13] BEF_ELMNT_LVL                                  int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	BEFELMNTLVL sql.NullInt64 `gorm:"column:BEF_ELMNT_LVL;type:int;default:0;"` // 前回影響レベル
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_INFO_AREATableInfo = &TableInfo{
	Name: "CONT_INFO_AREA",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "SID",
			Comment:            `情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SID",
			GoFieldType:        "int32",
			JSONFieldName:      "sid",
			ProtobufFieldName:  "sid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "SUB_NO",
			Comment:            `枝番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "sub_no",
			ProtobufFieldName:  "sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ELMNT_LVL",
			Comment:            `影響レベル`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ELMNTLVL",
			GoFieldType:        "int32",
			JSONFieldName:      "elmnt_lvl",
			ProtobufFieldName:  "elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "LIFT_FLG",
			Comment:            `解除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LIFTFLG",
			GoFieldType:        "string",
			JSONFieldName:      "lift_flg",
			ProtobufFieldName:  "lift_flg",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "AREA_CD1",
			Comment:            `エリアコード1`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "AREACD1",
			GoFieldType:        "string",
			JSONFieldName:      "area_cd_1",
			ProtobufFieldName:  "area_cd_1",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "AREA_CD2",
			Comment:            `エリアコード2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_2",
			ProtobufFieldName:  "area_cd_2",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "AREA_CD3",
			Comment:            `エリアコード3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_3",
			ProtobufFieldName:  "area_cd_3",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "AREA_CD4",
			Comment:            `エリアコード4`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD4",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_4",
			ProtobufFieldName:  "area_cd_4",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "AREA_CD5",
			Comment:            `エリアコード5`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD5",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_5",
			ProtobufFieldName:  "area_cd_5",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "WETH_AREA_CD1",
			Comment:            `気象予報区コード1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "WETHAREACD1",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_1",
			ProtobufFieldName:  "weth_area_cd_1",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "WETH_AREA_CD2",
			Comment:            `気象予報区コード2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       7,
			GoFieldName:        "WETHAREACD2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_2",
			ProtobufFieldName:  "weth_area_cd_2",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "WETH_AREA_CD3",
			Comment:            `気象予報区コード3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       7,
			GoFieldName:        "WETHAREACD3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_3",
			ProtobufFieldName:  "weth_area_cd_3",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "BEF_ELMNT_LVL",
			Comment:            `前回影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BEFELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bef_elmnt_lvl",
			ProtobufFieldName:  "bef_elmnt_lvl",
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
func (c *CONT_INFO_AREA_) TableName() string {
	return "CONT_INFO_AREA"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_INFO_AREA_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_INFO_AREA_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_INFO_AREA_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_INFO_AREA_) TableInfo() *TableInfo {
	return CONT_INFO_AREATableInfo
}
