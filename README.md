# download

download is a command line batch file downloader written in Go.

Given a file type, and a list of urls, the software will download all files on a given webpage matching the specified file type.

If files are found, each url will have a directory, and a sub-directory for a file type, where the files are downloaded to.
  
download will show which files have been found, their download size, and ask for confirmation.
  
Usage:
    `download -y -t pdf www.site.com/page/with/files.html www.anothersite.com/page/with/files.html`  
    
Flags:
  - `-y` bool. Add to automatically download all files, ignoring user confirmation.
    
    default (no flag): `false` 
  - `-t` string. Specify file extension 
  
    default (no flag): `pdf`

Future functionality:
  - concurrent downloads
  - suggest file types found in page
  - option to specify download directory
  - option to create sub-directory for each URL  
