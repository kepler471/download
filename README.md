# download

download is a command line batch file downloader written in Go.

Given a file type, and a list of urls, the software will download all files on a given webpage matching the specified file type.

If files are found, each url will have a directory, and a sub-directory for a file type, where the files are downloaded to.

Future functionality:
  - show which files have been found, and their download size, and ask for confirmation
  - concurrent downloads
