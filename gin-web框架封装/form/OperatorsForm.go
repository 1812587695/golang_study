package form

type OperatorsForm struct {
	Title           string  `form:"title" valid:"Required"`
	BusinessLicense string  `form:"business_license" valid:"Required"`
	Username        string  `form:"username" valid:"Required"`
	IDCard          string  `form:"id_card" valid:"Required"`
	IDCardFront     string  `form:"id_card_front" valid:"Required"`
	IDCardBack      string  `form:"id_card_back" valid:"Required"`
	Province        int     `form:"province"`
	City            int     `form:"city"`
	Area            int     `form:"area"`
	Phone           string  `form:"phone" valid:"Required"`
	Code            string  `form:"code" valid:"Required"`
	BankCardName    string  `form:"bank_card_name" valid:"Required"`
	BankCardNo      string  `form:"bank_card_no" valid:"Required"`
	BankName        string  `form:"bank_name" valid:"Required"`
	ManageCenter    int     `form:"manage_center" valid:"Required"`
	AgencyMoney     float64 `form:"agency_money" valid:"Required"`
	RefereeName     string  `form:"referee_name" valid:"Required"`
	RefereePhone    string  `form:"referee_phone" valid:"Required"`
}
