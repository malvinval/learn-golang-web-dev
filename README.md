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

> **Baca lebih lanjut tentang package `testing`: [https://pkg.go.dev/testing](https://pkg.go.dev/testing)**

## Handler

- Sesuai namanya, `Handler` ini bertugas untuk handling (penanganan) request dari client.
- Di Golang, `Handler` ini bentuknya interface. Didalam interface tersebut, ada sebuah function `ServeHTTP(ResponseWriter, *Request)`.
- Karena `Handler` ini bentuknya interface yang didalamnya ada function `ServeHTTP()`, kita harus implementasikan interface tersebut secara manual sesuai dengan *signature* atau *convention* dari interface tersebut (ingat konsep interface dalam OOP). Namun ada cara lain yang lebih gampang, yaitu membuat anonymous function bertipe `http.HandlerFunc` dengan parameter `(w http.ResponseWriter, r *http.Request)`. `w` itu untuk response ke client, sedangkan `r` untuk request dari client. `HandlerFunc` hanyalah sebuah user-defined data type yang sudah mengimplementasikan interface `Handler`.
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


> Silahkan liat kode deklarasi `HandlerFunc()` di `/usr/local/go/src/net/http/server.go` supaya lebih paham apa maksud `HandlerFunc()` ini dibuat. Dibawah ini adalah penjelasan singkat terkait kode tersebut:


```go
// source file: /usr/local/go/src/net/http/server.go

type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, r).
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
```

1. `type HandlerFunc func(ResponseWriter, *Request)`: Golang mendefinisikan sebuah tipe yang disebut `HandlerFunc`.

2. `func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request)`: Golang menambahkan function `ServeHTTP` ke tipe `HandlerFunc`. Ini adalah contoh penerapan metode yang dikenal sebagai ***receiver function*** dalam Go. Fungsi ServeHTTP ini akan dipanggil ketika tipe `HandlerFunc` dijalankan. `f HandlerFunc` adalah penerima (receiver) dari metode ini. Berarti function `ServeHTTP` dapat dipanggil pada instance dari tipe `HandlerFunc`. Dalam metode ini, `f` merujuk pada instance `HandlerFunc` yang memanggilnya.

**Kesimpulan sederhana**: memanggil function `ServeHTTP()` itu dilakukan dengan cara memanggil dan mengimplementasikan function `HandlerFunc`. Karena function `ServeHTTP()` sudah di-*attach* kedalam `HandlerFunc`.

## ServeMux

- `HandlerFunc` tidak mendukung handling banyak URL endpoint. Masa iya web kita cuma bisa handle endpoint `/` doang? Nah, `ServeMux` ini alternatifnya `HandlerFunc`.
- `ServeMux` adalah implementasi Handler yang mendukung multiple endpoint.

> Inget, sebelumnya udah kita bahas bahwa `Handler` adalah sebuah interface. Jadi implementasinya bisa dengan banyak cara, contohnya `HandlerFunc`, dan `ServeMux` ini.

Contoh:

```go
import (
	"fmt"
	"net/http"
	"testing"
)

func TestMux(t *testing.T) {
	mux := http.NewServeMux()

	// implementasi handler menggunakan anonymous function di parameter kedua
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { // mux.HandleFunc
		fmt.Fprint(w, "Ini path /")
	})

	// implementasi handler menggunakan http.HandlerFunc yang dideklarasi terlebih dahulu
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Ini path /about")
	}
	mux.Handle("/about", handler) // mux.Handle

	server := http.Server{
		Addr:    "127.0.0.1:8000",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
```

Dari contoh kode diatas, kita membuat mekanisme handling `ServeMux` dengan 2 cara yang berbeda yaitu dengan function `mux.HandleFunc()` dan `mux.Handle()`. Dengan `mux.HandleFunc()`, kita membuat handlernya secara langsung sebagai anonymous function di parameter kedua. Sedangkan dengan `mux.Handle()`, kita terlebih dahulu deklarasi sebuah handler (bertipe `http.HandlerFunc`). Bebas aja mau make cara yang mana.

## Request

- `Request` adalah sebuah struct dalam Golang yang merepresentasikan sebuah request dari user. Semua informasi terkait request user seperti URL, header, body, method, dan lain-lain dapat kita lihat melalui `Request`.