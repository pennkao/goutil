package exthtml

import (
	"bytes"
	"fmt"
	"regexp"
)

type htmlParser struct {
	reg      *regexp.Regexp
	url      string
	content  []byte
	partten  string
	replaces map[string]string
}

func NewHtmlParser() *htmlParser {
	htmlParser := &htmlParser{
		replaces: map[string]string{
			`\s+`:            " ", //过滤多余回车
			`<[ ]+`:          "<", //过滤<__("<"号后面带空格)
			`<\!–.*?–>`:      "",  // //注释
			`<(\!.*?)>`:      "",  //过滤DOCTYPE
			`<(\/?html.*?)>`: "",  //过滤html标签
			`<(\/?br.*?)>`:   "",  //过滤br标签
			`<(\/?head.*?)>`: "",  //过滤head标签
			// `<(\/?meta.*?)>`: "",                    //过滤meta标签
			`<(\/?body.*?)>`:                    "", //过滤body标签
			`<(\/?link.*?)>`:                    "", //过滤link标签
			`<(\/?form.*?)>`:                    "", //过滤form标签
			`<(applet.*?)>(.*?)<(\/applet.*?)>`: "", //过滤applet标签
			`<(\/?applet.*?)>`:                  "",
			`<(style.*?)>(.*?)<(\/style.*?)>`:   "", //过滤style标签
			`<(\/?style.*?)>`:                   "",
			// `<(title.*?)>(.*?)<(\/title.*?)>`: "", //过滤title标签
			// `<(\/?title.*?)>`: "",
			`<(object.*?)>(.*?)<(\/object.*?)>`:     "", //过滤object标签
			`<(\/?objec.*?)>`:                       "",
			`<(noframes.*?)>(.*?)<(\/noframes.*?)>`: "", //过滤noframes标签
			`<(\/?noframes.*?)>`:                    "",
			`<(i?frame.*?)>(.*?)<(\/i?frame.*?)>`:   "", //过滤frame标签
			`<(noscript.*?)>(.*?)<(\/noscript.*?)>`: "", //过滤noframes标签
			// `on([a-z]+)\s*="(.*?)"`: "",                 //过滤dom事件
			// `on([a-z]+)\s*='(.*?)'`: "",
		},
	}
	defer htmlParser.Close()
	return htmlParser
}

func (hp *htmlParser) CleanScript() *htmlParser {
	hp.replaces[`<(script.*?)>(.*?)<(\/script.*?)>`] = ""
	hp.replaces[`<(\/?script.*?)>`] = ""
	return hp
}

func (hp *htmlParser) CleanAllHtml() *htmlParser {
	hp.replaces[`<[^>]*>`] = ""
	return hp
}

func (hp *htmlParser) IsGbk() bool {
	d := bytes.ToLower(hp.content)
	if bytes.Index(d, []byte(`text/html; charset=gbk`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`text/html; charset="gbk"`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`text/html; charset=gb2312`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`text/html; charset="gb2312"`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`charset=gbk`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`charset="gbk"`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`charset="gb2312"`)) > 0 {
		return true
	}
	if bytes.Index(d, []byte(`charset=gb2312`)) > 0 {
		return true
	}
	return false
}

func (hp *htmlParser) Close() {
	hp.content = nil
}

func (hp *htmlParser) LoadData(content []byte) *htmlParser {
	hp.content = content
	return hp
}

func (hp *htmlParser) Content() []byte {
	return hp.content
}

func (hp *htmlParser) Replace() *htmlParser {
	for p, r := range hp.replaces {
		reg := regexp.MustCompile(p)
		hp.content = []byte(reg.ReplaceAllLiteral(hp.content, []byte(r)))
	}
	return hp
}

func (hp *htmlParser) Partten(p string) *htmlParser {
	hp.partten = p
	return hp
}

func (hp *htmlParser) FindSubmatch() [][]byte {
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindSubmatch(hp.content)
}

func (hp *htmlParser) FindAllSubmatch() [][][]byte {
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindAllSubmatch(hp.content, -1)
}

func (hp *htmlParser) FindAllString() []string {
	hp.reg = regexp.MustCompile(hp.partten)
	return hp.reg.FindAllString(fmt.Sprintf("%s", hp.content), -1)
}

func (hp *htmlParser) FindByAttr(tagName, attr, value string) [][][]byte {
	hp.partten = fmt.Sprintf(`((?U)<%s+.*%s=['"]%s['"]+.*>(.*)</%s>).*?`, tagName, attr, value, tagName)
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindAllSubmatch(hp.content, -1)
}

func (hp *htmlParser) FindByTagName(tagName string) [][][]byte {
	hp.partten = fmt.Sprintf(`((?U)<%s+.*>(.*)</%s>).*?`, tagName, tagName)
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindAllSubmatch(hp.content, -1)
}

func (hp *htmlParser) FindJsonStr(nodeName string) [][][]byte {
	hp.partten = fmt.Sprintf(`(?U)"%s":\s*?['"](.*)['"]`, nodeName)
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindAllSubmatch(hp.content, -1)
}

func (hp *htmlParser) FindJsonInt(nodeName string) [][][]byte {
	hp.partten = fmt.Sprintf(`(?U)"%s":(.*),`, nodeName)
	hp.reg = regexp.MustCompile(hp.partten)
	// fmt.Println(re.String())
	return hp.reg.FindAllSubmatch(hp.content, -1)
}
