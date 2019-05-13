# Ijah Store
IjahStore REST API to serve client app for Ijah Store Management

# Feature
Stock Management untuk Ijah Store meliputi:

1. **Pengelolaan Stock Item** (CRUD)
    
   Untuk mendapatkan keseluruhan item yang ada
   
    ```localhost:7090/api/v1/item```

2. **Info Catatan Jumlah Barang**
    
    Jumlah stok barang saat ini diakses melalui
    
    ```localhost:7090/api/v1/stock```
    
3. **Info Catatan Barang Masuk**

    Catatan keseluruhan barang masuk ke gudang dapat diakses melalui
    
    ```localhost:7090/api/v1/entry```
    
4. **Info Catatan Barang Keluar**

    Catatan keseluruhan barang keluar dari gudang dapat diakses melalui
    
    ```localhost:7090/api/v1/outcome```
    
5. **Pembuatan Order**

    Pembuatan order dilakukan karena setiap barang yg keluar dari gudang harus mempunyai order id yg valid.
    Untuk Pembuatan order melalui
    
    ```localhost:7090/api/v1/order```
    
6. **Laporan Nilai Barang**
    
    Untuk mendapatkan laporan nilai barang melalui
    
    ```localhost:7090/api/v1/report-value```
    
7. **Laporan Penjualan Barang**
    
    Untuk mendapatkan laporan penjualan barang melalui
    
    ```localhost:7090/api/v1/report-sales```

Untuk dokumentasi yang lebih lengkap bisa mengakses dokumentasi API di bawah

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
