/*
 * @Author: Mr Bian
 * @Date: 2018-08-20 10:13:17
 * @LastEditors: Mr Bian
 * @LastEditTime: 2018-08-20 19:55:26
 * @Description: Judgement on google server, inaccessible.
 * @version:
 */

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	//"golang.org/x/tour/reader"
)

// var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (my_reader MyReader) Read(b []byte) (int, error) {
	capacity := len(b)
	/*for i:=0;i< capacity;{
	b[i] = 'A'
	}
	*/
	src := strings.Repeat("AA", capacity)
	fmt.Println(copy(b, src))
	return capacity, nil
}

func read_test() {
	//reader.Validate(MyReader{})

	b := make([]byte, 1024, 2048)
	src := strings.Repeat("A", 3)
	copy(b, src)
	fmt.Println(src)
	fmt.Println(b)
}

type rot13Reader struct {
	r io.Reader
}

type rot13ReaderError string

/*
 func (rot_error rot13ReaderError) Error() string {
	 return fmt.Sprintf("non-alphabetic included in the string \" %v \"", string(rot_error))
 }
*/

func rot13(PtrByte *byte) {
	switch {
	case 'A' <= *PtrByte && *PtrByte <= 'M':
		*PtrByte += 13
	case 'M' < *PtrByte && *PtrByte <= 'Z':
		*PtrByte -= 13
	case 'a' <= *PtrByte && *PtrByte <= 'm':
		*PtrByte += 13
	case 'm' < *PtrByte && *PtrByte <= 'z':
		*PtrByte -= 13
	}
}

func (rot rot13Reader) Read(p []byte) (int, error) {

	length, err := rot.r.Read(p)

	for i := 0; i < length; i++ {
		rot13(&p[i])
	}

	return length, err
}

func TestRot13Reader(s string) {
	cypher := strings.NewReader(s)
	rot := rot13Reader{cypher} // construct a anonymous instance of struct
	io.Copy(os.Stdout, &rot)
	fmt.Println()

	//fmt.Println(rot.r)
	fmt.Printf("The cypher is : ")
	io.Copy(os.Stdout, cypher) // ???? Why is it empty?????
	fmt.Println()

	fmt.Printf("The information is :")
	if _, err := io.Copy(os.Stdout, &rot); err != nil {
		fmt.Print(err.Error())
	}
	fmt.Println()
}
