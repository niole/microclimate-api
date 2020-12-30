#!/bin/bash

rm *.pem

echo "Starting self signed key creation process. NO PW. DEV ONLY"

echo "Generating ca's self signed certificate for all domains ever"
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem \
    -subj "/C=US/ST=California/L=Berkeley/O=Me/OU=Me/CN=localhost/emailAddress=niolenelson@gmail.com"

echo "Generating private key and CSR for the API server"
openssl req -newkey rsa:4096 -nodes -keyout server-key.pem -out server-req.pem \
    -subj "/C=US/ST=California/L=Berkeley/O=Me/OU=Me/CN=*/emailAddress=niolenelson@gmail.com"

echo "Generating private key and CSR for API consuming clients"
openssl req -newkey rsa:4096 -nodes -keyout client-key.pem -out client-req.pem \
    -subj "/C=US/ST=California/L=Berkeley/O=Me/OU=Me/CN=*/emailAddress=niolenelson@gmail.com"

echo "Signing the client's CSR"
openssl x509 -req -in client-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out clientcert.pem \
    -extfile all-ext.cnf -sha256

echo "Signing the server's CSR"
openssl x509 -req -in server-req.pem -days 365 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out server-cert.pem \
    -extfile all-ext.cnf -sha256

echo "Verifying client certificate"
openssl verify -CAfile ca-cert.pem clientcert.pem

echo "Verifying server certificate"
openssl verify -CAfile ca-cert.pem server-cert.pem
