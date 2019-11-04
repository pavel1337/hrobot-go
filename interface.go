package client

import "gitlab.com/newsletter2go/hrobot-go/models"

type RobotClient interface {
	SetBaseURL(baseURL string)
	SetUserAgent(userAgent string)
	ServerGetList() ([]models.Server, error)
	ServerGet(ip string) (*models.Server, error)
	ServerSetName(ip string, input *models.ServerSetNameInput) (*models.Server, error)
	ServerReverse(ip string) (*models.Cancellation, error)
	KeyGetList() ([]models.Key, error)
	IPGetList() ([]models.IP, error)
	RDnsGetList() ([]models.Rdns, error)
	RDnsGet(ip string) (*models.Rdns, error)
	BootRescueGet(ip string) (*models.RescueOptions, error)
	BootRescueSet(ip string, input *models.RescueSetInput) (*models.RescueValues, error)
	ResetGet(ip string) (*models.Reset, error)
	ResetSet(ip string, input *models.ResetSetInput) (*models.ResetPost, error)
}
