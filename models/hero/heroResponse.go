package hero

type Hero struct {
	Type    string `json:"type"`
	Format  string `json:"format"`
	Version string `json:"version"`
	Data    Data   `json:"data"`
}
type Info struct {
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Magic      int `json:"magic"`
	Difficulty int `json:"difficulty"`
}
type Image struct {
	Full   string `json:"full"`
	Sprite string `json:"sprite"`
	Group  string `json:"group"`
	X      int    `json:"x"`
	Y      int    `json:"y"`
	W      int    `json:"w"`
	H      int    `json:"h"`
}
type Stats struct {
	Hp                   float64 `json:"hp"`
	Hpperlevel           float64 `json:"hpperlevel"`
	Mp                   float64 `json:"mp"`
	Mpperlevel           float64 `json:"mpperlevel"`
	Movespeed            float64 `json:"movespeed"`
	Armor                float64 `json:"armor"`
	Armorperlevel        float64 `json:"armorperlevel"`
	Spellblock           float64 `json:"spellblock"`
	Spellblockperlevel   float64 `json:"spellblockperlevel"`
	Attackrange          float64 `json:"attackrange"`
	Hpregen              float64 `json:"hpregen"`
	Hpregenperlevel      float64 `json:"hpregenperlevel"`
	Mpregen              float64 `json:"mpregen"`
	Mpregenperlevel      float64 `json:"mpregenperlevel"`
	Crit                 float64 `json:"crit"`
	Critperlevel         float64 `json:"critperlevel"`
	Attackdamage         float64 `json:"attackdamage"`
	Attackdamageperlevel float64 `json:"attackdamageperlevel"`
	Attackspeedoffset    float64 `json:"attackspeedoffset"`
	Attackspeedperlevel  float64 `json:"attackspeedperlevel"`
}
type Aatrox struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ahri struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Akali struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Alistar struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Amumu struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Anivia struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Annie struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ashe struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type AurelionSol struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Azir struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Bard struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Blitzcrank struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Brand struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Braum struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Caitlyn struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Camille struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Cassiopeia struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Chogath struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Corki struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Darius struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Diana struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Draven struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type DrMundo struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ekko struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Elise struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Evelynn struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ezreal struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type FiddleSticks struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Fiora struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Fizz struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Galio struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Gangplank struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Garen struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Gnar struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Gragas struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Graves struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Hecarim struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Heimerdinger struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Illaoi struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Irelia struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ivern struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Janna struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type JarvanIV struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Jax struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Jayce struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Jhin struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Jinx struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kalista struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Karma struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Karthus struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kassadin struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Katarina struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kayle struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kennen struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Khazix struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kindred struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Kled struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type KogMaw struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Leblanc struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type LeeSin struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Leona struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Lissandra struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Lucian struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Lulu struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Lux struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Malphite struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Malzahar struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Maokai struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type MasterYi struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type MissFortune struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type MonkeyKing struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Mordekaiser struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Morgana struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nami struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nasus struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nautilus struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nidalee struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nocturne struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Nunu struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Olaf struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Orianna struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Pantheon struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Poppy struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Quinn struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Rammus struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type RekSai struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Renekton struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Rengar struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Riven struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Rumble struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ryze struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Sejuani struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Shaco struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Shen struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Shyvana struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Singed struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Sion struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Sivir struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Skarner struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Sona struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Soraka struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Swain struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Syndra struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type TahmKench struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Taliyah struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Talon struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Taric struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Teemo struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Thresh struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Tristana struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Trundle struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Tryndamere struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type TwistedFate struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Twitch struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Udyr struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Urgot struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Varus struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Vayne struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Veigar struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Velkoz struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Vi struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Viktor struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Vladimir struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Volibear struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Warwick struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Xerath struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type XinZhao struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Yasuo struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Yorick struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Zac struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Zed struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Ziggs struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Zilean struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Zyra struct {
	Version string   `json:"version"`
	ID      string   `json:"id"`
	Key     string   `json:"key"`
	Name    string   `json:"name"`
	Title   string   `json:"title"`
	Blurb   string   `json:"blurb"`
	Info    Info     `json:"info"`
	Image   Image    `json:"image"`
	Tags    []string `json:"tags"`
	Partype string   `json:"partype"`
	Stats   Stats    `json:"stats"`
}
type Data struct {
	Aatrox       Aatrox       `json:"Aatrox"`
	Ahri         Ahri         `json:"Ahri"`
	Akali        Akali        `json:"Akali"`
	Alistar      Alistar      `json:"Alistar"`
	Amumu        Amumu        `json:"Amumu"`
	Anivia       Anivia       `json:"Anivia"`
	Annie        Annie        `json:"Annie"`
	Ashe         Ashe         `json:"Ashe"`
	AurelionSol  AurelionSol  `json:"AurelionSol"`
	Azir         Azir         `json:"Azir"`
	Bard         Bard         `json:"Bard"`
	Blitzcrank   Blitzcrank   `json:"Blitzcrank"`
	Brand        Brand        `json:"Brand"`
	Braum        Braum        `json:"Braum"`
	Caitlyn      Caitlyn      `json:"Caitlyn"`
	Camille      Camille      `json:"Camille"`
	Cassiopeia   Cassiopeia   `json:"Cassiopeia"`
	Chogath      Chogath      `json:"Chogath"`
	Corki        Corki        `json:"Corki"`
	Darius       Darius       `json:"Darius"`
	Diana        Diana        `json:"Diana"`
	Draven       Draven       `json:"Draven"`
	DrMundo      DrMundo      `json:"DrMundo"`
	Ekko         Ekko         `json:"Ekko"`
	Elise        Elise        `json:"Elise"`
	Evelynn      Evelynn      `json:"Evelynn"`
	Ezreal       Ezreal       `json:"Ezreal"`
	FiddleSticks FiddleSticks `json:"FiddleSticks"`
	Fiora        Fiora        `json:"Fiora"`
	Fizz         Fizz         `json:"Fizz"`
	Galio        Galio        `json:"Galio"`
	Gangplank    Gangplank    `json:"Gangplank"`
	Garen        Garen        `json:"Garen"`
	Gnar         Gnar         `json:"Gnar"`
	Gragas       Gragas       `json:"Gragas"`
	Graves       Graves       `json:"Graves"`
	Hecarim      Hecarim      `json:"Hecarim"`
	Heimerdinger Heimerdinger `json:"Heimerdinger"`
	Illaoi       Illaoi       `json:"Illaoi"`
	Irelia       Irelia       `json:"Irelia"`
	Ivern        Ivern        `json:"Ivern"`
	Janna        Janna        `json:"Janna"`
	JarvanIV     JarvanIV     `json:"JarvanIV"`
	Jax          Jax          `json:"Jax"`
	Jayce        Jayce        `json:"Jayce"`
	Jhin         Jhin         `json:"Jhin"`
	Jinx         Jinx         `json:"Jinx"`
	Kalista      Kalista      `json:"Kalista"`
	Karma        Karma        `json:"Karma"`
	Karthus      Karthus      `json:"Karthus"`
	Kassadin     Kassadin     `json:"Kassadin"`
	Katarina     Katarina     `json:"Katarina"`
	Kayle        Kayle        `json:"Kayle"`
	Kennen       Kennen       `json:"Kennen"`
	Khazix       Khazix       `json:"Khazix"`
	Kindred      Kindred      `json:"Kindred"`
	Kled         Kled         `json:"Kled"`
	KogMaw       KogMaw       `json:"KogMaw"`
	Leblanc      Leblanc      `json:"Leblanc"`
	LeeSin       LeeSin       `json:"LeeSin"`
	Leona        Leona        `json:"Leona"`
	Lissandra    Lissandra    `json:"Lissandra"`
	Lucian       Lucian       `json:"Lucian"`
	Lulu         Lulu         `json:"Lulu"`
	Lux          Lux          `json:"Lux"`
	Malphite     Malphite     `json:"Malphite"`
	Malzahar     Malzahar     `json:"Malzahar"`
	Maokai       Maokai       `json:"Maokai"`
	MasterYi     MasterYi     `json:"MasterYi"`
	MissFortune  MissFortune  `json:"MissFortune"`
	MonkeyKing   MonkeyKing   `json:"MonkeyKing"`
	Mordekaiser  Mordekaiser  `json:"Mordekaiser"`
	Morgana      Morgana      `json:"Morgana"`
	Nami         Nami         `json:"Nami"`
	Nasus        Nasus        `json:"Nasus"`
	Nautilus     Nautilus     `json:"Nautilus"`
	Nidalee      Nidalee      `json:"Nidalee"`
	Nocturne     Nocturne     `json:"Nocturne"`
	Nunu         Nunu         `json:"Nunu"`
	Olaf         Olaf         `json:"Olaf"`
	Orianna      Orianna      `json:"Orianna"`
	Pantheon     Pantheon     `json:"Pantheon"`
	Poppy        Poppy        `json:"Poppy"`
	Quinn        Quinn        `json:"Quinn"`
	Rammus       Rammus       `json:"Rammus"`
	RekSai       RekSai       `json:"RekSai"`
	Renekton     Renekton     `json:"Renekton"`
	Rengar       Rengar       `json:"Rengar"`
	Riven        Riven        `json:"Riven"`
	Rumble       Rumble       `json:"Rumble"`
	Ryze         Ryze         `json:"Ryze"`
	Sejuani      Sejuani      `json:"Sejuani"`
	Shaco        Shaco        `json:"Shaco"`
	Shen         Shen         `json:"Shen"`
	Shyvana      Shyvana      `json:"Shyvana"`
	Singed       Singed       `json:"Singed"`
	Sion         Sion         `json:"Sion"`
	Sivir        Sivir        `json:"Sivir"`
	Skarner      Skarner      `json:"Skarner"`
	Sona         Sona         `json:"Sona"`
	Soraka       Soraka       `json:"Soraka"`
	Swain        Swain        `json:"Swain"`
	Syndra       Syndra       `json:"Syndra"`
	TahmKench    TahmKench    `json:"TahmKench"`
	Taliyah      Taliyah      `json:"Taliyah"`
	Talon        Talon        `json:"Talon"`
	Taric        Taric        `json:"Taric"`
	Teemo        Teemo        `json:"Teemo"`
	Thresh       Thresh       `json:"Thresh"`
	Tristana     Tristana     `json:"Tristana"`
	Trundle      Trundle      `json:"Trundle"`
	Tryndamere   Tryndamere   `json:"Tryndamere"`
	TwistedFate  TwistedFate  `json:"TwistedFate"`
	Twitch       Twitch       `json:"Twitch"`
	Udyr         Udyr         `json:"Udyr"`
	Urgot        Urgot        `json:"Urgot"`
	Varus        Varus        `json:"Varus"`
	Vayne        Vayne        `json:"Vayne"`
	Veigar       Veigar       `json:"Veigar"`
	Velkoz       Velkoz       `json:"Velkoz"`
	Vi           Vi           `json:"Vi"`
	Viktor       Viktor       `json:"Viktor"`
	Vladimir     Vladimir     `json:"Vladimir"`
	Volibear     Volibear     `json:"Volibear"`
	Warwick      Warwick      `json:"Warwick"`
	Xerath       Xerath       `json:"Xerath"`
	XinZhao      XinZhao      `json:"XinZhao"`
	Yasuo        Yasuo        `json:"Yasuo"`
	Yorick       Yorick       `json:"Yorick"`
	Zac          Zac          `json:"Zac"`
	Zed          Zed          `json:"Zed"`
	Ziggs        Ziggs        `json:"Ziggs"`
	Zilean       Zilean       `json:"Zilean"`
	Zyra         Zyra         `json:"Zyra"`
}