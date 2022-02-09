"# hom-vu-golang-fresher-camp" 
+ Không nên chứa file upload vào ngay chính bên trong service mà nên chứa trong cloud bời vì:
  - Tăng dung lượng của service dẫn đến giảm performance
  - Có thể sử dụng chung trong nhiều service
  - Tính bảo mật cao
  - Dễ quản lý, backup dữ liệu
+ Không nên chứa ảnh vào binary DB bời vì:
  - Tăng dung lượng DB
  - Tốc độ query giảm
