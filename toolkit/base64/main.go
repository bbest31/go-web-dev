package main

import (
	"encoding/base64"
	"fmt"
)

/*
Base 64 is a nice way to encode data into characters that are acceptable for their storage.
For example cookies only accept alphanumeric characters plus some others. They don't accept double quotations.
So if you wanted to save json in a cookie you need to encode it using base64, store it in the cookie, and then decode it once you want to use it.
*/
func main() {
	s := "Nisi ad culpa veniam et. Officia eiusmod et enim in aliquip dolore velit sit dolore veniam. Duis eu incididunt laboris sunt adipisicing anim. Irure adipisicing ut fugiat deserunt enim proident anim nulla laboris fugiat eu. Excepteur amet labore Lorem veniam voluptate nostrud ullamco tempor exercitation amet quis mollit. Cillum pariatur eu pariatur ipsum qui ut. Duis laboris enim elit deserunt minim sunt velit dolor pariatur enim. Nostrud adipisicing nulla irure eiusmod amet ullamco. Ipsum laboris pariatur amet sunt occaecat sint cupidatat dolor et exercitation reprehenderit ad.Voluptate occaecat sunt consequat laborum et adipisicing irure ut ex esse. Nulla eu non nostrud nostrud quis laborum Lorem. Proident laboris qui excepteur laboris ullamco ut et. Id do deserunt Lorem velit et fugiat ut cupidatat eiusmod nisi ipsum laboris amet. Mollit anim sit consequat officia irure voluptate non enim in minim. Eu qui adipisicing voluptate aute pariatur voluptate sint ipsum elit id irure laboris sit."

	encodeStd := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	s64 := base64.NewEncoding(encodeStd).EncodeToString([]byte(s))
	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s)
	fmt.Println(s64)
}
