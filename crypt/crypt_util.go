package crypt
// Cryptography utilities

import (
   "crypto/aes"
   "crypto/cipher"
   "crypto/rand"
   "encoding/base64"
   "errors"
   "io"
)

func Encrypt(key []byte, message string) (encmess string, err error) {
    plainText := []byte(message)

    block, err := aes.NewCipher(key)
    if err != nil {
        return
    }

    cipherText := make([]byte, aes.BlockSize+len(plainText))
    iv := cipherText[:aes.BlockSize]
    if _, err = io.ReadFull(rand.Reader, iv); err != nil {
        return
    }

    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

    encmess = base64.URLEncoding.EncodeToString(cipherText)
    return
}

func Decrypt(key []byte, securemess string) (decodemess string, err error) {
   cipherText, err := base64.URLEncoding.DecodeString(securemess)
   if err != nil {
       return
   }

   block, err := aes.NewCipher(key)
   if err != nil {
      return
   }

   if len(cipherText) < aes.BlockSize {
      err = errors.New("Ciphertext block size is too short!")
      return
   }

   iv := cipherText[:aes.BlockSize]
   cipherText = cipherText[aes.BlockSize:]

   stream := cipher.NewCFBDecrypter(block, iv)
   stream.XORKeyStream(cipherText, cipherText)

   decodemess = string(cipherText)
   return
}

