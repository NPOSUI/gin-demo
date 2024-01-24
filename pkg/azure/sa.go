package azure

import (
	"fmt"
	"net/url"

	"github.com/Azure/azure-storage-file-go/azfile"
)

func GenerateServiceURL(saName string, saKey string) (*azfile.ServiceURL, error) {
	credential, err := azfile.NewSharedKeyCredential(saName, saKey)
	if err != nil {
		return nil, err
	}
	p := azfile.NewPipeline(credential, azfile.PipelineOptions{})
	u, _ := url.Parse(fmt.Sprintf("https://%s.file.core.chinacloudapi.cn", saName))
	serviceURL := azfile.NewServiceURL(*u, p)
	return &serviceURL, nil
}
