# download

download is a command line batch file downloader written in Go.

Given a list of urls, download will show which files have been found, their download size, and ask for confirmation.
  
### Usage:

`download -y -t pdf www.site.com/page/with/files.html www.anothersite.com/page/with/files.html`  
    
### Flags:
  - `-t` string. Specify file extension 
  
    default: `pdf`
    
  - `-l` bool. List files only. Overrides all bool flags
  
    default: `false`
    
  - `-y` bool. Add to automatically download all files, ignoring user confirmation.
    
    default: `false` 

##### TODO:
- concurrent downloads
- suggest file types found in page
- option to specify download directory
- option to create sub-directory for each URL
- install script
- want to handle URL endings, eg .html, .htm ...
- handle even when 0 urls provided with better feedback
- provide help, default for no args
- replace logs with errors
