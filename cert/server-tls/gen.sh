rm *.pem

# 1. Generate CA's private key and self-signed certificate
openssl req -x509 -newkey rsa:4096 -days 365 -nodes -keyout ca-key.pem -out ca-cert.pem -subj "/C=RU/ST=Moscow/L=Mocsow/O=GophKeeper/OU=GophKeeper Inc/CN=gophkeeper.ru/emailAddress=ksusonic@ya.ru"

echo "CA's self-signed certificate"
openssl x509 -in ca-cert.pem -noout -text

# 2. Generate web server's private key and certificate signing request (CSR)
openssl req -newkey rsa:4096 -nodes -keyout key.pem -out req.pem -subj "/C=RU/ST=Moscow/L=Mocsow/O=GophKeeper/OU=GophKeeper Inc/CN=gophkeeper.ru/emailAddress=ksusonic@ya.ru"

# 3. Use CA's private key to sign web server's CSR and get back the signed certificate
openssl x509 -req -in req.pem -days 60 -CA ca-cert.pem -CAkey ca-key.pem -CAcreateserial -out cert.pem

echo "Server's signed certificate"
openssl x509 -in cert.pem -noout -text