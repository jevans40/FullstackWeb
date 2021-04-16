package main

import (
	"testing"
)

func TestGetPublicKeys(t *testing.T) {
	test, _ := getPublicKeys()
	_ = test
	print("Done")
}

func TestAuth(t *testing.T) {
	auth := "eyJraWQiOiJXWWpWQ0xPY2VEVUYxUmh4VHlTdWNMWExVR2l1YTVwdnU5YWhtM0xwRkR3PSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiJjMDcyNGI0Yi02ZTEzLTQ0ODQtOTNhOS0xZmE1YTg0Y2UwNDkiLCJldmVudF9pZCI6ImY2NDQ3NWYyLWY3N2QtNDU0Yi1hOGJjLTY5YzAwMzNhNzAxMyIsInRva2VuX3VzZSI6ImFjY2VzcyIsInNjb3BlIjoiYXdzLmNvZ25pdG8uc2lnbmluLnVzZXIuYWRtaW4gcGhvbmUgb3BlbmlkIHByb2ZpbGUgZW1haWwiLCJhdXRoX3RpbWUiOjE2MTU1NDg1NjQsImlzcyI6Imh0dHBzOlwvXC9jb2duaXRvLWlkcC51cy1lYXN0LTIuYW1hem9uYXdzLmNvbVwvdXMtZWFzdC0yX0U0Rk5BT1lDTiIsImV4cCI6MTYxNTU1MjE2NCwiaWF0IjoxNjE1NTQ4NTY0LCJ2ZXJzaW9uIjoyLCJqdGkiOiIwMWY4NWYyNS00MGJiLTQ1YWYtODYzMi1jYTM5YTRmY2ZmZTAiLCJjbGllbnRfaWQiOiI0ajg0ZGV0dXAzbWJ2anN2MG44OHB2YjMycCIsInVzZXJuYW1lIjoiamVzc2UifQ.DX8rh5JvCD6uS2dbZkG3t7ZzzN3L2oF0AOs61tl3I5KiQjomszBOiVDIYtNXHLZ8BKrOGzouB7iz5Dyjj0Y8ac5CaB9cV7PE2PKu4kiuwZx52Rk8TH5Oha0lSD2Ju2VUnF_PT-1EdocuEdbhPuoj4_-ek1FZSGe3fmZDSHQUe6sndRSPGE0xdQsD6xGLkC7FU7Ul79yei8p72sjrBwpwa2iIZbTEKgejcvt7ADGNx1H3hqTIrIM4R5_fNkCg0MqB81jBAMGJefDOwKZvMpFK-OFX2du_JTzioBm-iURsFkLy8wU7lrD9Zff7VPpvDVnKOUOY3HIiLRCQvAkVFFhhLg"
	valid, message, err := AuthHandler(auth)
	_ = err
	_ = message
	_ = valid
}

/*
"keys":[
	{"alg":"RS256",
	"e":"AQAB",
	"kid":"UeZzi0T38lhKhXH3Jb3PbO55vZfd1DdpAj/oFJ2al0o=",
	"kty":"RSA",
	"n":"2CHTSgWnCPFnfk6pvayFAra3WuewtzXhq-w0N7fWba6ZnBPtjDLEv_ApSAdPBwFUzA2OawC6JkMrIhUep_UJBnfKrKn1mm4OwwpSh-c8HrT4r2svb9vm8F-YzMemeQmXqiQQN-3Og7rwD-SoHaiyaln5cwhsil4xMRiBUERcJMBd9mzpQR9toFnJjfHwBmpPsNNTAVDHcSUgR6UrRCAy467YrnrVX0E8BF0Lg6NTE6hesG7kUJDKAdT9HeiYpjnnILyQTNEtMc_Xn83FVcA11kcyUxkuQBfT3ooNj4768HW9VU1A_d8qthlbX7PfKarmX3Fnn4vU8OaWcpq-HCC4Sw",
	"use":"sig"},

	{"alg":"RS256",
	"e":"AQAB",
	"kid":"WYjVCLOceDUF1RhxTySucLXLUGiua5pvu9ahm3LpFDw=",
	"kty":"RSA",
	"n":"13mAJT69qhFyY4ExEuL0mLTAJN3Wj1iKfWtSeKEx9i_kE03Egs5h2H3kIPIpLbC3OYoj0FC4sDJCo7fNIVgNgclSaKePFPjMsXADIDfF3tbRIVJBrQu45n7M537LSayllaT77ydQrVtU21AzG68LQ5sMh7hMacaKQettaswwvFmT0vXcn_J0w0eZLjNKftaN0NEvmYomJ22aHct27dmSfq4C107vV3PzBkJLlKoXJxymP2MmEVUciSXv5OPWOQfn0rXbHhU5fY9Gf75PGcjSuCeIOYX35R6HYETKRij3NtA5dmfp_PeE2ykt0gQJyGa5xSl1-pkWdOUTgZ3DDX51-Q",
	"use":"sig"}]}
*/
