package samlidp

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const spMetadata = "<EntityDescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\" validUntil=\"2015-12-03T01:57:09Z\" entityID=\"https://example.com/saml2/metadata\"><SPSSODescriptor xmlns=\"urn:oasis:names:tc:SAML:2.0:metadata\" validUntil=\"0001-01-01T00:00:00Z\" protocolSupportEnumeration=\"urn:oasis:names:tc:SAML:2.0:protocol\" AuthnRequestsSigned=\"false\" WantAssertionsSigned=\"true\"><KeyDescriptor use=\"signing\"><KeyInfo xmlns=\"http://www.w3.org/2000/09/xmldsig#\"><X509Data><X509Certificate>MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UECAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoXDTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28xEjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTvSPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gfnqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90DvTLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==</X509Certificate></X509Data></KeyInfo></KeyDescriptor><KeyDescriptor use=\"encryption\"><KeyInfo xmlns=\"http://www.w3.org/2000/09/xmldsig#\"><X509Data><X509Certificate>MIIB7zCCAVgCCQDFzbKIp7b3MTANBgkqhkiG9w0BAQUFADA8MQswCQYDVQQGEwJVUzELMAkGA1UECAwCR0ExDDAKBgNVBAoMA2ZvbzESMBAGA1UEAwwJbG9jYWxob3N0MB4XDTEzMTAwMjAwMDg1MVoXDTE0MTAwMjAwMDg1MVowPDELMAkGA1UEBhMCVVMxCzAJBgNVBAgMAkdBMQwwCgYDVQQKDANmb28xEjAQBgNVBAMMCWxvY2FsaG9zdDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEA1PMHYmhZj308kWLhZVT4vOulqx/9ibm5B86fPWwUKKQ2i12MYtz07tzukPymisTDhQaqyJ8Kqb/6JjhmeMnEOdTvSPmHO8m1ZVveJU6NoKRn/mP/BD7FW52WhbrUXLSeHVSKfWkNk6S4hk9MV9TswTvyRIKvRsw0X/gfnqkroJcCAwEAATANBgkqhkiG9w0BAQUFAAOBgQCMMlIO+GNcGekevKgkakpMdAqJfs24maGb90DvTLbRZRD7Xvn1MnVBBS9hzlXiFLYOInXACMW5gcoRFfeTQLSouMM8o57h0uKjfTmuoWHLQLi6hnF+cvCsEFiJZ4AbF+DgmO6TarJ8O05t8zvnOwJlNCASPZRH/JmF8tX0hoHuAQ==</X509Certificate></X509Data></KeyInfo><EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes128-cbc\"></EncryptionMethod><EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes192-cbc\"></EncryptionMethod><EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#aes256-cbc\"></EncryptionMethod><EncryptionMethod Algorithm=\"http://www.w3.org/2001/04/xmlenc#rsa-oaep-mgf1p\"></EncryptionMethod></KeyDescriptor><AssertionConsumerService Binding=\"urn:oasis:names:tc:SAML:2.0:bindings:HTTP-POST\" Location=\"https://example.com/saml2/acs\" index=\"1\"></AssertionConsumerService></SPSSODescriptor></EntityDescriptor>"

func TestServicesCrud(t *testing.T) {
	test := NewServerTest()

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "https://idp.example.com/services/", nil)
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"services\":[]}\n", string(w.Body.Bytes()))

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("PUT", "https://idp.example.com/services/sp", strings.NewReader(spMetadata))
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "https://idp.example.com/services/sp", nil)
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, spMetadata, string(w.Body.Bytes()))

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "https://idp.example.com/services/", nil)
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"services\":[\"sp\"]}\n", string(w.Body.Bytes()))

	assert.Len(t, test.Server.serviceProviders, 2)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("DELETE", "https://idp.example.com/services/sp", nil)
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusNoContent, w.Code)

	w = httptest.NewRecorder()
	r, _ = http.NewRequest("GET", "https://idp.example.com/services/", nil)
	test.Server.ServeHTTP(w, r)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "{\"services\":[]}\n", string(w.Body.Bytes()))
	assert.Len(t, test.Server.serviceProviders, 1)
}
