# Agape Food - Landing Page

Dự án Landing Page chính thức của **Agape Food**, được tối ưu hiệu năng cao, chuẩn SEO, responsive đầy đủ và tích hợp thiết kế giao diện sáng (Premium Light/Fresh Theme) hiện đại.

Dự án được viết bằng ngôn ngữ **Go**, sử dụng framework **Fiber** kết hợp cùng công nghệ render template **Templ** và thư viện icon cao cấp **Phosphor Icons**.

---

## 🛠️ Công Nghệ Sử Dụng

- **Backend**: [Go (Golang)](https://golang.org/) & [Fiber v2](https://gofiber.io/) - Web framework siêu nhanh, nhẹ và hiệu năng cao.
- **Templates**: [Templ](https://templ.guide/) - Bộ gõ template kiểu tĩnh (Type-safe template) cho Go, hỗ trợ biên dịch nhanh.
- **Styling**: Vanilla CSS - Tối ưu tốc độ tải trang, tuân thủ hệ thống Design System linh hoạt và chuyên nghiệp.
- **Icons**: [Phosphor Icons](https://phosphoricons.com/) - Sử dụng font CDN gọn nhẹ với hai định dạng `bold` và `fill`.
- **Assets CDN**: Tích hợp module `github.com/dev-networldasia/dspgos/gos/templates` (`AssetURL`) để render chính xác host static assets.

---

## 📂 Cấu Trúc Thư Mục

```text
├── Makefile             # Quản lý các lệnh build/run dự án nhanh
├── main.go              # Điểm khởi chạy ứng dụng (routes, middlewares, 404 handler)
├── go.mod               # Khai báo Go modules và dependencies
├── .env                 # Cấu hình biến môi trường cục bộ (không commit)
├── .env.sample          # File môi trường mẫu
├── public/              # Chứa tài nguyên tĩnh (static files)
│   ├── css/
│   │   └── style.css    # CSS Design System tùy biến của ứng dụng
│   └── images/          # Logo, icon (SVG) và hình ảnh nội bộ
└── views/               # Nguồn mã nguồn template (.templ)
    ├── layouts/
    │   └── layout.templ # Master Layout (Header, Footer, Navigation, CDNs)
    ├── index.templ      # Trang Landing Page chính (Hero, Services, Audit, Menu, FAQ)
    └── 404.templ        # Giao diện trang lỗi 404 Not Found
```

---

## 🚀 Hướng Dẫn Cài Đặt & Chạy Dự Án

### 1. Chuẩn Bị Môi Trường
Yêu cầu hệ thống đã cài đặt sẵn **Go** (phiên bản >= 1.22) và CLI **Templ**.

Nếu chưa cài đặt CLI `templ`, bạn có thể cài bằng lệnh sau:
```bash
go install github.com/a-h/templ/cmd/templ@latest
```

### 2. Thiết Lập File Môi Trường (`.env`)
Sao chép cấu hình từ file mẫu `.env.sample`:
```bash
cp .env.sample .env
```
Mở file `.env` và thiết lập biến `DOMAIN` phù hợp với local của bạn (Ví dụ: `DOMAIN=http://localhost:8080`) để đảm bảo các static assets được load chính xác thông qua `templates.AssetURL`.

### 3. Tải Các Thư Viện Phụ Thuộc
Tải và làm sạch dependencies:
```bash
go mod tidy
```

### 4. Khởi Chạy Ứng Dụng
Sử dụng Makefile để tự động sinh mã template và chạy server:
```bash
make run
```
Sau đó truy cập ứng dụng tại địa chỉ: [http://localhost:8080](http://localhost:8080).

---

## ⚙️ Các Lệnh Makefile Hỗ Trợ

Bạn có thể chạy `make help` hoặc xem chi tiết các lệnh dưới đây:

- `make gen`: Biên dịch/Sinh mã cho toàn bộ các file `.templ` thành mã nguồn Go (`*_templ.go`).
- `make run`: Tự động biên dịch template và khởi chạy server Fiber ở chế độ Hot-reload cục bộ.
- `make build`: Biên dịch mã nguồn dự án thành file chạy Production binary (`landingpage`).
- `make clean`: Dọn dẹp các file binary đã build và toàn bộ các file template Go tự sinh (`*_templ.go`).

---

## 📈 Tối Ưu SEO & Tải Trang

- **SEO Best Practices**: Thiết lập đầy đủ cấu trúc thẻ meta (Title, Description, Author, OpenGraph) trong Master Layout.
- **Semantic HTML**: Sử dụng cấu trúc HTML5 chuẩn (`nav`, `section`, `footer`, `aside`) đảm bảo khả năng cào dữ liệu tốt của Google Bot.
- **Performance**:
  - Tự động nén dữ liệu truyền tải thông qua middleware **Compress** (Gzip/Brotli).
  - Tích hợp asset caching tốt và cấu trúc stylesheet ngoài.
  - Sử dụng lightweight SVG/Font CDN để tăng tốc độ load trang đầu tiên (First Contentful Paint).
