package service

// The JWTService interface defines methods needs to implement
type JWTService interface {
	GenerateJWT(id, email string) (string, error)
}

// The jWTService struct implements interface methods
type jWTService struct{}

// The NewJWTService returns jwtService struct which implement JWTService interface
func NewJWTService() JWTService {
	return &jWTService{}
}

// GenerateJWT takes id and email and returns JWT token
func (*jWTService) GenerateJWT(id, email string) (string, error) {

}
