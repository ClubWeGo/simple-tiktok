package utils

import (
	"context"
	"github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/ssh"
	"log"
	"net"
	"os"
)

type SSHConfig struct {
	Host       string
	Port       string
	User       string
	PriKeyPath string
}

var config = SSHConfig{
	Host:       "124.221.147.131",
	Port:       "22",
	User:       "simple-tiktok",
	PriKeyPath: "/root/.ssh/id_rsa",
}

func dialSSH() *ssh.Client {
	// A public key may be used to authenticate against the remote
	// server by using an unencrypted PEM-encoded private key file.
	//
	// If you have an encrypted private key, the crypto/x509 package
	// can be used to decrypt it.
	key, err := os.ReadFile(config.PriKeyPath)
	if err != nil {
		log.Fatalf("unable to read private key: %v", err)
	}
	// Create the Signer for this private key.
	signer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		log.Fatalf("unable to parse private key: %v", err)
	}

	client, err := ssh.Dial("tcp", config.Host+":"+config.Port, &ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			// Use the PublicKeys method for remote authentication.
			ssh.PublicKeys(signer)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		log.Fatalf("unable to connect: %v", err)
	}
	//defer client.Close()
	return client
}

type SSHDialer struct {
	client *ssh.Client
	ctx    *context.Context
}

func (dialer *SSHDialer) Dial(ctx context.Context, addr string) (net.Conn, error) {
	return dialer.client.Dial("tcp", addr)
}

func RegisterSSH() {
	mysql.RegisterDialContext("mysql+ssh", (&SSHDialer{dialSSH(), nil}).Dial)
}
