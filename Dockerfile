# ใช้ Go image เป็น base image
FROM golang:1.23-alpine

# ตั้งค่า working directory
WORKDIR /app

# คัดลอก go.mod และ go.sum เพื่อติดตั้ง dependencies
COPY go.mod go.sum ./

# ติดตั้ง dependencies
RUN go mod download

# คัดลอก source code ทั้งหมดเข้าไปใน container
COPY . .

# สร้าง executable file ของ Go API
RUN go build -o main .

# ระบุ port ที่จะ expose ใน container
EXPOSE 8000

# คำสั่งที่ใช้รัน API
CMD ["./main"]
