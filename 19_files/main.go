package main

import (
	"fmt"
	"os"
)

// file paths from cd 19_files

func main() {
	// f, err := os.Open("example.txt")

	// if err != nil {
	// 	panic(err)
	// }

	// fileInfo, err := f.Stat()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("file name", fileInfo.Name())
	// fmt.Println("file isDir", fileInfo.IsDir())
	// fmt.Println("file permission", fileInfo.Mode())

	// read file
	// f, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// buf := make([]byte, 15)

	// d, err := f.Read(buf)
	// if err != nil {
	// 	panic(err)
	// }

	// for i := 0; i < len(buf); i++ {
	// 	println("data", d, string(buf[i]))
	// }

	// read file
	// f, err := os.ReadFile("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(string(f))

	// read folders
	// dir, err := os.Open("../")
	// if err != nil {
	// 	panic(err)
	// }

	// defer dir.Close()

	// fileInfo, err := dir.ReadDir(-1)

	// for _, fi := range fileInfo {
	// 	fmt.Println(fi.Name())
	// }

	// create a file
	// f, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// f.WriteString("hi, GO")
	// f.WriteString("append line")

	// bytes := []byte("Hello, GO")

	// f.Write(bytes)

	// Transfer Data (streaming fashion)
	// srcFile, err := os.Open("example.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer srcFile.Close()

	// destFile, err := os.Create("example2.txt")
	// if err != nil {
	// 	panic(err)
	// }

	// defer destFile.Close()

	// reader := bufio.NewReader(srcFile)
	// writer := bufio.NewWriter(destFile)

	// for {
	// 	b, err := reader.ReadByte()
	// 	if err != nil {
	// 		// end of file error
	// 		if err.Error() != "EOF" {
	// 			panic(err)
	// 		}

	// 		break
	// 	}

	// 	er := writer.WriteByte(b)
	// 	if er != nil {
	// 		panic(er)
	// 	}
	// }

	// writer.Flush()

	// fmt.Println("written to new file")

	// delete file
	err := os.Remove("example2.txt")
	if err != nil {
		panic(err)
	}

	fmt.Println("file deleted")

}
