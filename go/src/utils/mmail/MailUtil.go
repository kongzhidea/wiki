package mmail

import (
	"crypto/tls"
	"fmt"
	"net"
	"net/smtp"
	"strings"
)

const (
	Host    = "smtp.exmail.qq.com" // 腾讯企业邮箱，其他的如smtp.163.com
	Port    = "25"
	Address = Host + ":" + Port
)

// 多个接收人使用  ; 隔开，  25
func SendToMail(user, password, to, subject, body string) error {
	auth := smtp.PlainAuth("", user, password, Host)
	var content_type = "Content-Type: text/html; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(Address, auth, user, send_to, msg)
	return err
}

//const (
//	Host    = "smtp.****.com" // ssl加密发送邮件，465端口
//	Port    = "465"
//	Address = Host + ":" + Port
//)

// 多个接收人使用  ; 隔开， 465
func SendToMailSSl(user, password, to, subject, body string) error {
	auth := smtp.PlainAuth("", user, password, Host)
	var content_type = "Content-Type: text/html; charset=UTF-8"

	msg := []byte("To: " + to + "\r\nFrom: " + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := sendMailUsingTLS(Address, auth, user, send_to, msg)
	return err
}

//return a smtp client
func dial(addr string) (*smtp.Client, error) {
	conn, err := tls.Dial("tcp", addr, nil)
	if err != nil {
		fmt.Println("Dialing Error:", err)
		return nil, err
	}
	//分解主机端口字符串
	host, _, _ := net.SplitHostPort(addr)
	return smtp.NewClient(conn, host)
}

//参考net/smtp的func SendMail()
//使用net.Dial连接tls(ssl)端口时,smtp.NewClient()会卡住且不提示err
//len(to)>1时,to[1]开始提示是密送
func sendMailUsingTLS(addr string, auth smtp.Auth, from string,
	to []string, msg []byte) (err error) {

	//create smtp client
	c, err := dial(addr)
	if err != nil {
		fmt.Println("Create smpt client error:", err)
		return err
	}
	defer c.Close()

	if auth != nil {
		if ok, _ := c.Extension("AUTH"); ok {
			if err = c.Auth(auth); err != nil {
				fmt.Println("Error during AUTH", err)
				return err
			}
		}
	}

	if err = c.Mail(from); err != nil {
		return err
	}

	for _, addr := range to {
		if err = c.Rcpt(addr); err != nil {
			return err
		}
	}

	w, err := c.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	return c.Quit()
}
