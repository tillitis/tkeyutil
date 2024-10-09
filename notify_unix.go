// SPDX-FileCopyrightText: 2023 Tillitis AB <tillitis.se>
// SPDX-License-Identifier: BSD-2-Clause

//go:build unix

package tkeyutil

import (
	"fmt"
	"os"

	"github.com/gen2brain/beeep"
)

func Notify(progname, msg string) {
	// Using progname as title
	if err := beeep.Notify(progname, msg, ""); err != nil {
		fmt.Fprintf(os.Stderr, "Notify message %q failed: %s\n", msg, err)
	}
}
