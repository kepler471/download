# download

download is a command line batch file downloader written in Go.

Given a list of urls, download will show which files have been found, their download size, and ask for confirmation.
  
### Usage:

`download -y -t pdf www.site.com/page/with/files.html www.anothersite.com/page/with/files.html`  
    
### Flags:
  - `-t` string. Specify file extension 
  
    default (no flag): `pdf`
    
  - `-y` bool. Add to automatically download all files, ignoring user confirmation.
    
    default (no flag): `false` 

##### Future functionality:
  - concurrent downloads
  - suggest file types found in page
  - option to specify download directory
  - option to create sub-directory for each URL  
