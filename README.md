## What's in the box?

Di repository ini, gue belajar backend web development Golang tanpa framework. Cukup pake built-in package yang namanya `net/http`.

## Server

- Dalam package `net/http`, `Server` adalah sebuah struct yang digunakan untuk representasi web server. Intinya, kalo mau bikin web pake Golang ya harus bikin `Server`.
- Ketika bikin `Server`, kita tentukan host dan port untuk menjalankan webnya.
- Setelah deklarasi `Server`, jalankan dengan menggunakan function `ListenAndServe()`. Function `ListenAndServe()` bisa aja mengembalikan error. Jadi, lebih baik masukkan kedalam sebuah variable lalu lakukan `panic()` apabila ada error.

Contoh:

```go
import (
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	// deklarasi Server
	server := http.Server{
		Addr: "127.0.0.1:8000",
	}

	// jalankan Server dengan ListenAndServe()
	err := server.ListenAndServe()

    // panic apabila eksekusi server error
	if err != nil {
		panic(err)
	}
}
```

Dari contoh kode diatas, kita pake unit testing dengan function `TestServer(t *testing.T)`, dan jangan lupa `import "testing"`. Lalu jalankan testingnya dengan command `go test nama_file.go -v`. Flag `-v` artinya verbose yang berguna untuk menampilkan seluruh output yang ada ketika proses testing dijalankan.

> **BACA LEBIH LANJUT TENTANG PACKAGE `testing`: [https://pkg.go.dev/testing](https://pkg.go.dev/testing)**