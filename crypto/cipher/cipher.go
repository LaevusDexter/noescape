// This file has automatically been generated on Wed Feb 26 02:09:43 +05 2020.
// DO NOT EDIT.
package cipher

import (
	"crypto/cipher"
	_ "unsafe"
)

//go:linkname NewCFBEncrypter crypto/cipher.NewCFBEncrypter
//go:noescape
func NewCFBEncrypter(block cipher.Block, iv []byte) cipher.Stream

//go:linkname streamwriterclose cipher.sub_streamwriterclose
func streamwriterclose(w cipher.StreamWriter) error {
	return w.Close()
}

//go:linkname StreamWriterClose cipher.sub_streamwriterclose
//go:noescape
func StreamWriterClose(w cipher.StreamWriter) error

//go:linkname NewGCMWithNonceSize crypto/cipher.NewGCMWithNonceSize
//go:noescape
func NewGCMWithNonceSize(cipher cipher.Block, size int) (cipher.AEAD, error)

//go:linkname NewCBCDecrypter crypto/cipher.NewCBCDecrypter
//go:noescape
func NewCBCDecrypter(b cipher.Block, iv []byte) cipher.BlockMode

//go:linkname NewCBCEncrypter crypto/cipher.NewCBCEncrypter
//go:noescape
func NewCBCEncrypter(b cipher.Block, iv []byte) cipher.BlockMode

//go:linkname NewCTR crypto/cipher.NewCTR
//go:noescape
func NewCTR(block cipher.Block, iv []byte) cipher.Stream

//go:linkname NewOFB crypto/cipher.NewOFB
//go:noescape
func NewOFB(b cipher.Block, iv []byte) cipher.Stream

//go:linkname streamreaderread cipher.sub_streamreaderread
func streamreaderread(r cipher.StreamReader, dst []byte) (int, error) {
	return r.Read(dst)
}

//go:linkname StreamReaderRead cipher.sub_streamreaderread
//go:noescape
func StreamReaderRead(r cipher.StreamReader, dst []byte) (int, error)

//go:linkname streamwriterwrite cipher.sub_streamwriterwrite
func streamwriterwrite(w cipher.StreamWriter, src []byte) (int, error) {
	return w.Write(src)
}

//go:linkname StreamWriterWrite cipher.sub_streamwriterwrite
//go:noescape
func StreamWriterWrite(w cipher.StreamWriter, src []byte) (int, error)

//go:linkname NewGCM crypto/cipher.NewGCM
//go:noescape
func NewGCM(cipher cipher.Block) (cipher.AEAD, error)

//go:linkname NewGCMWithTagSize crypto/cipher.NewGCMWithTagSize
//go:noescape
func NewGCMWithTagSize(cipher cipher.Block, tagSize int) (cipher.AEAD, error)

//go:linkname NewCFBDecrypter crypto/cipher.NewCFBDecrypter
//go:noescape
func NewCFBDecrypter(block cipher.Block, iv []byte) cipher.Stream
