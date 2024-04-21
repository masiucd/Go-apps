# File upload2.0

## Description

This is a simple file upload application. It allows users to upload files and view them. The files are stored in a database and can be downloaded by clicking on the file name.
To store the file in a database we need to convert the file into a byte array. This byte array is then stored in the database. When the file is downloaded, the byte array is converted back to the file and downloaded.

### Routes

- `/` - Home page
- `/upload` - Upload page
- `/attachments` - Files page
- `/files/:id` - View file page (id is the file id)

#### Tools

- [GO](https://golang.org/)
- [GORM](https://gorm.io/)
- [Go Fiber](https://gofiber.io/)
- [SQL Lite](https://www.sqlite.org/index.html)
