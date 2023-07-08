# Backend Cho [CsVideoCall app](https://github.com/quangnhat22/cs_video_call_app)

Repository này lưu trữ source code của Backend Server cho `CS Video Call` app.

## Cách cài đặt

1. Tải thư mục `docker`.
2. Điền các biến môi trường phù hợp.
3. Tải `Livekit`. [Hướng dẫn cách tải](https://docs.livekit.io/getting-started/server-setup/)
4. Cấu hình `Livekit` server với hook. [Ví dụ](https://github.com/dinhlockt02/cs_video_call_app_server/blob/main/docker/livekit.yaml)
5. Chạy `Livekit` server.
6. Thực thi `docker-compose up -d` để chạy server.

## Hướng dẫn sử dụng
Định nghĩa API có thể được nhập sử dụng Postman, file nhập được lưu tại `/api-docs/postman.json`.
