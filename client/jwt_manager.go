package main

import (
	"github.com/golang-jwt/jwt"
	"log"
)

const (
	skey = `
-----BEGIN PRIVATE KEY-----
MIIJQgIBADANBgkqhkiG9w0BAQEFAASCCSwwggkoAgEAAoICAQDQIE1fY8lfcpLR
xCB8799Y8uy28hP9sz8ovjrIr+2rkNuuR0M7vlliaN+g2TFyrEtz9sjs2r3EQho3
aYrc9wnc62XCN4NU/HeTYQqt0R80mvXBGz2f45sOXmCVnhzmCjwfWstGrkVcRr0D
E1ry81oXPo8NLFKkHqgvZR95lzBEZRK+WPrKyTj+DGCtdNkh7jdfoj37uR8y0X70
NuCUHnjIRlfa3+jfy1SzfVLXMBCFOu6yvzqte4fl4mke3s4Qw00qyKemd217hPF+
WGB5q6p65k8X6DyjDnkfj4Z4jHOMPx29XROHKqFFISkDt4KMkSZg3JbjgNEnhE38
qRCnrKmjiAL7DkKeH5YLQ6i3XckROy0gaoBRPotYAwLFp/reBQwI1pHW2e9Z9fwc
+p6YK91N7q+stwaZUP89dywtnpjhP2wPL9y4iz6umImVGDCMnNPSolxb/xYHsAJb
ekcQTuOF59MxJ3gxLapHSIIT5tyNa5uipTp6png/UJQ/tUnvHWA04jrWZmWDv41+
/FRCTar4IKppNcvjmQ+sbwtvuaDC441bh/Vt+erQ0u+rOktSimBj6mtKo+gCTLDy
k7brAjdAvNQMGPgVTufr6EYEold54FzPrYU06v2TSH8Fnr/imLFwmtW7SQj4ubRZ
8M0jG8lK3q+OYzpLGh9lJOiESV4xyQIDAQABAoICACj8U4BXfE6Jl5QrlWJFMqyn
miok3f72D5gMbjMbQiipLDnT2s+nGp8gm/lnDC/S8kDx9nt+UCSYB+WlqSz+kTiq
Oq3WlMxls36Ye6vjJMwNQBw5KxeTkPNxbn/IVVSP62sZTDKtGwei/pHee+igN7ug
HDex5MxhJSnANkSRm3W8mcZWa30Hx2twMJG4ExJ1gP4aSy6JNDWjv+aD8BH7Njnq
cF6v/YIk3pdCNNAfrQEvq9YThpCnZQnGXoEvy1DNxSKJxUoo9qx+b7g72zPatIXj
WOwSVe6eaQfoEcknj5Gs1ER4UeIqaPveHd+XgeQAqZglmKlQst+fVGRPPvJmNDmf
IxqaWvgQt1Ze5D0raMTvj/TFXHDDSKoFm7TBcKI95WGPO/yugJ1LN251+pK1Fzf0
Wzye1s7NeXD82MLuliIYwIk2QKJ3cM6xot4NXJ5AW1ZkMurWi6/hTRK1j34wUN/5
1++To9bcCasrOmdvgqYn9273t6m+fRl7DkhORw7O2UnTz5u9JdibSIBDu8FwuRoV
TYpyw1XZ+aVeY/799buB6qRwF0zq9qRL/0Tc2ZQ1jmY/6f3jOmmQNMCAqxwO9UQQ
B+0Cmk8fZiAY3C29rHp/mddAV3C4kS7k3Rl2Gndl6cWdWqvCRkf9lrqb0FRErU0E
a+UYIRuvZb0MGvcq0cihAoIBAQDoVhcAoD8USFoOAWn8J3kQEQWk/Q1jlDGs6Y4t
zDlT/gV2DAIuGCSKJh32l5cpHcmmFw4htiOhLzwDTU2zL+iocrdpZ28QUyYrcieU
0F3OabZz36iFCDcU0p/5NEdlXNiEHQMNxZdiA6dYsWBjQvkViU1M9ObV16TOsV/X
ZtYn4Mf43AwWW8b9Pf0Kr41hS0GPN/YDY2Nsi1IbPtA1ezi/TEZozHoSWrQezRUs
jUy/Vga1pbC9OTkOLTy4llPHVOsw7QMMKDCwsHa2iVCojye5XLUwTEzNRW6nPJvi
Hp0WFetZ2PVOhkOgJTcxVUgmmMvYMTL10AZj6PjqBYSgYMu9AoIBAQDlUvX0sHxo
Q/2m50qIMW+yaVacEluLb3yyGh83HB+BBSM4lSVGIxCHteiIoOpnmC57jOWiI7en
oqbqpucSyF4NQwN1LKJh5tlY3iX7IW4kbi5wjcdnCnlZWkKSrvnzL1Cfk+Axq7pJ
9FE19n62iVIcwSWjCj1UAKYWnoh/wmL83+7GMWAY3P4npjLfwEt3kgAmlF/Aed1W
ECDh4B0yQBckF7K40BGS1Vp/xwHg3XYKdo1SSFy+QGeFZrJMCj3E/nDk4sUVE8bA
TzD6vzrmvGBaaGwrSXfipn/+DRaPw1HNAgp87OH3xb83+Q39LasDweXwTJQ8iS0Y
4xBu73hOKbj9AoIBAHQuOKpzd5eo5n+CbYFOK9fA61WpxPw6quyQjiyQp0o9CHYT
YzOxlvQOQ4WAIHLLFn8boFFzRzXe7N+p1GT6XBl/4/+tXXiIW9n1550e5QnHCxm+
igKcLQ8YlC6F2f3yA8Nszo4gdKbqtNBBQrXgU7ZVRnZpMNVWVG+Xexm3rveaC+WC
A/laQ5N3YAXr64LvowT/MuJW3Oz2bkR5kJxt1d6zQbI+LfznppKQczooHmy/k/NZ
u5uj9cgFDwwvAA35hPKeUpvNrStt0q9M0yy3CWJD3ccdIHq1cbPnNvxH2r2kkGAJ
ggGWWlLGyWa+AqpqVeTyFU2ejNOBOJet0rhwLvECggEAE9uSNsUNkM3Kaih/GLoH
LRDYEh241aMqzhVa/J+vzrFOMnkfyCoDJW1IHISqp5vS8pnpzJeBML/x6kHXopW9
JzLVWtEpooal7X8XFN75Nahg1xg/2xlaFrWtLByAwmEnfxoEAOkY8Yx5d83HfTD4
7kp/YtXhJ6QCvdSuhzmi7rjJaVofMyf3ziEjKKLzJgB7iNuySu2CkBBHeKe6f35z
QonWzRAfZXKaKpRmAj3LGe7YH/bqKQNUs4WIDPOaE7PgvMEyRbf3rvFskVn71L0d
Ltb5/umWuwal8K7bdEl6jOEPXW/5xUXXYBZt8q2AsIuayKso7vEF+bC33JNZ/JK6
VQKCAQEA49eoq/LIzFiauayiFCxFV/VJg4Rekd0vOoTj0TzZTG67hrJOlnfDloBD
025v+XBIShP2OEaRRAtp2aPZFrQl5D6K/g4FK77vUAvSoCEi0DkowJcCHSHzkydn
ikkgQbKhn9KzEgXB4hrANpE5Kk83SoEqijbw2708mNR/xSUpGPJD4sCx40oFDVCc
nh1ue2FPwEJwp2Mh2ih1oUby7iKYX3O+bonTknVSq4OGHGdB9rGziPIPPYMOwcaR
zCJK87MHdmP85e0Wpkk6dY/4xTH1OLAc0qxAUcUPFBf0SS/Sw+xJiHU9ovHpeAp5
O2SAA+FsG/IWQtt656JfyftucYAmUQ==
-----END PRIVATE KEY-----`
	//symetric key if we want to use symetric crypto HMAC
	key = "62erSGDG35wWn55KkE7QPwMNMFs3n/BkdFXNKt7WrHY="
)

type UserClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Role     string `json:"role"`
}

type Service struct {
	secretKey   string
	UserClaims  UserClaims
	authMethods map[string]bool
	kid         string
}

var serviceToken Service

func init() {
	serviceToken.secretKey = skey
	serviceToken.authMethods = map[string]bool{
		"/event.grpc.Event/GetEvent": true,
	}
	// skey -> kid = toto
	// key -> kid = titi
	serviceToken.kid = "toto"
	serviceToken.UserClaims.Issuer = "client.toto.com"
	serviceToken.UserClaims.Id = "myid"
	serviceToken.UserClaims.Role = "admin"
	serviceToken.UserClaims.Username = "user"
}

func (s Service) token() string {
	// below uses asymetric crypto
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(serviceToken.secretKey))
	if err != nil {
		log.Fatal("Error while loading private key: ", err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, s.UserClaims)

	token.Header["kid"] = serviceToken.kid
	token.Header["x5u"] = "http://www.toto.com/titi.pem"

	ss, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatal("Error while signing: ", err)
	}
	return ss
}
