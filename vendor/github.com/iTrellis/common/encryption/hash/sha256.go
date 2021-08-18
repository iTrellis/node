/*
Copyright © 2016 Henry Huang <hhh@rutcode.com>

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/

package hash

import "crypto/sha256"

// NewSHA224 get SHA224 hash repo
func NewSHA224() Repo {
	return &defaultHash{
		Hash: sha256.New224(),
	}
}

// NewSHA256 get SHA256 hash repo
func NewSHA256() Repo {
	return &defaultHash{
		Hash: sha256.New(),
	}
}
