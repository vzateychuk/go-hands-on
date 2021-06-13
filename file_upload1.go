package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) {
	var uploadFormTmpl = []byte(`
		<html>
			<body>
				<form action="/upload" method="post" enctype="multipart/form-data">
					Image: <input type="file" name="my_file">
					<input type="submit" value="Upload">
				</form>
			<body>
		</html>
		`)

	w.Write(uploadFormTmpl)
}

func uploadPage(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(5 * 1024 * 1024)      // получаем 5 Mb загружаемого файла
	file, header, err := r.FormFile("my_file") // получаем файл header-ы
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() // обязательно закрыть файл чтобы не было утечки ресурсов

	fmt.Fprintf(w, "header.Filename %v\n", header.Filename)
	fmt.Fprintf(w, "header.Header %#v\n", header.Header)

	hasher := md5.New()
	io.Copy(hasher, file) // копируем в hasher содержимое файлов чтобы посчитать md5
	fmt.Fprintf(w, "md5 %x\n", hasher.Sum(nil))
}

func main() {
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/upload", uploadPage)
	fmt.Println("Server starting")
	http.ListenAndServe(":8080", nil)
}
