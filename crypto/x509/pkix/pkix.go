// This file has automatically been generated on Wed Feb 26 15:50:25 +05 2020.
// DO NOT EDIT.
package pkix

import (
	"crypto/x509/pkix"
	"time"
	_ "unsafe"
)

//go:linkname namestring pkix.sub_namestring
func namestring(n pkix.Name,) string {
	return n.String()
}

//go:linkname NameString pkix.sub_namestring
//go:noescape
func NameString(n pkix.Name,) string

//go:linkname nametordnsequence pkix.sub_nametordnsequence
func nametordnsequence(n pkix.Name,) pkix.RDNSequence {
	return n.ToRDNSequence()
}

//go:linkname NameToRDNSequence pkix.sub_nametordnsequence
//go:noescape
func NameToRDNSequence(n pkix.Name,) pkix.RDNSequence

//go:linkname rdnsequencestring pkix.sub_rdnsequencestring
func rdnsequencestring(r pkix.RDNSequence,) string {
	return r.String()
}

//go:linkname RDNSequenceString pkix.sub_rdnsequencestring
//go:noescape
func RDNSequenceString(r pkix.RDNSequence,) string

//go:linkname certificatelisthasexpired pkix.sub_certificatelisthasexpired
func certificatelisthasexpired(certList *pkix.CertificateList, now time.Time) bool {
	return certList.HasExpired(now)
}

//go:linkname CertificateListHasExpired pkix.sub_certificatelisthasexpired
//go:noescape
func CertificateListHasExpired(certList *pkix.CertificateList, now time.Time) bool
