## Question number 2

### Configuration file
Di service `gateway` dan `service-1` ada file (.env) yang berisi konfigurasi dengan runtime environment,
silahkan disesuaikan.

### Cara install
1. Jika ingin menyesuaikan konfigurasi di setiap service (misal: ganti port), silahkan mengedit file (.env) pada setiap services
2. Build dan jalan: `docker-compose up -d`
3. Silahkan akses endpoint yang telah terbuat: `http://{address}:{port}/search?title=Spiderman&page=1`

### Daftar Endpoint
```
GET /search                 - pencarian film (berdasarkan http://www.omdbapi.com/)
GET /detail/{imdbID}        - akses detail film dengan IMDB ID
```
Saat melakukan pencarian film, otomatis akan melakukan logging dengan menyimpan result ke database MySQL.