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

## Handler

- Sesuai namanya, `Handler` ini bertugas untuk handling (penanganan) request dari client.
- Di Golang, `Handler` ini bentuknya interface. Didalam interface tersebut, ada sebuah function `ServeHTTP(ResponseWriter, *Request)`.
- Namun, `ServeHTTP()` itu bisa kita implementasikan dalam bentuk anonymous function bertipe `http.HandlerFunc` dengan parameter `(w http.ResponseWriter, r *http.Request)`. `w` itu untuk response ke client, sedangkan `r` untuk request dari client.
- Kita coba berikan response "Hello World" dengan menggunakan function `Fprint()` karena `fmt.Println()` itu dipake untuk output console.

Contoh:

```go
import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		// w: response untuk client
		// r: request dari client
		fmt.Fprint(w, "Hello world!")
	}

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: handler,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
```