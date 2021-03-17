# snippetBox-microservice

Before start:

1) Read and generate TLS certificate Chapter 10.1 Alex Edwards Let's Go
2) Create tls directory in the root and move 2 generated .pem files there
3) Create a Google client ID and client secret:
  https://www.loginradius.com/blog/async/google-authentication-with-golang-and-goth/
4) Update env variables

P.S.
1) type https://localhost:#Your_address in the url, not just localhost:#Your_address
2) no refresh tokens, only access token
3) remove google warning about https (optional):
  https://peacocksoftware.com/blog/make-chrome-auto-accept-your-self-signed-certificate (then restart pc)