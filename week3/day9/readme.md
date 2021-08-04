# MONGODB

-   digunakan ketika data dinamis.
-   tidak direkomendasikan digunakan untuk data transaksi, data ACID, atau data yang sudah jelas.

## Kaidah CAP

## Consistency

Setiap proses mendapatkan data terbaru

## Availability

selalu tersedia untuk diakses

## Partition

Bisa diakses meski ada kendala jaringan

# Schemaless

-   bentuk data tiap `row` tidak tidak pasti. Sehingga, perlu dihandle manual.
-   tidak mempedulikan tipedata
-   Bisa nested

# Tipe NoSQL DBMS

## Key-Value

-   Key harus integer.
-   Value bebas.
-   Biasa digunakan untuk Cache.

## Column-Store

-   menyimpan data dalam bentuk column.

## Graph

-   Seperti GraphQL (tiap table punya label untuk 1 jenis data. Misal: Profile punya friend list, masing2 friend punya friendlistnya sendiri2, dll)

## Document-based

-   MongoDB
