#!/bin/sh
# Christophe Buffard OpenTV Inc.
# script to generate X.509 client certificate and its key pair

clear
days=1024
CLIENT_CERT="$1"
CLIENT_CERT_PEM="certs/$CLIENT_CERT.pem"
CLIENT_CERT_KEY="private/$CLIENT_CERT.key"
CLIENT_CERT_CSR="csr/$CLIENT_CERT.csr"
INTERMEDIATE_CERT_NAME="$2"
INTERMEDIATE_PRIVATE_KEY="private/$INTERMEDIATE_CERT_NAME.key"
INTERMEDIATE_CERTIFICATE_FILE="certs/$INTERMEDIATE_CERT_NAME.pem"

# generate's client key executed by the client
echo "INFO: generating client key"
openssl genrsa -out $CLIENT_CERT_KEY 2048


# generating the CSR executed by the client
echo "UK\nWales\nCwambran\nOpenTV\nUEX\n$CLIENT_CERT.opentv.com:8080\nwww.opentv.com\n\n\n" | openssl req -config openssl-rootca.cnf -key $CLIENT_CERT_KEY -new -sha256 -out $CLIENT_CERT_CSR

# Sign certificate, this is done by the CA
openssl x509 -days 1024 -req -in $CLIENT_CERT_CSR -CA $INTERMEDIATE_CERTIFICATE_FILE -CAkey $INTERMEDIATE_PRIVATE_KEY -out $CLIENT_CERT_PEM

###################
## Div useful commamds
#
# to generate a new x509 certicate
# openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout unsigned-client.key -out unsigned-client.crt
# to generate a .pem form a .crt
# openssl x509 -in mycert.crt -out mycert.pem -outform PEM
# to check a x509 certificate
# openssl x509 -in certificate.crt -text -noout
# CSR decoded
# openssl req -in mycsr.csr -noout -text
