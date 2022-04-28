// Package emailserve email service
//
// As a quick start:
// 	m := gomail.NewMessage()
//	m.SetHeader(From, viper.GetString("email.from"))
//	m.SetHeader(To, viper.GetStringSlice("email.to")...)
//	m.SetHeader(Subject, "Hello!")
//	m.SetBody("text/html", "Hello <b>Candra</b>!")
//	DialAndSend(viper.GetString("email.host"), viper.GetInt("email.port"), viper.GetString("email.username"), viper.GetString("email.password"), m)
package emailserve
