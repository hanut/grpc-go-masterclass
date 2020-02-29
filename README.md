# gRPC [Golang] Master Class: Build Modern API & Microservices

This repo contains my implementations of the code along (hands on)
and assignments based coding undertaken during Stephane Maarek's
[gRPC [Golang] Master Class: Build Modern API & Microservices] (https://www.udemy.com/course/grpc-golang/)

I made some modifications and added some tweaks but its basically everything
that he covered. His course is excellent and Stephane is an amazing
teacher. I really recommend checking it out if you are interested in GoLang and gRPC

## Notes

During the SSL section Stephane covers generation of self signed certificates.
The process he followed was quite complex and since self signed certificates are 
anyways not recommended in production, I decided to do what I always do for self
signed certs and used the shorcut for it - 
```
sudo openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout localhost.key -out localhost.crt
```
This will give you the .crt and .key files you need to get up and running.

You will still get an invalid CA error in your client though. To solve this, simply do the following: 
- Go to /usr/local/share/ca-certificates/
- Create a new folder, i.e. "sudo mkdir localhost"
- Set the permissions of the folder to 755
- Copy the . crt file into the new folder.
- Set the permissions of the file to 644
- Run "sudo update-ca-certificates"

Voila! Your localhost certificates are ready to use (for 365 days).
