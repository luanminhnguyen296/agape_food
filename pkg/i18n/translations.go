package i18n

type Translation struct {
	// Layout
	MetaTitle           string
	MetaDesc            string
	LogoSubtitle        string
	NavHome             string
	NavServices         string
	NavProcess          string
	NavWhyChoose        string
	NavMenu             string
	NavContact          string
	BtnZalo             string
	BtnZaloQuote        string
	BtnCallNow          string
	BtnCallHotline      string
	MobileMenuClose     string
	FooterDesc          string
	FooterContactTitle  string
	FooterConnectTitle  string
	FooterRights        string
	FooterBranch        string

	// Index Hero
	HeroBadge           string
	HeroTitle           string
	HeroTitleSpan       string
	HeroDesc            string
	HeroCheckATTP       string
	HeroCheckBep        string
	HeroCheckKiemThuc   string
	HeroCheckGiaoHang   string
	HeroCheckAudit      string
	BadgeFloatNuyenLieu string
	BadgeFloatKhachHang string

	// Pain Points
	PainTitle           string
	PainTitleSpan       string
	PainCard1           string
	PainCard2           string
	PainCard3           string
	PainCard4           string
	PainCard5           string
	PainCard6           string
	PainConclusion      string

	// Solutions
	SolTitle            string
	SolTitleSpan        string
	SolDesc             string
	SolCard1Title       string
	SolCard1Desc        string
	SolCard2Title       string
	SolCard2Desc        string
	SolCard3Title       string
	SolCard3Desc        string
	SolCard4Title       string
	SolCard4Desc        string
	SolCard5Title       string
	SolCard5Desc        string
	SolCard6Title       string
	SolCard6Desc        string
	SolBtnZalo          string

	// Process
	ProcTitle           string
	ProcTitleSpan       string
	ProcDesc            string
	ProcStep1Title      string
	ProcStep1Desc       string
	ProcStep2Title      string
	ProcStep2Desc       string
	ProcStep3Title      string
	ProcStep3Desc       string
	ProcStep4Title      string
	ProcStep4Desc       string
	ProcStep5Title      string
	ProcStep5Desc       string

	// Audit
	AuditBadge          string
	AuditTitle          string
	AuditDesc           string
	AuditNoteSmall      string

	// Documents
	DocsTitle           string
	DocsTitleSpan       string
	DocsDesc            string
	DocsGroup1Title     string
	DocsGroup1Item1     string
	DocsGroup1Item2     string
	DocsGroup1Item3     string
	DocsGroup2Title     string
	DocsGroup2Item1     string
	DocsGroup2Item2     string
	DocsGroup2Item3     string
	DocsGroup3Title     string
	DocsGroup3Item1     string
	DocsGroup3Item2     string
	DocsGroup3Item3     string
	DocsGroup3Item4     string
	DocsGroup4Title     string
	DocsGroup4Item1     string
	DocsGroup4Item2     string
	DocsGroup4Item3     string
	DocsHighlight       string

	// Why Choose
	WhyTitle            string
	WhyTitleSpan        string
	WhyDesc             string
	WhyItem1Title       string
	WhyItem1Desc        string
	WhyItem2Title       string
	WhyItem2Desc        string
	WhyItem3Title       string
	WhyItem3Desc        string
	WhyItem4Title       string
	WhyItem4Desc        string
	WhyItem5Title       string
	WhyItem5Desc        string
	WhyItem6Title       string
	WhyItem6Desc        string
	WhyItem7Title       string
	WhyItem7Desc        string
	WhyItem8Title       string
	WhyItem8Desc        string
	WhySlide1           string
	WhySlide2           string
	WhySlide3           string
	WhySlide4           string
	WhySlide5           string
	WhySlide6           string

	// Menu
	MenuTitle           string
	MenuTitleSpan       string
	MenuDesc            string
	MenuBtnZalo         string
	MenuTab1            string
	MenuTab2            string
	MenuTab3            string
	MenuTab4            string
	MenuTab5            string

	// Testimonials
	TestBadge           string
	TestTitle           string
	TestTitleSpan       string
	TestDesc            string
	Test1Quote          string
	Test1User           string
	Test1Title          string
	Test2Quote          string
	Test2User           string
	Test2Title          string
	Test3Quote          string
	Test3User           string
	Test3Title          string

	// Clients
	ClientsTitle        string

	// FAQ & Commitments
	CommitTitle         string
	CommitItem1         string
	CommitItem2         string
	CommitItem3         string
	CommitItem4         string
	CommitItem5         string
	CommitItem6         string
	CommitItem7         string
	FaqTitle            string
	FaqQ1               string
	FaqA1               string
	FaqQ2               string
	FaqA2               string
	FaqQ3               string
	FaqA3               string
	FaqQ4               string
	FaqA4               string
	FaqQ5               string
	FaqA5               string
	FaqQ6               string
	FaqA6               string
	FaqQ7               string
	FaqA7               string
	FaqQ8               string
	FaqA8               string

	// CTA Banner
	CtaBannerTitle      string
	CtaBannerBtnZalo    string
	CtaBannerBtnCall    string

	// 404 Page
	NotFoundTitle       string
	NotFoundBadge       string
	NotFoundSubtitle    string
	NotFoundDesc        string
	NotFoundBtn         string
}

func Get(lang string) Translation {
	if lang == "en" {
		return EN()
	}
	return VI()
}

func VI() Translation {
	return Translation{
		MetaTitle:           "Agape Food | Suất Ăn Công Nghiệp Long An – An Toàn, Đầy Đủ Hồ Sơ ATTP",
		MetaDesc:            "Agape Food – Suất ăn công nghiệp, suất ăn văn phòng, suất ăn trường học an toàn, chất lượng tại Long An. Đạt chuẩn ATTP, hỗ trợ Customer Audit & Supplier Audit cho doanh nghiệp FDI.",
		LogoSubtitle:        "An toàn - Chất lượng - Tận tâm",
		NavHome:             "Trang chủ",
		NavServices:         "Dịch vụ",
		NavProcess:          "Quy trình",
		NavWhyChoose:        "Về chúng tôi",
		NavMenu:             "Thực đơn",
		NavContact:          "Liên hệ",
		BtnZalo:             "Nhấn Zalo",
		BtnZaloQuote:        "Nhấn Zalo nhận báo giá",
		BtnCallNow:          "Gọi ngay",
		BtnCallHotline:      "Gọi Hotline",
		MobileMenuClose:     "Đóng menu",
		FooterDesc:          "Đồng hành cùng doanh nghiệp mang lại những suất ăn công nghiệp an toàn, chất lượng, đảm bảo vệ sinh an toàn thực phẩm chuẩn quốc tế.",
		FooterContactTitle:  "Thông tin liên hệ",
		FooterConnectTitle:  "Kết nối với chúng tôi",
		FooterRights:        "© 2026 Agape Food. All rights reserved. Chăm chút từng bữa ăn.",
		FooterBranch:        "Địa điểm kinh doanh 02",

		HeroBadge:           "CLEAN FOOD – STRONG BUSINESS",
		HeroTitle:           "SUẤT ĂN AN TOÀN ",
		HeroTitleSpan:       "DOANH NGHIỆP VỮNG VÀNG",
		HeroDesc:            "Agape Foods cung cấp suất ăn cho nhà máy, văn phòng, trường học và doanh nghiệp FDI với quy trình vận hành chuẩn, hồ sơ ATTP đầy đủ, hỗ trợ khách hàng trong các kỳ Customer Audit, Supplier Audit theo yêu cầu đối tác EU, US và quốc tế.",
		HeroCheckATTP:       "ATTP<br/>đầy đủ",
		HeroCheckBep:        "Bếp 1 chiều<br/>chuẩn hóa",
		HeroCheckKiemThuc:   "Kiểm thực<br/>3 bước",
		HeroCheckGiaoHang:   "Giao đúng giờ<br/>cam kết",
		HeroCheckAudit:      "Hỗ trợ<br/>Audit",
		BadgeFloatNuyenLieu: "Nguyên liệu<br/>tự nhiên",
		BadgeFloatKhachHang: "Khách hàng<br/>tin tưởng",

		PainTitle:           "Doanh nghiệp gặp khó khăn gì khi ",
		PainTitleSpan:       "quản lý suất ăn?",
		PainCard1:           "Nhà cung cấp giao trễ, món ăn không ổn định",
		PainCard2:           "Nhân viên phản ánh chất lượng bữa ăn",
		PainCard3:           "Menu lặp lại, thiếu đa dạng",
		PainCard4:           "Khó kiểm soát nguồn gốc nguyên liệu",
		PainCard5:           "Thiếu hồ sơ khi có đoàn kiểm tra hoặc Audit",
		PainCard6:           "Không biết cần chuẩn bị gì khi khách hàng EU, US đánh giá",
		PainConclusion:      "<strong>Kết luận:</strong> Agape Foods không chỉ cung cấp suất ăn, mà còn đồng hành cùng doanh nghiệp trong việc kiểm soát chất lượng, an toàn thực phẩm và hồ sơ tuân thủ.",

		SolTitle:            "Giải pháp suất ăn ",
		SolTitleSpan:        "phù hợp cho từng mô hình doanh nghiệp",
		SolDesc:             "Đáp ứng hoàn hảo mọi nhu cầu của các nhóm đối tác từ nhà máy, văn phòng đến trường học.",
		SolCard1Title:       "Suất ăn công nghiệp",
		SolCard1Desc:        "Cung cấp số lượng lớn suất ăn đạt chuẩn an toàn vệ sinh thực phẩm, tiếp thêm năng lượng cho công nhân sản xuất.",
		SolCard2Title:       "Suất ăn văn phòng",
		SolCard2Desc:        "Bữa ăn thanh đạm, dinh dưỡng cân đối và trình bày đẹp mắt cho khối nhân sự văn phòng năng động.",
		SolCard3Title:       "Suất ăn chế biến sẵn",
		SolCard3Desc:        "Chế biến sẵn và đóng gói tiện lợi, bảo quản tối ưu để phân phối nhanh cho các ca kíp trực tuyến.",
		SolCard4Title:       "Canteen tại nhà máy",
		SolCard4Desc:        "Thiết lập và quản lý vận hành canteen trực tiếp tại nhà máy đối tác với quy trình phục vụ chuẩn hóa.",
		SolCard5Title:       "Suất ăn cho chuyên gia",
		SolCard5Desc:        "Thực đơn cao cấp chuẩn nhà hàng phục vụ cho chuyên gia nước ngoài và các đối tác VIP.",
		SolCard6Title:       "Suất ăn trường học",
		SolCard6Desc:        "Cân đối dưỡng chất khoa học giúp học sinh phát triển thể trạng tốt nhất, an toàn và dễ ăn.",
		SolBtnZalo:          "Nhấn Zalo để được tư vấn mô hình phù hợp",

		ProcTitle:           "Quy trình hợp tác ",
		ProcTitleSpan:       "rõ ràng – nhanh chóng – chuyên nghiệp",
		ProcDesc:            "Chúng tôi tối ưu hóa từng bước hợp tác để tiết kiệm thời gian và mang lại sự an tâm tuyệt đối cho quý doanh nghiệp.",
		ProcStep1Title:      "Nhận thông tin",
		ProcStep1Desc:       "Tiếp nhận nhu cầu, quy mô suất ăn và các tiêu chuẩn đặc thù của doanh nghiệp.",
		ProcStep2Title:      "Khảo sát nhu cầu",
		ProcStep2Desc:       "Đội ngũ chuyên gia khảo sát canteen, bếp ăn và khảo sát khẩu vị của cán bộ công nhân viên.",
		ProcStep3Title:      "Dùng thử miễn phí",
		ProcStep3Desc:       "Hỗ trợ cung cấp các suất ăn dùng thử miễn phí để kiểm định chất lượng và đóng góp ý kiến.",
		ProcStep4Title:      "Ký kết hợp đồng",
		ProcStep4Desc:       "Thống nhất thực đơn, đơn giá, thời gian phục vụ và hoàn thiện các cam kết dịch vụ pháp lý.",
		ProcStep5Title:      "Thực hiện dịch vụ theo cam kết",
		ProcStep5Desc:       "Vận hành bếp ăn theo quy trình 1 chiều chuẩn hóa, cung cấp đầy đủ hồ sơ ATTP cho mọi kỳ Audit.",

		AuditBadge:          "ĐỒNG HÀNH CÙNG DOANH NGHIỆP TRONG CÁC KỲ AUDIT",
		AuditTitle:          "Sẵn sàng hồ sơ phục vụ Customer Audit, Supplier Audit và các tiêu chuẩn quốc tế",
		AuditDesc:           "Với các doanh nghiệp xuất khẩu, nhà máy FDI hoặc đơn vị có khách hàng từ EU, US, Nhật Bản, Hàn Quốc…, việc kiểm soát nhà cung cấp suất ăn là một phần quan trọng trong quá trình đánh giá. Agape Foods xây dựng hệ thống hồ sơ đầy đủ để hỗ trợ doanh nghiệp khi có yêu cầu kiểm tra từ khách hàng, tập đoàn quốc tế hoặc tổ chức chứng nhận.",
		AuditNoteSmall:      "<i class=\"ph-bold ph-info\" style=\"color: var(--accent);\"></i>Agape Foods hỗ trợ chuẩn bị và cung cấp hồ sơ phục vụ các kỳ đánh giá theo yêu cầu của khách hàng/doanh nghiệp.",

		DocsTitle:           "Hệ thống hồ sơ ",
		DocsTitleSpan:       "Agape Foods có thể cung cấp",
		DocsDesc:            "Chúng tôi chuẩn bị sẵn sàng, minh bạch các bộ hồ sơ chất lượng để hỗ trợ mọi quy trình kiểm tra của doanh nghiệp.",
		DocsGroup1Title:     "Hồ sơ pháp lý",
		DocsGroup1Item1:     "Giấy phép đăng ký kinh doanh",
		DocsGroup1Item2:     "Giấy chứng nhận đủ điều kiện ATTP",
		DocsGroup1Item3:     "Biên bản kiểm tra gần nhất của cơ quan quản lý nhà nước",
		DocsGroup2Title:     "Hồ sơ truy xuất nguồn gốc",
		DocsGroup2Item1:     "Hợp đồng nguyên liệu đầu vào",
		DocsGroup2Item2:     "Giấy đủ điều kiện ATTP của đối tác cung cấp nguyên liệu",
		DocsGroup2Item3:     "Hóa đơn mua hàng chứng minh nguồn nguyên liệu từ NCC đã ký hợp đồng",
		DocsGroup3Title:     "Hồ sơ chất lượng & vận hành",
		DocsGroup3Item1:     "Giấy kiểm định nguồn nước",
		DocsGroup3Item2:     "Giấy kiểm định thực phẩm định kỳ",
		DocsGroup3Item3:     "Hồ sơ kiểm thực 3 bước theo quy định Bộ Y tế",
		DocsGroup3Item4:     "Quy trình bếp 1 chiều",
		DocsGroup4Title:     "Hồ sơ nhân sự",
		DocsGroup4Item1:     "Hồ sơ khám sức khỏe nhân viên",
		DocsGroup4Item2:     "Giấy chứng nhận tập huấn kiến thức ATTP",
		DocsGroup4Item3:     "Giấy chứng nhận hoàn thành nghĩa vụ tài chính tại cơ quan BHXH",
		DocsHighlight:       "Không chỉ bán suất ăn – Agape Foods giúp doanh nghiệp sẵn sàng khi có đoàn Audit.",

		WhyTitle:            "Vì sao chọn ",
		WhyTitleSpan:        "Agape Foods?",
		WhyDesc:             "Khác biệt tạo nên thương hiệu uy tín hàng đầu trong lĩnh vực suất ăn công nghiệp.",
		WhyItem1Title:       "Nguyên liệu rõ nguồn gốc",
		WhyItem1Desc:        "Liên kết trực tiếp nông trại đạt chuẩn VietGAP.",
		WhyItem2Title:       "Thực đơn đa dạng, thay đổi định kỳ",
		WhyItem2Desc:        "Hơn 200 món ăn thay đổi định kỳ không trùng lặp.",
		WhyItem3Title:       "Đội ngũ bếp có kinh nghiệm",
		WhyItem3Desc:        "Đầu bếp được đào tạo bài bản về an toàn thực phẩm.",
		WhyItem4Title:       "Giao đúng giờ",
		WhyItem4Desc:        "Cam kết giờ ăn ổn định giúp doanh nghiệp tối ưu ca kíp.",
		WhyItem5Title:       "Kiểm thực 3 bước",
		WhyItem5Desc:        "Quy trình giám sát khép kín trước khi phục vụ.",
		WhyItem6Title:       "Lưu mẫu thức ăn theo quy định",
		WhyItem6Desc:        "Mẫu thức ăn được niêm phong lưu giữ 24h.",
		WhyItem7Title:       "Hồ sơ ATTP đầy đủ",
		WhyItem7Desc:        "Đầy đủ chứng nhận kiểm định nguồn nước, sức khỏe nhân viên.",
		WhyItem8Title:       "Hỗ trợ Audit cho doanh nghiệp FDI",
		WhyItem8Desc:        "Chuẩn bị hồ sơ và hiện diện trực tiếp trong mọi kỳ đánh giá.",
		WhySlide1:           "Bếp trung tâm",
		WhySlide2:           "Khu chế biến",
		WhySlide3:           "Canteen tại nhà máy",
		WhySlide4:           "Xe giao hàng",
		WhySlide5:           "Hồ sơ & chứng nhận",
		WhySlide6:           "Suất ăn mẫu",

		MenuTitle:           "Thiết kế thực đơn theo ",
		MenuTitleSpan:       "nhu cầu & ngân sách",
		MenuDesc:            "Thực đơn được thiết kế thông minh, khoa học và luân phiên thay đổi mỗi ngày giúp công nhân viên luôn ngon miệng.",
		MenuBtnZalo:         "Nhấn Zalo nhận menu mẫu theo ngân sách",
		MenuTab1:            "Công nghiệp",
		MenuTab2:            "Văn phòng",
		MenuTab3:            "Chuyên gia",
		MenuTab4:            "Trường học",
		MenuTab5:            "Theo ca / kíp",

		TestBadge:           "Ý KIẾN ĐỐI TÁC",
		TestTitle:           "Khách hàng ",
		TestTitleSpan:       "nói gì về chúng tôi?",
		TestDesc:            "Sự hài lòng và đồng hành lâu dài của các đối tác lớn, doanh nghiệp FDI là minh chứng rõ nhất cho chất lượng dịch vụ của Agape Foods.",
		Test1Quote:          "\"Từ ngày chuyển sang đối tác Agape Foods, chúng tôi hoàn toàn yên tâm trong các kỳ kiểm tra của tập đoàn Knauf toàn cầu. Suất ăn luôn tươi ngon, nóng sốt và đầy đủ hồ sơ ATTP phục vụ Audit.\"",
		Test1User:           "Bà Nguyễn Thị Lan Anh",
		Test1Title:          "Trưởng phòng Nhân sự - Knauf Việt Nam",
		Test2Quote:          "\"Chúng tôi đánh giá cao quy trình kiểm thực 3 bước và lưu mẫu thức ăn nghiêm ngặt của Agape. Suất ăn đa dạng, hợp khẩu vị của cả các chuyên gia nước ngoài lẫn công nhân nhà máy.\"",
		Test2User:           "Mr. Kenji Sato",
		Test2Title:          "Giám đốc Vận hành - Nhà máy Kizuna",
		Test3Quote:          "\"Dịch vụ suất ăn học đường của Agape Foods làm hài lòng cả ban giám hiệu và phụ huynh học sinh. Nguồn gốc nguyên liệu VietGAP minh bạch là lý do chúng tôi chọn đồng hành lâu dài.\"",
		Test3User:           "Bà Lê Thanh Hải",
		Test3Title:          "Hiệu trưởng - Trường Tiểu học Việt Mỹ",

		ClientsTitle:        "ĐỒNG HÀNH CÙNG CÁC ĐỐI TÁC LỚN & DOANH NGHIỆP FDI",

		CommitTitle:         "Cam kết của Agape Foods",
		CommitItem1:         "Không sử dụng thực phẩm không rõ nguồn gốc",
		CommitItem2:         "Thực hiện kiểm thực 3 bước",
		CommitItem3:         "Lưu mẫu thức ăn theo quy định",
		CommitItem4:         "Giao hàng đúng thời gian cam kết",
		CommitItem5:         "Phản hồi nhanh khi có góp ý từ người dùng",
		CommitItem6:         "Cung cấp hồ sơ khi doanh nghiệp cần phục vụ kiểm tra, Audit",
		CommitItem7:         "Thực hiện đúng cam kết trong hợp đồng",
		FaqTitle:            "Câu hỏi thường gặp",
		FaqQ1:               "Agape Foods nhận phục vụ từ bao nhiêu suất/ngày?",
		FaqA1:               "Chúng tôi nhận cung cấp suất ăn từ quy mô 100 suất/ngày trở lên cho các doanh nghiệp, nhà máy, xí nghiệp tại khu vực Long An và lân cận.",
		FaqQ2:               "Có được dùng thử suất ăn miễn phí không?",
		FaqA2:               "Có. Chúng tôi hỗ trợ phục vụ suất ăn dùng thử miễn phí từ 3-5 ngày để doanh nghiệp và người lao động đánh giá hương vị cũng như chất lượng phục vụ trước khi ký hợp đồng chính thức.",
		FaqQ3:               "Có thiết kế menu theo ngân sách không?",
		FaqA3:               "Có, thực đơn sẽ được cân đối định lượng dinh dưỡng dựa trên ngân sách cụ thể của từng doanh nghiệp nhằm bảo đảm hiệu quả chi phí tốt nhất cho đối tác.",
		FaqQ4:               "Có phục vụ ca đêm hoặc nhiều ca sản xuất không?",
		FaqA4:               "Có. Hệ thống bếp ăn của Agape Foods vận hành 24/7 để đáp ứng linh hoạt các ca kíp sản xuất và ca làm đêm của nhà máy.",
		FaqQ5:               "Có cung cấp hồ sơ ATTP và hồ sơ Audit không?",
		FaqA5:               "Có. Chúng tôi bàn giao đầy đủ bộ hồ sơ pháp lý gồm giấy chứng nhận cơ sở đủ điều kiện ATTP, giấy kiểm định nguồn nước, nguồn gốc thực phẩm và các chứng nhận sức khỏe của nhân viên định kỳ.",
		FaqQ6:               "Có hỗ trợ khi khách hàng nước ngoài kiểm tra nhà máy không?",
		FaqA6:               "Chúng tôi sẵn sàng cử nhân sự quản lý chất lượng trực tiếp tham gia cùng quý doanh nghiệp đón tiếp và giải trình hồ sơ khi có đoàn Audit từ đối tác nước ngoài.",
		FaqQ7:               "Có xuất hóa đơn VAT không?",
		FaqA7:               "Tất cả các dịch vụ suất ăn công nghiệp của Agape Foods đều được cung cấp đầy đủ hóa đơn tài chính hợp lệ cho doanh nghiệp.",
		FaqQ8:               "Có thể chế biến tại canteen của nhà máy không?",
		FaqA8:               "Có. Tùy thuộc vào cơ sở hạ tầng của đối tác, chúng tôi có thể setup đội ngũ đầu bếp và nguyên liệu đến chế biến nóng trực tiếp tại bếp ăn của nhà máy.",

		CtaBannerTitle:      "ĐỐI TÁC SUẤT ĂN AN TOÀN – ỔN ĐỊNH – SẴN SÀNG HỒ SƠ AUDIT",
		CtaBannerBtnZalo:    "Nhấn Zalo nhận báo giá",
		CtaBannerBtnCall:    "Gọi Hotline",

		NotFoundTitle:       "Trang Không Tồn Tại (404) | Agape Food",
		NotFoundBadge:       "Lỗi 404",
		NotFoundSubtitle:    "Trang Không Tồn Tại",
		NotFoundDesc:        "Đường dẫn bạn truy cập hiện tại không khả dụng hoặc đã bị thay đổi. Vui lòng quay trở lại Trang chủ.",
		NotFoundBtn:         "Quay lại Trang chủ",
	}
}

func EN() Translation {
	return Translation{
		MetaTitle:           "Agape Food | Industrial Catering in Long An – Safe, Full Food Safety Portfolio",
		MetaDesc:            "Agape Food – Safe, high-quality industrial catering, office lunch, and school meal services in Long An. Compliant with international standards, supporting Customer Audit & Supplier Audit for FDI companies.",
		LogoSubtitle:        "Safety - Quality - Dedication",
		NavHome:             "Home",
		NavServices:         "Services",
		NavProcess:          "Process",
		NavWhyChoose:        "About Us",
		NavMenu:             "Menu",
		NavContact:          "Contact",
		BtnZalo:             "Zalo Us",
		BtnZaloQuote:        "Get Quotation on Zalo",
		BtnCallNow:          "Call Now",
		BtnCallHotline:      "Call Hotline",
		MobileMenuClose:     "Close menu",
		FooterDesc:          "Accompanying enterprises to bring safe, high-quality industrial meals, ensuring international standard food hygiene and safety.",
		FooterContactTitle:  "Contact Information",
		FooterConnectTitle:  "Connect With Us",
		FooterRights:        "© 2026 Agape Food. All rights reserved. Caring for every meal.",
		FooterBranch:        "Business Location 02",

		HeroBadge:           "CLEAN FOOD – STRONG BUSINESS",
		HeroTitle:           "SAFE MEALS FOR ",
		HeroTitleSpan:       "SUSTAINABLE ENTERPRISES",
		HeroDesc:            "Agape Foods provides catering services for factories, offices, schools, and FDI enterprises with standardized operating procedures, comprehensive food safety documents, supporting clients in Customer Audits, Supplier Audits as required by EU, US, and international partners.",
		HeroCheckATTP:       "Food Safety<br/>Portfolio",
		HeroCheckBep:        "1-way Kitchen<br/>Standard",
		HeroCheckKiemThuc:   "3-step Food<br/>Inspection",
		HeroCheckGiaoHang:   "On-time Delivery<br/>Guaranteed",
		HeroCheckAudit:      "Audit<br/>Support",
		BadgeFloatNuyenLieu: "Natural<br/>Ingredients",
		BadgeFloatKhachHang: "Trusted<br/>Clients",

		PainTitle:           "What challenges do enterprises face when ",
		PainTitleSpan:       "managing catering?",
		PainCard1:           "Late delivery by suppliers, inconsistent meals",
		PainCard2:           "Employees complain about meal quality",
		PainCard3:           "Repetitive, non-diverse menu",
		PainCard4:           "Difficult to verify ingredient origins",
		PainCard5:           "Lack of documents during audits or inspections",
		PainCard6:           "Unsure what to prepare for EU, US customer evaluations",
		PainConclusion:      "<strong>Conclusion:</strong> Agape Foods does not just sell meals; we partner with enterprises to control quality, food safety, and compliance documentation.",

		SolTitle:            "Catering Solutions ",
		SolTitleSpan:        "Tailored for Every Business Model",
		SolDesc:             "Perfectly meeting all needs of partner groups from factories, offices to schools.",
		SolCard1Title:       "Industrial Catering",
		SolCard1Desc:        "Providing large volumes of meals meeting food hygiene and safety standards, status-checking factory workers.",
		SolCard2Title:       "Office Lunch",
		SolCard2Desc:        "Balanced, nutritious, and beautifully presented lunches for active office staff.",
		SolCard3Title:       "Pre-packaged Meals",
		SolCard3Desc:        "Freshly prepared and conveniently packaged, optimally preserved for quick distribution to online shifts.",
		SolCard4Title:       "Factory Canteen Management",
		SolCard4Desc:        "Setting up and operating canteens directly at partners' factories with standardized processes.",
		SolCard5Title:       "Expert Meals",
		SolCard5Desc:        "Premium restaurant-standard menus serving foreign experts and VIP partners.",
		SolCard6Title:       "School Catering",
		SolCard6Desc:        "Scientifically balanced nutrients to help students develop their best physique, safe and easy to eat.",
		SolBtnZalo:          "Zalo us for suitable model consulting",

		ProcTitle:           "Clear, Fast & Professional ",
		ProcTitleSpan:       "Cooperation Process",
		ProcDesc:            "We optimize every step of cooperation to save time and bring absolute peace of mind to your business.",
		ProcStep1Title:      "Receive Info",
		ProcStep1Desc:       "Receiving information on needs, meal scale, and specific standards of the enterprise.",
		ProcStep2Title:      "Site Survey",
		ProcStep2Desc:       "Our experts inspect the canteen and kitchen area, and survey employees' taste preferences.",
		ProcStep3Title:      "Free Trial",
		ProcStep3Desc:       "Providing free trial meals to verify quality and collect feedback.",
		ProcStep4Title:      "Contract Signing",
		ProcStep4Desc:       "Agreeing on menus, unit prices, service times, and completing legal service commitments.",
		ProcStep5Title:      "Service Execution",
		ProcStep5Desc:       "Operating the kitchen under a standardized 1-way flow, providing full food safety documents for all audits.",

		AuditBadge:          "ACCOMPANYING ENTERPRISES IN ALL AUDITS",
		AuditTitle:          "Ready with Documentation for Customer Audits, Supplier Audits & International Standards",
		AuditDesc:           "For exporting enterprises, FDI factories, or units with customers from the EU, US, Japan, South Korea, etc., monitoring the catering supplier is a crucial part of the audit process. Agape Foods builds a comprehensive document system to assist enterprises upon request from customers, international corporations, or certification bodies.",
		AuditNoteSmall:      "<i class=\"ph-bold ph-info\" style=\"color: var(--accent);\"></i>Agape Foods supports preparing and providing documents for audits upon customer/enterprise request.",

		DocsTitle:           "Documentation System ",
		DocsTitleSpan:       "Agape Foods Can Provide",
		DocsDesc:            "We prepare transparent and ready-to-use quality records to assist all business inspection processes.",
		DocsGroup1Title:     "Legal Documents",
		DocsGroup1Item1:     "Business Registration License",
		DocsGroup1Item2:     "Food Safety Eligibility Certificate",
		DocsGroup1Item3:     "Latest inspection report from state regulatory authorities",
		DocsGroup2Title:     "Traceability Documents",
		DocsGroup2Item1:     "Input ingredient contracts",
		DocsGroup2Item2:     "Food safety certificates of ingredient suppliers",
		DocsGroup2Item3:     "Purchasing invoices proving ingredient sources from contracted suppliers",
		DocsGroup3Title:     "Quality & Operation Documents",
		DocsGroup3Item1:     "Water test certificates",
		DocsGroup3Item2:     "Periodic food testing records",
		DocsGroup3Item3:     "3-step food inspection log following Ministry of Health rules",
		DocsGroup3Item4:     "1-way kitchen process chart",
		DocsGroup4Title:     "Staff Documents",
		DocsGroup4Item1:     "Employee periodic health checkup records",
		DocsGroup4Item2:     "Food safety training certificates",
		DocsGroup4Item3:     "Certificates of social security tax compliance",
		DocsHighlight:       "More than catering – Agape Foods helps your business get ready for audits.",

		WhyTitle:            "Why Choose ",
		WhyTitleSpan:        "Agape Foods?",
		WhyDesc:             "Differences make us a leading reputable brand in industrial catering.",
		WhyItem1Title:       "Traceable Ingredients",
		WhyItem1Desc:        "Directly partnered with VietGAP certified farms.",
		WhyItem2Title:       "Diverse, Rotating Menu",
		WhyItem2Desc:        "Over 200 dishes rotated regularly with no repetition.",
		WhyItem3Title:       "Experienced Kitchen Crew",
		WhyItem3Desc:        "Chefs are thoroughly trained in food safety.",
		WhyItem4Title:       "On-Time Delivery",
		WhyItem4Desc:        "Guaranteed stable meal hours to optimize shifts.",
		WhyItem5Title:       "3-step Food Inspection",
		WhyItem5Desc:        "Closed monitoring process prior to serving.",
		WhyItem6Title:       "Mandatory Food Sampling",
		WhyItem6Desc:        "Food samples are sealed and stored for 24 hours.",
		WhyItem7Title:       "Full Food Safety Certification",
		WhyItem7Desc:        "Water test and employee health certificates fully provided.",
		WhyItem8Title:       "Audit Support for FDI Companies",
		WhyItem8Desc:        "Preparing documents and physically presenting during audits.",
		WhySlide1:           "Central Kitchen",
		WhySlide2:           "Processing Area",
		WhySlide3:           "Factory Canteen",
		WhySlide4:           "Delivery Fleet",
		WhySlide5:           "Documents & Certificates",
		WhySlide6:           "Sample Meals",

		MenuTitle:           "Menu Customization by ",
		MenuTitleSpan:       "Demands & Budget",
		MenuDesc:            "Menus are intelligently and scientifically designed, rotating daily to keep employees happy and satisfied.",
		MenuBtnZalo:         "Zalo us for custom menu template by budget",
		MenuTab1:            "Industrial",
		MenuTab2:            "Office",
		MenuTab3:            "Experts",
		MenuTab4:            "Schools",
		MenuTab5:            "Shift Meals",

		TestBadge:           "PARTNER OPINIONS",
		TestTitle:           "What Our ",
		TestTitleSpan:       "Clients Say",
		TestDesc:            "The satisfaction and long-term partnership of major companies and FDI factories are the best proof of Agape Foods' service quality.",
		Test1Quote:          "\"Since switching to Agape Foods, we have been completely worry-free during global Knauf group audits. Meals are always fresh, hot, and full food safety records are ready for Audits.\"",
		Test1User:           "Mrs. Nguyen Thi Lan Anh",
		Test1Title:          "HR Director - Knauf Vietnam",
		Test2Quote:          "\"We highly appreciate Agape's strict 3-step food inspection and food sampling processes. Meals are diverse, satisfying both foreign experts and factory workers.\"",
		Test2User:           "Mr. Kenji Sato",
		Test2Title:          "Operations Manager - Kizuna Factory",
		Test3Quote:          "\"Agape Foods' school catering service pleases both the school board and parents. The transparent VietGAP ingredient source is the reason we partner long-term.\"",
		Test3User:           "Mrs. Le Thanh Hải",
		Test3Title:          "Principal - Viet My Primary School",

		ClientsTitle:        "PARTNERING WITH MAJOR FDI COMPANIES",

		CommitTitle:         "Agape Foods Commitments",
		CommitItem1:         "No food of unknown origin",
		CommitItem2:         "Implementation of 3-step food inspection",
		CommitItem3:         "Mandatory food sampling stored for 24h",
		CommitItem4:         "Delivery exactly on committed schedules",
		CommitItem5:         "Fast response to users' feedback",
		CommitItem6:         "Providing documents for inspections and audits",
		CommitItem7:         "Strictly executing contract agreements",
		FaqTitle:            "Frequently Asked Questions",
		FaqQ1:               "What is the minimum scale Agape Foods serves?",
		FaqA1:               "We supply meals starting from 100 portions/day for enterprises, factories, and workshops in Long An and neighboring areas.",
		FaqQ2:               "Are free trial meals available?",
		FaqA2:               "Yes. We support serving free trial meals for 3-5 days so the business and workers can evaluate the taste and service quality before signing a contract.",
		FaqQ3:               "Do you customize menus based on budgets?",
		FaqA3:               "Yes, menus are nutritionally portioned based on each business's specific budget to ensure the best cost efficiency for our partners.",
		FaqQ4:               "Do you serve night shifts or multiple shifts?",
		FaqA4:               "Yes. Agape Foods' kitchen operates 24/7 to flexibly meet factory production shifts and night shifts.",
		FaqQ5:               "Are food safety and audit portfolios provided?",
		FaqA5:               "Yes. We hand over a full legal set including food safety certificates, water test records, food tracing documents, and staff health check reports.",
		FaqQ6:               "Do you assist when foreign clients audit our factory?",
		FaqA6:               "We are ready to assign quality control staff to join and explain the catering documentation when you host audit delegations from overseas.",
		FaqQ7:               "Do you issue VAT invoices?",
		FaqA7:               "All industrial catering services of Agape Foods are fully provided with valid financial invoices for the company.",
		FaqQ8:               "Can you cook directly at the factory's canteen?",
		FaqA8:               "Yes. Depending on the partner's infrastructure, we can assign chefs and supply ingredients to cook hot meals directly at the factory kitchen.",

		CtaBannerTitle:      "SAFE, STABLE CATERING PARTNER – READY WITH AUDIT PORTFOLIO",
		CtaBannerBtnZalo:    "Get Quotation on Zalo",
		CtaBannerBtnCall:    "Call Hotline",

		NotFoundTitle:       "Page Not Found (404) | Agape Food",
		NotFoundBadge:       "Error 404",
		NotFoundSubtitle:    "Page Not Found",
		NotFoundDesc:        "The link you accessed is currently unavailable or has been changed. Please go back to the Home page.",
		NotFoundBtn:         "Back to Home",
	}
}
