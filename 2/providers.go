package main

type (
	NatsApp          struct{}
	UserHandler      struct{}
	DashboardHandler struct{}
	UserService      struct{}
	DashboardService struct{}
	UserRepo         struct{}
	DashboardRepo    struct{}
	Config           struct{}
	DB               struct{}
	ConnectionInfo   string
)

func (*NatsApp) Run() {
	println("app is running...")
	for {
		// listen to messages
	}
}

func NewNatsApp(userHandler *UserHandler, dashHandler *DashboardHandler) (*NatsApp, error) {
	println("init nats app")

	return &NatsApp{}, nil
}

func NewUserHandler(service *UserService) *UserHandler {
	println("init user handler")

	return &UserHandler{}
}

func NewDashboardHandler(service *DashboardService) *DashboardHandler {
	println("init dashboard handler")

	return &DashboardHandler{}
}

func NewUserService(repo *UserRepo) *UserService {
	println("init user service")

	return &UserService{}
}

func NewDashboardService(repo *DashboardRepo) *DashboardService {
	println("init dashboard service")

	return &DashboardService{}
}

func NewDefaultConfig() *Config {
	println("init config")

	return &Config{}
}

func NewUserRepo(cfg *Config, db *DB) (*UserRepo, error) {
	println("init user repo")

	return &UserRepo{}, nil
}

func NewDashboardRepo(cfg *Config, db *DB) (*DashboardRepo, error) {
	println("init dashboard repo")

	return &DashboardRepo{}, nil
}

func NewDB(info ConnectionInfo) (*DB, error) {
	println("init db")

	return &DB{}, nil
}
