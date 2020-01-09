package imageinfo

type Bing struct {
	Images []Image `images`
}

type Image struct {
	Enddate   string `enddate`
	Url       string `url`
	Copyright string `copyright`
	Hsh       string `hsh`
}

type ImageInfo struct {
	ImageName string
	ImagePath string
	LocalSize int64
	LocalHash string
	ServerSize int64
	ServerHash string
	Desc string
	DownloadUrl string
	DownloadResult int
	Err error
}
