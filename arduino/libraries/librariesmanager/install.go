/*
 * This file is part of arduino-cli.
 *
 * Copyright 2018 ARDUINO AG (http://www.arduino.cc/)
 *
 * arduino-cli is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 2 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, write to the Free Software
 * Foundation, Inc., 51 Franklin St, Fifth Floor, Boston, MA  02110-1301  USA
 *
 * As a special exception, you may use this file as part of a free software
 * library without restriction.  Specifically, if other files instantiate
 * templates or use macros or inline functions from this file, or you compile
 * this file and link it with other files to produce an executable, this
 * file does not by itself cause the resulting executable to be covered by
 * the GNU General Public License.  This exception does not however
 * invalidate any other reasons why the executable file might be covered by
 * the GNU General Public License.
 */

package librariesmanager

import (
	"fmt"

	"github.com/arduino/go-paths-helper"

	"github.com/bcmi-labs/arduino-cli/arduino/libraries"
	"github.com/bcmi-labs/arduino-cli/arduino/libraries/librariesindex"
	"github.com/bcmi-labs/arduino-cli/arduino/utils"
)

// Install installs a library and returns the installed path.
func (lm *LibrariesManager) Install(indexLibrary *librariesindex.Release) (*paths.Path, error) {
	if installedLibs, have := lm.Libraries[indexLibrary.Library.Name]; have {
		for _, installedLib := range installedLibs.Alternatives {
			if installedLib.Location != libraries.Sketchbook {
				continue
			}
			if installedLib.Version == indexLibrary.Version {
				return installedLib.Folder, fmt.Errorf("%s is already installed", indexLibrary.String())
			}
		}
	}

	libsDir := lm.getSketchbookLibrariesDir()
	if libsDir == nil {
		return nil, fmt.Errorf("sketchbook folder not set")
	}

	libPath := libsDir.Join(utils.SanitizeName(indexLibrary.Library.Name))
	return libPath, indexLibrary.Resource.Install(lm.DownloadsDir, libsDir, libPath)
}

func (lm *LibrariesManager) removeRelease(libName string, r *libraries.Library) error {
	libsDir := lm.getSketchbookLibrariesDir()
	if libsDir == nil {
		return fmt.Errorf("sketchbook folder not set")
	}

	libName = utils.SanitizeName(libName)
	return libsDir.Join(libName).RemoveAll()
}
