/*
 * This file is part of arduino-cli.
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
 *
 * Copyright 2017 BCMI LABS SA (http://www.arduino.cc/)
 */

package prettyPrints

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// actionOnItems pretty prints info about an action on one or more items.
func actionOnItems(itemPluralName string, actionPastParticiple string, itemOK []string, itemFails map[string]string) {
	if len(itemFails) > 0 {
		if len(itemOK) > 0 {
			logrus.Infof("The following %0s were succesfully %s:\n", itemPluralName, actionPastParticiple)
			logrus.Infoln(strings.Join(itemOK, " "))
			logrus.Info("However, t")
		} else { //UGLYYYY but it works
			logrus.Info("T")
		}
		logrus.Infof("he the following %s were not %s and failed :", itemPluralName, actionPastParticiple)
		for item, failure := range itemFails {
			logrus.Infof("%-10s -%s\n", item, failure)
		}
	} else {
		logrus.Infof("All %s successfully installed\n", itemPluralName)
	}
}
