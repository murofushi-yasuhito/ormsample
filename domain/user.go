package domain

type User struct {
	//model
	UserId           string `gorm:"primary_key;column:USER_ID"`
	ClientId         string `gorm:"column:CLIENT_ID"`
	UserNo           string `gorm:"column:USER_NO" json:"UserNo"`
	UserNm           string `gorm:"column:USER_NM"`
	UserNmc          string `gorm:"column:USER_NMC"`
	LoginId          string `gorm:"column:LOGIN_ID"`
	Passwd           string `gorm:"column:PASSWD"`
	ClientAuthKbn    string `gorm:"column:CLIENT_AUTH_KBN"`
	StartMailFlg     string `gorm:"column:START_MAIL_FLG"`
	IkkatuMailFlg    string `gorm:"column:IKKATU_MAIL_FLG"`
	PrefCd           string `gorm:"column:PREF_CD"`
	CityCd           string `gorm:"column:CITY_CD"`
	MailSts          string `gorm:"column:MAIL_STS"`
	BatchFlg         string `gorm:"column:BATCH_FLG"`
	Token            string `gorm:"column:TOKEN"`
	DeviceType       string `gorm:"column:DEVICE_TYPE"`
	Memo             string `gorm:"column:MEMO"`
	PartnerFlg       string `gorm:"column:PARTNER_FLG"`
	ScUserLinkFlg    string `gorm:"column:SC_USER_LINK_FLG"`
	ScMngAuth        string `gorm:"column:SC_MNG_AUTH"`
	ScTelNum         string `gorm:"column:SC_TEL_NUM"`
	DesigDeptAuthKbn string `gorm:"column:DESIG_DEPT_AUTH_KBN"`
	InitPasswdFlg    string `gorm:"column:INIT_PASSWD_FLG"`
	OneTimePasswd    string `gorm:"column:ONE_TIME_PASSWD"`
	OneTimeExpiry    string `gorm:"column:ONE_TIME_EXPIRY"`
	UpdaDte          string `gorm:"column:UPDA_DTE"`
	UpdaUserId       string `gorm:"column:UPDA_USER_ID"`
	CreaDte          string `gorm:"column:CREA_DTE"`
	CreaUserId       string `gorm:"column:CREA_USER_ID"`
	StartGrpId       string `gorm:"column:START_GRP_ID"`
}
