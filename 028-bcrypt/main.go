package main

import "bcrypt"

func main() {
	pass := `myPassword`
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.Mincost)
}