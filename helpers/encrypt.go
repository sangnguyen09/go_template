package helpers

import (
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"reflect"
	"strings"
	"time"

	// "g.ghn.vn/go-common/encryption"
	// "g.ghn.vn/go-common/zap-logger"

	// "github.com/sangnguyen09/go_template/config"
	// "github.com/sangnguyen09/go_template/models"
)

// var log = logger.GetLogger("Utils")
// var encryptionAlgo = encryption.NewAesCryption(encryption.SetAESKey(config.Config.Encryption.OIDKey))

type AnonymousType reflect.Value

func GetSHA1WithKey(text, key string) string {
	res := []byte(fmt.Sprintf("%s:%s", text, key))
	buf := sha1.Sum(res)
	return hex.EncodeToString(buf[:])
}

func GetMD5Hash(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func Contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func (a AnonymousType) IsA(typeToAssert interface{}) bool {
	return typeToAssert == reflect.Value(a).Kind()
}
func ToAnonymousType(obj interface{}) AnonymousType {
	return AnonymousType(reflect.ValueOf(obj))
}

func GenerateApiKey() string {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%d-%x-%x-%x-%x-%x", time.Now().Unix(), uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}

func EncryptPass(s string) string {
	byteString := []byte(s)
	byteSalt := byte(79)
	for i := 0; i < len(byteString); i++ {
		if i%2 == 0 {
			//XOR
			byteString[i] ^= byteSalt
		}
	}
	return strings.ToUpper(fmt.Sprintf("%x", sha256.Sum256(byteString)))
}

// func TenantLogin(phone, pass string) (*models.TenantLoginBody, error) {
// 	type loginPayload struct {
// 		APIKey       string `json:"ApiKey"`
// 		APISecretKey string `json:"ApiSecretKey"`
// 		Phone        string `json:"Email"`
// 		Password     string `json:"Password"`
// 	}
// 	payload := loginPayload{config.Config.Tenant.ApiKey, config.Config.Tenant.ApiSecretKey, phone, pass}
// 	b, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, _ := http.NewRequest("POST", fmt.Sprintf(config.Config.Tenant.Url, "Client/Login"), strings.NewReader(fmt.Sprintf("%s", b)))
// 	req.Header.Add("Content-Type", "application/json")
// 	res, err := http.DefaultClient.Do(req)
// 	if err != nil {
// 		log.Error("Error TenantLogin", err)
// 		return nil, err
// 	}
// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)
// 	var data models.TenantLoginBody
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res.StatusCode != 200 {
// 		err = errors.New("CALL_API_FAIL")
// 	} else if data.ClientID == 0 {
// 		err = errors.New("LOGIN_FAIL")
// 	}
// 	return &data, err
// }

// func TenantRegister(email, fullname, phone, pass string) (*models.TenantRegisterBody, error) {
// 	type registerPayload struct {
// 		APIKey       string `json:"ApiKey"`
// 		APISecretKey string `json:"ApiSecretKey"`
// 		Email        string `json:"Email"`
// 		Password     string `json:"Password"`
// 		ContactPhone string `json:"ContactPhone"`
// 		CustomerName string `json:"CustomerName"`
// 	}
// 	payload := registerPayload{config.Config.Tenant.ApiKey, config.Config.Tenant.ApiSecretKey, email, pass, phone, fullname}
// 	b, err := json.Marshal(payload)
// 	if err != nil {
// 		return nil, err
// 	}
// 	req, _ := http.NewRequest("POST", fmt.Sprintf(config.Config.Tenant.Url, "Client/Register"), strings.NewReader(fmt.Sprintf("%s", b)))
// 	req.Header.Add("Content-Type", "application/json")
// 	res, _ := http.DefaultClient.Do(req)
// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)
// 	var data models.TenantRegisterBody
// 	err = json.Unmarshal(body, &data)
// 	if err != nil {
// 		return nil, err
// 	}
// 	if res.StatusCode != 200 || data.ClientID == 0 {
// 		err = errors.New("CALL_API_FAIL")
// 	}
// 	fmt.Println(&data)
// 	return &data, err
// }

// func TenantChangePass(oldPass, newPass string, id int) error {
// 	type changePassPayload struct {
// 		APIKey         string `json:"ApiKey"`
// 		APISecretKey   string `json:"ApiSecretKey"`
// 		ClientID       int    `json:"ClientID"`
// 		OldPass        string `json:"OldPass"`
// 		NewPass        string `json:"NewPass"`
// 		ConfirmNewPass string `json:"ConfirmNewPass"`
// 	}
// 	payload := changePassPayload{config.Config.Tenant.ApiKey, config.Config.Tenant.ApiSecretKey, id, oldPass, newPass, newPass}
// 	b, err := json.Marshal(payload)
// 	if err != nil {
// 		return err
// 	}
// 	req, _ := http.NewRequest("POST", fmt.Sprintf(config.Config.Tenant.Url, "Client/ChangePass"), strings.NewReader(fmt.Sprintf("%s", b)))
// 	req.Header.Add("Content-Type", "application/json")
// 	res, _ := http.DefaultClient.Do(req)
// 	defer res.Body.Close()
// 	body, _ := ioutil.ReadAll(res.Body)

// 	if res.StatusCode != 200 {
// 		err = errors.New("CALL_API_FAIL")
// 		return err
// 	} else {
// 		var data models.TenantChangePassBody
// 		err = json.Unmarshal(body, &data)
// 		if err != nil {
// 			return err
// 		}
// 		if data.ErrorMessage == "" {
// 			return nil
// 		}
// 		return errors.New(data.ErrorMessage)
// 	}
// 	return nil
// }
// func Paging(page, total int) models.Paging {
// 	var pageInfo models.Paging
// 	totalPage := int(math.Ceil(float64(total) / float64(config.Config.Paging.Limit)))
// 	pageInfo.Limit = config.Config.Paging.Limit
// 	pageInfo.Total = total
// 	pageInfo.TotalPage = totalPage
// 	if page < 1 {
// 		page = 1
// 	}
// 	if page > totalPage {
// 		page = totalPage
// 	}
// 	pageInfo.Current = page
// 	pageInfo.Skip = (page - 1) * config.Config.Paging.Limit
// 	return pageInfo
// }
