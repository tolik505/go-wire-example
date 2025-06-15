package main

type (
	NatsApp        struct{}
	UserHandler    struct{}
	UserService    struct{}
	UserRepo       struct{}
	Config         struct{}
	DB             struct{}
	ConnectionInfo string
)

func (*NatsApp) Run() {
	println("app is running...")
	for {
		// listen to messages
	}
}

func NewNatsApp(userHandler *UserHandler) (*NatsApp, error) {
	println("init nats app")

	return &NatsApp{}, nil
}

func NewUserHandler(service *UserService) *UserHandler {
	println("init user handler")

	return &UserHandler{}
}

func NewUserService(repo *UserRepo) *UserService {
	println("init user service")

	return &UserService{}
}

func NewDefaultConfig() *Config {
	println("init config")

	return &Config{}
}

func NewUserRepo(cfg *Config, db *DB) (*UserRepo, error) {
	println("init user repo")

	return &UserRepo{}, nil
}

func NewDB(info ConnectionInfo) (*DB, error) {
	println("init db")

	return &DB{}, nil
}
