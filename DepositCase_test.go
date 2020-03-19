package ManDepositServer

import (
	"encoding/json"
	"github.com/MatrixAINetwork/go-AIMan/manager"
	"math/big"
	"fmt"
	"os"
	"testing"
)

func TestSendDeposit100_v2(t *testing.T) {
	/*addressList := []string{
	"MAN.3vgdE1Tc4cVPCT88c1WWMVn1z5riL",
	"MAN.2J5ATuHNFfHF9B8iu1Hn2AbaqxVqu",
	"MAN.MoZUPLomob9bd1nEeeUMtb5EXM9x",
	"MAN.3STdjomHfF4rtsWtXzArjvXyjvszN",
	"MAN.3rrog5tmMioUxbUshJeCGeedMWY1M",
	"MAN.iBrGVPXpbQ6ej8FUoxMf3AUDXYYD",
	"MAN.XQT4SsmrgziXRejVDsYv6zTTypV3",
	"MAN.4An8q24TbjLGnjeWNhFrKqzjTae5o",
	"MAN.4TpXxjZLsXgF8qMeZQ6QrLciJYDg9",
	"MAN.2H45AdHNS8Rtkt36iW8Jz8tuB169E",
	"MAN.3HJNfsStCtUkacdkRpvEfub8FbKuj",
	"MAN.2afaRjkBtvDNYCHpBBj5c7xECCsH3",
	"MAN.288L4cWmxkRjj1THeMnDuRm2qW1Lz",
	"MAN.2AD7R32w8nLHFm7Bwzz7npnjeYchD",
	"MAN.3A9orsZJK9eMSzBtfXxtyjJrwhWXE",
	"MAN.pYHddAwf2Q9zre1HCh9SYhKjXxuh",
	"MAN.42UM8CGLCnBUvFGQtQSe6Rmi9SQoR",
	"MAN.4SxJ3jGcQ1TFW9vqk4Hn7L6R7uEpQ",
	"MAN.3BCs5TsgLbCsFT7z8RJvpwof2JTrH",
	"MAN.Myy37DXgwYszjGraeFnssyi7PWXS",
	"MAN.MmVPJ8pRxr8tWBJjzrcFQfrTwY6E",
	"MAN.BoWYwSD4t38QDDJexMttLoovkujT",
	"MAN.24WMGuB6oFtBpbvfR5nt8FtfaNXM",
	"MAN.3gYaMQDJoWb9VL3exZEjzrBQKVEbe",
	"MAN.2QdycWJysinkzyPpvPxUoxXrzgoH",
	"MAN.JKBzEnBwUEsjY8E1bcyD1D4eAi1Q",
	"MAN.5uzYbSjbRAjACKc8Bm9QL4FnoNJ5",
	"MAN.wNYfiNtNrjv2tknFkchkFJwbFm9Y",
	"MAN.2CQgoEvL1hRR7UGkY6QuwqQfxVvZS",
	"MAN.3gpRaEUy7jdGVTj42ibsAzvwWHfRx",
	"MAN.3zWWQ4z53RGkQpvaNXkyFjboDBKwr",
	"MAN.3KT4oucei91ggGY4o7prdxKgFfEKg",
	"MAN.2Vww4jjw29wWwtq7R1YDvddu2eV1g",
	"MAN.44EbaBtGrS7pHTWgmS62QVRFxMX1G",
	"MAN.21cwqk8d4pzjMQiiegPa1oB5uj4PR",
	"MAN.nh2Wbbj1qgLDjXFduN1FfhdCRmhH",
	"MAN.3vFGSoui9oDmf1DFpLkpJ5YDXoTgb",
	"MAN.2U6D3nu9VEyHD9m1zgdiLGwgCmpNP",
	"MAN.HGgoyWjd9ciJc49LByaT2qTPgE4G",
	"MAN.2JbEewQFpcZuKmSL3ph8yiFJPs4Hv",
	"MAN.CQg1CkeRH9iL9e1FHAumhPkCTZM2",
	"MAN.3iifiELQnSJKD95jfTHCVFLS2C2a2",
	"MAN.46ABUWjTiXErkvPWhT26veiC6QJPt",
	"MAN.QpY4muk4zFMoTPnfhAHWjBv91DGA",
	"MAN.3AAHsKoSZRfTqknUPxJwvsa4Q5wE5",
	"MAN.2HJ25jJbvZreZnL36FfdiZKMY4T8p",
	"MAN.4PFK5q6Mj8b6RVx1Q7fJGCTacGPs6",
	"MAN.f1TEdrfUB514jrXCRUKCmZFAm8Tc",
	"MAN.3BpgxNyUMnt8srWMaZphHGQMpYD3u",
	"MAN.3My3ayEurjNxB22Hp8B6Naqnec7op"}*/
	/*addressList := []string{
	"MAN.3gYaMQDJoWb9VL3exZEjzrBQKVEbe",
	"MAN.3KT4oucei91ggGY4o7prdxKgFfEKg",
	"MAN.2afaRjkBtvDNYCHpBBj5c7xECCsH3"}*/

	/*addressList := []string{
	"MAN.2AD7R32w8nLHFm7Bwzz7npnjeYchD",
	"MAN.2U6D3nu9VEyHD9m1zgdiLGwgCmpNP"}*/

	addressList := []string{
		"MAN.3vgdE1Tc4cVPCT88c1WWMVn1z5riL",
		"MAN.2J5ATuHNFfHF9B8iu1Hn2AbaqxVqu",
		"MAN.MoZUPLomob9bd1nEeeUMtb5EXM9x",
		"MAN.3STdjomHfF4rtsWtXzArjvXyjvszN",
		"MAN.3rrog5tmMioUxbUshJeCGeedMWY1M",
		"MAN.iBrGVPXpbQ6ej8FUoxMf3AUDXYYD",
		"MAN.XQT4SsmrgziXRejVDsYv6zTTypV3",
		"MAN.4An8q24TbjLGnjeWNhFrKqzjTae5o",
		"MAN.4TpXxjZLsXgF8qMeZQ6QrLciJYDg9",
		"MAN.2H45AdHNS8Rtkt36iW8Jz8tuB169E",
		"MAN.3HJNfsStCtUkacdkRpvEfub8FbKuj",
		"MAN.288L4cWmxkRjj1THeMnDuRm2qW1Lz",
		"MAN.3A9orsZJK9eMSzBtfXxtyjJrwhWXE",
		"MAN.pYHddAwf2Q9zre1HCh9SYhKjXxuh",
		"MAN.42UM8CGLCnBUvFGQtQSe6Rmi9SQoR",
		"MAN.4SxJ3jGcQ1TFW9vqk4Hn7L6R7uEpQ",
		"MAN.3BCs5TsgLbCsFT7z8RJvpwof2JTrH",
		"MAN.Myy37DXgwYszjGraeFnssyi7PWXS",
		"MAN.MmVPJ8pRxr8tWBJjzrcFQfrTwY6E",
		"MAN.BoWYwSD4t38QDDJexMttLoovkujT",
		"MAN.24WMGuB6oFtBpbvfR5nt8FtfaNXM",
		"MAN.2QdycWJysinkzyPpvPxUoxXrzgoH",
		"MAN.JKBzEnBwUEsjY8E1bcyD1D4eAi1Q",
		"MAN.5uzYbSjbRAjACKc8Bm9QL4FnoNJ5",
		"MAN.wNYfiNtNrjv2tknFkchkFJwbFm9Y",
		"MAN.2CQgoEvL1hRR7UGkY6QuwqQfxVvZS",
		"MAN.3gpRaEUy7jdGVTj42ibsAzvwWHfRx",
		"MAN.3zWWQ4z53RGkQpvaNXkyFjboDBKwr",
		"MAN.2Vww4jjw29wWwtq7R1YDvddu2eV1g",
		"MAN.44EbaBtGrS7pHTWgmS62QVRFxMX1G",
		"MAN.21cwqk8d4pzjMQiiegPa1oB5uj4PR",
		"MAN.nh2Wbbj1qgLDjXFduN1FfhdCRmhH",
		"MAN.3vFGSoui9oDmf1DFpLkpJ5YDXoTgb",
		"MAN.HGgoyWjd9ciJc49LByaT2qTPgE4G",
		"MAN.2JbEewQFpcZuKmSL3ph8yiFJPs4Hv",
		"MAN.CQg1CkeRH9iL9e1FHAumhPkCTZM2",
		"MAN.3iifiELQnSJKD95jfTHCVFLS2C2a2",
		"MAN.46ABUWjTiXErkvPWhT26veiC6QJPt",
		"MAN.QpY4muk4zFMoTPnfhAHWjBv91DGA",
		"MAN.3AAHsKoSZRfTqknUPxJwvsa4Q5wE5",
		"MAN.2HJ25jJbvZreZnL36FfdiZKMY4T8p",
		"MAN.4PFK5q6Mj8b6RVx1Q7fJGCTacGPs6",
		"MAN.f1TEdrfUB514jrXCRUKCmZFAm8Tc",
		"MAN.3BpgxNyUMnt8srWMaZphHGQMpYD3u",
		"MAN.3My3ayEurjNxB22Hp8B6Naqnec7op"}

	Man1 := new(big.Int).SetUint64(1e18)
	amount := new(big.Int).SetUint64(10000)
	amount = amount.Mul(amount, Man1)
	depositType := big.NewInt(0) //0-活期,1-定期1个月,3-定期3个月,6-定期6个月
	depositRole := big.NewInt(1) // 1是矿工，2是验证者

	ReadPassword := func(filePth string) (string, error) {
		f, err := os.Open(filePth)
		if err != nil {
			return "", err
		}

		var X struct {
			Password string `json:"Password"`
		}
		//创建一个解码器
		decoder := json.NewDecoder(f)

		err = decoder.Decode(&X)
		if err != nil {
			return "", err
		}

		return X.Password, nil
	}
	for _, addressA0 := range addressList {
		//filePath := "C:/Users/liyong/Desktop/Create/keystore/Balance/" + addressA0 + "/in.json"
		filePath := "./keystore/Balance/" + addressA0 + "/in.json"
		accountPassword, err := ReadPassword(filePath)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}

		fmt.Println("addr", addressA0, "password", accountPassword)
		err = SendDepositTrans_v2(manager.Jerry_Manager, addressA0,
			amount, depositType, depositRole, addressA0, accountPassword)
		if err != nil {
			t.Error(err)
			t.FailNow()
		}
	}

}
