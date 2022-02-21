Khi nào cần tạo các cột số đếm ngay trên table dữ liệu (VD: liked_count trên restaurants)?

Ta cần thêm cột đếm ngay trên table thường xuyên được đọc, trong đó thuộc tính đếm được cập nhật liên tục, khi đó nếu cột số đếm mà phải tổng hợp từ bảng khác thì tốn thêm thời gian, cũng như database phải tính toán tổng hợp lại số liệu làm cho hệ thống DB phải chịu thêm tải.
Ví dụ cột liked_count nếu được lưu trên bảng restaurants, do bảng restaurants là bảng được dọc dữ liệu thường xuyên, việc lưu cột liked_count sẽ khiến việc cập nhật dữ liệu dễ dàng mà ko phải gọi thêm API để tính toán và tổng hợp từ bảng restaurant_likes.
