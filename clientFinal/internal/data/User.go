package data

import (
	pb "clientFinal/pkg/auth/proto"
	"golang.org/x/crypto/bcrypt"
)

func Set(plaintextPassword string) *pb.Password {
	p := pb.Password{}
	hash, err := bcrypt.GenerateFromPassword([]byte(plaintextPassword), 12)
	if err != nil {
		return &p
	}
	p.PlainText = plaintextPassword
	p.Hash = hash
	return &p
}
