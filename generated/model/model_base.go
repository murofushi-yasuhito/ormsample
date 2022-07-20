package model

import "fmt"

// Action CRUD actions
type Action int32

var (
	// Create action when record is created
	Create = Action(0)

	// RetrieveOne action when a record is retrieved from db
	RetrieveOne = Action(1)

	// RetrieveMany action when record(s) are retrieved from db
	RetrieveMany = Action(2)

	// Update action when record is updated in db
	Update = Action(3)

	// Delete action when record is deleted in db
	Delete = Action(4)

	// FetchDDL action when fetching ddl info from db
	FetchDDL = Action(5)

	tables map[string]*TableInfo
)

func init() {
	tables = make(map[string]*TableInfo)

	tables["ADDR_AREA_MST"] = ADDR_AREA_MSTTableInfo
	tables["ADDR_CITY_MST"] = ADDR_CITY_MSTTableInfo
	tables["ADDR_PREF_MST"] = ADDR_PREF_MSTTableInfo
	tables["ADMIN_GRP_DTL_MST"] = ADMIN_GRP_DTL_MSTTableInfo
	tables["ADMIN_GRP_MST"] = ADMIN_GRP_MSTTableInfo
	tables["ADMIN_MNG_MST"] = ADMIN_MNG_MSTTableInfo
	tables["ANKEN"] = ANKENTableInfo
	tables["ANKEN_DTL"] = ANKEN_DTLTableInfo
	tables["ANKEN_HIS_CLIENT_TMPL_MST"] = ANKEN_HIS_CLIENT_TMPL_MSTTableInfo
	tables["ANKEN_HIS_CLIENT_TMPL_OPT_MST"] = ANKEN_HIS_CLIENT_TMPL_OPT_MSTTableInfo
	tables["ANKEN_HIS_DEPT_MST"] = ANKEN_HIS_DEPT_MSTTableInfo
	tables["ANKEN_HIS_GRP_DTL_MST"] = ANKEN_HIS_GRP_DTL_MSTTableInfo
	tables["ANKEN_HIS_GRP_MST"] = ANKEN_HIS_GRP_MSTTableInfo
	tables["ANKEN_HIS_GRP_SEND_MST"] = ANKEN_HIS_GRP_SEND_MSTTableInfo
	tables["ANKEN_HIS_GRP_USER_MST"] = ANKEN_HIS_GRP_USER_MSTTableInfo
	tables["ANKEN_HIS_HUB_MST"] = ANKEN_HIS_HUB_MSTTableInfo
	tables["ANKEN_HIS_POST_MST"] = ANKEN_HIS_POST_MSTTableInfo
	tables["ANKEN_HIS_SUMM_GRP_DTL_MST"] = ANKEN_HIS_SUMM_GRP_DTL_MSTTableInfo
	tables["ANKEN_HIS_SUMM_GRP_MST"] = ANKEN_HIS_SUMM_GRP_MSTTableInfo
	tables["ANKEN_HIS_USER_DEPT_MST"] = ANKEN_HIS_USER_DEPT_MSTTableInfo
	tables["ANKEN_HIS_USER_HUB_MST"] = ANKEN_HIS_USER_HUB_MSTTableInfo
	tables["ANKEN_HIS_USER_MST"] = ANKEN_HIS_USER_MSTTableInfo
	tables["ANKEN_HIS_USER_POST_MST"] = ANKEN_HIS_USER_POST_MSTTableInfo
	tables["ANKEN_MANU_AREA_DTL"] = ANKEN_MANU_AREA_DTLTableInfo
	tables["ANKEN_MANU_GRP_DTL"] = ANKEN_MANU_GRP_DTLTableInfo
	tables["ANKEN_MANU_NINI_DTL"] = ANKEN_MANU_NINI_DTLTableInfo
	tables["ANKEN_MSG"] = ANKEN_MSGTableInfo
	tables["ANKEN_MSG_USER_MAIL"] = ANKEN_MSG_USER_MAILTableInfo
	tables["ANKEN_NO_MNG"] = ANKEN_NO_MNGTableInfo
	tables["ANKEN_SEND_USER"] = ANKEN_SEND_USERTableInfo
	tables["ANKEN_SEND_USER_DEPT"] = ANKEN_SEND_USER_DEPTTableInfo
	tables["ANKEN_SEND_USER_HUB"] = ANKEN_SEND_USER_HUBTableInfo
	tables["ANKEN_SEND_USER_POST"] = ANKEN_SEND_USER_POSTTableInfo
	tables["ANKEN_START_DEPT"] = ANKEN_START_DEPTTableInfo
	tables["ANKEN_START_MAIL_MST"] = ANKEN_START_MAIL_MSTTableInfo
	tables["ANKEN_START_MAIL_SEND"] = ANKEN_START_MAIL_SENDTableInfo
	tables["ANKEN_START_MAIL_SEND_BK"] = ANKEN_START_MAIL_SEND_BKTableInfo
	tables["ANKEN_STATUS_CUR"] = ANKEN_STATUS_CURTableInfo
	tables["ANKEN_STATUS_CUR_BK"] = ANKEN_STATUS_CUR_BKTableInfo
	tables["ANKEN_STATUS_REC"] = ANKEN_STATUS_RECTableInfo
	tables["ANKEN_STATUS_REC_copy"] = ANKEN_STATUS_REC_copyTableInfo
	tables["ANKEN_TARGET"] = ANKEN_TARGETTableInfo
	tables["ANKEN_TARGET_AREA"] = ANKEN_TARGET_AREATableInfo
	tables["ANKEN_TARGET_DEPT"] = ANKEN_TARGET_DEPTTableInfo
	tables["ANKEN_TARGET_HUB"] = ANKEN_TARGET_HUBTableInfo
	tables["ANKEN_TARGET_POST"] = ANKEN_TARGET_POSTTableInfo
	tables["ANKEN_TARGET_USER"] = ANKEN_TARGET_USERTableInfo
	tables["ANPI_AREA_SYNC"] = ANPI_AREA_SYNCTableInfo
	tables["ANPI_DLCT_CATE2_MST"] = ANPI_DLCT_CATE2_MSTTableInfo
	tables["ANPI_DLCT_MST"] = ANPI_DLCT_MSTTableInfo
	tables["ANPI_DLCT_UPDA_HIS"] = ANPI_DLCT_UPDA_HISTableInfo
	tables["ANPI_EXCEPT_AREA_MST"] = ANPI_EXCEPT_AREA_MSTTableInfo
	tables["ANPI_USER_AREA_MST"] = ANPI_USER_AREA_MSTTableInfo
	tables["ANPI_USER_AREA_MST_CHECK"] = ANPI_USER_AREA_MST_CHECKTableInfo
	tables["ANPI_USER_AREA_MST_RESULT"] = ANPI_USER_AREA_MST_RESULTTableInfo
	tables["AUTH_DTL_MST"] = AUTH_DTL_MSTTableInfo
	tables["AUTH_MST"] = AUTH_MSTTableInfo
	tables["BATCH_STATUS"] = BATCH_STATUSTableInfo
	tables["CATEGORY1_MST"] = CATEGORY1_MSTTableInfo
	tables["CATEGORY2_MST"] = CATEGORY2_MSTTableInfo
	tables["CATEGORY_CNV_MST"] = CATEGORY_CNV_MSTTableInfo
	tables["CATE_MST"] = CATE_MSTTableInfo
	tables["CLIENT_MNG_MST"] = CLIENT_MNG_MSTTableInfo
	tables["CLIENT_MST"] = CLIENT_MSTTableInfo
	tables["CLIENT_TMPL_ADD_QES_MST"] = CLIENT_TMPL_ADD_QES_MSTTableInfo
	tables["CLIENT_TMPL_MST"] = CLIENT_TMPL_MSTTableInfo
	tables["CLIENT_TMPL_OPT_MST"] = CLIENT_TMPL_OPT_MSTTableInfo
	tables["CONT_ANPI_HIS"] = CONT_ANPI_HISTableInfo
	tables["CONT_DELI_HIS"] = CONT_DELI_HISTableInfo
	tables["CONT_INFO_AREA"] = CONT_INFO_AREATableInfo
	tables["CONT_INFO_MAIN"] = CONT_INFO_MAINTableInfo
	tables["CONT_INFO_WARN"] = CONT_INFO_WARNTableInfo
	tables["CONT_INFO_WARN_CITY"] = CONT_INFO_WARN_CITYTableInfo
	tables["CONT_MAIL_SEND_HIS"] = CONT_MAIL_SEND_HISTableInfo
	tables["CONT_MAIL_SEND_HIS_DTL"] = CONT_MAIL_SEND_HIS_DTLTableInfo
	tables["CONT_MAIL_SEND_HIS_WARN"] = CONT_MAIL_SEND_HIS_WARNTableInfo
	tables["CONT_MAIL_SEND_MAIL_ADDR"] = CONT_MAIL_SEND_MAIL_ADDRTableInfo
	tables["DEFAULT_DLCT_CATE2_MST"] = DEFAULT_DLCT_CATE2_MSTTableInfo
	tables["DEFAULT_DLCT_DAY_MST"] = DEFAULT_DLCT_DAY_MSTTableInfo
	tables["DEFAULT_DLCT_MST"] = DEFAULT_DLCT_MSTTableInfo
	tables["DELI_SET_MENU_CATE_MST"] = DELI_SET_MENU_CATE_MSTTableInfo
	tables["DELI_SET_MENU_DTL_MST"] = DELI_SET_MENU_DTL_MSTTableInfo
	tables["DELI_SET_MENU_MST"] = DELI_SET_MENU_MSTTableInfo
	tables["DEL_CLIENT_TMPL_DTL_MST"] = DEL_CLIENT_TMPL_DTL_MSTTableInfo
	tables["DEL_CLIENT_TMPL_OPT2_MNG_MST"] = DEL_CLIENT_TMPL_OPT2_MNG_MSTTableInfo
	tables["DEL_CLIENT_TMPL_OPT2_MST"] = DEL_CLIENT_TMPL_OPT2_MSTTableInfo
	tables["DEPT_MST"] = DEPT_MSTTableInfo
	tables["FAMILY_BOARD"] = FAMILY_BOARDTableInfo
	tables["FAMILY_MST"] = FAMILY_MSTTableInfo
	tables["FAMILY_TMPL_MST"] = FAMILY_TMPL_MSTTableInfo
	tables["FAMILY_USER_MST"] = FAMILY_USER_MSTTableInfo
	tables["GRP_DTL_MST"] = GRP_DTL_MSTTableInfo
	tables["GRP_MST"] = GRP_MSTTableInfo
	tables["GRP_SEND_MST"] = GRP_SEND_MSTTableInfo
	tables["GRP_USER_MST"] = GRP_USER_MSTTableInfo
	tables["HOLIDAY_MST"] = HOLIDAY_MSTTableInfo
	tables["HUB_MST"] = HUB_MSTTableInfo
	tables["LEVEL_MST"] = LEVEL_MSTTableInfo
	tables["MAIL_STAT_MNG"] = MAIL_STAT_MNGTableInfo
	tables["MENU_DTL_MST"] = MENU_DTL_MSTTableInfo
	tables["MENU_MST"] = MENU_MSTTableInfo
	tables["MST_UPD_MNG"] = MST_UPD_MNGTableInfo
	tables["NOTICE"] = NOTICETableInfo
	tables["OEM_MST"] = OEM_MSTTableInfo
	tables["POST_MST"] = POST_MSTTableInfo
	tables["QES_TMPL_TYPE_MST"] = QES_TMPL_TYPE_MSTTableInfo
	tables["RAIL_AREA_MST"] = RAIL_AREA_MSTTableInfo
	tables["RAIL_COMP_MST"] = RAIL_COMP_MSTTableInfo
	tables["RAIL_LINE_MST"] = RAIL_LINE_MSTTableInfo
	tables["RECE_MAIL_HIST"] = RECE_MAIL_HISTTableInfo
	tables["REG_USER_MAIL"] = REG_USER_MAILTableInfo
	tables["SC_SPOT_PERM_MST"] = SC_SPOT_PERM_MSTTableInfo
	tables["SM_SENDER_LOG"] = SM_SENDER_LOGTableInfo
	tables["START_GRP_DTL_MST"] = START_GRP_DTL_MSTTableInfo
	tables["START_GRP_MST"] = START_GRP_MSTTableInfo
	tables["START_GRP_THREAD_MNG"] = START_GRP_THREAD_MNGTableInfo
	tables["SUMM_GRP_DTL_MST"] = SUMM_GRP_DTL_MSTTableInfo
	tables["SUMM_GRP_MST"] = SUMM_GRP_MSTTableInfo
	tables["TMPL_MST"] = TMPL_MSTTableInfo
	tables["TMPL_OPT_MST"] = TMPL_OPT_MSTTableInfo
	tables["USER_DEPT_MST"] = USER_DEPT_MSTTableInfo
	tables["USER_DLCT_CATE2_MST"] = USER_DLCT_CATE2_MSTTableInfo
	tables["USER_DLCT_DAY_MST"] = USER_DLCT_DAY_MSTTableInfo
	tables["USER_DLCT_D_MST"] = USER_DLCT_D_MSTTableInfo
	tables["USER_DLCT_MST"] = USER_DLCT_MSTTableInfo
	tables["USER_HUB_MST"] = USER_HUB_MSTTableInfo
	tables["USER_MAIL_MST"] = USER_MAIL_MSTTableInfo
	tables["USER_MST"] = USER_MSTTableInfo
	tables["USER_POST_MST"] = USER_POST_MSTTableInfo
	tables["WETH_MENU_AREA_MST"] = WETH_MENU_AREA_MSTTableInfo
	tables["WETH_OFFI_MST"] = WETH_OFFI_MSTTableInfo
	tables["WETH_QUAKE_AREA_MST"] = WETH_QUAKE_AREA_MSTTableInfo
	tables["WETH_QUAKE_CITY_MST"] = WETH_QUAKE_CITY_MSTTableInfo
	tables["WETH_SUB_AREA_MST"] = WETH_SUB_AREA_MSTTableInfo
	tables["WETH_VOLC_AREA_MST"] = WETH_VOLC_AREA_MSTTableInfo
	tables["WETH_VOLC_MST"] = WETH_VOLC_MSTTableInfo
	tables["WETH_WARN_MST"] = WETH_WARN_MSTTableInfo
	tables["WETH_XML_ADDR_MST"] = WETH_XML_ADDR_MSTTableInfo
	tables["WETH_XML_CITY_AREA_MST"] = WETH_XML_CITY_AREA_MSTTableInfo
	tables["WETH_XML_CITY_AREA_MST_20220322"] = WETH_XML_CITY_AREA_MST_20220322TableInfo
	tables["countries"] = countriesTableInfo
	tables["country"] = countryTableInfo
	tables["items"] = itemsTableInfo
}

// String describe the action
func (i Action) String() string {
	switch i {
	case Create:
		return "Create"
	case RetrieveOne:
		return "RetrieveOne"
	case RetrieveMany:
		return "RetrieveMany"
	case Update:
		return "Update"
	case Delete:
		return "Delete"
	case FetchDDL:
		return "FetchDDL"
	default:
		return fmt.Sprintf("unknown action: %d", int(i))
	}
}

// Model interface methods for database structs generated
type Model interface {
	TableName() string
	BeforeSave() error
	Prepare()
	Validate(action Action) error
	TableInfo() *TableInfo
}

// TableInfo describes a table in the database
type TableInfo struct {
	Name    string        `json:"name"`
	Columns []*ColumnInfo `json:"columns"`
}

// ColumnInfo describes a column in the database table
type ColumnInfo struct {
	Index              int    `json:"index"`
	GoFieldName        string `json:"go_field_name"`
	GoFieldType        string `json:"go_field_type"`
	JSONFieldName      string `json:"json_field_name"`
	ProtobufFieldName  string `json:"protobuf_field_name"`
	ProtobufType       string `json:"protobuf_field_type"`
	ProtobufPos        int    `json:"protobuf_field_pos"`
	Comment            string `json:"comment"`
	Notes              string `json:"notes"`
	Name               string `json:"name"`
	Nullable           bool   `json:"is_nullable"`
	DatabaseTypeName   string `json:"database_type_name"`
	DatabaseTypePretty string `json:"database_type_pretty"`
	IsPrimaryKey       bool   `json:"is_primary_key"`
	IsAutoIncrement    bool   `json:"is_auto_increment"`
	IsArray            bool   `json:"is_array"`
	ColumnType         string `json:"column_type"`
	ColumnLength       int64  `json:"column_length"`
	DefaultValue       string `json:"default_value"`
}

// GetTableInfo retrieve TableInfo for a table
func GetTableInfo(name string) (*TableInfo, bool) {
	val, ok := tables[name]
	return val, ok
}
