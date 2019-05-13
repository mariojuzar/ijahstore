# Ijah Store
IjahStore REST API to serve client app for Ijah Store Management

# Feature
Stock Management untuk Ijah Store meliputi:

- Pengelolaan Stock Item (CRUD)
- Info Catatan Jumlah Barang
- Info Catatan Barang Masuk
- Info Catatan Barang Keluar
- Pembuatan Order
- Laporan Nilai Barang
    

# Requirement
- Go 1.12+

# Setup
- Clone repository ini
- Jika kamu punya GoLand IDE, open project, lalu run `go build main.go`

# Assumption
1. Jika stock item telah di buat maka tidak ada penghapusan data dari stock item
2. Barang keluar hanya bisa valid jika order id yg dimasukkan valid
3. Current stock di tidak bisa di update manual, datanya di dapatkan dari barang masuk dan barang keluar
4. **Tidak semua negatif case di handle dengan baik karena waktu yg sangat singkat, jadi tolong di execute sesuai dengan yang seharusnya**

# Flow
1. Data Stock Item harus diisi terlebih dahulu
2. Setelah ada data stock item, bisa cek current stock
3. Buat Entry Stock Log untuk pencatatan barang masuk ke gudang
4. Barang keluar baru bisa valid jika ada order id yg valid, jadi harus create order terlebih dahulu
5. Buat Outcome Stock Log berdasarkan order id
6. Jika semua proses di atas selesai, maka bisa melihat report

# API Documentation
Available in Postman Doc [here](https://documenter.getpostman.com/view/6895601/S1LyTSjy)
