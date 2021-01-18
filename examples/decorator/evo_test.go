package decorator

func defaultPasswd() string {
	return "superSecret"
}

func defaultHash() string {
	return "$argon2id$v=19$m=65536,t=1,p=4$FG7b21+4cacIrzTN0mulDg$uzoRnwcQSmF5vIEhprh0uOcJDBUmGAQ7EgbOe66LtFE"
}

func wrongHash() string {
	return "$$$$$"
}
