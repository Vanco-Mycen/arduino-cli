// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package configuration

import (
	"path/filepath"

	"github.com/spf13/viper"
)

func setDefaults(settings *viper.Viper, dataDir, userDir string) {
	// logging
	settings.SetDefault("logging.level", "info")
	settings.SetDefault("logging.format", "text")

	// Boards Manager
	settings.SetDefault("board_manager.additional_urls", []string{})

	// arduino directories
	settings.SetDefault("directories.Data", dataDir)
	settings.SetDefault("directories.Downloads", filepath.Join(dataDir, "staging"))
	settings.SetDefault("directories.User", userDir)

	// daemon settings
	settings.SetDefault("daemon.port", "50051")

	//telemetry settings
	settings.SetDefault("telemetry.enabled", true)
	settings.SetDefault("telemetry.addr", ":9090")
}
