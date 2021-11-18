package main

import ("fmt" //доступ к базовым функциям
"net/http") //отображение данных на странице и отслеживание пользователя

type User struct { //структура юсера
  name string
  age uint16
  money int16
  avg_gtades, happiness float64
}

func (u *User) setNewName(newName string) { //функция замены имени
  u.name = newName
}

func (u User) getALLInfo() string { // функция выводит информацию о юзере
  return fmt.Sprintf("User name is: %s He is %d and he has money " +
    "equal: %d", u.name, u.age, u.money)
}

func home_page(w http.ResponseWriter, r *http.Request)  { // слушаем url
  bob := User{"Bob", 25, -50, 4.2, 0.8}
  bob.setNewName("Alex")
  fmt.Fprintf(w, bob.getALLInfo()) //форматированная строка
}

func contacts_page(w http.ResponseWriter, r *http.Request)  {

}

func handleRequest()  {
  http.HandleFunc("/", home_page) //принимает 2 параметра и отследить переход по url и использовать метод для отображения информации
  http.HandleFunc("/contacts/", contacts_page) //принимает 2 параметра и отследить переход по url и использовать метод для отображения информации
  http.ListenAndServe(":8080", nil) //принимает 2 параметра. порт и настройки сервера
}
func main() {
  //var bob User = ....
  //bob := User {name: "Bob", age: 25, money: -50, avg_gtades: 4.2, happiness: 0.8}
  handleRequest()
}
