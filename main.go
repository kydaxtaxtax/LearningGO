package main

<<<<<<< HEAD
import ("fmt"
  "net/http"
  "html/template"
)

func index(w http.ResponseWriter, r *http.Request) {
  t, err := template.ParseFiles("templates/index.html", "templates/header.html", "templates/footer.html")

  if err != nil {
    fmt.Fprintf(w, err.Error())
  }
  t.ExecuteTemplate(w, "index", nil)
}

func handleFunc() {
  http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static/"))))
  http.HandleFunc("/", index)
  http.ListenAndServe(":8080", nil)
}

func main() {
  handleFunc()
=======
import ("fmt" //доступ к базовым функциям
  "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
  Name string `json: "name"`
  Age uint16 `json: "age"`
}

func main() {
  db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")
  if err != nil {
    panic(err)
  }

  defer db.Close()

  //Установка данных
  // insert, err := db.Query("INSERT INTO `users` (`name`, `age`) VALUES ('Bob', 35)")
  // if err != nil {
  //   panic(err)
  // }
  // defer insert.Close()

//Выборка данных
res, err := db.Query("SELECT `name`, `age` FROM `users`")
if err != nil {
  panic(err)
}

for res.Next() {
    var user User
    err = res.Scan(&user.Name, &user.Age)
    if err != nil {
      panic(err)
    }

    fmt.Println(fmt.Sprintf("User: %s with age %d", user.Name, user.Age))
  }
>>>>>>> d01660e716c213ac839fbb1df1a092364a3156ec
}
