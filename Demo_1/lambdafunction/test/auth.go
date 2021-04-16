package main

import (
	"bytes"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"

	"github.com/gbrlsnchs/jwt/v3"
)

var l_cacheKeys MapOfKidToPublicKey

type TokenHeader struct {
	Kid string
	Alg string
}

type PublicKey struct {
	Alg string
	E   string
	Kid string
	Kty string
	N   string
	Use string
}

type PublicKeys struct {
	Keys []PublicKey
}

type PublicKeyMeta struct {
	instance PublicKey
	pem      string
	pubKey   *rsa.PublicKey
}

type MapOfKidToPublicKey struct {
	keymap map[string]PublicKeyMeta
}

type Claim struct {
	Token_use string
	Auth_time int64
	Iss       string
	Exp       int64
	Username  string
	Client_id string
}

func AuthHandler(accessToken string) (bool, string, error) {
	print(`user claim verify invoked`)
	tokenSections := strings.Split(accessToken, ".")
	if len(tokenSections) < 2 {
		return false, "err", fmt.Errorf("Requested token is invalid")
	}
	print(tokenSections[0])
	headerJson, err := base64Decode(tokenSections[0])
	var header TokenHeader
	json.Unmarshal([]byte(headerJson), &header)
	keys, err := getPublicKeys()
	if !(err == nil) {
		return false, "err", err
	}

	key, ok := keys.keymap[header.Kid]
	if !ok {
		return false, "claim made for unknown kid", fmt.Errorf("Claim made for unknown kid")
	}
	var claim Claim
	rs := jwt.NewRS256(jwt.RSAPublicKey(key.pubKey))
	_, err = jwt.Verify([]byte(accessToken), rs, &claim)
	if !(err == nil) {
		return false, "claim is  invalid", err
	}
	currenttime := time.Now().UnixNano() / int64(time.Second)
	if currenttime > claim.Exp || currenttime < claim.Auth_time {
		return false, "claim is expired", err
	}
	if claim.Iss != "https://cognito-idp.us-east-2.amazonaws.com/us-east-2_E4FNAOYCN" {
		return false, "claim issuer is invalid", fmt.Errorf("Claim issuer is invalid")
	}
	if claim.Token_use != "access" {
		return false, "claim use is not access", fmt.Errorf("Claim use is not access")
	}

	print(`claim confirmed`)

	return true, claim.Username, nil
}

func base64Decode(str string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(str + "=")
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func getPublicKeys() (MapOfKidToPublicKey, error) {
	var keys PublicKeys
	if len(l_cacheKeys.keymap) == 0 {
		url := "https://cognito-idp.us-east-2.amazonaws.com/us-east-2_E4FNAOYCN/.well-known/jwks.json"
		publicKeys, err := http.Get(url)
		if !(err == nil) {
			panic(err)
		}

		body, err := ioutil.ReadAll(publicKeys.Body)
		data := body
		print(string(body))
		if err := json.Unmarshal(data, &keys); err != nil {
			return MapOfKidToPublicKey{}, err
		}
		returnMap := MapOfKidToPublicKey{keymap: make(map[string]PublicKeyMeta)}
		for _, v := range keys.Keys {
			returnMap.keymap[v.Kid] = jwkToPem(v)
		}
		l_cacheKeys = returnMap
		return returnMap, nil
	}
	return l_cacheKeys, nil
}

func jwkToPem(jwk PublicKey) PublicKeyMeta {

	nb, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		panic(err)
	}

	e := 0
	if jwk.E == "AQAB" || jwk.E == "AAEAAQ" {
		e = 65537
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		panic(err)
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)

	var ret = PublicKeyMeta{
		instance: jwk,
		pem:      out.String(),
		pubKey:   pk,
	}
	return ret
}
