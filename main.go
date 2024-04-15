package main

import (
	"fmt"
	m "homework14/models"

	validation "github.com/ruziba3vich/hmwnmb14lst/Validation"
)

func main() {
	status := 1
	for status == 1 {
		var user m.User
		fmt.Print("username kiriting : ")
		fmt.Scan(&user.Username)
		fmt.Print("email kiriting : ")
		fmt.Scan(&user.Email)
		fmt.Print("age kiriting : ")
		fmt.Scan(&user.Age)
		if ValidateUser(user) {
			fmt.Println("Validatsiya muvaffaqayatli amalga oshirildi ")
			status = 2
		} else {
			fmt.Println("Validatsiyada hatolik, ma'lumotni qayta tekshiring")
		}
		if status == 2 {
			fmt.Print("dasturdan yana foydalanasizmi ? 1 -> Ha / 2 -> Yo'q : ")
			fmt.Scan(&status)
		}
	}
	fmt.Println("Dasturdan foydalanganingiz uchun raxmat !")
}

func ValidateUser(user m.User) bool {
	usernameError := validation.ValidateUsername(user.Username, map[string]bool{})
	passwordError := validation.ValidateEmail(user.Email)
	ageError := validation.ValidateAge(int(user.Age))

	return usernameError == nil && passwordError == nil && ageError == nil
}

