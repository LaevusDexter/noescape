// This file has automatically been generated on Wed Feb 26 15:50:24 +05 2020.
// DO NOT EDIT.
package x509

import (
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"io"
	"time"
	_ "unsafe"
)

//go:linkname ParseCertificates crypto/x509.ParseCertificates
//go:noescape
func ParseCertificates(asn1Data []byte) ([]*x509.Certificate, error)

//go:linkname certificateverifyhostname x509.sub_certificateverifyhostname
func certificateverifyhostname(c *x509.Certificate, h string) error {
	return c.VerifyHostname(h)
}

//go:linkname CertificateVerifyHostname x509.sub_certificateverifyhostname
//go:noescape
func CertificateVerifyHostname(c *x509.Certificate, h string) error

//go:linkname certificaterequestchecksignature x509.sub_certificaterequestchecksignature
func certificaterequestchecksignature(c *x509.CertificateRequest,) error {
	return c.CheckSignature()
}

//go:linkname CertificateRequestCheckSignature x509.sub_certificaterequestchecksignature
//go:noescape
func CertificateRequestCheckSignature(c *x509.CertificateRequest,) error

//go:linkname ParseCertificateRequest crypto/x509.ParseCertificateRequest
//go:noescape
func ParseCertificateRequest(asn1Data []byte) (*x509.CertificateRequest, error)

//go:linkname CreateCertificate crypto/x509.CreateCertificate
//go:noescape
func CreateCertificate(rand io.Reader, template, parent *x509.Certificate, pub, priv interface{}) ([]byte, error)

//go:linkname CreateCertificateRequest crypto/x509.CreateCertificateRequest
//go:noescape
func CreateCertificateRequest(rand io.Reader, template *x509.CertificateRequest, priv interface{}) ([]byte, error)

//go:linkname ParseECPrivateKey crypto/x509.ParseECPrivateKey
//go:noescape
func ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)

//go:linkname certificateverify x509.sub_certificateverify
func certificateverify(c *x509.Certificate, opts x509.VerifyOptions,) ([][]*x509.Certificate, error) {
	return c.Verify(opts)
}

//go:linkname CertificateVerify x509.sub_certificateverify
//go:noescape
func CertificateVerify(c *x509.Certificate, opts x509.VerifyOptions,) ([][]*x509.Certificate, error)

//go:linkname certpoolsubjects x509.sub_certpoolsubjects
func certpoolsubjects(s *x509.CertPool,) [][]byte {
	return s.Subjects()
}

//go:linkname CertPoolSubjects x509.sub_certpoolsubjects
//go:noescape
func CertPoolSubjects(s *x509.CertPool,) [][]byte

//go:linkname certificatechecksignature x509.sub_certificatechecksignature
func certificatechecksignature(c *x509.Certificate, algo x509.SignatureAlgorithm, signed, signature []byte) error {
	return c.CheckSignature(algo, signed, signature)
}

//go:linkname CertificateCheckSignature x509.sub_certificatechecksignature
//go:noescape
func CertificateCheckSignature(c *x509.Certificate, algo x509.SignatureAlgorithm, signed, signature []byte) error

//go:linkname DecryptPEMBlock crypto/x509.DecryptPEMBlock
//go:noescape
func DecryptPEMBlock(b *pem.Block, password []byte) ([]byte, error)

//go:linkname certificatecheckcrlsignature x509.sub_certificatecheckcrlsignature
func certificatecheckcrlsignature(c *x509.Certificate, crl *pkix.CertificateList) error {
	return c.CheckCRLSignature(crl)
}

//go:linkname CertificateCheckCRLSignature x509.sub_certificatecheckcrlsignature
//go:noescape
func CertificateCheckCRLSignature(c *x509.Certificate, crl *pkix.CertificateList) error

//go:linkname hostnameerrorerror x509.sub_hostnameerrorerror
func hostnameerrorerror(h x509.HostnameError,) string {
	return h.Error()
}

//go:linkname HostnameErrorError x509.sub_hostnameerrorerror
//go:noescape
func HostnameErrorError(h x509.HostnameError,) string

//go:linkname publickeyalgorithmstring x509.sub_publickeyalgorithmstring
func publickeyalgorithmstring(algo x509.PublicKeyAlgorithm,) string {
	return algo.String()
}

//go:linkname PublicKeyAlgorithmString x509.sub_publickeyalgorithmstring
//go:noescape
func PublicKeyAlgorithmString(algo x509.PublicKeyAlgorithm,) string

//go:linkname unhandledcriticalextensionerror x509.sub_unhandledcriticalextensionerror
func unhandledcriticalextensionerror(h x509.UnhandledCriticalExtension,) string {
	return h.Error()
}

//go:linkname UnhandledCriticalExtensionError x509.sub_unhandledcriticalextensionerror
//go:noescape
func UnhandledCriticalExtensionError(h x509.UnhandledCriticalExtension,) string

//go:linkname IsEncryptedPEMBlock crypto/x509.IsEncryptedPEMBlock
//go:noescape
func IsEncryptedPEMBlock(b *pem.Block) bool

//go:linkname MarshalPKIXPublicKey crypto/x509.MarshalPKIXPublicKey
//go:noescape
func MarshalPKIXPublicKey(pub interface{}) ([]byte, error)

//go:linkname ParseDERCRL crypto/x509.ParseDERCRL
//go:noescape
func ParseDERCRL(derBytes []byte) (*pkix.CertificateList, error)

//go:linkname ParsePKIXPublicKey crypto/x509.ParsePKIXPublicKey
//go:noescape
func ParsePKIXPublicKey(derBytes []byte) (interface{}, error)

//go:linkname ParseCertificate crypto/x509.ParseCertificate
//go:noescape
func ParseCertificate(asn1Data []byte) (*x509.Certificate, error)

//go:linkname constraintviolationerrorerror x509.sub_constraintviolationerrorerror
func constraintviolationerrorerror(receiverV x509.ConstraintViolationError,) string {
	return receiverV.Error()
}

//go:linkname ConstraintViolationErrorError x509.sub_constraintviolationerrorerror
//go:noescape
func ConstraintViolationErrorError(receiverV x509.ConstraintViolationError,) string

//go:linkname EncryptPEMBlock crypto/x509.EncryptPEMBlock
//go:noescape
func EncryptPEMBlock(rand io.Reader, blockType string, data, password []byte, alg x509.PEMCipher,) (*pem.Block, error)

//go:linkname MarshalECPrivateKey crypto/x509.MarshalECPrivateKey
//go:noescape
func MarshalECPrivateKey(key *ecdsa.PrivateKey) ([]byte, error)

//go:linkname NewCertPool crypto/x509.NewCertPool
//go:noescape
func NewCertPool() *x509.CertPool

//go:linkname SystemCertPool crypto/x509.SystemCertPool
//go:noescape
func SystemCertPool() (*x509.CertPool, error)

//go:linkname signaturealgorithmstring x509.sub_signaturealgorithmstring
func signaturealgorithmstring(algo x509.SignatureAlgorithm,) string {
	return algo.String()
}

//go:linkname SignatureAlgorithmString x509.sub_signaturealgorithmstring
//go:noescape
func SignatureAlgorithmString(algo x509.SignatureAlgorithm,) string

//go:linkname systemrootserrorerror x509.sub_systemrootserrorerror
func systemrootserrorerror(se x509.SystemRootsError,) string {
	return se.Error()
}

//go:linkname SystemRootsErrorError x509.sub_systemrootserrorerror
//go:noescape
func SystemRootsErrorError(se x509.SystemRootsError,) string

//go:linkname unknownauthorityerrorerror x509.sub_unknownauthorityerrorerror
func unknownauthorityerrorerror(e x509.UnknownAuthorityError,) string {
	return e.Error()
}

//go:linkname UnknownAuthorityErrorError x509.sub_unknownauthorityerrorerror
//go:noescape
func UnknownAuthorityErrorError(e x509.UnknownAuthorityError,) string

//go:linkname ParseCRL crypto/x509.ParseCRL
//go:noescape
func ParseCRL(crlBytes []byte) (*pkix.CertificateList, error)

//go:linkname certificatechecksignaturefrom x509.sub_certificatechecksignaturefrom
func certificatechecksignaturefrom(c *x509.Certificate, parent *x509.Certificate,) error {
	return c.CheckSignatureFrom(parent)
}

//go:linkname CertificateCheckSignatureFrom x509.sub_certificatechecksignaturefrom
//go:noescape
func CertificateCheckSignatureFrom(c *x509.Certificate, parent *x509.Certificate,) error

//go:linkname certificateinvaliderrorerror x509.sub_certificateinvaliderrorerror
func certificateinvaliderrorerror(e x509.CertificateInvalidError,) string {
	return e.Error()
}

//go:linkname CertificateInvalidErrorError x509.sub_certificateinvaliderrorerror
//go:noescape
func CertificateInvalidErrorError(e x509.CertificateInvalidError,) string

//go:linkname insecurealgorithmerrorerror x509.sub_insecurealgorithmerrorerror
func insecurealgorithmerrorerror(e x509.InsecureAlgorithmError,) string {
	return e.Error()
}

//go:linkname InsecureAlgorithmErrorError x509.sub_insecurealgorithmerrorerror
//go:noescape
func InsecureAlgorithmErrorError(e x509.InsecureAlgorithmError,) string

//go:linkname certpoolappendcertsfrompem x509.sub_certpoolappendcertsfrompem
func certpoolappendcertsfrompem(s *x509.CertPool, pemCerts []byte) bool {
	return s.AppendCertsFromPEM(pemCerts)
}

//go:linkname CertPoolAppendCertsFromPEM x509.sub_certpoolappendcertsfrompem
//go:noescape
func CertPoolAppendCertsFromPEM(s *x509.CertPool, pemCerts []byte) bool

//go:linkname certificatecreatecrl x509.sub_certificatecreatecrl
func certificatecreatecrl(c *x509.Certificate, rand io.Reader, priv interface{}, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) ([]byte, error) {
	return c.CreateCRL(rand, priv, revokedCerts, now, expiry)
}

//go:linkname CertificateCreateCRL x509.sub_certificatecreatecrl
//go:noescape
func CertificateCreateCRL(c *x509.Certificate, rand io.Reader, priv interface{}, revokedCerts []pkix.RevokedCertificate, now, expiry time.Time) ([]byte, error)

//go:linkname certificateequal x509.sub_certificateequal
func certificateequal(c *x509.Certificate, other *x509.Certificate,) bool {
	return c.Equal(other)
}

//go:linkname CertificateEqual x509.sub_certificateequal
//go:noescape
func CertificateEqual(c *x509.Certificate, other *x509.Certificate,) bool
