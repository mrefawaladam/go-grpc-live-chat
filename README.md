# gRPC Live Chat

Proyek ini adalah aplikasi chat real-time yang dibangun menggunakan gRPC. Aplikasi ini memungkinkan pengguna untuk mengirim dan menerima pesan secara langsung.

## Fitur

- Pengguna dapat bergabung dengan chat menggunakan username.
- Pesan yang dikirim oleh satu pengguna akan disiarkan ke semua pengguna lainnya.
- Menggunakan gRPC untuk komunikasi antara server dan klien.

## Struktur Proyek

grpc-live-chat/
├── client/
│ └── main.go # Kode untuk klien chat
├── proto/
│ └── chat.proto # Definisi protokol gRPC
└── server/
└── main.go # Kode untuk server chat

## Persyaratan

- Go (versi 1.16 atau lebih baru)
- gRPC dan Protobuf

## Instalasi

1. Clone repositori ini:

   ```bash
   git clone https://github.com/username/grpc-live-chat.git
   cd grpc-live-chat
   ```

2. Instal dependensi gRPC:

   ```bash
   go get google.golang.org/grpc
   go get google.golang.org/protobuf
   ```

3. Generate kode dari file `.proto`:
   ```bash
   protoc --go_out=. --go-grpc_out=. proto/chat.proto
   ```

## Menjalankan Server

Untuk menjalankan server, navigasikan ke direktori `server` dan jalankan perintah berikut:

```bash
go run main.go
```

Server akan berjalan di `localhost:50051`.

## Menjalankan Klien

Untuk menjalankan klien, navigasikan ke direktori `client` dan jalankan perintah berikut:

```bash
go run main.go
```

Ikuti instruksi untuk memasukkan username dan mulai mengirim pesan.

## Kontribusi

Jika Anda ingin berkontribusi pada proyek ini, silakan buat pull request atau buka isu.

## Lisensi

Proyek ini dilisensikan di bawah MIT License. Lihat file LICENSE untuk detail lebih lanjut.
