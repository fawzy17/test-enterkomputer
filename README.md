# Test Backend Programmer - EnterKomputer

Pada test ini saya menggunakan:

1. golang versi 1.22.5

2. mysql

3. make

4. migrate

**note: 

- jika belum menginstall make & migrate skip ke poin 8
- dokumentasi API menggunakan POSTMAN nama file "test-enterkomputer.postman_collection.json"

## Installation

1. pastikan semua sudah terinstall

2. compile code dengan menjalankan perintah
```bash
make build
```

3. buat database dengan nama "test-enterkomputer"

4. untuk membersihkan tabel dalam database bisa menggunakan perintah
```bash
make migrate-down
```

5. untuk import kembali tabel bisa menggunakan perintah
```bash
make migrate-up
```

6. masukkan data products secara manual
```bash
INSERT INTO `products` (`id`, `name`, `category`, `variant`, `price`, `createdAt`) VALUES
(1, 'Jeruk', 'Minuman', 'Dingin', 12000, '2024-08-05 14:21:02'),
(2, 'Jeruk', 'Minuman', 'Panas', 10000, '2024-08-05 14:21:02'),
(3, 'Teh', 'Minuman', 'Manis', 8000, '2024-08-05 14:21:02'),
(4, 'Teh', 'Minuman', 'Tawar', 5000, '2024-08-05 14:21:02'),
(5, 'Kopi', 'Minuman', 'Dingin', 8000, '2024-08-05 14:21:02'),
(6, 'Kopi', 'Minuman', 'Panas', 6000, '2024-08-05 14:21:02'),
(7, 'Mie', 'Makanan', 'Goreng', 15000, '2024-08-05 14:21:02'),
(8, 'Mie', 'Minuman', 'Kuah', 15000, '2024-08-05 14:21:02'),
(9, 'Nasi', 'Makanan', 'Goreng', 15000, '2024-08-05 14:21:02'),
(10, 'Nasi Goreng + Jeruk Dingin', 'Promo', NULL, 23000, '2024-08-05 14:21:02'),
(11, 'Extra Es Batu', 'Minuman', NULL, 2000, '2024-08-05 14:21:02');
```

7. untuk menjalankan project jalankan perintah
```bash
make run
```

## Installation without make & migrate
8. buat database dengan nama "test-enterkomputer"
9. import table dari file "data.sql" secara manual, sudah include data products dan orders
10. jalankan perintah
```bash
go run cmd/main.go
```