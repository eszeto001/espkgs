package str

import (
  "log"
  "strconv"
)

func Chop(s string) string {
   return s[:len(s)-1]
}

func ParseFloat(s string) float64 {
   val, err := strconv.ParseFloat(s, 64)
   if err != nil {
      log.Fatal(err)
   }
   return val
}

func ParseInt(s string) int {
    val, err := strconv.Atoi(s)
    if err != nil {
       log.Fatal(err);
    }
    return val
}

