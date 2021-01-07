package shell

import (
	"github.com/fyuan1316/operatorlib/task"
	"path/filepath"
	"strings"
)

func (s *ScriptManager) LoadDir(dir string) error {
	var (
		files map[string]string
		err   error
	)
	if files, err = task.GetFilesInFolder(dir, task.Asc(), task.Suffix(".sh")); err != nil {
		return err
	}
	if err = s.LoadContents(files); err != nil {
		return err
	}
	s.Dir = dir
	return nil
}

func (s *ScriptManager) LoadContents(files map[string]string) error {
	for path := range files {
		if err := s.Load(path); err != nil {
			return err
		}
	}
	return nil
}

func (s *ScriptManager) Load(filePath string) error {
	if err := s.ensureExecutable(filePath); err != nil {
		return err
	}
	lowercaseFilePath := strings.ToLower(filePath)
	_, fileName := filepath.Split(lowercaseFilePath)

	if strings.HasPrefix(fileName, "precheck") {
		s.preCheck = append(s.preCheck, filePath)
	} else if strings.HasPrefix(fileName, "postcheck") {
		s.postCheck = append(s.postCheck, filePath)
	} else if strings.HasPrefix(fileName, "prerun") {
		s.preRun = append(s.preRun, filePath)
	} else if strings.HasPrefix(fileName, "postrun") {
		s.postRun = append(s.postRun, filePath)
	} else {
		s.run = append(s.run, filePath)
	}
	return nil
}
