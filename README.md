go-photo-sort

![Go](https://github.com/sebfoucault/go-photo-sort/workflows/Build%20and%20Test/badge.svg)

## Supported patterns for date

In the table below, examples are based on the following date: "2006:01:02 15:04:05"

| Pattern | Explanation                   | Example |
| ------- | ----------------------------- | ------- |
| dd      | Day in month with two digits  | 02      |
| MM      | Month in year with two digits | 01      |
| MMM     | Month in year with short name | Jan     |
| MMMM    | Month in year with full  name | January |
| yyyy    | Year with four digits         | 2006    |
| yy      | Year with two digits          | 06      |
| HH      | Hour in day (0-23)            | 15      |
| hh      | Hour in day (1-12)            | 03      |
| mm      | Minute in hour                | 04      |
| ss      | Second in minute              | 05      |
| a       | Am/pm marker (lower case)     | pm      |
| EE      | Day in week with short name   | Mon     |
| EEE     | Day in week with short name   | Mon     |
| EEEE    | Day in week with long name    | Monday  |

## Credits

gops makes use of the following open source projects:

| Project                 | URL                                                |
| ----------------------- | -------------------------------------------------- |
| cli                     | https://github.com/urfave/cli                      |
| gocrest                 | https://github.com/corbym/gocrest                  |
| go-exif                 | https://github.com/dsoprea/go-exif                 |
| go-jpeg-image-structure | https://github.com/dsoprea/go-jpeg-image-structure |
| go-logging              | https://github.com/dsoprea/go-logging              |
| go-png-image-structure  | https://github.com/dsoprea/go-png-image-structure  |

Thanks a lot to all the contributors of these projects!
