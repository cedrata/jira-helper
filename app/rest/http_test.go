package rest

import (
	// "crypto/tls"
	// "net/http"
	// "net/http/httptest"
	// "net/url"
	//"strings"
	//"testing"

	// "github.com/cedrata/jira-helper/app/config"
	// "github.com/stretchr/testify/assert"
)

const certPEM string = `-----BEGIN CERTIFICATE-----
MIIDazCCAlOgAwIBAgIUEELuDfshkrGJhOSC1raBp8IqmvEwDQYJKoZIhvcNAQEL
BQAwRTELMAkGA1UEBhMCSVQxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoM
GEludGVybmV0IFdpZGdpdHMgUHR5IEx0ZDAeFw0yMzEyMTcwNzM4MDBaFw0yNDEy
MTYwNzM4MDBaMEUxCzAJBgNVBAYTAklUMRMwEQYDVQQIDApTb21lLVN0YXRlMSEw
HwYDVQQKDBhJbnRlcm5ldCBXaWRnaXRzIFB0eSBMdGQwggEiMA0GCSqGSIb3DQEB
AQUAA4IBDwAwggEKAoIBAQDcdIjcHrsakmEgDNHavjYKy43nsNSLCqyFddjKYE0o
HuDpodnzrRidYxzomHk29WsEecmipjxQA7N8fV9tIjU1daHMrlqGIjvmxPM04Ocn
4EAHhHtss+TSR9ycIPCcHChW1EoZCi3JN1BkFAS74f1vRHIZvrTrNCmEflbWL3NG
8SalbiR4fwm6GqbQ/gIi5DtfvZlAyC4c3WkRQu67ILcyLUml5SkcOAm4sKUke8ix
qywhqSQnQUwxrF76tAWIfIAUHThvUjRBCKdwrFqw0UO60uRPiSUFXfxn3HboON9y
mWi17JXTqsZ5GCZZiTR7rvfh0B9xJFlXlXmS4N05lBQXAgMBAAGjUzBRMB0GA1Ud
DgQWBBSuIBDyLvw9rK5sWBvzDQdaih26DzAfBgNVHSMEGDAWgBSuIBDyLvw9rK5s
WBvzDQdaih26DzAPBgNVHRMBAf8EBTADAQH/MA0GCSqGSIb3DQEBCwUAA4IBAQC5
hyKUzw53JNhe+O581DZyNMpYIj8cRLUH0d46GTDDs1aUryOttVXmWPG1vompnd/T
XZ/zO6sx2ULmIvY/CaII0wyL9QIiEeFdUTW+FkcEpeZxFAPd2dMbDkax7OmDjUAO
rlpNL2vyZ4VWED2iw4M7bTSKuZx+dfMR00+ICIYkVfiVg5yQWE9HDx6cmisZSA3H
qhNckh5wcnY4oAMAdQDl+9LDXxpY1hOyHhDbvfAoGyaYANHn57Kc7PGwnqattZ9P
DT0qbmr8PizYmS55dlztI0pF+Fr4Wa4PxiEDviPaU/LgJhvDk8yXmBCEjSDsApQM
u54MM9jXiFXKp9IG+Rm5
-----END CERTIFICATE-----
`

const keyPEM string = `-----BEGIN PRIVATE KEY-----
MIIEvAIBADANBgkqhkiG9w0BAQEFAASCBKYwggSiAgEAAoIBAQDcdIjcHrsakmEg
DNHavjYKy43nsNSLCqyFddjKYE0oHuDpodnzrRidYxzomHk29WsEecmipjxQA7N8
fV9tIjU1daHMrlqGIjvmxPM04Ocn4EAHhHtss+TSR9ycIPCcHChW1EoZCi3JN1Bk
FAS74f1vRHIZvrTrNCmEflbWL3NG8SalbiR4fwm6GqbQ/gIi5DtfvZlAyC4c3WkR
Qu67ILcyLUml5SkcOAm4sKUke8ixqywhqSQnQUwxrF76tAWIfIAUHThvUjRBCKdw
rFqw0UO60uRPiSUFXfxn3HboON9ymWi17JXTqsZ5GCZZiTR7rvfh0B9xJFlXlXmS
4N05lBQXAgMBAAECggEAYJgRFX3Dlqd3tg7X7oP0cvNwuImC/29MW2fg9v+OHxu0
ibn4oSwBgoiYdJPGXN3Yp8vjHQjAmYCdptjaNJvN+6AQpnnowSgD3iACvnMi5ZZ9
B641VFJYWwjQsXo/Yu91f2IiG2mZ2TYK2+bnkhk5rgSsB/rHE921qU+gJdYBqR2V
jqneMr/Wdc3P3VI4gLWCDPa/bX2aYJqH+igwK+og0DTPET1INb/nXy1bWl4EuNmc
2LTYkn7yfmiq3F3LYbm/t+BhpWvSfUZUE3N3iX+cUYwVTBrjccJ4i4dYX0UtKeah
03nCZjvMxEoqx25MeVkfRU2oe8/l6yDpqwXpPSqCLQKBgQDxft9HbCOP1Ku0EqlW
3uC7vaM+eaERh7oZGnsxGwHN1CaD3g28i0wZhuYmusCElv1NJvqrcht068wq3f1d
l5OmqvW8poqVJFYWXWmzypTl3g3oYJduBbv48eI45ar+UJj6k62ze42hGpSrwGhH
s75LQ8vxtwiUCxe5szslIwEWJQKBgQDpsiesTLyszwy77r4+F7gbLcElrGBMYz3C
FFB2oSyspNiznsjbe3wBGwmqYUkgOzLC6pDctqrrnAU7+HD13yrewRP2vgvcz/DL
iH0F+D9RtI8cpBilequQMr7kJTF1LgnvoyEz31J2X/1XrndRKJ5VeAjOs6ewtrYF
6SnXmih2iwKBgFEBLKYJePhC7wFtDQ5NnnZ7Cunm5Ic7zsmi31W+aGGBWxX6gwMO
eo1JeaPeGrue5gJeI9Ekal9SxN5QLi+Zq3ZJfDo2Zt/WG2ZPGSisuDtOu72JwOGv
3LiJckeBilTZ9iZ/KNG+jOhQQTRHSvNaMGeQqzU+HwuBmQi6PQmc7z1dAoGAVJ8Z
NI/Y5i5XKxoJM0y9csH/pZekiySIcWWPuVUlayKKAYimrKsrPO9AcbymkRA+kkwD
xpgyjfxB/PQ6Wx3DVUPO6dLpUrzNMbYrp2S78OcTx0g4UHt58k4dx1kcbpUMLgUA
+dqM7qZVg1F+jRnLM6Gydr6hIyEWCk/iwdplen8CgYBSV76f+fWKQ1C/NcHKC+qs
ZyTvJiVnaLWGbgD7ewjPT5R4setTOr2g5hhvOEi7WdPzpTgrexg7rHSZMtp0q4qc
EbZalZd/pfcnZlTMzcj7ckNQobIP7+Fat/ds1E7bkywtqJPrfGt2HvuusuB7GlLy
j8vFG3wadEctlKbHuyQVnQ==
-----END PRIVATE KEY-----
`

// func TestGetAllIssues(t *testing.T) {
// 	server := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 
// 		w.Header().Set("Content-Type", JSONContentType)
// 		w.WriteHeader(http.StatusOK)
// 		w.Write([]byte(`{"pippo": "pippo"}`))
// 
// 		t.Logf("Requested path: %s", r.URL)
// 
// 		// // Check the request method and path
// 		// assert.Equal(t, http.MethodGet, r.Method)
// 		// assert.Equal(t, "/rest/api/2/search?jql=project=your_project+order+by+duedate&fields=id,key", r.URL.Path)
// 
// 		// // Respond with a sample JSON response
// 		// w.Header().Set("Content-Type", JSONContentType)
// 		// w.WriteHeader(http.StatusOK)
// 		// fmt.Fprint(w, `{"key": "value"}`)
// 	}))
// 	defer server.Close()
// 
// 	// Generate a self-signed certificate for testing
// 	cert, err := tls.X509KeyPair([]byte(certPEM), []byte(keyPEM))
// 	if err != nil {
// 		t.Logf("Error loading certificate: %s", err)
// 		return
// 	}
// 
// 	// Configure the server to use the generated certificate
// 	server.TLS = &tls.Config{Certificates: []tls.Certificate{cert}}
// 
// 	// Start the server with TLS enabled
// 	server.StartTLS()
// 
// 	client := &http.Client{
// 		Transport: &http.Transport{
// 			TLSClientConfig: &tls.Config{
// 				InsecureSkipVerify: true,
// 			},
// 		},
// 	}
// 	parsedUrl, _ := url.Parse(server.URL)
// 	payload, err := Get(
// 		GetIssues,
// 		&config.Config{
// 			Token:   "token",
// 			JiraUrl: parsedUrl.Host,
// 			Project: "INAEDM",
// 		},
// 		client,
// 	)
// 
// 	t.Logf("payload: %s", payload)
// 
// 	assert.Nil(t, err)
// 	assert.NotEmpty(t, payload)
// }
