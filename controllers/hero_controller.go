package controllers

import (
	"bmg/constants"
	"bmg/models/hero"
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"github.com/labstack/echo/v4"
	// "github.com/tidwall/gjson"
)

type data struct {
	Number int `json:"number"`
	People []people `json:"people"`
	Message string `json:"message"`
}
type people struct {
	Craft string `json:"craft"`
	Name string `json:"name"`
}

func Tes(c echo.Context) error {
	url := "http://api.open-notify.org/astros.json"

	spaceClient := http.Client{
		Timeout: time.Second * 2, // Timeout after 2 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		// log.Fatal(err)
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error a",
			err.Error(),
		))
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		// log.Fatal(getErr)
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error b",
			getErr.Error(),
		))
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		// log.Fatal(readErr)
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error c",
			readErr.Error(),
		))
	}

	data1 := data{}
	jsonErr := json.Unmarshal([]byte(body), &data1)
	if jsonErr != nil {
		// log.Fatal(jsonErr)
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error d",
			jsonErr.Error(),
		))
	}

	// fmt.Println(people1.Number)
	return c.JSON(http.StatusCreated, BaseResponse(
		http.StatusCreated,
		"Success",
		data1.People,
	))
}

func HeroController(c echo.Context) error {

	spaceClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, constants.URL_LOL, nil)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error",
			err.Error(),
		))
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error",
			getErr.Error(),
		))
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error",
			readErr.Error(),
		))
	}

	data := hero.Hero{}
	jsonErr := json.Unmarshal([]byte(body), &data)
	if jsonErr != nil {
		return c.JSON(http.StatusInternalServerError, BaseResponse(
			http.StatusInternalServerError,
			"error",
			jsonErr.Error(),
		))
	}

	heroName := GetHero(c.QueryParam("name"),data.Data)

	return c.JSON(http.StatusOK, heroName)
}

func GetHero(name string, hero hero.Data ) interface{} {

	re := regexp.MustCompile(name)
	status := false

	heroList := [...]string{"Aatrox","Ahri","Akali","Alistar","Amumu","Anivia","Annie","Ashe","AurelionSol","Azir","Bard",
	"Blitzcrank","Brand","Braum","Caitlyn","Camille","Cassiopeia","Chogath","Corki","Darius","Diana","Draven","DrMund",
	"Ekko","Elise","Evelynn","Ezreal","FiddleSticks","Fiora","Fizz","Galio","Gangplank","Garen","Gnar","Gragas","Graves",
	"Hecarim","Heimerdinger","Illaoi","Irelia","Ivern","Janna","JarvanIV","Jax","Jayce","Jhin","Jinx","Kalista","Karma",
	"Karthus","Kassadin","Katarina","Kayle","Kennen","Khazix","Kindred","Kled","KogMaw","Leblanc","LeeSin","Leona",
	"Lissandra","Lucian","Lulu","Lux","Malphite","Malzahar","Maokai","MasterYi","MissFortune","MonkeyKing","Mordekaiser",
	"Morgana","Nami","Nasus","Nautilus","Nidalee","Nocturne","Nunu","Olaf","Orianna","Pantheon","Poppy","Quinn","Rammus",
	"RekSai","Renekton","Rengar","Riven","Rumble","Ryze","Sejuani","Shaco","Shen","Shyvana","Singed","Sion","Sivir",
	"Skarner","Sona","Soraka","Swain","Syndra","TahmKench","Taliyah","Talon","Taric","Teemo","Thresh","Tristana","Trundle",
	"Tryndamere","TwistedFate","Twitch","Udyr","Urgot","Varus","Vayne","Veigar","Velkoz","Vi","Viktor","Vladimir","Volibear",
	"Warwick","Xerath","XinZhao","Yasuo","Yorick","Zac","Zed","Ziggs","Zilean","Zyra"}

	for _, value := range heroList{
		status = re.Match([]byte(value))
		if status {

			switch value {
				case heroList[0]:
					return hero.Aatrox
				case heroList[1]:
					return hero.Ahri
				case heroList[2]:
					return hero.Akali
				case heroList[3]:
					return hero.Alistar
				case heroList[4]:
					return hero.Amumu
				case heroList[5]:
					return hero.Anivia
				case heroList[6]:
					return hero.Annie
				case heroList[7]:
					return hero.Ashe
				case heroList[8]:
					return hero.AurelionSol
				case heroList[9]:
					return hero.Azir
				case heroList[10]:
					return hero.Bard
				case heroList[11]:
					return hero.Blitzcrank
				case heroList[12]:
					return hero.Brand
				case heroList[13]:
					return hero.Braum
				case heroList[14]:
					return hero.Caitlyn
				case heroList[15]:
					return hero.Camille
				case heroList[16]:
					return hero.Cassiopeia
				case heroList[17]:
					return hero.Chogath
				case heroList[18]:
					return hero.Corki
				case heroList[19]:
					return hero.Darius
				case heroList[20]:
					return hero.Diana
				case heroList[21]:
					return hero.DrMundo
				case heroList[22]:
					return hero.Draven
				case heroList[23]:
					return hero.Ekko
				case heroList[24]:
					return hero.Elise
				case heroList[25]:
					return hero.Evelynn
				case heroList[26]:
					return hero.Ezreal
				case heroList[27]:
					return hero.FiddleSticks
				case heroList[28]:
					return hero.Fiora
				case heroList[29]:
					return hero.Fizz
				case heroList[30]:
					return hero.Galio
				case heroList[31]:
					return hero.Gangplank
				case heroList[32]:
					return hero.Garen
				case heroList[33]:
					return hero.Gnar
				case heroList[34]:
					return hero.Gragas
				case heroList[35]:
					return hero.Graves
				case heroList[36]:
					return hero.Hecarim
				case heroList[37]:
					return hero.Heimerdinger
				case heroList[38]:
					return hero.Illaoi
				case heroList[39]:
					return hero.Irelia
				case heroList[40]:
					return hero.Ivern
				case heroList[41]:
					return hero.Janna
				case heroList[42]:
					return hero.JarvanIV
				case heroList[43]:
					return hero.Jax
				case heroList[44]:
					return hero.Jayce
				case heroList[45]:
					return hero.Jhin
				case heroList[46]:
					return hero.Jinx
				case heroList[47]:
					return hero.Kalista
				case heroList[48]:
					return hero.Karma
				case heroList[49]:
					return hero.Karthus
				case heroList[50]:
					return hero.Kassadin
				case heroList[51]:
					return hero.Katarina
				case heroList[52]:
					return hero.Kayle
				case heroList[53]:
					return hero.Kennen
				case heroList[54]:
					return hero.Khazix
				case heroList[55]:
					return hero.Kindred
				case heroList[56]:
					return hero.Kled
				case heroList[57]:
					return hero.KogMaw
				case heroList[58]:
					return hero.Leblanc
				case heroList[59]:
					return hero.LeeSin
				case heroList[60]:
					return hero.Leona
				case heroList[61]:
					return hero.Lissandra
				case heroList[62]:
					return hero.Lucian
				case heroList[63]:
					return hero.Lulu
				case heroList[64]:
					return hero.Lux
				case heroList[65]:
					return hero.Malphite
				case heroList[66]:
					return hero.Malzahar
				case heroList[67]:
					return hero.Maokai
				case heroList[68]:
					return hero.MasterYi
				case heroList[69]:
					return hero.MissFortune
				case heroList[70]:
					return hero.MonkeyKing
				case heroList[71]:
					return hero.Mordekaiser
				case heroList[72]:
					return hero.Morgana
				case heroList[73]:
					return hero.Nami
				case heroList[74]:
					return hero.Nasus
				case heroList[75]:
					return hero.Nautilus
				case heroList[76]:
					return hero.Nidalee
				case heroList[77]:
					return hero.Nocturne
				case heroList[78]:
					return hero.Nunu
				case heroList[79]:
					return hero.Olaf
				case heroList[80]:
					return hero.Orianna
				case heroList[81]:
					return hero.Pantheon
				case heroList[82]:
					return hero.Poppy
				case heroList[83]:
					return hero.Quinn
				case heroList[84]:
					return hero.Rammus
				case heroList[85]:
					return hero.RekSai
				case heroList[86]:
					return hero.Renekton
				case heroList[87]:
					return hero.Rengar
				case heroList[88]:
					return hero.Riven
				case heroList[89]:
					return hero.Rumble
				case heroList[90]:
					return hero.Ryze
				case heroList[91]:
					return hero.Sejuani
				case heroList[92]:
					return hero.Shaco
				case heroList[93]:
					return hero.Shen
				case heroList[94]:
					return hero.Shyvana
				case heroList[95]:
					return hero.Singed
				case heroList[96]:
					return hero.Sion
				case heroList[97]:
					return hero.Sivir
				case heroList[98]:
					return hero.Skarner
				case heroList[99]:
					return hero.Sona
				case heroList[100]:
					return hero.Soraka
				case heroList[101]:
					return hero.Swain
				case heroList[102]:
					return hero.Syndra
				case heroList[103]:
					return hero.TahmKench
				case heroList[104]:
					return hero.Taliyah
				case heroList[105]:
					return hero.Talon
				case heroList[106]:
					return hero.Taric
				case heroList[107]:
					return hero.Teemo
				case heroList[108]:
					return hero.Thresh
				case heroList[109]:
					return hero.Tristana
				case heroList[110]:
					return hero.Trundle
				case heroList[111]:
					return hero.Tryndamere
				case heroList[112]:
					return hero.TwistedFate
				case heroList[113]:
					return hero.Twitch
				case heroList[114]:
					return hero.Udyr
				case heroList[115]:
					return hero.Udyr
				case heroList[116]:
					return hero.Varus
				case heroList[117]:
					return hero.Vayne
				case heroList[118]:
					return hero.Veigar
				case heroList[119]:
					return hero.Velkoz
				case heroList[120]:
					return hero.Vi
				case heroList[121]:
					return hero.Viktor
				case heroList[122]:
					return hero.Vladimir
				case heroList[123]:
					return hero.Volibear
				case heroList[124]:
					return hero.Warwick
				case heroList[125]:
					return hero.Xerath
				case heroList[126]:
					return hero.XinZhao
				case heroList[127]:
					return hero.Yasuo
				case heroList[128]:
					return hero.Yorick
				case heroList[129]:
					return hero.Zac
				case heroList[130]:
					return hero.Zed
				case heroList[131]:
					return hero.Ziggs
				case heroList[132]:
					return hero.Zilean
				case heroList[133]:
					return hero.Zyra
			}
		}
	}

	return ""
}