package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"study/util"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       // 解析url传递的参数，对于POST则解析响应包的主体(request body) //注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	fmt.Println(r.Form) // 这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello junxi!") // 这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	fmt.Println('我')
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("web/form/login.html")
		t.Execute(w, token)
		// t1, _ := template.New("foo").Parse(`{{define "T"}}Hello, {{.}}!{{end}}`)
		// t1.ExecuteTemplate(w, "T", "<script>alert('you have been pwned')</script>")
		// t1.ExecuteTemplate(w, "T", template.HTML("<script>alert('you have been pwned')</script>"))
	} else {
		// 请求的是登录数据， 那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		fmt.Println(token)
		if token != "" {
			// 验证token的合法性
		} else {
			// 不存在token报错
		}
		if m, _ := regexp.MatchString("^[a-zA-Z]+$", r.Form.Get("username")); !m {
			fmt.Fprintln(w, "username check:", m)
		} else {
			fmt.Fprintln(w, "username: ", r.Form.Get("username"))
		}
		if m, _ := regexp.MatchString("^[0-9]+$", r.Form.Get("password")); !m {
			fmt.Fprintln(w, "password check:", m)
		} else {
			fmt.Fprintln(w, "password: ", r.Form.Get("password")) // 这个写入到w的是输出到客户端的
		}
		if m, _ := regexp.MatchString("^[\\x{4e00}-\\x{9fa5}]+$", r.Form.Get("nickname")); !m {
			fmt.Fprintln(w, "nickname check:", m)
		} else {
			fmt.Fprintln(w, "nickname: ", r.Form.Get("nickname"))
		}
		if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Fprintln(w, "email check:", m)
		} else {
			fmt.Fprintln(w, "email: ", r.Form.Get("email"))
		}
		if m, _ := regexp.MatchString(`^(1[356789]\d{9})$`, r.Form.Get("phone")); !m {
			fmt.Fprintln(w, "phone check:", m)
		} else {
			fmt.Fprintln(w, "phone: ", r.Form.Get("phone"))
		}
		fruits := []string{"apple", "pear", "banane"}
		for _, v := range fruits {
			if v == r.Form.Get("fruit") {
				fmt.Fprintln(w, "fruit: ", r.Form.Get("fruit"))
			} else {
				fmt.Fprintln(w, "fruit check:", false)
			}
		}
		genders := []string{"1", "2"}
		for _, v := range genders {
			if v == r.Form.Get("gender") {
				fmt.Fprintln(w, "gender: ", r.Form.Get("gender"))
			} else {
				fmt.Fprintln(w, "gender check:", false)
			}
		}
		slice := []string{"football", "basketball", "tennis"}
		newSlice := make([]interface{}, len(slice))
		for i, v := range slice {
			newSlice[i] = v
		}
		interest := make([]interface{}, len(r.Form["interest"]))
		for i, v := range r.Form["interest"] {
			interest[i] = v
		}
		a := util.Slice_diff(interest, newSlice)
		if a == nil {
			fmt.Fprintln(w, "interest: ", r.Form.Get("interest"))
		} else {
			fmt.Fprintln(w, "interest check:", false)
		}
		t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
		fmt.Printf("Go launched at %s\n", t.Local())
		if m, _ := regexp.MatchString(`^(\d{17})([0-9]|X)$`, r.Form.Get("id")); !m {
			fmt.Fprintln(w, "id: ", r.Form.Get("id"))
		} else {
			fmt.Fprintln(w, "id check:", false)
		}
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) // 获取请求的方法
	if r.Method == "GET" {
		crutime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(crutime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))
		t, _ := template.ParseFiles("web/form/upload.html")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, hander, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		fmt.Fprintf(w, "%v", hander.Header)
		f, err := os.OpenFile("web/form/" + hander.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		_, err = io.Copy(f, file)
		if err != nil {
			fmt.Fprintf(w, "文件上传失败")
			return
		}
		// io.WriteString(w, hander.Filename)
	}
}

func main() {
	http.HandleFunc("/", sayhello) // 设置访问的路由
	http.HandleFunc("/login", login)
	http.HandleFunc("/upload", upload)
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
