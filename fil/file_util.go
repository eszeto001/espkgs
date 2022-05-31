package fil

import (
  "os"
  "log"
)

func NewReadFileHandle(fileName string) *os.File {
   rfh, err := os.Open(fileName)
   if err != nil {
      log.Printf("Cannot read file '%s'\n", fileName)
      log.Fatal(err)
   }
   return rfh
}

func NewWriteFileHandle(fileName string) *os.File {
    wfh, err := os.Create(fileName)
    if err != nil {
       log.Printf("Cannot write file '%s'\n", fileName)
       log.Fatal(err)
    }
    return wfh
}

func FileExist(fname string) bool {
   if _, err := os.Stat(fname); err == nil {
      return true
   }
   return false
}


