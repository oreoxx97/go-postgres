# ใช้ official Go image เวอร์ชัน alpine (ขนาดเล็ก)
FROM golang:1.20-alpine

# ตั้ง working directory ใน container
WORKDIR /app

# คัดลอก go.mod และ go.sum ก่อน เพื่อ cache dependencies
COPY go.mod go.sum ./

# ดาวน์โหลด dependencies
RUN go mod download

# คัดลอก source code ทั้งหมดเข้า container
COPY . .

# build binary ออกมาเป็นไฟล์ชื่อ app (เปลี่ยนได้ตามชอบ)
RUN go build -o app .

# กำหนดพอร์ตที่แอปจะรัน (ถ้ามี)
EXPOSE 8080

# คำสั่งรันแอปตอน container start
CMD ["./app"]
