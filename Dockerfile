# Menggunakan image golang:1.x sebagai base image
FROM golang:1.21.1

# Menentukan direktori kerja di dalam container
WORKDIR /app

# Menyalin seluruh file aplikasi Go ke dalam direktori kerja di dalam container
COPY . .

# Menjalankan perintah build aplikasi Go
RUN go build -o main .

# Mengungkapkan port 8080 yang akan digunakan oleh aplikasi
EXPOSE 8080

# Perintah yang akan dijalankan ketika container dijalankan
CMD ["./main"]