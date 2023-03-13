#!/bin/sh
# Christophe Buffard OpenTV Inc.
# script to generate X.509 root certificate and its key pair

clear
days=7300

#if no input argument we default the root certificate name
if [ $# -eq 0 ];
then
    ROOT_CERT_NAME="root-52-8-151-239"
else
    ROOT_CERT_NAME="$1"
fi
ROOT_PRIVATE_KEY="private/$ROOT_CERT_NAME.key"
ROOT_CERTIFICATE_FILE="certs/$ROOT_CERT_NAME.pem"

######
echo "Cleaning up all certificate structure"
rm certs/*
rm csr/*
rm index.*
rm index-intermediate.txt
rm intermediate/*
rm newcerts/*
rm private/*
rm serial
rm serial.old
rm serial-intermediate
rm serial-intermediate.old

touch index.txt
touch index-intermediate.txt
touch crl/crl.pem
touch crl/crlnumber

echo 1000 > serial
echo 2000 > serial-intermediate
echo 3000 > crl/crlnumber



# generate root key for CA
echo "INFO: generating root key"
openssl genrsa -out $ROOT_PRIVATE_KEY 2048

# if [ -f "root.pem" ]; then
#     echo "Current root.pem subject is: "
#     openssl x509 -text -in root.pem  | grep "Subject" | grep C=
# fi

# self sign CA certificate
echo "INFO: generating root certificate"
echo "US\nCalifornia\nMountain View\nOpenTV\nUEX\nroot.opentv.com\nwww.opentv.com\n" | openssl req -config openssl-rootca.cnf -new -x509 -sha256 -extensions v3_ca -key $ROOT_PRIVATE_KEY -days $days -out $ROOT_CERTIFICATE_FILE

# create link to the root key.
#ln -s $ROOT_CERTIFICATE_FILE certs/root-certificate.pem
#ln -s $ROOT_PRIVATE_KEY private/root-certificate.key
cp $ROOT_CERTIFICATE_FILE certs/root-certificate.pem
cp $ROOT_PRIVATE_KEY private/root-certificate.key

# join this two certificates into a single file
# cat root.pem root.key > root.pem2
# mv root.pem2 root.pem
# rm -f root.key


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
