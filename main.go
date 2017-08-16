package main
import (
        "os"
        "fmt"
        "io"
)
func main() {
  w, err := CopyFile("test.txt", "test-zhou.txt")  
  if err != nil {  
      fmt.Println(err.Error())  
  }  
  fmt.Println(w)  
}
func CopyFile(src, des string) (w int64, err error) {  
  srcFile, err := os.Open(src)  
  if err != nil {  
      fmt.Println(err)  
  }  
  defer srcFile.Close()  

  desFile, err := os.Create(des)  
  if err != nil {  
      fmt.Println(err)  
  }  
  defer desFile.Close()  

  return io.Copy(desFile, srcFile)  
}  
func readTxt(){
  userFile := "test.txt"
  fin,err := os.Open(userFile)
  defer fin.Close()
  if err != nil {
          fmt.Println(userFile,err)
          return
  }
  buf := make([]byte, 1024)
  fmt.Println(buf,"buf")
  for{
          n, _ := fin.Read(buf)
          if 0 == n { break }
          os.Stdout.Write(buf[:n])
  }
}
func createTxt(){
  userFile := "test.txt"
  fout,err := os.Create(userFile)
  defer fout.Close()
  if err != nil {
          fmt.Println(userFile,err)
          return
  }
  for i:= 0;i<10;i++ {
          fout.WriteString("Just a test!\r\n")
          fout.Write([]byte("Just a test!\r\n"))
  }
}