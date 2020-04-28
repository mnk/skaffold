/*
Copyright 2019 The Skaffold Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package survey

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/GoogleContainerTools/skaffold/pkg/skaffold/config"
)

const (
	Prompt = `Help improve Skaffold! Take a 10-second anonymous survey by running
   $skaffold survey`

	Form = `Thank you for offering your feedback on Skaffold! Understanding your experiences and opinions helps us make Skaffold better for you and other users.
   Our survey can be found here: https://forms.gle/BMTbGQXLWSdn7vEs6

To permanently disable the survey prompt, run:
   skaffold config set --survey --global disable-prompt true`
)

// for testing
var (
	isStdOut = stdOut
)

type Runner struct {
	configFile string
}

func New(configFile string) *Runner {
	return &Runner{
		configFile: configFile,
	}
}

func DisplaySurveyPrompt(out io.Writer) {
	if isStdOut(out) {
		fmt.Fprintln(out, Prompt)
	}
}

func (s *Runner) DisplaySurveyForm(_ context.Context, out io.Writer) error {
	_, err := fmt.Fprintln(out, Form)
	if err != nil {
		return err
	}
	// Currently we will only update the global survey taken
	// When prompting for the survey, we need to use the same field.
	return config.UpdateGlobalSurveyTaken(s.configFile)
}

func stdOut(out io.Writer) bool {
	return out == os.Stdout
}
