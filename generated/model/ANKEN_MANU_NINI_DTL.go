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


CREATE TABLE `ANKEN_MANU_NINI_DTL` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `POST_ID` int(11) DEFAULT '0' COMMENT '役職ID	 0:すべて',
  `POST_LEVEL` int(11) DEFAULT '0' COMMENT '職位順',
  `POST_COND_KBN` char(1) DEFAULT '0' COMMENT '役職条件区分	 0:指定なし 1:以上 2:以下',
  `HUB_ID` int(11) DEFAULT '0' COMMENT '拠点ID	 0:すべて',
  `HUB_KBN` char(1) DEFAULT '0' COMMENT '拠点条件区分	 N:拠点未設定',
  `DEPT_ID1` int(11) DEFAULT '0' COMMENT '部署ID1	 0:すべて',
  `DEPT_ID2` int(11) DEFAULT '0' COMMENT '部署ID2',
  `DEPT_ID3` int(11) DEFAULT '0' COMMENT '部署ID3',
  `DEPT_ID4` int(11) DEFAULT '0' COMMENT '部署ID4',
  `DEPT_ID5` int(11) DEFAULT '0' COMMENT '部署ID5',
  `DEPT_ID6` int(11) DEFAULT '0' COMMENT '部署ID6',
  `DEPT_ID7` int(11) DEFAULT '0' COMMENT '部署ID7',
  `DEPT_ID8` int(11) DEFAULT '0' COMMENT '部署ID8',
  `DEPT_ID9` int(11) DEFAULT '0' COMMENT '部署ID9',
  `DEPT_ID10` int(11) DEFAULT '0' COMMENT '部署ID10',
  `DEPT_COND_KBN` char(1) DEFAULT '0' COMMENT '部署条件区分	 0:なし 1:のみ 2:すべて',
  `HOME_PREF_CD` char(2) DEFAULT '0' COMMENT '居住地都道府県	 0:すべて N:未設定',
  `HOME_CITY_CD` char(5) DEFAULT '0' COMMENT '居住地市区町村	 0:すべて',
  `AREA_PREF_CD` char(2) DEFAULT '0' COMMENT '地域都道府県	 0:すべて',
  `AREA_CITY_CD` char(5) DEFAULT '0' COMMENT '地域市区町村	 0:すべて',
  `USER_NM` varchar(100) DEFAULT NULL COMMENT 'ユーザ名',
  `USER_NO` varchar(650) DEFAULT NULL COMMENT 'ユーザ番号',
  `STS_CDS` varchar(20) DEFAULT NULL COMMENT '状態コード(複数)	 カンマ区切り',
  `STS_CMNT_FLG` char(1) DEFAULT NULL COMMENT '状態コメントフラグ	 0:コメント無 1:コメント有',
  `MAIL_STS` char(1) DEFAULT NULL COMMENT 'メールステータス(条件)	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG',
  `FAMI_STS_CDS` varchar(20) DEFAULT NULL COMMENT '家族状態コード(複数)	 カンマ区切り',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  `QES1_STS_CDS` varchar(20) DEFAULT NULL COMMENT '設問１状態コード(複数) カンマ区切り',
  `QES2_STS_CDS` varchar(20) DEFAULT NULL COMMENT '設問２状態コード(複数) カンマ区切り',
  `QES3_STS_CDS` varchar(20) DEFAULT NULL COMMENT '設問３状態コード(複数) カンマ区切り',
  `QES4_STS_CDS` varchar(20) DEFAULT NULL COMMENT '設問４状態コード(複数) カンマ区切り',
  `QES5_STS_CDS` varchar(20) DEFAULT NULL COMMENT '設問５状態コード(複数) カンマ区切り',
  `TARGET_USER_KBNS` varchar(20) DEFAULT NULL COMMENT '対象ユーザ 0:案件起動ユーザ 1:自主応答ユーザ',
  `SUMM_GRP_ID` int(11) DEFAULT '0' COMMENT '集計グループID',
  `APP_KBN` char(1) DEFAULT NULL COMMENT 'アプリ条件区分（0:未登録、1:登録済み）',
  PRIMARY KEY (`ANKEN_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_id": 5,    "post_id": 5,    "post_level": 54,    "post_cond_kbn": "aJPCpBMERwgRXutGWrvLhUyDB",    "hub_id": 4,    "hub_kbn": "dxJlwvoGrpQmMnjlqtKLeITTm",    "dept_id_1": 95,    "dept_id_2": 84,    "dept_id_3": 75,    "dept_id_4": 36,    "dept_id_5": 22,    "dept_id_6": 42,    "dept_id_7": 21,    "dept_id_8": 85,    "dept_id_9": 15,    "dept_id_10": 63,    "dept_cond_kbn": "xjJPLSGfppswkWAbyBwPygjdY",    "home_pref_cd": "bdhBOqVNYUPUFLOophDNFiteL",    "home_city_cd": "KGocViJwuNJLvOMUghLDTkFSk",    "area_pref_cd": "DKHaHcoFDMblUalJdaZJapPdc",    "area_city_cd": "rsVCwprbbjEnhRROKyQRoQJck",    "user_nm": "QAEvdxxbsiLFABjpKhTbdGSSg",    "user_no": "EKduhaxSoRhJNghalDTJqZkNT",    "sts_cds": "DYsskdXaTKwQDYLWoTjLnygFC",    "sts_cmnt_flg": "TGIGxIXrUHOvhPLhAVtbIRJFq",    "mail_sts": "EAdZZoyCJKssHygIHtGNgXGFY",    "fami_sts_cds": "juVirDBwfARxfvdcwFArdxiwi",    "crea_dte": "2184-04-02T10:19:58.611150455+09:00",    "crea_user_id": 93,    "qes_1_sts_cds": "TQJknQrNfbqSEMaGBXRBMteDB",    "qes_2_sts_cds": "BguJsfcxVuvQaYiqJmmnWMPtI",    "qes_3_sts_cds": "EUGRwrUDeFNNJWgVpXhPQLwLX",    "qes_4_sts_cds": "foPhnAyCrmcxdIshYlXCiQJHx",    "qes_5_sts_cds": "XDxaUpQcTFjPAPMuhwcFJrJBB",    "target_user_kbns": "BMcPOeBtncxxnRJcdxhwUwMwU",    "summ_grp_id": 0,    "app_kbn": "FNApKypLOTUBrwiepgGQLlHwK"}



*/

// ANKEN_MANU_NINI_DTL_ struct is a row record of the ANKEN_MANU_NINI_DTL table in the anpidb database
type ANKEN_MANU_NINI_DTL_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] POST_ID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	POSTID sql.NullInt64 `gorm:"column:POST_ID;type:int;default:0;"` // 役職ID	 0:すべて
	//[ 2] POST_LEVEL                                     int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	POSTLEVEL sql.NullInt64 `gorm:"column:POST_LEVEL;type:int;default:0;"` // 職位順
	//[ 3] POST_COND_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	POSTCONDKBN sql.NullString `gorm:"column:POST_COND_KBN;type:char;size:1;default:0;"` // 役職条件区分	 0:指定なし 1:以上 2:以下
	//[ 4] HUB_ID                                         int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	HUBID sql.NullInt64 `gorm:"column:HUB_ID;type:int;default:0;"` // 拠点ID	 0:すべて
	//[ 5] HUB_KBN                                        char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	HUBKBN sql.NullString `gorm:"column:HUB_KBN;type:char;size:1;default:0;"` // 拠点条件区分	 N:拠点未設定
	//[ 6] DEPT_ID1                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID1 sql.NullInt64 `gorm:"column:DEPT_ID1;type:int;default:0;"` // 部署ID1	 0:すべて
	//[ 7] DEPT_ID2                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID2 sql.NullInt64 `gorm:"column:DEPT_ID2;type:int;default:0;"` // 部署ID2
	//[ 8] DEPT_ID3                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID3 sql.NullInt64 `gorm:"column:DEPT_ID3;type:int;default:0;"` // 部署ID3
	//[ 9] DEPT_ID4                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID4 sql.NullInt64 `gorm:"column:DEPT_ID4;type:int;default:0;"` // 部署ID4
	//[10] DEPT_ID5                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID5 sql.NullInt64 `gorm:"column:DEPT_ID5;type:int;default:0;"` // 部署ID5
	//[11] DEPT_ID6                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID6 sql.NullInt64 `gorm:"column:DEPT_ID6;type:int;default:0;"` // 部署ID6
	//[12] DEPT_ID7                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID7 sql.NullInt64 `gorm:"column:DEPT_ID7;type:int;default:0;"` // 部署ID7
	//[13] DEPT_ID8                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID8 sql.NullInt64 `gorm:"column:DEPT_ID8;type:int;default:0;"` // 部署ID8
	//[14] DEPT_ID9                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID9 sql.NullInt64 `gorm:"column:DEPT_ID9;type:int;default:0;"` // 部署ID9
	//[15] DEPT_ID10                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID10 sql.NullInt64 `gorm:"column:DEPT_ID10;type:int;default:0;"` // 部署ID10
	//[16] DEPT_COND_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DEPTCONDKBN sql.NullString `gorm:"column:DEPT_COND_KBN;type:char;size:1;default:0;"` // 部署条件区分	 0:なし 1:のみ 2:すべて
	//[17] HOME_PREF_CD                                   char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: [0]
	HOMEPREFCD sql.NullString `gorm:"column:HOME_PREF_CD;type:char;size:2;default:0;"` // 居住地都道府県	 0:すべて N:未設定
	//[18] HOME_CITY_CD                                   char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: [0]
	HOMECITYCD sql.NullString `gorm:"column:HOME_CITY_CD;type:char;size:5;default:0;"` // 居住地市区町村	 0:すべて
	//[19] AREA_PREF_CD                                   char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: [0]
	AREAPREFCD sql.NullString `gorm:"column:AREA_PREF_CD;type:char;size:2;default:0;"` // 地域都道府県	 0:すべて
	//[20] AREA_CITY_CD                                   char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: [0]
	AREACITYCD sql.NullString `gorm:"column:AREA_CITY_CD;type:char;size:5;default:0;"` // 地域市区町村	 0:すべて
	//[21] USER_NM                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	USERNM sql.NullString `gorm:"column:USER_NM;type:varchar;size:100;"` // ユーザ名
	//[22] USER_NO                                        varchar(650)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 650     default: []
	USERNO sql.NullString `gorm:"column:USER_NO;type:varchar;size:650;"` // ユーザ番号
	//[23] STS_CDS                                        varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	STSCDS sql.NullString `gorm:"column:STS_CDS;type:varchar;size:20;"` // 状態コード(複数)	 カンマ区切り
	//[24] STS_CMNT_FLG                                   char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STSCMNTFLG sql.NullString `gorm:"column:STS_CMNT_FLG;type:char;size:1;"` // 状態コメントフラグ	 0:コメント無 1:コメント有
	//[25] MAIL_STS                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSTS sql.NullString `gorm:"column:MAIL_STS;type:char;size:1;"` // メールステータス(条件)	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG
	//[26] FAMI_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	FAMISTSCDS sql.NullString `gorm:"column:FAMI_STS_CDS;type:varchar;size:20;"` // 家族状態コード(複数)	 カンマ区切り
	//[27] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[28] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID
	//[29] QES1_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	QES1STSCDS sql.NullString `gorm:"column:QES1_STS_CDS;type:varchar;size:20;"` // 設問１状態コード(複数) カンマ区切り
	//[30] QES2_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	QES2STSCDS sql.NullString `gorm:"column:QES2_STS_CDS;type:varchar;size:20;"` // 設問２状態コード(複数) カンマ区切り
	//[31] QES3_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	QES3STSCDS sql.NullString `gorm:"column:QES3_STS_CDS;type:varchar;size:20;"` // 設問３状態コード(複数) カンマ区切り
	//[32] QES4_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	QES4STSCDS sql.NullString `gorm:"column:QES4_STS_CDS;type:varchar;size:20;"` // 設問４状態コード(複数) カンマ区切り
	//[33] QES5_STS_CDS                                   varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	QES5STSCDS sql.NullString `gorm:"column:QES5_STS_CDS;type:varchar;size:20;"` // 設問５状態コード(複数) カンマ区切り
	//[34] TARGET_USER_KBNS                               varchar(20)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	TARGETUSERKBNS sql.NullString `gorm:"column:TARGET_USER_KBNS;type:varchar;size:20;"` // 対象ユーザ 0:案件起動ユーザ 1:自主応答ユーザ
	//[35] SUMM_GRP_ID                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SUMMGRPID sql.NullInt64 `gorm:"column:SUMM_GRP_ID;type:int;default:0;"` // 集計グループID
	//[36] APP_KBN                                        char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	APPKBN sql.NullString `gorm:"column:APP_KBN;type:char;size:1;"` // アプリ条件区分（0:未登録、1:登録済み）

}

var ANKEN_MANU_NINI_DTLTableInfo = &TableInfo{
	Name: "ANKEN_MANU_NINI_DTL",
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
			Index: 1,
			Name:  "POST_ID",
			Comment: `役職ID	 0:すべて`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "HUB_ID",
			Comment: `拠点ID	 0:すべて`,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "HUB_KBN",
			Comment: `拠点条件区分	 N:拠点未設定`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "HUBKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "hub_kbn",
			ProtobufFieldName:  "hub_kbn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index: 6,
			Name:  "DEPT_ID1",
			Comment: `部署ID1	 0:すべて`,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index: 16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index: 17,
			Name:  "HOME_PREF_CD",
			Comment: `居住地都道府県	 0:すべて N:未設定`,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index: 18,
			Name:  "HOME_CITY_CD",
			Comment: `居住地市区町村	 0:すべて`,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index: 19,
			Name:  "AREA_PREF_CD",
			Comment: `地域都道府県	 0:すべて`,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index: 20,
			Name:  "AREA_CITY_CD",
			Comment: `地域市区町村	 0:すべて`,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "USER_NM",
			Comment:            `ユーザ名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "USERNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_nm",
			ProtobufFieldName:  "user_nm",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "USER_NO",
			Comment:            `ユーザ番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(650)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       650,
			GoFieldName:        "USERNO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "user_no",
			ProtobufFieldName:  "user_no",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index: 23,
			Name:  "STS_CDS",
			Comment: `状態コード(複数)	 カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_cds",
			ProtobufFieldName:  "sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index: 24,
			Name:  "STS_CMNT_FLG",
			Comment: `状態コメントフラグ	 0:コメント無 1:コメント有`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STSCMNTFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_cmnt_flg",
			ProtobufFieldName:  "sts_cmnt_flg",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index: 25,
			Name:  "MAIL_STS",
			Comment: `メールステータス(条件)	 0:メールアドレス未登録 1:メールアドレス送信OK 2:メールアドレス送信NG`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSTS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "mail_sts",
			ProtobufFieldName:  "mail_sts",
			ProtobufType:       "string",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index: 26,
			Name:  "FAMI_STS_CDS",
			Comment: `家族状態コード(複数)	 カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "FAMISTSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "fami_sts_cds",
			ProtobufFieldName:  "fami_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
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
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
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
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "QES1_STS_CDS",
			Comment:            `設問１状態コード(複数) カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "QES1STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_1_sts_cds",
			ProtobufFieldName:  "qes_1_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "QES2_STS_CDS",
			Comment:            `設問２状態コード(複数) カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "QES2STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_2_sts_cds",
			ProtobufFieldName:  "qes_2_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "QES3_STS_CDS",
			Comment:            `設問３状態コード(複数) カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "QES3STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_3_sts_cds",
			ProtobufFieldName:  "qes_3_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "QES4_STS_CDS",
			Comment:            `設問４状態コード(複数) カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "QES4STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_4_sts_cds",
			ProtobufFieldName:  "qes_4_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "QES5_STS_CDS",
			Comment:            `設問５状態コード(複数) カンマ区切り`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "QES5STSCDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_5_sts_cds",
			ProtobufFieldName:  "qes_5_sts_cds",
			ProtobufType:       "string",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "TARGET_USER_KBNS",
			Comment:            `対象ユーザ 0:案件起動ユーザ 1:自主応答ユーザ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "TARGETUSERKBNS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "target_user_kbns",
			ProtobufFieldName:  "target_user_kbns",
			ProtobufType:       "string",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
			Name:               "SUMM_GRP_ID",
			Comment:            `集計グループID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUMMGRPID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "summ_grp_id",
			ProtobufFieldName:  "summ_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
			Name:               "APP_KBN",
			Comment:            `アプリ条件区分（0:未登録、1:登録済み）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "APPKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "app_kbn",
			ProtobufFieldName:  "app_kbn",
			ProtobufType:       "string",
			ProtobufPos:        37,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_MANU_NINI_DTL_) TableName() string {
	return "ANKEN_MANU_NINI_DTL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_MANU_NINI_DTL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_MANU_NINI_DTL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_MANU_NINI_DTL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_MANU_NINI_DTL_) TableInfo() *TableInfo {
	return ANKEN_MANU_NINI_DTLTableInfo
}
