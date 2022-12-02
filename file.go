package log

import (
	"os"
)

// LogToFile sets up default logging to log to a file. This is helpful as we
// can't print to the terminal since our TUI is occupying it. If the file
// doesn't exist it will be created.
//
// Don't forget to close the file when you're done with it.
//
//	  f, err := LogToFile("debug.log", "debug")
//	  if err != nil {
//			fmt.Println("fatal:", err)
//			os.Exit(1)
//	  }
//	  defer f.Close()
func LogToFile(path string, prefix string) (*os.File, error) {
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0o644)
	if err != nil {
		return nil, err
	}
	SetLogger(NewDefault(f))
	if prefix != "" {
		SetFields("prefix", prefix)
	}

	return f, nil
}
