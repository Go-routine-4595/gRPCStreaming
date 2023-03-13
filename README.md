# gRPCStreaming poject README

This is a toy project using gRPC streaming capability and the interception methods.

It uses TLS to secure the connection with x.509 certificate

It uses  a JWT token to for the authorization.

To connect you need to present the right certificate (mutual authentication between the client and the server)

To call the method you need to present a JWT token.



The JWT uses header claim to identify the key used, this will potentially allow to identify the client and which key to use to authenticate the token, and it uses the standard claim so the server can authorize, or not the access the the resource.



It uses the interceptor mechanism

- to verify the token signature
- identify the client (using header claim kid)
- check the client role to authorize the access to the resource