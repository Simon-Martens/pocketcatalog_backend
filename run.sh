#!/bin/bash
CGO_ENABLED=1 go run  -tags "sqlite_fts5 sqlite_json sqlite_foreign_keys sqlite_vtable sqlite_math_functions" . serve --dev
