package params

import (
	"fmt"
	"io/ioutil"
	"strings"
)

const DEFULT_PORT int = 8000

/*
title = 'Directory listing for %s' % displaypath
        r.append('<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" '
                 '"http://www.w3.org/TR/html4/strict.dtd">')
        r.append('<html>\n<head>')
        r.append('<meta http-equiv="Content-Type" '
                 'content="text/html; charset=%s">' % enc)
        r.append('<title>%s</title>\n</head>' % title)
        r.append('<body>\n<h1>%s</h1>' % title)
        r.append('<hr>\n<ul>')
        for name in list:
            fullname = os.path.join(path, name)
            displayname = linkname = name
            # Append / for directories or @ for symbolic links
            if os.path.isdir(fullname):
                displayname = name + "/"
                linkname = name + "/"
            if os.path.islink(fullname):
                displayname = name + "@"
                # Note: a link to a directory displays with @ and links with /
            r.append('<li><a href="%s">%s</a></li>'
                    % (urllib.parse.quote(linkname,
                                          errors='surrogatepass'),
                       html.escape(displayname)))
        r.append('</ul>\n<hr>\n</body>\n</html>\n')
        encoded = '\n'.join(r).encode(enc, 'surrogateescape')
        f = io.BytesIO()
        f.write(encoded)
        f.seek(0)
        self.send_response(HTTPStatus.OK)
        self.send_header("Content-type", "text/html; charset=%s" % enc)
        self.send_header("Content-Length", str(len(encoded)))
        self.end_headers()
        return f

*/
func SimpleListDirPage(displaypath string, dirPath string) string {
	title := fmt.Sprintf("Directory listing for %s", displaypath)

	bufferWriter := strings.Builder{}
	bufferWriter.WriteString("<!DOCTYPE HTML PUBLIC \"" +
		"-//W3C//DTD HTML 4.01//EN \"http://www.w3.org/TR/html4/strict.dtd\">")
	bufferWriter.WriteString("<html>\n<head>")
	bufferWriter.WriteString("<link rel=\"shortcut icon\" " +
		"href=\"data:image/x-icon;,\" type=\"image/x-icon\"> ")
	bufferWriter.WriteString("<meta http-equiv=\"Content-Type\" " +
		"content=\"text/html; charset=utf-8\"")
	bufferWriter.WriteString("<title></title>\n</head>")
	bufferWriter.WriteString(fmt.Sprintf("<body>\n<h1>%s</h1>", title))
	bufferWriter.WriteString("<hr>\n<ul>")
	//file list
	infos, err := ioutil.ReadDir(dirPath)
	if err == nil {

		for _, info := range infos {
			name := info.Name()
			bufferWriter.WriteString(
				fmt.Sprintf("<li><a href=\"%s\">%s</a></li>", name, name))
		}

	} else {
		bufferWriter.WriteString("<li><h1>Read Error</h1></li>")
	}

	bufferWriter.WriteString("</ul>\n<hr>\n</body>\n</html>\n")

	return bufferWriter.String()

	//writer := bufio.NewWriter(&strings.Builder{})

}
