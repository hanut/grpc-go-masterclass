In case you want to generate more comprehensive server + client certs, 
you can use the following script - 


```
#!/bin/bash
echo "Creating certs folder ..."

cd certs

echo "Generating certificates ..."

openssl genrsa -passout pass:1111 -des3 -out ca.key 4096

openssl req -passin pass:1111 -new -x509 -days 365 -key ca.key -out ca.crt -subj  "/C=IN/ST=Uttarakand/L=Dehradun/O=KSDAC LLP/OU=Engineering/CN=ksdacllp.com"

openssl genrsa -passout pass:1111 -des3 -out server.key 4096

openssl req -passin pass:1111 -new -key server.key -out server.csr -subj  "/C=IN/ST=Uttarakhand/L=Dehradun/O=KSDAC LLP/OU=Engineering/CN=localhost"

openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

openssl rsa -passin pass:1111 -in server.key -out server.key

openssl genrsa -passout pass:1111 -des3 -out client.key 4096

openssl req -passin pass:1111 -new -key client.key -out client.csr -subj  "/C=IN/ST=Uttarakhand/L=Dehradun/O=KSDAC LLP/OU=Engineering/CN=localhost"

openssl x509 -passin pass:1111 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out client.crt

openssl rsa -passin pass:1111 -in client.key -out client.key
```
