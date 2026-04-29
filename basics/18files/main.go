package main

import (
	"bufio"
	"fmt"
	// "io"
	"os"
)

// func main() {
// 	fmt.Println("Welcome to the files in golang")

// 	content := "This is the content of the files"

// 	// create a file
// 	file,err := os.Create("./mygofile.txt");

// 	if err != nil {
// 		panic(err)
// 	}

// 	// write to the file with io package

// 	length, err := io.WriteString(file, content)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Written %d characters\n", length)

// 	defer file.Close();  // defer will execute at the end of the main function, it will close the file after all operations are done
// 	// using defer is a good practice to ensure that resources are released properly, even if an error occurs

// 	fmt.Println("File created successfully")

// 	// read the file
// 	fmt.Println("reading the data from file")
// 	readFile("./mygofile.txt")  // we will get the content as byte slice, we can convert it to string to get the actual content of the file

// 	fmt.Println(string(readFile("./mygofile.txt")))  // converting byte slice to string to get the actual content of the file

// }

// func readFile(filename string) []byte {
// 	// read the file

// 	content, err := os.ReadFile(filename)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(content)
// 	return content;
// }

func main() {

	f, err := openFile("mygofile.txt")
	if err != nil {
		panic(err)
	}


	// delete a file 

		defer func() {
		dlerr := os.Remove("destination.txt")
		if dlerr != nil {
			panic(dlerr)
		}

		fmt.Println("File deleted successfully")
	}()

	// read file content
	// to read the file content this way we need to create a buffer to read the file content, we can adjust the buffer size as needed
	// buf := make([]byte, 1024)

	// n, err := f.Read(buf) // here we get length of the content read and an error if any
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("Read %d characters\n", n)

	// // now we can print the content of file which is in buf

	// fmt.Println("File content:", string(buf)) // converting byte slice to string and slicing it to get only the content read

	// // easier way to read the file content is to use os.ReadFile which reads the entire file content and returns it as a byte slice, we can convert it to string to get the actual content of the file
	// content, err := os.ReadFile("mygofile.txt")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("File content using os.ReadFile:", string(content)) // converting byte slice to string to get the actual content of the file

	defer f.Close() // close the file after all operations are done, using defer to ensure that it is closed even if an error occurs

	// read folders
	// readFolder("../") // read the current folder and print the file names in it

	// write to a file 
	f.WriteString("this is update")

	// write to a file using bytes

	// content := []byte(" this is the udpated content of the file")
	// f.Write(content)



	// read and write to another file (streaming fashion)

	// taking f as source file 

	destFile , err := os.Create("destination.txt")

	if err != nil {
		panic(err)
	}

	defer destFile.Close()


	// we will use inbuilt package called bufio which provides buffered I/O operations

	reader := bufio.NewReader(f)
    writer := bufio.NewWriter(destFile)

	for {
		// read bytes from source file

		b , err := reader.ReadByte() // read a single byte from the source file

		if err != nil {
			if err.Error() != "EOF" { // if the error is not EOF, then we panic, otherwise we break the loop
				panic(err)
			}
			break
		}
		e := writer.WriteByte(b) 
		if e != nil {
			panic(e)
		}
	}

writer.Flush() // flush the writer to ensure that all the buffered data is written to the destination file

fmt.Println("File copied successfully")


}

func openFile(fpath string) (*os.File, error) {
	// open a file with os.Open, it returns a file pointer and an error

	// f, err := os.Open(fpath) // os.open is for read only
	f, err := os.OpenFile(fpath, os.O_RDWR, 0644) // os.OpenFile is for read and write, we need to specify the flags and permissions

	if err != nil {
		return nil, err
	}

	fileInfo, err := f.Stat()
	if err != nil {
		return nil, err
	}

	fmt.Println("File Info:", fileInfo.Name(), fileInfo.Size(), fileInfo.Mode(), fileInfo.ModTime(), fileInfo.IsDir(), fileInfo.Sys())

	return f, nil
}

// read folders

func readFolder(folderPath string) {
	dir, err := os.Open(folderPath)
	if err != nil {
		panic(err)
	}
	defer dir.Close()

	fileInfo, err := dir.Readdir(-1) // 0 < n <= 1000, -1 to read all files in the folder
	if err != nil {
		panic(err)
	}

	for _, file := range fileInfo {
		fmt.Println(file.Name(), file.IsDir())
	}
}


