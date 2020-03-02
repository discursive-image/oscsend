# SPDX-FileCopyrightText: 2020 jecoz
#
# SPDX-License-Identifier: MIT

oscsend: main.go
	go build -o bin/$@ $^