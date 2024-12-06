# Memahami Golang Interfaces

fitur interface di Golang mungkin sebuah fitur yang paling banyak disalahgunakan karena ga sedikit coder
golang berasal dari bahasa OOP seperti C++, Java, PHP, Typescript

dimana di dalam bahasa tersebut mereka terbiasa membuat API yang dimulai dari abstraksi (abstract) terlebih dahulu baru 
aktual implementasinya (concrete)

itu tentu bisa bekerja juga di Golang, namun ada beberapa kekurangan dengan pendekatan itu, salah satunya adalah bikin repot
coder pengguna API (consumer) untuk mengetes kodenya yang memakai package API (dependency/library) tersebut

keribetan muncul ketika membuat `mock` untuk interface (abstraks) yang dedefine oleh pembuat API, dimana sangat mungkin banyak method
yang tidak dipakai oleh pengguna API

menyebabkan pengguna API harus menulis mock untuk semua method daripada interface yang diekspos tersebut, padahal kodenya hanya
memakai beberapa method saja

sebagai ilustrasi dan contoh: gue pingin buat fitur ngitung kata (wordcount), trus ada kode dependency/library yang dibuat "temen" seperti ini:

```go
type TokenStorage interface {
	AddToken(name string)
	GetTokenCount(name string) int
	UpdateTokenName(old, new string)
	UpdateTokenCount(name string, count int)
	RemoveToken(name string)
	ResetToken(name string)
	Size() int
}
```

kalau kita pakai interface itu langsung seperti ini:

```go
type TokenService struct {
	dep database.TokenStorage
}

func NewTokenService(dep database.TokenStorage) *TokenService {
	return &TokenService{
		dep: dep,
	}
}

func (s TokenService) Tokenize(tokens []string) {
	for _, v := range tokens {
		s.dep.AddToken(v)
	}
}

func (s TokenService) GetTokenCount(name string) int {
	return s.dep.GetTokenCount(name)
}
```

nanti dalam menulis test dan ingin membuat mock ke external dependencynya kita harus mengimplementasikan setiap method dari `TokenStorage`

```go
type mockClient struct{}

func (c mockClient) AddToken(name string)                    {}
func (c mockClient) GetTokenCount(name string) int           { return 0 }
func (c mockClient) UpdateTokenName(old, new string)         {}
func (c mockClient) UpdateTokenCount(name string, count int) {}
func (c mockClient) RemoveToken(name string)                 {}
func (c mockClient) ResetToken(name string)                  {}
func (c mockClient) Size() int                               { return 0 }

func Test_TokenService(t *testing.T) {
	ts := slightlyunhappyconsumer.NewTokenService(mockClient{})
	if ts == nil {
		t.Fail()
	}
}
```

bagaimana kalau kita balik pendekatannya dalam membuat API, kita cukup menulis implementasinya (concrete) saja tanpa repot-repot mikirin interface (abstract) apa yang musti kita expose ke pemakai

kita bebaskan pengguna API untuk menulis interface yang mereka butuhkan saja ke package implementasi API seperti ini:

```go
type TokenPusher interface {
	AddToken(name string)
}

type TokenGetter interface {
	GetTokenCount(name string) int
}

type Dependency interface {
	TokenPusher
	TokenGetter
}
```

dan

```go
type TokenService struct {
	dep Dependency
}

func NewTokenService(dep Dependency) *TokenService {
	return &TokenService{
		dep: dep,
	}
}

func (s TokenService) Tokenize(tokens []string) {
	for _, v := range tokens {
		s.dep.AddToken(v)
	}
}

func (s TokenService) GetTokenCount(name string) int {
	return s.dep.GetTokenCount(name)
}
```

sehingga untuk membuat mock dalam test menjadi lebih mudah:

```go
type mockClient struct{}

func (c mockClient) AddToken(name string)          {}
func (c mockClient) GetTokenCount(name string) int { return 0 }

func Test_TokenService(t *testing.T) {
	ts := happyconsumer.NewTokenService(mockClient{})
	if ts == nil {
		t.Fail()
	}
}
```

apakah "implementasi" dari interface di Golang lebih bagus dari C++ dan/atau Java? gatau, gue sendiri masih belum paham banget motivasi dibalik kenapa golang memilih cara seperti ini