package handlers

import "SuperSiteLogsApp/services"

type KeyPathSelect struct{}

func (k *KeyPathSelect) SetKey(base64string string, fileName string) error {
	return services.SetKey(base64string, fileName)
}

func (k *KeyPathSelect) GetKeyName() (string, error) {
	return services.GetKeyName()
}
