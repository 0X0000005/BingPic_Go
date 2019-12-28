package service

type Bing struct {
	Images []Image `images`
}

type Image struct {
	Enddate   string `enddate`
	Url       string `url`
	Copyright string `copyright`
	Hsh       string `hsh`
}
