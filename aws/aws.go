package knife

import (
	"crypto"
	"crypto/rsa"
	"crypto/x509"
	"crypto/sha256"
	"encoding/base64"
	"encoding/pem"
	"fmt"
    "github.com/golang/glog"
	"flag"
)

var (
	awsInstanceDocURL = "http://169.254.169.254/latest/dynamic/instance-identity/document"
	awsPublicCert = []byte( `-----BEGIN CERTIFICATE-----
MIIDIjCCAougAwIBAgIJAKnL4UEDMN/FMA0GCSqGSIb3DQEBBQUAMGoxCzAJBgNV
BAYTAlVTMRMwEQYDVQQIEwpXYXNoaW5ndG9uMRAwDgYDVQQHEwdTZWF0dGxlMRgw
FgYDVQQKEw9BbWF6b24uY29tIEluYy4xGjAYBgNVBAMTEWVjMi5hbWF6b25hd3Mu
Y29tMB4XDTE0MDYwNTE0MjgwMloXDTI0MDYwNTE0MjgwMlowajELMAkGA1UEBhMC
VVMxEzARBgNVBAgTCldhc2hpbmd0b24xEDAOBgNVBAcTB1NlYXR0bGUxGDAWBgNV
BAoTD0FtYXpvbi5jb20gSW5jLjEaMBgGA1UEAxMRZWMyLmFtYXpvbmF3cy5jb20w
gZ8wDQYJKoZIhvcNAQEBBQADgY0AMIGJAoGBAIe9GN//SRK2knbjySG0ho3yqQM3
e2TDhWO8D2e8+XZqck754gFSo99AbT2RmXClambI7xsYHZFapbELC4H91ycihvrD
jbST1ZjkLQgga0NE1q43eS68ZeTDccScXQSNivSlzJZS8HJZjgqzBlXjZftjtdJL
XeE4hwvo0sD4f3j9AgMBAAGjgc8wgcwwHQYDVR0OBBYEFCXWzAgVyrbwnFncFFIs
77VBdlE4MIGcBgNVHSMEgZQwgZGAFCXWzAgVyrbwnFncFFIs77VBdlE4oW6kbDBq
MQswCQYDVQQGEwJVUzETMBEGA1UECBMKV2FzaGluZ3RvbjEQMA4GA1UEBxMHU2Vh
dHRsZTEYMBYGA1UEChMPQW1hem9uLmNvbSBJbmMuMRowGAYDVQQDExFlYzIuYW1h
em9uYXdzLmNvbYIJAKnL4UEDMN/FMAwGA1UdEwQFMAMBAf8wDQYJKoZIhvcNAQEF
BQADgYEAFYcz1OgEhQBXIwIdsgCOS8vEtiJYF+j9uO6jz7VOmJqO+pRlAbRlvY8T
C1haGgSI/A1uZUKs/Zfnph0oEI0/hu1IIJ/SKBDtN5lvmZ/IzbOPIJWirlsllQIQ
7zvWbGd9c9+Rm3p04oTvhup99la7kZqevJK0QRdD/6NpCKsqP/0=
-----END CERTIFICATE-----` )
	
)

func VerifyInstance( ec2Sig string, ec2Doc []byte ) bool {
	decodedSig, _ := base64.StdEncoding.DecodeString( ec2Sig )
	
	certBlock, _ := pem.Decode( awsPublicCert )
	awsCert, errMsg := x509.ParseCertificate( certBlock.Bytes )
	if errMsg != nil {
		fmt.Println( "Failed To Parse awsPublicCert", errMsg.Error() )
		return false
	}
	awsPubKey := awsCert.PublicKey.( *rsa.PublicKey )
	
	docHash := sha256.Sum256( ec2Doc )
	errMsg = rsa.VerifyPKCS1v15( awsPubKey, crypto.SHA256, docHash[ : ], decodedSig )
	if errMsg != nil {
		glog.Error( errMsg.Error() )
		return false
	}
	return true
}

func main() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")

	sigData := `SF3qVlCDpm/vpb052Ez9E41UPB6RiUZjh4rG45samMQMAw3e89ZCwVmqIGrZuR/9fiXNKDwtl1KGpcPWMvKLMRzRpxByYYBr4ZD+t1d6Tl0HMnlBLQeh/KQOrkBCioyWIDM2shYLuhWy+q4oDyzKpNlsM44Z1GgrR3WnaDyFnwU=`
	testDoc := []byte( `{
  "devpayProductCodes" : null,
  "privateIp" : "28.8.9.60",
  "availabilityZone" : "eu-west-1b",
  "version" : "2010-08-31",
  "instanceId" : "i-05a78619c9d863912",
  "billingProducts" : null,
  "instanceType" : "m4.large",
  "accountId" : "480897156998",
  "architecture" : "x86_64",
  "kernelId" : null,
  "ramdiskId" : null,
  "imageId" : "ami-aee420d7",
  "pendingTime" : "2017-09-20T18:45:23Z",
  "region" : "eu-west-1"
}` )
	fmt.Println( VerifyInstance( sigData, testDoc ) )
}
