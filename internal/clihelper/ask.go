package clihelper

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/tcnksm/go-input"
)

// yesNoValidator - waits for 'y'/'Y'/'n'/'N'
func yesNoValidator(rawString string) error {
	if s := strings.TrimSpace(strings.ToLower(rawString)); s == "y" || s == "n" {
		return nil
	}
	return fmt.Errorf("answer is not 'y' or 'n'")
}

// askYesNo - returns true if yes
func (h *Helper) askYesNo(question string) bool {
	hasAccountAnwser, err := h.ui.Ask(
		question,
		&input.Options{
			Required:     true,
			ValidateFunc: yesNoValidator,
			Loop:         true,
			HideOrder:    true,
		})
	if err != nil {
		h.out.Fatal(err)
	}
	return hasAccountAnwser == "y"
}

// askData - returns string from input
func (h *Helper) askData(question string) string {
	answer, err := h.ui.Ask(
		question,
		&input.Options{
			Required:  true,
			Loop:      true,
			HideOrder: true,
		},
	)
	if err != nil {
		h.out.Fatal(err)
	}
	return answer
}

// askPrivate - returns string from input with hiding input
func (h *Helper) askPrivate(question string) string {
	answer, err := h.ui.Ask(
		question,
		&input.Options{
			Required:  true,
			Loop:      true,
			HideOrder: true,
			Hide:      true,
		},
	)
	if err != nil {
		h.out.Fatal(err)
	}
	return answer
}

// askJSON - returns map. Not required input
func (h *Helper) askJSON(question string) (m map[string]interface{}) {
	_, err := h.ui.Ask(
		question,
		&input.Options{
			Default: "{}",
			Loop:    true,
			ValidateFunc: func(s string) error {
				return json.Unmarshal([]byte(s), &m)
			},
		},
	)
	if err != nil {
		h.out.Fatal(err)
	}
	return nil
}
