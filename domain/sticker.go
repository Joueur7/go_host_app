package domain


type Sticker struct {
    ID        int    `json:"id"`
    Name      string `json:"name"`
    ImageURL  string `json:"image_url"`
    Priority  int    `json:"-"`
    StartTime string `json:"-"`
    EndTime   string `json:"-"`
}

